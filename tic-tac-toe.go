package main

import "fmt"

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

// func (b *Board) put(x, y int, u string) // interface
// func (b *Board) get(x, y int) string    // interface

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = -1
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == -1 {
		return "x"
	} else {
		return "."
	}
}

func (b *Board) show() {
	for i := 0; i < 9; i++ {
		x := i % 3
		y := i / 3
		fmt.Print(b.get(x, y))
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

//勝敗チェック
func (b *Board) judge() string {
	if b.tokens[0]+b.tokens[3]+b.tokens[6] == 3 {
		return "o"
	} else if b.tokens[1]+b.tokens[4]+b.tokens[7] == 3 {
		return "o"
	} else if b.tokens[2]+b.tokens[5]+b.tokens[8] == 3 {
		return "o"
	} else if b.tokens[0]+b.tokens[1]+b.tokens[2] == 3 {
		return "o"
	} else if b.tokens[3]+b.tokens[4]+b.tokens[5] == 3 {
		return "o"
	} else if b.tokens[6]+b.tokens[7]+b.tokens[8] == 3 {
		return "o"
	} else if b.tokens[0]+b.tokens[4]+b.tokens[8] == 3 {
		return "o"
	} else if b.tokens[2]+b.tokens[4]+b.tokens[6] == 3 {
		return "o"
	} else if b.tokens[0]+b.tokens[3]+b.tokens[6] == -3 {
		return "x"
	} else if b.tokens[1]+b.tokens[4]+b.tokens[7] == -3 {
		return "x"
	} else if b.tokens[2]+b.tokens[5]+b.tokens[8] == -3 {
		return "x"
	} else if b.tokens[0]+b.tokens[1]+b.tokens[2] == -3 {
		return "x"
	} else if b.tokens[3]+b.tokens[4]+b.tokens[5] == -3 {
		return "x"
	} else if b.tokens[6]+b.tokens[7]+b.tokens[8] == -3 {
		return "x"
	} else if b.tokens[0]+b.tokens[4]+b.tokens[8] == -3 {
		return "x"
	} else if b.tokens[2]+b.tokens[4]+b.tokens[6] == -3 {
		return "x"
	} else {
		return "undergo"
	}
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	b.show()
}
