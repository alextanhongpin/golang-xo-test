package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alextanhongpin/go-xo-test/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	p := models.Person{
		ID:    uuid.New(), // UUID must be provided.
		Name:  "John Doe",
		Email: "john.doe@mail.com",
	}

	// The query generated does not return the id for postgres driver.
	if err := p.Insert(db); err != nil && !errors.Is(sql.ErrNoRows, err) {
		log.Printf("%+v\n", p)
		log.Fatal("insertError:", err)
	}
	log.Println("Inserted user")

	u, err := models.PersonByEmail(db, "john.doe@mail.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", u)
}
