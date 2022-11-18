package elasticsearch

import (
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchClient struct{}


func (ctrl ElasticsearchClient) CreateIndex() {
	// ここの設定の見直しを行う
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	index := "test"
	mapping := `
	{
		"settings": {
			"number_of_shards": 1
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "text"
				}
			}
		}
	}`
	res, err := es.Indices.Create(
			index,
			es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
			log.Fatal(err)
	}
	fmt.Printf("\"更新前です\": %v\n", "ああ")
	log.Println(res)
}
