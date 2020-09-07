package base

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key" validate:"required";"uuid4"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.CreatedAt = time.Now()
	return nil
}

func (base *Base) BeforeUpdate(*gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}
