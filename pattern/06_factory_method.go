package pattern

import (
	"errors"
	"fmt"
)

/*
	Паттерн «фабричный метод».
	Определяет интерфейс для создания объекта, но оставляет подклассам решение о том,
	экземпляры какого класса должны создаваться. Фабричный метод позволяет делегировать
	создание экземпляров подклассам.
*/

/*
	Применим когда
	- классу зарание не известно объекты каких классов ему нужно созавать
	- класс спроектирован так, чтобы объекты, которые он создает, определялись подкалссами.
	- класс дилегирует свои обязанности одному из нескольких вспомогательных подклассов, и
	вам нужно локализовать информацию о том, какой класспринимает эти обязаности на себя.
*/

/*
	Плюсы:
	- избавляет проектировщика от необходимости встраивать в код зависящие от приложения классы.
	- создание объектов внутри класса с помощью фабричного метода более гибкое решение чем
	непосредственное создание. Создание в подклассах операций-зацепок для предоставления
	расширенной версии объекта.
	- возможно соединение паралельных иерархий.
*/

/*
	Минусы:
	- клиентам может потребоваться создавать подклассы класса для создания лишь одного объекта.
	Возможно придется иметь дело с дополнительным уровнем подклассов.
*/

type Motorbike struct {
	model    string
	price    float32
	maxSpeed int
}

type MotorbikeFactory struct{}

func (mf *MotorbikeFactory) create(model string) (Motorbike, error) {
	if model == "Круизер" {
		return Motorbike{
			model:    model,
			price:    170000,
			maxSpeed: 280,
		}, nil
	}

	if model == "Чоппер" {
		return Motorbike{
			model:    model,
			price:    180000,
			maxSpeed: 190,
		}, nil
	}

	if model == "Нейкед" {
		return Motorbike{
			model:    model,
			price:    190000,
			maxSpeed: 265,
		}, nil
	}

	return Motorbike{}, errors.New("Such model not exist")
}

func UseFactoryMethod() {
	factory := MotorbikeFactory{}
	k, errK := factory.create("Круизер")
	if errK != nil {
		fmt.Println(errK)
	} else {
		fmt.Println(k)
	}

	ch, errCh := factory.create("Чоппер")
	if errCh != nil {
		fmt.Println(errCh)
	} else {
		fmt.Println(ch)
	}

	n, errN := factory.create("Нейкед")
	if errN != nil {
		fmt.Println(errN)
	} else {
		fmt.Println(n)
	}
}
