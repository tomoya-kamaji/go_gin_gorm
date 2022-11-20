package elasticsearch

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

type ElasticSearchAdapter struct {
	client *elastic.Client
}

func NewElasticSearchAdapter() *ElasticSearchAdapter {
	es, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return &ElasticSearchAdapter{
		client: es,
	}
}

func (adapter ElasticSearchAdapter) CreateIndex() {
	adapter.client.DeleteIndex("user").Do(context.Background())
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"doc":{
				"properties":{
					"name":{
						"type":"keyword"
					},
					"address":{
						"type":"text",
						"store": true,
						"fielddata": true
					},
					"age":{
							"type":"long"
					}
				}
			}
		}
	}`
	createIndex, err := adapter.client.CreateIndex("user").Body(mapping).IncludeTypeName(true).Do(context.Background())
	if err != nil {
		fmt.Println("createIndexに失敗しました")
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	}
}
