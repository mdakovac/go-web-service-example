package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type DbConnectionType = *pgx.Conn

func Connect(DATABASE_URL string) DbConnectionType {
	fmt.Println("Connecting to DB with connection string:", DATABASE_URL)
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func Disconnect(c DbConnectionType) {
	c.Close(context.Background())
}
