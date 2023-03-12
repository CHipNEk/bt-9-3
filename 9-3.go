package main

import (
	"fmt"
)

var board = [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var player = "X"
var winner = " "

func printboard() {

	fmt.Println("|", board[0], "|", board[1], "|", board[2], "|")
	fmt.Println("-------------")
	fmt.Println("|", board[3], "|", board[4], "|", board[5], "|")
	fmt.Println("-------------")
	fmt.Println("|", board[6], "|", board[7], "|", board[8], "|")
}

func isboardfull() bool {

	for _, value := range board {
		if value == " " {
			return false
		}
	}
	return true
}

func getmove() int {

	var move int
	for {
		fmt.Println("nhập số từ 0 đến 8: ")
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("sai số, nhập lại số trong khoảng 0 đến 8: ")
			continue
		}
		if move < 0 || move > 8 {
			fmt.Println("ko di chuyển dc, nhập lại số trong khoảng 0 đến 8: ")
			continue
		}
		if board[move] != " " {
			fmt.Println("ô đã bị lấy, nhập số khác: ")
			continue
		}
		break

	}
	return move
}

func checkwinner() string {
	for i := 0; i < 3; i++ {
		if board[i*3] != " " && board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] {
			return board[i*3]
		}
		if board[i] != " " && board[i] == board[i+3] && board[i+3] == board[i+6] {
			return board[i]
		}
	}
	if board[0] != " " && board[0] == board[4] && board[4] == board[8] {
		return board[0]
	}
	if board[2] != " " && board[2] == board[4] && board[4] == board[6] {
		return board[2]
	}
	return " "

}

func switchplayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}

func main() {

	fmt.Println("BẮT ĐẦU")
	printboard()
	for winner == " " && !isboardfull() {

		fmt.Printf("người chơi %s đi", player)
		move := getmove()
		board[move] = player
		printboard()
		winner = checkwinner()
		switchplayer()

	}
	if winner != " " {
		fmt.Printf("người chơi %s thắng\n", winner)
	} else {
		fmt.Println("hòa")
	}
}
