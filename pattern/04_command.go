package pattern

import "fmt"

/*
	Паттерн «комманда».
*/

/*
	Плюсы:
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
