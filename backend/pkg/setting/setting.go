package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DatabaseSetting = &Database{}

type Minio struct {
	EndPoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

var MinioSetting = &Minio{}

var cfg *ini.File

func Setup() {
	_, err := os.Stat("/.dockerenv")
	isDocker := !os.IsNotExist(err)

	configPath := "conf/appLocal.ini"
	if isDocker {
		configPath = "conf/app.ini"
	}

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%s': %v", configPath, err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("minio", MinioSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
