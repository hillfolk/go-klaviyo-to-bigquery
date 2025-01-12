package bq

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/pkg/errors"

	"go-klaviyo-to-bigquery/internal"
)

type DataInserter struct {
	bqClient *bigquery.Client
}

func NewBigQueryDataInsert(client *bigquery.Client) *DataInserter {
	return &DataInserter{
		bqClient: client,
	}
}

func (b *DataInserter) CreateTable(ctx context.Context, datasetID, tableID, location string, schemaGen internal.SchemaGenerator) error {
	if err := b.bqClient.Dataset(datasetID).Table(tableID).Create(ctx, &bigquery.TableMetadata{
		Schema:   schemaGen.Schema(),
		Location: location,
	}); err != nil {
		return errors.Wrap(err, "failed to create table")
	}

	return nil
}

func (b *DataInserter) InsertData(ctx context.Context, datasetID, tableID string, data []bigquery.ValueSaver) error {
	const chunkSize = 200
	for startIdx := 0; startIdx < len(data); startIdx += chunkSize {
		endIdx := startIdx + chunkSize
		if endIdx > len(data) {
			endIdx = len(data)
		}
		if err := b.bqClient.Dataset(datasetID).Table(tableID).Inserter().Put(ctx, data[startIdx:endIdx]); err != nil {
			return errors.Wrap(err, "failed to insert data")
		}
	}
	return nil
}

// 테이블 존재 여부 확인
func (b *DataInserter) TableExists(ctx context.Context, datasetID, tableID string) bool {
	table := b.bqClient.Dataset(datasetID).Table(tableID)
	_, err := table.Metadata(ctx)
	if err != nil {
		return false
	}
	return true
}
