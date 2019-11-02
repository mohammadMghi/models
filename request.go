package models

import "github.com/gin-gonic/gin"

type IRequest interface {
	GetContext() *gin.Context
	AddNewFilter(key string, value interface{})
	SetBody(body IBaseModel)
	SetBaseRequest(req *Request)
	GetBaseRequest() *Request
}

type Request struct {
	IRequest

	Context    *gin.Context
	Params     *gin.Params
	ID         interface{}
	Fields     *Fields
	Filters    *Filters
	Sort       *[]SortItem
	Page       uint64
	PerPage    uint64
	Body       IBaseModel
	ExtraQuery map[string]interface{}

	// config tags from loading or not
	Tags map[string]bool
}

func (request *Request) GetContext() *gin.Context {
	return request.Context
}

func (request *Request) SetBaseRequest(req *Request) {
	request.IRequest = req
}

func (request *Request) GetBaseRequest() *Request {
	return request
}

func (request *Request) SetBody(body IBaseModel) {
	request.Body = body
}
