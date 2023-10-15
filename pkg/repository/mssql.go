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

func NewMSsqlDB(cfg Config) (*sqlx.DB, error)  {

	connString := fmt.Sprintf("server=%s; database=%s",
	cfg.Host, cfg.DBName)

	//connString = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
      //  cfg.Host, cfg.Username, cfg.Password, cfg.DBName)

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
