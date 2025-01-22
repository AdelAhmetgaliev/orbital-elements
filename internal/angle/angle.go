package angle

import "math"

type Angle float64

func FromDegrees(d float64) Angle {
	r := d * (math.Pi / 180.0)

	return Angle(r)
}

func FromRadians(r float64) Angle {
	return Angle(r)
}

func Acos(c float64) Angle {
	return Angle(math.Acos(c))
}

func Asin(s float64) Angle {
	return Angle(math.Asin(s))
}

func Atan(t float64) Angle {
    return Angle(math.Atan(t))
}

func (a Angle) Degrees() float64 {
	return float64(a * (180.0 / math.Pi))
}

func (a Angle) Radians() float64 {
	return float64(a)
}

func (a Angle) Cos() float64 {
    return math.Cos(float64(a))
}

func (a Angle) Sin() float64 {
    return math.Sin(float64(a))
}

func (a Angle) Tan() float64 {
    return math.Tan(float64(a))
}
