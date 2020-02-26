package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/hukurou-s/user-auth-api-sample/interfaces/database"
)

func NewSqlHandler() database.SqlHandler {
	db, err := gorm.Open("postgres", "user=LEO dbname=user-auth-sample-db password='' sslmode=disable")

	if err != nil {
		panic(err)
	}
	return db
}
