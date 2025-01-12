package metrics

import (
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"

	"go-klaviyo-to-bigquery/internal"
)

// Klaviyo API 를 통해 받아온 데이터 구조를 정의합니다.
type MetricTable struct {
	Items []MetricItem `json:"items"`
}

type MetricItem struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	IntegrationId       string `json:"integration_id"`
	IntegrationName     string `json:"integration_name"`
	IntegrationCategory string `json:"integration_category"`
	Created             string `json:"created"`
	Updated             string `json:"updated"`
}

func (m MetricItem) Save() (row map[string]bigquery.Value, insertID string, err error) {
	row = map[string]bigquery.Value{
		"id":                   m.Id,
		"name":                 m.Name,
		"integration_id":       m.IntegrationId,
		"integration_name":     m.IntegrationName,
		"integration_category": m.IntegrationCategory,
		"created":              m.Created,
		"updated":              m.Updated,
	}
	return row, m.Id, nil
}

func NewMetricTable(data []internal.Data) *MetricTable {
	var items []MetricItem

	for _, item := range data {
		attr := item.Attributes.(map[string]interface{})
		integration := attr["integration"].(map[string]interface{})
		v := MetricItem{
			Id:                  item.Id,
			Name:                attr["name"].(string),
			IntegrationId:       integration["id"].(string),
			IntegrationName:     integration["name"].(string),
			IntegrationCategory: integration["category"].(string),
			Created:             attr["created"].(string),
			Updated:             attr["updated"].(string),
		}
		items = append(items, v)
	}

	return &MetricTable{
		Items: items,
	}
}

func (t MetricTable) Schema() bigquery.Schema {
	return bigquery.Schema{
		{Name: "id", Required: true, Type: bigquery.StringFieldType},
		{Name: "name", Required: true, Type: bigquery.StringFieldType},
		{Name: "integration_id", Required: true, Type: bigquery.StringFieldType},
		{Name: "integration_name", Required: true, Type: bigquery.StringFieldType},
		{Name: "integration_category", Required: true, Type: bigquery.StringFieldType},
		{Name: "created", Required: true, Type: bigquery.TimestampFieldType},
		{Name: "updated", Required: true, Type: bigquery.TimestampFieldType},
	}
}

func (t MetricTable) TableName() string {
	return fmt.Sprintf("metric_date_%s", time.Now().Format("20060102"))
}

func (t MetricTable) TransformFunc() ([]bigquery.ValueSaver, error) {
	var transformedData []bigquery.ValueSaver
	for i := len(t.Items) - 1; i >= 0; i-- {
		Log(t.Items[i])
		transformedData = append(transformedData, t.Items[i])
	}
	return transformedData, nil
}

func Log(item bigquery.ValueSaver) {
	fmt.Println(item)
}
