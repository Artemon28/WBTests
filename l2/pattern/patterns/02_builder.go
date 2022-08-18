package patterns

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Паттерн применяется, когда мы хотим задать различные характеристики предмету. При этом мы не хотим создавать
каждый раз новый класс или подкласс, отличающиеся несколькими характеристиками.При этом сложный объект задаётся пошагово
Один и тот же код может использоваться повторно или не использоваться (никто не заставляет обязательно его прописывать)

Упрощает задание большого объекта, скоращая его конструктор. Есть возможность задать различные представления объекта

Вроде должно хорошо работать с деревьями компоновщика

Также есть вариант через директора. тогда задаётся жесткая форма задания, зато мы сразу получаемый готовый продукт.
При таком варианте для каждого отличного продукта необходимо прописать своего билдера
*/

/*
Плюсы:
1) Нет огромного конструктора с пустыми полями
2) Повторение кода для различных объектов
3) Изолирует код сборки от бизнес логики ????
*/

/*
минусы:
1) Лишний класс билдера и множество методов задания параметров
2) Если задаётся директор, то мы вообще ограничены конкретными билдерами.
*/

/*
пример на практике
1) подключение к http портам или БД, где задаётся много полей и тп.
2) возможно характеристика предмета в кэше, как в задании l0, если бы не было удобного парсера
*/

type Phone struct {
	name           string
	camera         int
	processor      string
	screenDiagonal int
	connector      string
	frontalCamera  int
	secondCamera   int
}

type Builder struct {
	name           string
	camera         int
	processor      string
	screenDiagonal int
	connector      string
	frontalCamera  int
	secondCamera   int
}

func NewBuilder(name, processor, connector string, camera, screenDiagonal int) *Builder {
	return &Builder{
		name:           name,
		processor:      processor,
		connector:      connector,
		camera:         camera,
		screenDiagonal: screenDiagonal,
	}
}

func (b *Builder) SetFrontalCamera(frontCam int) *Builder {
	b.frontalCamera = frontCam
	return b
}

func (b *Builder) SetSecondCamera(secCam int) *Builder {
	b.secondCamera = secCam
	return b
}

func newPhone(build *Builder) *Phone {
	return &Phone{
		name:           build.name,
		camera:         build.camera,
		processor:      build.processor,
		screenDiagonal: build.screenDiagonal,
		connector:      build.connector,
		frontalCamera:  build.frontalCamera,
		secondCamera:   build.secondCamera,
	}
}

func (b *Builder) Build() *Phone {
	return newPhone(b)
}

//usage
func main() {
	ph := NewBuilder("apple", "A13", "lightning", 12, 6).SetSecondCamera(20).Build()
	fmt.Println(ph)
}
