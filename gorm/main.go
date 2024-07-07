package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type User struct {
	gorm.Model
	FirstName sql.NullString
	LastName  string `gorm:"not null;unique"`
}

func (User) TableName() string {
	return "app_user"
}

func (u User) String() string {
	if u.FirstName.Valid {
		return fmt.Sprintf("%s %s", u.FirstName.String, u.LastName)
	}
	return u.LastName
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

	var user User
	db.First(&user)
	fmt.Printf("found first user: %v\n", user)

	var users []User
	result := db.Where("first_name = ?", "Dirk").Find(&users)
	fmt.Printf("found %d users with first name = 'Dirk'", result.RowsAffected)

	err = db.Transaction(func(tx *gorm.DB) error {
		newUser := User{FirstName: sql.NullString{String: "Dirk", Valid: true}, LastName: "Nowitzki"}
		return tx.Create(&newUser).Error
	})
	if err != nil {
		log.Println("transaction returned error:", err)
	}
}
