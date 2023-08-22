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
func CreateElement_Shape(ElementGraph0 *etree.Element, ElementGroup *etree.Element, ElementName string) *etree.Element {

	Width := findWidth_Shape(ElementName)
	Height := findHeight_Shape(ElementName)
	sWidth := fmt.Sprintf("%.1f", float32(Width))
	sHeight := fmt.Sprintf("%.1f", float32(Height))

	sFontSize := strconv.Itoa(FONT_SIZE_SHAPE)

	//ищем graph
	var ElementGraph *etree.Element
	ElementGraph2 := ElementGraph0.SelectElement("graph")
	if ElementGraph2 != nil {
		ElementGraph = ElementGraph2
	} else {
		ElementGraph = ElementGraph0
	}

	//node
	ElementNode := ElementGraph.CreateElement("node")
	ElementNode.CreateAttr("id", "n"+strconv.Itoa(ElementNode.Index()))

	//data
	ElementData := ElementNode.CreateElement("data")
	ElementData.CreateAttr("key", "d5")

	//ShapeNode
	ElementShapeNode := ElementData.CreateElement("y:ShapeNode")

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
	ElementBorderStyle := ElementShapeNode.CreateElement("y:BorderStyle")
	ElementBorderStyle.CreateAttr("color", "#000000")
	ElementBorderStyle.CreateAttr("type", "line")
	ElementBorderStyle.CreateAttr("width", "1.0")

	//NodeLabel
	ElementNodeLabel := ElementShapeNode.CreateElement("y:NodeLabel")
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
	ElementNodeLabel.CreateText(ElementName)

	//y:Shape
	ElementYShape := ElementShapeNode.CreateElement("y:Shape")
	ElementYShape.CreateAttr("type", "rectangle")

	return ElementNode
}

// CreateElement_Group - создаёт элемент xgml - группа
func CreateElement_Group(ElementGraph, ElementGroup *etree.Element, GroupCaption string) *etree.Element {

	Width := findWidth_Group(GroupCaption)
	Height := findHeight_Group(GroupCaption)
	sWidth := fmt.Sprintf("%.1f", float32(Width))
	sHeight := fmt.Sprintf("%.1f", float32(Height))

	//node
	ElementNode := ElementGraph.CreateElement("node")
	NodeId := "n" + strconv.Itoa(ElementNode.Index())
	ElementNode.CreateAttr("id", NodeId+"::"+NodeId)
	ElementNode.CreateAttr("yfiles.foldertype", "group")

	//data
	ElementData := ElementNode.CreateElement("data")
	ElementData.CreateAttr("key", "d5")

	//YProxyAutoBoundsNode
	ElementYProxyAutoBoundsNode := ElementData.CreateElement("y:ProxyAutoBoundsNode")

	//YRealizers
	ElementYRealizers := ElementYProxyAutoBoundsNode.CreateElement("y:Realizers")
	ElementYRealizers.CreateAttr("active", "0")

	//----------------------- visible ---------------------------------------------

	//YGroupNode
	ElementYGroupNode := ElementYRealizers.CreateElement("y:GroupNode")

	//YGeometry
	ElementYGeometry := ElementYGroupNode.CreateElement("y:Geometry")
	ElementYGeometry.CreateAttr("height", sHeight)
	ElementYGeometry.CreateAttr("width", sWidth)
	ElementYGeometry.CreateAttr("x", "0.0")
	ElementYGeometry.CreateAttr("y", "0.0")

	//YFill
	ElementYFill := ElementYGroupNode.CreateElement("y:Fill")
	ElementYFill.CreateAttr("color", "#F5F5F5")
	ElementYFill.CreateAttr("transparent", "false")

	//YBorderStyle
	ElementYBorderStyle := ElementYGroupNode.CreateElement("y:BorderStyle")
	ElementYBorderStyle.CreateAttr("color", "#000000")
	ElementYBorderStyle.CreateAttr("type", "dashed")
	ElementYBorderStyle.CreateAttr("width", "1.0")

	//YNodeLabel
	ElementYNodeLabel := ElementYGroupNode.CreateElement("y:NodeLabel")
	ElementYNodeLabel.CreateAttr("alignment", "right")
	ElementYNodeLabel.CreateAttr("autoSizePolicy", "content")
	ElementYNodeLabel.CreateAttr("backgroundColor", "#EBEBEB")
	ElementYNodeLabel.CreateAttr("borderDistance", "0.0")
	ElementYNodeLabel.CreateAttr("fontFamily", "Dialog")
	ElementYNodeLabel.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_GROUP))
	ElementYNodeLabel.CreateAttr("fontStyle", "plain")
	ElementYNodeLabel.CreateAttr("hasLineColor", "false")
	ElementYNodeLabel.CreateAttr("height", sHeight)
	ElementYNodeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementYNodeLabel.CreateAttr("iconTextGap", "4")
	ElementYNodeLabel.CreateAttr("modelName", "sandwich")
	ElementYNodeLabel.CreateAttr("modelPosition", "n")
	ElementYNodeLabel.CreateAttr("textColor", "#000000")
	ElementYNodeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementYNodeLabel.CreateAttr("width", sWidth)
	ElementYNodeLabel.CreateAttr("x", "0")
	ElementYNodeLabel.CreateAttr("xml:space", "preserve")
	ElementYNodeLabel.CreateAttr("y", "0")
	ElementYNodeLabel.CreateText(GroupCaption)

	//YShape
	ElementYShape := ElementYGroupNode.CreateElement("y:Shape")
	ElementYShape.CreateAttr("type", "roundrectangle")

	//YState
	ElementYState := ElementYGroupNode.CreateElement("y:State")
	ElementYState.CreateAttr("closed", "false")
	ElementYState.CreateAttr("closedHeight", "80.0")
	ElementYState.CreateAttr("closedWidth", "100.0")
	ElementYState.CreateAttr("innerGraphDisplayEnabled", "false")

	//YInsets
	ElementYInsets := ElementYGroupNode.CreateElement("y:Insets")
	ElementYInsets.CreateAttr("bottom", "15")
	ElementYInsets.CreateAttr("bottomF", "15.0")
	ElementYInsets.CreateAttr("left", "15")
	ElementYInsets.CreateAttr("leftF", "15.0")
	ElementYInsets.CreateAttr("right", "15")
	ElementYInsets.CreateAttr("rightF", "15.0")
	ElementYInsets.CreateAttr("top", "15")
	ElementYInsets.CreateAttr("topF", "15.0")

	//YBorderInsets
	ElementYBorderInsets := ElementYGroupNode.CreateElement("y:BorderInsets")
	ElementYBorderInsets.CreateAttr("bottom", "54")
	ElementYBorderInsets.CreateAttr("bottomF", "54.0")
	ElementYBorderInsets.CreateAttr("left", "0")
	ElementYBorderInsets.CreateAttr("leftF", "0.0")
	ElementYBorderInsets.CreateAttr("right", "23")
	ElementYBorderInsets.CreateAttr("rightF", "23.35")
	ElementYBorderInsets.CreateAttr("top", "0")
	ElementYBorderInsets.CreateAttr("topF", "0.0")

	//----------------------- not visible ---------------------------------------------

	//YGroupNode
	ElementYGroupNode2 := ElementYRealizers.CreateElement("y:GroupNode")

	//YGeometry
	ElementYGeometry2 := ElementYGroupNode2.CreateElement("y:Geometry")
	ElementYGeometry2.CreateAttr("height", sHeight)
	ElementYGeometry2.CreateAttr("width", sWidth)
	ElementYGeometry2.CreateAttr("x", "0.0")
	ElementYGeometry2.CreateAttr("y", "0.0")

	//YFill
	ElementYFill2 := ElementYGroupNode2.CreateElement("y:Fill")
	ElementYFill2.CreateAttr("color", "#F5F5F5")
	ElementYFill2.CreateAttr("transparent", "false")

	//YBorderStyle
	ElementYBorderStyle2 := ElementYGroupNode2.CreateElement("y:BorderStyle")
	ElementYBorderStyle2.CreateAttr("color", "#000000")
	ElementYBorderStyle2.CreateAttr("type", "dashed")
	ElementYBorderStyle2.CreateAttr("width", "1.0")

	//YNodeLabel
	ElementYNodeLabel2 := ElementYGroupNode2.CreateElement("y:NodeLabel")
	ElementYNodeLabel2.CreateAttr("alignment", "right")
	ElementYNodeLabel2.CreateAttr("autoSizePolicy", "content")
	ElementYNodeLabel2.CreateAttr("backgroundColor", "#EBEBEB")
	ElementYNodeLabel2.CreateAttr("borderDistance", "0.0")
	ElementYNodeLabel2.CreateAttr("fontFamily", "Dialog")
	ElementYNodeLabel2.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_GROUP))
	ElementYNodeLabel2.CreateAttr("fontStyle", "plain")
	ElementYNodeLabel2.CreateAttr("hasLineColor", "false")
	ElementYNodeLabel2.CreateAttr("hasText", "false") //только у 2
	ElementYNodeLabel2.CreateAttr("height", sHeight)
	ElementYNodeLabel2.CreateAttr("horizontalTextPosition", "center")
	ElementYNodeLabel2.CreateAttr("iconTextGap", "4")
	ElementYNodeLabel2.CreateAttr("modelName", "sandwich")
	ElementYNodeLabel2.CreateAttr("modelPosition", "n")
	ElementYNodeLabel2.CreateAttr("textColor", "#000000")
	ElementYNodeLabel2.CreateAttr("verticalTextPosition", "bottom")
	ElementYNodeLabel2.CreateAttr("width", sWidth)
	ElementYNodeLabel2.CreateAttr("x", "0")
	//ElementYNodeLabel2.CreateAttr("xml:space", "preserve") //только у 2
	ElementYNodeLabel2.CreateAttr("y", "0")
	//ElementYNodeLabel2.CreateText(GroupCaption) //только у 2

	//YShape
	ElementYShape2 := ElementYGroupNode2.CreateElement("y:Shape")
	ElementYShape2.CreateAttr("type", "roundrectangle")

	//YState
	ElementYState2 := ElementYGroupNode2.CreateElement("y:State")
	ElementYState2.CreateAttr("closed", "true")
	ElementYState2.CreateAttr("closedHeight", "80.0")
	ElementYState2.CreateAttr("closedWidth", "100.0")
	ElementYState2.CreateAttr("innerGraphDisplayEnabled", "false")

	//YInsets
	ElementYInsets2 := ElementYGroupNode2.CreateElement("y:Insets")
	ElementYInsets2.CreateAttr("bottom", "15")
	ElementYInsets2.CreateAttr("bottomF", "15.0")
	ElementYInsets2.CreateAttr("left", "15")
	ElementYInsets2.CreateAttr("leftF", "15.0")
	ElementYInsets2.CreateAttr("right", "15")
	ElementYInsets2.CreateAttr("rightF", "15.0")
	ElementYInsets2.CreateAttr("top", "15")
	ElementYInsets2.CreateAttr("topF", "15.0")

	//YBorderInsets
	ElementYBorderInsets2 := ElementYGroupNode2.CreateElement("y:BorderInsets")
	ElementYBorderInsets2.CreateAttr("bottom", "54")
	ElementYBorderInsets2.CreateAttr("bottomF", "54.0")
	ElementYBorderInsets2.CreateAttr("left", "0")
	ElementYBorderInsets2.CreateAttr("leftF", "0.0")
	ElementYBorderInsets2.CreateAttr("right", "23")
	ElementYBorderInsets2.CreateAttr("rightF", "23.35")
	ElementYBorderInsets2.CreateAttr("top", "0")
	ElementYBorderInsets2.CreateAttr("topF", "0.0")

	//----------------------- продолжение ---------------------------------------------
	//YBorderInsets
	ElementGraphGraph := ElementNode.CreateElement("graph")
	ElementGraphGraph.CreateAttr("edgedefault", "directed")
	ElementGraphGraph.CreateAttr("id", NodeId+":")

	return ElementNode
}

// CreateElement_Edge - создаёт элемент xgml - стрелка
func CreateElement_Edge(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	////edge
	//ElementEdge := xml.AddSectionXML(ElementGraph, "edge")
	//xml.AddAttributeXML_int(ElementEdge, "source", IndexElementFrom)
	//xml.AddAttributeXML_int(ElementEdge, "target", IndexElementTo)
	//
	////graphics
	//ElementGraphics := xml.AddSectionXML(ElementEdge, "graphics")
	//xml.AddAttributeXML(ElementGraphics, "fill", "string", "#000000")
	//xml.AddAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

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
func CreateDocument() (*etree.Document, *etree.Element) {

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

	return DocXML, ElementGraph
}
