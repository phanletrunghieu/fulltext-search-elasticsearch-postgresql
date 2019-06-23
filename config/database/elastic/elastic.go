package elastic

import (
	"fmt"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch"
)

var elasticClient *elasticsearch.Client

const ElasticIndexName = "blogs"

const (
	dbHost = "localhost"
	dbPort = 9200
)

// Init .
func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%d", dbHost, dbPort),
		},
	}

	var err error
	elasticClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panicf("Error creating the client: %s", err)
	}

	res, err := elasticClient.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}

	log.Println(res)
}

// GetDB .
func GetDB() *elasticsearch.Client {
	return elasticClient
}
