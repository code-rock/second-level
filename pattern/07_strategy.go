package pattern

import "fmt"

/*
	Паттерн «стратегия».
	Определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми.
	Стратегия позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются.
*/
/*
	Применим когда
	- в наличие множество родственных классов, отличающихся только поведением.
	- в налииче несколько разновидностей алгоритмов.
	- в классе определено много вариантов поведения, представленных разветвленными
	условными операторами.
*/
/*
	Плюсы:
	- семейства родственных алгоритмов.
	- фльтернатива порождению подклассов.
	- позволяет избавиться от условных конструкций.
	- выбор реализаций.
*/

/*
	Минусы:
	- клиенты должны знать о разных стратегиях.
	- затраты на передачу информации между стратегией и контекстом.
	- увеличение читсла объектов.
*/

func baseStrategy(amount float32) float32 {
	return amount
}

func premiumStrategy(amount float32) float32 {
	return amount * 0.85
}

func platinumStrategy(amount float32) float32 {
	return amount * 0.65
}

type SAutoCart struct {
	discount func(float32) float32
	amount   float32
}

func (ac SAutoCart) checkout() float32 {
	return ac.discount(ac.amount)
}

func (ac *SAutoCart) setAmount(amount float32) {
	ac.amount = amount
}

func UseStrategy() {
	baseCustomer := SAutoCart{
		discount: baseStrategy,
		amount:   500000,
	}
	fmt.Println(baseCustomer.checkout())

	premiumCustomer := SAutoCart{
		discount: premiumStrategy,
		amount:   500000,
	}
	fmt.Println(premiumCustomer.checkout())

	platinumCustomer := SAutoCart{
		discount: platinumStrategy,
		amount:   500000,
	}
	fmt.Println(platinumCustomer.checkout())
}
