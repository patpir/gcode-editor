package gcode

type UnknownCommand struct {
	Command string
	Params []string
}

type MoveCommand struct {
	Movement
}
