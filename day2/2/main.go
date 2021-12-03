package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type submarine struct {
	Aim int
	X   int
	Y   int
}

//go:embed input
var input string

func main() {
	inputs := strings.Split(input, "\n")
	sub := submarine{0, 0, 0}

	for _, input := range inputs {
		sub.Move(input)
	}
	fmt.Printf("Final position Aim: %d, X: %d, Y: %d - multiple is %d", sub.Aim, sub.X, sub.Y, sub.X*sub.Y)
}

func (s *submarine) Move(command string) {
	if strings.TrimSpace(command) == "" {
		return
	}
	parts := strings.Split(command, " ")

	intVal, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("error converting %s: %v\n", parts[1], err)
	}

	switch parts[0] {
	case "forward":
		s.X += intVal
		s.Y += s.Aim * intVal
	case "down":
		s.Aim += intVal
	case "up":
		s.Aim -= intVal
	default:
		fmt.Println("No command for " + parts[0] + " - skipping...")
	}
}
