package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (g *Grid) Show() {
	fmt.Println("\n   A B C   D E F   G H I")
	fmt.Println("  ╔═══════╦═══════╦═══════╗")
	for r := 0; r < rows; r++ {
		fmt.Printf("%d ║ ", r+1)
		for c := 0; c < columns; c++ {
			if g[r][c].digit == empty {
				fmt.Print(". ")
			} else {
				if g[r][c].fixed {
					fmt.Printf("\033[1m%d\033[0m ", g[r][c].digit)
				} else {
					fmt.Printf("%d ", g[r][c].digit)
				}
			}
			if c == 2 || c == 5 {
				fmt.Print("║ ")
			}
		}
		fmt.Println("║")
		if r == 2 || r == 5 {
			fmt.Println("  ╠═══════╬═══════╬═══════╣")
		}
		if r == 8 {
			fmt.Println("  ╚═══════╩═══════╩═══════╝")
		}
	}
}

func parseInput(input string) (int, int, int8, error) {
	input = strings.TrimSpace(input)
	if len(input) < 4 {
		return 0, 0, 0, errors.New("неправильный формат")
	}

	colStr := strings.ToUpper(string(input[0]))
	if colStr < "A" || colStr > "I" {
		return 0, 0, 0, errors.New("столбец должен быть от A до I")
	}
	col := int(colStr[0] - 'A')

	rowStr := string(input[1])
	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return 0, 0, 0, errors.New("строка должна быть от 1 до 9")
	}

	digit, errd := strconv.Atoi(string(input[len(input)-1]))
	if errd != nil || digit < 1 || digit > 9 {
		return 0, 0, 0, errors.New("цифра должна быть от 1 до 9")
	}

	return row - 1, col, int8(digit), nil
}
