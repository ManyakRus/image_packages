package yed

import (
	"github.com/ManyakRus/starter/log"
	"github.com/beevik/etree"
	_ "github.com/beevik/etree"
	"math"
	"strconv"
	"strings"
)

//var doc = etree.NewDocument()

func CreateNewXGML(ShowKind string) *etree.Document {

	DocXML := etree.NewDocument()
	DocXML.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	ElementXGML := DocXML.CreateElement("section")
	ElementXGML.CreateAttr("name", "xgml")

	//ElementXGML := AddSectionXML(DocXML, "xgml")
	ElementGraph := AddSectionXML(ElementXGML, "graph")
	AddAttributeXML(ElementGraph, "hierarchic", "int", "1")

	log.Info("ElementGraph.GetPath(): ", ElementGraph.GetPath())

	//DocXML.УстановитьСтроку()
	//
	//DocXML.ЗаписатьОбъявлениеXML()
	//
	//сМинус = " -"

	//DocXML.ЗаписатьНачалоЭлемента("section")
	//
	//DocXML.ЗаписатьАтрибут("name", "xgml")

	//ЗаписатьАтрибутXGML(DocXML, "Creator", "String", "yFiles"); //тест

	//ЗаписатьАтрибутXGML(DocXML, "Version", "String", "2.16"); //тест

	//DocXML.ЗаписатьНачалоЭлемента("section")
	//
	//DocXML.ЗаписатьАтрибут("name", "graph")

	//ЗаписатьАтрибутXGML(DocXML, "hierarchic", "String", "1") //тест

	//ВыборкаТипОбъекта = РезультатЗапроса.Выбрать(ОбходРезультатаЗапроса.ПоГруппировкам)
	//
	//Пока
	//ВыборкаТипОбъекта.Следующий()
	//Цикл
	//
	//ТипОбъекта = ВыборкаТипОбъекта.ТипОбъекта
	//
	//СоздатьГруппуXGML(DocXML, ТипОбъекта)
	//
	//ВыборкаВидОбъекта = ВыборкаТипОбъекта.Выбрать(ОбходРезультатаЗапроса.ПоГруппировкам)
	//
	//Пока
	//ВыборкаВидОбъекта.Следующий()
	//Цикл
	//
	//ВидОбъекта = ВыборкаВидОбъекта.ВидОбъекта
	//
	//ТипВид = ВыборкаВидОбъекта.ТипВид
	//
	//МассивНужныеОбъекты = Новый
	//Массив
	//
	//ИмяРеквизита = ""
	//
	//Запятая = ""
	//
	//ВыборкаНужныйОбъект = ВыборкаВидОбъекта.Выбрать(ОбходРезультатаЗапроса.ПоГруппировкам)
	//
	//Пока
	//ВыборкаНужныйОбъект.Следующий()
	//Цикл
	//
	//НужныйОбъект = ВыборкаНужныйОбъект.НужныйОбъект
	//
	//НужныйОбъект = УбратьСловоСсылка(НужныйОбъект)
	//
	//ВидНужныйОбъект = НайтиВидИзПолногоИмени(НужныйОбъект)
	//
	//Выборка = ВыборкаНужныйОбъект.Выбрать()
	//
	//Пока
	//Выборка.Следующий()
	//Цикл
	//
	//ИмяРеквизита1 = Выборка.ИмяРеквизита
	//
	//Имя = ВидОбъекта
	//
	//Если
	//ЗначениеЗаполнено(ИмяРеквизита1) = Ложь
	//Тогда
	//
	//Продолжить
	//
	//КонецЕсли
	//
	//Имя = Имя + "." + ИмяРеквизита1
	//
	//ИмяРеквизита = ИмяРеквизита + Запятая + сМинус + ИмяРеквизита1
	//
	//Запятая = Символы.ПС
	//
	//КонецЦикла
	//
	//Если
	//ТипВид < > НужныйОбъект
	//Тогда
	//
	//МассивНужныеОбъекты.Добавить(НужныйОбъект)
	//
	//КонецЕсли
	//
	//КонецЦикла
	//
	////Если ЗначениеЗаполнено(ИмяРеквизита) = Истина Тогда
	//
	////ИмяРеквизита = ВидОбъекта + Символы.ПС + ИмяРеквизита;
	//
	//СоздатьЭлементXGML(DocXML, ТипОбъекта, ТипВид, ИмяРеквизита, ShowKind)
	//
	//Для
	//каждого
	//Массив1
	//из
	//МассивНужныеОбъекты
	//Цикл
	//
	////СоздатьСвязьXGML(DocXML, Массив1, ТипВид);
	//
	//НоваяСтрока = ТЗСвязи.Добавить()
	//
	//НоваяСтрока.Откуда = Массив1
	//
	//НоваяСтрока.Куда = ТипВид
	//
	//КонецЦикла
	//
	////КонецЕсли;
	//
	//КонецЦикла
	//
	//КонецЦикла
	//
	////создадим связи
	//
	//Для
	//каждого
	//ТЗ1
	//из
	//ТЗСвязи
	//Цикл
	//
	//СоздатьСвязьXGML(DocXML, ТЗ1.Откуда, ТЗ1.Куда)
	//
	//КонецЦикла
	//
	//DocXML.ЗаписатьКонецЭлемента() //graph
	//
	//DocXML.ЗаписатьКонецЭлемента() //xgml

	return DocXML

} // ТекстXGML()

func AddDirectory(buffer *strings.Builder, Name string) {

}

func CreateLinkXGML(ElementGraph *etree.Element, IndexElementFrom, IndexElementTo int) {

	ElementEdge := AddSectionXML(ElementGraph, "edge")
	AddAttributeXML_int(ElementEdge, "source", "int", IndexElementFrom)
	AddAttributeXML_int(ElementEdge, "target", "int", IndexElementTo)

	ElementGraphics := AddSectionXML(ElementEdge, "graphics")
	AddAttributeXML(ElementGraphics, "fill", "string", "#000000")
	AddAttributeXML(ElementGraphics, "targetArrow", "string", "standard")

	//Element.ЗаписатьКонецЭлемента() //graphics
	//
	//Element.ЗаписатьКонецЭлемента() //edge

} // CreateLinkXGML()

// возвращает
func FindWidthBlock(ElementName string) int {
	Otvet := 10

	LenMax := FindMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*7.1
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// возвращает
func FindHeightBlock(ElementName string) int {

	Otvet := 30

	RowsTotal := CountLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

func CountLines(s string) int {
	Otvet := 0

	strings.Count(s, "/n")

	return Otvet
}

//санек вид как UML 25.10.2019 10:05:41

func CreateElementXGML_UML(Element *etree.Element, ElementGroup, ElementId, ElementName string) *etree.Element {

	if ElementId == "" {
		ElementId = ElementName
	}

	Width := FindWidthBlock(ElementName)
	Height := FindHeightBlock(ElementName)

	//node

	ElementNode := AddSectionXML(Element, "node")
	AddAttributeXML(ElementNode, "id", "string", ElementId)
	AddAttributeXML(ElementNode, "label", "string", "")

	//AddAttributeXML(ElementNode, "label", "string", ElementName)

	//graphics

	ElementGraphics := AddSectionXML(ElementNode, "graphics")
	AddAttributeXML(ElementGraphics, "type", "string", "rectangle")
	AddAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00
	AddAttributeXML(ElementGraphics, "outline", "string", "#000000")
	AddAttributeXML(ElementGraphics, "customconfiguration", "string", "com.yworks.entityRelationship.big_entity")
	AddAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	AddAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//styleproperties

	ElementStyleProperties := AddSectionXML(ElementGraphics, "styleproperties")

	//property

	AddSectionXML(ElementStyleProperties, "property")
	AddAttributeXML(ElementStyleProperties, "name", "string", "y.view.ShadowNodePainter.SHADOW_PAINTING")
	AddAttributeXML(ElementStyleProperties, "valueClass", "string", "java.lang.Boolean")
	AddAttributeXML(ElementStyleProperties, "value", "string", "true")

	//Element.ЗаписатьКонецЭлемента() //property

	//

	//Element.ЗаписатьКонецЭлемента() //styleproperties

	//

	//Element.ЗаписатьКонецЭлемента() //graphics

	//LabelGraphics-1

	ElementLabelGraphics := AddSectionXML(ElementGraphics, "LabelGraphics")
	AddAttributeXML(ElementLabelGraphics, "text", "String", ElementId)
	AddAttributeXML(ElementLabelGraphics, "fontSize", "int", "12")
	AddAttributeXML(ElementLabelGraphics, "configuration", "String", "com.yworks.entityRelationship.label.name")
	AddAttributeXML(ElementLabelGraphics, "anchor", "String", "t")
	AddAttributeXML(ElementLabelGraphics, "contentWidth", "int", "24")
	AddAttributeXML(ElementLabelGraphics, "contentHeight", "int", "18")

	//Element.ЗаписатьКонецЭлемента()

	//LabelGraphics-2

	ElementLabelGraphics2 := AddSectionXML(ElementGraphics, "LabelGraphics")
	AddAttributeXML(ElementLabelGraphics2, "text", "String", ElementName)
	AddAttributeXML(ElementLabelGraphics2, "fontSize", "int", "12")
	AddAttributeXML(ElementLabelGraphics2, "configuration", "String", "com.yworks.entityRelationship.label.attributes")
	AddAttributeXML(ElementLabelGraphics2, "alignment", "String", "left")
	AddAttributeXML(ElementLabelGraphics2, "contentWidth", "int", "24")
	AddAttributeXML(ElementLabelGraphics2, "contentHeight", "int", "18")

	//Element.ЗаписатьКонецЭлемента()

	//gid

	if (ElementGroup) != "" {
		AddAttributeXML(Element, "gid", "String", ElementGroup)
	}

	//

	//Element.ЗаписатьКонецЭлемента() //node

	return ElementNode
}

func CreateElementXGML_Standart(Element *etree.Element, ElementGroup, ElementName, ShowKind string) *etree.Element {

	ElementNode := AddSectionXML(Element, "node")
	AddAttributeXML(ElementNode, "id", "int", strconv.Itoa(ElementNode.Index()))
	AddAttributeXML(ElementNode, "label", "string", ElementName)

	ElementGraphics := AddSectionXML(ElementNode, "graphics")
	AddAttributeXML(ElementGraphics, "type", "string", "rectangle")
	AddAttributeXML(ElementGraphics, "fill", "string", "#FFFFFF") //было #FFCC00

	//санек Begin 24.10.2019 10:09:41

	Width := FindWidthBlock(ElementName)
	Height := FindHeightBlock(ElementName)
	//Width2 := FindWidthBlock(ElementId)

	//Height2 := FindHeightBlock(ElementId)

	//if Width2 > Width {
	//
	//	Width = Width2
	//
	//}
	//
	//if Width2 > Width {
	//
	//	Width = Width2
	//
	//}

	if ShowKind == "Группами" {

		Width = int(math.Round(float64(Width) * 1.3))

		//Высота = Окр(Высота*1.3, 0)

		AddAttributeXML(ElementGraphics, "outline", "string", "#FFFFFF")

	} else {

		AddAttributeXML(ElementGraphics, "outline", "string", "#000000")

	}

	AddAttributeXML(ElementGraphics, "h", "double", strconv.Itoa(Height))
	AddAttributeXML(ElementGraphics, "w", "double", strconv.Itoa(Width))

	//санек КОНЕЦ

	//Element.ЗаписатьКонецЭлемента() //graphics

	ElementLabelGraphics := AddSectionXML(ElementNode, "LabelGraphics")

	AddAttributeXML(ElementLabelGraphics, "text", "String", ElementName)
	AddAttributeXML(ElementLabelGraphics, "fontSize", "int", "12")

	//Element.ЗаписатьКонецЭлемента() //LabelGraphics

	if ElementGroup != "" {
		AddAttributeXML(ElementNode, "gid", "String", ElementGroup)
	}

	//Element.ЗаписатьКонецЭлемента() //node

	return ElementNode
} // ВидОтображения = "Группами"

func CreateElementXGML_WithGroup(Element *etree.Element, ElementGroup, ElementName, ВидОтображения string) { // ВидОтображения = "Группами"

	//if ElementId == "" {
	//	ElementId = ElementName
	//}

	CreateElementXGML_Standart(Element, ElementGroup, ElementName, ВидОтображения)
	CreateGroupXGML(Element, ElementName, ElementGroup)

}

func CreateElementXGML(Element *etree.Element, ElementGroup, ElementId, ElementName, ShoqKind string) {

	if ShoqKind == "Группами" {

		CreateElementXGML_WithGroup(Element, ElementGroup, ElementName, ShoqKind)

	} else {

		Текст := ElementId + "\n" + ElementName

		CreateElementXGML_Standart(Element, ElementGroup, Текст, ShoqKind)

	}

}

func CreateGroupXGML(Element *etree.Element, GroupName, GroupCaption string) {

	if GroupCaption == "" {
		GroupCaption = GroupName
		//GroupCaption = ИмяОбъектаВоМножественномЧисле(GroupCaption)
	}

	AddSectionXML(Element, "node")

	AddAttributeXML(Element, "id", "string", GroupName)

	AddAttributeXML(Element, "label", "string", GroupCaption)

	AddSectionXML(Element, "graphics")

	AddAttributeXML(Element, "type", "string", "roundrectangle")

	AddAttributeXML(Element, "fill", "string", "#F5F5F5")

	AddAttributeXML(Element, "outline", "string", "#000000")

	////санек Begin 24.10.2019 10:09:41

	//Ширина = 10 + FindWidthBlock(GroupCaption)

	//Высота = 10 + FindHeightBlock(GroupCaption)

	//AddAttributeXML(Element, "h", "double", СтрокаЧ(Высота))

	//AddAttributeXML(Element, "w", "double", СтрокаЧ(Ширина))

	////санек КОНЕЦ

	//Element.ЗаписатьКонецЭлемента() //graphics

	AddSectionXML(Element, "LabelGraphics")

	AddAttributeXML(Element, "text", "String", GroupCaption)

	AddAttributeXML(Element, "fill", "String", "#EBEBEB")

	AddAttributeXML(Element, "fontSize", "int", "16")

	AddAttributeXML(Element, "anchor", "String", "t")

	//Element.ЗаписатьКонецЭлемента() //LabelGraphics

	AddAttributeXML(Element, "isGroup", "boolean", "true")

	//if GroupId != "" {
	//	AddAttributeXML(Element, "gid", "String", GroupId)
	//}

	//Element.ЗаписатьКонецЭлемента() //node

} // CreateGroupXGML()

func AddSectionXML(Element *etree.Element, name string) *etree.Element {

	Element1 := Element.CreateElement("section")
	Element1.CreateAttr("name", name)

	//AddAttributeXML(Element1, "id", "int", strconv.Itoa(Element1.Index()))

	//Element.ЗаписатьBeginЭлемента("section")
	//Element.ЗаписатьАтрибут("name", name)

	return Element1
} // ЗаписатьBeginСекции()

func AddAttributeXML(Element *etree.Element, key, stype, text string) *etree.Element {

	ElementAttribute := Element.CreateElement("attribute")
	ElementAttribute.CreateAttr("key", key)
	ElementAttribute.CreateAttr("type", stype)
	ElementAttribute.SetText(text)

	return ElementAttribute
}

func AddAttributeXML_int(Element *etree.Element, key, stype string, value int) *etree.Element {

	ElementAttribute := AddAttributeXML(Element, key, stype, strconv.Itoa(value))

	return ElementAttribute
}

// возвращает количество символов в строке максимум
func FindMaxLenRow(ElementName string) int {
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
