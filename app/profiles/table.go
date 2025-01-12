package profiles

import (
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"

	"go-klaviyo-to-bigquery/internal"
)

type ProfileTable struct {
	Items []ProfileItem `json:"items"`
}

// Klaviyo API 를 통해 받아온 데이터 구조를 정의합니다.
type ProfileItem struct {
	// 이벤트의 고유 식별자입니다.
	Id                string `json:"id"`
	ExternalId        string `json:"external_id"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Title             string `json:"title"`
	Organization      string `json:"organization"`
	Properties        string `json:"properties"`
	Image             string `json:"image"`
	Created           string `json:"created"`
	Updated           string `json:"updated"`
	LocationAddress1  string `json:"location_address1"`
	LocationAddress2  string `json:"location_address2"`
	LocationCity      string `json:"location_city"`
	LocationCountry   string `json:"location_country"`
	LocationLatitude  string `json:"location_latitude"`
	LocationLongitude string `json:"location_longitude"`
	LocationRegion    string `json:"location_region"`
	LocationZip       string `json:"location_zip"`
}

func NewProfileTable(data []internal.Data) *ProfileTable {
	var items []ProfileItem

	for _, item := range data {
		attr := item.Attributes.(map[string]interface{})

		properties := attr["properties"].(map[string]interface{})
		propertiesJson, _ := json.Marshal(properties)

		v := ProfileItem{
			Id:                item.Id,
			ExternalId:        getStringFromMap(attr, "external_id"),
			Email:             getStringFromMap(attr, "email"),
			PhoneNumber:       getStringFromMap(attr, "phone_number"),
			FirstName:         getStringFromMap(attr, "first_name"),
			LastName:          getStringFromMap(attr, "last_name"),
			Title:             getStringFromMap(attr, "title"),
			Organization:      getStringFromMap(attr, "organization"),
			Properties:        string(propertiesJson),
			Image:             getStringFromMap(attr, "image"),
			Created:           getStringFromMap(attr, "created"),
			Updated:           getStringFromMap(attr, "updated"),
			LocationAddress1:  getStringFromMap(attr, "location_address1"),
			LocationAddress2:  getStringFromMap(attr, "location_address2"),
			LocationCity:      getStringFromMap(attr, "location_city"),
			LocationCountry:   getStringFromMap(attr, "location_country"),
			LocationLatitude:  getStringFromMap(attr, "location_latitude"),
			LocationLongitude: getStringFromMap(attr, "location_longitude"),
			LocationRegion:    getStringFromMap(attr, "location_region"),
			LocationZip:       getStringFromMap(attr, "location_zip"),
		}
		items = append(items, v)
	}

	return &ProfileTable{
		Items: items,
	}
}

func getStringFromMap(a map[string]interface{}, k string) string {
	if v, ok := a[k].(string); ok {
		return v
	}
	return ""
}

func (t ProfileTable) Schema() bigquery.Schema {
	return bigquery.Schema{
		{Name: "id", Required: true, Type: bigquery.StringFieldType},
		{Name: "external_id", Required: false, Type: bigquery.StringFieldType},
		{Name: "email", Required: false, Type: bigquery.StringFieldType},
		{Name: "phone_number", Required: false, Type: bigquery.StringFieldType},
		{Name: "first_name", Required: false, Type: bigquery.StringFieldType},
		{Name: "last_name", Required: false, Type: bigquery.StringFieldType},
		{Name: "title", Required: false, Type: bigquery.StringFieldType},
		{Name: "organization", Required: false, Type: bigquery.StringFieldType},
		{Name: "properties", Required: false, Type: bigquery.JSONFieldType},
		{Name: "image", Required: false, Type: bigquery.StringFieldType},
		{Name: "created", Required: false, Type: bigquery.TimestampFieldType},
		{Name: "updated", Required: false, Type: bigquery.TimestampFieldType},
		{Name: "location_address1", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_address2", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_city", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_country", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_latitude", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_longitude", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_region", Required: false, Type: bigquery.StringFieldType},
		{Name: "location_zip", Required: false, Type: bigquery.StringFieldType},
	}
}

func (ProfileTable) TableName() string {
	return fmt.Sprintf("profile_date_%s", time.Now().Format("20060102"))
}

func (t ProfileItem) Save() (row map[string]bigquery.Value, insertID string, err error) {

	row = map[string]bigquery.Value{
		"id":                 t.Id,
		"external_id":        t.ExternalId,
		"email":              t.Email,
		"phone_number":       t.PhoneNumber,
		"first_name":         t.FirstName,
		"last_name":          t.LastName,
		"title":              t.Title,
		"organization":       t.Organization,
		"properties":         t.Properties,
		"image":              t.Image,
		"created":            t.Created,
		"updated":            t.Updated,
		"location_address1":  t.LocationAddress1,
		"location_address2":  t.LocationAddress2,
		"location_city":      t.LocationCity,
		"location_country":   t.LocationCountry,
		"location_latitude":  t.LocationLatitude,
		"location_longitude": t.LocationLongitude,
		"location_region":    t.LocationRegion,
		"location_zip":       t.LocationZip,
	}
	return row, t.Id, nil
}

func (t ProfileTable) TransformFunc() ([]bigquery.ValueSaver, error) {
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
