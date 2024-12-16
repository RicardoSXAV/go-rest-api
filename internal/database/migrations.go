package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // migrations source path
		"mysql", // database name
		driver, // database instance
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	return nil
}

func GetMigrationVersion(db *sql.DB) (uint, bool, error) {
	var version uint
	var dirty bool

	row := db.QueryRow("SELECT version, dirty FROM schema_migrations LIMIT 1")
	err := row.Scan(&version, &dirty)
	if err != nil {
		return 0, false, fmt.Errorf("could not get migration version: %w", err)
	}

	return version, dirty, nil
}