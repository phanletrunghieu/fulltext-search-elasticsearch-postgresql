package domain

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Post .
type Post struct {
	gorm.Model
	Title   string
	Content string
}

// SearchByText .
func SearchByText(c context.Context, keyword string, offset, limit int) error {
	// elasticClient := elasticDB.GetDB()

	// esQuery := elastic.NewMultiMatchQuery(keyword, "title", "content").
	// 	Fuzziness("2").
	// 	MinimumShouldMatch("2")

	// result, err := elasticClient.Search().
	// 	Index(elasticDB.ElasticIndexName).
	// 	Query(esQuery).
	// 	From(offset).Size(limit).
	// 	Do(c)
	// if err != nil {
	// 	return err
	// }

	// pp.Println(result)

	return nil
}
