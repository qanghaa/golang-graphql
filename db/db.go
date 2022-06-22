// db.go

package db

import (
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Database: "gobooks",
		Password: "AElamot2",
	})
	_, DBStatus := db.Exec("SELECT 1")
	if DBStatus != nil {
		panic("PostgreSQL is down")
	}
	return db
}
