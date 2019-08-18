package models

import (
	"time"
)

type IBaseModel interface {
	HandleCreateDefaultValues()
	HandleUpdateDefaultValues()
	HandleDeleteDefaultValues()
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

func (base *BaseModel) HandleDeleteDefaultValues() {
	now := time.Now().UTC()
	base.DeletedAt = &now
}
