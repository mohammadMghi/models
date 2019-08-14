package models

import (
	"time"
)

type IBaseModel interface {
}

type BaseModel struct {
	IBaseModel `json:"-"`

	ID        uint64     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
