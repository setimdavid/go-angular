package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"os"
)

func main() {
	pass := os.Getenv("DB_PASS")
	db, err := gorm.Open(
		"postgres",
		"host=my-postgres-db user=stimbeam_apps password="+pass+" dbname=go sslmode=disable")
		// "host=172.19.0.2 port=5433 user=stimbeam_apps password="+pass+" dbname=go sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db: db,
		r:  mux.NewRouter(),
	}
	app.start()
}
