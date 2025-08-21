package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Start → get input from CLI.

Parse → Reduce → Get Degree → Solve → Print.

End.
*/

func main() {
	// if (len(os.Args) < 2) {
	// 	fmt.Println("Please check your argument")
	// }
	equation := os.Args[1]
	fmt.Println("Input equation: ", equation)

	parseEquation(equation)
}

func parseEquation(eq string) {
	fmt.Println("parse equation: ", eq)
	parts := strings.SplitN(eq, "=", 2)
	splitInputLeft := strings.TrimSpace(parts[0])
	splitInputRight := strings.TrimSpace(parts[1])
	fmt.Println("split 1left equation: ", splitInputLeft)
	fmt.Println("split 1right equation: ", splitInputRight)

}

// func reduceEquation(t map[int]float64)float64{
// fmt.Println("reduce equation")
// // 	Input: the terms map.
// // Output: a reduced version with combined coefficients (e.g. if you had two X^2 terms, they’re added together).
// // What it does: ensures the polynomial is in its simplest reduced form.
// }

// func getDegree(t map[int]float64) int{
// 	fmt.Println("get equation")
// // 	Input: reduced terms.
// // Output: the highest exponent with a non-zero coefficient.
// // What it does: determines the polynomial degree.
// }

// func solveEquation(t map[int]float64) []string{
// fmt.Println("solve equation")

// // solveEquation(terms map[int]float64) []string
// // Input: reduced terms.
// // Output: a slice of solution strings.
// // What it does:
// // Checks the degree.
// // If degree 0 → infinite/no solution.
// // If degree 1 → solve linear (-b/a).
// // If degree 2 → call quadratic solver.
// // If degree > 2 → return "I can’t solve".
// }

func printResult(){
	fmt.Println("print Result")
}

