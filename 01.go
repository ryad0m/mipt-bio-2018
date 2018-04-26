package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(arr[0])
	l, _ := strconv.Atoi(arr[1])
	t, _ := strconv.Atoi(arr[2])

	set := make(map[string]bool)

	for i := 0; i <= len(s) - k; i++ {
		indexes := getSubStr(s, s[i:i + k])
		start := 0
		good := false
		for end, index := range indexes {
			for index + k - indexes[start] > l {
				start++
			}
			//fmt.Println(start, end)
			if end - start + 1 >= t {
				good = true
			}
		}
		if good {
			set[s[i:i + k]] = true
		}
	}
	for key := range set {
		fmt.Print(key, " ")
	}
}

