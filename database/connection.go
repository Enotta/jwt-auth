package database

import (
	"main/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Local postgres connection
func Connect() *gorm.DB {
	var dsn string = "host=localhost user=postgres password=123 dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var tables int
	db.Raw(`SELECT count(table_name)
  		    FROM information_schema.tables
 			WHERE table_schema='public'
   			AND table_type='BASE TABLE';`).Scan(&tables)

	if tables == 0 {
		db.AutoMigrate(&model.User{})
	}

	return db
}
