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
