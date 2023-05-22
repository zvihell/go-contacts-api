package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func InitDB(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLmode))
	if err != nil {
		panic(err.Error())
	}
	return db
}
