package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var isRoman bool

func main() {
	expression := `\b(V?I{1,3}[VX]?|[VX]|[0-9]|10)\b\s?([-+*/])\s?\b(V?I{1,3}[VX]?|[VX]|[0-9]|10)`

	fmt.Println("Введите математическое выражение:")
	var inputString string
	fmt.Scanln(&inputString)

	match := regexp.MustCompile(expression).FindStringSubmatch(inputString)
	operand1 := convertToNumber(match[1])
	operand2 := convertToNumber(match[3])
	_operator := match[2]
	if containsRomanNumerals(match[1]) != containsRomanNumerals(match[3]){
		fmt.Println("Ошибка! Выражение содержит числа разных форм. Я конечно могу их посчитать, но мне нельзя. :D")
		return
	}
	var result int
	switch _operator {
	case "/":
		result = operand1 / operand2
	case "*":
		result = operand1 * operand2
	case "-":
		result = operand1 - operand2
	case "+":
		result = operand1 + operand2
	}
	if !isRoman {
		fmt.Println("Результат выражения:", result)
	} else {
		if result < 0{
			fmt.Println("Ошибка! Результат выражения равен: ", result, ". Но в римской системе счисления нет отрицательных чисел.")
		} else {
			fmt.Println("Результат выражения:", arabicToRoman(result), "(", result, ")")
		}
	}
}

func convertToNumber(operand string) int {
	if containsRomanNumerals(operand) {
		isRoman = true
		romanNumerals := map[byte]int{
			'I': 1, 'V': 5, 'X': 10,
		}

		arabicNum := 0
		prevValue := 0

		for i := len(operand) - 1; i >= 0; i-- {
			value := romanNumerals[operand[i]]

			if value >= prevValue {
				arabicNum += value
			} else {
				arabicNum -= value
			}

			prevValue = value
		}
		return arabicNum
	} else {
		num, _ := strconv.Atoi(operand)
		return num
	}
}

func containsRomanNumerals(operand string) bool {
	for i := 0; i < len(operand); i++ {
		if operand[i] == 'I' || operand[i] == 'V' || operand[i] == 'X' {
			return true
		}
	}
	return false
}

func arabicToRoman(arabicNum int) string {
	romanNumerals := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C",
		90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX",
		5: "V", 4: "IV", 1: "I",
	}

	romanNumeral := ""

	for num, numeral := range romanNumerals {
		for arabicNum >= num {
			romanNumeral += numeral
			arabicNum -= num
		}
	}

	return romanNumeral
}