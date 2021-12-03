package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type binaryValue []binaryCount

func (b binaryValue) Gamma() uint {
	var result uint
	for i, c := range b {
		cResult := c.Calculate()
		fmt.Printf("cResult is %b\n", cResult)
		fmt.Printf("result is %b\n", result)
		result = result | cResult
		fmt.Printf("result & cResult: result is %b\n", result)
		if i < len(b)-1 {
			result = result << 1
		}

		fmt.Printf("result << 1: result is %b\n\n", result)
	}
	return result
}

type binaryCount struct {
	Zero int
	One  int
}

func (b *binaryCount) Zpp() {
	b.Zero++
}

func (b *binaryCount) Opp() {
	b.One++
}

// Calculate returns 1 if there are One counts, otherwise Zero.
func (b binaryCount) Calculate() uint {
	if b.One > b.Zero {
		return uint(1)
	}
	return uint(0)
}

func main() {
	inputs := strings.Split(input, "\n")
	result := make(binaryValue, 12, 12)
	for _, input := range inputs {
		if input == "" {
			continue
		}
		binaryDigits := strings.Split(input, "")
		for i, digit := range binaryDigits {
			if digit == "1" {
				result[i].Opp()
			} else if digit == "0" {
				result[i].Zpp()
			} else {
				fmt.Println("Unknown digit: " + digit)
			}
		}
	}

	gamma := result.Gamma()
	fmt.Printf("Gamma is %b = %d", gamma, gamma)

	// epsilon is the binary inverse of gamma
	epsilon := ^gamma
	epsilon = epsilon << 52
	epsilon = epsilon >> 52
	fmt.Printf("epsilon is %b = %d", epsilon, epsilon)
	fmt.Printf("result is: %d", epsilon*gamma)
}
