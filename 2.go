package main

import (
	"fmt"
	"bufio"
	"os"
)


func getSubStr(s string, sub string) ([]int) {
	res := make([]int, 0)
	for i := 0; i <= len(s) - len(sub); i++ {
		if sub == s[i:i + len(sub)] {
			res = append(res, i)
		}
	}
	return res
}

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

