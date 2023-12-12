package main

import (
	"fmt"
	"math/rand"
)

// Snake
// Ladder
// Player

const (
	WIN_POSITION = 100
)

type SLPlayer struct {
	id   int64
	name string
}

type SnakeAndLadderBoard struct {
	snakes  map[int32]int32
	ladders map[int32]int32
}

func NewSnakeAndLadderBoard() *SnakeAndLadderBoard {
	return &SnakeAndLadderBoard{
		snakes:  map[int32]int32{},
		ladders: map[int32]int32{},
	}
}

func (b *SnakeAndLadderBoard) addSnake(start int32, end int32) {
	if b == nil {
		panic("initialize board first")
	}

	if !isValidPosition(start) || !isValidPosition(end) || start <= end {
		panic("invalid start/end position(s)")
	}

	b.snakes[start] = end
}

func (b *SnakeAndLadderBoard) addLadder(start int32, end int32) {
	if b == nil {
		panic("initialize board first")
	}

	if !isValidPosition(start) || !isValidPosition(end) || start >= end {
		panic("invalid start/end position(s)")
	}

	b.ladders[start] = end
}

func isValidPosition(pos int32) bool {
	return pos <= 100 && pos >= 1
}

type SnakeAndLadderGame struct {
	players   []SLPlayer
	gameState map[string]int32
	board     SnakeAndLadderBoard
}

func NewSnakeAndLadderGame(players []SLPlayer, board SnakeAndLadderBoard) *SnakeAndLadderGame {
	game := &SnakeAndLadderGame{
		players:   players,
		gameState: map[string]int32{},
		board:     board,
	}

	for _, player := range players {
		game.gameState[player.name] = 0
	}

	return game
}

func (g *SnakeAndLadderGame) startGame() string {

	for {
		for _, player := range g.players {
			currPos := g.gameState[player.name]
			count := 0
			num := int32(0)
			for {
				num = rollDice()
				if currPos+num > WIN_POSITION {
					break
				}

				if num < 6 || count == 3 {
					printTheRoll(player.name, num, currPos, currPos+num)
					currPos += num
					break
				}

				printTheRoll(player.name, num, currPos, currPos+num)
				currPos += num
				count += 1
			}

			if currPos == WIN_POSITION {
				fmt.Printf("%s won the game\n", player.name)
				return player.name
			}

			if snakeEnd, ok := g.board.snakes[currPos]; ok {
				fmt.Printf("%s landed on a **SNAKE**, moved down from %d to %d\n", player.name, currPos, snakeEnd)
				currPos = snakeEnd
			} else if ladderEnd, ok2 := g.board.ladders[currPos]; ok2 {
				fmt.Printf("%s landed on a **LADDER**, moved up from %d to %d\n", player.name, currPos, ladderEnd)
				currPos = ladderEnd
			}

			g.gameState[player.name] = currPos
		}
	}

}

func rollDice() int32 {
	return rand.Int31n(6) + 1
}

func printTheRoll(name string, roll int32, oldPos int32, newPos int32) {
	fmt.Printf("%s rolled a %d and moved from %d to %d\n", name, roll, oldPos, newPos)
}
