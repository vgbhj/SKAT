package models

import (
	"fmt"
	"log"

	"github.com/vgbhj/SKAT/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Port)))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

}
