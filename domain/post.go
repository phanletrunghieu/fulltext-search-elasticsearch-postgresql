package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"

	elasticDB "github.com/phanletrunghieu/fulltext-search-elasticsearch-postgresql/config/database/elastic"
)

// Post .
type Post struct {
	gorm.Model
	Title   string
	Content string
}

// SearchByText .
func SearchByText(c context.Context, keyword string, offset, limit int) ([]*Post, uint, uint, error) {
	elasticClient := elasticDB.GetDB()

	query := `
	{
		"query": {
		  "multi_match": {
			"query": "` + keyword + `",
			"fuzziness": "2",
			"minimum_should_match": "2",
			"fields": [
			  "title",
			  "content"
			]
		  }
		}
	}`

	buff := bytes.NewBufferString(query)

	res, err := elasticClient.Search(
		elasticClient.Search.WithContext(context.Background()),
		elasticClient.Search.WithIndex("posts"),
		elasticClient.Search.WithBody(buff),
		elasticClient.Search.WithTrackTotalHits(true),
		elasticClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, 0, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var err error
		var e map[string]interface{}
		if er := json.NewDecoder(res.Body).Decode(&e); er != nil {
			err = er
		} else {
			err = fmt.Errorf("%s", e["error"].(map[string]interface{})["reason"])
		}
		return nil, 0, 0, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, 0, err
	}

	total := uint(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	took := uint(result["took"].(float64))

	posts := []*Post{}
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		post := &Post{}

		source := hit.(map[string]interface{})["_source"]
		b, _ := json.Marshal(source)
		json.Unmarshal(b, post)

		posts = append(posts, post)
	}

	return posts, took, total, nil
}
