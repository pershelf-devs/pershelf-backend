package initializer

import (
	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/globals"
)

// This file contains the functions to initialize the database tables
func InitializeTables() error {
	// users
	if err := globals.PershelfDB.AutoMigrate(&crud.User{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	// refresh_tokens
	if err := globals.PershelfDB.AutoMigrate(&crud.RefreshToken{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	return nil
}
