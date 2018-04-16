package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

var Inf = 1000000000
var saved [20000][10]int

func getAns(left int, coins []int) int {
	if left == 0 {
		return 0
	}
	if len(coins) == 0 {
		return Inf
	}

	if saved[left][len(coins)] == 0 {

		res := Inf
		for i := 0; i*coins[0] <= left; i++ {
			ans := getAns(left-i*coins[0], coins[1:]) + i
			if ans < res {
				res = ans
			}
		}
		saved[left][len(coins)] = res
	}
	return saved[left][len(coins)]
}


func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	coins := make([]int, 0)
	for _, s := range strings.Split(scanner.Text(), ",") {
		i, _ := strconv.Atoi(s)
		coins = append(coins, i)
	}
	fmt.Println(getAns(n, coins))
}

