package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/core-pershelf/cmd/constructor"
	"github.com/core-pershelf/cmd/starter"
	"github.com/core-pershelf/globals"
)

var (
	logFilePath = "/pershelf/log/service.log"
)

func Run() {
	// construct the server
	srv := constructor.ConstructServer()

	// Initialize log file
	initLogFile(logFilePath)

	// Initialize server config
	initServerConfig()

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

// initServerConfig reads the server config file and initializes the global variables
func initServerConfig() {
	// Read file: /pershelf/etc/server.json
	serverConfig, err := os.ReadFile("/pershelf/etc/server.json")
	if err != nil {
		log.Fatalf("(Error) : error reading the server config file : %v", err)
	}

	// Unmarshal the server config
	err = json.Unmarshal(serverConfig, &globals.ServerConf)
	if err != nil {
		log.Fatalf("(Error) : error unmarshalling the server config file : %v", err)
	}

	// Check fields
	if globals.ServerConf.Server.ServerIP == "" || globals.ServerConf.Server.ServerPort == "" || globals.ServerConf.Server.HelperPort == "" {
		log.Fatalf("(Error) : server config file is not valid")
	}
}
