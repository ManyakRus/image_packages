package graphml

import (
	"fmt"
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

	Width := findWidth_Shape(ElementName)
	Height := findHeight_Shape(ElementName)
	sWidth := fmt.Sprintf("%.1f", Width)
	sHeight := fmt.Sprintf("%.1f", Height)

	sFontSize := strconv.Itoa(FONT_SIZE_SHAPE)

	//node
	ElementNode := ElementGraph.CreateElement("node")
	ElementGraph.CreateAttr("id", strconv.Itoa(ElementNode.Index()))

	//data
	ElementData := ElementGraph.CreateElement("data")
	ElementData.CreateAttr("key", "d5")

	//ShapeNode
	ElementShapeNode := ElementData.CreateElement("ShapeNode")

	//YGeometry
	ElementYGeometry := ElementShapeNode.CreateElement("y:Geometry")
	ElementYGeometry.CreateAttr("height", sHeight)
	ElementYGeometry.CreateAttr("width", sWidth)
	ElementYGeometry.CreateAttr("x", "0.0")
	ElementYGeometry.CreateAttr("y", "0.0")

	//YFill
	ElementYFill := ElementShapeNode.CreateElement("y:Fill")
	ElementYFill.CreateAttr("color", "#FFFFFF")
	ElementYFill.CreateAttr("transparent", "false")

	//BorderStyle
	ElementBorderStyle := ElementShapeNode.CreateElement("BorderStyle")
	ElementBorderStyle.CreateAttr("color", "#000000")
	ElementBorderStyle.CreateAttr("type", "line")
	ElementBorderStyle.CreateAttr("width", "1.0")

	//NodeLabel
	ElementNodeLabel := ElementShapeNode.CreateElement("NodeLabel")
	ElementNodeLabel.CreateAttr("alignment", "center")
	ElementNodeLabel.CreateAttr("autoSizePolicy", "content")
	ElementNodeLabel.CreateAttr("fontFamily", "Dialog")
	ElementNodeLabel.CreateAttr("fontSize", sFontSize)
	ElementNodeLabel.CreateAttr("fontStyle", "plain")
	ElementNodeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementNodeLabel.CreateAttr("hasLineColor", "false")
	ElementNodeLabel.CreateAttr("height", sHeight)
	ElementNodeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementNodeLabel.CreateAttr("iconTextGap", "4")
	ElementNodeLabel.CreateAttr("modelName", "internal")
	ElementNodeLabel.CreateAttr("modelPosition", "c")
	ElementNodeLabel.CreateAttr("textColor", "#000000")
	ElementNodeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementNodeLabel.CreateAttr("visible", "true")
	ElementNodeLabel.CreateAttr("width", sWidth)
	ElementNodeLabel.CreateAttr("x", "0.0")
	ElementNodeLabel.CreateAttr("xml:space", "preserve")
	ElementNodeLabel.CreateAttr("y", "0.0")

	////node
	//ElementNode := xml.AddSectionXML(ElementGraph, "node")
	//xml.AddAttributeXML(ElementNode, "id", "int", strconv.Itoa(ElementNode.Index()))
	//xml.AddAttributeXML(ElementNode, "label", "string", ElementName)
	//
	////graphics
	//ElementGraphics := xml.AddSectionXML(ElementNode, "graphics")
	//xml.AddAttributeXML(ElementGraphics, "type", "string", "rectangle")
	//xml.AddAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	//xml.AddAttributeXML(ElementGraphics, "outline", "string", "#000000")
	//xml.AddAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	//xml.AddAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))
	//
	////LabelGraphics
	//ElementLabelGraphics := xml.AddSectionXML(ElementNode, "LabelGraphics")
	//xml.AddAttributeXML(ElementLabelGraphics, "text", "String", ElementName)
	//xml.AddAttributeXML(ElementLabelGraphics, "fontSize", "int", strconv.Itoa(FONT_SIZE_SHAPE))
	//
	////group
	//if ElementGroup != nil {
	//	xml.AddAttributeXML_int(ElementNode, "gid", ElementGroup.Index())
	//}

	return ElementNode
}

// CreateElement_Group - создаёт элемент xgml - группа
func CreateElement_Group(ElementGraph, ElementGroup *etree.Element, GroupCaption string) *etree.Element {

	Width := findWidth_Group(GroupCaption)
	Height := findHeight_Group(GroupCaption)

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

// CreateElement_Edge - создаёт элемент xgml - стрелка
func CreateElement_Edge(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	//edge
	ElementEdge := xml.AddSectionXML(ElementGraph, "edge")
	xml.AddAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	xml.AddAttributeXML_int(ElementEdge, "target", IndexElementTo)

	//graphics
	ElementGraphics := xml.AddSectionXML(ElementEdge, "graphics")
	xml.AddAttributeXML(ElementGraphics, "fill", "string", "#000000")
	xml.AddAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

}

// CreateElement_Edge_blue - создаёт элемент xgml - стрелка синяя с заголовком
func CreateElement_Edge_blue(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int, label string) {

	Width := float64(findWidth_Edge(label))
	Height := float64(findHeight_Edge(label))

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

// findWidth_Shape - возвращает число - ширину элемента
func findWidth_Shape(ElementName string) int {
	Otvet := FONT_SIZE_SHAPE * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE/2)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Shape - возвращает число - высоту элемента
func findHeight_Shape(ElementName string) int {

	Otvet := 10 + FONT_SIZE_SHAPE*3

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*FONT_SIZE_SHAPE*2

	return Otvet

}

// findWidth_Group - возвращает число - ширину элемента
func findWidth_Group(ElementName string) int {
	Otvet := 10

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Group - возвращает число - высоту элемента
func findHeight_Group(ElementName string) int {

	Otvet := 30

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

// findWidth_Edge - возвращает число - ширину элемента
func findWidth_Edge(Label string) int {
	Otvet := 10

	LenMax := findMaxLenRow(Label)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Edge - возвращает число - высоту элемента
func findHeight_Edge(Label string) int {

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

// CreateDocument - создаёт новый документ .xgml
func CreateDocument() *etree.Document {

	DocXML := etree.NewDocument()
	DocXML.CreateProcInst("xml", `version="1.0" encoding="UTF-8" standalone="no"`)

	ElementGraphMl := DocXML.CreateElement("graphml")
	ElementGraphMl.CreateAttr("xmlns", "http://graphml.graphdrawing.org/xmlns")
	ElementGraphMl.CreateAttr("xmlns:java", "http://www.yworks.com/xml/yfiles-common/1.0/java")
	ElementGraphMl.CreateAttr("xmlns:sys", "http://www.yworks.com/xml/yfiles-common/markup/primitives/2.0")
	ElementGraphMl.CreateAttr("xmlns:x", "http://www.yworks.com/xml/yfiles-common/markup/2.0")
	ElementGraphMl.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	ElementGraphMl.CreateAttr("xmlns:y", "http://www.yworks.com/xml/graphml")
	ElementGraphMl.CreateAttr("xmlns:y", "http://www.yworks.com/xml/graphml")
	ElementGraphMl.CreateAttr("xmlns:yed", "http://www.yworks.com/xml/yed/3")
	ElementGraphMl.CreateAttr("xsi:schemaLocation", "http://graphml.graphdrawing.org/xmlns http://www.yworks.com/xml/schema/graphml/1.1/ygraphml.xsd")

	ElementD0 := ElementGraphMl.CreateElement("key")
	ElementD0.CreateAttr("for", "port")
	ElementD0.CreateAttr("id", "d0")
	ElementD0.CreateAttr("yfiles.type", "portgraphics")

	ElementD1 := ElementGraphMl.CreateElement("key")
	ElementD1.CreateAttr("for", "port")
	ElementD1.CreateAttr("id", "d1")
	ElementD1.CreateAttr("yfiles.type", "portgeometry")

	ElementD2 := ElementGraphMl.CreateElement("key")
	ElementD2.CreateAttr("for", "port")
	ElementD2.CreateAttr("id", "d2")
	ElementD2.CreateAttr("yfiles.type", "portuserdata")

	ElementD3 := ElementGraphMl.CreateElement("key")
	ElementD3.CreateAttr("attr.name", "url")
	ElementD3.CreateAttr("attr.type", "string")
	ElementD3.CreateAttr("for", "node")
	ElementD3.CreateAttr("id", "d3")

	ElementD4 := ElementGraphMl.CreateElement("key")
	ElementD4.CreateAttr("attr.name", "description")
	ElementD4.CreateAttr("attr.type", "string")
	ElementD4.CreateAttr("for", "node")
	ElementD4.CreateAttr("id", "d4")

	ElementD5 := ElementGraphMl.CreateElement("key")
	ElementD5.CreateAttr("for", "node")
	ElementD5.CreateAttr("id", "d5")
	ElementD5.CreateAttr("yfiles.type", "nodegraphics")

	ElementD6 := ElementGraphMl.CreateElement("key")
	ElementD6.CreateAttr("for", "graphml")
	ElementD6.CreateAttr("id", "d6")
	ElementD6.CreateAttr("yfiles.type", "resources")

	ElementD7 := ElementGraphMl.CreateElement("key")
	ElementD7.CreateAttr("attr.name", "url")
	ElementD7.CreateAttr("attr.type", "string")
	ElementD7.CreateAttr("for", "edge")
	ElementD7.CreateAttr("id", "d7")

	ElementD8 := ElementGraphMl.CreateElement("key")
	ElementD8.CreateAttr("attr.name", "description")
	ElementD8.CreateAttr("attr.type", "string")
	ElementD8.CreateAttr("for", "edge")
	ElementD8.CreateAttr("id", "d8")

	ElementD9 := ElementGraphMl.CreateElement("key")
	ElementD9.CreateAttr("for", "edge")
	ElementD9.CreateAttr("id", "d9")
	ElementD9.CreateAttr("yfiles.type", "edgegraphics")

	ElementGraph := ElementGraphMl.CreateElement("graph")
	ElementGraph.CreateAttr("edgedefault", "directed")
	ElementGraph.CreateAttr("id", "G")

	return DocXML
}
