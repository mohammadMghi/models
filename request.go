package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IRequest interface {
	GetContext() *gin.Context
	GetAuth() IAuthorization
	AddNewFilter(key string, value interface{})
	SetBody(body IBaseModel)
	SetBaseRequest(req *Request)
	GetBaseRequest() *Request
	SetTemp(key string, value interface{})
	GetTemp(key string) (value interface{})
	GetID() interface{}
	GetIDString() string
}

type Request struct {
	IRequest

	Context    *gin.Context
	Auth       IAuthorization
	Params     *Params
	ID         interface{}
	Fields     *Fields
	Filters    *Filters
	Sort       *[]SortItem
	Page       uint64
	PerPage    uint64
	Body       IBaseModel
	ExtraQuery map[string]interface{}

	Tags map[string]bool
	// temporary data for further use
	Temp map[string]interface{}
}

func (request *Request) GetContext() *gin.Context {
	return request.Context
}

func (request *Request) GetAuth() IAuthorization {
	return request.Auth
}

func (request *Request) SetBaseRequest(req *Request) {
	request.Context = req.Context
	request.Params = req.Params
	request.ID = req.ID
	request.Fields = req.Fields
	request.Filters = req.Filters
	request.Sort = req.Sort
	request.Page = req.Page
	request.PerPage = req.PerPage
	request.Body = req.Body
	request.ExtraQuery = req.ExtraQuery
	request.Tags = req.Tags
}

func (request *Request) GetBaseRequest() *Request {
	return request
}

func (request *Request) SetBody(body IBaseModel) {
	request.Body = body
}

func (request *Request) SetTemp(key string, value interface{}) {
	if request.Temp == nil {
		request.Temp = map[string]interface{}{}
	}
	request.Temp[key] = value
}

func (request *Request) GetTemp(key string) (value interface{}) {
	if request.Temp == nil {
		return
	}
	value, _ = request.Temp[key]
	return
}

func (request *Request) GetID() interface{} {
	id := request.ID
	if id == nil && request.Body != nil {
		id = request.Body.GetID()
	}
	return id
}

func (request *Request) GetIDString() string {
	var id interface{}
	if request.Body != nil {
		id = request.Body.GetIDString()
	}
	if id == nil {
		id = request.ID
	}
	return fmt.Sprintf("%v", id)
}
