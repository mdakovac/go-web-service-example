package main

import (
	"context"
	"db/db"
	"fmt"
	"log"
	"util/env_vars"
)

func main(){
	conn := db.Connect(env_vars.EnvVariables.DATABASE_URL)
	defer db.Disconnect(conn)

	// Execute the SQL setup script
    _, err := conn.Exec(context.Background(), `
        DROP TABLE IF EXISTS album;
        CREATE TABLE album (
          id         SERIAL PRIMARY KEY,
          title      VARCHAR(128) NOT NULL,
          artist     VARCHAR(255) NOT NULL,
          price      DECIMAL(5,2) NOT NULL
        );

        INSERT INTO album
          (title, artist, price)
        VALUES
          ('Blue Train', 'John Coltrane', 56.99),
          ('Giant Steps', 'John Coltrane', 63.99),
          ('Jeru', 'Gerry Mulligan', 17.99),
          ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
    `)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("SQL setup script executed successfully")
}