package packages_folder

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"golang.org/x/tools/go/packages"
)

// PackageFolder - содержит информацию о пакете
type PackageFolder struct {
	Package  *packages.Package
	PkgPath  string
	Name     string
	FileName string
	//Imports  []*packages.Package
}

// FoldersPackages - объединяет в себе информацию о папке и пакете, рекурсивно
type FoldersPackages struct {
	FoldersPackage *FoldersPackages
	Folder         folders.Folder
	PackageName    string
	//PackageImports []string
}

func CreateConfigPackages(dir string) *packages.Config {
	var err error
	ctxMain := contextmain.GetContext()
	cfg := &packages.Config{}
	cfg.Context = ctxMain
	cfg.Dir = dir
	cfg.Mode = packages.NeedImports + packages.NeedName + packages.NeedFiles
	//cfg.Mode = packages.NeedImports + packages.NeedName + packages.NeedExportFile + packages.NeedFiles
	cfg.Tests = false
	if err != nil {
		log.Panic("FindAllFolders_FromDir() error: ", err)
	}

	return cfg
}

// FindAllFolders_FromDir - возвращает дерево всех папок и файлов в директории
func FindAllFolders_FromDir(dir string) *folders.Folder {

	MassExclude := make([]string, 0)
	MassExclude = append(MassExclude, "vendor")
	MassExclude = append(MassExclude, ".git")
	MassExclude = append(MassExclude, ".idea")
	MassExclude = append(MassExclude, ".vscode")
	FolderRoot := folders.FindFoldersTree(dir, true, false, false, MassExclude)

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

// FindPackageFromFolder - возвращает пакет Package + Folder
func FindPackageFromFolder(FolderRoot *folders.Folder) PackageFolder {
	Otvet := PackageFolder{}
	Otvet.FileName = FolderRoot.FileName

	FolderName := FolderRoot.FileName

	//HasGoFiles := HasGoFiles(FolderRoot.FileName)
	//if HasGoFiles == false {
	//	return Otvet
	//}

	//HasGoFiles := HasGoFiles(FolderRoot)
	//if HasGoFiles == false {
	//	return Otvet
	//}

	RepositoryName := "" //FindRepositoryName(FolderRoot)
	ConfigPackages := CreateConfigPackages(FolderName)
	MassPackages, _ := packages.Load(ConfigPackages)

	for _, v := range MassPackages {
		Otvet.Package = v
		RepositoryName = v.Name
		break
	}
	Otvet.Name = RepositoryName
	//Otvet.Imports = MassPackages

	return Otvet
}

//// HasGoFilesSearch - возвращает true если в папке есть файлы с расширением .go
//func HasGoFilesSearch(FolderName string) bool {
//	Otvet := false
//
//	ok, err := micro.FileExists(FolderName)
//	if err != nil {
//		log.Error("FileExists() error: ", err)
//		return Otvet
//	}
//	if ok == false {
//		return Otvet
//	}
//
//	files, err := os.ReadDir(FolderName)
//	if err != nil {
//		log.Error("ReadDir() error: ", err)
//	}
//
//	for _, v := range files {
//		if strings.HasSuffix(v.Name(), ".go") {
//			Otvet = true
//			break
//		}
//	}
//
//	return Otvet
//}
//
//// HasGoFiles - возвращает true если в папке есть файлы с расширением .go
//func HasGoFiles(Folder *folders.Folder) bool {
//	Otvet := false
//
//	for _, v := range Folder.Files {
//		if strings.HasSuffix(v.Name, ".go") {
//			Otvet = true
//			break
//		}
//	}
//
//	return Otvet
//}
