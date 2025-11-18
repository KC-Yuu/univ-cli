package custom

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(expression string) (float64, error) {
	expression = strings.TrimSpace(expression)

	operators := []string{"+", "-", "*", "/"}
	var operator string
	var operatorIndex int

	for _, op := range operators {
		idx := strings.Index(expression, op)
		if idx > 0 { 
			operator = op
			operatorIndex = idx
			break
		}
	}

	if operator == "" {
		return 0, fmt.Errorf("expression invalide: aucun opérateur trouvé (+, -, *, /)")
	}

	leftStr := strings.TrimSpace(expression[:operatorIndex])
	rightStr := strings.TrimSpace(expression[operatorIndex+1:])

	left, err := strconv.ParseFloat(leftStr, 64)
	if err != nil {
		return 0, fmt.Errorf("nombre invalide: '%s'", leftStr)
	}

	right, err := strconv.ParseFloat(rightStr, 64)
	if err != nil {
		return 0, fmt.Errorf("nombre invalide: '%s'", rightStr)
	}

	var result float64

	switch operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division par zéro impossible")
		}
		result = left / right
	default:
		return 0, fmt.Errorf("opérateur inconnu: %s", operator)
	}

	return result, nil
}
