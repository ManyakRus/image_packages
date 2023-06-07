package packages

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

type FoldersPackages struct {
	Folder         *folders.Folder
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

	FolderRoot := folders.FindFoldersTree(dir, true, false, false, "vendor")

	FoldersPackage := FoldersPackages{}
	FillFoldersPackages(&FoldersPackage, FolderRoot, cfg)

}

func FindRepositoryName(FolderRoot *folders.Folder, cfg *packages.Config) string {
	var Otvet string

	dir := FolderRoot.Name
	cfg.Dir = dir
	MassPackages, _ := packages.Load(cfg)
	for _, v := range MassPackages {
		Otvet = v.ID
		break
	}

	return Otvet
}

// изменяет FoldersPackage
func FillFoldersPackages(FoldersPackage *FoldersPackages, FolderRoot *folders.Folder, cfg *packages.Config) {
	RepositoryName := FindRepositoryName(FolderRoot, cfg)

	for _, folder1 := range FolderRoot.Folders {
		FoldersPackage1 := &FoldersPackages{}
		FoldersPackage1.Folder = folder1

		cfg.Dir = folder1.Name
		MassPackages, _ := packages.Load(cfg)
		for _, v := range MassPackages {
			FoldersPackage1.PackageName = v.ID
			PackageImports := make([]string, 0)

			RepositoryLen := len(RepositoryName)
			for _, import1 := range v.Imports {
				ImportID := import1.ID
				if len(ImportID) < RepositoryLen {
					continue
				}
				if ImportID[0:RepositoryLen] != RepositoryName {
					continue
				}
				log.Print("add package: ", import1.ID)
				PackageImports = append(PackageImports, import1.ID)
			}
			FoldersPackage1.PackageImports = PackageImports
		}
		*FoldersPackage = append(*FoldersPackage, FoldersPackage1)
	}

}
