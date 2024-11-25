package parse_go

import (
	"fmt"
	"github.com/ManyakRus/starter/micro"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"strings"
)

// GoImport - содержит информацию о вызове горутины go
type GoImport struct {
	Go_package_name   string //псевдоним импорта вызываемой функции из другого пакета
	Go_package_import string //полный путь импорта вызываемой функции из другого пакета
	Go_func_name      string //имя вызываемой функции
}

// FindFunctions_Cache - рассчитанный кэш, для ускорения, т.к. 1 файл считаем много раз
var FindFunctions_Cache = make(map[*ast.File][]GoImport)

// ParseFile_Cache - кэш пропарсенных файлов, для ускорения
var ParseFile_Cache = make(map[string]*ast.File)

// ParseDir - парсит все файлы .go, кроме тсетов
func ParseDir(Dir string) (map[string]*ast.Package, error) {

	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	MapPackages, err := parser.ParseDir(fset, Dir, filter_fn, parser.AllErrors+parser.Trace)
	if err != nil {
		fmt.Println(err)
		return MapPackages, err
	}

	//// Print the imports from the AST.
	//for _, v := range MapPackages {
	//	fmt.Println(v.Imports)
	//}

	return MapPackages, err
}

// filter_fn - проверяет имя файла чтоб это был не тестовый файл
func filter_fn(fi fs.FileInfo) bool {
	Otvet := true

	Filename := fi.Name()
	len1 := len(Filename)
	if len1 > 8 && Filename[len1-8-1:] == "_test.go" {
		Otvet = false
	}

	return Otvet
}

// ParseFile - парсит файл .go
func ParseFile(Filename string) (*ast.File, error) {
	var Otvet *ast.File
	var err error

	//поищем в кэш
	Otvet, isFinded := ParseFile_Cache[Filename]
	if isFinded == true {
		return Otvet, err
	}

	//
	fset := token.NewFileSet() // positions are relative to fset

	Otvet, err = parser.ParseFile(fset, Filename, nil, parser.AllErrors)
	if err != nil {
		//fmt.Println(err)
		return Otvet, err
	}

	//запомним в кэш
	ParseFile_Cache[Filename] = Otvet

	return Otvet, err
}

// FindGo - находит массив команд go (горутины)
func FindGo(AstFile *ast.File) []GoImport {
	Otvet := make([]GoImport, 0)

	if AstFile == nil {
		return Otvet
	}

	for _, decl1 := range AstFile.Decls {
		switch decl1.(type) {
		case *ast.FuncDecl:
			{
				func1 := decl1.(*ast.FuncDecl)
				body := func1.Body
				if body == nil {
					continue
				}
				for _, list1 := range body.List {
					switch list1.(type) {
					case *ast.GoStmt:
						{
							GoStmt1 := list1.(*ast.GoStmt)
							go_package_name, go_package_import, go_func_name := FindGoValues(AstFile, GoStmt1)
							GoImport1 := GoImport{}
							GoImport1.Go_package_name = go_package_name
							GoImport1.Go_package_import = go_package_import
							GoImport1.Go_func_name = go_func_name
							Otvet = append(Otvet, GoImport1)
						}
					}
				}
			}
		}
	}

	return Otvet
}

// FindGoImport_fromFunc - находит имя функции и её префикс импорт
func FindGoImport_fromFunc(AstFile *ast.File, SelectorExpr1 *ast.SelectorExpr) GoImport {
	Otvet := GoImport{}

	switch SelectorExpr1.X.(type) {
	case *ast.Ident:
		{
			Ident_X := SelectorExpr1.X.(*ast.Ident)
			Ident_Sel := SelectorExpr1.Sel

			go_package_name := Ident_X.Name
			go_func_name := Ident_Sel.Name
			go_package_import := FindPackageImport_FromName(AstFile, go_package_name)

			//GoImport1 := GoImport{}
			Otvet.Go_package_name = go_package_name
			Otvet.Go_package_import = go_package_import
			Otvet.Go_func_name = go_func_name
			//Otvet = append(Otvet, GoImport1)
		}
	case *ast.CallExpr:
		{
			CallExpr2 := SelectorExpr1.X.(*ast.CallExpr)
			switch CallExpr2.Fun.(type) {
			case *ast.SelectorExpr:
				{
					SelectorExpr2 := CallExpr2.Fun.(*ast.SelectorExpr)
					Otvet = FindGoImport_fromFunc(AstFile, SelectorExpr2)
				}
			}
		}
	default:
		{
			//log.Warnf("%#v", SelectorExpr1.X)
		}
	}

	return Otvet
}

// FindFunctions - находит массив команд go (горутины)
func FindFunctions(AstFile *ast.File) []GoImport {
	Otvet := make([]GoImport, 0)

	if AstFile == nil {
		return Otvet
	}

	Otvet, isFinded := FindFunctions_Cache[AstFile]
	if isFinded == true {
		return Otvet
	}

	fset := token.NewFileSet()
	visitor := &Visitor{fset: fset}
	visitor.MassGoImport = make([]GoImport, 0)
	visitor.AstFile = AstFile
	ast.Walk(visitor, AstFile)

	Otvet = visitor.MassGoImport
	FindFunctions_Cache[AstFile] = Otvet

	return Otvet
}

type Visitor struct {
	fset         *token.FileSet
	MassGoImport []GoImport
	AstFile      *ast.File
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	var SelectorExpr1 *ast.SelectorExpr

	switch x := n.(type) {
	case *ast.CallExpr:
		var ok bool
		SelectorExpr1, ok = x.Fun.(*ast.SelectorExpr)
		if ok == false {
			return v
		}
	default:
		return v
	}

	GoImport1 := FindGoImport_fromFunc(v.AstFile, SelectorExpr1)
	v.MassGoImport = append(v.MassGoImport, GoImport1)

	return v
}

//// parseFuncNode - возвращает Истина если это функция
//func parseFuncNode(n ast.Node) bool {
//	Otvet := true
//
//	switch n.(type) {
//	case *ast.FuncDecl:
//	}
//
//	return Otvet
//}

// FindGoValues - находит параметры функции: имя пакета, имя испорта, имя функции
func FindGoValues(AstFile *ast.File, GoStmt1 *ast.GoStmt) (go_package_name, go_package_import, go_func_name string) {

	iFunc1 := GoStmt1.Call.Fun
	switch iFunc1.(type) {
	case *ast.SelectorExpr:
		{
			func1 := iFunc1.(*ast.SelectorExpr)
			X := func1.X.(*ast.Ident)
			go_package_name = X.Name

			Sel := func1.Sel
			go_func_name = Sel.Name + "()"

			go_package_import = FindPackageImport_FromName(AstFile, go_package_name)
		}
	case *ast.Ident:
		{
			func1 := iFunc1.(*ast.Ident)
			go_func_name = func1.Name + "()"
		}
	}

	return
}

// FindPackageImport_FromName - находит имя пакета импорта из полного пути URL
func FindPackageImport_FromName(AstFile *ast.File, go_package_name string) string {
	Otvet := ""

	//поиск псевдонима импорта
	for _, import1 := range AstFile.Imports {
		//если задан псевдоним импорта
		Name := ""
		iName := import1.Name
		if iName != nil {
			Name = iName.String()
		}
		if Name == go_package_name {
			Otvet = import1.Path.Value
			break
		}

	}

	//поиск импорта без псевдонима
	for _, import1 := range AstFile.Imports {
		//импорт с псевдонимом пропускаем
		Name := ""
		iName := import1.Name
		if iName != nil {
			Name = iName.String()
		}
		if Name != "" {
			continue
		}

		//импорт без псевдонима - последнее слово в строке
		ImportString := micro.Trim(import1.Path.Value)
		last_word := FindLastWordImport(ImportString)
		if last_word == go_package_name {
			Otvet = import1.Path.Value
			Otvet = DeleteQuotes(Otvet)
			break
		}
	}

	return Otvet
}

// FindImportID_from_URL - находит ID из импорта URL
func FindImportID_from_URL(URL string) string {
	Otvet := ""
	ImportString := micro.Trim(URL)
	Otvet = FindLastWordImport(ImportString)

	pos1 := strings.Index(Otvet, ".")
	if pos1 >= 0 {
		Otvet = Otvet[0:pos1]
	}

	return Otvet
}

// FindLastWordImport - находит последнее слово в URL
func FindLastWordImport(ImportString string) string {
	Otvet := ImportString
	Otvet = DeleteQuotes(Otvet)

	pos1 := strings.LastIndex(Otvet, "/")
	if pos1 < 0 {
		return Otvet
	}
	if len(Otvet) < (pos1 + 1 + 1) {
		return Otvet
	}
	Otvet = Otvet[pos1+1:]

	return Otvet
}

// DeleteQuotes - удаляет все кавычки
func DeleteQuotes(s string) string {
	Otvet := s

	Otvet = strings.ReplaceAll(Otvet, `"`, ``)

	return Otvet
}

//// FindFunctions - находит массив функций go
//func FindFunctions(AstFile *ast.File) []GoImport {
//	Otvet := make([]GoImport, 0)
//
//	if AstFile == nil {
//		return Otvet
//	}
//
//	for _, decl1 := range AstFile.Decls {
//		switch decl1.(type) {
//		case *ast.FuncDecl:
//			{
//				func1 := decl1.(*ast.FuncDecl)
//				body := func1.Body
//				if body == nil {
//					continue
//				}
//				for _, list1 := range body.List {
//					switch list1.(type) {
//					case *ast.GoStmt:
//						{
//							GoStmt1 := list1.(*ast.GoStmt)
//							go_package_name, go_package_import, go_func_name := FindGoValues(AstFile, GoStmt1)
//							GoImport1 := GoImport{}
//							GoImport1.Go_package_name = go_package_name
//							GoImport1.Go_package_import = go_package_import
//							GoImport1.Go_func_name = go_func_name
//							Otvet = append(Otvet, GoImport1)
//						}
//					}
//				}
//			}
//		}
//	}
//
//	return Otvet
//}
