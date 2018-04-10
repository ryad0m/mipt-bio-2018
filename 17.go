package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)


var Amins = []rune{'G', 'A', 'S', 'P', 'V', 'T', 'C', 'L', 'N', 'D', 'Q', 'E', 'M', 'H', 'F', 'R', 'Y', 'W'}
var Masses = []int64{57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186}


func factorial(x int64) int64 {
	if x <= 1 {
		return 1
	}
	return x * factorial(x - 1)
}

func recurse(left int64, cur int, res string) int64 {
	if cur == len(Amins) && left > 0 || left < 0 {
		return 0
	}
	if left == 0 {
		//fmt.Println(len(res))
		ans := factorial(int64(len(res)))
		cnt := make(map[rune]int)
		for _, ch := range res {
			cnt[ch] += 1
		}
		for _, value := range cnt {
			ans /= factorial(int64(value))
		}
		return ans
	}
    return recurse(left - Masses[cur], cur, res + string(Amins[cur])) + recurse(left, cur + 1, res)
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Println(recurse(int64(n), 0, ""))
}

