package util

import "github.com/vgbhj/SKAT/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
