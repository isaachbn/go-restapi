package base

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

const UUID_BLANK = "00000000-0000-0000-0000-000000000000"

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

func (base *Base) GenerateIdentifier()  {
	valueId, _ := base.ID.Value()
	if valueId == UUID_BLANK {
		base.ID, _ = uuid.NewV4()
	}
}