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
		Meta PagedResultMetadata `json:"meta"`
		Records []T `json:"records"`
	}
)