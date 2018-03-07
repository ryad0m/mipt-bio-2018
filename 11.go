package main

import (
	"bufio"
	"os"
	"fmt"
)

func pref(s string) string {
	return s[:len(s) - 1]
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	//k, _ := strconv.Atoi(scanner.Text())
	parts := make([]string, 0)
	for scanner.Scan() {
		parts = append(parts, scanner.Text())
	}
	prefixes := make(map[string]int)
	suffixes := make(map[string]int)
	unused := make(map[string]int)
	for _, s := range parts {
		prefixes[pref(s)]++
		suffixes[s[1:]]++
		unused[s]++
	}
	start := parts[0]
	for _, part := range parts {
		if suffixes[pref(part)] == 0 {
			start = part
		}
	}
	unused[start]--
	prefixes[pref(start)]--
	cont := true
	for cont {
		cont = false
		for part, cnt := range unused {
			if cnt > 0 && pref(part) == start[len(start)-len(part)+1:] && prefixes[part[1:]] > 0 {
				unused[part]--
				prefixes[pref(part)]--
				start += part[len(part)-1:]
				cont = true
			}
		}
	}
	for part, cnt := range unused {
		if cnt > 0 && pref(part) == start[len(start) - len(part) + 1:] {
			start += part[len(part) - 1:]
			unused[part]--
		}
	}
	sum := 0
	for _, cnt := range unused {
		sum += cnt
	}
	fmt.Println(sum)
	fmt.Println(start)
}

