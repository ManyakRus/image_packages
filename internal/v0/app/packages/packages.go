package packages

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

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

	Dirs, err := FilePathWalkDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	RepositoryName := ""

	for _, dir1 := range Dirs {
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

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			return nil
		}
		if len(info.Name()) >= 1 && info.Name()[0:1] == "." {
			return filepath.SkipDir
		}
		if len(info.Name()) >= 6 && info.Name()[0:6] == "vendor" {
			return filepath.SkipDir
		}
		files = append(files, path)
		return nil
	})
	return files, err
}
