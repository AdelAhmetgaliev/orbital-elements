package velocity

import "math"

type Velocity struct {
	X float64
	Y float64
	Z float64
}

func New(x float64, y float64, z float64) *Velocity {
	return &Velocity{x, y, z}
}

func (v *Velocity) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Velocity) Length2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
