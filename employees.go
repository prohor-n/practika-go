package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника по ID
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}

MainLoop:
	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d\n", &cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s\n", &empl.Name)
			fmt.Println("Возраст:")
			fmt.Scanf("%d\n", &empl.Age)
			fmt.Println("Позиция:")
			fmt.Scanf("%s\n", &empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scanf("%d\n", &empl.Salary)

			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					fmt.Printf("Сотрудник успешно добавлен! Присвоен ID: %d\n", i)
					added = true
					break
				}
			}
			if !added {
				fmt.Println("Ошибка: достигнут лимит в 512 сотрудников.")
			}

		case 2:
			fmt.Println("\nВведите ID сотрудника для удаления (от 0 до 511):")
			var idToRemove int
			fmt.Scanf("%d\n", &idToRemove)

			if idToRemove < 0 || idToRemove >= size {
				fmt.Printf("Ошибка: ID должен быть в диапазоне от 0 до %d.\n", size-1)
				continue
			}

			if empls[idToRemove] != nil {
				name := empls[idToRemove].Name
				empls[idToRemove] = nil
				fmt.Printf("Сотрудник %s (ID: %d) успешно удален.\n", name, idToRemove)
			} else {
				fmt.Printf("Сотрудник с ID %d не найден (ячейка уже пуста).\n", idToRemove)
			}

		case 3:
			fmt.Println("\n=== Список сотрудников ===")
			hasEmployees := false
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					hasEmployees = true
					fmt.Printf("ID: %d | Имя: %s | Возраст: %d | Позиция: %s | Зарплата: %d\n",
						i, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
				}
			}
			if !hasEmployees {
				fmt.Println("Список пуст.")
			}
			fmt.Println("==========================")

		case 4:
			fmt.Println("Выход из программы...")
			break MainLoop
		}
	}
}
