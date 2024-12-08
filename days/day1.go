package days

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HistorianList struct {
	list []int
}

func (l *HistorianList) Append(v string) {
	n, _ := strconv.Atoi(v)
	l.list = append(l.list, n)
}

func (l *HistorianList) Load(slice []string) {
	l.list = make([]int, 0)
	for _, v := range slice {
		n, _ := strconv.Atoi(v)
		l.list = append(l.list, n)
	}
}

func (l *HistorianList) GetDifference(l2 *HistorianList) int {
	slices.Sort(l.list)
	slices.Sort(l2.list)
	diff := 0
	for i, _ := range l.list {
		diff += int(math.Abs(float64(l.list[i] - l2.list[i])))
	}
	return diff
}

func Day1() {
	res, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("ReadFile: ", err)
	}
	stringRes := string(res)
	sliceNums := strings.Fields(stringRes)

	list1 := HistorianList{}
	list2 := HistorianList{}

	for i, e := range sliceNums {
		if i%2 == 0 {
			list1.Append(e)
		} else {
			list2.Append(e)
		}
	}

	diff := list1.GetDifference(&list2)
	fmt.Println("The difference is:", diff)
}
