package main

import (
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-entgo/ent"
)

type (
	Database struct {
		username string
		password string
		hostname string
		dbname   string
		port     string
	}
)

func (db *Database) GetDSN() string {
	const format = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	return fmt.Sprintf(format, db.username, db.password, db.hostname, db.port, db.dbname)
}

func main() {

	db := Database{
		username: "test",
		password: "test",
		hostname: "localhost",
		dbname:   "test",
		port:     "33016",
	}

	client, err := ent.Open(dialect.MySQL, db.GetDSN(), ent.Debug())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect success", client)

}
