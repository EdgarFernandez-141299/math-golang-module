package operations

import "errors"

func Sumar(a, b float64) float64 {
	return a + b
}

func Restar(a, b float64) float64 {
	return a - b
}

func Multiplicar(a, b float64) float64 {
	return a * b
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}
