package models

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"sync"
)

type IRequest interface {
	Populate(requestToPopulate IRequest) (populated IRequest)
	GetContext() *gin.Context
	GetAuth() IAuthorization
	Authenticated() bool
	GetCurrentAccount(request IRequest) interface{}
	GetCurrentAccountId(request IRequest) interface{}
	HasRole(roles ...string) bool
	AddNewFilter(key string, value interface{})
	GetFilter(key string) (value interface{})
	RemoveFilterByKey(key string)
	SetBody(body IBaseModel)
	GetBody() (body IBaseModel)
	SetBaseRequest(req *Request)
	GetBaseRequest() *Request
	SetTemp(key string, value interface{})
	GetTemp(key string) (value interface{})
	RemoveTemp(key string)
	GetID() interface{}
	SetID(id interface{})
	GetIDString() string
	SetTag(key string, value bool)
	GetTag(key string) (value *bool)
	RemoveTag(key string)
	AddExtraQuery(key string, value interface{})
	RemoveExtraQueryByKey(key string, value interface{})
	MustLocalize(lc *i18n.LocalizeConfig) string
	Lock()
	Unlock()
	RLock()
	RUnlock()
	SetLanguage(lang string, messagesPath string) *Request
}

type Language struct {
	AcceptLanguage string
	Localizer      *i18n.Localizer
}

type Request struct {
	IRequest

	Mutex           *sync.RWMutex
	Context         *gin.Context
	Auth            IAuthorization
	Params          *Params
	ID              interface{}
	Fields          *Fields
	Filters         *Filters
	Sort            *[]SortItem
	Page            uint64
	PerPage         uint64
	Body            IBaseModel
	ExtraQuery      map[string]interface{}
	CurrentLanguage *Language

	Tags map[string]bool
	// temporary data for further use
	Temp map[string]interface{}
}

func (request *Request) Populate(requestToPopulate IRequest) (populated IRequest) {
	populated = requestToPopulate
	req := requestToPopulate.GetBaseRequest()
	if req.Context == nil {
		req.Context = request.Context
	}
	if req.CurrentLanguage == nil {
		req.CurrentLanguage = request.CurrentLanguage
	}
	if req.Temp == nil {
		req.Temp = request.Temp
	} else {
		req0 := request.GetBaseRequest()
		if req0.Temp != nil {
			for k, v := range req0.Temp {
				if _, ok := req.Temp[k]; !ok {
					req.SetTemp(k, v)
				}
			}
		}
	}
	req.Mutex = request.Mutex
	if req.Tags == nil {
		req.Tags = request.Tags
	}
	if req.Page == 0 {
		req.Page = request.Page
	}
	if req.PerPage == 0 {
		req.PerPage = request.PerPage
	}
	return
}

func (request *Request) Lock() {
	if request.Mutex == nil {
		request.Mutex = &sync.RWMutex{}
	}
	request.Mutex.Lock()
}

func (request *Request) Unlock() {
	request.Mutex.Unlock()
}

func (request *Request) RLock() {
	if request.Mutex == nil {
		request.Mutex = &sync.RWMutex{}
	}
	request.Mutex.RLock()
}

func (request *Request) RUnlock() {
	request.Mutex.RUnlock()
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

func (request *Request) GetFilter(key string) (value interface{}) {
	if request.Filters == nil {
		return
	}
	value, _ = (*request.Filters)[key]
	return
}

func (request *Request) RemoveFilterByKey(key string) {
	if request.Filters == nil {
		return
	}
	request.Filters.Delete(key)
}

// auth
func (request *Request) GetAuth() IAuthorization {
	return request.Auth
}

func (request *Request) Authenticated() bool {
	return request.Auth != nil && request.Auth.Authenticated()
}

func (request *Request) GetCurrentAccount(req IRequest) interface{} {
	if request.Auth == nil {
		return nil
	}
	return request.Auth.GetCurrentAccount(req)
}

func (request *Request) GetCurrentAccountId(req IRequest) interface{} {
	if request.Auth == nil {
		return nil
	}
	return request.Auth.GetCurrentAccountId(req)
}

func (request *Request) HasRole(roles ...string) bool {
	return request.Auth != nil && request.Auth.HasRole(roles...)
}

//

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
	request.CurrentLanguage = req.CurrentLanguage
}

func (request *Request) GetBaseRequest() *Request {
	return request
}

func (request *Request) SetBody(body IBaseModel) {
	request.Body = body
	request.SetTemp("body", body)
}

func (request *Request) GetBody() (body IBaseModel) {
	return request.Body
}

func (request *Request) SetTemp(key string, value interface{}) {
	if request.Temp == nil {
		request.Temp = map[string]interface{}{}
	}
	request.Lock()
	request.Temp[key] = value
	request.Unlock()
}

func (request *Request) GetTemp(key string) (value interface{}) {
	if request.Temp == nil {
		return
	}
	request.RLock()
	value, _ = request.Temp[key]
	request.RUnlock()
	return
}

func (request *Request) RemoveTemp(key string) {
	if request.Temp == nil {
		return
	}
	request.Lock()
	delete(request.Temp, key)
	request.Unlock()
}

func (request *Request) GetID() interface{} {
	id := request.ID
	if id == nil && request.Body != nil {
		id = request.Body.GetID()
	}
	return id
}

func (request *Request) SetID(id interface{}) {
	request.ID = id
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

func (request *Request) SetTag(key string, value bool) {
	request.Lock()
	if request.Tags == nil {
		request.Tags = map[string]bool{}
	}
	request.Tags[key] = value
	request.Unlock()
}

func (request *Request) GetTag(key string) (value *bool) {
	if request.Tags == nil {
		return nil
	}
	request.RLock()
	defer request.RUnlock()
	if value, ok := request.Tags[key]; ok {
		return &value
	}
	return nil
}

func (request *Request) RemoveTag(key string) {
	if request.Tags == nil {
		return
	}
	request.Lock()
	delete(request.Tags, key)
	request.Unlock()
}

func (request *Request) AddExtraQuery(key string, value interface{}) {
	request.Lock()
	if request.ExtraQuery == nil {
		request.ExtraQuery = map[string]interface{}{}
	}
	request.ExtraQuery[key] = value
	request.Unlock()
}

func (request *Request) RemoveExtraQueryByKey(key string, value interface{}) {
	if request.ExtraQuery == nil {
		return
	}
	request.Lock()
	delete(request.ExtraQuery, key)
	request.Unlock()
}

func (request *Request) Language() string {
	return request.CurrentLanguage.AcceptLanguage
}

func (request *Request) MustLocalize(lc *i18n.LocalizeConfig) string {
	req := request.GetBaseRequest()
	if req.CurrentLanguage != nil {
		return req.CurrentLanguage.Localizer.MustLocalize(lc)
	}
	return DefaultLanguage.Localizer.MustLocalize(lc)
}

func (request *Request) SetLanguage(lang string, messagesPath string) *Request {
	var LanguageBundle = i18n.NewBundle(language.English)
	LanguageBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	LanguageBundle.LoadMessageFile(messagesPath)
	request.CurrentLanguage = &Language{
		AcceptLanguage: lang,
		Localizer:      i18n.NewLocalizer(LanguageBundle, lang),
	}
	return request
}
