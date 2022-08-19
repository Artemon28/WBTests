package patterns

import (
	"database/sql"
	"log"
)

/*
	Реализовать паттерн «Команда».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Плюсы:
1)
*/

/*
минусы:
1)
*/

/*
пример на практике
1) механизмы кнопок с повторяющимися действиями (так реализован swing на JAVA)
2) Сохранение истории/ логирование, отмена действия, создание сценариев действий
3) отправка команд по сети, приводится пример сдействием игрока в игре
4) многопоточное выполнение - задачи записываются как объекты и передаются на паралельные потоки (master/worker)
*/

// Отправитель
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// команда
type Command interface {
	execute()
}

type writeToDBCommand struct {
	id   int
	text string
	DB   DataBase
}

func NewWriteToDBCommand(id int, text string, DB DataBase) *writeToDBCommand {
	return &writeToDBCommand{
		id:   id,
		text: text,
		DB:   DB,
	}
}

func (wrDB *writeToDBCommand) execute() {
	wrDB.DB.write(wrDB.id, wrDB.text)
}

func (wrDB *writeToDBCommand) UnExecute() {
	wrDB.DB.delete(wrDB.id)
}

type UpdateDBCommand struct {
	id     int
	text   string
	buffer *sql.Rows
	DB     DataBase
}

func NewUpdateDBCommand(id int, text string, DB DataBase) *writeToDBCommand {
	return &writeToDBCommand{
		id:   id,
		text: text,
		DB:   DB,
	}
}

func (upDB *UpdateDBCommand) execute() {
	upDB.buffer = upDB.DB.read(upDB.id)
	upDB.DB.update(upDB.id, upDB.text)
}

func (upDB *UpdateDBCommand) UnExecute() {
	var bufferText string
	upDB.buffer.Scan(&bufferText)
	upDB.DB.update(upDB.id, bufferText)
}

type DataBase interface {
	write(int, string)
	read(int) *sql.Rows
	update(int, string)
	delete(int)
}

type OracleDB struct {
	DB *sql.DB
}

func (o *OracleDB) write(id int, text string) {
	o.DB.Exec(`INSERT Тут пишем запрос и отправляем`, id, text)
}

func (o *OracleDB) read(id int) *sql.Rows {
	buffer, _ := o.DB.Query(`SELECT достаём прошлую запись`, id)
	if buffer == nil {
		log.Println("there was nothing")
	}
	return buffer
}

func (o *OracleDB) update(id int, text string) {
	o.DB.Exec(`UPDATE Тут пишем запрос и отправляем`, id, text)
}

func (o *OracleDB) delete(id int) {
	o.DB.Exec(`DELETE Тут пишем запрос и отправляем`, id)
}
