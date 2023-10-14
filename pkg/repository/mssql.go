package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"	
)


type Config struct {
	Host     string
	Instance string
    Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewMSsqlDB(cfg Config) (*sqlx.DB, error) {

	connStr:=fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode);

	db, err := sqlx.Open("mssql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
