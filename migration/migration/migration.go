package migration

import (
	"flag"
	"fmt"
	"os"

	// "github.com/local/migration/migration"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type Config struct {
}

func InitConfig() (*Config, error) {
	return nil, nil
}

type (
	Migration struct {
		db     *pg.DB
		Config *Config
		logger logrus.FieldLogger
	}
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func New() (app *Migration, err error) {
	app = &Migration{}

	// app.logger = util.Log

	app.db, err = newDb()
	if err != nil {
		return nil, err
	}
	// app.Database.Logger = app.logger

	return app, err
}

func (a *Migration) Migration(command string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir = fmt.Sprintf("%s/migration", dir)
	migrations.DefaultCollection.DiscoverSQLMigrations(dir)
	migrations.DefaultCollection.DisableSQLAutodiscover(true)
	fmt.Println("=== a.Database.Source:: ")
	commands := []string{command}
	fmt.Println("=== commands:: ", commands)

	oldVersion, newVersion, err := migrations.Run(a.db, commands...)
	if err != nil {
		return err
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
	return nil
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

// func (a *Migration) Close() error {
// 	return a.Database.Close()
// }
