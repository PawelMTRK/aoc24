package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puter struct {
	mulBlocks []string
	mulArgs   [][2]int
}

func IndexN(s string, substr string, n int) int {
	i, last := 0, 0
	for range n {
		i = strings.Index(s[last:], substr)
		if i == -1 {
			return -1
		}
		last += i + 1
	}
	// last variable contains the last index of the substring + 1
	// so we decrement it
	i = last - 1
	if i < 0 {
		i = 0
	}
	return i
}

func NewPuter(v string) Puter {
	memoryBlocks := strings.Split(v, "mul")
	return Puter{mulBlocks: memoryBlocks[1:]}
}

func (p *Puter) ParseBlocks() {
	for _, v := range p.mulBlocks {
		// fmt.Println("Parsing...", v)
		closeParenI := strings.Index(v, ")")
		if closeParenI == -1 {
			continue
		}
		rawArgs := v[1:closeParenI]
		commaI := strings.Index(rawArgs, ",")
		if commaI == -1 {
			continue
		}
		// fmt.Println(rawArgs[:commaI], rawArgs[commaI+1:])
		n1, _ := strconv.Atoi(rawArgs[:commaI])
		n2, _ := strconv.Atoi(rawArgs[commaI+1:])
		p.mulArgs = append(p.mulArgs, [2]int{n1, n2})
	}
}

func (p *Puter) Calculate() int {
	total := 0
	for _, a := range p.mulArgs {
		// fmt.Println(a[0], "*", a[1])
		total += (a[0] * a[1])
	}
	return total
}

func Day3() {
	res, _ := os.ReadFile("input3.txt")
	stringRes := string(res)
	puter := NewPuter(stringRes)
	puter.ParseBlocks()
	result := puter.Calculate()
	fmt.Println("Result is:", result)
}
