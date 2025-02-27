package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pershelf/pershelf/cmd/database"
	"github.com/pershelf/pershelf/cmd/database/initializer"
	"github.com/pershelf/pershelf/cmd/server"
	"github.com/pershelf/pershelf/config"
	"github.com/pershelf/pershelf/config/server/parsers"
)

func Run() {
	// parse the server paths (configurations & logging)
	serverConfFilePath, dbConfFilePath, logPath := parsers.ParseFlags()

	if logPath == "" {
		log.Fatalf("(Error) : error parsing configuration path. Ensure the configuration file path is specified.")
	}

	var confStruct config.Config

	// Read the server configuration from the configuration file
	confByte, err := os.ReadFile(*serverConfFilePath)
	if err != nil {
		log.Fatalf("Error reading service configuration file: %v", err)
	}

	if err = json.Unmarshal(confByte, &confStruct); err != nil {
		log.Fatalf("Error unmarshalling pershelf server configuration: %v", err)
	}

	// Read the database configuration from the auditdb configuration file
	dbConfByte, err := os.ReadFile(*dbConfFilePath)
	if err != nil {
		log.Fatalf("Error reading pershelf database configuration file: %v", err)
	}

	// Unmarshal the configuration
	if err = json.Unmarshal(dbConfByte, &confStruct); err != nil {
		log.Fatalf("Error unmarshalling pershelf database configuration: %v", err)
	}

	// initialize log file
	initLogFile(logPath)

	// Log the configuration
	log.Printf("pershelf configuration: %+v", confStruct)

	// Initialize the database connection
	if err = database.DBHandler(confStruct.Conn, logPath); err != nil {
		log.Printf("Error initializing pershelf database connection: %v", err)
	}

	// Initialize the database
	if err := initializer.InitializeDatabase(); err != nil {
		log.Printf("Error initializing database: %v", err)
	}

	// run the database server
	server.RunDBHttpServer(confStruct.Srv)
}

// initLogFile redirects log output to file
func initLogFile(logFilePath string) {
	LogFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("(Error) : error opening the log file : %v", err)
	}
	log.SetOutput(LogFile)
	log.Printf("Log output is set to %s", logFilePath)
}
