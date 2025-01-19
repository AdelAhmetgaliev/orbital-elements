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

func (v *Velocity) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Velocity) Length2() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}
