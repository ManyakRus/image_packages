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

	//ElementGraph := DocXML.FindElement("/section/section")
	//
	Group1 := CreateElement_Group(ElementGraph, nil, "Group1")
	//
	//Element1 := CreateElement_Shape(ElementGraph, Group1, "Entity1")
	CreateElement_Shape(Group1, nil, "Shape1")
	//CreateElement_Edge(ElementGraph, Element1.Index(), Element2.Index())
	//
	//CreateElement_Edge_blue(ElementGraph, Element1.Index(), Element2.Index(), "test()")

	FileName := dir + "test" + micro.SeparatorFile() + "test.graphml"
	//DocXML.IndentTabs()
	DocXML.Indent(2)
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
