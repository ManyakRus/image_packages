package xgml

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestCreateNewXGML(t *testing.T) {
	dir := micro.ProgramDir()
	DocXML := CreateDocXGML()
	ElementGraph := DocXML.FindElement("/section/section")

	Group1 := CreateGroupXGML(ElementGraph, nil, "GroupCaption1")

	Element1 := CreateElementXGML_Shape(ElementGraph, Group1, "Entity1")
	Element2 := CreateElementXGML_Shape(ElementGraph, nil, "Entity2")
	CreateLinkXGML(ElementGraph, Element1.Index(), Element2.Index())

	CreateLinkXGML_blue(ElementGraph, Element1.Index(), Element2.Index(), "test()")

	FileName := dir + "test" + micro.SeparatorFile() + "test.xgml"
	//DocXML.IndentTabs()
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
