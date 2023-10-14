package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"	
)


type Config struct {
	Host     string
    Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewMSsqlDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mssql", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, 
		cfg.Port, 
		cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
