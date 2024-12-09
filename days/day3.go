package days

import (
	"aoc24/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puter struct {
	mulBlocks []string
	mulArgs   [][2]int
}

func NewPuter(s string) Puter {
	muls := util.GetIndexes(s, "mul")
	enabled := true
	memoryBlocks := make([]string, 0)
	var block string
	for i := range muls {
		if i == len(muls)-1 {
			block = s[muls[i]:]
		} else {
			block = s[muls[i]:muls[i+1]]
		}
		if enabled {
			memoryBlocks = append(memoryBlocks, block)
		}
		util.ToggleBool(&enabled, strings.Contains(block, "do()"), strings.Contains(block, "don't()"))
	}
	return Puter{mulBlocks: memoryBlocks}
}

func (p *Puter) ParseBlocks() {
	for _, v := range p.mulBlocks {
		validStart := strings.Contains(v, "mul(")
		if !validStart {
			continue
		}
		closeParenI := strings.Index(v, ")")
		if closeParenI == -1 {
			continue
		}
		// start before "mul("
		rawArgs := v[4:closeParenI]
		commaI := strings.Index(rawArgs, ",")
		if commaI == -1 {
			continue
		}
		n1, _ := strconv.Atoi(rawArgs[:commaI])
		n2, _ := strconv.Atoi(rawArgs[commaI+1:])
		p.mulArgs = append(p.mulArgs, [2]int{n1, n2})
	}
}

func (p *Puter) Calculate() int {
	total := 0
	for _, a := range p.mulArgs {
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
