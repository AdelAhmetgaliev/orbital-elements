package velocity

import "math"

type Velocity struct {
	x float64
	y float64
	z float64
}

func New(x float64, y float64, z float64) *Velocity {
	return &Velocity{x, y, z}
}

func (c *Velocity) Length() float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y + c.z*c.z)
}

func (c *Velocity) Length2() float64 {
	return c.x*c.x + c.y*c.y + c.z*c.z
}
