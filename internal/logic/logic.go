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

// MapPackagesElements - содержит индекс элемента xgml для каждого пакета
var MapPackagesElements = make(map[*packages_folder.PackageFolder]*etree.Element, 0)

func StartFillAll(FileName string) {
	FolderRoot := packages_folder.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)

	//var buffer *strings.Builder

	//xgml.AddDirectory(buffer, FolderRoot.Name)

	DocXML := xgml.CreateDocXGML("")
	ElementGraph := DocXML.FindElement("/section/section")

	//заполним каталоги и пакеты
	FillFolder(ElementGraph, nil, FolderRoot)

	//заполним связи
	//FillLinks(ElementGraph)

	DocXML.IndentTabs()
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		log.Error("WriteToFile() FileName: ", FileName, " error: ", err)
	}

}

// FillLinks - заполняет связи (стрелки) между пакетами
func FillLinks(ElementGraph *etree.Element) {
	for PackageFrom, ElementFrom := range MapPackagesElements {
		for _, PackageImport := range PackageFrom.Imports {
			ElementImport, ok := MapPackagesElements[PackageImport]
			if ok == false {
				log.Panic("MapPackagesElements[PackageImport] error: ok =false")
			}
		}
	}
}

func FillFolder(ElementGraph, ElementGroup *etree.Element, Folder *folders.Folder) {
	FolderName := Folder.Name

	//ConfigPackages := packages_folder.CreateConfigPackages(Folder.FileName)
	Package1 := packages_folder.FindPackageFromFolder(Folder)
	PackageName := Package1.Name
	//PackageNameFull := Package1.
	//PackageName := FindFileNameShort(PackageNameFull)
	if PackageName == "" && len(Folder.Folders) == 0 {
		return
	}

	//добавим группа (каталог)
	ElementGroup = xgml.CreateGroupXGML(ElementGraph, ElementGroup, FolderName)
	if PackageName != "" {
		//добавим пакет(package)
		ElementShape := xgml.CreateElementXGML_Shape(ElementGraph, ElementGroup, PackageName)
		MapPackagesElements[&Package1] = ElementShape
	}

	//сортировка
	MassKeys := make([]string, 0, len(Folder.Folders))
	for k := range Folder.Folders {
		MassKeys = append(MassKeys, k)
	}
	sort.Strings(MassKeys)

	//обход всех папок
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
