package models

import "reflect"

type PaginationInfo struct {
	TotalCount uint64 `json:"total,omitempty"`
	PageCount  uint64 `json:"pages,omitempty"`
	Page       uint64 `json:"page,omitempty"`
	PerPage    uint64 `json:"per_page,omitempty"`
}

type PaginateResult struct {
	ReflectItems reflect.Value  `json:"-"`
	Items        interface{}  `json:"result,omitempty"`
	Pagination   PaginationInfo `json:"pagination,omitempty"`
}
