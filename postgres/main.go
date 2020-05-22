package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-pg/pg"
)

func main() {
	type DBConfig struct {
		user        string
		password    string
		database    string
		addr        string
		search_path string
	}

	env := os.Getenv("ENV")
	dbconfig := DBConfig{"postgres", "huyhoang@123", "media",
		fmt.Sprintf("%s:%d", "192.168.3.212", 5433),
		"media"}
	if env == "PRODUCTION" {
		dbconfig = DBConfig{"postgres", "huyhoang@123", "media", fmt.Sprintf("%s:%d", "127.0.0.1", 5433), "media"}
	}

	db := pg.Connect(&pg.Options{
		User:                  dbconfig.user,
		Password:              dbconfig.password,
		Database:              dbconfig.database,
		Addr:                  dbconfig.addr,
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * 6000,
		OnConnect: func(conn *pg.Conn) error {
			zone, _ := time.Now().Zone()
			_, err := conn.Exec("set search_path = ?; set timezone = ?", dbconfig.search_path, zone)
			if err != nil {
				fmt.Println("Connect Fail")
				fmt.Println("ERR:: ", err)
				return err

			}
			fmt.Println("Connect success")
			return nil
		},
	})

	type media struct {
		Id       int
		Name     string
		TypeFile string
		IsPublic bool
	}
	// image := &media{Name: "huy", TypeFile: "type", IsPublic: false}
	// err := db.Insert(image)

	image := &media{
		Id:   1,
		Name: "huy hoang",
	}
	err := db.Update(image)
	fmt.Println("=== ERR ", err)
	defer db.Close()

}
