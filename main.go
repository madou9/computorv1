package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Please check your argument")
	}
	equation := os.Args[1]
	fmt.Println("Input equation: ", equation)

	parseEquation(equation)
}

func parseEquation(eq string) {
	fmt.Println("parse equation: ", eq)
}

func reduceEquation(){
	fmt.Println("reduce equation")
}

func getDegree(){
	fmt.Println("get equation")
}

func solveEquation(){
	fmt.Println("solve equation")
}

func printResult(){
	fmt.Println("print Result")
}

