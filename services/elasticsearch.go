package services

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

type ESClient struct {
	*elastic.Client
}

func ConnectToESClient() (*ESClient, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))
	if err != nil {
		return nil, err
	}

	return &ESClient{client}, nil
}

func (es *ESClient) SearchClient(out interface{}, key, value string) (interface{}, error) {
	ctx := context.Background()
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	data := []interface{}{}

	searchService := es.Search().Index().SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		err = json.Unmarshal(hit.Source, &out)
		if err != nil {
			return nil, err
		}
		data = append(data, out)

	}

	return data, nil
}
