package main

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/constants"
	"github.com/ManyakRus/image_packages/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	//_ "net/http/pprof" //удалить
	"time"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()
	config.FillFlags()

	//go http.ListenAndServe(":8080", nil) //удалить

	StartAt := time.Now()
	FileName := config.Settings.FILENAME_GRAPHML
	log.Info("directory: ", config.Settings.DIRECTORY_SOURCE)
	log.Info("file graphml: ", FileName)
	ok := logic.StartFillAll(FileName)
	if ok == false {
		println(constants.TEXT_HELP)
	}

	log.Info("Time passed: ", time.Since(StartAt))

	//go parse_go.ParseDir("") //удалить
	//go print("1")
	//micro.Sleep(40000) //удалить

}
