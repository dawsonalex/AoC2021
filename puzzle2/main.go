package main

import (
	_ "embed"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	// TODO: THIS DOESN'T WORK
	measurements := strings.Split(input, "\n")
	var lastSum, greaterSums int
	for i := 0; i < len(measurements); i += 2 {
		if i+3 >= len(measurements) {
			break
		}
		if i == 0 {
			// set the first group
			lastSum = sum(toIntSlice(measurements[i : i+3]))
			continue
		}
		thisSum := sum(toIntSlice(measurements[i : i+3]))
		if thisSum > lastSum {
			log.Printf("this sum %d > greater than last %d", thisSum, lastSum)
			greaterSums++
		}
		lastSum = thisSum
	}
	log.Print(greaterSums)
}

func toIntSlice(vals []string) []int {
	result := make([]int, len(vals), len(vals))
	for i, val := range vals {
		m, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("cannot convert %s to int", val)
			os.Exit(1)
		}
		result[i] = m
	}
	return result
}

func sum(values []int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}
