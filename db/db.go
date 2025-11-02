package db

import (
	"database/sql"
	"os"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB
	
func Init() error {
	_ = godotenv.Load()
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		return errors.New("MYSQL_DSN environment variable not set")
	}
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	return nil 
}
