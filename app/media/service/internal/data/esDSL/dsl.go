package esDSL

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io"
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
	if err != nil {
		return nil, fmt.Errorf("QueryFuzzy es.Search error: %w", err)
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("QueryFuzzy es.Search response error: %s", body)
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

func QueryByCategoryMustViewsDESCLimit[T any](es *elasticsearch.Client, index string, page, size int, queryItems ...interface{}) ([]T, error) {
	if len(queryItems)%2 != 0 {
		return nil, errors.New("query items must (key value),but it lost")
	}
	must := make([]map[string]interface{}, 0)
	for i := 0; i < len(queryItems); i += 2 {
		key, ok := queryItems[i].(string)
		if !ok {
			return nil, errors.New("query key must be string")
		}
		must = append(must, map[string]interface{}{"match": map[string]interface{}{key: queryItems[i+1]}})
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
		"sort": []map[string]interface{}{
			{"views": "desc"},
		},
		"from": page * size,
		"size": size,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.New(fmt.Sprintf("json.Marshal doc error: %s", err))
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("es.Search error: %w", err)
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("es.Search response error: %s", body)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// todo: 要求写入到传入的日志
	log.Println(res)

	// 解析查询结果
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.New(fmt.Sprintf("json.NewDecoder(res.Body).Decode(&r) error: %s", err))
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

// BulkUpdate 传进来的doc结构体字段后面的tag一定要加上omitempty用来除去空值影响。例子`json:"imageUuid,omitempty"`
func BulkUpdate[T any](es *elasticsearch.Client, docs []T, index string, ids []string) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      index, // The default index name
		Client:     es,    // The Elasticsearch client
		NumWorkers: 4,     // The number of worker goroutines (default: number of CPUs)
		FlushBytes: 5e+6,  // The flush threshold in bytes
	})
	if err != nil {
		fmt.Println("esutil.NewBulkIndexer" + err.Error())
	}
	for index, v := range docs {
		data, err := json.Marshal(map[string]interface{}{
			"doc":           v,
			"doc_as_upsert": true,
		})
		if err != nil {
			fmt.Println("json.Marshal" + err.Error())
		}
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "update",
				DocumentID: ids[index],
				Body:       bytes.NewReader(data),
			},
		)
		if err != nil {
			fmt.Println("indexer add" + err.Error())
		}
	}
	bi.Close(context.Background())
}
