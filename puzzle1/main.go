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
	measurements := strings.Split(input, "\n")
	largerCount := 0
	for i, m := range measurements {
		if i == 0 || strings.TrimSpace(m) == "" {
			continue
		}
		if toInt(m) > toInt(measurements[i-1]) {
			largerCount++
		}
	}
	log.Print(largerCount)
}

func toInt(val string) int {
	m, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("cannot convert %s to int", val)
		os.Exit(1)
	}
	return m
}
