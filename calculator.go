package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для преобразования римского числа в арабское
func romanToInt(s string) int {
	roman := map[byte]int{'I': 1, 'V': 5, 'X': 10}
	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := roman[s[i]]
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

// Функция для преобразования арабского числа в римское
func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i, v := range val {
		for num >= v {
			result += sym[i]
			num -= v
		}
	}
	return result
}

// Проверка, является ли строка римским числом
func isRoman(s string) bool {
	for _, ch := range s {
		if ch != 'I' && ch != 'V' && ch != 'X' {
			return false
		}
	}
	return true
}

// Выполнение арифметической операции
func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неверная арифметическая операция")
	}
}

func main() {
	var input string
	fmt.Println("Введите арифметическое выражение (например, 3 + 2 или III + IV):")
	fmt.Scanln(&input)

	// Разделяем ввод на части
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	isRomanA, isRomanB := isRoman(aStr), isRoman(bStr)
	if isRomanA != isRomanB {
		panic("Нельзя смешивать арабские и римские числа")
	}

	var a, b int
	if isRomanA {
		a = romanToInt(aStr)
		b = romanToInt(bStr)
	} else {
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil || a < 1 || a > 10 {
			panic("Неверное арабское число")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil || b < 1 || b > 10 {
			panic("Неверное арабское число")
		}
	}

	// Выполняем операцию
	result := calculate(a, b, op)

	// Выводим результат
	if isRomanA {
		if result < 1 {
			panic("Результат римского числа не может быть менее 1")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}