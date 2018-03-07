package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func dist(a, b string) int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}

func minDist(s, t string) int {
	res := len(t)
	for i := 0; i <= len(s) - len(t); i++ {
		cur := dist(t, s[i:i + len(t)])
		if cur < res {
			res = cur
		}
	}
	return res
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	list := strings.Split(scanner.Text(), " ")
	sum := 0
	for _, s := range list {
		sum += minDist(s, t)
	}
	fmt.Println(sum)
}

