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

	var id int64
	var firstName, lastName string
	err = conn.QueryRow(context.Background(), "select id, first_name, last_name from app_user").Scan(&id, &firstName, &lastName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "query row failed: %v\n", err)
		os.Exit(1)
	}
	log.Printf("found user %d: %s %s\n", id, firstName, lastName)
}
