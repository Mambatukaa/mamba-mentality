package migrate

// ApplyMigrations applies migrations to the specified database connection.
func ApplyMigrations(connString, migrationsPath string) error {
}

// RollbackMigrations rolls back the last applied migration from the specified database connection.
func RollbackMigrations(connString, migrationsPath string) error {
}
