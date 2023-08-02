package pattern

import (
	"errors"
	"fmt"
)

/*
	Паттерн «фабричный метод».
	Создания объекта, который будет помогать создавать обьекты.
	Когда нужно создавать одностипные объекты.
*/

/*
	Плюсы:
*/

/*
	Минусы:
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
