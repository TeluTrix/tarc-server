package schema

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	// Gorm default vlaues
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Actual values
	Name  string
	Color string

	// Relationships
	Items []*Item `gorm:"many2many:item_tags"`
}

func InitTag(name string, path string) Tag {
	var tag Tag

	randomColor, err := genRandomColor()
	if err != nil {
		slog.Error(err.Error())
	}

	tag.Name = name
	tag.Color = randomColor

	return tag
}
