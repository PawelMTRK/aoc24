package days

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type WordSearch struct {
	puzzle []string
	total  int
	m      sync.Mutex
	wg     sync.WaitGroup
}

func NewWordSearch(s string) *WordSearch {
	lines := strings.Split(s, "\n")
	ws := WordSearch{puzzle: lines, total: 0}
	ws.wg.Add(4)
	return &ws
}

func (w *WordSearch) isXMASorSAMX(x, m, a, s byte) bool {
	if x == 88 && m == 77 && a == 65 && s == 83 {
		return true
	} else if x == 83 && m == 65 && a == 77 && s == 88 {
		return true
	}
	return false
}

func (w *WordSearch) FindByDirection(dir string) {
	switch dir {
	case "-":
		for r := range len(w.puzzle) {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r][c+1], w.puzzle[r][c+2], w.puzzle[r][c+3]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "|":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c], w.puzzle[r+2][c], w.puzzle[r+3][c]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "\\":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c+1], w.puzzle[r+2][c+2], w.puzzle[r+3][c+3]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "/":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c+3], w.puzzle[r+1][c+2], w.puzzle[r+2][c+1], w.puzzle[r+3][c]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	}
	w.wg.Done()
}

func Day4() {
	res, _ := os.ReadFile("input4.txt")
	stringRes := string(res)
	wordSearch := NewWordSearch(stringRes)
	go wordSearch.FindByDirection("-")
	go wordSearch.FindByDirection("|")
	go wordSearch.FindByDirection("\\")
	go wordSearch.FindByDirection("/")
	wordSearch.wg.Wait()
	fmt.Println("Total number of XMAS:", wordSearch.total)
}
