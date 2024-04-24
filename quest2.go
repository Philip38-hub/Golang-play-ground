package main

import "fmt"

func main() {
	PrintComb2()
	//PrintCombn()
}

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k {
					if i == '7' && j == '8' && k == '9' {
						fmt.Printf("%c", i)
						fmt.Printf("%c", j)
						fmt.Printf("%c", k)
					} else {
						fmt.Printf("%c", i)
						fmt.Printf("%c", j)
						fmt.Printf("%c", k)
						fmt.Printf(",")
						fmt.Printf(" ")
					}
				}
			}
		}
	}
}
func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i+1; j <= 99; j++ {
			fmt.Print(i/10)
			fmt.Print(i%10)
			fmt.Print(" ")
			fmt.Print(j/10)
			fmt.Print(j%10)
			if i != 98 && i != 99 {
				fmt.Print(",")
				fmt.Print(" ")
			}
		}
	}
}
func PrintCombn(n int) {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}