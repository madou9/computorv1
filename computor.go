package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	Start -> get input from CLI.

	Parse -> Reduce -> Get Degree -> Solve -> Print.

	End.
*/

const EPSILON = 1e-8

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
	var left , right string
	if len(parts) == 1 {
		// no '=' found, treat entire input as left side
		left = strings.TrimSpace(parts[0])
		right = "0"
	} else if len(parts) == 2 {
		// normal case with '='
		left = strings.TrimSpace(parts[0])
		right = strings.TrimSpace(parts[1])
	} else {
		fmt.Println("Invalid equation format")
		os.Exit(1)
	}
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

		// remove internal spaces for simpler parsing
		s := strings.ReplaceAll(term, " ", "")
		if s == "" || s == "+" || s == "-" {
			continue
		}
		// work in upper case for 'x' vs 'X'
        u := strings.ToUpper(s)

		// contains X ?
		if strings.Contains(u, "X") {
			partsX := strings.SplitN(u, "X", 2)
			coeffPart := strings.TrimSuffix(partsX[0], "*")
			// fmt.Println("coeffPart: ", coeffPart)
			// fmt.Println("partsX: ", partsX)
			// get coefficient
			var coeff float64
			if coeffPart == "" || coeffPart == "+" {
				coeff = 1
			} else if coeffPart == "-" {
				coeff = -1
			} else {
				// parse coefficient
				c, _ := strconv.ParseFloat(coeffPart, 64)
				coeff = c
			}
			expPart := "1"
			if len(partsX) > 1 && strings.HasPrefix(partsX[1], "^"){
				expPart = partsX[1][1:] // remove '^'
				if expPart == "" {
					expPart = "1"
				}
			}
			exp, _ := strconv.Atoi(expPart)
			terms[exp] += sign * coeff
			continue
		}

		// no X: treat as constant term (exponent 0)
		constPart := strings.TrimSuffix(u, "*")
		constCoeff, _ := strconv.ParseFloat(constPart, 64)
		terms[0] += sign * constCoeff
		continue
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
	firstTerm := true
	 // Build the reduced equation string
	for _, exp := range exponents {
			coeff := reduced[exp]
			
			// Skip zero coefficients
			if Abs(coeff) < EPSILON {
				continue
			}
			
			sign := "+"
			if coeff < 0 {
				sign = "-"
				coeff = -coeff
			}
			
			if !firstTerm {
				builder.WriteString(" " + sign + " ")
			} else if sign == "-" {
				builder.WriteString("-")
			}
			
			builder.WriteString(fmt.Sprintf("%g * X^%d", coeff, exp))
			firstTerm = false
		}
    // Handle case where all coefficients are zero
    if firstTerm {
        builder.WriteString("0")
    }
	builder.WriteString(" = 0")
	
	fmt.Println("Reduced form:", builder.String())

	return reduced, builder.String()
}

func getDegree(t map[int]float64) int{

	if len(t) == 0 {
		return 0
	}
	degree := 0
	for exp, coeff := range t {
		// Only consider non-zero coefficients
		if Abs(coeff) > EPSILON && exp > degree {
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
			 // Check if constant term exists and is non-zero
        if val, exists := t[0]; exists && Abs(val) >= EPSILON {
            fmt.Println("There is no solution")
        } else {
            fmt.Println("Each real number is a solution")
        }
		case 1:
			 a := t[1]
        b := t[0]
        
        // Safety check for division by zero
        if Abs(a) < EPSILON {
            fmt.Println("Error: Invalid linear equation")
            return
        }
        
        x := -b / a
        if Abs(x) < EPSILON {
            x = 0
        }
        fmt.Printf("The solution is:\n%v\n", x)
		case 2:
			a := t[2]
			b := t[1]
			c := t[0]
			 // Safety check for division by zero
        if Abs(a) < EPSILON {
            fmt.Println("Error: Invalid quadratic equation")
            return
        }
			discriminant := b*b -4 * a * c
			// fmt.Printf("discriminant: %v\n : ", discriminant)
			if discriminant > 0 {
				sqrt := mySqrt(discriminant)
				x1 := (-b - sqrt) / (2 * a)
				x2 := (-b + sqrt) / (2 * a)
				fmt.Printf("Discriminant is strictly positive, the two solutions are: \n%f\n%f\n", x1, x2)	
			} else if Abs(discriminant) < EPSILON {
				x := -b / (2 * a)
				fmt.Println("Discriminant is zero, one real solution:")
				fmt.Printf("%f\n", x)
			}else {
			// COMPLEX solutions
				realPart := -b / (2 * a)
				if Abs(realPart) < EPSILON {
				realPart = 0  // Fix -0 display
			}
				imagPart := mySqrt(-discriminant) / (2 * a)
				fmt.Println("Discriminant is negative, the two complex solutions are:")
				fmt.Printf("%f - %f * i\n", realPart, imagPart)
				fmt.Printf("%f + %f * i\n", realPart, imagPart)
			}
		default:
			fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	}

}

func mySqrt(x float64) float64 {
		if x == 0 {
			return 0
		}

		guess := x / 2 // initial guess
		epsilon := EPSILON // precision threshold

		for {
			newguess := (guess + x/guess) / 2
			if Abs(newguess - guess) < epsilon {
				return newguess
			}
			guess = newguess // update gues

			// fmt.Println("newGuess: ", newguess)
		}
	}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
