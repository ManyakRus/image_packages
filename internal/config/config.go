package config

import (
	"os"
)

const FILENAME_XGML = "packages.xgml"

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// NeedReconnect - флаг необходимости переподключения
var NeedReconnect bool

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
	FILENAME_XGML    string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")
	Settings.FILENAME_XGML = os.Getenv("FILENAME_XGML")

	if Settings.DIRECTORY_SOURCE == "" {
		Settings.DIRECTORY_SOURCE = CurrentDirectory()
		//log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	if Settings.FILENAME_XGML == "" {
		Settings.FILENAME_XGML = FILENAME_XGML
	}

	//
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
	if len(Args) != 2 {
		return
	}

	if len(Args) > 0 {
		Settings.DIRECTORY_SOURCE = Args[0]
	}
	if len(Args) > 1 {
		Settings.FILENAME_XGML = Args[1]
	}
}
