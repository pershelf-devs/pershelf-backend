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

	// books
	if err := globals.PershelfDB.AutoMigrate(&crud.Book{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	// reviews
	if err := globals.PershelfDB.AutoMigrate(&crud.Review{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	// comments
	if err := globals.PershelfDB.AutoMigrate(&crud.Comment{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	// user_shelves
	if err := globals.PershelfDB.AutoMigrate(&crud.UserShelf{}); err != nil {
		globals.Log("Error initializing tables: ", err)
		return err
	}

	// user_books
	if err := globals.PershelfDB.AutoMigrate(&crud.UserBook{}); err != nil {
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
