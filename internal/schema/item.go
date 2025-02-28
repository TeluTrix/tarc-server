package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	// Gorm default vlaues
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Actual values
	Path string

	// Relationships
	FolderID uuid.UUID `gorm:"type:varchar(191)"`
	Tags     []*Tag    `gorm:"many2many:item_tags"`
	//Tags     []Tag  `gorm:"many2many:items_tags;"`
}
