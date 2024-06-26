package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	DBHost     = "localhost"
	DBPort     = 8000
	DBUser     = "postgres"
	DBPassword = "Rifki123"
	DBName     = "postgres"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName))
	if err != nil {
		fmt.Println("error 1")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error 2")
		panic(err)
	}
	fmt.Println("Successfully connected to Database!")
	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
