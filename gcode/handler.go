package gcode

type Handler struct {
	HandleComment func(string)
	HandleUnknown func(UnknownCommand)
	HandleMove    func(MoveCommand)
}
