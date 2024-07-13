package paginator

import "math"

const (
	defaultPage  = 1
	defaultLimit = 15
)

type PaginatorQuery struct {
	Page  int   `json:"page" form:"page"`
	Limit int64 `json:"limit" form:"limit"`
}

func (p *PaginatorQuery) Adjust() {
	if p.Page < 1 {
		p.Page = defaultPage
	}

	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
}

func (p *PaginatorQuery) Offset() int64 {
	return int64((p.Page - 1)) * p.Limit
}

type Paginator struct {
	Total       int64
	Count       int64
	PerPage     int64
	CurrentPage int
}

func (p Paginator) TotalPages() int {
	if p.Total == 0 {
		return 0
	}

	return int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
}

func (p Paginator) ToResponse() PaginatorResponse {
	return PaginatorResponse{
		Total:       p.Total,
		Count:       p.Count,
		PerPage:     p.PerPage,
		CurrentPage: p.CurrentPage,
		TotalPages:  p.TotalPages(),
	}
}

type PaginatorResponse struct {
	Total       int64 `json:"total"`
	Count       int64 `json:"count"`
	PerPage     int64 `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
}
