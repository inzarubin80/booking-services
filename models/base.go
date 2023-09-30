package models
import (
	"database/sql"
	"os"
	"log"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"	
)

var db *sql.DB


func init() {

	var err error
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	connString := fmt.Sprintf("server=%s; database=%s; user Id=%s;  password=%s; ", dbHost, dbName, username, password)
	
	fmt.Println("Строка соединения: " + connString)

	db, err = sql.Open("sqlserver", connString)
    
	if err != nil {
        log.Fatal("Error creating connection pool: ", err.Error())
    }    
  
}

func GetDB () *sql.DB{
	return db
}
