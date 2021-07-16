package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	// con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	con := os.Getenv("DB2DSN")
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
