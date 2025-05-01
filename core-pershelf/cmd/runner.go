package cmd

import (
	"log"
	"os"

	"github.com/core-pershelf/cmd/constructor"
	"github.com/core-pershelf/cmd/starter"
	"github.com/core-pershelf/globals"
	"github.com/core-pershelf/mongo"
)

var (
	logFilePath = "/pershelf/log/service.log"
)

func Run() {
	// construct the server
	srv := constructor.ConstructServer()

	// Initialize log file
	initLogFile(logFilePath)

	// connect to mongo
	err := mongo.ConnectMongoDB()
	if err != nil {
		log.Fatalf("(Error) : error connecting to mongo : %v", err)
	}

	// Initialize MongoDB collections
	globals.InitCollections()

	// start the server
	starter.StartServer(srv)
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
