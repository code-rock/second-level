package pattern

import "fmt"

/*
	Паттерн «строитель».
	Отделяет конструирование сложного объекта от его представления,
	так что в результате одного и того же процесса конструирования
	могут получиться разные представления.
*/

/*
	Применим когда
	- алгоритм создания сложного объекта не должен зависить от того,
	из каких частей состоит объект и как они стыкуются между собой.
	- процесс конструирования должен обеспечивать различные представления
	конструируемого объекта.
*/

/*
	Плюсы:
	- Позволяет изменить внутреннее представлние продукта
	- Изолирует код, реализующий конструирование и представление
	- Представляет более точный контроль над процесом конструирования
*/

/*
	Минусы:
	- алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект и как они стыкуются между собой;
	- процесс конструирования должен обеспечивать различные представления конструируемого объекта.
*/

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type IComputerBuilder interface {
	CPU(val string) IComputerBuilder
	RAM(val int) IComputerBuilder
	MB(val string) IComputerBuilder

	Build() Computer
}

type SComputerBuilder struct {
	cpu string
	ram int
	mb  string
}

func (b SComputerBuilder) CPU(val string) IComputerBuilder {
	b.cpu = val
	return b
}

func (b SComputerBuilder) RAM(val int) IComputerBuilder {
	b.ram = val
	return b
}

func (b SComputerBuilder) MB(val string) IComputerBuilder {
	b.mb = val
	return b
}

func (b SComputerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func NewComputerBuilder() IComputerBuilder {
	return SComputerBuilder{}
}

func UseBuider() {
	compBuilder := NewComputerBuilder()
	computer := compBuilder.CPU("core i3").RAM(8).MB("gigabyte").Build()
	fmt.Println(computer)

	officeCompBuilder := NewOfficeComputerBuilder()
	officeCompBuilder.RAM(4)
	officeComputer := officeCompBuilder.Build()
	fmt.Println(officeComputer)
}

type SOfficeComputerBuilder struct {
	SComputerBuilder
}

func NewOfficeComputerBuilder() IComputerBuilder {
	return SOfficeComputerBuilder{}
}

func (b SOfficeComputerBuilder) Build() Computer {
	return Computer{
		CPU: "intel pentium 3",
		RAM: 2,
		MB:  "asrock",
	}
}
