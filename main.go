package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var (
	w, h, playerCount, myId int
	me                      Player
	others                  []Player
	wallCount               int
	board                   [][]int
	vWalls                  [][]int
	hWalls                  [][]int
)

type Player struct {
	x, y, wallsLeft int
}

func main() {

	w = 9
	h = 9
	playerCount = 3
	myId = 2
	fmt.Fprintln(os.Stderr, "Init: ", w, h, playerCount, myId)

	others = make([]Player, playerCount-1)
	board = make([][]int, h)
	vWalls = make([][]int, h)
	hWalls = make([][]int, h+1)

	for {
		for i := range board {
			board[i] = make([]int, w)
			for j := range board[i] {
				board[i][j] = -1
			}
		}
		for i := range vWalls {
			vWalls[i] = make([]int, w+1)
			for j := range vWalls[i] {
				vWalls[i][j] = 0
			}
			vWalls[i][0] = 1
			vWalls[i][w] = 1
		}

		for i := range hWalls {
			hWalls[i] = make([]int, w)
		}
		for i := range hWalls {
			for j := range hWalls[i] {
				hWalls[i][j] = 0
				hWalls[0][j] = 1
				hWalls[h][j] = 1
			}
		}

		var x, y, wallsLeft int
		x = 0
		y = 3
		wallsLeft = 6
		newOne := Player{x, y, wallsLeft}
		others = append(others, newOne)
		fmt.Fprintf(os.Stderr, "Other: %v\n", newOne)
		board[y][x] = 0

		x = 8
		y = 5
		wallsLeft = 6
		newOne = Player{x, y, wallsLeft}
		others = append(others, newOne)
		fmt.Fprintf(os.Stderr, "Other: %v\n", newOne)
		board[y][x] = 1

		x = 6
		y = 0
		wallsLeft = 6
		me = Player{x, y, wallsLeft}
		fmt.Fprintf(os.Stderr, "Me: %v\n", me)
		board[y][x] = 2

		print(board)

		// wallCount: number of walls on the board
		wallCount = 3

		var wallX, wallY int
		var wallOrientation string
		wallX = 2
		wallY = 6
		wallOrientation = "V"
		handleWall(wallX, wallY, wallOrientation)
		wallX = 3
		wallY = 4
		wallOrientation = "V"
		handleWall(wallX, wallY, wallOrientation)
		wallX = 6
		wallY = 2
		wallOrientation = "H"
		handleWall(wallX, wallY, wallOrientation)

		fmt.Fprintf(os.Stderr, "vWalls: \n")
		simplePrint(vWalls)
		fmt.Fprintf(os.Stderr, "hWalls: \n")
		simplePrint(hWalls)

		wx := 2
		wy := 6
		tx, ty, err := getSafeWall(wx, wy, "V")
		fmt.Fprintf(os.Stderr, "Test wall %v %v %v %v %v\n", wx, wy, tx, ty, err)

		wx = 2
		wy = 5
		tx, ty, err = getSafeWall(wx, wy, "V")
		fmt.Fprintf(os.Stderr, "Test wall %v %v %v %v %v\n", wx, wy, tx, ty, err)

		wx = 3
		wy = 8
		tx, ty, err = getSafeWall(wx, wy, "V")
		fmt.Fprintf(os.Stderr, "Test wall %v %v %v %v %v\n", wx, wy, tx, ty, err)

		fmt.Fprintf(os.Stderr, "First: \n")
		scoreMap := buildMap(0)
		print(scoreMap)

		result := strategy(scoreMap)

		fmt.Println(result)
		fmt.Fprintf(os.Stderr, "Second: \n")
		scoreMap = buildMap(1)
		print(scoreMap)

		result = strategy(scoreMap)

		fmt.Println(result)
		fmt.Fprintf(os.Stderr, "Last: \n")
		scoreMap = buildMap(2)
		print(scoreMap)

		result = strategy(scoreMap)

		fmt.Println(result)

		break
	}
}

func handleWall(wallX, wallY int, wallOrientation string) {
	if wallOrientation == "V" {
		vWalls[wallY][wallX] = 1
		vWalls[wallY+1][wallX] = 1
	} else {
		hWalls[wallY][wallX] = 1
		hWalls[wallY][wallX+1] = 1
	}
}

func strategy(scoreMap [][]int) (result string) {
	result = "RIGHT"
	min := scoreMap[me.y][me.x]
	if me.x > 0 && scoreMap[me.y][me.x-1] < min && vWalls[me.y][me.x] == 0 {
		result = "LEFT"
		min = scoreMap[me.y][me.x-1]
	}
	if me.x < w-1 && scoreMap[me.y][me.x+1] < min && vWalls[me.y][me.x+1] == 0 {
		result = "RIGHT"
		min = scoreMap[me.y][me.x+1]
	}
	if me.y > 0 && scoreMap[me.y-1][me.x] < min && hWalls[me.y][me.x] == 0 {
		result = "UP"
		min = scoreMap[me.y-1][me.x]
	}
	if me.y < h-1 && scoreMap[me.y+1][me.x] < min && hWalls[me.y+1][me.x] == 0 {
		result = "DOWN"
		min = scoreMap[me.y+1][me.x]
	}
	return result
}

func simplePrint(data [][]int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] < 0 {
				fmt.Fprintf(os.Stderr, " .")
			} else {
				fmt.Fprintf(os.Stderr, " %v", data[i][j])
			}
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func print(data [][]int) {
	for i := range data {
		for j := range hWalls[i] {
			if hWalls[i][j] == 0 {
				fmt.Fprintf(os.Stderr, "    .")
			} else {
				fmt.Fprintf(os.Stderr, "    =")
			}
		}
		fmt.Fprintf(os.Stderr, " \n")
		for j := range data[i] {
			if vWalls[i][j] == 0 {
				fmt.Fprintf(os.Stderr, " ")
			} else {
				fmt.Fprintf(os.Stderr, "|")
			}
			if data[i][j] < 0 {
				fmt.Fprintf(os.Stderr, "  . ")
			} else {
				fmt.Fprintf(os.Stderr, " %3v", data[i][j])
			}
		}
		fmt.Fprintf(os.Stderr, "|")
		fmt.Fprintf(os.Stderr, "\n")
	}
	for j := range hWalls[h] {
		if hWalls[h][j] == 0 {
			fmt.Fprintf(os.Stderr, "    .")
		} else {
			fmt.Fprintf(os.Stderr, "    =")
		}
	}
	fmt.Fprintf(os.Stderr, "\n")
}

func getSafeWall(x, y int, wType string) (tx, ty int, err error) {
	if wType == "V" {
		if vWalls[y][x] == 1 {
			return 0, 0, fmt.Errorf("Already a wall in %v %v.", x, y)
		}
		if y < h-1 {
			if vWalls[y+1][x] != 1 {
				return x, y, nil
			}
			if y > 0 && vWalls[y-1][x] != 1 {
				return x, y - 1, nil
			}
			return 0, 0, fmt.Errorf("Cannot put a VWall in %v %v.", x, y)
		} else {
			if vWalls[y-1][x] != 1 {
				return x, y - 1, nil
			}
			return 0, 0, fmt.Errorf("Cannot put a VWall in %v %v.", x, y)
		}

	} else {
		if hWalls[y][x] == 1 {
			return 0, 0, fmt.Errorf("Already a wall in %v %v.", x, y)
		}
		if x < w-1 {
			if hWalls[y][x+1] != 1 {
				return x, y, nil
			}
			if x > 0 && hWalls[y][x-1] != 1 {
				return x - 1, y, nil
			}
			return 0, 0, fmt.Errorf("Cannot put a HWall in %v %v.", x, y)
		} else {
			if hWalls[y][x-1] != 1 {
				return x - 1, y, nil
			}
			return 0, 0, fmt.Errorf("Cannot put a HWall in %v %v.", x, y)
		}
	}
}

func buildMap(target int) (result [][]int) {
	result = make([][]int, h)
	for i := range result {
		result[i] = make([]int, w)
		for j := range result[i] {
			result[i][j] = 100
		}
	}
	switch target {
	case 0:
		for i := 0; i < h; i++ {
			result[i][w-1] = 0
		}
		computePoint(&result, w-2, 0)
	case 1:
		for i := 0; i < h; i++ {
			result[i][0] = 0
		}
		computePoint(&result, 1, 0)
	case 2:
		for i := 0; i < w; i++ {
			result[h-1][i] = 0
		}
		computePoint(&result, 0, h-2)
	}
	return result
}

func computePoint(data *[][]int, x, y int) {
	min := 100

	if x > 0 && (*data)[y][x-1] < min && vWalls[y][x] == 0 {
		min = (*data)[y][x-1]
	}
	if x < w-1 && (*data)[y][x+1] < min && vWalls[y][x+1] == 0 {
		min = (*data)[y][x+1]
	}
	if y > 0 && (*data)[y-1][x] < min && hWalls[y][x] == 0 {
		min = (*data)[y-1][x]
	}
	if y < h-1 && (*data)[y+1][x] < min && hWalls[y+1][x] == 0 {
		min = (*data)[y+1][x]
	}
	(*data)[y][x] = min + 1

	if x > 0 && vWalls[y][x] == 0 && (*data)[y][x-1] > (*data)[y][x]+1 {
		//fmt.Fprintf(os.Stderr, " call computePoint: %v %v %v, %v %v %v \n", x, y, (*data)[y][x], x-1, y, (*data)[y][x-1])
		computePoint(data, x-1, y)
	}
	if x < w-1 && vWalls[y][x+1] == 0 && (*data)[y][x+1] > (*data)[y][x]+1 {
		//fmt.Fprintf(os.Stderr, " call computePoint: %v %v %v, %v %v %v \n", x, y, (*data)[y][x], x+1, y, (*data)[y][x+1])
		computePoint(data, x+1, y)
	}
	if y > 0 && hWalls[y][x] == 0 && (*data)[y-1][x] > (*data)[y][x]+1 {
		//fmt.Fprintf(os.Stderr, " call computePoint: %v %v %v, %v %v %v \n", x, y, (*data)[y][x], x, y-1, (*data)[y-1][x])
		computePoint(data, x, y-1)
	}
	if y < h-1 && hWalls[y+1][x] == 0 && (*data)[y+1][x] > (*data)[y][x]+1 {
		//fmt.Fprintf(os.Stderr, " call computePoint: %v %v %v, %v %v %v \n", x, y, (*data)[y][x], x, y+1, (*data)[y+1][x])
		computePoint(data, x, y+1)
	}
}
