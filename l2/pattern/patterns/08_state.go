package patterns

import "fmt"

/*
	Реализовать паттерн «состояния».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

/*
Плюсы:
1) Убирает огромный иф с выбором состояния(машина состояния)
*/

/*
минусы:
1)
*/

/*
пример на практике
1)
*/

/*
в нашем примере есть онлайн магазин, оформляется заказ
1 Нет ничего в корзине (мы можем добавить товар в корзину, можем авторизоваться)
2 В корзине что-то есть (добавить товар в корзину, удалить товар из корзины, перейти в оплату товара, можем авторизоваться)
3 Оплата товара (только оплата товара)
*/

type Shop struct {
	noAuthNoItems   State
	noAuthWithItems State
	authNoItems     State
	authWithItems   State

	currentState State

	items int
}

func newShop() *Shop {
	v := &Shop{
		items: 0,
	}
	noAuthNoItemsState := &NoAuthNoItemsState{
		shop: v,
	}
	noAuthWithItemsState := &NoAuthWithItemsState{
		shop: v,
	}
	authNoItemsState := &AuthNoItemsState{
		shop: v,
	}
	authWithItemsState := &AuthWithItemsState{
		shop: v,
	}

	v.SetState(noAuthNoItemsState)
	v.noAuthNoItems = noAuthNoItemsState
	v.noAuthWithItems = noAuthWithItemsState
	v.authNoItems = authNoItemsState
	v.authWithItems = authWithItemsState
	return v
}

//функции интерфейса состояния
func (v *Shop) addToKart(id int) error {
	return v.currentState.addToKart(id)
}

func (v *Shop) deleteFromKart(id int) error {
	return v.currentState.deleteFromKart(id)
}

func (v *Shop) authorize() error {
	return v.currentState.authorize()
}

func (v *Shop) pay() error {
	return v.currentState.pay()
}

//функции работы с сайтом
func (s *Shop) SetState(st State) {
	s.currentState = st
}

func (s *Shop) AddItem(id int) {
	s.items++
}

func (s *Shop) DelItem(id int) {
	s.items--
}

func (s *Shop) ContainItem(id int) bool {
	if s.items > 0 {
		return true
	}
	return false
}

type State interface {
	addToKart(id int) error
	deleteFromKart(id int) error
	authorize() error
	pay() error
}

// клиент не авторизован и не выбрал товар
type NoAuthNoItemsState struct {
	shop *Shop
}

func (nani *NoAuthNoItemsState) addToKart(id int) error {
	nani.shop.AddItem(id)
	fmt.Println("Добавили товар в корзину")
	nani.shop.SetState(nani.shop.noAuthWithItems)
	return nil
}

func (nani *NoAuthNoItemsState) deleteFromKart(id int) error {
	return fmt.Errorf("Nothing in the kart")
}

func (nani *NoAuthNoItemsState) authorize() error {
	nani.shop.SetState(nani.shop.authNoItems)
	return nil
}

func (nani *NoAuthNoItemsState) pay() error {
	return fmt.Errorf("add items to kart and authorize")
}

//клиент не авторизован, но выбрал товар
type NoAuthWithItemsState struct {
	shop *Shop
}

func (nawi *NoAuthWithItemsState) addToKart(id int) error {
	nawi.shop.AddItem(id)
	fmt.Println("Добавили товар в корзину")
	return nil
}

func (nawi *NoAuthWithItemsState) deleteFromKart(id int) error {
	nawi.shop.DelItem(id)
	if !nawi.shop.ContainItem(id) {
		nawi.shop.SetState(nawi.shop.noAuthNoItems)
	}
	return nil
}

func (nawi *NoAuthWithItemsState) authorize() error {
	nawi.shop.SetState(nawi.shop.authWithItems)
	return nil
}

func (nawi *NoAuthWithItemsState) pay() error {
	return fmt.Errorf("authorize please")
}

//клиент не авторизован, но выбрал товар
type AuthNoItemsState struct {
	shop *Shop
}

func (ani *AuthNoItemsState) addToKart(id int) error {
	ani.shop.AddItem(id)
	fmt.Println("Добавили товар в корзину")
	ani.shop.SetState(ani.shop.authWithItems)
	return nil
}

func (ani *AuthNoItemsState) deleteFromKart(id int) error {
	return fmt.Errorf("no items in the kart")
}

func (ani *AuthNoItemsState) authorize() error {
	return fmt.Errorf("already authorized")
}

func (ani *AuthNoItemsState) pay() error {
	return fmt.Errorf("add items to kart")
}

// клиент авторизован и выбрал товары
type AuthWithItemsState struct {
	shop *Shop
}

func (awi *AuthWithItemsState) addToKart(id int) error {
	awi.shop.AddItem(id)
	fmt.Println("Добавили товар в корзину")
	return nil
}

func (awi *AuthWithItemsState) deleteFromKart(id int) error {
	awi.shop.DelItem(id)
	if !awi.shop.ContainItem(id) {
		awi.shop.SetState(awi.shop.authNoItems)
	}
	return nil
}

func (awi *AuthWithItemsState) authorize() error {
	return fmt.Errorf("already authorized")
}

func (awi *AuthWithItemsState) pay() error {
	fmt.Println("Происходит оплата")
	return nil
}
