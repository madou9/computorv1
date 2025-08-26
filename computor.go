package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	// fmt.Println("Input equation: ", equation)

	terms := parseEquation(equation)
	reducedTerms, _ := reduceEquation(terms)
	getDegree(reducedTerms)

	fmt.Println("reduced terms : ", reducedTerms)
	// reduceTerms = reduceEquation(terms)
}

func parseEquation(eq string) map[int]float64 {

    // Split input into left and right parts
    parts := strings.SplitN(eq, "=", 2)
    left := strings.TrimSpace(parts[0])
    right := strings.TrimSpace(parts[1])

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
		coeff = strings.ReplaceAll(coeff, " ", "")
		exp := strings.TrimSpace(pieces[1])

		// removed the "X^" from exponent
		exp = strings.TrimPrefix(exp, "X^")

		// convert the string to a numeric value to be able to make the operation.
		convCoeff, _ := strconv.ParseFloat(coeff, 64)
		convExp, _ := strconv.Atoi(exp) 
		// fmt.Println("Raw term:", term, "| coeff:", coeff, "| exp:", exp)

		terms[convExp] += sign * convCoeff
	}

}

func reduceEquation(t map[int]float64)(map[int]float64, string){
	reduced := make(map[int]float64)
	for exp, coeff := range t {
		reduced[exp] = coeff
	}
    // Build string: sort exponents ascending
    exponents := make([]int, 0, len(reduced))
    for exp := range reduced {
        exponents = append(exponents, exp)
    }
	sort.Ints(exponents) // Sort exponents in ascending order 0 ,1 , 2

	var builder strings.Builder
	 // Build the reduced equation string
	for i, exp := range exponents {
		coeff := reduced[exp]
		sign := "+"
		if coeff < 0 {
			sign = "-"
			coeff = -coeff
	}
	if i > 0 {
		builder.WriteString(" " + sign + " ")
	}else if sign == "-" {
		builder.WriteString("-")
	}
	builder.WriteString(fmt.Sprintf("%g * X^%d", coeff, exp))
}
	builder.WriteString(" = 0")
	
	fmt.Println("Reduced form: ", builder.String())

	return reduced, builder.String()
}

func getDegree(t map[int]float64) int{

	if len(t) == 0 {
		return 0
	}
	degree := 0
	for exp, _ := range t {
		if exp > degree {
			degree = exp
		}
	}
	fmt.Printf("Polynomial degree: %d\n", degree)
	return degree
}

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

