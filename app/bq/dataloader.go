package bq

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/pkg/errors"
)

type DataLoader struct {
	bqClient *bigquery.Client
}

func NewBigQueryDataLoader(client *bigquery.Client) *DataLoader {
	return &DataLoader{
		bqClient: client,
	}
}

// QueryTable queries a table in BigQuery
func selectQuery(ctx context.Context, bqClient *bigquery.Client, fullTableId string) error {
	var idx int = 0
	query := fmt.Sprintf("SELECT * FROM %s;", fullTableId)

	q := bqClient.Query(query)

	it, err := q.Read(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to read query result")
	}
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err != nil {
			break
		}
		fmt.Printf("Row %d: %v\n", idx, row)
		idx++
	}
	return nil
}

func (b *DataLoader) QueryTable(ctx context.Context, client *bigquery.Client, datasetID, tableID string, data interface{}) error {

	fullTableId := fmt.Sprintf("%s.%s.%s", client.Project(), datasetID, tableID)
	err := selectQuery(ctx, client, fullTableId)
	if err != nil {
		return errors.Wrap(err, "failed to query table")
	}
	return nil
}
