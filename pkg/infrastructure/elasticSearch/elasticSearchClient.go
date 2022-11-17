package elasticsearch

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchClient struct{}

func (ctrl ElasticsearchClient) Update() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
