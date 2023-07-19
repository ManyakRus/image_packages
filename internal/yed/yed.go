package yed

import (
	"strings"
)

func AddDirectory(buffer *strings.Builder, Name string) {

}

func CreateLinkXGML(ЗаписьXML, Начало, Конец) {



SaveStartSectionXGML(ЗаписьXML, "edge")

SaveAttributeXGML(ЗаписьXML, "source", "string", Начало)

SaveAttributeXGML(ЗаписьXML, "target", "string", Конец)



SaveStartSectionXGML(ЗаписьXML, "graphics")

SaveAttributeXGML(ЗаписьXML, "fill", "string", "#000000")

SaveAttributeXGML(ЗаписьXML, "targetArrow", "string", "standard")

ЗаписьXML.ЗаписатьКонецЭлемента() //graphics



ЗаписьXML.ЗаписатьКонецЭлемента() //edge



} // CreateLinkXGML()







//возвращает
func FindWidthBlock(ИмяЭлемента){

Otvet = 10



СимволовМакс = НайтиКоличествоСимволовВСтрокеМакс(ИмяЭлемента)

Otvet = Otvet + СимволовМакс*7.1

Otvet = Окр(Otvet, 0)





Возврат Otvet

}



//возвращает
func FindHeightBlock(ИмяЭлемента){

Otvet = 30



ВсегоСтрок = СтрЧислоСтрок(ИмяЭлемента)

Otvet = Otvet + (ВсегоСтрок-1)*18





Возврат Otvet

}



//санек вид как UML 25.10.2019 10:05:41  

func CreateElementXGML_UML(ЗаписьXML, ГруппаЭлемента = "", ИДЭлемента = "", ИмяЭлемента = ""){



Если ИДЭлемента = "" Тогда

ИДЭлемента = ИмяЭлемента

КонецЕсли



Ширина = FindWidthBlock(ИмяЭлемента)

Высота = FindHeightBlock(ИмяЭлемента)





//node

SaveStartSectionXGML(ЗаписьXML, "node")



SaveAttributeXGML(ЗаписьXML, "id", "string", ИДЭлемента)

SaveAttributeXGML(ЗаписьXML, "label", "string", "")

//SaveAttributeXGML(ЗаписьXML, "label", "string", ИмяЭлемента)



//graphics

SaveStartSectionXGML(ЗаписьXML, "graphics")

SaveAttributeXGML(ЗаписьXML, "type", "string", "rectangle")

SaveAttributeXGML(ЗаписьXML, "fill", "string", "#FFFFFF") //было #FFCC00

SaveAttributeXGML(ЗаписьXML, "outline", "string", "#000000")

SaveAttributeXGML(ЗаписьXML, "customconfiguration", "string", "com.yworks.entityRelationship.big_entity")



SaveAttributeXGML(ЗаписьXML, "h", "double", СтрокаЧ(Высота))

SaveAttributeXGML(ЗаписьXML, "w", "double", СтрокаЧ(Ширина))



//styleproperties

SaveStartSectionXGML(ЗаписьXML, "styleproperties")



//property

SaveStartSectionXGML(ЗаписьXML, "property")

SaveAttributeXGML(ЗаписьXML, "name", "string", "y.view.ShadowNodePainter.SHADOW_PAINTING")

SaveAttributeXGML(ЗаписьXML, "valueClass", "string", "java.lang.Boolean")

SaveAttributeXGML(ЗаписьXML, "value", "string", "true")

ЗаписьXML.ЗаписатьКонецЭлемента() //property



//

ЗаписьXML.ЗаписатьКонецЭлемента() //styleproperties



//

ЗаписьXML.ЗаписатьКонецЭлемента() //graphics



//LabelGraphics-1

SaveStartSectionXGML(ЗаписьXML, "LabelGraphics")

SaveAttributeXGML(ЗаписьXML, "text", "String", ИДЭлемента)

SaveAttributeXGML(ЗаписьXML, "fontSize", "int", "12")

SaveAttributeXGML(ЗаписьXML, "configuration", "String", "com.yworks.entityRelationship.label.name")

SaveAttributeXGML(ЗаписьXML, "anchor", "String", "t")

SaveAttributeXGML(ЗаписьXML, "contentWidth", "int", "24")

SaveAttributeXGML(ЗаписьXML, "contentHeight", "int", "18")

ЗаписьXML.ЗаписатьКонецЭлемента()



//LabelGraphics-2

SaveStartSectionXGML(ЗаписьXML, "LabelGraphics")

SaveAttributeXGML(ЗаписьXML, "text", "String", ИмяЭлемента)

SaveAttributeXGML(ЗаписьXML, "fontSize", "int", "12")

SaveAttributeXGML(ЗаписьXML, "configuration", "String", "com.yworks.entityRelationship.label.attributes")

SaveAttributeXGML(ЗаписьXML, "alignment", "String", "left")

SaveAttributeXGML(ЗаписьXML, "contentWidth", "int", "24")

SaveAttributeXGML(ЗаписьXML, "contentHeight", "int", "18")

ЗаписьXML.ЗаписатьКонецЭлемента()



//gid

Если НЕ ПустаяСтрока(ГруппаЭлемента) Тогда

SaveAttributeXGML(ЗаписьXML, "gid", "String", ГруппаЭлемента)

КонецЕсли



//

ЗаписьXML.ЗаписатьКонецЭлемента() //node



}



func CreateElementXGML_Standart(ЗаписьXML, ГруппаЭлемента = "", ИДЭлемента = "", ИмяЭлемента = "", ВидОтображения = "Группами"){



Если ИДЭлемента = "" Тогда

ИДЭлемента = ИмяЭлемента

КонецЕсли



SaveStartSectionXGML(ЗаписьXML, "node")



SaveAttributeXGML(ЗаписьXML, "id", "string", ИДЭлемента)

SaveAttributeXGML(ЗаписьXML, "label", "string", ИмяЭлемента)



SaveStartSectionXGML(ЗаписьXML, "graphics")

SaveAttributeXGML(ЗаписьXML, "type", "string", "rectangle")

SaveAttributeXGML(ЗаписьXML, "fill", "string", "#FFFFFF") //было #FFCC00



//санек НАЧАЛО 24.10.2019 10:09:41

Ширина = FindWidthBlock(ИмяЭлемента)

Высота = FindHeightBlock(ИмяЭлемента)



Ширина2 = Окр(FindWidthBlock(ИДЭлемента), 2)

Высота2 = Окр(FindHeightBlock(ИДЭлемента), 2)

Если Ширина2 > Ширина Тогда

Ширина = Ширина2

КонецЕсли

Если Ширина2 > Ширина Тогда

Ширина = Ширина2

КонецЕсли



Если ВидОтображения = "Группами" Тогда

Ширина = Окр(Ширина*1.3, 0)

//Высота = Окр(Высота*1.3, 0)

SaveAttributeXGML(ЗаписьXML, "outline", "string", "#FFFFFF")

Иначе

SaveAttributeXGML(ЗаписьXML, "outline", "string", "#000000")

КонецЕсли

SaveAttributeXGML(ЗаписьXML, "h", "double", СтрокаЧ(Высота))

SaveAttributeXGML(ЗаписьXML, "w", "double", СтрокаЧ(Ширина))

//санек КОНЕЦ 



ЗаписьXML.ЗаписатьКонецЭлемента() //graphics



SaveStartSectionXGML(ЗаписьXML, "LabelGraphics")

SaveAttributeXGML(ЗаписьXML, "text", "String", ИмяЭлемента)

SaveAttributeXGML(ЗаписьXML, "fontSize", "int", "12")

ЗаписьXML.ЗаписатьКонецЭлемента() //LabelGraphics



Если НЕ ПустаяСтрока(ГруппаЭлемента) Тогда

SaveAttributeXGML(ЗаписьXML, "gid", "String", ГруппаЭлемента)

КонецЕсли



ЗаписьXML.ЗаписатьКонецЭлемента() //node



}





func CreateElementXGML_WithGroup(ЗаписьXML, ГруппаЭлемента = "", ИДЭлемента = "", ИмяЭлемента = "", ВидОтображения = "Группами"){



Если ИДЭлемента = "" Тогда

ИДЭлемента = ИмяЭлемента

КонецЕсли



CreateElementXGML_Standart(ЗаписьXML, ИДЭлемента, ИДЭлемента, ИмяЭлемента, ВидОтображения)



CreateGroupXGML(ЗаписьXML, ИДЭлемента, ИДЭлемента, ГруппаЭлемента)





}





func CreateElementXGML(ЗаписьXML, ГруппаЭлемента = "", ИДЭлемента = "", ИмяЭлемента = "", ВидОтображения = "Группами"){



Если ВидОтображения = "Группами" Тогда

CreateElementXGML_WithGroup(ЗаписьXML, ГруппаЭлемента, ИДЭлемента, ИмяЭлемента, ВидОтображения)

Иначе

Текст = ИДЭлемента + Символы.ПС + ИмяЭлемента

CreateElementXGML_Standart(ЗаписьXML, ГруппаЭлемента, ИДЭлемента, Текст, ВидОтображения)

КонецЕсли





}



func CreateGroupXGML(ЗаписьXML, ИмяГруппы, ЗаголовокГруппы = "", ИДГруппы = ""){



Если ПустаяСтрока(ЗаголовокГруппы) Тогда

ЗаголовокГруппы = ИмяГруппы

ЗаголовокГруппы = ИмяОбъектаВоМножественномЧисле(ЗаголовокГруппы)

КонецЕсли



SaveStartSectionXGML(ЗаписьXML, "node")



SaveAttributeXGML(ЗаписьXML, "id", "string", ИмяГруппы)

SaveAttributeXGML(ЗаписьXML, "label", "string", ЗаголовокГруппы)



SaveStartSectionXGML(ЗаписьXML, "graphics")

SaveAttributeXGML(ЗаписьXML, "type", "string", "roundrectangle")

SaveAttributeXGML(ЗаписьXML, "fill", "string", "#F5F5F5")

SaveAttributeXGML(ЗаписьXML, "outline", "string", "#000000")



////санек НАЧАЛО 24.10.2019 10:09:41

//Ширина = 10 + FindWidthBlock(ЗаголовокГруппы)

//Высота = 10 + FindHeightBlock(ЗаголовокГруппы)

//SaveAttributeXGML(ЗаписьXML, "h", "double", СтрокаЧ(Высота))

//SaveAttributeXGML(ЗаписьXML, "w", "double", СтрокаЧ(Ширина))

////санек КОНЕЦ 



ЗаписьXML.ЗаписатьКонецЭлемента() //graphics



SaveStartSectionXGML(ЗаписьXML, "LabelGraphics")

SaveAttributeXGML(ЗаписьXML, "text", "String", ЗаголовокГруппы)

SaveAttributeXGML(ЗаписьXML, "fill", "String", "#EBEBEB")

SaveAttributeXGML(ЗаписьXML, "fontSize", "int", "16")

SaveAttributeXGML(ЗаписьXML, "anchor", "String", "t")

ЗаписьXML.ЗаписатьКонецЭлемента() //LabelGraphics



SaveAttributeXGML(ЗаписьXML, "isGroup", "boolean", "true")



Если ЗначениеЗаполнено(ИДГруппы) Тогда

SaveAttributeXGML(ЗаписьXML, "gid", "String", ИДГруппы)

КонецЕсли



ЗаписьXML.ЗаписатьКонецЭлемента() //node



} // CreateGroupXGML()



func SaveStartSectionXGML(ЗаписьXML, name){



ЗаписьXML.ЗаписатьНачалоЭлемента("section")

ЗаписьXML.ЗаписатьАтрибут("name", name)



} // ЗаписатьНачалоСекции()



func SaveAttributeXGML(ЗаписьXML, key, type, text) {



ЗаписьXML.ЗаписатьНачалоЭлемента("attribute")

ЗаписьXML.ЗаписатьАтрибут("key", key)

ЗаписьXML.ЗаписатьАтрибут("type", type)

ЗаписьXML.ЗаписатьТекст(text)

ЗаписьXML.ЗаписатьКонецЭлемента() //attribute



} // SaveAttributeXGML()



