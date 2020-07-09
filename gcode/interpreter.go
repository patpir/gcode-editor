package gcode

import (
	"fmt"
	"io"
	//"log"
	"reflect"

	"github.com/256dpi/gcode"
)

func gc(letter string, value int) gcode.GCode {
	return gcode.GCode{
		Letter: letter,
		Value: float64(value),
	}
}

type Interpreter struct {
	Handler      Handler
	HomePosition HeadPosition
	file         *gcode.File
	currentLine  int
	mover        Mover
}

func Interpret(reader io.Reader) (in *Interpreter, err error) {
	in = new(Interpreter)
	in.file, err = gcode.ParseFile(reader)
	if err != nil {
		return nil, err
	}
	in.mover = NewAbsoluteMover(
		Movement{},
		Millimeters,
	)
	return in, err
}

func (in *Interpreter) Walk() error {
	for err := in.Next(); err != io.EOF; err = in.Next() {
		if err != nil {
			return err
		}
	}
	return nil
}

func (in *Interpreter) nextLine() (gcode.Line, error) {
	if len(in.file.Lines) > in.currentLine {
		line := in.file.Lines[in.currentLine]
		in.currentLine += 1

		// prepare line for interpretation
		if len(line.Codes) == 0 {
			if line.Comment != "" {
				in.call(in.Handler.HandleComment, line.Comment)
			}
			return line, nil
		}

		params := lineParams(line)
		switch line.Codes[0] {
		case gc("G", 0), gc("G", 1):
			movement, err := in.mover.Move(params)
			if err != nil {
				return line, err
			}
			move := MoveCommand{
				Movement: movement,
			}
			in.call(in.Handler.HandleMove, move)
		case gc("G", 20):
			in.mover.SetUnits(Inches)
		case gc("G", 21):
			in.mover.SetUnits(Millimeters)
		case gc("G", 28):
			move := MoveCommand{
				Movement: in.mover.MoveTo(in.HomePosition),
			}
			in.call(in.Handler.HandleMove, move)
		case gc("G", 90):
			in.mover = NewAbsoluteMover(
				in.mover.LastMovement(),
				in.mover.Units(),
			)
		case gc("G", 91):
			in.mover = NewRelativeMover(
				in.mover.LastMovement(),
				in.mover.Units(),
			)
		case gc("G", 92):
			in.mover.SetReference(params)
		default:
			if line.Codes[0].Letter == "G" {
				return line, fmt.Errorf("unknown command: %v", line.Codes[0])
			}
/*
			cmd := UnknownCommand{
				Command: line.Codes[0]
				Params: line.Codes[1:],
			}
			in.call(in.Handler.HandleUnknown, cmd)
*/
		}

		return line, nil
	}

	return gcode.Line{}, io.EOF
}

func (in *Interpreter) Next() error {
	_, err := in.nextLine()
	return err
}

func (in *Interpreter) Position() Position {
	return in.mover.LastMovement().Position
}

func (in Interpreter) call(handleFunc interface{}, param interface{}) {
	if handleFunc != nil {
		params := []reflect.Value{ reflect.ValueOf(param) }
		handleFuncValue := reflect.ValueOf(handleFunc)
		if !handleFuncValue.IsNil() {
			reflect.ValueOf(handleFunc).Call(params)
		}
	}
}

/*
func (in Interpreter) handleComment(comment string) {
	if in.Handler.HandleComment != nil {
		in.Handler.HandleComment(comment)
	}
}

func (in Interpreter) handleUnknown(cmd UnknownCommand) {
	if in.Handler.HandleComment != nil {
		in.Handler.HandleUnknown(cmd)
	}
}

func (in Interpreter) handleMove(move MoveCommand) {
	if in.Handler.HandleMove != nil {
		in.Handler.HandleMove(move)
	}
}
*/
