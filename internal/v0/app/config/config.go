package config

import (
	"log"
	"os"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// NeedReconnect - флаг необходимости переподключения
var NeedReconnect bool

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")

	if Settings.DIRECTORY_SOURCE == "" {
		log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	//
}
