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

	//connStr:=fmt.Sprintf("host=%s  user=%s dbname=%s password=%s sslmode=%s database=%s",

	connString := fmt.Sprintf("mssql://%s:%s@%s/%s", cfg.Username,cfg.Password, cfg.Host, cfg.DBName)

	db, err := sqlx.Open("mssql", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
