package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parse(s string) [][]int {
	res := make([][]int, 0)
	s = strings.Replace(s, "+", "", -1)
	s = s[:len(s) - 1]
	for _, q := range strings.Split(s, ")") {
		if len(q) > 0 {
			res = append(res, make([]int, 0))
			for _, w := range strings.Split(q[1:], " ") {
				i, _ := strconv.Atoi(w)
				res[len(res)-1] = append(res[len(res)-1], i)
			}
		}
	}
	return res
}

func addEdge(perm []int, graph [][]int) {
	for i := range perm {
		j := (i + 1) % len(perm)
		out := 0
		if perm[i] >= 0 {
			out = perm[i] * 2 - 1
		} else {
			out = -perm[i] * 2 - 2
		}
		in := 0
		if perm[j] >= 0 {
			in = perm[j] * 2 - 2
		} else {
			in = -perm[j] * 2 - 1
		}
		graph[out] = append(graph[out], in)
		graph[in] = append(graph[in], out)
	}
}

func dfs(graph [][]int, cur int, visited []bool) {
	if visited[cur] {
		return
	}
	visited[cur] = true
	for _, to := range graph[cur] {
		dfs(graph, to, visited)
	}
}

func sum(a [][]int) (res int) {
	for _, i := range a {
		res += len(i)
	}
	return
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewReader(fin)
	s, _ := scanner.ReadString('\n')
	p := parse(s)
	s, _ = scanner.ReadString('\n')
	q := parse(s)
	n := sum(p)
	graph := make([][]int, 2 * n)
	for i := 0; i < 2 * n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, i := range p {
		addEdge(i, graph)
	}
	for _, i := range q {
		addEdge(i, graph)
	}
	visited := make([]bool, n * 2)
	res := n
	for i := range graph {
		if !visited[i] {
			dfs(graph, i, visited)
			res--
		}
	}
	fmt.Println(res)
}

