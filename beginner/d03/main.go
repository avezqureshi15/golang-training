package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("modulo by zero")
		}
		return float64(int(a) % int(b)), nil
	default:
		return 0, fmt.Errorf("Invalid Operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Calculator REPL")
		fmt.Println("Format: number operator number")
		fmt.Println("Type 'exit' to quit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Bye bro")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Use Format: number operator number")
			continue
		}

		//Parse numbers
		a, err1 := strconv.ParseFloat(parts[0], 64)
		op := parts[1]
		b, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid number")
			continue
		}

		result,err := calculate(a,op,b)
		if err != nil{
			fmt.Println("Error",err)
			continue
		}
		fmt.Println("Result: ",result)
	}
}
