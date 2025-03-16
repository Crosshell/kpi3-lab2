package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func EvaluatePrefixExpression(expression string) (float64, error) {
	if expression == "" {
		return 0, errors.New("empty expression")
	}

	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, errors.New("invalid expression")
	}

	stack := []float64{}
	for i := len(tokens) - 1; i >= 0; i-- {
		tok := tokens[i]

		switch tok {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, errors.New("syntax error: not enough operands")
			}
			op1, op2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res float64
			switch tok {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 + op2
			case "*":
				res = op1 + op2
			case "/":
				if op2 == 0 {
					return 0, errors.New("division by zero")
				}
				res = op1 / op2
			case "^":
				res = math.Pow(op1, op2)
			}
			stack = append(stack, res)

		default:
			num, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return 0, errors.New("invalid token: " + tok)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("syntax error: invalid expression")
	}

	return stack[0], nil
}
