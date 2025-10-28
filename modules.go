package main

import (
	"errors"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

var (
	ErrBounds     = errors.New("за пределами")
	ErrDigit      = errors.New("неправильная цифра")
	ErrInRow      = errors.New("эта цифра уже есть в ряду")
	ErrInColumn   = errors.New("эта цифра уже есть в столбце")
	ErrInRegion   = errors.New("в данной части эта цифра уже есть")
	ErrFixedDigit = errors.New("начальные цифры нельзя переписать")
)
