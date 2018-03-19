package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)


func match(cur, next string) (bool, string) {
	if cur == "" {
		return true, next
	}
	res := cur[:len(cur) - len(next) + 1]
	for i := 0; i < len(next); i++ {
		if i + 1 == len(next) {
			res += next[len(next) - 1:]
		} else {
			nextc := next[i]
			curc := cur[len(cur) - len(next) + 1 + i]
			if nextc == '?' {
				res += string(curc)
			} else if (curc == '?') {
				res += string(nextc)
			} else if (nextc == curc) {
				res += string(nextc)
			} else {
				return false, ""
			}
		}
	}
	return true, res
}

func trystring(cur string, left map[string]int) (err bool, res string) {
	//fmt.Println(cur)
	found := false
	for key, value := range left {
		if value > 0 {
			found = true
			err, next := match(cur, key)
			if err {
				left[key] -= 1
				has, res := trystring(next, left)
				if has {
					return has, res
				}
				left[key] += 1
			}
		}
	}
	if !found {
		return true, cur
	}
	return false, ""
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(strs[1])

	delim := "";
	for i := 0; i < k; i++ {
		delim += "?"
	}

	parts := make(map[string]int)

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "|")
		piece := strs[0] + delim + strs[1]
		parts[piece] += 1
	}

	fmt.Println(trystring("", parts))
}

