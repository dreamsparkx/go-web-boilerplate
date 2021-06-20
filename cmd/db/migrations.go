package db

import (
	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(app *config.AppConfig) error {
	dbMigrations := &migrate.FileMigrationSource{Dir: "cmd/db/migrations"}
	if _, err := dbMigrations.FindMigrations(); err != nil {
		return err
	}
	nOut, errOut := migrate.Exec(app.DB.DB, "postgres", dbMigrations, migrate.Up)
	config.AppLogger.Infof("Applied %d migrations to the DB", nOut)
	if errOut != nil {
		config.AppLogger.Info("Error running migrations on the DB")
		return errOut
	}
	return nil
}
