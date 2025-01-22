package main

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/AdelAhmetgaliev/orbital-elements/internal/angle"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/calculations"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/constants"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/inputdata"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/vectorelement"
)

func main() {
	inputFilePath := filepath.FromSlash("data/input.txt")

	inputCoords, inputVelocity := inputdata.Read(&inputFilePath)

	distance := inputCoords.Length()
	velocitySquare := inputVelocity.Length2()

	reverseSemiMajorAxis := calculations.ReverseSemiMajorAxis(distance, velocitySquare)
	eccentricity := calculations.Eccentricity(inputCoords, inputVelocity, reverseSemiMajorAxis)
	eccentricAnomaly := calculations.EccentricAnomaly(inputCoords, inputVelocity, reverseSemiMajorAxis, eccentricity)
	averageAnomaly := calculations.AverageAnomaly(eccentricAnomaly, eccentricity)

	firstVectorElement := vectorelement.First(
		inputCoords, inputVelocity, reverseSemiMajorAxis, eccentricAnomaly)
	secondVectorElement := vectorelement.Second(
		inputCoords, inputVelocity, reverseSemiMajorAxis, eccentricAnomaly, eccentricity)

	if math.Abs(firstVectorElement.Length()-1.0) > constants.Epsilon {
		panic("Length of the first vector element is greater than 1.0")
	}

	if math.Abs(secondVectorElement.Length()-1.0) > constants.Epsilon {
		panic("Length of the second vector element is greater than 1.0")
	}

	if math.Abs(vectorelement.Dot(firstVectorElement, secondVectorElement)) > constants.Epsilon {
		panic("Vector elements are not orthogonal to each other")
	}

	eclipticTilt := angle.FromDegrees(constants.EclipticTiltDegrees)
	cose := eclipticTilt.Cos()
	sine := eclipticTilt.Sin()

	tempValue1 := firstVectorElement.Z*cose - firstVectorElement.Y*sine
	tempValue2 := secondVectorElement.Z*cose - secondVectorElement.Y*sine

	sinInclination := math.Sqrt(tempValue1*tempValue1 + tempValue2*tempValue2)

	sinArgOfPeriapsis := tempValue1 / sinInclination
	cosArgOfPeriapsis := tempValue2 / sinInclination

	argOfPeriapsis := angle.Atan2(sinArgOfPeriapsis, cosArgOfPeriapsis)

	sinAscendingNode :=
		(firstVectorElement.Y*cosArgOfPeriapsis - secondVectorElement.Y*sinArgOfPeriapsis) / cose
	cosAscendingNode := firstVectorElement.X*cosArgOfPeriapsis - secondVectorElement.X*sinArgOfPeriapsis

	ascendingNode := angle.Atan2(sinAscendingNode, cosAscendingNode)

	cosInclination := -(firstVectorElement.X*sinArgOfPeriapsis + secondVectorElement.X*cosArgOfPeriapsis) / sinAscendingNode

	inclination := angle.Atan2(sinInclination, cosInclination)

	argOfPeriapsisDegrees := argOfPeriapsis.Degrees()
	ascendingNodeDegrees := ascendingNode.Degrees()
	inclinationDegrees := inclination.Degrees()
	averageAnomalyDegrees := averageAnomaly.Degrees()
	semiMajorAxis := 1.0 / reverseSemiMajorAxis

	fmt.Printf("ω = %.8f°\n", argOfPeriapsisDegrees)
	fmt.Printf("Ω = %.8f°\n", ascendingNodeDegrees)
	fmt.Printf("i = %.8f°\n", inclinationDegrees)
	fmt.Printf("e = %.8f\n", eccentricity)
	fmt.Printf("a = %.8f a.e.\n", semiMajorAxis)
	fmt.Printf("M = %.8f°\n", averageAnomalyDegrees)
}
