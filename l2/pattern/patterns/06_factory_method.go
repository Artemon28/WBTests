package patterns

import "fmt"

/*
	Реализовать паттерн «Фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Так как у нас GO, мы не можем реализовать фабрику со всей ООП функциональностью. Поэтому мы делаем упрощенную фабрику.
Есть интерфейс.
*/

/*
Метод нужен, чтобы создавать объекты разных типов, но обладающих одним функционалом. Это необходимо, так как мы можем
заранее не знать эти типы и их зависимости с которыми должен работать код. (Например можно выпускать только электро
машины, но до сих пор используются и ДВС)
*/

/*
Плюсы:
1) Избавляет от привязки к конкретным классам
2) Выделение создание объекта в отдельное мест
3)Упрощение создания новых объектов, нен нужно так сильно разбираться в коде, только в интерфейсе
*/

/*
минусы:
1) Для каждого класса необходима своя структура, что нагружает код повторениями
*/

/*
пример на практике
1) Расширение библиотеки, например можно написать фабрику по созданию какой-то структуры и переопределить её параметры
под нужные себе и использовать уже свою структуру далее
2) Экономия места, так как фабрика будет следидить вернуть клиенту ссылку на уже существующий объект или создать новый?????
*/

type ITransport interface {
	setName(name string)
	setEngineType(engType string)
	getName() string
	getEngineType() string
}

type Transport struct {
	name    string
	engType string
}

func (t *Transport) setName(name string) {
	t.name = name
}

func (t *Transport) setEngineType(engType string) {
	t.engType = engType
}

func (t *Transport) getName() string {
	return t.name
}
func (t *Transport) getEngineType() string {
	return t.engType
}

type Tesla struct {
	Transport
}

func NewTesla() ITransport {
	return &Tesla{
		Transport: Transport{
			name:    "Tesla",
			engType: "Electric",
		},
	}
}

type Lada struct {
	Transport
}

func NewLada() ITransport {
	return &Tesla{
		Transport: Transport{
			name:    "Lada",
			engType: "ICE",
		},
	}
}

//Fabric
func getTransport(autoName string) ITransport {
	if autoName == "Tesla" {
		return NewTesla()
	} else if autoName == "Lada" {
		return NewLada()
	}
	return nil
}

func main() {
	electricCar := getTransport("Tesla")
	fmt.Println(electricCar)
}
