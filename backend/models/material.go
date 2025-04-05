package models

type Material struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	FileName string `json:"file_name"`
}

func AddMaterial(data map[string]interface{}) error {
	material := Material{
		Title:    data["title"].(string),
		Desc:     data["desc"].(string),
		FileName: data["file_name"].(string),
	}
	if err := db.Create(&material).Error; err != nil {
		return err
	}
	return nil
}
