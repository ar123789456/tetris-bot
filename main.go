package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var allSQ = 5

func main() {
	Tetris()
}

type tetraminoStr struct {
	symbol    string
	tetramino [][2]int
}

func Tetris() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		tetrisERROR(err)
	}
	fileStr := string(file)
	fmt.Println(fileStr)
	alltetramino := splitTetramino(fileStr)
	finishTetris(alltetramino)

}

func finishTetris(alltetramino []tetraminoStr) {
	b := false
	a := [][]string{}
	fmt.Println(alltetramino)
	// for !b {
	fmt.Println(allSQ)
	a = genaretSQ(allSQ)
	a, b = decidedTetris(alltetramino, a, 0, 0)
	allSQ++
	// }
	fmt.Println(b)
	for _, i := range a {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

//TODO Write this shit
func decidedTetris(alltetramino []tetraminoStr, sq [][]string, x, y int) ([][]string, bool) {

	if len(alltetramino) == 0 {
		return sq, true
	}
	n := len(sq)
	if x >= n {
		return sq, false
	}
	fmt.Println(x, y)

	for j, i := range alltetramino {
		x1 := i.tetramino[0][0] + x
		y1 := i.tetramino[0][1] + y
		if x1 < 0 || x1 >= n || y1 < 0 || y1 >= n || sq[x1][y1] != "." {
			a, b := decidedTetris(alltetramino, sq, x+(y+1)/n, (y+1)%n)
			if b {
				return a, b
			}
			continue
		}
		x2 := i.tetramino[1][0] + x
		y2 := i.tetramino[1][1] + y
		if x2 < 0 || x2 >= n || y2 < 0 || y2 >= n || sq[x2][y2] != "." {
			a, b := decidedTetris(alltetramino, sq, x+(y+1)/n, (y+1)%n)
			if b {
				return a, b
			}
			continue
		}
		x3 := i.tetramino[2][0] + x
		y3 := i.tetramino[2][1] + y
		if x3 < 0 || x3 >= n || y3 < 0 || y3 >= n || sq[x3][y3] != "." {
			a, b := decidedTetris(alltetramino, sq, x+(y+1)/n, (y+1)%n)
			if b {
				return a, b
			}
			continue
		}
		x4 := i.tetramino[3][0] + x
		y4 := i.tetramino[3][1] + y
		if x4 < 0 || x4 >= n || y4 < 0 || y4 >= n || sq[x4][y4] != "." {
			a, b := decidedTetris(alltetramino, sq, x+(y+1)/n, (y+1)%n)
			if b {
				return a, b
			}
			continue
		}

		sq[x1][y1] = i.symbol
		sq[x2][y2] = i.symbol
		sq[x3][y3] = i.symbol
		sq[x4][y4] = i.symbol
		a := alltetramino[:j]
		a = append(a, alltetramino[j+1:]...)
		if len(a) == 0 {
			return sq, true
		}
		v, b := decidedTetris(a, sq, x+(y+1)/n, (y+1)%n)
		if b {
			return v, b
		}
	}
	return sq, false
}

func splitTetramino(fileStr string) []tetraminoStr {
	var alltetramino [][]string
	lineList := strings.Split(fileStr, "\n")

	var tetramino []string

	for _, i := range lineList {
		if len(i) == 0 {
			alltetramino = append(alltetramino, tetramino)
			tetramino = []string{}
			continue
		}
		tetramino = append(tetramino, i)
	}
	return validateTetramino(alltetramino)
}

func genaretSQ(n int) [][]string {
	var sq [][]string
	k := 0
	for k != n {
		j := 0
		line := []string{}
		for j != n {
			line = append(line, ".")
			j++
		}
		sq = append(sq, line)
		k++
	}
	return sq
}

//TODO unconnected squares
func validateTetramino(alltetramino [][]string) []tetraminoStr {
	var alltetr []tetraminoStr
	alph := 'A'

	for _, i := range alltetramino {
		var tetr tetraminoStr
		a := 0
		defx := 100
		defy := 100

		if len(i) != 4 {
			tetrisERROR(errors.New("invalid value"))
		}
		for x, j := range i {

			if len(j) != 4 {
				tetrisERROR(errors.New("invalid value"))
			}
			for y, n := range j {
				if n == '#' {
					var coord [2]int
					if defx > 4 {
						defx = x
						defy = y
					}
					coord[0] = defx - x
					coord[1] = defy - y

					tetr.tetramino = append(tetr.tetramino, coord)
					a++
					continue
				}
			}
		}
		tetr.symbol = string(alph)
		alph = rune(byte(alph) + 1)
		if a != 4 {
			tetrisERROR(errors.New("invalid value. A tetromino is a geometric shape composed of four squares"))
		}
		alltetr = append(alltetr, tetr)

	}
	return alltetr
}

func tetrisERROR(err error) {
	fmt.Println("ERROR: ", err)
}
