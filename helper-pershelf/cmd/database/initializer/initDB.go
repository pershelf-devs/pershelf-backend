package initializer

/**
* InitializeDatabase initializes the database.
* This function performs the following operations:
* - Initializes the database tables
* - Initializes the triggers
* - Initializes the constraints
*
* @return error Returns an error if the initialization process fails
 */
func InitializeDatabase() error {
	if err := InitializeTables(); err != nil {
		return err
	}
	if err := InitializeTriggers(); err != nil {
		return err
	}
	if err := InitializeConstraints(); err != nil {
		return err
	}
	return nil
}
