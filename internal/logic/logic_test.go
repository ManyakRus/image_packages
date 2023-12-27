package logic

import (
	"github.com/ManyakRus/image_packages/internal/config"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestStartFillAll(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	dir := micro.ProgramDir()
	FileName := dir + "test" + micro.SeparatorFile() + "test_start.xgml"
	StartFillAll(FileName)
}
