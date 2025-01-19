package coordinates

import "math"

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

func New(x float64, y float64, z float64) *Coordinates {
	return &Coordinates{x, y, z}
}

func (c *Coordinates) Length() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z)
}

func (c *Coordinates) Length2() float64 {
	return c.X*c.X + c.Y*c.Y + c.Z*c.Z
}
