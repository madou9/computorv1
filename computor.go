package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
)

/*
	Start -> get input from CLI.

	Parse -> Reduce -> Get Degree -> Solve -> Print.

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
	degree := getDegree(reducedTerms)
	solveEquation(reducedTerms, degree)

	// fmt.Println("reduced terms : ", reducedTerms)
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

func solveEquation(t map[int]float64, degree int){
	// fmt.Printf("terms : %v\n degree : %v\n", t, degree)
	switch degree {
		case 0:
			if math.Abs(t[0]) < 1e-8 {
				fmt.Println("All real numbers are solutions")
			} else {
				fmt.Println("No solution")
			}
		case 1:
			a := t[1]
			b := t[0]
			x := -b / a
			fmt.Printf("The solution is: \n%v\n", x)
		case 2:
			a := t[2]
			b := t[1]
			c := t[0]
			discriminant := b*b -4 * a * c
			// fmt.Printf("discriminant: %v\n : ", discriminant)
			if discriminant > 0 {
				sqrt := mySqrt(discriminant)
				x1 := (-b - sqrt) / (2 * a)
				x2 := (-b + sqrt) / (2 * a)
				fmt.Printf("Discriminant is strictly positive, the two solutions are: \n%f\n%f\n", x1, x2)	
			} else if math.Abs(discriminant) < 1e-8 {
				x := -b / (2 * a)
				fmt.Println("Discriminant is zero, one real solution:")
				fmt.Printf("%f\n", x)
			}
		default:
			fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	}

}

func mySqrt(x float64) float64 {
		if x == 0 {
			return 0
		}

		guess := x / 2
		epsilon := 1e-7

		for {
			newguess := (guess + x/guess) / 2
			if math.Abs(newguess - guess) < epsilon {
				return newguess
			}
			guess = newguess

			// fmt.Println("newGuess: ", newguess)
		}
	}
