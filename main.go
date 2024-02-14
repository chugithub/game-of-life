package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameOfLife struct {
	size     int
	universe [][]int
}

func NewGameOfLife(size int) *GameOfLife {
	rand.Seed(time.Now().UnixNano())
	game := &GameOfLife{size: size}
	game.universe = make([][]int, size)
	for i := 0; i < size; i++ {
		game.universe[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if rand.Intn(2) == 1 {
				game.universe[i][j] = 'O'
			} else {
				game.universe[i][j] = ' '
			}
		}
	}
	return game
}

func (g *GameOfLife) Next() {
	size := g.size
	nextUniverse := make([][]int, size)

	for i := 0; i < size; i++ {
		nextUniverse[i] = make([]int, size)
		for j := 0; j < size; j++ {
			nextUniverse[i][j] = g.universe[i][j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			alive := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					if g.universe[(i+x+size)%size][(j+y+size)%size] == 'O' {
						alive++
					}
				}
			}
			if g.universe[i][j] == 'O' {
				if alive < 2 || alive > 3 {
					nextUniverse[i][j] = ' '
				}
			} else {
				if alive == 3 {
					nextUniverse[i][j] = 'O'
				}
			}
		}
	}

	g.universe = nextUniverse
}

func (g *GameOfLife) CountAlive() int {
	alive := 0
	size := g.size
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g.universe[i][j] == 'O' {
				alive++
			}
		}
	}
	return alive
}

func (g *GameOfLife) Print() {
	fmt.Println("Alive:", g.CountAlive())
	size := g.size
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%c", g.universe[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var size int
	fmt.Scan(&size)

	game := NewGameOfLife(size)

	for i := 0; i < 10; i++ {
		fmt.Printf("Generation #%d\n", i+1)
		game.Print()
		game.Next()
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
