package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func count_mismatch(s, t string, d int) (ans int) {
	for i := 0; i <= len(s) - len(t); i++ {
		pattern := s[i:i + len(t)]
		cnt := 0
		for i := range t {
			if t[i] != pattern[i] {
				cnt++
			}
		}
		if cnt <= d {
			ans++
		}
	}
	return
}

func get_all_patterns(n int) ([]string) {
	if n == 0 {
		return []string {""}
	}
	ans := make([]string, 0)
	for _, pattern := range get_all_patterns(n - 1) {
		for _, c := range []string {"A", "T", "G", "C"} {
			ans = append(ans, c + pattern)
		}
	}
	return ans
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	k, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	d, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

	max := 0
	ans := make([]string, 0)

	for _, pattern := range get_all_patterns(k) {
		count := count_mismatch(s, pattern, d)
		if count > max {
			ans = make([]string, 0)
			max = count
		}
		if count == max {
			ans = append(ans, pattern)
		}
	}
	fmt.Println(ans)
}

