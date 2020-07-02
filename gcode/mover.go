package gcode

var _ Mover = new(AbsoluteMover)
var _ Mover = new(RelativeMover)

type Mover interface {
	SetUnits(u Units)
	Units() Units
	SetReference(coordinates Params)
	Move(coordinates Params) (Movement, error)
	MoveTo(pos HeadPosition) Movement
	LastMovement() Movement
}

type baseMover struct {
	units        Units
	lastMovement Movement
	zeroPos      Position
}

func (m *baseMover) SetUnits(u Units) {
	m.units = u
}

func (m *baseMover) Units() Units {
	return m.units
}

func (m *baseMover) SetReference(coordinates Params) {
	// TODO!
	//currentPos := 
}

func (m *baseMover) MoveTo(pos HeadPosition) Movement {
	m.lastMovement.Position.HeadPosition = pos
	return m.lastMovement
}

func (m *baseMover) LastMovement() Movement {
	return m.lastMovement
}

type AbsoluteMover struct {
	baseMover
}

func NewAbsoluteMover(lastMovement Movement, units Units) *AbsoluteMover {
	return &AbsoluteMover{
		baseMover: baseMover{
			lastMovement: lastMovement,
			units:        units,
		},
	}
}

func (m *AbsoluteMover) Move(coordinates Params) (Movement, error) {
	for key, value := range coordinates {
		switch key {
		case "E": m.lastMovement.E = m.units.ToSI(value)
		case "X": m.lastMovement.X = m.units.ToSI(value)
		case "Y": m.lastMovement.Y = m.units.ToSI(value)
		case "Z": m.lastMovement.Z = m.units.ToSI(value)
		case "F": m.lastMovement.FeedRate = m.units.ToSI(value)
		}
	}
	return m.lastMovement, nil
}

type RelativeMover struct {
	baseMover
}

func NewRelativeMover(lastMovement Movement, units Units) *RelativeMover {
	return &RelativeMover{
		baseMover: baseMover{
			lastMovement: lastMovement,
			units:        units,
		},
	}
}

func (m *RelativeMover) Move(coordinates Params) (Movement, error) {
	for key, value := range coordinates {
		switch key {
		case "E": m.lastMovement.E += m.units.ToSI(value)
		case "X": m.lastMovement.X += m.units.ToSI(value)
		case "Y": m.lastMovement.Y += m.units.ToSI(value)
		case "Z": m.lastMovement.Z += m.units.ToSI(value)
		case "F": m.lastMovement.FeedRate = m.units.ToSI(value)
		}
	}
	return m.lastMovement, nil
}
