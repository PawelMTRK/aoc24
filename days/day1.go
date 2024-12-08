package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1() {
	res, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("ReadFile: ", err)
	}
	stringRes := string(res)
	sliceNums := strings.Fields(stringRes)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for i, e := range sliceNums {
		n, _ := strconv.Atoi(e)
		if i%2 == 0 {
			list1 = append(list1, n)
		} else {
			list2 = append(list2, n)
		}
	}
	fmt.Println(list1, list2)
}
