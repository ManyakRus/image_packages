package main

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/packages"
	ConfigMain "github.com/ManyakRus/starter/config"
)

func main() {
	ConfigMain.LoadEnv()
	config.FillSettings()

	packages.FindAllPackages_FromDir(config.Settings.DIRECTORY_SOURCE)
}
