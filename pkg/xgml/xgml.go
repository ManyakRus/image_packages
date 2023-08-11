package xgml

import (
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

// CreateDocXGML - создаёт новый документ .xgml
func CreateDocXGML() *etree.Document {

	DocXML := etree.NewDocument()
	DocXML.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	//DocXML.IndentTabs()

	ElementXGML := DocXML.CreateElement("section")
	ElementXGML.CreateAttr("name", "xgml")

	//ElementXGML := addSectionXML(DocXML, "xgml")
	ElementGraph := addSectionXML(ElementXGML, "graph")
	addAttributeXML(ElementGraph, "hierarchic", "int", "1")

	//log.Info("ElementGraph.GetPath(): ", ElementGraph.GetPath())

	return DocXML
}

// CreateElementXGML_Shape - создаёт элемент xgml - прямоугольник
func CreateElementXGML_Shape(ElementGraph *etree.Element, ElementGroup *etree.Element, ElementName string) *etree.Element {

	Width := findWidthShape(ElementName)
	Height := findHeightShape(ElementName)

	//node
	ElementNode := addSectionXML(ElementGraph, "node")
	addAttributeXML(ElementNode, "id", "int", strconv.Itoa(ElementNode.Index()))
	addAttributeXML(ElementNode, "label", "string", ElementName)

	//graphics
	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "rectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	addAttributeXML(ElementGraphics, "outline", "string", "#000000")
	addAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	addAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//LabelGraphics
	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", ElementName)
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_SHAPE))

	//group
	if ElementGroup != nil {
		addAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	return ElementNode
}

// CreateGroupXGML - создаёт элемент xgml - группа
func CreateGroupXGML(ElementGraph, ElementGroup *etree.Element, GroupCaption string) *etree.Element {

	Width := findWidthGroup(GroupCaption)
	Height := findHeightGroup(GroupCaption)

	//node
	ElementNode := addSectionXML(ElementGraph, "node")
	addAttributeXML_int(ElementNode, "id", ElementNode.Index())
	addAttributeXML(ElementNode, "label", "string", GroupCaption)
	addAttributeXML(ElementNode, "isGroup", "boolean", "true")

	//graphics
	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "roundrectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#F5F5F5")
	addAttributeXML(ElementGraphics, "outline", "string", "#F5F5F5")
	addAttributeXML_double(ElementGraphics, "h", float64(Height))
	addAttributeXML_double(ElementGraphics, "w", float64(Width))

	//LabelGraphics
	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", GroupCaption)
	addAttributeXML(ElementLabelGraphics, "fill", "String", "#EBEBEB")
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_GROUP))
	addAttributeXML(ElementLabelGraphics, "anchor", "String", "n")
	addAttributeXML_double(ElementLabelGraphics, "borderDistance", 0)
	addAttributeXML_double(ElementLabelGraphics, "leftBorderInset", 50)
	addAttributeXML_double(ElementLabelGraphics, "rightBorderInset", 50)
	addAttributeXML_string(ElementLabelGraphics, "model", "sandwich")

	//group
	if ElementGroup != nil {
		addAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	return ElementNode
}

// CreateElementXGML_UML - создаёт элемент xgml - UML
func CreateElementXGML_UML(ElementGraph *etree.Element, ElementGroup *etree.Element, ElementId, ElementName string) *etree.Element {

	if ElementId == "" {
		ElementId = ElementName
	}

	Width := findWidthShape(ElementName)
	Height := findHeightShape(ElementName)

	//node
	ElementNode := addSectionXML(ElementGraph, "node")
	addAttributeXML(ElementNode, "id", "string", ElementId)
	addAttributeXML(ElementNode, "label", "string", "")

	//graphics
	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "rectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	addAttributeXML(ElementGraphics, "outline", "string", "#000000")
	addAttributeXML(ElementGraphics, "customconfiguration", "string", "com.yworks.entityRelationship.big_entity")
	addAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	addAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//style
	ElementStyleProperties := addSectionXML(ElementGraphics, "styleproperties")

	//property
	addSectionXML(ElementStyleProperties, "property")
	addAttributeXML(ElementStyleProperties, "name", "string", "y.view.ShadowNodePainter.SHADOW_PAINTING")
	addAttributeXML(ElementStyleProperties, "valueClass", "string", "java.lang.Boolean")
	addAttributeXML(ElementStyleProperties, "value", "string", "true")

	//LabelGraphics
	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", ElementId)
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", "12")
	addAttributeXML(ElementLabelGraphics, "configuration", "String", "com.yworks.entityRelationship.label.name")
	addAttributeXML(ElementLabelGraphics, "anchor", "String", "t")
	addAttributeXML(ElementLabelGraphics, "contentWidth", "int", "24")
	addAttributeXML(ElementLabelGraphics, "contentHeight", "int", "18")

	//LabelGraphics2
	ElementLabelGraphics2 := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics2, "text", "String", ElementName)
	addAttributeXML(ElementLabelGraphics2, "fontSize", "int", "12")
	addAttributeXML(ElementLabelGraphics2, "configuration", "String", "com.yworks.entityRelationship.label.attributes")
	addAttributeXML(ElementLabelGraphics2, "alignment", "String", "left")
	addAttributeXML(ElementLabelGraphics2, "contentWidth", "int", "24")
	addAttributeXML(ElementLabelGraphics2, "contentHeight", "int", "18")

	//group
	if (ElementGroup) != nil {
		addAttributeXML_int(ElementGraph, "gid", ElementGroup.Index())
	}
	return ElementNode
}

// CreateLinkXGML - создаёт элемент xgml - стрелка
func CreateLinkXGML(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	//edge
	ElementEdge := addSectionXML(ElementGraph, "edge")
	addAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	addAttributeXML_int(ElementEdge, "target", IndexElementTo)

	//graphics
	ElementGraphics := addSectionXML(ElementEdge, "graphics")
	addAttributeXML(ElementGraphics, "fill", "string", "#000000")
	addAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

}

// CreateLinkXGML - создаёт элемент xgml - стрелка синяя с заголовком
func CreateLinkXGML_blue(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int, label string) {

	Width := float64(findWidthEdge(label))
	Height := float64(findHeightEdge(label))

	//edge
	ElementEdge := addSectionXML(ElementGraph, "edge")
	addAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	addAttributeXML_int(ElementEdge, "target", IndexElementTo)
	addAttributeXML_string(ElementEdge, "label", label)

	//graphics
	ElementGraphics := addSectionXML(ElementEdge, "graphics")
	addAttributeXML(ElementGraphics, "fill", "string", "#0000FF")
	addAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

	//LabelGraphics
	ElementLabelGraphics := addSectionXML(ElementEdge, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", label)
	addAttributeXML(ElementLabelGraphics, "color", "String", "#0000FF")
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_EDGE))
	addAttributeXML(ElementLabelGraphics, "fontName", "String", "Dialog")
	addAttributeXML(ElementLabelGraphics, "configuration", "String", "AutoFlippingLabel")
	addAttributeXML_double(ElementLabelGraphics, "contentWidth", Width)
	addAttributeXML_double(ElementLabelGraphics, "contentHeight", Height)
	addAttributeXML_string(ElementLabelGraphics, "model", "two_pos")
	addAttributeXML_string(ElementLabelGraphics, "position", "head")

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

// addSectionXML - добавляет секцию в xgml
func addSectionXML(Element *etree.Element, name string) *etree.Element {

	Element1 := Element.CreateElement("section")
	Element1.CreateAttr("name", name)

	return Element1
}

// addAttributeXML - добавляет аттрибут в элемент xgml
func addAttributeXML(Element *etree.Element, key, stype, text string) *etree.Element {

	ElementAttribute := Element.CreateElement("attribute")
	ElementAttribute.CreateAttr("key", key)
	ElementAttribute.CreateAttr("type", stype)
	ElementAttribute.SetText(text)

	return ElementAttribute
}

// addAttributeXML_int - добавляет аттрибут типа int в элемент xgml
func addAttributeXML_int(Element *etree.Element, key string, value int) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "int", strconv.Itoa(value))

	return ElementAttribute
}

// addAttributeXML_double - добавляет аттрибут типа double в элемент xgml
func addAttributeXML_double(Element *etree.Element, key string, value float64) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "double", strconv.Itoa(int(value)))

	return ElementAttribute
}

// addAttributeXML_string - добавляет аттрибут типа string в элемент xgml
func addAttributeXML_string(Element *etree.Element, key string, value string) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "String", value)

	return ElementAttribute
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
