package main

import "fmt"

func generate(numRows int) [][]int {
	solution := make([][]int, 0)

	a := make([]int, 0)
	for i := 0; i < numRows; i++ {
		fmt.Println("----------------")
		if i == 0 || i == 1 {
			fmt.Printf("i = %d\n", i)
			a = append(a, 1)
			fmt.Printf("a = %v\n", a)
			solution = append(solution, a)
			fmt.Printf("solution = %v\n", solution)
		} else {
			fmt.Printf("i = %d\n", i)
			b := make([]int, i+1)
			
			b[0] = 1
			fmt.Printf("b[0] = %d\n", b[0])
			
			for j := 1; j < i; j++ {
				fmt.Printf("i = %d\n", i)
				fmt.Printf("j = %d\n", j)
				value := solution[i-1][j-1] + solution[i-1][j] 
				b[j] = value
				fmt.Printf("b[%d] = %d\n", j, value)
			}
			b[i] = 1

			solution = append(solution, b)
		}
	}

	return solution
}

func main() {
	fmt.Println(generate(9))
}
