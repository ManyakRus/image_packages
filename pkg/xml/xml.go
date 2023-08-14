package xml

import (
	"github.com/beevik/etree"
	"strconv"
)

// CreateDocXGML - создаёт новый документ .xgml
func CreateDocXGML() *etree.Document {

	DocXML := etree.NewDocument()
	DocXML.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	//DocXML.IndentTabs()

	ElementXGML := DocXML.CreateElement("section")
	ElementXGML.CreateAttr("name", "xgml")

	//ElementXGML := AddSectionXML(DocXML, "xgml")
	ElementGraph := AddSectionXML(ElementXGML, "graph")
	AddAttributeXML(ElementGraph, "hierarchic", "int", "1")

	//log.Info("ElementGraph.GetPath(): ", ElementGraph.GetPath())

	return DocXML
}

// AddSectionXML - добавляет секцию в xgml
func AddSectionXML(Element *etree.Element, name string) *etree.Element {

	Element1 := Element.CreateElement("section")
	Element1.CreateAttr("name", name)

	return Element1
}

// AddAttributeXML - добавляет аттрибут в элемент xgml
func AddAttributeXML(Element *etree.Element, key, stype, text string) *etree.Element {

	ElementAttribute := Element.CreateElement("attribute")
	ElementAttribute.CreateAttr("key", key)
	ElementAttribute.CreateAttr("type", stype)
	ElementAttribute.SetText(text)

	return ElementAttribute
}

// AddAttributeXML_int - добавляет аттрибут типа int в элемент xgml
func AddAttributeXML_int(Element *etree.Element, key string, value int) *etree.Element {

	ElementAttribute := AddAttributeXML(Element, key, "int", strconv.Itoa(value))

	return ElementAttribute
}

// AddAttributeXML_double - добавляет аттрибут типа double в элемент xgml
func AddAttributeXML_double(Element *etree.Element, key string, value float64) *etree.Element {

	ElementAttribute := AddAttributeXML(Element, key, "double", strconv.Itoa(int(value)))

	return ElementAttribute
}

// AddAttributeXML_string - добавляет аттрибут типа string в элемент xgml
func AddAttributeXML_string(Element *etree.Element, key string, value string) *etree.Element {

	ElementAttribute := AddAttributeXML(Element, key, "String", value)

	return ElementAttribute
}
