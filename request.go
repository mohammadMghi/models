package models

import "github.com/gin-gonic/gin"

type IRequest interface {
	AddNewFilter(key string, value interface{})
}

type Request struct {
	IRequest

	Context *gin.Context
	Params  *gin.Params
	Fields  *Fields
	Filters *Filters
	Sort    *[]SortItem
	Page    *uint64
	PerPage *uint64
}

func (request *Request) AddNewFilter(key string, value interface{})  {
	(*request.Filters)[key] = value
}
