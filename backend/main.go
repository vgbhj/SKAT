package main

import (
	"fmt"

	"github.com/vgbhj/SKAT/models"
	"github.com/vgbhj/SKAT/pkg/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	fmt.Println("Hello, World!")
}
