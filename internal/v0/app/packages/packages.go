package packages

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

func FindAllPackages_FromDir(dir string) {
	ctxMain := contextmain.GetContext()
	cfg := &packages.Config{}
	cfg.Context = ctxMain
	cfg.Dir = dir
	cfg.Mode = packages.NeedImports
	MassPackages, err := packages.Load(cfg)
	if err != nil {
		log.Panic("FindAllPackages_FromDir() error: ", err)
	}

	for _, v := range MassPackages {
		log.Print(v)
	}
}
