package schema

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	// Gorm default vlaues
	ID        uuid.UUID `gorm:"type:varchar(191);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Actual values
	Name  string
	Path  string
	Color string

	// Relationships
	Items []Item
}

func InitFolder(name string, path string) Folder {
	var folder Folder

	randomColor, err := genRandomColor()
	if err != nil {
		slog.Error(err.Error())
	}

	folder.Name = name
	folder.Path = path
	folder.Color = randomColor

	return folder
}

func genRandomColor() (string, error) {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("#%s", hex.EncodeToString(bytes)), nil
}
