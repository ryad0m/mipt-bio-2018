package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

var REPLACE = map[rune]map[rune]int{
	'A': {'A': 2, 'C': -2, 'D': 0, 'E': 0, 'F': -3, 'G': 1, 'H': -1, 'I': -1, 'K': -1, 'L': -2, 'M': -1, 'N': 0, 'P': 1, 'Q': 0, 'R': -2, 'S': 1, 'T': 1, 'V': 0, 'W': -6, 'Y': -3, '-': -5},
	'C': {'A': -2, 'C': 12, 'D': -5, 'E': -5, 'F': -4, 'G': -3, 'H': -3, 'I': -2, 'K': -5, 'L': -6, 'M': -5, 'N': -4, 'P': -3, 'Q': -5, 'R': -4, 'S': 0, 'T': -2, 'V': -2, 'W': -8, 'Y': 0, '-': -5},
	'D': {'A': 0, 'C': -5, 'D': 4, 'E': 3, 'F': -6, 'G': 1, 'H': 1, 'I': -2, 'K': 0, 'L': -4, 'M': -3, 'N': 2, 'P': -1, 'Q': 2, 'R': -1, 'S': 0, 'T': 0, 'V': -2, 'W': -7, 'Y': -4, '-': -5},
	'E': {'A': 0, 'C': -5, 'D': 3, 'E': 4, 'F': -5, 'G': 0, 'H': 1, 'I': -2, 'K': 0, 'L': -3, 'M': -2, 'N': 1, 'P': -1, 'Q': 2, 'R': -1, 'S': 0, 'T': 0, 'V': -2, 'W': -7, 'Y': -4, '-': -5},
	'F': {'A': -3, 'C': -4, 'D': -6, 'E': -5, 'F': 9, 'G': -5, 'H': -2, 'I': 1, 'K': -5, 'L': 2, 'M': 0, 'N': -3, 'P': -5, 'Q': -5, 'R': -4, 'S': -3, 'T': -3, 'V': -1, 'W': 0, 'Y': 7, '-': -5},
	'G': {'A': 1, 'C': -3, 'D': 1, 'E': 0, 'F': -5, 'G': 5, 'H': -2, 'I': -3, 'K': -2, 'L': -4, 'M': -3, 'N': 0, 'P': 0, 'Q': -1, 'R': -3, 'S': 1, 'T': 0, 'V': -1, 'W': -7, 'Y': -5, '-': -5},
	'H': {'A': -1, 'C': -3, 'D': 1, 'E': 1, 'F': -2, 'G': -2, 'H': 6, 'I': -2, 'K': 0, 'L': -2, 'M': -2, 'N': 2, 'P': 0, 'Q': 3, 'R': 2, 'S': -1, 'T': -1, 'V': -2, 'W': -3, 'Y': 0, '-': -5},
	'I': {'A': -1, 'C': -2, 'D': -2, 'E': -2, 'F': 1, 'G': -3, 'H': -2, 'I': 5, 'K': -2, 'L': 2, 'M': 2, 'N': -2, 'P': -2, 'Q': -2, 'R': -2, 'S': -1, 'T': 0, 'V': 4, 'W': -5, 'Y': -1, '-': -5},
	'K': {'A': -1, 'C': -5, 'D': 0, 'E': 0, 'F': -5, 'G': -2, 'H': 0, 'I': -2, 'K': 5, 'L': -3, 'M': 0, 'N': 1, 'P': -1, 'Q': 1, 'R': 3, 'S': 0, 'T': 0, 'V': -2, 'W': -3, 'Y': -4, '-': -5},
	'L': {'A': -2, 'C': -6, 'D': -4, 'E': -3, 'F': 2, 'G': -4, 'H': -2, 'I': 2, 'K': -3, 'L': 6, 'M': 4, 'N': -3, 'P': -3, 'Q': -2, 'R': -3, 'S': -3, 'T': -2, 'V': 2, 'W': -2, 'Y': -1, '-': -5},
	'M': {'A': -1, 'C': -5, 'D': -3, 'E': -2, 'F': 0, 'G': -3, 'H': -2, 'I': 2, 'K': 0, 'L': 4, 'M': 6, 'N': -2, 'P': -2, 'Q': -1, 'R': 0, 'S': -2, 'T': -1, 'V': 2, 'W': -4, 'Y': -2, '-': -5},
	'N': {'A': 0, 'C': -4, 'D': 2, 'E': 1, 'F': -3, 'G': 0, 'H': 2, 'I': -2, 'K': 1, 'L': -3, 'M': -2, 'N': 2, 'P': 0, 'Q': 1, 'R': 0, 'S': 1, 'T': 0, 'V': -2, 'W': -4, 'Y': -2, '-': -5},
	'P': {'A': 1, 'C': -3, 'D': -1, 'E': -1, 'F': -5, 'G': 0, 'H': 0, 'I': -2, 'K': -1, 'L': -3, 'M': -2, 'N': 0, 'P': 6, 'Q': 0, 'R': 0, 'S': 1, 'T': 0, 'V': -1, 'W': -6, 'Y': -5, '-': -5},
	'Q': {'A': 0, 'C': -5, 'D': 2, 'E': 2, 'F': -5, 'G': -1, 'H': 3, 'I': -2, 'K': 1, 'L': -2, 'M': -1, 'N': 1, 'P': 0, 'Q': 4, 'R': 1, 'S': -1, 'T': -1, 'V': -2, 'W': -5, 'Y': -4, '-': -5},
	'R': {'A': -2, 'C': -4, 'D': -1, 'E': -1, 'F': -4, 'G': -3, 'H': 2, 'I': -2, 'K': 3, 'L': -3, 'M': 0, 'N': 0, 'P': 0, 'Q': 1, 'R': 6, 'S': 0, 'T': -1, 'V': -2, 'W': 2, 'Y': -4, '-': -5},
	'S': {'A': 1, 'C': 0, 'D': 0, 'E': 0, 'F': -3, 'G': 1, 'H': -1, 'I': -1, 'K': 0, 'L': -3, 'M': -2, 'N': 1, 'P': 1, 'Q': -1, 'R': 0, 'S': 2, 'T': 1, 'V': -1, 'W': -2, 'Y': -3, '-': -5},
	'T': {'A': 1, 'C': -2, 'D': 0, 'E': 0, 'F': -3, 'G': 0, 'H': -1, 'I': 0, 'K': 0, 'L': -2, 'M': -1, 'N': 0, 'P': 0, 'Q': -1, 'R': -1, 'S': 1, 'T': 3, 'V': 0, 'W': -5, 'Y': -3, '-': -5},
	'V': {'A': 0, 'C': -2, 'D': -2, 'E': -2, 'F': -1, 'G': -1, 'H': -2, 'I': 4, 'K': -2, 'L': 2, 'M': 2, 'N': -2, 'P': -1, 'Q': -2, 'R': -2, 'S': -1, 'T': 0, 'V': 4, 'W': -6, 'Y': -2, '-': -5},
	'W': {'A': -6, 'C': -8, 'D': -7, 'E': -7, 'F': 0, 'G': -7, 'H': -3, 'I': -5, 'K': -3, 'L': -2, 'M': -4, 'N': -4, 'P': -6, 'Q': -5, 'R': 2, 'S': -2, 'T': -5, 'V': -6, 'W': 17, 'Y': 0, '-': -5},
	'Y': {'A': -3, 'C': 0, 'D': -4, 'E': -4, 'F': 7, 'G': -5, 'H': 0, 'I': -1, 'K': -4, 'L': -1, 'M': -2, 'N': -2, 'P': -5, 'Q': -4, 'R': -4, 'S': -3, 'T': -3, 'V': -2, 'W': 0, 'Y': 10, '-': -5},
	'-': {'A': -5, 'C': -5, 'D': -5, 'E': -5, 'F': -5, 'G': -5, 'H': -5, 'I': -5, 'K': -5, 'L': -5, 'M': -5, 'N': -5, 'P': -5, 'Q': -5, 'R': -5, 'S': -5, 'T': -5, 'V': -5, 'W': -5, 'Y': -5, '-': -100000}}

type State struct {
	score int
	px    int
	py    int
	rx    rune
	ry    rune
}

func (s *State) Update(nscore, npx, npy int, nrx, nry rune) {
	if s.score < nscore {
		s.score = nscore
		s.px = npx
		s.py = npy
		s.rx = nrx
		s.ry = nry
	}
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()

	dp := make([][]State, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]State, len(b)+1)
		for j := 0; j <= len(b); j++ {
			dp[i][j].score = -1e9
		}
	}
	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if i > 0 {
				dp[i][j].Update(dp[i-1][j].score+REPLACE[rune(a[i-1])]['-'], i-1, j, rune(a[i-1]), '-')
			}
			if j > 0 {
				dp[i][j].Update(dp[i][j-1].score+REPLACE['-'][rune(b[j-1])], i, j-1, '-', rune(b[j-1]))
			}
			if i > 0 && j > 0 {
				dp[i][j].Update(dp[i-1][j-1].score+REPLACE[rune(a[i-1])][rune(b[j-1])], i-1, j-1, rune(a[i-1]), rune(b[j-1]))
			}
			dp[i][j].Update(0, -1, -1, 0, 0)
		}
	}
	posx, posy := len(a), len(b)
	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if dp[i][j].score > dp[posx][posy].score {
				posx = i
				posy = j
			}
		}
	}
	resa, resb := "", ""
	res := dp[posx][posy].score
	for (posx > 0 || posy > 0) && dp[posx][posy].score >= 0 {
		if dp[posx][posy].rx > 0 && dp[posx][posy].ry > 0 {
			resa = string(dp[posx][posy].rx) + resa
			resb = string(dp[posx][posy].ry) + resb
		}
		posx, posy = dp[posx][posy].px, dp[posx][posy].py
	}
	fout, _ := os.Create("output.txt")
	fout.WriteString(strconv.Itoa(res))
	fout.WriteString("\n")
	fout.WriteString(resa)
	fout.WriteString("\n")
	fout.WriteString(resb)
	fout.WriteString("\n")
	fout.Close()
	fmt.Println(res)
	fmt.Println(resa)
	fmt.Println(resb)
}

