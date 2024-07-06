package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:password@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to start transaction: %v\n", err)
		os.Exit(1)
	}

	var id int64
	var firstName, lastName string
	err = conn.QueryRow(context.Background(), "select id, first_name, last_name from app_user").Scan(&id, &firstName, &lastName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "query row failed: %v\n", err)
		tx.Rollback(context.Background())
		os.Exit(1)
	}
	log.Printf("found user %d: %s %s\n", id, firstName, lastName)

	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to commit: %v\n", err)
		os.Exit(1)
	}
}
