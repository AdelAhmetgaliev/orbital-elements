package coordinates

import "math"

type Coordinates struct {
	x float64
	y float64
	z float64
}

func New(x float64, y float64, z float64) *Coordinates {
	return &Coordinates{x, y, z}
}

func (c *Coordinates) Length() float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y + c.z*c.z)
}

func (c *Coordinates) Length2() float64 {
	return c.x*c.x + c.y*c.y + c.z*c.z
}
