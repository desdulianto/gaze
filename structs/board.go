package structs

import (
	"io"
)

type Board struct {
	Height uint16 // row
	Width  uint16 // col
	Cells  [][]Cell
}

func (b *Board) Init() {
	b.Cells = make([][]Cell, b.Height)
	for i := uint16(0); i < b.Height; i++ {
		b.Cells[i] = make([]Cell, b.Width)
	}

	for h := uint16(0); h < b.Height; h++ {
		for w := uint16(0); w < b.Width; w++ {
			// set the flag field with 15, which in binary will be 0b00001111,
			// the 4 bits indicates that all 4 walls are up, so the cells are
			// isolated/sealed from each other initially. ex:
			//  _ _
			// |_|_|
			// |_|_|
			//
			b.Cells[h][w].Flag = 15

			// set the relative [x,y] position of the cell on the board
			b.Cells[h][w].X = h
			b.Cells[h][w].Y = w
		}
	}
}

// Neighbors find the neighboring cells of the current cell
func (b *Board) Neighbors(c *Cell) []*Cell {
	result := []*Cell{}
	if ok, cell := b.getNeighbor(c.X+1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X-1, c.Y); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X, c.Y+1); ok {
		result = append(result, cell)
	}

	if ok, cell := b.getNeighbor(c.X, c.Y-1); ok {
		result = append(result, cell)
	}
	return result
}

func (b *Board) GetDirection(from, to *Cell) FlagPosition {
	// X denotes row, Y denotes col
	//
	//        col 0  | col 1 | col 2
	// --------------------------------
	// row 0  [x0,y0] [x0,y1] [x0,y2]
	// row 1  [x1,y0] [x1,y1] [x1,y2]
	if from.X < to.X {
		return SOUTH
	}
	if from.X > to.X {
		return NORTH
	}
	if from.Y < to.Y {
		return EAST
	}
	if from.Y > to.Y {
		return WEST
	}
	//TODO: This is really an error case here
	return EAST
}

func (b *Board) BreakWall(from, to *Cell, direction FlagPosition) {
	from.SetBit(VISITED)
	to.SetBit(VISITED)

	switch direction {
	case EAST:
		from.ClearBit(EAST)
		to.ClearBit(WEST)
	case SOUTH:
		from.ClearBit(SOUTH)
		to.ClearBit(NORTH)
	case WEST:
		from.ClearBit(WEST)
		to.ClearBit(EAST)
	case NORTH:
		from.ClearBit(NORTH)
		to.ClearBit(SOUTH)
	}
}

func (b *Board) Write(writer io.Writer) {
	writer.Write([]byte("  "))
	for i := uint16(1); i < b.Width; i++ {
		writer.Write([]byte(" _"))
	}
	writer.Write([]byte("\n"))

	for h := uint16(0); h < b.Height; h++ {
		writer.Write([]byte("|"))
		for w := uint16(0); w < b.Width; w++ {
			c := b.Cells[h][w]
			if w == b.Width-1 && h == b.Height-1 {
				writer.Write([]byte(" |"))
				break
			}
			if c.IsSet(SOUTH) {
				writer.Write([]byte("_"))
			} else {
				writer.Write([]byte(" "))
			}

			if c.IsSet(EAST) {
				writer.Write([]byte("|"))
			} else {
				writer.Write([]byte(" "))
			}
		}
		writer.Write([]byte("\n"))
	}
}

func (b *Board) getNeighbor(x, y uint16) (bool, *Cell) {
	if x >= 0 && x < b.Height &&
		y >= 0 && y < b.Width &&
		!b.Cells[x][y].IsSet(VISITED) {
		return true, &b.Cells[x][y]
	}
	return false, nil
}