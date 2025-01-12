package events

import (
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"

	"go-klaviyo-to-bigquery/internal"
)

/*

 */

func NewEventTable(data []internal.Data) *EventTable {
	var items []EventItem

	for _, item := range data {
		attr := item.Attributes.(map[string]interface{})
		profile := item.Relationships.Profile.Data
		metric := item.Relationships.Metric.Data
		properties := attr["event_properties"].(map[string]interface{})
		propertiesJson, _ := json.Marshal(properties)

		v := EventItem{
			Id:              item.Id,
			MetricId:        metric.Id,
			ProfileId:       profile.Id,
			EventProperties: string(propertiesJson),
			Datetime:        attr["datetime"].(string),
			Uuid:            attr["uuid"].(string),
		}
		items = append(items, v)
	}

	return &EventTable{
		Items: items,
	}
}

type EventTable struct {
	Items []EventItem `json:"items"`
}

func (t EventTable) TableName() string {
	return fmt.Sprintf("event_date_%s", time.Now().Format("20060102"))
}

type EventItem struct {
	Id              string `json:"id"`
	MetricId        string `json:"metric_id"`
	ProfileId       string `json:"profile_id"`
	EventProperties string `json:"event_properties"`
	Datetime        string `json:"datetime"`
	Uuid            string `json:"uuid"`
}

func (e EventItem) Save() (row map[string]bigquery.Value, insertID string, err error) {
	row = map[string]bigquery.Value{
		"id":               e.Id,
		"metric_id":        e.MetricId,
		"profile_id":       e.ProfileId,
		"event_properties": e.EventProperties,
		"datetime":         e.Datetime,
		"uuid":             e.Uuid,
	}
	return row, bigquery.NoDedupeID, nil
}

func (EventItem) Row() (row []string) {
	var values []string
	for i := range row {
		values = append(values, row[i])
	}

	return values
}

func (t EventTable) TransformFunc() ([]bigquery.ValueSaver, error) {
	var transformedData []bigquery.ValueSaver
	for i := len(t.Items) - 1; i >= 0; i-- {
		Log(t.Items[i])
		transformedData = append(transformedData, t.Items[i])
	}
	return transformedData, nil
}

func (EventTable) Schema() bigquery.Schema {
	return bigquery.Schema{
		{Name: "id", Required: true, Type: bigquery.StringFieldType},
		{Name: "metric_id", Required: true, Type: bigquery.StringFieldType},
		{Name: "profile_id", Required: true, Type: bigquery.StringFieldType},
		{Name: "event_properties", Required: true, Type: bigquery.JSONFieldType},
		{Name: "datetime", Required: true, Type: bigquery.TimestampFieldType},
		{Name: "uuid", Required: true, Type: bigquery.StringFieldType},
	}
}

func Log(item bigquery.ValueSaver) {
	fmt.Println(item)
}
