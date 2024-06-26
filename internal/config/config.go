package config

import (
	"github.com/ManyakRus/starter/log"
	"os"
)

const FILENAME_XGML = "packages.graphml"

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
	FILENAME_GRAPHML string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//сменим текущую директорию чтоб работали локальные пути
	ChangeCurrentDirectory()

	//
	Settings = SettingsINI{}
	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")
	Settings.FILENAME_GRAPHML = os.Getenv("FILENAME_GRAPHML")

	if Settings.DIRECTORY_SOURCE == "" {
		Settings.DIRECTORY_SOURCE = CurrentDirectory()
		//log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	if Settings.FILENAME_GRAPHML == "" {
		Settings.FILENAME_GRAPHML = FILENAME_XGML
	}

	//
}

// ChangeCurrentDirectory - устанавливает текущую директорию на директорию откуда запущена программа
// вместо директории где находится программа
func ChangeCurrentDirectory() {
	var err error

	// сменим директорию на текущую
	dir := CurrentDirectory()
	err = os.Chdir(dir)
	if err != nil {
		log.Error("Chdir error: ", err)
	} else {
		log.Info("Chdir: ", dir)
	}

}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}

// FillFlags - заполняет параметры из командной строки
func FillFlags() {
	Args := os.Args[1:]
	if len(Args) > 2 {
		return
	}

	if len(Args) > 0 {
		Settings.DIRECTORY_SOURCE = Args[0]
	}
	if len(Args) > 1 {
		Settings.FILENAME_GRAPHML = Args[1]
	}
}
