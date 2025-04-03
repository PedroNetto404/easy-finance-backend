package types

type QueryArgs struct {
	Limit    int64  `json:"limit"`
	Offset   int64  `json:"offset"`
	SortBy   string `json:"sort_by"`
  Ascending bool   `json:"asc"`	
	Filter map[string]any `json:"filter"`
}

func (q *QueryArgs) CheckDefaults() {
	if q.Limit == 0 {
		q.Limit = 10
	}
	if q.Offset == 0 {
		q.Offset = 0
	}
	if q.SortBy == "" {
		q.SortBy = "id"
	}
	if q.Filter == nil {
		q.Filter = make(map[string]any)
	}
}