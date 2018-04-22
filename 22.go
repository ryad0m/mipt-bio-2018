package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	a := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(a[0])
	m, _ := strconv.Atoi(a[1])
	down := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		down[i] = make([]int, m + 1)
		for j, k := range strings.Split(scanner.Text(), " ") {
			k, _ := strconv.Atoi(k)
			down[i][j] = k
		}
	}
	scanner.Scan()
	right := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		scanner.Scan()
		right[i] = make([]int, m)
		for j, k := range strings.Split(scanner.Text(), " ") {
			k, _ := strconv.Atoi(k)
			right[i][j] = k
		}
	}

	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m + 1)
	}
	dp[0][0] = 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i > 0 {
				dp[i][j] = dp[i - 1][j] + down[i - 1][j]
			}
			if j > 0 && dp[i][j] < dp[i][j - 1] + right[i][j - 1] {
				dp[i][j] = dp[i][j - 1] + right[i][j - 1]
			}
		}
	}
	fmt.Println(dp[n][m])
}

