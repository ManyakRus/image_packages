package packages_folder

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

type PackageFolder struct {
	Package  *packages.Package
	PkgPath  string
	Name     string
	FileName string
	Imports  []*packages.Package
}

type FoldersPackages struct {
	FoldersPackage *FoldersPackages
	Folder         folders.Folder
	PackageName    string
	PackageImports []string
}

func CreateConfigPackages(dir string) *packages.Config {
	var err error
	ctxMain := contextmain.GetContext()
	cfg := &packages.Config{}
	cfg.Context = ctxMain
	cfg.Dir = dir
	cfg.Mode = packages.NeedImports + packages.NeedName + packages.NeedExportFile + packages.NeedFiles
	cfg.Tests = false
	if err != nil {
		log.Panic("FindAllFolders_FromDir() error: ", err)
	}

	return cfg
}

func FindAllFolders_FromDir(dir string) *folders.Folder {

	FolderRoot := folders.FindFoldersTree(dir, true, false, false, "vendor")

	//FoldersPackage := FoldersPackages{}
	//FoldersPackage := FillFoldersPackages(FolderRoot, cfg)

	return FolderRoot
}

func FindRepositoryName(FolderRoot *folders.Folder) string {
	var Otvet string

	ConfigPackages := CreateConfigPackages(FolderRoot.FileName)

	MassPackages, _ := packages.Load(ConfigPackages)
	for _, v := range MassPackages {
		Otvet = v.Name
		break
	}

	if Otvet == "." {
		Otvet = ""
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
//			FoldersPackage1.PackageName = v.PkgPath
//			PackageImports := make([]string, 0)
//
//			RepositoryLen := len(RepositoryName)
//			for _, import1 := range v.Imports {
//				ImportID := import1.PkgPath
//				if len(ImportID) < RepositoryLen {
//					continue
//				}
//				if ImportID[0:RepositoryLen] != RepositoryName {
//					continue
//				}
//				log.Print("add package: ", import1.PkgPath)
//				PackageImports = append(PackageImports, import1.PkgPath)
//			}
//			FoldersPackage1.PackageImports = PackageImports
//		}
//		//FoldersPackage = append(FoldersPackage, FoldersPackage1)
//	}
//
//	return FoldersPackage
//}

func FindPackageFromFolder(FolderRoot *folders.Folder) PackageFolder {
	Otvet := PackageFolder{}
	RepositoryName := FindRepositoryName(FolderRoot)
	Otvet.Name = RepositoryName
	Otvet.FileName = FolderRoot.FileName

	//PackageImports := make([]*folders.Folder, 0)

	ConfigPackages := CreateConfigPackages(FolderRoot.FileName)
	MassPackages, _ := packages.Load(ConfigPackages)
	//for _, v := range MassPackages {
	//	Otvet.PkgPath = v.PkgPath
	//	RepositoryLen := len(RepositoryName)
	//	for _, import1 := range v.Imports {
	//		ImportID := import1.ID
	//		if len(ImportID) < RepositoryLen {
	//			continue
	//		}
	//		if ImportID[0:RepositoryLen] != RepositoryName {
	//			continue
	//		}
	//		//log.Print("add package: ", import1.PkgPath)
	//		//PackageFolder1 := PackageFolder{}
	//		PackageImports = append(PackageImports, import1)
	//	}
	//}

	for _, v := range MassPackages {
		Otvet.Package = v
		break
	}
	Otvet.Imports = MassPackages

	return Otvet
}
