package yed

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestCreateNewXGML(t *testing.T) {
	dir := micro.ProgramDir()
	DocXML := CreateNewXGML("")
	ElementGraph := DocXML.FindElement("/section/section")
	Element1 := CreateElementXGML_Standart(ElementGraph, "", "Entity1", "")
	Element2 := CreateElementXGML_Standart(ElementGraph, "", "Entity2", "")
	CreateLinkXGML(ElementGraph, Element1.Index(), Element2.Index())

	FileName := dir + "test" + micro.SeparatorFile() + "test.xgml"
	DocXML.IndentTabs()
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
