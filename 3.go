package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())
	fmt.Println(t, s, d)
	ans := make([]int, 0)
	for i := 0; i <= len(s) - len(t); i++ {
		pattern := s[i:i + len(t)]
		cnt := 0
		for i := range t {
			if t[i] != pattern[i] {
				cnt++
			}
		}
		if cnt <= d {
			ans = append(ans, i)
		}
	}
	fmt.Println(ans)
}

