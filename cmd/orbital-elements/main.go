package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("data/example.txt")
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

	fmt.Println(inputData)
}
