package models

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Filename    string    `json:"filename"`
	UploadDate  time.Time `json:"upload_date" gorm:"not null" example:"2024-06-01T20:00:00Z"`
}

func ExistMaterialByID(id int) (bool, error) {
	var material Material
	err := db.Select("id").Where("id = ?", id).First(&material).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if material.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddMaterial(data map[string]interface{}) error {
	material := Material{
		Title:       data["title"].(string),
		Description: data["desc"].(string),
		Filename:    data["filename"].(string),
		UploadDate:  data["upload_date"].(time.Time),
	}
	if err := db.Create(&material).Error; err != nil {
		return err
	}
	return nil
}

func GetMaterials() ([]Material, error) {
	var (
		materials []Material
		err       error
	)
	err = db.Find(&materials).Error

	if err != nil {
		return nil, err
	}

	return materials, nil
}
