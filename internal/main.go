package main

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/log"
)

func main() {
	ConfigMain.LoadEnv()
	config.FillSettings()
	config.FillFlags()

	FileName := config.Settings.FILENAME_XGML
	log.Info("directory: ", config.Settings.DIRECTORY_SOURCE)
	log.Info("file xgml: ", FileName)
	logic.StartFillAll(FileName)

	//go parse_go.ParseDir("") //удалить
	//go print("1")
}
