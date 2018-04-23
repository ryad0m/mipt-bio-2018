package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func Print(a []int) {
	fmt.Print("(")
	for i, x := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		if x > 0 {
			fmt.Print("+")
		}
		fmt.Print(x)
	}
	fmt.Println(")")
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	a := make([]int, 0)
	for _, x := range strings.Split(strings.Replace(scanner.Text()[1:len(scanner.Text())-1], "+", "", -1), " ") {
		x, _ := strconv.Atoi(x)
		a = append(a, x)
	}
	for k := range a {
		i := 0
		for ; i < len(a); i++ {
			if a[i] == k+1 || -a[i] == k+1 {
				break
			}
		}
		if i == k && a[i] > 0 {
			continue
		}
		if i < k {
			reverse(a[i : k+1])
			for j := i; j <= k; j++ {
				a[j] = -a[j]
			}
		} else {
			reverse(a[k : i+1])
			for j := k; j <= i; j++ {
				a[j] = -a[j]
			}
		}
		Print(a)
		if a[k] < 0 {
			a[k] = -a[k]
			Print(a)
		}
	}
}

