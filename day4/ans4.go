package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, boards := boardBuilder(read("input4.dat"))

	//p1
	var wonNumber, wonBoard int

out:
	for _, n := range nums {
		fmt.Println(n)
		for i, b := range boards {
			boards[i] = setBoard(n, -1, b)
			if isBingo(b, -1) {
				fmt.Println("board", i, "bingo")
				wonNumber = n
				wonBoard = i
				break out
			}
		}
	}

	sumUnselected := 0

	for i := 0; i < len(boards[wonBoard]); i++ {
		for j := 0; j < len(boards[wonBoard][0]); j++ {
			if boards[wonBoard][i][j] != -1 {
				sumUnselected += boards[wonBoard][i][j]
			}
		}
	}

	fmt.Println(wonNumber * sumUnselected)

	//p2
	numsP2, boardsP2 := boardBuilder(read("input4.dat"))
	fmt.Println("board length", len(boardsP2))
	fmt.Println(boardsP2)
	var lastWonNumber, lastWonBoard int
	wonBoardMap := make(map[int]bool)

last:
	for _, n := range numsP2 {
		fmt.Println(n)
		for i, b := range boardsP2 {
			boardsP2[i] = setBoard(n, -1, b)
			_, ok := wonBoardMap[i]

			if !ok && isBingo(b, -1) {
				fmt.Println("board", i, "bingo")
				wonBoardMap[i] = true

				if len(wonBoardMap) == len(boardsP2) {
					lastWonNumber = n
					lastWonBoard = i
					fmt.Println("last won number", lastWonNumber, "last won board", lastWonBoard)
					break last
				}
			}

		}
	}

	sumLastUnselected := 0
	for i := 0; i < len(boardsP2[lastWonBoard]); i++ {
		for j := 0; j < len(boardsP2[lastWonBoard][0]); j++ {
			if boardsP2[lastWonBoard][i][j] != -1 {
				sumLastUnselected += boardsP2[lastWonBoard][i][j]
			}
		}
	}

	fmt.Println(lastWonNumber * sumLastUnselected)

}

func read(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func boardBuilder(lines []string) ([]int, [][][]int) {
	var nums []int
	var boards [][][]int
	var board [][]int

	for i, l := range lines {
		if i == 0 {
			splits := strings.Split(l, ",")
			for _, s := range splits {
				num, _ := strconv.Atoi(s)
				nums = append(nums, num)
			}
		} else if i > 1 && len(l) != 0 {
			var row []int
			stringNums := strings.Fields(l)
			for _, stringNum := range stringNums {
				num, _ := strconv.Atoi(stringNum)
				row = append(row, num)
			}
			board = append(board, row)
		} else if i > 1 && len(l) == 0 {
			boards = append(boards, board)
			board = [][]int{}
		}
	}
	//last line
	boards = append(boards, board)

	return nums, boards
}

func isBingo(board [][]int, mark int) bool {
	bingo := false
	counter := 0
	//check row
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == mark {
				counter++
			}
		}
		if counter == len(board[0]) {
			return true
		}
		counter = 0
	}
	counter = 0

	//check column
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			if board[j][i] == mark {
				counter++
			}
		}

		if counter == len(board) {
			return true
		}
		counter = 0
	}

	return bingo
}

func setBoard(num, mark int, board [][]int) [][]int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == num {
				board[i][j] = mark
			}
		}
	}
	return board
}
