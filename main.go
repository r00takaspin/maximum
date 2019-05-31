package main

import (
	"fmt"
	"github.com/r00takaspin/maximum/maximum"
	"os"
	"strconv"
)

func parseInput(args []string) ([]float64, error) {
	res := make([]float64, len(args))

	for i := 0; i < len(args); i++ {
		f, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			return []float64{}, err
		}
		res[i] = f
	}

	return res, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("please pass numbers: ./max 1 2 59 32.3 2")
		os.Exit(0)
	}

	numbers, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Printf("error while parsing argument list :%v", err)
		os.Exit(0)
	}

	max1, max2, err := maximum.MaxPair(numbers)
	if err != nil {
		fmt.Printf("error while finding maximum elements: %v", err)
		os.Exit(0)
	}

	fmt.Printf("%f %f", max1, max2)
}
