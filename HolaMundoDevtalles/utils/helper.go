// Package utils
package utils

import "errors"

func Booltolabel(value bool, iftrue string, iffalse string) string {

	if value {
		return iftrue

	}
	return iffalse

}

func Sum(number1 int, number2 int) (int, error) {
	if number1 <= 0 || number2 <= 0 {
		return 0, errors.New("no puedes mandar un numero menor o igual a 0")
	}

	total := number1 + number2

	return total, nil
}

func CalculateAdult(age int) (bool, error) {

	if age <= 17 {
		return false, errors.New("Edad no valida")
	}

	return true, nil
}

func CountText(texto string) (int, error) {
	textrunes := []rune(texto)

	if len(textrunes) <= 0 {
		return 0, errors.New("no se puede mandar un string vacio")
	}

	count := 0
	for _, v := range textrunes {
		if v == ' ' {
			continue
		}
		count += 1
	}
	return count, nil
}
