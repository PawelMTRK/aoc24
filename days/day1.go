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
	for i := range l.list {
		diff += int(math.Abs(float64(l.list[i] - l2.list[i])))
	}
	return diff
}

func (l *HistorianList) Count(needle int) int {
	n := 0
	for _, v := range l.list {
		if v == needle {
			n++
		}
	}
	return n
}

func (l *HistorianList) GetSimiliarity(l2 *HistorianList) int {
	slices.Sort(l.list)
	slices.Sort(l2.list)
	sim := 0
	for _, v := range l.list {
		amount := l2.Count(v)
		sim += v * amount
	}
	return sim
}

func Day1() {
	res, _ := os.ReadFile("input.txt")
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
	sim := list1.GetSimiliarity(&list2)
	fmt.Println("The difference is:", diff)
	fmt.Println("The similiarity is:", sim)
}
