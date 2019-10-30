package model

import (
	"database/sql"
	"fmt"	
	_ "github.com/go-sql-driver/mysql"		
)

func ConnectDB(username string, password string, host string, database string) (*sql.DB, error) {
	// db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/golang_crud")

	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	db, err := sql.Open("mysql", conn)

	return db, err
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}