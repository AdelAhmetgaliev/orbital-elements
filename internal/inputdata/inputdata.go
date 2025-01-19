package inputdata

import (
	"bufio"
	"os"
	"strconv"

	"github.com/AdelAhmetgaliev/orbital-elements/internal/coordinates"
	"github.com/AdelAhmetgaliev/orbital-elements/internal/velocity"
)

func Read(filePath *string) (*coordinates.Coordinates, *velocity.Velocity) {
	inputFile, err := os.Open(*filePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err)
		}
	}()

	var inputData [6]float64

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < 6 && scanner.Scan(); i++ {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		inputData[i] = value
	}

	inputCoords := coordinates.New(inputData[0], inputData[1], inputData[2])
	inputVelocity := velocity.New(inputData[3], inputData[4], inputData[5])

	return inputCoords, inputVelocity
}
