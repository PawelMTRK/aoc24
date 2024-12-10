package days

import (
	"fmt"
	"os"
	"strings"
)

type WordSearch struct {
	puzzle []string
}

func NewWordSearch(s string) WordSearch {
	lines := strings.Split(s, "\n")
	return WordSearch{puzzle: lines}
}

func (w *WordSearch) isXMASorSAMX(x, m, a, s byte) bool {
	if x == 88 && m == 77 && a == 65 && s == 83 {
		return true
	} else if x == 83 && m == 65 && a == 77 && s == 88 {
		return true
	}
	return false
}

func (w *WordSearch) FindByDirection(dir string) int {
	total := 0

	switch dir {
	case "-":
		for r := range len(w.puzzle) {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r][c+1], w.puzzle[r][c+2], w.puzzle[r][c+3]) {
					total++
				}
			}
		}
	case "|":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c], w.puzzle[r+2][c], w.puzzle[r+3][c]) {
					fmt.Println("|XMAS at", r, c)
					total++
				}
			}
		}
	case "\\":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c+1], w.puzzle[r+2][c+2], w.puzzle[r+3][c+3]) {
					fmt.Println("\\XMAS at", r, c)
					total++
				}
			}
		}
	case "/":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c+3], w.puzzle[r+1][c+2], w.puzzle[r+2][c+1], w.puzzle[r+3][c]) {
					fmt.Println("/XMAS at", r, c+3)
					total++
				}
			}
		}

	}
	return total
}

func Day4() {
	res, _ := os.ReadFile("input4.txt")
	stringRes := string(res)
	wordSearch := NewWordSearch(stringRes)
	fmt.Println(wordSearch.puzzle[0])
	total := wordSearch.FindByDirection("-") + wordSearch.FindByDirection("|") + wordSearch.FindByDirection("\\") + wordSearch.FindByDirection("/")
	fmt.Println(total)
}
