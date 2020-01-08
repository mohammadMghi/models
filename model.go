package models

import (
	"fmt"
	"time"
)

type IBaseModel interface {
	HandleCreateDefaultValues()
	HandleUpdateDefaultValues()
	HandleUpsertDefaultValues()
	HandleDeleteDefaultValues()
	GetID() interface{}
	GetIDString() string
	SetID(id interface{})
	Populate(request IRequest)
}

type BaseModel struct {
	IBaseModel `json:"-"`

	ID        uint64     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (base *BaseModel) HandleCreateDefaultValues() {
	base.CreatedAt = time.Now().UTC()
	base.UpdatedAt = time.Now().UTC()
}

func (base *BaseModel) HandleUpdateDefaultValues() {
	base.UpdatedAt = time.Now().UTC()
}

func (base *BaseModel) HandleUpsertDefaultValues() {
	base.CreatedAt = time.Now().UTC()
	base.UpdatedAt = time.Now().UTC()
}

func (base *BaseModel) HandleDeleteDefaultValues() {
	now := time.Now().UTC()
	base.DeletedAt = &now
}

func (base *BaseModel) GetID() interface{} {
	return base.ID
}

func (base *BaseModel) GetIDString() string {
	return fmt.Sprintf("%v", base.ID)
}

func (base *BaseModel) SetID(id interface{}) {
}

func (base *BaseModel) Populate(request IRequest) {
}
