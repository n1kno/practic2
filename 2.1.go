package main

import (
	"fmt"
	"strings"
)

// romanToArabic конвертирует римское число в арабское
func romanToArabic(roman string) int {
	// Карта соответствия римских символов и их значений
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	roman = strings.ToUpper(roman)
	result := 0
	prevValue := 0

	// Идем по строке справа налево
	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := values[rune(roman[i])]
		
		// Если текущее значение меньше предыдущего - вычитаем
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	return result
}

func main() {
	// Примеры использования
	examples := []string{"III", "IV", "IX", "LVIII", "MCMXC", "MMXXIV"}
	
	for _, roman := range examples {
		arabic := romanToArabic(roman)
		fmt.Printf("%s = %d\n", roman, arabic)
	}
}