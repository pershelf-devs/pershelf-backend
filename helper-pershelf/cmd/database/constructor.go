package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/pershelf/pershelf/config/database"
	"github.com/pershelf/pershelf/globals"
)

// createDatabasePershelfIfNotExists creates the pershelf database if it doesn't exist
func createDatabasePershelfIfNotExists(dbConf config.DBConnectionConfig) error {
	// open connection
	dsn := fmt.Sprintf("%s:%s@%s/", dbConf.Username, dbConf.Password, dbConf.Network)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		globals.Log(err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		globals.Log(err)
		return err
	}

	defer sqlDB.Close()

	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbConf.DbName + ";").Error
	if err != nil {
		globals.Log(err)
		return err
	}

	err = db.Exec("ALTER DATABASE " + dbConf.DbName + " CHARACTER SET utf8mb4 COLLATE utf8mb4_turkish_ci;").Error
	if err != nil {
		globals.Log(err)
		return err
	}

	globals.Log("pershelf DB " + dbConf.DbName + " created/exists")

	return nil
}

// initPershelfDBConnection initializes the pershelf database connection (checks if it exists and creates it if it does not)
func initPershelfDBConnection(dbConf config.DBConnectionConfig) (*gorm.DB, error) {
	createDatabasePershelfIfNotExists(dbConf)

	// open connection
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", dbConf.Username, dbConf.Password, dbConf.Network, dbConf.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		globals.Log(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		globals.Log(err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		globals.Log("pershelf DB conn not ok: " + dbConf.Network)
	} else {
		globals.Log("pershelf DB conn ok: " + dbConf.Network)
	}

	return db, nil
}

// DBHandler creates (if not exists), initializes connection to and sets up the pershelf DB for use.
func DBHandler(dbConf config.DBConnectionConfig, logPath string) error {
	// initialize log file
	initLogFile(logPath)

	// init DB
	var err error
	globals.PershelfDB, err = initPershelfDBConnection(dbConf)
	if err != nil {
		globals.Log("(Error) : error connecting to the pershelf Database : ", err)
		return err
	}

	sqlDB, err := globals.PershelfDB.DB()
	if err != nil {
		globals.Log("(Error) : error getting the pershelf Database connection : ", err)
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	globals.Log("pershelf DB connection initialized")

	return nil
}

// initLogFile redirects log output to file
func initLogFile(logFilePath string) {
	LogFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening log file: ", err)
	}
	log.SetOutput(LogFile)
}
