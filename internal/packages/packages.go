package packages

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

type FoldersPackages struct {
	*folders.Folder
	PackageName    string
	PackageImports []string
}

func FindAllPackages_FromDir(dir string) {
	var err error
	ctxMain := contextmain.GetContext()
	cfg := &packages.Config{}
	cfg.Context = ctxMain
	cfg.Dir = dir
	cfg.Mode = packages.NeedImports
	cfg.Tests = false
	if err != nil {
		log.Panic("FindAllPackages_FromDir() error: ", err)
	}

	//err = filepath.Walk(dir, FindImports_FromPackage)

	FolderAll := folders.FindFoldersTree(dir, true, false, false, "vendor")

	RepositoryName := ""

	for _, dir1 := range FolderAll {
		cfg.Dir = dir1
		MassPackages, _ := packages.Load(cfg)
		for _, v := range MassPackages {
			if RepositoryName == "" {
				RepositoryName = v.ID
			}
			RepositoryLen := len(RepositoryName)
			for _, import1 := range v.Imports {
				ImportID := import1.ID
				if len(ImportID) < RepositoryLen {
					continue
				}
				if ImportID[0:RepositoryLen] != RepositoryName {
					continue
				}
				log.Print(import1)
			}
		}
	}
}
