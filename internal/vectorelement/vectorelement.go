package vectorelement

import (
	"math"

	"github.com/AdelAhmetgaliev/orbital-elements/internal/coordinates"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/velocity"
)

type VectorElement struct {
	X float64
	Y float64
	Z float64
}

const gravitationalConstant = 0.01720209895

func First(c *coordinates.Coordinates, v *velocity.Velocity,
	reverseSemiMajorAxis float64, eccentricAnomaly float64) *VectorElement {
	r := c.Length()
	ak := math.Sqrt(1.0/reverseSemiMajorAxis) / gravitationalConstant
	cosE := math.Cos(eccentricAnomaly)
	sinE := math.Sin(eccentricAnomaly)

	x := (c.X/r)*cosE - (v.X*ak)*sinE
	y := (c.Y/r)*cosE - (v.Y*ak)*sinE
	z := (c.Z/r)*cosE - (v.Z*ak)*sinE

	return &VectorElement{x, y, z}
}

func Second(c *coordinates.Coordinates, v *velocity.Velocity,
	reverseSemiMajorAxis float64, eccentricAnomaly float64, e float64) *VectorElement {
	cosE := math.Cos(eccentricAnomaly)
	sinE := math.Sin(eccentricAnomaly)

	tempValue1 := c.Length() * math.Sqrt(1-e*e)
	tempValue2 := gravitationalConstant * math.Sqrt(1-e*e)
	tempValue3 := sinE / tempValue1
	tempValue4 := ((1.0 / math.Sqrt(reverseSemiMajorAxis)) * (cosE - e)) / tempValue2

	x := c.X*tempValue3 + v.X*tempValue4
	y := c.Y*tempValue3 + v.Y*tempValue4
	z := c.Z*tempValue3 + v.Z*tempValue4

	return &VectorElement{x, y, z}
}

func Dot(left *VectorElement, right *VectorElement) float64 {
	return left.X*right.X + left.Y*right.Y + left.Z*right.Z
}

func (ve *VectorElement) Length() float64 {
	return math.Sqrt(ve.X*ve.X + ve.Y*ve.Y + ve.Z*ve.Z)
}
