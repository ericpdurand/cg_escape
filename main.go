package main

import "fmt"
import "os"
import "math"

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
	x, y, wallsLeft, id int
}

type Wall struct {
    x,y int
    wType string
}

func main() {

	w = 9
	h = 9
	playerCount = 3
	myId = 2
	fmt.Fprintln(os.Stderr, "Init: ", w, h, playerCount, myId)

	board = make([][]int, h)
	vWalls = make([][]int, h)
	hWalls = make([][]int, h+1)

	for {
		others = make([]Player, 0)
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
		x = 7
		y = 4
		wallsLeft = 5
		newOne := Player{x, y, wallsLeft, 1}
		others = append(others, newOne)
		fmt.Fprintf(os.Stderr, "Other: %v\n", newOne)
		board[y][x] = 0

		x = 6
		y = 2
		wallsLeft = 6
		newOne = Player{x, y, wallsLeft,2}
		others = append(others, newOne)
		fmt.Fprintf(os.Stderr, "Other: %v\n", newOne)
		board[y][x] = 1

		x = 1
		y = 7
		wallsLeft = 3
		me = Player{x, y, wallsLeft,0}
		fmt.Fprintf(os.Stderr, "Me: %v\n", me)
		board[y][x] = 2

		print(board)

		// wallCount: number of walls on the board
		wallCount = 3


		walls := []Wall{
                Wall{2, 7, "V"},
                Wall{7, 2, "H"},
                Wall{7, 4, "V"},
                Wall{7, 2, "V"},
        	}
        
		for _,w := range walls {
            		if w.wType=="V" {
                		vWalls[w.y][w.x] = 1
                		vWalls[w.y+1][w.x] = 1
            		} else {
                		hWalls[w.y][w.x] = 1
                		hWalls[w.y][w.x+1] = 1
            		}
        	}

		fmt.Fprintf(os.Stderr, "vWalls: \n")
        	simplePrint(vWalls)
        	fmt.Fprintf(os.Stderr, "hWalls: \n")
        	simplePrint(hWalls)

		fmt.Fprintf(os.Stderr, "Me: \n")
		scoreMap := buildMap(me.id)
		print(scoreMap)

		result,score := strategy(scoreMap, me.x,me.y)
	
		fmt.Fprintf(os.Stderr, "myResult: %v %v\n", result, score)
	
		fmt.Println(result)

		
		fmt.Fprintf(os.Stderr, "Other1: \n")
                scoreMap = buildMap(others[0].id)
                print(scoreMap)

                result, score = strategy(scoreMap,others[0].x,others[0].y)
		
		fmt.Fprintf(os.Stderr, "0Result: %v %v\n", result, score)

                fmt.Println(result)

		fmt.Fprintf(os.Stderr, "Me: \n")
                scoreMap = buildMap(others[1].id)
                print(scoreMap)

                result, score = strategy(scoreMap,others[1].x,others[1].y)
		fmt.Fprintf(os.Stderr, "1Result: %v %v\n", result, score)

                fmt.Println(result)
		
		result = messWith(1)
		fmt.Println("Messup result: %v",result)

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

func messWith(id int)(result string){
    player := others[id]
    fmt.Fprintf(os.Stderr, "Player  : %v\n",player)
    switch player.id {
        case 0 :{
            tx,ty,err := getSafeWall(player.x+1,player.y,"V")
            if err == nil {
                err = checkCrossing(tx,ty,"V")
		if err == nil {
			result =  fmt.Sprintf("%v %v V",tx,ty)
		}
            } else {
                fmt.Fprintf(os.Stderr, "Wall error  : %v\n",err)
                result = ""
            }
        }
        case 1 :{
            tx,ty,err := getSafeWall(player.x,player.y,"V")
            if err == nil {
		err = checkCrossing(tx,ty,"V")
		if err == nil {
                	result =  fmt.Sprintf("%v %v V",tx,ty)
		}
            } else {
                fmt.Fprintf(os.Stderr, "Wall error  : %v\n",err)
                result = ""
            }
        }
        case 2 :{
            tx,ty,err := getSafeWall(player.x,player.y+1,"H")
            if err == nil {
		err = checkCrossing(tx,ty,"H")
		if err == nil {
	 		result =  fmt.Sprintf("%v %v H",tx,ty)
		}
            } else {
                fmt.Fprintf(os.Stderr, "Wall error  : %v\n",err)
                result = ""
            }
        }
        default :{
            fmt.Fprintf(os.Stderr, "Woot?  : %v\n",id)
        }
    }
    return result
}

func checkCrossing(x,y int , wType string) error {
	if wType == "V" {
		if hWalls[y+1][x-1] !=0 && hWalls[y+1][x] != 0 {
			count := 1
			for i:=x+1;i<w;i++ {
				if hWalls[y+1][i] !=0 {
					count++
				} else {
					break
				}
			}
			if math.Trunc(float64(count)/2) == math.Trunc(float64(count+1)/2) {
				return nil
			} else {
				return fmt.Errorf("Forbidden wall in %v %v.",x,y)
			}			
		} else {
			return nil
		}
	} else {
		if vWalls[y-1][x+1] !=0 && vWalls[y][x+1] != 0 {
                        count := 1
                        for i:=y+1;i<h;i++ {
                                if vWalls[i][x+1] !=0 {
                                        count++
                                } else {
                                        break
                                }
                        }
                        if math.Trunc(float64(count)/2) == math.Trunc(float64(count+1)/2) {
                                return nil
                        } else {
                                return fmt.Errorf("Forbidden wall in %v %v.",x,y)
                        }
                } else {
                        return nil
                }
	}
}

func strategy(scoreMap [][]int, px, py int) (result string, score int) {
    result = "RIGHT"
    min := scoreMap[py][px]
    if px > 0 && scoreMap[py][px-1] < min && vWalls[py][px] == 0 {
        result = "LEFT"
        min = scoreMap[py][px-1]
    }
    if px < w-1 && scoreMap[py][px+1] < min && vWalls[py][px+1] == 0 {
        result = "RIGHT"
        min = scoreMap[py][px+1]
    }
    if py > 0 && scoreMap[py-1][px] < min && hWalls[py][px] == 0 {
        result = "UP"
        min = scoreMap[py-1][px]
    }
    if py < h-1 && scoreMap[py+1][px] < min && hWalls[py+1][px] == 0 {
        result = "DOWN"
        min = scoreMap[py+1][px]
    }
    return result, min
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
		for i := 0; i < h; i++ {
			if vWalls[i][w-1] == 0 {
				computePoint(&result, w-2, i)
			}
		}
	case 1:
		for i := 0; i < h; i++ {
			result[i][0] = 0
		}
		for i := 0; i < h; i++ {
                        if vWalls[i][1] == 0 {
                                computePoint(&result, 1, i)
                        }
                }
	case 2:
		for i := 0; i < w; i++ {
			result[h-1][i] = 0
		}
		for i := 0; i < w; i++ {
                        if hWalls[h-1][i] == 0 {
                                computePoint(&result, i, h-2)
                        }
                }
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
	if min<100 {
		(*data)[y][x] = min + 1
	}

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
