package main

import (
	"fmt"
	"math/rand"

	// "sort"
	"time"
)

// FillMatrix2DUnique заполняет двумерный массив случайными уникальными числами
func FillMatrix2DUnique(rows, cols int) [][]int {
	// Проверка на возможность заполнения
	totalCells := rows * cols
	if totalCells > 10000 {
		fmt.Println("Предупреждение: много элементов, процесс может быть медленным")
	}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Создаём слайс уникальных чисел от 1 до rows*cols
	numbers := make([]int, totalCells)
	for i := 0; i < totalCells; i++ {
		numbers[i] = i + 1
	}

	// Перемешиваем числа (Fisher-Yates shuffle)
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// Заполняем матрицу
	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[idx]
			idx++
		}
	}

	return matrix
}

// FillMatrix2DUniqueRange заполняет матрицу случайными уникальными числами из диапазона
func FillMatrix2DUniqueRange(rows, cols int, minVal, maxVal int) ([][]int, error) {
	totalCells := rows * cols
	rangeSize := maxVal - minVal + 1

	if totalCells > rangeSize {
		return nil, fmt.Errorf("невозможно заполнить %d ячеек уникальными числами из диапазона [%d, %d]",
			totalCells, minVal, maxVal)
	}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Создаём слайс чисел в диапазоне
	numbers := make([]int, rangeSize)
	for i := 0; i < rangeSize; i++ {
		numbers[i] = minVal + i
	}

	// Перемешиваем и берём столько, сколько нужно
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// Заполняем матрицу первыми totalCells элементами
	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[idx]
			idx++
		}
	}

	return matrix, nil
}

// PrintMatrix красиво выводит матрицу
func PrintMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	// Вычисляем ширину столбца
	maxWidth := 1
	for _, row := range matrix {
		for _, val := range row {
			width := len(fmt.Sprintf("%d", val))
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	// Выводим матрицу
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%*d ", maxWidth, val)
		}
		fmt.Println()
	}
}

// CheckUniqueness проверяет, что все числа в матрице уникальны
func CheckUniqueness(matrix [][]int) bool {
	seen := make(map[int]bool)
	for _, row := range matrix {
		for _, val := range row {
			if seen[val] {
				return false // Нашли дубликат
			}
			seen[val] = true
		}
	}
	return true
}

// GetStatistics выводит статистику по матрице
func GetStatistics(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	var minVal, maxVal int
	var sum int
	count := 0

	for _, row := range matrix {
		for _, val := range row {
			if count == 0 {
				minVal = val
				maxVal = val
			}
			if val < minVal {
				minVal = val
			}
			if val > maxVal {
				maxVal = val
			}
			sum += val
			count++
		}
	}

	average := float64(sum) / float64(count)

	fmt.Printf("\nСтатистика:\n")
	fmt.Printf("  Размер: %dx%d = %d элементов\n", len(matrix), len(matrix[0]), count)
	fmt.Printf("  Минимум: %d\n", minVal)
	fmt.Printf("  Максимум: %d\n", maxVal)
	fmt.Printf("  Сумма: %d\n", sum)
	fmt.Printf("  Среднее: %.2f\n", average)
	fmt.Printf("  Уникальные: %v\n", CheckUniqueness(matrix))
}

func main() {
	// Инициализируем random seed
	rand.Seed(time.Now().UnixNano())

	fmt.Println("========== ЗАПОЛНЕНИЕ 2D МАССИВА СЛУЧАЙНЫМИ УНИКАЛЬНЫМИ ЧИСЛАМИ ==========\n")

	// Пример 1: Матрица 4x4
	fmt.Println("Пример 1: Матрица 4x4 (числа от 1 до 16)")
	matrix1 := FillMatrix2DUnique(4, 4)
	PrintMatrix(matrix1)
	GetStatistics(matrix1)

	// Пример 2: Матрица 5x3
	fmt.Println("\n\nПример 2: Матрица 5x3 (числа от 1 до 15)")
	matrix2 := FillMatrix2DUnique(5, 3)
	PrintMatrix(matrix2)
	GetStatistics(matrix2)

	// Пример 3: Матрица с диапазоном
	fmt.Println("\n\nПример 3: Матрица 3x5 (случайные числа от 100 до 199)")
	matrix3, err := FillMatrix2DUniqueRange(3, 5, 100, 199)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		PrintMatrix(matrix3)
		GetStatistics(matrix3)
	}

	// Пример 4: Большая матрица
	fmt.Println("\n\nПример 4: Матрица 10x10 (числа от 1 до 100)")
	matrix4 := FillMatrix2DUnique(10, 10)
	PrintMatrix(matrix4)
	GetStatistics(matrix4)

	// Пример 5: Проверка на ошибку
	fmt.Println("\n\nПример 5: Попытка заполнить 3x3 числами от 1 до 5 (невозможно)")
	_, err = FillMatrix2DUniqueRange(3, 3, 1, 5)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
