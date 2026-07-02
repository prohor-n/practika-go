package main

import (
	"fmt"
	"strings"
)

// RomanToArabic конвертирует римские цифры в арабские
func RomanToArabic(roman string) (int, error) {
	// Карта соответствия римских цифр арабским
	romanMap := map[rune]int{
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

	// Идём справа налево для правильной обработки вычитающихся правил
	for i := len(roman) - 1; i >= 0; i-- {
		char := rune(roman[i])
		value, exists := romanMap[char]

		if !exists {
			return 0, fmt.Errorf("неверный символ в римской цифре: %c", char)
		}

		// Если текущее значение меньше предыдущего, вычитаем
		// (например, IV = 5 - 1 = 4)
		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result, nil
}

// ArabicToRoman конвертирует арабские цифры в римские (бонус!)
func ArabicToRoman(num int) (string, error) {
	if num <= 0 || num >= 4000 {
		return "", fmt.Errorf("число должно быть от 1 до 3999")
	}

	// Значения и их римские представления в убывающем порядке
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}

	return result, nil
}

func main() {
	fmt.Println("========== КОНВЕРТЕР РИМСКИХ ЦИФР ==========\n")

	// Тестовые примеры римских цифр
	testCases := []string{
		"I",      // 1
		"IV",     // 4
		"IX",     // 9
		"XXVII",  // 27
		"XLII",   // 42
		"XCIX",   // 99
		"CCCXL",  // 340
		"CDXLIV", // 444
		"MCMXC",  // 1990
		"MMXXIII", // 2023
	}

	fmt.Println("Римские цифры → Арабские:")
	for _, roman := range testCases {
		arabic, err := RomanToArabic(roman)
		if err != nil {
			fmt.Printf("  %s: ошибка - %v\n", roman, err)
		} else {
			fmt.Printf("  %s = %d\n", roman, arabic)
		}
	}

	fmt.Println("\nАрабские цифры → Римские (бонус):")
	arabicNumbers := []int{1, 4, 9, 27, 42, 99, 340, 444, 1990, 2023}
	for _, num := range arabicNumbers {
		roman, err := ArabicToRoman(num)
		if err != nil {
			fmt.Printf("  %d: ошибка - %v\n", num, err)
		} else {
			fmt.Printf("  %d = %s\n", num, roman)
		}
	}

	// Интерактивный пример
	fmt.Println("\n========== ПРИМЕРЫ ДЛЯ ТЕСТИРОВАНИЯ ==========")
	result, _ := RomanToArabic("XLVIII")
	fmt.Printf("XLVIII = %d\n", result)

	result, _ = RomanToArabic("MDCCCLXXXVIII")
	fmt.Printf("MDCCCLXXXVIII = %d\n", result)
}
