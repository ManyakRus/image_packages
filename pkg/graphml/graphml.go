package graphml

import (
	"github.com/ManyakRus/image_packages/pkg/xml"
	"github.com/beevik/etree"
	_ "github.com/beevik/etree"
	"math"
	"strconv"
	"strings"
)

// FONT_SIZE_SHAPE - размер шрифта прямоугольника
var FONT_SIZE_SHAPE = 16

// FONT_SIZE_SHAPE - размер шрифта групп
var FONT_SIZE_GROUP = 10

// FONT_SIZE_EDGE - размер шрифта стрелок
var FONT_SIZE_EDGE = 8

//var doc = etree.NewDocument()

// CreateElement_Shape - создаёт элемент xgml - прямоугольник
func CreateElement_Shape(ElementGraph *etree.Element, ElementGroup *etree.Element, ElementName string) *etree.Element {

	Width := findWidthShape(ElementName)
	Height := findHeightShape(ElementName)

	//node
	ElementNode := xml.AddSectionXML(ElementGraph, "node")
	xml.AddAttributeXML(ElementNode, "id", "int", strconv.Itoa(ElementNode.Index()))
	xml.AddAttributeXML(ElementNode, "label", "string", ElementName)

	//graphics
	ElementGraphics := xml.AddSectionXML(ElementNode, "graphics")
	xml.AddAttributeXML(ElementGraphics, "type", "string", "rectangle")
	xml.AddAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	xml.AddAttributeXML(ElementGraphics, "outline", "string", "#000000")
	xml.AddAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	xml.AddAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//LabelGraphics
	ElementLabelGraphics := xml.AddSectionXML(ElementNode, "LabelGraphics")
	xml.AddAttributeXML(ElementLabelGraphics, "text", "String", ElementName)
	xml.AddAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_SHAPE))

	//group
	if ElementGroup != nil {
		xml.AddAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	return ElementNode
}

// CreateElementGroup - создаёт элемент xgml - группа
func CreateElementGroup(ElementGraph, ElementGroup *etree.Element, GroupCaption string) *etree.Element {

	Width := findWidthGroup(GroupCaption)
	Height := findHeightGroup(GroupCaption)

	//node
	ElementNode := xml.AddSectionXML(ElementGraph, "node")
	xml.AddAttributeXML_int(ElementNode, "id", ElementNode.Index())
	xml.AddAttributeXML(ElementNode, "label", "string", GroupCaption)
	xml.AddAttributeXML(ElementNode, "isGroup", "boolean", "true")

	//graphics
	ElementGraphics := xml.AddSectionXML(ElementNode, "graphics")
	xml.AddAttributeXML(ElementGraphics, "type", "string", "roundrectangle")
	xml.AddAttributeXML(ElementGraphics, "fill", "string", "#F5F5F5")
	xml.AddAttributeXML(ElementGraphics, "outline", "string", "#F5F5F5")
	xml.AddAttributeXML_double(ElementGraphics, "h", float64(Height))
	xml.AddAttributeXML_double(ElementGraphics, "w", float64(Width))

	//LabelGraphics
	ElementLabelGraphics := xml.AddSectionXML(ElementNode, "LabelGraphics")
	xml.AddAttributeXML(ElementLabelGraphics, "text", "String", GroupCaption)
	xml.AddAttributeXML(ElementLabelGraphics, "fill", "String", "#EBEBEB")
	xml.AddAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_GROUP))
	xml.AddAttributeXML(ElementLabelGraphics, "anchor", "String", "n")
	xml.AddAttributeXML_double(ElementLabelGraphics, "borderDistance", 0)
	xml.AddAttributeXML_double(ElementLabelGraphics, "leftBorderInset", 50)
	xml.AddAttributeXML_double(ElementLabelGraphics, "rightBorderInset", 50)
	xml.AddAttributeXML_string(ElementLabelGraphics, "model", "sandwich")

	//group
	if ElementGroup != nil {
		xml.AddAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	return ElementNode
}

// CreateElementEdge - создаёт элемент xgml - стрелка
func CreateElementEdge(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	//edge
	ElementEdge := xml.AddSectionXML(ElementGraph, "edge")
	xml.AddAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	xml.AddAttributeXML_int(ElementEdge, "target", IndexElementTo)

	//graphics
	ElementGraphics := xml.AddSectionXML(ElementEdge, "graphics")
	xml.AddAttributeXML(ElementGraphics, "fill", "string", "#000000")
	xml.AddAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

}

// CreateElementEdge_blue - создаёт элемент xgml - стрелка синяя с заголовком
func CreateElementEdge_blue(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int, label string) {

	Width := float64(findWidthEdge(label))
	Height := float64(findHeightEdge(label))

	//edge
	ElementEdge := xml.AddSectionXML(ElementGraph, "edge")
	xml.AddAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	xml.AddAttributeXML_int(ElementEdge, "target", IndexElementTo)
	xml.AddAttributeXML_string(ElementEdge, "label", label)

	//graphics
	ElementGraphics := xml.AddSectionXML(ElementEdge, "graphics")
	xml.AddAttributeXML(ElementGraphics, "fill", "string", "#0000FF")
	xml.AddAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

	//LabelGraphics
	ElementLabelGraphics := xml.AddSectionXML(ElementEdge, "LabelGraphics")
	xml.AddAttributeXML(ElementLabelGraphics, "text", "String", label)
	xml.AddAttributeXML(ElementLabelGraphics, "color", "String", "#0000FF")
	xml.AddAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_EDGE))
	xml.AddAttributeXML(ElementLabelGraphics, "fontName", "String", "Dialog")
	xml.AddAttributeXML(ElementLabelGraphics, "configuration", "String", "AutoFlippingLabel")
	xml.AddAttributeXML_double(ElementLabelGraphics, "contentWidth", Width)
	xml.AddAttributeXML_double(ElementLabelGraphics, "contentHeight", Height)
	xml.AddAttributeXML_string(ElementLabelGraphics, "model", "two_pos")
	xml.AddAttributeXML_string(ElementLabelGraphics, "position", "head")

}

// findWidthShape - возвращает число - ширину элемента
func findWidthShape(ElementName string) int {
	Otvet := FONT_SIZE_SHAPE * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE/2)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeightShape - возвращает число - высоту элемента
func findHeightShape(ElementName string) int {

	Otvet := 10 + FONT_SIZE_SHAPE*3

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*FONT_SIZE_SHAPE*2

	return Otvet

}

// findWidthGroup - возвращает число - ширину элемента
func findWidthGroup(ElementName string) int {
	Otvet := 10

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeightGroup - возвращает число - высоту элемента
func findHeightGroup(ElementName string) int {

	Otvet := 30

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

// findWidthEdge - возвращает число - ширину элемента
func findWidthEdge(Label string) int {
	Otvet := 10

	LenMax := findMaxLenRow(Label)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeightEdge - возвращает число - высоту элемента
func findHeightEdge(Label string) int {

	Otvet := 30

	RowsTotal := countLines(Label)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

// countLines - возвращает количество переводов строки
func countLines(s string) int {
	Otvet := 0

	strings.Count(s, "/n")

	return Otvet
}

// findMaxLenRow - возвращает количество символов в строке максимум
func findMaxLenRow(ElementName string) int {
	Otvet := 0

	Mass := strings.Split(ElementName, "\n")

	for _, Mass1 := range Mass {
		len1 := len(Mass1)
		if len1 > Otvet {
			Otvet = len1
		}
	}

	return Otvet
}
