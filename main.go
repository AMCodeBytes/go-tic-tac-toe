package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	var moves [9][2]int

	fmt.Println("Welcome to TicTacToe")
	fmt.Println("1. Play")
	fmt.Println("2. Quit")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("You have chosen to play the game.")
		board()
		moves[0][0], moves[0][1] = move("X")
		fmt.Println(moves)
	case 2:
		fmt.Println("Goodbye, see you soon.")
		os.Exit(0)
	default:
		panic("Not a valid input.")
	}
}

func board() {
	fmt.Println("TicTacToe")
	fmt.Println("2     |   |  ")
	fmt.Println("   -----------")
	fmt.Println("1     |   |  ")
	fmt.Println("   -----------")
	fmt.Println("0     |   |  ")
	fmt.Println("   0    1   2")
}

func move(player string) (int, int) {
	var column, row int
	fmt.Printf("%v's Move\n", player)
	fmt.Print("Column: ")
	fmt.Scan(&column)
	fmt.Print("Row: ")
	fmt.Scan(&row)
	return column, row
}
