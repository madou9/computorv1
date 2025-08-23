package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

/*
Start → get input from CLI.

Parse → Reduce → Get Degree → Solve → Print.

End.
*/

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Please check your argument")
	}
	equation := os.Args[1]
	fmt.Println("Input equation: ", equation)

	terms := parseEquation(equation)
	fmt.Println("terms : ", terms)
	// reduceTerms = reduceEquation(terms)
}

func parseEquation(eq string) map[int]float64 {
    fmt.Println("parse equation:", eq)

    // Split input into left and right parts
    parts := strings.SplitN(eq, "=", 2)
    left := strings.TrimSpace(parts[0])
    right := strings.TrimSpace(parts[1])

	// print the left and right
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)

	// store element in terms
    terms := make(map[int]float64)

    // 2. Process both sides
    processSide(left, 1, terms)   // left side: sign = +1
    processSide(right, -1, terms) // right side: sign = -1

    return terms
}

func processSide(side string, sign float64, terms map[int]float64) {
	// normalize the input so all terms can be split with a single delimiter ("+") instead of doing two separate splits for + and -
	normalized := strings.ReplaceAll(side, "-", "+-")
    parts := strings.Split(normalized, "+")
	// fmt.Println("parts: ", parts)

	for _, term := range parts {
        term = strings.TrimSpace(term)
        if term == "" {
            continue
        }
		// check the leng of each term 
		pieces := strings.Split(term, "*")
		if (len(pieces) != 2){
			log.Println("Invalid term: ", term)
			continue
		}

		coeff := strings.TrimSpace(pieces[0])
		exp := strings.TrimSpace(pieces[1])

		// removed the "X^" from exponent
		exp = strings.TrimPrefix(exp, "X^")

		// convert the string to a numeric value
		coeffN, _ := strconv.ParseFloat(coeff, 64)
		expInt, _ := strconv.Atoi(exp) 

		// fmt.Println("coeffN: ", coeffN)
		// fmt.Println("expInt: ", expInt)
		// fmt.Println("sign: ", sign)

		terms[expInt] += sign * coeffN
	}

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

