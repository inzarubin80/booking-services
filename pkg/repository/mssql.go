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

	//connStr:=fmt.Sprintf("host=%s  user=%s dbname=%s password=%s sslmode=%s",
	//cfg.Host, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode);


	connStr := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Instance, cfg.DBName)


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
