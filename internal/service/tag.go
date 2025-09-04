package service

import (
	"log/slog"

	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/google/uuid"
	"gopkg.in/imdario/mergo.v0"
)

func CreateTag(tag schema.Tag) (schema.Tag, error) {
	result := DB.Create(&tag)

	if result.Error != nil {
		var nilTag schema.Tag

		slog.Error(result.Error.Error())
		return nilTag, result.Error
	}

	return tag, nil

}
func GetTag(id uuid.UUID) (schema.Tag, error) {
	var tag schema.Tag
	tag.ID = id

	result := DB.First(&tag)

	if result.Error != nil {
		var nilTag schema.Tag

		slog.Error(result.Error.Error())
		return nilTag, result.Error
	}

	return tag, nil
}
func GetTags() ([]schema.Tag, error) {
	var tags []schema.Tag

	result := DB.Find(&tags)

	if result.Error != nil {
		var nilTags []schema.Tag

		slog.Error(result.Error.Error())
		return nilTags, result.Error
	}

	return tags, nil
}
func UpdateTag(id uuid.UUID, tag schema.Tag) (schema.Tag, error) {
	var updatedTag schema.Tag
	updatedTag.ID = id

	result := DB.First(&updatedTag)

	if result.Error != nil {
		var nilTag schema.Tag

		slog.Error(result.Error.Error())

		return nilTag, result.Error
	}

	err := mergo.Merge(&updatedTag, tag, mergo.WithOverride)
	if err != nil {
		slog.Error(err.Error())
	}

	DB.Save(&updatedTag)

	return updatedTag, nil

}
func DeleteTag(id uuid.UUID) error {
	var deleteTag schema.Tag
	deleteTag.ID = id

	result := DB.First(&deleteTag)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}
