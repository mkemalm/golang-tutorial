package main

import (
	"fmt"
	"os"
	"strconv"
)

func mathfunc(num1 int, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return num1 + num2
	}
	return num1 + num2
}

func main() {
	//fmt.Scanf("%d %d", &a, &b)
	if len(os.Args) > 1 { /* os.Args[0] is "hello" or "hello.exe" */
		a, err1 := strconv.Atoi(os.Args[1])
		b, err2 := strconv.Atoi(os.Args[2])
		if err1 == nil && err2 == nil {
			fmt.Println("Add: " + strconv.Itoa(mathfunc(a, b, "+")))
			fmt.Println("Subs: " + strconv.Itoa(mathfunc(a, b, "-")))
			fmt.Println("Mult: " + strconv.Itoa(mathfunc(a, b, "*")))
			fmt.Println("Div: " + strconv.Itoa(mathfunc(a, b, "/")))
			fmt.Println("Default op is addition: " + strconv.Itoa(mathfunc(a, b, "?")))
		}
	}
}
