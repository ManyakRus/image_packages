package logic

import (
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/packages"
	"github.com/ManyakRus/image_packages/pkg/xgml"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/beevik/etree"
	"sort"
)

func StartFillAll(FileName string) {
	FolderRoot := packages.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)

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
	FolderName := Folder.Name
	ElementGroup = xgml.CreateGroupXGML(ElementGraph, ElementGroup, FolderName)

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
