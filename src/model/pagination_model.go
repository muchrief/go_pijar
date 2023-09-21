package model

import "fmt"

type SortType string

const (
	ASC  SortType = "asc"
	DESC SortType = "desc"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	SortType   SortType    `json:"sort_type,omitempty" query:"sort_type"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 20
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	return fmt.Sprintf("%s %s", p.Sort, p.GetSortType())
}

func (p *Pagination) GetSortType() SortType {
	if p.SortType == "" {
		p.SortType = DESC
	}

	return p.SortType
}
