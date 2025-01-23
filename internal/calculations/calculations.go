package calculations

import (
	"math"

	"github.com/AdelAhmetgaliev/orbital-elements/internal/angle"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/constants"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/coordinates"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/velocity"
)

func ReverseSemiMajorAxis(distance float64, velocitySquare float64) float64 {
	return (2.0 / distance) - (velocitySquare / constants.GravitationalConstant2)
}

func Eccentricity(c *coordinates.Coordinates, v *velocity.Velocity, r float64) float64 {
	tempValue := c.X*v.X + c.Y*v.Y + c.Z*v.Z
	part1 := 1.0 - (c.Length() * r)
	part2 := tempValue * math.Sqrt(r) / constants.GravitationalConstant

	return math.Sqrt(part1*part1 + part2*part2)
}

func EccentricAnomaly(c *coordinates.Coordinates, v *velocity.Velocity, r float64, e float64) angle.Angle {
	tempValue := c.X*v.X + c.Y*v.Y + c.Z*v.Z
	part1 := 1.0 - (c.Length() * r)
	part2 := tempValue * math.Sqrt(r) / constants.GravitationalConstant

	cosE := part1 / e
	sinE := part2 / e

	eccentricAnomaly := angle.Atan2(sinE, cosE)

	return eccentricAnomaly
}

func AverageAnomaly(eccentricAnomaly angle.Angle, eccentricity float64) angle.Angle {
	return eccentricAnomaly - angle.Angle(eccentricity*eccentricAnomaly.Sin())
}
