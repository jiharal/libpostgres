package libpostgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate is used to migrate table
// Ex: Migrate(opt, "file://db/migrations")
func Migrate(db *sql.DB, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}
	return nil
}

func DownOneStep(db *sql.DB, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	err = m.Steps(-1)
	if err != nil {
		return err
	}
	return nil
}
