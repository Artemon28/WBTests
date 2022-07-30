package main

import "fmt"

//интерфейс вебсайта, не заблокированного ркн
type website interface {
	loadPage() string
}

//функция - браузер, которая загружает вебсайт
func browser(web website) {
	fmt.Println(web.loadPage())
}

//пример вебсайта
type wildberries struct {
}

func (WB wildberries) loadPage() string {
	return "WB page is loaded"
}

//заблокированный вебсайт, который браузер не может загрузить
type Linkedin struct {
}

func (LN Linkedin) loadBlockedPage() string {
	return "Linkedin page is loaded"
}

//структура адаптер, в которую подаётся заблокированный сайт
type VPNAdapter struct {
	blockWeb Linkedin
}

//но эта структура возвращает значение от заблокированного сайта и браузер знает как это обработать (кажется впн не так работает)
func (vpn VPNAdapter) loadPage() string {
	return vpn.blockWeb.loadBlockedPage()
}

func main() {
	wb := wildberries{}
	browser(wb)

	ln := Linkedin{}
	vpnLN := VPNAdapter{blockWeb: ln}

	browser(vpnLN)
}
