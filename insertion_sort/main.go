package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rep = strings.NewReplacer("[", "", "]", "")

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()

	var arr []int
	cols := strings.Split(line, " ")
	for _, col := range cols {
		v, err := strconv.Atoi(col)
		if err != nil {
			panic(err)
		}
		arr = append(arr, v)
	}

	sort(arr)
}

func sort(nums []int) []int {
	for i := range nums {
		v := nums[i]
		j := 0

		for nums[j] < v {
			j++
		}

		k := i
		for ; k > j; k-- {
			nums[k] = nums[k-1]
		}

		nums[j] = v
		printArray(nums)
	}

	return nums
}

func printArray(nums []int) {
	fmt.Println(rep.Replace(fmt.Sprint(nums)))
}
