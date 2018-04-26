package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewReader(fin)
	s, _ := scanner.ReadString('\n')
	ans := make([]int, 1)
	ans[0] = 1
	scew := 0
	min_scew := 10000000
	for pos, c := range s {
		if c == 'G' {
			scew += 1
		} else if c == 'C' {
			scew -= 1
		}
		if scew == min_scew {
			ans = append(ans, pos + 1)
		} else if scew < min_scew {
			min_scew = scew
			ans = make([]int, 1)
			ans[0] = pos + 1
		}
	}
	fmt.Println(ans)
}

