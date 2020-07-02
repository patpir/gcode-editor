package gcode

import (
	"fmt"
)

var (
	Millimeters = Units{1, "mm"}
	Inches = Units{25.4, "in"}
)

type Units struct {
	conversionFactor float64
	suffix string
}

func (u Units) ToSI(value float64) float64 {
	if u.conversionFactor < 0.00001 {
		panic(fmt.Errorf("Invalid conversion factor %f for unit %q",
			u.conversionFactor,
			u.suffix,
		))
	}
	return value * u.conversionFactor
}

func (u Units) String(value float64) string {
	return fmt.Sprintf("%f %s", value, u.suffix)
}
