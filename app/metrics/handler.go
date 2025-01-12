package metrics

import (
	"context"
	"encoding/json"
	"net/url"

	"cloud.google.com/go/bigquery"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"go-klaviyo-to-bigquery/app/bq"
	"go-klaviyo-to-bigquery/app/client"
	"go-klaviyo-to-bigquery/internal"
)

// Handler represents the event handler
type Handler struct {
	cfg      *internal.Config
	client   *client.Client
	req      *resty.Request
	resp     *resty.Response
	respErr  *resty.ResponseError
	inserter *bq.DataInserter
}

const (
	Category = "metrics"
	FilterOp = "greater-than"
	Sort     = "-created"
	PageSize = 50
)

// Client 를 사용해서 Klaviyo API 를 호출하는 Handler 를 생성합니다.
func (h *Handler) Handle(ctx context.Context) error {
	data, err := request(ctx, h)
	if err != nil {
		return errors.Wrap(err, "failed to make request to Klaviyo API")
	}

	if !h.inserter.TableExists(ctx, h.cfg.DatasetID, h.cfg.TablePrefix+MetricTable{}.TableName()) {
		_ = h.inserter.CreateTable(ctx, h.cfg.DatasetID, h.cfg.TablePrefix+MetricTable{}.TableName(), h.cfg.DatasetLocation, MetricTable{})
	}
	if err := h.inserter.InsertData(ctx, h.cfg.DatasetID, h.cfg.TablePrefix+MetricTable{}.TableName(), data); err != nil {
		return errors.Wrap(err, "failed to insert data")
	}

	return nil
}

func request(ctx context.Context, h *Handler) ([]bigquery.ValueSaver, error) {
	var pageCursor = ""
	var err error
	var data []bigquery.ValueSaver
	for {
		h.client.ClearQuery()
		q := h.client.GetQuery()
		if pageCursor != "" {
			q.SetPageCursor(pageCursor)
		} else {

		}
		h.resp, err = h.client.R().
			SetQueryString(q.RawQuery()).
			EnableTrace().
			SetContext(ctx).
			SetError(h.respErr).Get("/" + Category)
		if err != nil {
			return data, errors.Wrap(err, "failed to make request to Klaviyo API")
		}
		response := &internal.ResponseData{}
		if json.Unmarshal(h.resp.Body(), response) != nil {
			return data, errors.Wrap(err, "failed to unmarshal response body")
		}

		table := NewMetricTable(response.Data) // TODO: s`함수 분리 필요

		transformData, err := table.TransformFunc()
		if err != nil {
			return data, errors.Wrap(err, "failed to transform data")
		}
		// 테이블이 없으면 생성합니다.
		data = append(data, transformData...)
		if response.Links.Next == "" {
			break
		} else {
			v, err := url.Parse(response.Links.Next)
			if err != nil {
				return data, errors.Wrap(err, "failed to parse query")
			}
			pageCursor = v.Query().Get("page[cursor]")
		}
	}
	return data, nil
}

func NewHandler(client *client.Client, inserter *bq.DataInserter, cfg *internal.Config) *Handler {

	return &Handler{
		cfg:      cfg,
		client:   client,
		inserter: inserter,
	}
}
