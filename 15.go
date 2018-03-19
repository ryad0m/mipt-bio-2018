package main

import (
	"bufio"
	"os"
	"fmt"
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
	a := ""
	b := ""

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "|")
		if a == "" {
			a += strs[0]
			b += strs[1]
		} else {
			a += strs[0][len(strs[0]) - 1:]
			b += strs[1][len(strs[1]) - 1:]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i:] == b[:len(a) - i] {
			fmt.Print(a[:i])
			fmt.Print(b)
			return
		}
	}
}

