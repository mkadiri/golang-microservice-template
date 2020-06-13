package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB

func init() {
	fmt.Println("--- Init database")

	var err error
	Db, err = sql.Open("mysql", getDatabaseSourceName())

	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}

	// have a read
	// https://www.alexedwards.net/blog/configuring-
	// https://github.com/go-sql-driver/mysql/issues/461
	Db.SetMaxIdleConns(0)
	Db.SetMaxOpenConns(1)
	Db.SetConnMaxLifetime(time.Minute * 5)

	err = Db.Ping()

	if err != nil {
		_ = Db.Close()
		log.Fatal("Unable to ping database, likely your config is incorrect")
	}
}

func getDatabaseSourceName() string {
	config := map[string]string{
		"MYSQL_HOST": os.Getenv("MYSQL_HOST"),
		"MYSQL_PORT": os.Getenv("MYSQL_PORT"),
		"MYSQL_USER": os.Getenv("MYSQL_USER"),
		"MYSQL_PASSWORD": os.Getenv("MYSQL_PASSWORD"),
		"MYSQL_DATABASE": os.Getenv("MYSQL_DATABASE"),
	}

	for key, value := range config {
		if len(value) == 0 {
			log.Fatal("[" + key +"] needs to be set")
		}
	}

	return config["MYSQL_USER"] + ":" + config["MYSQL_PASSWORD"] + "@tcp(" + config["MYSQL_HOST"] + ":" +
		config["MYSQL_PORT"] + ")/" + config["MYSQL_DATABASE"] + "?parseTime=true"
}