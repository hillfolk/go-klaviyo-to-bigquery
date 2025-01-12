package internal

import (
	"cloud.google.com/go/bigquery"
)

type SchemaGenerator interface {
	Schema() bigquery.Schema
}

type CsvWriter interface {
	CsvWriter(filePath string) error
}

type Table interface {
	SchemaGenerator
	TableName() string
}

type Item interface {
	Save() (map[string]bigquery.Value, string, error)
	Row() []string
}
