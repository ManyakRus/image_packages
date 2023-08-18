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
	//Group1 := CreateElement_Group(ElementGraph, nil, "GroupCaption1")
	//
	//Element1 := CreateElement_Shape(ElementGraph, Group1, "Entity1")
	CreateElement_Shape(ElementGraph, nil, "Entity2")
	//CreateElement_Edge(ElementGraph, Element1.Index(), Element2.Index())
	//
	//CreateElement_Edge_blue(ElementGraph, Element1.Index(), Element2.Index(), "test()")

	FileName := dir + "test" + micro.SeparatorFile() + "test.graphml"
	DocXML.IndentTabs()
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
