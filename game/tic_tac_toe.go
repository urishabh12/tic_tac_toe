package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ticTacToe struct {
	board         [3][3]string
	currentPlayer string
	finished      bool
}

func NewTicTacToeGame() *ticTacToe {
	//Game starts with player X
	return &ticTacToe{
		currentPlayer: "X",
	}
}

func (t *ticTacToe) StartGame() {
	var winner string
	for !t.finished {
		t.drawBoard()
		i, j := t.inputForPosition()
		t.addInputToBoard(i, j)
		winner = t.checkForWin(i, j)
		t.changeTurn()
		t.clearScreenLinux()
	}

	t.drawBoard()
	fmt.Printf("[Result] Game has finished\nThe Winner is %s\n", winner)
}

func (t *ticTacToe) drawBoard() {
	for i := 0; i < len(t.board); i++ {
		for j := 0; j < len(t.board[0]); j++ {
			printVal := t.board[i][j]
			if printVal == "" {
				printVal = "-"
			}

			fmt.Printf("%s ", printVal)
		}
		fmt.Println("")
	}
}

func (t *ticTacToe) inputForPosition() (int, int) {
	var i int
	var j int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("It's player %s's turn please give your input index starts with 1\n", t.currentPlayer)
		var pos string
		for scanner.Scan() {
			pos = scanner.Text()
			break
		}
		positions := strings.Split(pos, " ")
		if len(positions) != 2 {
			fmt.Println("[Error] Incorrect input should have 2 integers")
			continue
		}

		var err error
		i, err = strconv.Atoi(positions[0])
		if err != nil {
			fmt.Println("[Error] Incorrect input should have 2 integers ", err)
			continue
		}

		j, err = strconv.Atoi(positions[1])
		if err != nil {
			fmt.Println("[Error] Incorrect input should have 2 integers ", err)
			continue
		}

		if !t.validInputRange(i, j) {
			fmt.Println("[Error] Invalid range")
			continue
		}

		i--
		j--
		if t.board[i][j] != "" {
			fmt.Println("[Error] Cordinate already contains input")
			continue
		}

		break
	}

	return i, j
}

func (t *ticTacToe) addInputToBoard(i, j int) {
	t.board[i][j] = t.currentPlayer
}

func (t *ticTacToe) changeTurn() {
	if t.currentPlayer == "X" {
		t.currentPlayer = "0"
		return
	}

	t.currentPlayer = "X"
}

func (t *ticTacToe) clearScreenLinux() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (t *ticTacToe) checkForWin(i, j int) string {
	//Check top down
	topDownTrue := true
	for p := 0; p < len(t.board); p++ {
		if t.board[p][j] != t.currentPlayer {
			topDownTrue = false
			break
		}
	}

	//Check left to right
	leftRightTrue := true
	for p := 0; p < len(t.board[0]); p++ {
		if t.board[i][p] != t.currentPlayer {
			leftRightTrue = false
			break
		}
	}

	//Check for Diagonal
	forwardCenterDiagonalIsTrue := true
	for p := 0; p < len(t.board); p++ {
		if t.board[p][p] != t.currentPlayer {
			forwardCenterDiagonalIsTrue = false
		}
	}

	reverseCenterDiagonalIsTrue := true
	q := 0
	for p := len(t.board) - 1; p >= 0; p-- {
		if t.board[p][q] != t.currentPlayer {
			reverseCenterDiagonalIsTrue = false
		}
		q++
	}

	t.finished = topDownTrue || leftRightTrue || forwardCenterDiagonalIsTrue || reverseCenterDiagonalIsTrue
	return t.currentPlayer
}

func (t *ticTacToe) validInputRange(i, j int) bool {
	return !(i <= 0 || i > len(t.board) || j <= 0 || j > len(t.board[0]))
}
