package database

import (
	elasticsearch "github.com/elastic/go-elasticsearch"
	"github.com/jinzhu/gorm"

	elasticDB "github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/elastic"
	"github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/pg"
)

// GetPGDB .
func GetPGDB() *gorm.DB {
	return pg.GetDB()
}

// GetElastic .
func GetElastic() *elasticsearch.Client {
	return elasticDB.GetDB()
}
