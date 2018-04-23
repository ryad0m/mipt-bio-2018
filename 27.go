package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	a := make([]int, 0)
	for _, x := range strings.Split(strings.Replace(scanner.Text()[1:len(scanner.Text())-1], "+", "", -1), " ") {
		x, _ := strconv.Atoi(x)
		a = append(a, x)
	}
	a = append(a, len(a) + 1)
	ans := 0
	prev := 0
	for _, x := range a {
		if x - prev != 1 {
			ans++
		}
		prev = x
	}
	fmt.Println(ans)
}

