package internal

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-orm/config"
	"log"
)

func makeMigrate() *migrate.Migrate {
	db, err := sql.Open(config.GetDatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//entMigrationPath := "file://ent_example/migrations"
	sqlcMigrationPath := "file://sqlc_example/db/migrations"

	m, err := migrate.NewWithDatabaseInstance(
		sqlcMigrationPath,
		"postgres", driver)

	return m
}

func UpMigration() {
	log.Println("Start migrate up")
	m := makeMigrate()
	err := m.Up()
	LogFatal(err)
	log.Println("Complete migrate up")
}

func DownMigration() {
	log.Println("Start migrate down")
	m := makeMigrate()
	err := m.Down()
	LogFatal(err)
	log.Println("Complete migrate down")
}
