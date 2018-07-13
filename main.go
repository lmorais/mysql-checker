package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var (
	user     = flag.String("user", "", "The database user name")
	password = flag.String("password", "", "The database password")
	db       = flag.String("database", "", "The database to connect to")
	query    = flag.String("query", "SELECT name FROM user", "The test query")
	addr     = flag.String("address", "localhost:8080", "The address to listen on")
)

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", *user, *password, *db))

	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}

	defer db.Close()

	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		_, err := db.Exec(*query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("{\"status\": \"OK\"}"))
		return
	})

	http.ListenAndServe(*addr, nil)
}
