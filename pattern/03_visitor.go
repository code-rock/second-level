package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Расширяет программу не изменяя содержимого структуры изнутри
*/

/*
	Плюсы:
*/

/*
	Минусы:
*/

type SCar struct{}

func (c *SCar) info() string {
	return "Базовый набор машины"
}

func (c *SCar) accept(visitor func(c *SCar)) {
	visitor(c)
}

type Tesla struct {
	SCar
}

func (c *Tesla) info() string {
	return "Электрокар"
}

type BMW struct {
	SCar
}

func (c *BMW) info() string {
	return "Красивая машинка"
}

type Audi struct {
	SCar
}

func (c *Audi) info() string {
	return "Хз что.."
}

type ICanShow interface {
	info() string
}

func exportVisitor(car *SCar) {
	if ok := ICanShow(car); ok != nil {
		fmt.Println(car.info())
	}
}

func UseVisitor() {
	q := Tesla{}
	q.accept(exportVisitor)
	w := BMW{}
	w.accept(exportVisitor)
	e := Audi{}
	e.accept(exportVisitor)
}
