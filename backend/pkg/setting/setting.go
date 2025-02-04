package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode  string
	HttpPort int
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf.app.ini")

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
