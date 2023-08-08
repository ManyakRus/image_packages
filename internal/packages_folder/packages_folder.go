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

//func FindRepositoryName(FolderRoot *folders.Folder) string {
//	var Otvet string
//
//	ConfigPackages := CreateConfigPackages(FolderRoot.FileName)
//
//	MassPackages, _ := packages.Load(ConfigPackages)
//	for _, v := range MassPackages {
//		Otvet = v.Name
//		break
//	}
//
//	if Otvet == "." {
//		Otvet = ""
//	}
//
//	return Otvet
//}

func FindPackageFromFolder(FolderRoot *folders.Folder) PackageFolder {
	Otvet := PackageFolder{}

	RepositoryName := "" //FindRepositoryName(FolderRoot)
	ConfigPackages := CreateConfigPackages(FolderRoot.FileName)
	MassPackages, _ := packages.Load(ConfigPackages)

	for _, v := range MassPackages {
		Otvet.Package = v
		RepositoryName = v.Name
		break
	}
	Otvet.Name = RepositoryName
	Otvet.FileName = FolderRoot.FileName
	Otvet.Imports = MassPackages

	return Otvet
}
