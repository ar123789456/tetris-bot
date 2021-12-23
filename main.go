package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var allSQ = 7

func main() {
	Tetris()
}

type tetraminoStr struct {
	b         bool
	symbol    string
	tetramino [][2]int
}

func Tetris() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		tetrisERROR(err)
	}
	fileStr := string(file)
	alltetramino := splitTetramino(fileStr)
	finishTetris(alltetramino)

}

func finishTetris(alltetramino []tetraminoStr) {
	b := false
	a := [][]string{}
	for !b {
		fmt.Println("Hello")
		a = genaretSQ(allSQ)
		a, b = decidedTetris(alltetramino, a, 0, 0)
		allSQ++
	}
	PrintTetramino(a)
}

//TODO Write this shit
func decidedTetris(alltetramino []tetraminoStr, squ [][]string, x, y int) ([][]string, bool) {
	// fmt.Println(x, y)
	n := len(squ)

	if checkerListTetramino(alltetramino) {
		return squ, true
	}

	if x >= n {
		return squ, false
	}

	if squ[x][y] != "." {
		return decidedTetris(alltetramino, squ, x+((y+1)/n), (y+1)%n)
	}

	for j, v := range alltetramino {
		if v.b {
			continue
		}
		b, sq := checkerTetramino(v, squ, x, y)
		if b {
			a := []tetraminoStr{}
			for _, o := range alltetramino {
				a = append(a, o)
			}
			a[j].b = true
			sq, b = decidedTetris(a, sq, x+((y+1)/n), (y+1)%n)
			if b {
				return sq, true
			}
		}
	}
	return decidedTetris(alltetramino, squ, x+((y+1)/n), (y+1)%n)
}

func checkerTetramino(tetramino tetraminoStr, squ [][]string, x, y int) (bool, [][]string) {
	sq := [][]string{}
	for i, j := range squ {
		sq = append(sq, []string{})
		for _, f := range j {
			sq[i] = append(sq[i], f)
		}
	}

	for _, v := range tetramino.tetramino {
		x1 := v[0] + x
		y1 := v[1] + y
		if x1 < 0 || x1 >= len(squ) {
			return false, squ
		}
		if y1 < 0 || y1 >= len(squ) {
			return false, squ
		}
		if squ[x1][y1] != "." {
			return false, squ
		}
	}
	for _, v := range tetramino.tetramino {
		x1 := v[0] + x
		y1 := v[1] + y
		sq[x1][y1] = tetramino.symbol
	}
	return true, sq
}

func checkerListTetramino(alltetramino []tetraminoStr) bool {
	for _, v := range alltetramino {
		if !v.b {
			return false
		}
	}
	return true
}

func PrintTetramino(squ [][]string) {
	for _, i := range squ {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println("")
	}
}

func splitTetramino(fileStr string) []tetraminoStr {
	var alltetramino [][]string
	lineList := strings.Split(fileStr, "\r\n")

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
			tetrisERROR(errors.New("invalid value x != 4"))
		}
		for x, j := range i {

			if len(j) != 4 {
				tetrisERROR(errors.New("invalid value y != 4"))
			}
			for y, n := range j {
				if n == '#' {
					var coord [2]int
					if defx > 4 {
						defx = x
						defy = y
					}
					coord[0] = (defx - x) * -1
					coord[1] = (defy - y) * -1

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
	panic(1)
}
