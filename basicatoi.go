package piscine

import (
	"strconv"
)

func BasicAtoi(s string) int { //convert string to int
	x, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return x
}

func BasicItoa(s int) string{ // converts int to string
	x := strconv.Itoa(s)
	return x
}

func BasicFormartFloat(s float64) string{ //converts float to string
	x := strconv.FormatFloat(s, 'f', 2, 64)
	return x
}