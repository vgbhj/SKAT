package material_service

import (
	"mime/multipart"

	"github.com/vgbhj/SKAT/models"
	"github.com/vgbhj/SKAT/pkg/minio"
)

type Material struct {
	Title    string
	Desc     string
	FileData *multipart.FileHeader
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

	if err := minio.AddFile(fileName, data, fileSize); err != nil {
		return err
	}

	material := map[string]interface{}{
		"title":     m.Title,
		"desc":      m.Desc,
		"file_name": fileName,
	}

	if err := models.AddMaterial(material); err != nil {
		return err
	}

	return nil
}
