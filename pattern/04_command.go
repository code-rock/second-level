package pattern

import "fmt"

/*
	Паттерн «комманда».
	Инкапсулирует запрос в объекте, позволяя тем самым параметризовать клиенты для разных
	запросов, ставить запросы в очередь или протоколировать их, а также поддерживать отмену
	операций.
*/
/*
	Применим когда нужны:
	- параметризация объектов выполняемым действием.
	- определение, постоновка в очередь  и выполнение в разное время.
	- поддержка отмены операций
	- поддержка протоколирования изменений, чтобы их можно было выполнить повторно после сбоя
	в сиситеме.
	- структурирование системы на основе высокоуровневых операций, построенных из приметивных.
*/

/*
	Плюсы:
	- отделяет объект, инициализирующий операцию, от объекта, располагающего информацией о том,
	как ее выполнить.
	- команда это обект который можно расширять.
	- из простых команд можно собирать состовные.
	- новые команды добавляются легко, поскольку никакие существующие классы изменять не нужно.
*/

/*
	Минусы:
*/

type ICommand interface {
	execute()
}

type SDriver struct {
	command ICommand
}

func (d SDriver) execute() {
	d.command.execute()
}

type SEngin struct {
	state bool
}

func (e *SEngin) on() {
	e.state = true
}

func (e *SEngin) off() {
	e.state = false
}

type sOnStartCommand struct {
	engine SEngin
}

func (osc sOnStartCommand) execute() {
	osc.engine.on()
}

type sOnSwitchOffCommand struct {
	engine SEngin
}

func (osoc sOnSwitchOffCommand) execute() {
	osoc.engine.off()
}

func UseCommand() {
	engine := SEngin{state: false}
	fmt.Println(engine)

	onStartCommand := sOnStartCommand{
		engine: engine,
	}
	driver := SDriver{
		command: onStartCommand,
	}
	driver.execute()
	fmt.Println(engine)
}
