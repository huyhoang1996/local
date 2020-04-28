package migration

import (
	"fmt"
	"time"

	"github.com/gas/gas-tools/util"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Database struct {
	Source      *pg.DB
	Destination *pg.DB
	Logger      logrus.FieldLogger
}

func newDb() (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "postgres",
		Database:              "migration",
		Addr:                  fmt.Sprintf("%s:%d", "192.168.3.212", 5433),
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
		OnConnect: func(conn *pg.Conn) error {
			// zone, _ := time.Now().Zone()
			var err error
			// if len("") > 0 {
			// 	_, err = conn.Exec("set search_path = ?; set timezone = ?", config.Schema, zone)
			// } else {
			// 	_, err = conn.Exec("set timezone = ?", zone)
			// }
			if err != nil {
				return errors.Wrap(err, "unable to connect to database")
			}
			return nil
		},
	})
	// db.AddQueryHook(dbLogger{})
	return db, nil
}

// func NewSeed(src *Config, des *Config) (*Database, error) {
// 	source, err := new(src)
// 	if err != nil {
// 		return nil, err
// 	}
// 	destination, err := new(des)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Database{Source: source, Destination: destination}, nil
// }

// func NewDb(config *Config) (*Database, error) {
// 	db, err := new(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("=== db:: ", db)
// 	return &Database{Source: db}, nil
// }

func (db *Database) Close() error {
	if db.Source != nil {
		err := db.Source.Close()
		if err != nil {
			return err
		}
	}
	if db.Destination != nil {
		err := db.Destination.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	util.Log.Debug(q.FormattedQuery())
}
