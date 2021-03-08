package game

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Game represents the state of the game
type Game struct {
	boardHeight       int
	boardWidth        int
	snake             *Snake
	score             int
	round             int
	lastMoveDirection string
	food              Coordinate
	freeCoordinateMap map[string]Coordinate
}

var (
	reader          = bufio.NewReader(os.Stdin)
	oppositeMoveMap = map[string]string{
		"E": "W",
		"W": "E",
		"N": "S",
		"S": "N",
	}
	moveStringMap = map[string]string{
		"E": "Press E to move East",
		"W": "Press W to move West",
		"S": "Press S to move South",
		"N": "Press N to move North",
	}
)

// Start initializes the game
func Start() {
	height, width := getBoardDimensions()
	c := Coordinate{
		X: rand.Intn(width),
		Y: rand.Intn(height),
	}
	g := &Game{
		boardHeight: height,
		boardWidth:  width,
		snake:       NewSnake(c),
		score:       0,
		round:       0,
	}
	g.initBoardCoordinateMap()
	g.removeFreeCoordinate(c)
	g.placeFood()
	for {
		snakeFeeded := false
		g.printState()
		newHeadPos := g.nextMove()
		g.removeFreeCoordinate(newHeadPos)
		isFree := !g.snake.IsBodyOnPosition(newHeadPos)
		if g.snakeHasFood(newHeadPos) {
			snakeFeeded = true
			g.score++
			g.snake.Feed()
			g.placeFood()
		}
		freeCoordinate := g.snake.PositionSnakeAndReturnFreeCoordinate(newHeadPos)
		g.addFreeCoordinate(freeCoordinate)
		if g.isOutOfBoundary(g.snake.Head()) ||
			(isFree == false && !snakeFeeded && !freeCoordinate.Equal(newHeadPos)) {
			fmt.Println("Game Over !!! Your snake died!!!")
			break
		}
		if g.size() == g.snake.Length() {
			fmt.Println("This world is too small to fit in your snake")
			break
		}
	}
}

func (g *Game) initBoardCoordinateMap() {
	m := make(map[string]Coordinate, g.boardHeight*g.boardWidth)
	for y := 0; y < g.boardHeight; y++ {
		for x := 0; x < g.boardWidth; x++ {
			c := Coordinate{
				X: x,
				Y: y,
			}
			key := g.getCoordinateKey(c)
			m[key] = c
		}
	}
	g.freeCoordinateMap = m
}

func getBoardDimensions() (int, int) {
	for {
		fmt.Println("Enter the height of the snake world : ")
		var height, width int
		_, err := fmt.Scanln(&height)
		if err != nil {
			fmt.Println("Some error occurred. Try again!")
			continue
		}
		fmt.Println("Enter the width of the snake world : ")
		_, err = fmt.Scanln(&width)
		if err != nil {
			fmt.Println("Some error occurred. Try again!")
			continue
		}
		return height, width
	}
}

func (g *Game) nextMove() Coordinate {
	for {
		fmt.Println("Enter your next move.")
		for key := range oppositeMoveMap {
			if key != oppositeMoveMap[g.lastMoveDirection] && g.snake.Length() > 1 {
				fmt.Println(moveStringMap[key])
			}
		}
		var move string
		_, err := fmt.Scanln(&move)
		if err != nil {
			fmt.Println("Some error occurred! Try again.")
			continue
		}
		move = strings.ToUpper(move)
		_, ok := moveStringMap[move]
		if !ok || (move == oppositeMoveMap[g.lastMoveDirection] && g.snake.Length() > 1) {
			fmt.Println("Invalid move choice. Please try again")
			continue
		}
		g.round++
		g.lastMoveDirection = move
		return g.snake.NewHeadPos(move)
	}
}

func (g *Game) getFreeCoordinates() (Coordinate, error) {
	for _, coordinate := range g.freeCoordinateMap {
		return coordinate, nil
	}
	return Coordinate{}, errors.New("No free Coordinate")
}

func (g *Game) placeFood() error {
	var err error
	g.food, err = g.getFreeCoordinates()
	if err != nil {
		return err
	}
	g.removeFreeCoordinate(g.food)
	return nil
}

func (g *Game) getCoordinateKey(c Coordinate) string {
	return strconv.Itoa(c.X) + strconv.Itoa(c.Y)
}

func (g *Game) removeFreeCoordinate(c Coordinate) {
	delete(g.freeCoordinateMap, g.getCoordinateKey(c))
}

func (g *Game) addFreeCoordinate(c Coordinate) {
	if c.X >= 0 && c.Y >= 0 {
		g.freeCoordinateMap[g.getCoordinateKey(c)] = c
	}
}

func (g *Game) snakeHasFood(c Coordinate) bool {
	return g.food.Equal(c)
}

func (g *Game) isOutOfBoundary(pos Coordinate) bool {
	return pos.X >= g.boardWidth || pos.Y >= g.boardHeight ||
		pos.X < 0 || pos.Y < 0
}

func (g *Game) size() int {
	return g.boardHeight * g.boardWidth
}

func (g *Game) printState() {
	fmt.Println("---------------------------------------------------------------")
	head := g.snake.Head()
	for i := 0; i < g.boardHeight; i++ {
		for j := 0; j < g.boardWidth; j++ {
			c := Coordinate{
				X: j,
				Y: i,
			}
			printChar := "."
			if head.Equal(c) {
				printChar = "*"
			} else if g.snake.IsBodyOnPosition(c) {
				printChar = "#"
			} else if g.food.Equal(c) {
				printChar = "@"
			}
			fmt.Print(printChar)
		}
		fmt.Println()
	}
	fmt.Println("Current Round : ", g.round)
	fmt.Println("Score : ", g.score)
	fmt.Println("Length of snake : ", g.snake.Length())
	fmt.Println("Head Coordinates : ", head)
	fmt.Println("---------------------------------------------------------------")
}
