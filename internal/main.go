package main

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/constants"
	"github.com/ManyakRus/image_packages/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	//"os"
	//"runtime/pprof"

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

	//f, err := os.Create("./cpu.pprof")
	//if err != nil {
	//	log.Fatal("could not create memory profile: ", err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	StartAt := time.Now()
	FileName := config.Settings.FILENAME_GRAPHML
	log.Info("directory: ", config.Settings.DIRECTORY_SOURCE)
	log.Info("file graphml: ", FileName)
	ok := logic.StartFillAll(FileName)
	if ok == false {
		println(constants.TEXT_HELP)
	}

	log.Info("Time passed: ", time.Since(StartAt))

}
