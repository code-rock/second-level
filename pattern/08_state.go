package pattern

import "fmt"

/*
	Паттерн «состояние».
	Помогает менять обьектам свое поведение в зависимости от состояния.
*/

/*
	Плюсы:
*/

/*
	Минусы:
*/

type SOrederStatus struct {
	name       string
	nextStatus INextStatus
}

func (os SOrederStatus) next() INextStatus {
	return os.nextStatus
}

type INextStatus interface {
	next() INextStatus
}

type SWaitingForPaymant struct {
	SOrederStatus
}

func (wp SWaitingForPaymant) cancelOrder() {
	fmt.Println("Надо бы переоткрыть заказ, но мне лень")
	// return SOrederStatus{
	// 	name:       "canceled",
	// 	nextStatus: Reopen,
	// }
}

type SShipping struct {
	SOrederStatus
}

type SDelivered struct {
	SOrederStatus
}

type SOrder struct {
	state INextStatus
}

func (os SOrder) nextState() INextStatus {
	return os.state.next()
}

func UseState() {
	// Delivered := SDelivered{
	// 	SOrederStatus{
	// 		name:       "delivered",
	// 		nextStatus: nil,
	// 	},
	// }

	// Shipping := SShipping{
	// 	SOrederStatus{
	// 		name:       "shipping",
	// 		nextStatus: Delivered,
	// 	},
	// }

	// WaitingForPaymant := SWaitingForPaymant{
	// 	SOrederStatus{
	// 		name:       "waitingForPaymant",
	// 		nextStatus: Shipping,
	// 	},
	// }

	// order := SOrder{
	// 	state: WaitingForPaymant,
	// }

	// fmt.Println(order.state.name)
	// order.state.cancelOrder()

	// order2 := SOrder{
	// 	state: WaitingForPaymant,
	// }

	// fmt.Println(order2.state.name)
	// order2.nextState()
	// fmt.Println(order2.state.name)
	// order2.nextState()
	// fmt.Println(order2.state.name)
}
