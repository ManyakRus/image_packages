package xgml

import (
	"github.com/beevik/etree"
	_ "github.com/beevik/etree"
	"math"
	"strconv"
	"strings"
)

var FONT_SIZE_SHAPE = 16

var FONT_SIZE_GROUP = 10

//var doc = etree.NewDocument()

// CreateDocXGML - создаёт новый документ .xgml
func CreateDocXGML(ShowKind string) *etree.Document {

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

func CreateElementXGML_Shape(ElementGraph *etree.Element, ElementGroup *etree.Element, ElementName string) *etree.Element {

	ElementNode := addSectionXML(ElementGraph, "node")
	addAttributeXML(ElementNode, "id", "int", strconv.Itoa(ElementNode.Index()))
	addAttributeXML(ElementNode, "label", "string", ElementName)

	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "rectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00

	Width := findWidthShape(ElementName)
	Height := findHeightShape(ElementName)

	//if ShowKind == "Группами" {
	//
	//	Width = int(math.Round(float64(Width) * 1.3))
	//
	//	//Высота = Окр(Высота*1.3, 0)
	//
	//	addAttributeXML(ElementGraphics, "outline", "string", "#FFFFFF")
	//
	//} else {

	addAttributeXML(ElementGraphics, "outline", "string", "#000000")

	//}

	addAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	addAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", ElementName)
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_SHAPE))

	if ElementGroup != nil {
		addAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	//ElementGraph.ЗаписатьКонецЭлемента() //node

	return ElementNode
} // ВидОтображения = "Группами"

func CreateGroupXGML(ElementGraph, ElementGroup *etree.Element, GroupCaption string) *etree.Element {

	ElementNode := addSectionXML(ElementGraph, "node")
	addAttributeXML_int(ElementNode, "id", ElementNode.Index())
	addAttributeXML(ElementNode, "label", "string", GroupCaption)
	addAttributeXML(ElementNode, "isGroup", "boolean", "true")

	Width := findWidthGroup(GroupCaption)
	Height := findHeightGroup(GroupCaption)

	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "roundrectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#F5F5F5")
	addAttributeXML(ElementGraphics, "outline", "string", "#F5F5F5")
	//addAttributeXML(ElementGraphics, "outline", "string", "#000000")
	addAttributeXML_double(ElementGraphics, "h", float64(Height))
	addAttributeXML_double(ElementGraphics, "w", float64(Width))

	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", GroupCaption)
	addAttributeXML(ElementLabelGraphics, "fill", "String", "#EBEBEB")
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_GROUP))
	addAttributeXML(ElementLabelGraphics, "anchor", "String", "n")
	//addAttributeXML_string(ElementLabelGraphics, "autoSizePolicy", "node_width")
	addAttributeXML_double(ElementLabelGraphics, "borderDistance", 0)
	addAttributeXML_double(ElementLabelGraphics, "leftBorderInset", 50)
	addAttributeXML_double(ElementLabelGraphics, "rightBorderInset", 50)
	addAttributeXML_string(ElementLabelGraphics, "model", "sandwich")

	if ElementGroup != nil {
		addAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	}

	//ElementGraph.ЗаписатьКонецЭлемента() //node

	return ElementNode
} // CreateGroupXGML()

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

	//addAttributeXML(ElementNode, "label", "string", ElementName)

	//graphics

	ElementGraphics := addSectionXML(ElementNode, "graphics")
	addAttributeXML(ElementGraphics, "type", "string", "rectangle")
	addAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	addAttributeXML(ElementGraphics, "outline", "string", "#000000")
	addAttributeXML(ElementGraphics, "customconfiguration", "string", "com.yworks.entityRelationship.big_entity")
	addAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	addAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//styleproperties

	ElementStyleProperties := addSectionXML(ElementGraphics, "styleproperties")

	//property

	addSectionXML(ElementStyleProperties, "property")
	addAttributeXML(ElementStyleProperties, "name", "string", "y.view.ShadowNodePainter.SHADOW_PAINTING")
	addAttributeXML(ElementStyleProperties, "valueClass", "string", "java.lang.Boolean")
	addAttributeXML(ElementStyleProperties, "value", "string", "true")

	//ElementGraph.ЗаписатьКонецЭлемента() //property

	//

	//ElementGraph.ЗаписатьКонецЭлемента() //styleproperties

	//

	//ElementGraph.ЗаписатьКонецЭлемента() //graphics

	//LabelGraphics-1

	ElementLabelGraphics := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics, "text", "String", ElementId)
	addAttributeXML(ElementLabelGraphics, "fontSize", "int", "12")
	addAttributeXML(ElementLabelGraphics, "configuration", "String", "com.yworks.entityRelationship.label.name")
	addAttributeXML(ElementLabelGraphics, "anchor", "String", "t")
	addAttributeXML(ElementLabelGraphics, "contentWidth", "int", "24")
	addAttributeXML(ElementLabelGraphics, "contentHeight", "int", "18")

	//ElementGraph.ЗаписатьКонецЭлемента()

	//LabelGraphics-2

	ElementLabelGraphics2 := addSectionXML(ElementNode, "LabelGraphics")
	addAttributeXML(ElementLabelGraphics2, "text", "String", ElementName)
	addAttributeXML(ElementLabelGraphics2, "fontSize", "int", "12")
	addAttributeXML(ElementLabelGraphics2, "configuration", "String", "com.yworks.entityRelationship.label.attributes")
	addAttributeXML(ElementLabelGraphics2, "alignment", "String", "left")
	addAttributeXML(ElementLabelGraphics2, "contentWidth", "int", "24")
	addAttributeXML(ElementLabelGraphics2, "contentHeight", "int", "18")

	//ElementGraph.ЗаписатьКонецЭлемента()

	//gid

	if (ElementGroup) != nil {
		addAttributeXML_int(ElementGraph, "gid", ElementGroup.Index())
	}

	//

	//ElementGraph.ЗаписатьКонецЭлемента() //node

	return ElementNode
}

//func CreateElementXGML_WithGroup(Element, ElementGroup *etree.Element, ElementName, ВидОтображения string) { // ВидОтображения = "Группами"
//
//	//if ElementId == "" {
//	//	ElementId = ElementName
//	//}
//
//	CreateElementXGML_Shape(Element, ElementGroup, ElementName, ВидОтображения)
//	//CreateGroupXGML(Element, ElementName, ElementGroup)
//
//}
//
//func CreateElementXGML(Element, ElementGroup *etree.Element, ElementId int, ElementName, ShowKind string) {
//
//	if ShowKind == "Группами" {
//
//		CreateElementXGML_WithGroup(Element, ElementGroup, ElementName, ShowKind)
//
//	} else {
//
//		Текст := strconv.Itoa(ElementId) + "\n" + ElementName
//
//		CreateElementXGML_Shape(Element, ElementGroup, Текст, ShowKind)
//
//	}
//
//}

func AddDirectory(buffer *strings.Builder, Name string) {

}

func CreateLinkXGML(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	ElementEdge := addSectionXML(ElementGraph, "edge")
	addAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	addAttributeXML_int(ElementEdge, "target", IndexElementTo)

	ElementGraphics := addSectionXML(ElementEdge, "graphics")
	addAttributeXML(ElementGraphics, "fill", "string", "#000000")
	addAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

	//Element.ЗаписатьКонецЭлемента() //graphics
	//
	//Element.ЗаписатьКонецЭлемента() //edge

} // CreateLinkXGML()

// возвращает число - ширину элемента
func findWidthShape(ElementName string) int {
	Otvet := FONT_SIZE_SHAPE * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE/2)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// возвращает число - высоту элемента
func findHeightShape(ElementName string) int {

	Otvet := 10 + FONT_SIZE_SHAPE*3

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*FONT_SIZE_SHAPE*2

	return Otvet

}

// возвращает число - ширину элемента
func findWidthGroup(ElementName string) int {
	Otvet := 10

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// возвращает число - высоту элемента
func findHeightGroup(ElementName string) int {

	Otvet := 30

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

func countLines(s string) int {
	Otvet := 0

	strings.Count(s, "/n")

	return Otvet
}

//санек вид как UML 25.10.2019 10:05:41

func addSectionXML(Element *etree.Element, name string) *etree.Element {

	Element1 := Element.CreateElement("section")
	Element1.CreateAttr("name", name)

	//addAttributeXML(Element1, "id", "int", strconv.Itoa(Element1.Index()))

	//Element.ЗаписатьBeginЭлемента("section")
	//Element.ЗаписатьАтрибут("name", name)

	return Element1
} // ЗаписатьBeginСекции()

func addAttributeXML(Element *etree.Element, key, stype, text string) *etree.Element {

	ElementAttribute := Element.CreateElement("attribute")
	ElementAttribute.CreateAttr("key", key)
	ElementAttribute.CreateAttr("type", stype)
	ElementAttribute.SetText(text)

	return ElementAttribute
}

func addAttributeXML_int(Element *etree.Element, key string, value int) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "int", strconv.Itoa(value))

	return ElementAttribute
}

func addAttributeXML_double(Element *etree.Element, key string, value float64) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "double", strconv.Itoa(int(value)))

	return ElementAttribute
}

func addAttributeXML_string(Element *etree.Element, key string, value string) *etree.Element {

	ElementAttribute := addAttributeXML(Element, key, "String", value)

	return ElementAttribute
}

// возвращает количество символов в строке максимум
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
