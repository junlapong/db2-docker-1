package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	// dsn := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;PROTOCOL=TCPIP;UID=db2inst1;PWD=pa55w0rd"
	dsn := os.Getenv("DB2DSN")
	db, err := sql.Open("go_ibm_db", dsn)
	defer db.Close()

	if err != nil {
		fmt.Println("error !!!")
		fmt.Println(err)
	}

	fmt.Println(display(db))
	fmt.Printf("~~~ go-db2: %s\n", dsn)
}

func display(db *sql.DB) error {
	st, err := db.Prepare("SELECT current date FROM sysibm.sysdummy1")
	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	cols, _ := rows.Columns()
	fmt.Printf("%v\n", cols)
	fmt.Println("-------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var d string
		err = rows.Scan(&d)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", d)
	}
	return nil
}
