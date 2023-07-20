package logic

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/packages_folder"
	"github.com/ManyakRus/image_packages/pkg/xgml"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/beevik/etree"
	"path/filepath"
	"sort"
)

func StartFillAll(FileName string) {
	FolderRoot := packages_folder.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)

	//var buffer *strings.Builder

	//xgml.AddDirectory(buffer, FolderRoot.Name)

	DocXML := xgml.CreateDocXGML("")
	ElementGraph := DocXML.FindElement("/section/section")

	FillFolder(ElementGraph, nil, FolderRoot)

	DocXML.IndentTabs()
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		log.Error("WriteToFile() FileName: ", FileName, " error: ", err)
	}

}

func FillFolder(ElementGraph, ElementGroup *etree.Element, Folder *folders.Folder) {
	//добавим группа (каталог)
	FolderName := Folder.Name

	//добавим пакет(package)
	//ConfigPackages := packages_folder.CreateConfigPackages(Folder.FileName)
	PackageNameFull := packages_folder.FindRepositoryName(Folder)
	PackageName := FindFileNameShort(PackageNameFull)
	if PackageName == "" && len(Folder.Folders) == 0 {
		return
	}

	ElementGroup = xgml.CreateGroupXGML(ElementGraph, ElementGroup, FolderName)
	if PackageName != "" {
		xgml.CreateElementXGML_Shape(ElementGraph, ElementGroup, PackageName)
	}

	//
	MassKeys := make([]string, 0, len(Folder.Folders))
	for k := range Folder.Folders {
		MassKeys = append(MassKeys, k)
	}
	sort.Strings(MassKeys)

	for _, key1 := range MassKeys {
		Folder1, ok := Folder.Folders[key1]
		if ok == false {
			log.Panic("Folder.Folders[key1] ok =false")
		}
		FillFolder(ElementGraph, ElementGroup, Folder1)
	}
}

// FindFileNameShort - возвращает имя файла(каталога) без пути
func FindFileNameShort(path string) string {
	Otvet := ""
	if path == "" {
		return Otvet
	}
	Otvet = filepath.Base(path)

	return Otvet
}
