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

	ID        uint64     `json:"id,omitempty" dl:"read_only"`
	CreatedAt time.Time  `json:"created_at,omitempty" dl:"read_only"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" dl:"read_only"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" dl:"read_only"`
}

func (base *BaseModel) HandleCreateDefaultValues() {
	base.CreatedAt = time.Now().UTC()
}

func (base *BaseModel) HandleUpdateDefaultValues() {
	now := time.Now().UTC()
	base.UpdatedAt = &now
}

func (base *BaseModel) HandleUpsertDefaultValues() {
	now := time.Now().UTC()
	base.CreatedAt = now
	base.UpdatedAt = &now
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
