package packages

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

type Package struct {
	Name    string
	Imports []string
}

type FoldersPackages struct {
	FoldersPackage *FoldersPackages
	Folder         folders.Folder
	PackageName    string
	PackageImports []string
}

func FindAllFolders_FromDir(dir string) *folders.Folder {
	var err error
	ctxMain := contextmain.GetContext()
	cfg := &packages.Config{}
	cfg.Context = ctxMain
	cfg.Dir = dir
	cfg.Mode = packages.NeedImports
	cfg.Tests = false
	if err != nil {
		log.Panic("FindAllFolders_FromDir() error: ", err)
	}

	FolderRoot := folders.FindFoldersTree(dir, true, false, false, "vendor")

	//FoldersPackage := FoldersPackages{}
	//FoldersPackage := FillFoldersPackages(FolderRoot, cfg)

	return FolderRoot
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

//func FillFoldersPackages(FolderRoot *folders.Folder, cfg *packages.Config) FoldersPackages {
//	RepositoryName := FindRepositoryName(FolderRoot, cfg)
//
//	FoldersPackage := FoldersPackages{}
//
//	for _, folder1 := range FolderRoot.Folders {
//		FoldersPackage1 := FoldersPackages{}
//		FoldersPackage1.Folder = *folder1
//
//		cfg.Dir = folder1.Name
//		MassPackages, _ := packages.Load(cfg)
//		for _, v := range MassPackages {
//			FoldersPackage1.PackageName = v.ID
//			PackageImports := make([]string, 0)
//
//			RepositoryLen := len(RepositoryName)
//			for _, import1 := range v.Imports {
//				ImportID := import1.ID
//				if len(ImportID) < RepositoryLen {
//					continue
//				}
//				if ImportID[0:RepositoryLen] != RepositoryName {
//					continue
//				}
//				log.Print("add package: ", import1.ID)
//				PackageImports = append(PackageImports, import1.ID)
//			}
//			FoldersPackage1.PackageImports = PackageImports
//		}
//		//FoldersPackage = append(FoldersPackage, FoldersPackage1)
//	}
//
//	return FoldersPackage
//}

func FindPackageFromFolder(FolderRoot *folders.Folder, cfg *packages.Config) Package {
	Otvet := Package{}
	RepositoryName := FindRepositoryName(FolderRoot, cfg)
	Otvet.Name = RepositoryName

	PackageImports := make([]string, 0)

	cfg.Dir = FolderRoot.Name
	MassPackages, _ := packages.Load(cfg)
	for _, v := range MassPackages {

		RepositoryLen := len(RepositoryName)
		for _, import1 := range v.Imports {
			ImportID := import1.ID
			if len(ImportID) < RepositoryLen {
				continue
			}
			if ImportID[0:RepositoryLen] != RepositoryName {
				continue
			}
			//log.Print("add package: ", import1.ID)
			PackageImports = append(PackageImports, import1.ID)
		}
	}

	Otvet.Imports = PackageImports

	return Otvet
}
