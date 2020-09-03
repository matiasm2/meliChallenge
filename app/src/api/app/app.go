package app

import (
	"api/app/items"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// Needed to mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	r *gin.Engine
)

const (
	port string = ":8080"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	items.Configure(r, db)
	r.Run(port)
}

func configDataBase() *sql.DB {
	MYSQLUSER := os.Getenv("MYSQL_USER")
	MYSQLPASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQLDB := os.Getenv("MYSQL_DATABASE")
	DBHOST := "db"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		MYSQLUSER,
		MYSQLPASSWORD,
		DBHOST,
		MYSQLDB))
	if err != nil {
		panic("Could not connect to the db")
	}

	for {
		err := db.Ping()
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		// This is bad practice... You should create a schema.sql with all the definitions
		createTable(db)
		return db
	}

}

func createTable(db *sql.DB) {
	// create table if not exists
	sqlTable := `
	CREATE TABLE IF NOT EXISTS items(
		id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		name TEXT,
		description TEXT
	);`
	_, err := db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}
