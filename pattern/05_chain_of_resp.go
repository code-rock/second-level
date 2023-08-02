package pattern

import "fmt"

/*
	Паттерн «цепочка вызовов».

	позволяет передовать запросы последовательно по цепочке обработчиков
	каждый последующий обработчик решает задачу того может ли он сам обработать запрос
	либо его нужно передать дальше по цепочке
*/

// Структыр платежных систем

// Фигню какую-то наприсала
type SMaster struct {
	SAccount
}

type SPaypal struct {
	SAccount
}

type SQiwi struct {
	SAccount
}

type SAccount struct {
	name    string
	balance int
	incomer *SAccount
}

type IAccont interface {
	canPay(amount int) bool
	pay(orderPrice int)
	setNext(account IAccont)
}

func (a *SAccount) setNext(account *SAccount) {
	a.incomer = account
}

func (a *SAccount) canPay(amount int) bool {
	return a.balance >= amount
}

func (a *SAccount) pay(orderPrice int) {
	if a.canPay(orderPrice) {
		fmt.Printf("Paid %d using %s.\n", orderPrice, a.name)
	} else if a.incomer != nil {
		fmt.Printf("Can not pay using %s.\n", a.name)
		a.incomer.pay(orderPrice)
	} else {
		fmt.Printf("Unfortunatelly, not enough money")
	}
}

func (a *SAccount) show() {
	fmt.Println(a)
}

func UseChainOfResp() {
	master := SAccount{
		name:    "Master Card",
		balance: 8000,
	}

	paypal := SAccount{
		name:    "Paypal",
		balance: 630,
	}

	qiwi := SAccount{
		name:    "Qiwi",
		balance: 800630,
	}

	master.setNext(&paypal)
	paypal.setNext(&qiwi)

	master.pay(79000)
	master.show()
}
