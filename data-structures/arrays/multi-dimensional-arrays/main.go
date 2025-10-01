package main

import "fmt"

func main() {
	//var matrix [3][5]int
	/*var matrix = [3][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}*/
	matrix := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	//fmt.Println(matrix[1][1])

	// nested loops
	/*for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}*/

	// range loops
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d]=%d \n", i, j, val)
		}
	}

}
