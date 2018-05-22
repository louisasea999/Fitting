package database

import (
    "database/sql"
	"log"
    _ "github.com/lib/pq"
)

var Pdb *sql.DB
var err error

func init() {
 
 Pdb, err = sql.Open("postgres", "port=5432 user=postgres password=Aabc123_* dbname=Fitting sslmode=disable")
 
 if err != nil {
  log.Fatal(err.Error())
 }
 
 err =Pdb.Ping()
 
 if err != nil {
  log.Fatal(err.Error())
 }
}
