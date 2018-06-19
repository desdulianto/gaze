package models

import "github.com/wliao008/gaze/structs"

type BoardModel struct {
	Name          string
	Cells         [][]CellModel
	RawCells      [][]structs.Cell
	Height, Width uint16
	TableCss      string
	WeaveChecked  string
}

type CellModel struct {
	X, Y       uint16
	CssClasses string
}
