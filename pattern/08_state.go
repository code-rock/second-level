package pattern

import "fmt"

/*
	Паттерн «состояние».
	Помогает менять объектам свое поведение в зависимости от внутреннего состояния.
	Извне создаетсявпечатление, что изменился классобъекта.
*/
/*
	Применим когда
	- поведение объекта зависит от его состояния и должно изменяться во время выполнения.
	- когда в коде операций встречаются состояния из многих ветвей условные операторы, в
	которыхвыбор ветви зависит от состояния.
*/
/*
	Плюсы:
	- локализация поведения, зависящего от состояния, и деление егона части, соотвествующие
	состояниям.
	- явно выраженые переходы между состояниями.
	- возможность совместного использования объектов состояния.
*/

/*
	Минусы:
	- выбор между увеличением количества классов и громоздкими условными операорами
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
	fmt.Println("Надо бы переоткрыть заказ...")
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
	Delivered := SDelivered{
		SOrederStatus{
			name:       "delivered",
			nextStatus: nil,
		},
	}

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
