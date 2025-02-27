package parsers

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// ParseFlags parses the command-line flags and returns the path to the database configurations JSON file and the path to
// the log file.
func ParseFlags() (*string, *string, string) {
	// define command-line flags serverConf, dbConf, logFile
	serverConfFile := flag.String("c", "", "Path to the server configurations JSON file")
	dbConfFile := flag.String("d", "", "Path to the database configurations JSON file")
	logFile := flag.String("l", "", "Path to the log file")
	flag.Parse()

	// check if the server config JSON file flag is provided
	if *serverConfFile == "" {
		log.Println(fmt.Sprintf("Usage: %s -c <server_conf_json_file> -d <db_conf_json_file> -l <log_file>", os.Args[0]))
		return nil, nil, ""
	}

	// check if the database config JSON file flag is provided
	if *dbConfFile == "" {
		log.Println(fmt.Sprintf("Usage: %s -c <server_conf_json_file> -d <db_conf_json_file> -l <log_file>", os.Args[0]))
		return nil, nil, ""

	}

	// check if the log file flag is provided
	if *logFile == "" {
		log.Println(fmt.Sprintf("Usage: %s -c <server_conf_json_file> -d <db_conf_json_file> -l <log_file>", os.Args[0]))
		return nil, nil, ""
	}

	return serverConfFile, dbConfFile, *logFile
}
