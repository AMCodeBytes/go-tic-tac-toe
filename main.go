package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	var moves [9][2]int
	var boardLayout [3][3]string

	for {
		fmt.Println("Welcome to TicTacToe")
		fmt.Println("1. Play")
		fmt.Println("2. Quit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("You have chosen to play the game.")

			for i := 0; i < 9; i++ {
				board(boardLayout)
				moves[i][0], moves[i][1] = move(i)
				if i%2 == 0 {
					boardLayout[moves[i][0]][moves[i][1]] = "X"
				} else {
					boardLayout[moves[i][0]][moves[i][1]] = "O"
				}

			}

			fmt.Println("ALL MOVES EXHAUSTED")
			board(boardLayout)
			fmt.Println(moves)
			return
		case 2:
			fmt.Println("Goodbye, see you soon.")
			os.Exit(0)
			return
		default:
			panic("Not a valid input.")
		}
	}
}

func board(boardLayout [3][3]string) {
	fmt.Println("TicTacToe")
	fmt.Printf("2    %v |  %v | %v \n", boardLayout[2][0], boardLayout[2][1], boardLayout[2][2])
	fmt.Printf("   -------------\n")
	fmt.Printf("1    %v |  %v | %v \n", boardLayout[1][0], boardLayout[1][1], boardLayout[1][2])
	fmt.Printf("   -------------\n")
	fmt.Printf("0    %v |  %v | %v \n", boardLayout[0][0], boardLayout[0][1], boardLayout[0][2])
	fmt.Printf("     0    1   2\n")
}

func move(round int) (int, int) {
	var player string
	var column, row int

	if round%2 == 0 {
		player = "X"
	} else {
		player = "O"
	}

	fmt.Printf("Round %v, %v's Move\n", round+1, player)
	fmt.Print("Column: ")
	fmt.Scan(&column)
	fmt.Print("Row: ")
	fmt.Scan(&row)

	return column, row
}

func winCheck() {
	// Check to see if there is a winner
}
