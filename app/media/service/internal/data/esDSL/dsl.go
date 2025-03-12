package esDSL

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

// 使用这个工具类需要注意，可以适配任何类型
// 但是传进来的参数必须和Elasticsearch中的对应

func QueryFuzzy[T any](es *elasticsearch.Client, index, title, value string) ([]T, error) {
	// 创建查询语句
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"fuzzy": map[string]interface{}{
				title: map[string]interface{}{
					"value":     value,
					"fuzziness": "AUTO",
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.New(fmt.Sprintf("QueryFuzzy json.Marshal doc error: %s", err))
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil || res.IsError() {
		return nil, errors.New(fmt.Sprintf("QueryFuzzy es.Search doc error: %s", err))
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// todo: 要求写入到传入的日志
	log.Println(res)

	// 解析查询结果
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.New(fmt.Sprintf("QueryFuzzy json.NewDecoder(res.Body).Decode(&r) error: %s", err))
	}

	models := make([]T, 0)

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		if _, v := hit.(map[string]interface{})["_source"]; v {
			var model T
			body, err := json.Marshal(hit.(map[string]interface{})["_source"])
			if err != nil {
				log.Println(err)
				continue
			}
			if err := json.Unmarshal(body, &model); err != nil {
				log.Println(err)
				continue
			}
			models = append(models, model)
		}
	}
	return models, nil
}

func PutDoc[T any](es *elasticsearch.Client, doc T, index string, id string) error {
	marshal, err := json.Marshal(doc)
	if err != nil {
		return errors.New(fmt.Sprintf("PutDoc json.Marshal doc error: %s", err))
	}
	indexReq := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(marshal),
		Refresh:    "true",
	}
	_, err = indexReq.Do(context.Background(), es)
	if err != nil {
		return errors.New(fmt.Sprintf("PutDoc indexReq.Do error: %s", err))
	}
	return nil
}
