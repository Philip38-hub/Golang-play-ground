package golang

import "fmt"

func Rect(y int, x int) {
	for i := 0; i < x; i++ {
		fmt.Println()
		for j := 0; j < y; j++ {
			if i == 0 {
				fmt.Print("/")
			} else if i == x-1 {
				fmt.Print("\\")
			} else if  i == 0 || i == x-1 || j == 0 || j == y-1  {
				fmt.Print("*")	
			} else {
				fmt.Print(" ")
			}
			// if (i == 0 || j == 0)  && ( i == x-1 || j == y-1) {
			// 	fmt.Print("\\") // vertices
			// } else if (i == x-1 || j == 0) && ( i == 0 || j == y-1)  {
			// 	fmt.Print("/") // vertices
			// } else if i == 0 || i == x-1 || j == 0 || j == y-1 {
			// 	fmt.Print("*") //edges
			// } else {
			// 	fmt.Print(" ")
			// }
		}
	}
}
