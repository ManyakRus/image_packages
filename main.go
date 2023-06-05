package main

import (
	"github.com/ManyakRus/image_packages/internal/v0/app/config"
	"github.com/ManyakRus/image_packages/internal/v0/app/packages"
	ConfigMain "github.com/ManyakRus/starter/config"
)

func main() {
	ConfigMain.LoadEnv()
	config.FillSettings()

	packages.FindAllPackages_FromDir(config.Settings.DIRECTORY_SOURCE)
}
