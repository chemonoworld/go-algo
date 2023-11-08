package main

import (
	"fmt"
	"go-algo/c3_counting_sort"
)

func main() {
	fmt.Println("Hello, World!")
	result := c3_counting_sort.CountingSort([]int{15, 2, 3, 4, 5, 5, 5, 5, 5}, -1, 15)
	fmt.Println(result)

}
