package graphml

import (
	"fmt"
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
func CreateElement_Shape(ElementGraph0 *etree.Element, ElementName string) *etree.Element {

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
	sId := FindId(ElementGraph0, ElementNode)
	ElementNode.CreateAttr("id", sId)
	//ElementNode.CreateAttr("id", "n"+strconv.Itoa(ElementNode.Index()))

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
func CreateElement_Group(ElementGraph0 *etree.Element, GroupCaption string) *etree.Element {

	Width := findWidth_Group(GroupCaption)
	Height := findHeight_Group(GroupCaption)
	sWidth := fmt.Sprintf("%.1f", float32(Width))
	sHeight := fmt.Sprintf("%.1f", float32(Height))

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
	//NodeId := "n" + strconv.Itoa(ElementNode.Index())
	NodeId := FindId(ElementGraph, ElementNode)
	ElementNode.CreateAttr("id", NodeId)
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
func CreateElement_Edge(ElementGraph, ElementFrom, ElementTo *etree.Element, label, Description string) *etree.Element {

	//node
	ElementEdge := ElementGraph.CreateElement("edge")
	//EdgeId := FindId(ElementGraph, ElementEdge)
	//EdgeID := EdgeId
	EdgeID := "e" + strconv.Itoa(ElementEdge.Index())
	ElementEdge.CreateAttr("id", EdgeID)
	//Source := "n" + strconv.Itoa(IndexElementFrom) + "::" + "n" + strconv.Itoa(IndexElementTo)
	IdFrom := FindId(ElementGraph, ElementFrom)
	IdTo := FindId(ElementGraph, ElementTo)
	ElementEdge.CreateAttr("source", IdFrom)
	ElementEdge.CreateAttr("target", IdTo)

	//data
	ElementData := ElementEdge.CreateElement("data")
	ElementData.CreateAttr("key", "d8")
	ElementData.CreateAttr("xml:space", "preserve")
	//ElementData.CreateText("<![CDATA[descr]]>")
	//ElementData.CreateElement("![CDATA[descr]]")
	ElementData.CreateCData(Description)

	//data2
	ElementData2 := ElementEdge.CreateElement("data")
	ElementData2.CreateAttr("key", "d9")

	//y:PolyLineEdge
	ElementYPolyLineEdge := ElementData2.CreateElement("y:PolyLineEdge")

	//y:Path
	ElementYPath := ElementYPolyLineEdge.CreateElement("y:Path")
	ElementYPath.CreateAttr("sx", "0.0")
	ElementYPath.CreateAttr("sy", "0.0")
	ElementYPath.CreateAttr("tx", "0.0")
	ElementYPath.CreateAttr("ty", "0.0")

	//y:LineStyle
	ElementYLineStyle := ElementYPolyLineEdge.CreateElement("y:LineStyle")
	ElementYLineStyle.CreateAttr("color", "#000000")
	ElementYLineStyle.CreateAttr("type", "line")
	ElementYLineStyle.CreateAttr("width", "1.0")

	//y:Arrows
	ElementYArrows := ElementYPolyLineEdge.CreateElement("y:Arrows")
	ElementYArrows.CreateAttr("source", "none")
	ElementYArrows.CreateAttr("target", "standard")

	//y:EdgeLabel
	ElementYEdgeLabel := ElementYPolyLineEdge.CreateElement("y:EdgeLabel")
	ElementYEdgeLabel.CreateAttr("alignment", "center")
	ElementYEdgeLabel.CreateAttr("configuration", "AutoFlippingLabel")
	ElementYEdgeLabel.CreateAttr("distance", "0.0")
	ElementYEdgeLabel.CreateAttr("fontFamily", "Dialog")
	ElementYEdgeLabel.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_EDGE))
	ElementYEdgeLabel.CreateAttr("fontStyle", "plain")
	ElementYEdgeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementYEdgeLabel.CreateAttr("hasLineColor", "false")
	ElementYEdgeLabel.CreateAttr("height", "17.96875")
	ElementYEdgeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementYEdgeLabel.CreateAttr("iconTextGap", "4")
	ElementYEdgeLabel.CreateAttr("modelName", "centered")
	ElementYEdgeLabel.CreateAttr("modelPosition", "head")
	ElementYEdgeLabel.CreateAttr("preferredPlacement", "anywhere")
	ElementYEdgeLabel.CreateAttr("ratio", "0.5")
	ElementYEdgeLabel.CreateAttr("textColor", "#000000")
	ElementYEdgeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementYEdgeLabel.CreateAttr("visible", "true")
	ElementYEdgeLabel.CreateAttr("width", "41.8")
	ElementYEdgeLabel.CreateAttr("x", "71.5")
	ElementYEdgeLabel.CreateAttr("xml:space", "preserve")
	ElementYEdgeLabel.CreateAttr("y", "0.5")
	ElementYEdgeLabel.CreateAttr("bottomInset", "0")
	ElementYEdgeLabel.CreateAttr("leftInset", "0")
	ElementYEdgeLabel.CreateAttr("rightInset", "0")
	ElementYEdgeLabel.CreateAttr("topInset", "0")
	ElementYEdgeLabel.CreateText(label)

	//y:PreferredPlacementDescriptor
	ElementYPreferredPlacementDescriptor := ElementYEdgeLabel.CreateElement("y:PreferredPlacementDescriptor")
	ElementYPreferredPlacementDescriptor.CreateAttr("angle", "0.0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleOffsetOnRightSide", "0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleReference", "absolute")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleRotationOnRightSide", "co")
	ElementYPreferredPlacementDescriptor.CreateAttr("distance", "-1.0")
	ElementYPreferredPlacementDescriptor.CreateAttr("frozen", "true")
	ElementYPreferredPlacementDescriptor.CreateAttr("placement", "anywhere")
	ElementYPreferredPlacementDescriptor.CreateAttr("side", "anywhere")
	ElementYPreferredPlacementDescriptor.CreateAttr("sideReference", "relative_to_edge_flow")

	//y:BendStyle
	ElementYBendStyle := ElementYPolyLineEdge.CreateElement("y:BendStyle")
	ElementYBendStyle.CreateAttr("smoothed", "false")

	return ElementEdge
}

// CreateElement_Edge_blue - создаёт элемент xgml - стрелка синяя с заголовком
func CreateElement_Edge_blue(ElementGraph, ElementFrom, ElementTo *etree.Element, label, Description string) *etree.Element {

	//node
	ElementEdge := ElementGraph.CreateElement("edge")
	//EdgeId := FindId(ElementGraph, ElementEdge)
	//EdgeID := EdgeId
	EdgeID := "e" + strconv.Itoa(ElementEdge.Index())
	ElementEdge.CreateAttr("id", EdgeID)
	//Source := "n" + strconv.Itoa(IndexElementFrom) + "::" + "n" + strconv.Itoa(IndexElementTo)
	IdFrom := FindId(ElementGraph, ElementFrom)
	IdTo := FindId(ElementGraph, ElementTo)
	ElementEdge.CreateAttr("source", IdFrom)
	ElementEdge.CreateAttr("target", IdTo)

	//data
	ElementData := ElementEdge.CreateElement("data")
	ElementData.CreateAttr("key", "d8")
	ElementData.CreateAttr("xml:space", "preserve")
	//ElementData.CreateText("<![CDATA[descr]]>")
	//ElementData.CreateElement("![CDATA[descr]]")
	ElementData.CreateCData(Description)

	//data2
	ElementData2 := ElementEdge.CreateElement("data")
	ElementData2.CreateAttr("key", "d9")

	//y:PolyLineEdge
	ElementYPolyLineEdge := ElementData2.CreateElement("y:PolyLineEdge")

	//y:Path
	ElementYPath := ElementYPolyLineEdge.CreateElement("y:Path")
	ElementYPath.CreateAttr("sx", "0.0")
	ElementYPath.CreateAttr("sy", "0.0")
	ElementYPath.CreateAttr("tx", "0.0")
	ElementYPath.CreateAttr("ty", "0.0")

	//y:LineStyle
	ElementYLineStyle := ElementYPolyLineEdge.CreateElement("y:LineStyle")
	ElementYLineStyle.CreateAttr("color", "#0000FF")
	ElementYLineStyle.CreateAttr("type", "line")
	ElementYLineStyle.CreateAttr("width", "1.0")

	//y:Arrows
	ElementYArrows := ElementYPolyLineEdge.CreateElement("y:Arrows")
	ElementYArrows.CreateAttr("source", "none")
	ElementYArrows.CreateAttr("target", "standard")

	//y:EdgeLabel
	ElementYEdgeLabel := ElementYPolyLineEdge.CreateElement("y:EdgeLabel")
	ElementYEdgeLabel.CreateAttr("alignment", "center")
	ElementYEdgeLabel.CreateAttr("configuration", "AutoFlippingLabel")
	ElementYEdgeLabel.CreateAttr("distance", "0.0")
	ElementYEdgeLabel.CreateAttr("fontFamily", "Dialog")
	ElementYEdgeLabel.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_EDGE))
	ElementYEdgeLabel.CreateAttr("fontStyle", "plain")
	ElementYEdgeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementYEdgeLabel.CreateAttr("hasLineColor", "false")
	ElementYEdgeLabel.CreateAttr("height", "17.96875")
	ElementYEdgeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementYEdgeLabel.CreateAttr("iconTextGap", "4")
	ElementYEdgeLabel.CreateAttr("modelName", "centered")
	ElementYEdgeLabel.CreateAttr("modelPosition", "head")
	ElementYEdgeLabel.CreateAttr("preferredPlacement", "anywhere")
	ElementYEdgeLabel.CreateAttr("ratio", "0.5")
	ElementYEdgeLabel.CreateAttr("textColor", "#0000FF")
	ElementYEdgeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementYEdgeLabel.CreateAttr("visible", "true")
	ElementYEdgeLabel.CreateAttr("width", "41.8")
	ElementYEdgeLabel.CreateAttr("x", "71.5")
	ElementYEdgeLabel.CreateAttr("xml:space", "preserve")
	ElementYEdgeLabel.CreateAttr("y", "0.5")
	ElementYEdgeLabel.CreateAttr("bottomInset", "0")
	ElementYEdgeLabel.CreateAttr("leftInset", "0")
	ElementYEdgeLabel.CreateAttr("rightInset", "0")
	ElementYEdgeLabel.CreateAttr("topInset", "0")
	ElementYEdgeLabel.CreateText(label)

	//y:PreferredPlacementDescriptor
	ElementYPreferredPlacementDescriptor := ElementYEdgeLabel.CreateElement("y:PreferredPlacementDescriptor")
	ElementYPreferredPlacementDescriptor.CreateAttr("angle", "0.0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleOffsetOnRightSide", "0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleReference", "absolute")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleRotationOnRightSide", "co")
	ElementYPreferredPlacementDescriptor.CreateAttr("distance", "-1.0")
	ElementYPreferredPlacementDescriptor.CreateAttr("frozen", "true")
	ElementYPreferredPlacementDescriptor.CreateAttr("placement", "anywhere")
	ElementYPreferredPlacementDescriptor.CreateAttr("side", "anywhere")
	ElementYPreferredPlacementDescriptor.CreateAttr("sideReference", "relative_to_edge_flow")

	//y:BendStyle
	ElementYBendStyle := ElementYPolyLineEdge.CreateElement("y:BendStyle")
	ElementYBendStyle.CreateAttr("smoothed", "false")

	return ElementEdge
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

// FindId - находит ИД в формате "n1::n1::n1"
func FindId(ElementGraph0, Element *etree.Element) string {
	Otvet := ""
	if Element == nil {
		return Otvet
	}

	//if Element == ElementGraph0 {
	//	return Otvet
	//}

	if Element.Tag == "node" {
		Otvet = "n" + strconv.Itoa(Element.Index())
		//return Otvet
	}

	ParentSID := FindId(ElementGraph0, Element.Parent())
	if ParentSID != "" {
		if Otvet == "" {
			Otvet = ParentSID
		} else {
			Otvet = ParentSID + "::" + Otvet
		}
	}

	return Otvet
}
