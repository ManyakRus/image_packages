package graphml

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestCreateNewGraphml(t *testing.T) {
	dir := micro.ProgramDir()
	DocXML, ElementGraph := CreateDocument()
	if ElementGraph == nil {

	}

	Shape2 := CreateElement_Shape(ElementGraph, "Shape2")
	//ElementGraph := DocXML.FindElement("/section/section")
	//
	Group1 := CreateElement_Group(ElementGraph, "Group1")
	//
	//Element1 := CreateElement_Shape(ElementGraph, Group1, "Entity1")
	Shape1 := CreateElement_Shape(Group1, "Shape1")
	CreateElement_Edge(ElementGraph, Shape1, Shape2, "edge1", "descr")
	CreateElement_Edge_blue(ElementGraph, Shape2, Shape1, "edge2", "descr2")
	//
	//CreateElement_Edge_blue(ElementGraph, Element1.Index(), Element2.Index(), "test()")

	if Shape1 == nil || Shape2 == nil {

	}

	FileName := dir + "test" + micro.SeparatorFile() + "test.graphml"
	//DocXML.IndentTabs()
	DocXML.Indent(2)
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
