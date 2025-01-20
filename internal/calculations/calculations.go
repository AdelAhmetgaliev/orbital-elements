package calculations

import (
	"math"

	"github.com/AdelAhmetgaliev/orbital-elements/internal/constants"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/coordinates"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/velocity"
)

func ReverseSemiMajorAxis(distance float64, velocitySquare float64) float64 {
	return (2.0 / distance) - (velocitySquare / constants.GravitationalConstant2)
}

func Eccentricity(c *coordinates.Coordinates, v *velocity.Velocity, reverseSemiMajorAxis float64) float64 {
	tempValue := c.X*v.X + c.Y*v.Y + c.Z*v.Z
	part1 := 1.0 - (c.Length() * reverseSemiMajorAxis)
	part2 := tempValue * math.Sqrt(reverseSemiMajorAxis) / constants.GravitationalConstant

	return math.Sqrt(part1*part1 + part2*part2)
}

func EccentricAnomaly(c *coordinates.Coordinates, v *velocity.Velocity, reverseSemiMajorAxis float64, e float64) float64 {
	tempValue := c.X*v.X + c.Y*v.Y + c.Z*v.Z
	part1 := 1.0 - (c.Length() * reverseSemiMajorAxis)
	part2 := tempValue * math.Sqrt(reverseSemiMajorAxis) / constants.GravitationalConstant

	cosE := part1 / e
	sinE := part2 / e

	eccentricAnomaly := math.Atan2(sinE, cosE)
	if eccentricAnomaly < 0.0 {
		eccentricAnomaly += 2.0 * math.Pi
	}

	return eccentricAnomaly
}

func AverageAnomaly(eccentricAnomaly float64, eccentricity float64) float64 {
	return eccentricAnomaly - eccentricity*math.Sin(eccentricAnomaly)
}
