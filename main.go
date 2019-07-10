package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseStdin(args []string) [][]int {
	result := make([][]int, len(args))
	for i, arg := range args {
		s := strings.Split(arg, "")
		inner := make([]int, len(s))
		for j, str := range s {
			inty, _ := strconv.Atoi(str)
			inner[j] = inty
		}
		result[i] = inner
	}
	return result
}

func main() {
	args := os.Args[1:]
	stdin := parseStdin(args)

	result := SelectionSort(stdin[0], true)
	// result = InsertionSort(MakeRandIntArray(20, 300), true)
	// binsumresult, err := BinarySum(stdin[0], stdin[1])
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(result)
	// fmt.Println(binsumresult)
}
