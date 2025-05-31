package globals

import "log"

// Log logs stuff
func Log(msg ...interface{}) {
	log.Println(msg...)
}