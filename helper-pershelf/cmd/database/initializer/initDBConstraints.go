package initializer

import "github.com/pershelf/pershelf/globals"

// This file contains the functions to initialize the database constraints
func InitializeConstraints() error {
	if err := dropConstraints(); err != nil {
		globals.Log("Error dropping constraints:", err)
	}

	// user_book_relations
	if err := globals.PershelfDB.Exec(`
		ALTER TABLE user_book_relations
		ADD CONSTRAINT unique_user_book_relation UNIQUE (user_id, book_id),
		ADD CONSTRAINT fk_user_book_relations_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		ADD CONSTRAINT fk_user_book_relations_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE;
	`).Error; err != nil {
		globals.Log("Error adding constraints to user_book_relations table:", err)
	}

	return nil
}

func dropConstraints() error {
	// user_book_relations
	if err := globals.PershelfDB.Exec(`
		ALTER TABLE user_book_relations
		DROP CONSTRAINT unique_user_book_relation,
		DROP CONSTRAINT fk_user_book_relations_user,
		DROP CONSTRAINT fk_user_book_relations_book;
	`).Error; err != nil {
		globals.Log("Error dropping constraints from user_book_relations table:", err)
	}

	return nil
}
