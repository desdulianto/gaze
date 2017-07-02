package solvers

import (
	"github.com/wliao008/mazing/structs"
	"github.com/wliao008/mazing/util"
)

type DeadEndFiller struct {
	Board *structs.Board
}

func NewDeadEndFiller() *DeadEndFiller {
	def := &DeadEndFiller{}
	return def
}

func (def *DeadEndFiller) Solve() {
	// get initial list of dead ends
	// while there are more dead ends
	//    mark current cell as a dead end
	//    find dead neighbors of the current cell and add to list
	stack := &util.Stack{}
	def.Board.DeadEnds(stack)
	c := stack.Pop()
	for c != nil {	
		c.(*structs.Cell).SetBit(structs.DEAD)
		def.Board.DeadEnds(stack)
		c = stack.Pop()
	}
}