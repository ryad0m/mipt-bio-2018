package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)


func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	kk, _ := strconv.Atoi(scanner.Text())
	k := uint(kk)
	s := ""
	for i := 0; i < kk; i++ {
		s += "1"
	}
	parts := make(map[string]int)
	parts[s]++
	for i := 0; uint(i) < ((1 << k) - k); i++ {
		if parts[s[len(s) - kk + 1:] + "0"] > 0 {
			s += "1"
		} else {
			s += "0"
		}
		parts[s[len(s) - kk:]]++
	}
	fmt.Println(s)
}

