package packages_folder

import (
	"github.com/ManyakRus/image_packages/internal/parse_go"
	//"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

type Import struct {
	Name string //имя в коде
	URL  string
	//ID string //имя рассчитанное
}

// PackageFolder - содержит информацию о пакете
type PackageFolder struct {
	Package  *packages.Package
	PkgPath  string
	Name     string
	FileName string
	Imports  map[string]Import
	GoFiles  []string
}

// FoldersPackages - объединяет в себе информацию о папке и пакете, рекурсивно
type FoldersPackages struct {
	FoldersPackage *FoldersPackages
	Folder         folders.Folder
	PackageName    string
	//PackageImports []string
}

//func CreateConfigPackages(dir string) *packages.Config {
//	var err error
//	ctxMain := contextmain.GetContext()
//	cfg := &packages.Config{}
//	cfg.Context = ctxMain
//	cfg.Dir = dir
//	cfg.Mode = packages.NeedImports + packages.NeedName + packages.NeedFiles
//	//cfg.Mode = packages.NeedImports + packages.NeedName + packages.NeedExportFile + packages.NeedFiles
//	cfg.Tests = false
//	if err != nil {
//		log.Panic("FindAllFolders_FromDir() error: ", err)
//	}
//
//	return cfg
//}

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
func FindPackageFromFolder(Folder *folders.Folder) PackageFolder {
	Otvet := PackageFolder{}
	Otvet.FileName = Folder.FileName

	FolderName := Folder.FileName

	ok, err := micro.FileExists(FolderName)
	if err != nil {
		log.Error("FileExists() error: ", err)
		return Otvet
	}
	if ok == false {
		return Otvet
	}

	//заполним GoFiles
	files, err := os.ReadDir(FolderName)
	if err != nil {
		log.Error("ReadDir() error: ", err)
	}

	GoFiles := make([]string, 0)
	Imports := make(map[string]Import)

	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".go") {
			Filename := FolderName + micro.SeparatorFile() + v.Name()
			GoFiles = append(GoFiles, Filename)

			Imports = FindImports_FromFile(Filename, Imports)
		}
	}

	PackageName := ""
	if len(GoFiles) > 0 {
		PackageName = parse_go.FindLastWordImport(FolderName)
	}

	Otvet.GoFiles = GoFiles
	Otvet.Name = PackageName
	Otvet.Imports = Imports

	//FolderName := FolderRoot.FileName
	//
	//Otvet.Name = FolderName
	//Package1 := CreatePackageFromFolder(FolderRoot)
	//Otvet.Package = Package1
	//return Otvet
	//
	//RepositoryName := "" //FindRepositoryName(FolderRoot)
	//ConfigPackages := CreateConfigPackages(FolderName)
	//MassPackages, _ := packages.Load(ConfigPackages)
	//
	//for _, v := range MassPackages {
	//	Otvet.Package = v
	//	RepositoryName = v.Name
	//	break
	//}
	//Otvet.Name = RepositoryName
	////Otvet.Imports = MassPackages

	return Otvet
}

// CreatePackageFromFolder - ускорение packages.Load()
func CreatePackageFromFolder(Folder *folders.Folder) *packages.Package {
	Otvet := &packages.Package{}

	return Otvet
}

func FindImports_FromFile(Filename string, Imports map[string]Import) map[string]Import {
	//Otvet := make(map[string]Import)

	//ищем Imports
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parse_go.ParseFile(Filename)
	//f, err := parser.ParseFile(fset, Filename, nil, parser.ImportsOnly)
	if err != nil {
		log.Error("ParseFile() error: ", err)
		return Imports
	}

	AstSpec := astutil.Imports(fset, f)
	for _, group1 := range AstSpec {
		for _, import1 := range group1 {
			URL := import1.Path.Value
			URL = parse_go.DeleteQuotes(URL)

			Import1 := Import{}
			Name := ""
			NameID := import1.Name
			if NameID != nil {
				Name = NameID.Name
			} else {
				Name = parse_go.FindImportID_from_URL(URL)
			}
			Import1.Name = Name
			Import1.URL = URL
			Imports[Name] = Import1
		}
	}

	return Imports
}
