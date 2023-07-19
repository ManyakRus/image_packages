package main

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config"
)

func main() {
	ConfigMain.LoadEnv()
	config.FillSettings()

	logic.StartFillAll()

}
