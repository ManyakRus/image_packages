package logic

import (
	"bytes"
	"github.com/ManyakRus/image_packages/internal/config"
	"github.com/ManyakRus/image_packages/internal/packages_folder"
	"github.com/ManyakRus/image_packages/internal/parse_go"
	"github.com/ManyakRus/image_packages/pkg/graphml"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/beevik/etree"
	"golang.org/x/tools/go/packages"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// MapPackagesElements - связь Пакета golang / Элемент файла .graphml
var MapPackagesElements = make(map[*packages.Package]*etree.Element, 0)

// MapPackageIDElements - связь ИД Пакета golang / Элемент файла .graphml
var MapPackageIDElements = make(map[string]*etree.Element, 0)

// CountLinesFunctions - количество строк и количество функций
type CountLinesFunctions struct {
	LinesCount int
	FuncCount  int
}

// FindLinesCount_Cache - кэш рассчитанных количество строк и количество функций
var FindLinesCount_Cache = make(map[string]CountLinesFunctions)

// FindModuleFuncCallCount_Cache - кэш количества вызовов функций, для ускорения
var FindModuleFuncCallCount_Cache = make(map[string]int, 0)

// StartFillAll - Старт работы приложения
func StartFillAll(FileName string) bool {
	Otvet := false

	FolderRoot := packages_folder.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)
	if FolderRoot == nil {
		log.Error("Error: not found folder: ", FolderRoot)
		return Otvet
	}

	//var buffer *strings.Builder

	//graphml.AddDirectory(buffer, FolderRoot.Name)

	DocXML, ElementGraph := graphml.CreateDocument()

	var StartAt time.Time
	//заполним каталоги и пакеты
	StartAt = time.Now()
	//log.Info("Start fill groups")
	FillFolder(ElementGraph, nil, FolderRoot)
	log.Info("FillFolder() Time passed: ", time.Since(StartAt))

	//заполним связи
	StartAt = time.Now()
	//log.Info("Start fill links")
	FillLinks(ElementGraph)
	log.Info("FillLinks() Time passed: ", time.Since(StartAt))

	//заполним связи горутин
	StartAt = time.Now()
	//log.Info("Start fill goroutine links")
	FillLinks_goroutine(ElementGraph)
	log.Info("FillLinks_goroutine() Time passed: ", time.Since(StartAt))

	if len(MapPackagesElements) > 0 {
		Otvet = true
	}

	if Otvet == false {
		println("warning: Empty file not saved !")
		return Otvet
	}

	log.Info("Start save file")
	//DocXML.IndentTabs()
	DocXML.Indent(2)
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		log.Error("WriteToFile() FileName: ", FileName, " error: ", err)
	}

	return Otvet
}

// FillLinks - заполняет связи (стрелки) между пакетами
func FillLinks(ElementGraph *etree.Element) {
	for PackageFrom, ElementFrom := range MapPackagesElements {
		for _, PackageImport := range PackageFrom.Imports {
			ElementImport, ok := MapPackageIDElements[PackageImport.ID]
			if ok == false {
				//посторонние импорты
				continue
			}
			descr := PackageFrom.Name + " -> " + PackageImport.Name
			call_count := FindModuleFuncCallCount(PackageFrom, PackageImport)
			if call_count > 0 {
				//линия
				graphml.CreateElement_Edge(ElementGraph, ElementFrom, ElementImport, "", descr)
			} else {
				//пунктиром
				graphml.CreateElement_Edge_dotted(ElementGraph, ElementFrom, ElementImport, "", descr)
			}
		}
	}
}

// FillLinks_goroutine - заполняет связи (стрелки) между пакетами для горутин go, синим цветом
func FillLinks_goroutine(ElementGraph *etree.Element) {
	for PackageFrom, ElementFrom := range MapPackagesElements {
		for _, Filename1 := range PackageFrom.GoFiles {

			AstFile, err := parse_go.ParseFile(Filename1)
			if err != nil {
				log.Warn("ParseFile() ", Filename1, " error: ", err)
				continue
			}

			MassGoImport := parse_go.FindGo(AstFile)
			for _, GoImport1 := range MassGoImport {
				//Go_package_name := GoImport1.Go_package_name
				Go_package_import := GoImport1.Go_package_import
				Go_func_name := GoImport1.Go_func_name

				ElementImport, ok := MapPackageIDElements[Go_package_import]
				if ok == false {
					//посторонние импорты
					//continue
					ElementImport = ElementFrom
				}
				label := Go_func_name
				descr := PackageFrom.Name + " -> " + GoImport1.Go_package_name
				graphml.CreateElement_Edge_blue(ElementGraph, ElementFrom, ElementImport, label, descr)
			}

		}
	}
}

// FindModuleFuncCallCount - находит количество вызовов функций нужного модуля
func FindModuleFuncCallCount(PackageFrom, PackageTo *packages.Package) int {
	Otvet := 0

	sID_Packages_from_to := PackageFrom.ID + "_" + PackageTo.ID
	Ovvet, isFinded := FindModuleFuncCallCount_Cache[sID_Packages_from_to]
	if isFinded == true {
		return Ovvet
	}

	PackageTo_ID := PackageTo.ID

	for _, Filename1 := range PackageFrom.GoFiles {

		AstFile, err := parse_go.ParseFile(Filename1)
		if err != nil {
			log.Warn("ParseFile() ", Filename1, " error: ", err)
			continue
		}

		MassGoImport := parse_go.FindFunctions(AstFile)
		for _, GoImport1 := range MassGoImport {
			//	//Go_package_name := GoImport1.Go_package_name
			Go_package_import := GoImport1.Go_package_import
			if Go_package_import == PackageTo_ID {
				Otvet = Otvet + 1
			}
			//	Go_func_name := GoImport1.Go_func_name
			//
			//	ElementImport, ok := MapPackageIDElements[Go_package_import]
			//	if ok == false {
			//		//посторонние импорты
			//		//continue
			//		ElementImport = ElementFrom
			//	}
			//	label := Go_func_name
			//	descr := PackageFrom.Name + " -> " + GoImport1.Go_package_name
			//	graphml.CreateElement_Edge_blue(ElementGraph, ElementFrom, ElementImport, label, descr)
		}

	}

	FindModuleFuncCallCount_Cache[sID_Packages_from_to] = Otvet

	return Otvet
}

func FillFolder(ElementGraph, ElementGroup *etree.Element, Folder *folders.Folder) {

	FolderName := Folder.Name

	//ConfigPackages := packages_folder.CreateConfigPackages(Folder.FileName)
	PackageFolder1 := packages_folder.FindPackageFromFolder(Folder)
	PackageName := PackageFolder1.Name
	//PackageNameFull := PackageFolder1.
	//PackageName := FindFileNameShort(PackageNameFull)
	if PackageName == "" && len(Folder.Folders) == 0 {
		return
	}

	GroupName := FolderName
	lines_count, func_count := FindLinesCount_package(PackageFolder1.Package)
	if lines_count > 0 || func_count > 0 {
		GroupName = GroupName + " ("
		if func_count > 0 {
			GroupName = GroupName + strconv.Itoa(func_count) + " func"
			GroupName = GroupName + ", "
		}
		if lines_count > 0 {
			GroupName = GroupName + strconv.Itoa(lines_count) + " lines"
		}
		GroupName = GroupName + ")"
	}

	//добавим группа (каталог)
	var ElementGroup2 *etree.Element
	if ElementGroup != nil {
		ElementGroup2 = graphml.CreateElement_Group(ElementGroup, GroupName)
	} else {
		ElementGroup2 = graphml.CreateElement_Group(ElementGraph, GroupName)
	}
	if PackageName != "" {
		//добавим пакет(package)
		ElementShape := graphml.CreateElement_Shape(ElementGroup2, PackageName)
		MapPackagesElements[PackageFolder1.Package] = ElementShape
		MapPackageIDElements[PackageFolder1.Package.ID] = ElementShape
		//MapPackagesElements[&PackageFolder1] = ElementShape
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
		FillFolder(ElementGraph, ElementGroup2, Folder1)
	}

}

//// FindFileNameShort - возвращает имя файла(каталога) без пути
//func FindFileNameShort(path string) string {
//	Otvet := ""
//	if path == "" {
//		return Otvet
//	}
//	Otvet = filepath.Base(path)
//
//	return Otvet
//}

func FindLinesCount_package(Package1 *packages.Package) (int, int) {
	LinesCount := 0
	FuncCount := 0

	if Package1 == nil {
		return 0, 0
	}

	if Package1.GoFiles == nil {
		return 0, 0
	}

	for _, s := range Package1.GoFiles {
		count, func_count := FindLinesCount(s)
		LinesCount = LinesCount + count
		FuncCount = FuncCount + func_count
	}

	return LinesCount, FuncCount
}

// FindLinesCount - возвращает количество строк и количество функций в файле
func FindLinesCount(FileName string) (int, int) {
	LinesCount := 0
	FuncCount := 0

	//
	CountLinesFunctions1, isFinded := FindLinesCount_Cache[FileName]
	if isFinded == true {
		return CountLinesFunctions1.LinesCount, CountLinesFunctions1.FuncCount
	}

	//
	bytes1, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bytes1)
	LinesCount, err = LinesCount_reader(reader)
	if err != nil {
		log.Fatal(err)
	}

	FuncCount = FindFuncCount(&bytes1)

	//
	FindLinesCount_Cache[FileName] = CountLinesFunctions{
		LinesCount: LinesCount,
		FuncCount:  FuncCount,
	}

	return LinesCount, FuncCount
}

// LinesCount_reader - возвращает количество строк в файле
func LinesCount_reader(r io.Reader) (int, error) {
	defaultSize := 1024
	defaultEndLine := "\n"

	Size := defaultSize
	Sep := defaultEndLine

	buf := make([]byte, Size)
	var count int

	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], []byte(Sep))

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}

	}
}

//func LinesCount_reader(r io.Reader) (int, error) {
//	Otvet := 0
//	var err error
//
//	buf := make([]byte, 8192)
//
//	for {
//		c, err := r.Read(buf)
//		if err != nil {
//			if err == io.EOF && c == 0 {
//				break
//			} else {
//				return Otvet, err
//			}
//		}
//
//		for _, b := range buf[:c] {
//			if b == '\n' {
//				Otvet++
//			}
//		}
//	}
//
//	if err == io.EOF {
//		err = nil
//	}
//
//	return Otvet, err
//}

// FindFuncCount - находит количество функций(func) в файле
func FindFuncCount(bytes0 *[]byte) int {
	Otvet := 0

	//s := string(*bytes0)
	//Otvet = strings.Count(s, "\nfunc ")

	//sFind := "(\n|\t| )func( |\t)"
	//
	//Otvet = CountMatches(s, regexp.MustCompile(sFind))

	Otvet = bytes.Count(*bytes0, []byte("\nfunc "))

	return Otvet
}

// CountMatches - находит количество совпадений в regexp
func CountMatches(s string, re *regexp.Regexp) int {
	total := 0
	for start := 0; start < len(s); {
		remaining := s[start:] // slicing the string is cheap
		loc := re.FindStringIndex(remaining)
		if loc == nil {
			break
		}
		// loc[0] is the start index of the match,
		// loc[1] is the end index (exclusive)
		start += loc[1]
		total++
	}
	return total
}
