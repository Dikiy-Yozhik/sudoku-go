package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🎯 Добро пожаловать в Генератор Судоку!")
	fmt.Println("Выберите сложность:")
	fmt.Println("  1 - Легкий")
	fmt.Println("  2 - Средний")
	fmt.Println("  3 - Сложный")
	fmt.Print("Ваш выбор: ")

	scanner.Scan()
	difficultyInput := scanner.Text()

	difficulty := 2
	if diff, err := strconv.Atoi(difficultyInput); err == nil && diff >= 1 && diff <= 3 {
		difficulty = diff
	}

	fmt.Println("Генерируем случайное судоку...")
	game := NewRandomGame(difficulty)

	fmt.Println("\nКоманды:")
	fmt.Println("  A1 5 - поставить цифру 5 в клетку A1")
	fmt.Println("  clear A1 - очистить клетку A1")
	fmt.Println("  new - начать новую игру")
	fmt.Println("  quit - выйти из игры")
	fmt.Println("  help - показать помощь")

	for {
		game.Show()

		if game.isComplete() {
			fmt.Println("\n🎉 Поздравляем! Вы решили судоку!")
			return
		}

		fmt.Print("\nВведите команду: ")
		scanner.Scan()
		command := scanner.Text()

		switch {

		case command == "quit":
			fmt.Println("До свидания!")
			return

		case command == "help":
			fmt.Println("\nПомощь:")
			fmt.Println("Формат ввода: [столбец][строка] [цифра]")
			fmt.Println("  Столбцы: A-I (слева направо)")
			fmt.Println("  Строки: 1-9 (сверху вниз)")
			fmt.Println("  Пример: 'A1 5' - поставить 5 в верхний левый угол")
			fmt.Println("  'clear A1' - очистить клетку A1")
			fmt.Println("  'new' - начать новую игру")
			continue

		case command == "new":
			fmt.Println("Начинаем новую игру...")
			game = NewRandomGame(difficulty)
			continue

		case strings.HasPrefix(command, "clear "):
			input := strings.TrimPrefix(command, "clear ")
			if len(input) < 2 {
				fmt.Println("❌ Неправильный формат. Используйте: clear A1")
				continue
			}

			colStr := strings.ToUpper(string(input[0]))
			if colStr < "A" || colStr > "I" {
				fmt.Println("столбец должен быть от A до I")
				continue
			}
			col := int(colStr[0] - 'A')

			rowStr := string(input[1])
			row, err := strconv.Atoi(rowStr)
			if err != nil {
				fmt.Println("строка должна быть от 1 до 9")
				continue
			}
			err = game.Clear(row-1, col)
			if err != nil {
				fmt.Printf("❌ Ошибка: %v\n", err)
				continue
			} else {
				fmt.Println("✅ Клетка очищена")
			}

		default:
			row, column, digit, err := parseInput(command)
			if err != nil {
				fmt.Printf("❌ Ошибка ввода: %v\n", err)
				fmt.Println("Попробуйте снова или введите 'help' для помощи")
				continue
			}

			err = game.Set(row, column, digit)
			if err != nil {
				fmt.Printf("❌ Ошибка: %v\n", err)
			} else {
				fmt.Println("✅ Цифра установлена")
			}
		}
	}
}
