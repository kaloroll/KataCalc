package main

import (
	"fmt"
	"regexp"
	"strconv"
	"bufio"
	"os"
)

var isRoman bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression := `^\b(V?I{1,3}[VX]?|[VX]I?|[0-9]|10)\b\s?([-+*/])\s?\b(V?I{1,3}[VX]?|[VX]I?|[0-9]|10)\s+$`
	checkExpression := regexp.MustCompile(expression)
	fmt.Println("Введите математическое выражение:")
	inputString, _ := reader.ReadString('\n')
	match := regexp.MustCompile(expression).FindStringSubmatch(inputString)
	if !checkExpression.MatchString(inputString){
		fmt.Println("Ошибка! Введеная строка имеет не верный формат.")
		return
	}
	operand1 := convertToNumber(match[1])
	operand2 := convertToNumber(match[3])
	_operator := match[2]
	if containsRomanNumerals(match[1]) != containsRomanNumerals(match[3]){
		fmt.Println("Ошибка! Выражение содержит числа разных форм. Я конечно могу их посчитать, но мне нельзя. :D")
		return
	}
	if operand1 < 1 || operand2 < 1 || operand1 > 10 || operand2 > 10{		
		fmt.Println("Ошибка! Программа может работать только с числами 1-10 или I-X")
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
		100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}

	romanNumeral := ""
	
	nums := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	
	for _, num := range nums {
		numeral := romanNumerals[num]
		for arabicNum >= num {
			romanNumeral += numeral
			arabicNum -= num
		}
	}
	return romanNumeral
}