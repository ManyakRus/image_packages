package parse_go

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestParseDir(t *testing.T) {
	Dir := micro.ProgramDir()
	Dir = Dir + "internal"
	ParseDir(Dir)
}

func TestParseFile(t *testing.T) {
	Filename := micro.ProgramDir() + "internal/main.go"
	ParseFile(Filename)
}
