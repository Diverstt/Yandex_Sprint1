package application

import (
	"fmt"

	"github.com/Diverstt/Yandex_Sprint1/pkg/rpn"
)

// CalcService отвечает за вычисление выражений.
type CalcService struct{}

// Calculate вычисляет выражение в обратной польской нотации.
func (c *CalcService) Calculate(expression string) (float64, error) {
	err := rpn.FindMistake(expression)
	if err != nil {
		return 0, fmt.Errorf("expression is not valid")
	}

	result, err := rpn.Calc(expression)
	if err != nil {
		return 0, fmt.Errorf("internal server error")
	}

	return result, nil
}
