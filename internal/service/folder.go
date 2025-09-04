package service

import (
	"log/slog"

	"github.com/TeluTrix/tarc-server/internal/schema"
	"github.com/google/uuid"
	"gopkg.in/imdario/mergo.v0"
)

func CreateFolder(folder schema.Folder) (schema.Folder, error) {
	result := DB.Create(&folder)

	if result.Error != nil {
		var nilFolder schema.Folder

		slog.Error(result.Error.Error())
		return nilFolder, result.Error
	}

	return folder, nil

}
func GetFolder(id uuid.UUID) (schema.Folder, error) {
	var folder schema.Folder
	folder.ID = id

	result := DB.First(&folder)

	if result.Error != nil {
		var nilFolder schema.Folder

		slog.Error(result.Error.Error())
		return nilFolder, result.Error
	}

	return folder, nil
}
func GetFolders() ([]schema.Folder, error) {
	var folders []schema.Folder

	result := DB.Find(&folders)

	if result.Error != nil {
		var nilFolders []schema.Folder

		slog.Error(result.Error.Error())
		return nilFolders, result.Error
	}

	return folders, nil
}
func UpdateFolder(id uuid.UUID, folder schema.Folder) (schema.Folder, error) {
	var updatedFolder schema.Folder
	updatedFolder.ID = id

	result := DB.First(&updatedFolder)

	if result.Error != nil {
		var nilFolder schema.Folder

		slog.Error(result.Error.Error())

		return nilFolder, result.Error
	}

	err := mergo.Merge(&updatedFolder, folder, mergo.WithOverride)
	if err != nil {
		slog.Error(err.Error())
	}

	DB.Save(&updatedFolder)

	return updatedFolder, nil

}
func DeleteFolder(id uuid.UUID) error {
	var deleteFolder schema.Folder
	deleteFolder.ID = id

	result := DB.First(&deleteFolder)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}
