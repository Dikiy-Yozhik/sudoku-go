package main

func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(column, digit):
		return ErrInColumn
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}

	g[row][column].digit = digit
	return nil
}

func (g *Grid) Clear(row, col int) error {
	switch {
	case !inBounds(row, col):
		return ErrBounds
	case g.isFixed(row, col):
		return ErrFixedDigit
	}

	g[row][col].digit = empty
	return nil
}

func (g *Grid) isComplete() bool {
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if g[r][c].digit == empty {
				return false
			}
		}
	}
	return true
}

// ----------------------

func inBounds(row, col int) bool {
	if row < 0 || row >= rows || col < 0 || col >= columns {
		return false
	}
	return true
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inColumn(col int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][col].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRegion(row, col int, digit int8) bool {
	r := row / 3 * 3
	for i := 0; i < 3; i++ {
		c := col / 3 * 3
		for j := 0; j < 3; j++ {
			if g[r][c].digit == digit {
				return true
			}
			c++
		}
		r++
	}

	return false
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) isFixed(row, col int) bool {
	return g[row][col].fixed
}
