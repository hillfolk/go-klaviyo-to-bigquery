package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_query_SetCategory(t *testing.T) {
	type fields struct {
		category string
		field    []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}

	type args struct {
		category string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "카테고리_입력_검증",
			fields: fields{
				category: "profiles",
			},
			args: args{
				"profiles",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.field,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			q.SetCategory(tt.args.category)
			// 카테고리 값 검증
			assert.Equal(t, tt.args.category, q.category)
		})

	}
}

func Test_query_Category(t *testing.T) {
	type fields struct {
		category string
		field    []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "카테고리_출력_검증",
			fields: fields{
				category: "events",
				field:    []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			want: "events",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.field,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.Category(), "Category()")
		})
	}
}

func Test_query_AddField(t *testing.T) {
	type fields struct {
		category string
		field    []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "필드_단수_추가_검증",
			fields: fields{
				category: "profiles",
				field:    []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				"email",
			},
			want: "fields[profiles]=email",
		},
		{
			name: "필드_복수_추가_검증",
			fields: fields{
				category: "profiles",
				field: []string{
					"email",
				},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				"name",
			},
			want: "fields[profiles]=email,name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.field,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			q.AddField(tt.args.field)

			// 필드 값 검증
			assert.Contains(t, q.fields, tt.args.field)
			assert.Equal(t, q.Fields(), tt.want)
		})
	}
}

func Test_query_AddSort(t *testing.T) {
	type fields struct {
		category string
		field    []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "정렬_추가_검증",
			fields: fields{
				category: "profiles",
				field:    []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				"created",
			},
		},
		{
			name: "정렬_추가_검증",
			fields: fields{
				category: "profiles",
				field:    []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				"-created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.field,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			q.AddSort(tt.args.field)
			assert.Contains(t, q.sort, tt.args.field)
		})
	}
}

func Test_query_Sort(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "정렬_출력_검증1",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort: []string{
					"created",
					"-updated",
					"datetime",
				},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			want: "sort=created,-updated,datetime",
		},
		{
			name: "정렬_출력_검증2",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort: []string{
					"created",
					"-updated",
					"datetime",
					"email",
					"-phone_number",
				},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			want: "sort=created,-updated,datetime,email,-phone_number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.Sort(), "failed to Sort()")
		})
	}
}

func Test_query_Filters(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "필터_출력_검증1",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort:     []string{},
				filter: []filter{
					{
						op:    "equals",
						field: "email",
						value: "xxx@example.com",
					},
					{
						op:    "equals",
						field: "phone_number",
						value: "123-456-7890",
					},
				},
				pageSize: 0,
				cursor:   "",
			},
			want: "filter=equals(email,\"xxx@example.com\"),equals(phone_number,\"123-456-7890\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.Filters(), "Filters()")
		})
	}
}

func Test_removeComma(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "콤마_제거_검증",
			args: args{
				"fields[profiles]=email,name,",
			},
			want: "fields[profiles]=email,name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeComma(tt.args.r), "removeComma(%v)", tt.args.r)
		})
	}
}

func Test_query_PageSize(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "페이지_사이즈_출력_검증",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 20,
				cursor:   "",
			},
			want: "page[size]=20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.PageSize(), "PageSize()")
		})
	}
}

func Test_query_PageCursor(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "페이지_커서_출력_검증",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "e3b0c44298fc1c149b934ca495991b7852b855",
			},
			want: "page[curser]=e3b0c44298fc1c149b934ca495991b7852b855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.PageCursor(), "PageCursor()")
		})
	}
}

func Test_query_RawQuery(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "쿼리_출력_검증",
			fields: fields{
				category: "profiles",
				fields: []string{
					"email",
					"name",
				},
				sort: []string{
					"created",
					"-updated",
				},
				filter: []filter{
					{
						op:    "equals",
						field: "email",
						value: "xxx@example.com",
					},
					{
						op:    "equals",
						field: "phone_number",
						value: "123-456-7890",
					},
				},
				pageSize: 20,
				cursor:   "e3b0c44298fc1c149b934ca495991b7852b855",
			},
			want: "fields[profiles]=email,name&filter=equals(email,\"xxx@example.com\"),equals(phone_number,\"123-456-7890\")&sort=created,-updated&page[size]=20&page[curser]=e3b0c44298fc1c149b934ca495991b7852b855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			assert.Equalf(t, tt.want, q.RawQuery(), "RawQuery()")
		})
	}
}

func Test_query_SetPageSize(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	type args struct {
		size int
	}
	want := "page[size]=20"
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "페이지_사이즈_설정_검증",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			q.SetPageSize(tt.args.size)
			assert.Equal(t, tt.args.size, q.pageSize)
			assert.Equal(t, want, q.PageSize())
		})
	}
}

func Test_query_SetPageCursor(t *testing.T) {
	type fields struct {
		category string
		fields   []string
		sort     []string
		filter   []filter
		pageSize int
		cursor   string
	}
	type args struct {
		cursor string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "페이지_커서_설정_검증",
			fields: fields{
				category: "profiles",
				fields:   []string{},
				sort:     []string{},
				filter:   []filter{},
				pageSize: 0,
				cursor:   "",
			},
			args: args{
				"e3b0c44298fc1c149b934ca495991b7852b855",
			},
			want: "page[curser]=e3b0c44298fc1c149b934ca495991b7852b855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				category: tt.fields.category,
				fields:   tt.fields.fields,
				sort:     tt.fields.sort,
				filter:   tt.fields.filter,
				pageSize: tt.fields.pageSize,
				cursor:   tt.fields.cursor,
			}
			q.SetPageCursor(tt.args.cursor)
			assert.Equal(t, tt.args.cursor, q.cursor)
			assert.Equal(t, tt.want, q.PageCursor())
		})
	}
}
