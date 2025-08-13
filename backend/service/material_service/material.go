package material_service

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/vgbhj/SKAT/models"
	"github.com/vgbhj/SKAT/pkg/minio"
)

type Material struct {
	Title    string
	Desc     string
	FileData *multipart.FileHeader
	UserName string
	// user_id integer [ref: > users.id]
	// faculty_id integer [ref: > faculty.id]
	// subject_id integer [ref: > subject.id]
	// year_id integer [ref: > year.id]
	// university_id integer [ref: > university.id]
}

func (m *Material) Add() error {
	fileName := m.FileData.Filename
	fileSize := m.FileData.Size
	data, err := m.FileData.Open()
	if err != nil {
		return err
	}
	defer data.Close()

	objectPath := fmt.Sprintf(
		"user_%s/date_%s/title_%s/%s",
		m.UserName, time.Now().Format("2006-01-02_15-04-05"), m.Title, fileName,
	)

	if err := minio.AddFile(objectPath, data, fileSize); err != nil {
		return err
	}

	material := map[string]interface{}{
		"title":       m.Title,
		"desc":        m.Desc,
		"filename":    objectPath,
		"upload_date": time.Now(),
	}

	if err := models.AddMaterial(material); err != nil {
		return err
	}

	return nil
}

func GetAll() ([]models.Material, error) {
	materials, err := models.GetMaterials()
	if err != nil {
		return nil, err
	}

	return materials, nil
}
