package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type User struct {
	gorm.Model
	FirstName sql.NullString
	LastName  string `gorm:"not null"`
}

func (User) TableName() string {
	return "app_user"
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to migrate table 'app_user': %v\n", err)
		os.Exit(1)
	}
}
