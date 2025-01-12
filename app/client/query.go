package client

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

// TODO Query 값을 맵으로 처리할 필요 있음
type query struct {
	category string
	fields   []string
	sort     []string
	filter   []filter
	pageSize int
	cursor   string
}

type filter struct {
	op    string
	field string
	value string
}

func newQuery() *query {
	return &query{
		category: "",
		fields:   []string{},
		sort:     []string{},
		filter:   []filter{},
		pageSize: 0,
		cursor:   "",
	}
}

// SetCategory 카테고리를 지정합니다.
func (q *query) SetCategory(category string) *query {
	q.category = category
	return q
}

func (q *query) Category() string {
	return q.category
}

// AddField Field 값을 추가합니다.
func (q *query) AddField(field string) *query {
	q.fields = append(q.fields, field)
	return q
}

// Fields Field 쿼리를 반환합니다.
func (q *query) Fields() string {
	r := ""
	if len(q.fields) == 0 {
		return r
	}
	r = fmt.Sprintf("fields[%s]=", q.category)
	for _, f := range q.fields {
		r += fmt.Sprintf("%s,", f)
	}
	r = removeComma(r)
	return r
}

// AddSort 오름차순은 datetime 내일차순은 -datetime 으로 입력한다.
func (q *query) AddSort(field string) *query {
	q.sort = append(q.sort, field)
	return q
}

// Sort 정렬 값을 반환합니다.
func (q *query) Sort() string {
	r := ""
	if len(q.sort) == 0 {
		return r
	}
	r = "sort="
	for _, f := range q.sort {
		r += fmt.Sprintf("%s,", f)
	}
	r = removeComma(r)
	return r
}

// AddFilter 필터를 추가 합니다.
func (q *query) AddFilter(op, field, value string) *query {
	q.filter = append(q.filter, filter{op: op, field: field, value: value})
	return q
}

// Filters Filter 는 필터링 조건을 반환 합니다.
func (q *query) Filters() string {
	r := ""
	if len(q.filter) == 0 {
		return r
	}

	r += "filter="

	for _, f := range q.filter {
		v, err := opConv(f.op, f.field, f.value)
		if err != nil {
			continue
		}
		r += v + ","
	}
	r = removeComma(r)
	return r
}

const (
	Equals         = "equals"
	LessThan       = "less-than"
	LessOrEqual    = "less-or-equal"
	GreaterThan    = "greater-than"
	GreaterOrEqual = "greater-or-equal"
	Contains       = "contains"
	EndsWith       = "ends-with"
	StartsWith     = "starts-with"
	Any            = "any"
	Has            = "has"
)

func opConv(op, f, v string) (string, error) {
	if op == "" {
		return "", fmt.Errorf("op is empty")
	}
	if f == "" {
		return "", fmt.Errorf("field is empty")
	}
	if v == "" {
		return "", fmt.Errorf("value is empty")
	}

	switch op {
	case Equals:
		{
			v = fmt.Sprintf("equals(%s,\"%s\")", f, v)
			return v, nil
		}
	case LessThan:
		return fmt.Sprintf("less-then(%v,%v)", f, v), nil
	case LessOrEqual:
		return fmt.Sprintf("less-or-equal(%v,%v)", f, v), nil
	case GreaterThan:
		return fmt.Sprintf("greater-than(%v,%v)", f, v), nil
	case GreaterOrEqual:
		return fmt.Sprintf("greater-or-equal(%v,%v)", f, v), nil
	case Contains:
		return fmt.Sprintf("contains(%v,%v)", f, v), nil
	case EndsWith:
		return fmt.Sprintf("ends-with(%v,%v)", f, v), nil
	case StartsWith:
		return fmt.Sprintf("starts-with(%v,%v)", f, v), nil
	case Any:
		return fmt.Sprintf("less-or-equal(%v,%v)", f, v), nil
	case Has:
		return fmt.Sprintf("has(%v,%v)", f, v), nil
	}
	return "", errors.Wrap(fmt.Errorf("op is not found"), "opConv")
}

// removeComma 콤마를 삭제합니다.
func removeComma(r string) string {
	return string(bytes.TrimRight([]byte(r), ","))
}

// PageSize 페이지 사이즈 쿼리를 반환합니다.
func (q *query) PageSize() string {
	return fmt.Sprintf("page[size]=%d", q.pageSize)
}

func (q *query) SetPageSize(size int) *query {
	q.pageSize = size
	return q
}

// PageCursor 페이지 커서 쿼리 값을반환 합니다.
func (q *query) PageCursor() string {
	return fmt.Sprintf("page[cursor]=%s", q.cursor)
}

func (q *query) SetPageCursor(cursor string) *query {
	q.cursor = cursor
	return q
}

// RawQuery RawQuery 값을 반환 합니다.
func (q *query) RawQuery() string {
	var r string
	r = ""

	if q.Fields() != "" {
		r += q.Fields()
	}

	if q.Filters() != "" {
		if r != "" {
			r += "&"
		}
		r += q.Filters()
	}

	if q.Sort() != "" {
		if r != "" {
			r += "&"
		}
		r += q.Sort()
	}

	if q.pageSize != 0 {
		if r != "" {
			r += "&"
		}
		r += q.PageSize()
	}

	if q.cursor != "" {
		if r != "" {
			r += "&"
		}
		r += q.PageCursor()
	}
	return r
}
