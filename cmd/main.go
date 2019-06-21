package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/pg"
	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/domain"
)

func main() {
	closeDB := pg.Init()
	defer closeDB()

	db := pg.GetDB()
	db.AutoMigrate(&domain.Post{})
}
