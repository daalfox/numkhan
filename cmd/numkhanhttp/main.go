package main

import (
	"net/http"

	"github.com/daalfox/numkhan/internal/numkhan"
)

func main() {
	dbName := "numkhandb.db"
	db, _ := numkhan.SetupDb(dbName)

	s := numkhan.NewHttpService(db)

	http.ListenAndServe(":8080", s.Router)
}
