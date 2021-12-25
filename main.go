package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var allSQ = 2
var tetraminoBase []tetraminoStr

func main() {
	initBase()
	Tetris()
}

func initBase() {
	file, err := os.ReadFile("base.txt")
	if err != nil {
		tetrisERROR(errors.New("500?"))
	}
	fileStr := string(file)
	tetraminoBase = splitTetramino(fileStr)
}

type tetraminoStr struct {
	symbol    string
	tetramino [][2]int
}

func Tetris() {
	args := os.Args[1:]
	if len(args) != 1 {
		tetrisERROR(errors.New("invalid form."))
	}
	file, err := os.ReadFile(args[0])
	if err != nil {
		tetrisERROR(err)
	}
	fileStr := string(file)

	alltetramino := splitTetramino(fileStr)
	if checkupTetramino(alltetramino) {
		finishTetris(alltetramino)
		return
	}
	tetrisERROR(errors.New("bad example"))
}

func checkupTetramino(alltetramino []tetraminoStr) bool {
	n := 0
	m := 0
	// fmt.Println(alltetramino)
	// fmt.Println(tetraminoBase)
	for _, i := range alltetramino {
		for _, j := range tetraminoBase {
			// fmt.Println(i.tetramino, j.)

			for k := range j.tetramino {
				// m := 0
				if i.tetramino[k] == j.tetramino[k] {
					m++
				}
				// fmt.Print(m, " ")
			}
			if m == 4 {
				n++
			}
			m = 0
		}
	}
	// fmt.Println(n, len(alltetramino))
	if n == len(alltetramino) {
		return true
	}

	return false
}

func finishTetris(alltetramino []tetraminoStr) {

	b := false
	a := [][]string{}
	for !b {
		if allSQ*allSQ < len(alltetramino)*4 {
			allSQ++
			continue
		}
		a = genaretSQ(allSQ)
		a, b = decidedTetris(alltetramino, a, 0, 0, allSQ*allSQ)
		allSQ++
	}
	PrintTetramino(a)
}

func decidedTetris(alltetramino []tetraminoStr, squ [][]string, x, y int, dot int) ([][]string, bool) {
	// fmt.Println(path)
	n := len(squ)

	if len(alltetramino) == 0 {
		return squ, true
	}
	if dot < len(alltetramino)*4 {
		return squ, false
	}
	if x >= n {
		return squ, false
	}

	if squ[x][y] != "." {
		return decidedTetris(alltetramino, squ, x+((y+1)/n), (y+1)%n, dot)
	}

	for j, v := range alltetramino {

		b, sq := checkerTetramino(v, squ, x, y)
		if b {
			a := PopTetr(alltetramino, j)
			sq, b = decidedTetris(a, sq, x+((y+1)/n), (y+1)%n, dot-4)
			if b {
				return sq, true
			}
		}
	}

	return decidedTetris(alltetramino, squ, x+((y+1)/n), (y+1)%n, dot-1)
}

func PopTetr(alltetramino []tetraminoStr, i int) []tetraminoStr {
	a := []tetraminoStr{}
	for j, v := range alltetramino {
		if j == i {
			continue
		}
		a = append(a, v)
	}
	return a
}

func checkerTetramino(tetramino tetraminoStr, squ [][]string, x, y int) (bool, [][]string) {
	sq := [][]string{}
	for i, j := range squ {
		sq = append(sq, []string{})
		sq[i] = append(sq[i], j...)

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
		sq[x1][y1] = tetramino.symbol

	}

	return true, sq
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
			tetrisERROR(errors.New("invalid value, bad format file"))
		}
		for x, j := range i {

			if len(j) != 4 {
				tetrisERROR(errors.New("invalid value, bad format file"))
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
	fmt.Println("ERROR:", err)
	os.Exit(1)
}
