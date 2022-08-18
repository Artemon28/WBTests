package patterns

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «Посетитель».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Есть мапа или слайс какого-то интерфейса. Мы хотим перебрать все элементы и выполнить над ними одинаковую операцию
Но мы не хотим прописывать эту операцию в их интерфейсе или для их структуры, так как это нарушит логику
или эта функция позже изменится. Поэтому делаем визиотра, даём ему доступ и запускаем

Важно что изначальная структура должна остаться неизменной, кроме разрешения доступа. Это делается чтобы не сломать
отлаженный код, а также не портить его логику. Применяется, когда изначальный код не должен часто изменяться, а вот
новые функции могут часто обновляться или изменяться.
*/

/*
Плюсы:
1) Не нарушается начальная логика и добавляется новый класс для отдельной общей логики
2) Упрощается добавление новых функций
3) Работа с каждой структурой отдельно, что даёт большую функциональность
*/

/*
минусы:
1) Не подходит если начальная архитектура будет часто меняться => добавочные методы придётся изменять каждый раз
2) нарушается инкапсуляция, добавляются функции вне наших структур
3) лишняя нагрузка на код
*/

/*
пример на практике
1) запись в json/xml
*/

type Record interface {
	WriteDB()
	accept(Visitor)
}

type Order struct {
	Order_uid    string    `json:"order_uid"`
	Track_number string    `json:"track_number"`
	Del_phone    string    `json:"Del_phone"`
	Customer_id  []string  `json:"customer_id"`
	Date_created time.Time `json:"date_created"`
	Delivery     Delivery
	Payment      Payment
}

func (o Order) WriteDB() {
	fmt.Println("write this struct to DB")
}

func (o *Order) accept(v Visitor) {
	v.visitForOrder(o)
}

type Delivery struct {
	Name    []string `json:"name"`
	Phone   []string `json:"phone"`
	Address string   `json:"address"`
}

func (d Delivery) WriteDB() {
	fmt.Println("write this struct to DB")
}
func (d *Delivery) accept(v Visitor) {
	v.visitForDelivery(d)
}

type Payment struct {
	Transaction   []string `json:"transaction"`
	Currency      string   `json:"currency"`
	Amount        int      `json:"amount"`
	Delivery_cost int      `json:"delivery_cost"`
	Goods_total   int      `json:"goods_total"`
}

func (p Payment) WriteDB() {
	fmt.Println("write this struct to DB")
}

func (p *Payment) accept(v Visitor) {
	v.visitForPayment(p)
}

//visitor
type Visitor interface {
	visitForOrder(*Order)
	visitForDelivery(delivery *Delivery)
	visitForPayment(payment *Payment)
}

type memoryCounter struct {
	memory int
}

func (m *memoryCounter) visitForOrder(o *Order) {
	fmt.Println("goes through all fields and calculate sum memory")
}

func (m *memoryCounter) visitForDelivery(d *Delivery) {
	fmt.Println("goes through all fields and calculate sum memory")
}

func (m *memoryCounter) visitForPayment(p *Payment) {
	fmt.Println("goes through all fields and calculate sum memory")
}

type writeToJSON struct {
}

func (w *writeToJSON) visitForOrder(o *Order) {
	fmt.Println("Write this struct to JSON")
}

func (w *writeToJSON) visitForDelivery(d *Delivery) {
	fmt.Println("Write this struct to JSON")
}

func (w *writeToJSON) visitForPayment(p *Payment) {
	fmt.Println("Write this struct to JSON")
}
