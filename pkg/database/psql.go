package database

import (
	"database/sql"
	"fmt"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	UserName string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.UserName, info.DBName, info.SSLMode, info.Password))
	if err != nil {
		return db, err
	}
	return db, nil
}
