package patterns

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
В данном примере у нас есть робот. Его ноги и руки программируются по отдельности, юному пользователю
будет тяжело программировать робота самому, поэтому добавим в фасаде функции ходьбы и приветствия
Также есть многоуровневость, когда мы используем фасад для одной группы сильно связанных классов и методов,
а потом объединяем это в конечном фасаде
*/

/*
Данный паттерн необходим, когда у нас есть большой функционал (мы могли написать его сами или взять из библиотеки)
В данном функционале тяжело разобраться и потребуется время на это. Поэтому мы выделяем в фасаде функции,
необходимые клиенту, чтобы он мог пользоваться ими, не разбираясь в запутанной логике

Для данного примера можно добавить больше уровней, напрмер как реализовано сгибание колена? это надо активировать
моторчик, согнуть на нужное количесвто градусов, то есть активировать его движение с мощностью N на время T
*/
/*
Плюсы:
1) Удобный функционал для клиента
2) Обеспечение целостности данных и правильности работы системы, так как все функции составляет программист
*/

/*
минусы:
1) Риск "божественного объекта"
Фасад становится большим, начинает отвечать за работу всего кода, в этом случае весь код завязывается на одном узле
2) Урезанный функционал
*/

/*
пример на практике
1) Упрощенная работа с большими библиотеками
2) Для написания своей библиотеки, чтобы спрятать логику и предоставить красивое API
*/

import (
	"fmt"
	"strconv"
)

type leg struct {
	model string
}

func newLeg(m string) *leg {
	return &leg{
		model: m,
	}
}

func (l leg) bendTheKnee() {
	fmt.Println("робот согнул колено ноги " + l.model)
}

func (l leg) straightenTheKnee() {
	fmt.Println("робот разогнул колено ноги " + l.model)
}

func (l leg) moveForward(n int) {
	fmt.Println("робот передвинул ногу " + l.model + " вперёд на " + strconv.Itoa(n) + " сантиметров")
}

type arm struct {
	model string
}

func newArm(m string) *arm {
	return &arm{
		model: m,
	}
}

func (a arm) riseArm() {
	fmt.Println("Робот поднял руку " + a.model)
}

func (a arm) shakeArm() {
	fmt.Println("Робот помахал рукой " + a.model)
}

func (a arm) lowerArm() {
	fmt.Println("Робот опустил руку " + a.model)
}

type speaker struct {
}

func (a speaker) say(phrase string) {
	fmt.Println(phrase)
}

type Robot struct {
	leftLeg  *leg
	rightLeg *leg
	leftArm  *arm
	rightArm *arm
	mouth    *speaker
}

func newRobot(legsModel, armsModel string) *Robot {
	return &Robot{
		leftLeg:  newLeg("L" + legsModel),
		rightLeg: newLeg("R" + legsModel),
		leftArm:  newArm("L" + armsModel),
		rightArm: newArm("R" + armsModel),
	}
}

func (r Robot) Walk() {
	r.leftLeg.bendTheKnee()
	r.leftLeg.moveForward(10)
	r.leftLeg.straightenTheKnee()

	r.rightLeg.bendTheKnee()
	r.rightLeg.moveForward(10)
	r.rightLeg.straightenTheKnee()
}

func (r Robot) Greetings() {
	r.leftArm.riseArm()
	r.leftArm.shakeArm()
	r.mouth.say("Hello there!")
	r.leftArm.lowerArm()
}

func main() {
	rob := newRobot("to1", "to2")
	rob.Greetings()
	rob.Walk()
}
