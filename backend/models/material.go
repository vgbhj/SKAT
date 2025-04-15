package models

import "time"

type Material struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Filename    string    `json:"filename"`
	UploadDate  time.Time `json:"upload_date" gorm:"not null" example:"2024-06-01T20:00:00Z"`
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
