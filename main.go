package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-entgo/ent"
	"log"
	"time"
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
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("connect success", client)

	Do(ctx, client)

}

func Do(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Tom").
		Save(ctx)
	if err != nil {
		return err
	}

	card1, err := client.Card.
		Create().
		SetOwner(a8m).
		SetNumber("1020").
		SetExpired(time.Now().Add(time.Minute)).
		Save(ctx)

	log.Println("card:", card1)
	if err != nil {
		return err
	}

	card2, err := a8m.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	log.Println("card:", card2)

	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	log.Println("owner:", owner)

	return nil
}
