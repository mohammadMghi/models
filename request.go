package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IRequest interface {
	GetContext() *gin.Context
	GetAuth() IAuthorization
	AddNewFilter(key string, value interface{})
	RemoveFilterByKey(key string)
	SetBody(body IBaseModel)
	GetBody() (body IBaseModel)
	SetBaseRequest(req *Request)
	GetBaseRequest() *Request
	SetTemp(key string, value interface{})
	GetTemp(key string) (value interface{})
	GetID() interface{}
	GetIDString() string
	AddTag(key string, value bool)
	GetTag(key string) (value *bool)
	RemoveTag(key string)
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

func (request *Request) AddNewFilter(key string, value interface{}) {
	if request.Filters == nil {
		request.Filters = &Filters{}
	}
	request.Filters.Add(key, value)
}

func (request *Request) RemoveFilterByKey(key string) {
	if request.Filters == nil {
		return
	}
	request.Filters.Delete(key)
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

func (request *Request) GetBody() (body IBaseModel) {
	return request.Body
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
	if id == nil || id == "" {
		id = request.ID
	}
	return fmt.Sprintf("%v", id)
}

func (request *Request) AddSort(name string, ascending ...bool) {
	if request.Sort == nil {
		request.Sort = &[]SortItem{}
	}
	*request.Sort = append(*request.Sort, SortItem{
		Name:      name,
		Ascending: len(ascending) > 0 && ascending[0],
	})
}

func (request *Request) AddTag(key string, value bool) {
	if request.Tags == nil {
		request.Tags = map[string]bool{}
	}
	request.Tags[key] = value
}

func (request *Request) GetTag(key string) (value *bool) {
	if request.Tags == nil {
		return nil
	}
	if value, ok := request.Tags[key]; ok {
		return &value
	}
	return nil
}

func (request *Request) RemoveTag(key string) {
	if request.Tags == nil {
		return
	}
	delete(request.Tags, key)
}
