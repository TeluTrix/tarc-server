package service

import (
	"log/slog"

	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/google/uuid"
	"gopkg.in/imdario/mergo.v0"
)

func CreateItem(item schema.Item) (schema.Item, error) {
	result := DB.Create(&item)

	if result.Error != nil {
		var nilItem schema.Item

		slog.Error(result.Error.Error())
		return nilItem, result.Error
	}

	return item, nil

}
func GetItem(id uuid.UUID) (schema.Item, error) {
	var item schema.Item
	item.ID = id

	result := DB.First(&item)

	if result.Error != nil {
		var nilItem schema.Item

		slog.Error(result.Error.Error())
		return nilItem, result.Error
	}

	return item, nil
}
func GetItems() ([]schema.Item, error) {
	var items []schema.Item

	result := DB.Find(&items)

	if result.Error != nil {
		var nilItems []schema.Item

		slog.Error(result.Error.Error())
		return nilItems, result.Error
	}

	return items, nil
}
func UpdateItem(id uuid.UUID, item schema.Item) (schema.Item, error) {
	var updatedItem schema.Item
	updatedItem.ID = id

	result := DB.First(&updatedItem)

	if result.Error != nil {
		var nilItem schema.Item

		slog.Error(result.Error.Error())

		return nilItem, result.Error
	}

	err := mergo.Merge(&updatedItem, item, mergo.WithOverride)
	if err != nil {
		slog.Error(err.Error())
	}

	DB.Save(&updatedItem)

	return updatedItem, nil

}
func DeleteItem(id uuid.UUID) error {
	var deleteItem schema.Item
	deleteItem.ID = id

	result := DB.First(&deleteItem)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}
