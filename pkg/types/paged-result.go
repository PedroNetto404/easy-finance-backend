package types

type (
	PagedResultMetadata struct {
		TotalCount int64 `json:"total_count"`
		PageCount  int64 `json:"page_count"`
		PageNumber int64 `json:"page_number"`
		PageSize   int64 `json:"page_size"`
		HasNext    bool  `json:"has_next"`
		HasPrev    bool  `json:"has_prev"`
	}

	PagedResult[T any] struct {
		Meta    PagedResultMetadata `json:"meta"`
		Records []T                 `json:"records"`
	}
)

func NewPagedResult[T any](
	limit int64,
	offset int64,
	totalCount int64,
	records []T,
) *PagedResult[T] {
	pageCount := totalCount / limit
	if totalCount%limit > 0 {
		pageCount++
	}

	hasNext := pageCount > offset+1
	hasPrev := offset > 0

	return &PagedResult[T]{
		Meta: PagedResultMetadata{
			TotalCount: totalCount,
			PageCount:  pageCount,
			PageNumber: offset + 1,
			PageSize:   limit,
			HasNext:    hasNext,
			HasPrev:    hasPrev,
		},
		Records: records,
	}
}
