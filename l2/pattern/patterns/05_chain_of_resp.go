package patterns

import (
	"database/sql"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «Цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Плюсы:
1) Убирает нагруженные вызовы If else, которые могут менять логику запуска
2) Динамическое создание проверок во время написания проекта
3) Уменьшение зависимости клиента от обработчика (обработчик вынесен как раз в эту цепочку)
*/

/*
минусы:
1) Запрос может быть не обработан
2) Большой код, вместо удобных Ифов
*/

/*
пример на практике
1) Мой пример кода с вызовом кэша, также в нём может быть несколько кэшей или баз данных
2) Авторизация пользователя мб
*/

//вызов данных из БД, мы ускоряем запрос, поэтому сначала смотрим в КЭШ, потом достаём просто строковое значение из
//БД, потом скачиваем в формате JSON, потом уже полная обработка БД

type Request struct {
	id int
}

type Reader interface {
	Execute(*Request)
	SetNext(Reader)
}

// Смотрим есть ли значение в кэше
type ReaderFromCache struct {
	cache map[int]int
	next  Reader
}

func (r *ReaderFromCache) Execute(req *Request) {
	if v, ok := r.cache[req.id]; ok {
		fmt.Println(v)
	} else {
		r.next.Execute(req)
	}
}

func (r *ReaderFromCache) SetNext(next Reader) {
	r.next = next
}

//читаем из БД значение
type ReaderFromDB struct {
	DB   sql.DB
	next Reader
}

func (r *ReaderFromDB) Execute(req *Request) {
	//Ищем в БД значение если есть, то печатаем
	if true {
		fmt.Println("Я нашёл твое значение")
	}
	//Если это значение удалено, то можем выйти с ошибкой
	if true {
		log.Println("Это значение было удалено вчера :-(")
	}
	r.next.Execute(req)
}

func (r *ReaderFromDB) SetNext(next Reader) {
	r.next = next
}

func main() {
	rCache := &ReaderFromCache{}
	rBD := &ReaderFromDB{}
	rCache.SetNext(rBD)
	req := &Request{id: 1234}
	rCache.Execute(req)
}
