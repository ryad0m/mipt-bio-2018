package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

type Counter map[byte]int
type Profile []map[byte]float64


func countChars(motifs []string) ([]Counter) {
	res := make([]Counter, len(motifs[0]))
	for i := 0; i < len(motifs[0]); i++ {
		res[i] = make(Counter)
		for j := 0; j < len(motifs); j++ {
			res[i][motifs[j][i]] += 1
		}
	}
	return res
}

func (m Counter) MaxValue() (int) {
	val := -1
	for _, value := range m {
		if value > val {
			val = value
		}
	}
	return val
}

func score(motifs []string) (score int) {
	for _, counts := range countChars(motifs) {
		score += len(motifs) - counts.MaxValue()
	}
	return
}

func getProfile(motifs []string) (Profile) {
	res := make(Profile, len(motifs[0]))
	for i, count := range countChars(motifs) {
		res[i] = make(map[byte]float64)
		for _, key := range []byte{'A', 'T', 'G', 'C'} {
			res[i][key] = float64(count[key] + 1) / float64(len(motifs))
		}
	}
	return res
}

func mostProbable(dna string, k int, profile Profile) (string) {
	max_prob := -1.
	index := 0
	for i := 0; i <= len(dna) - k; i++ {
		prob := 1.
		for j := 0; j < k; j++ {
			prob *= profile[j][dna[i + j]]
		}
		if prob > max_prob {
			max_prob = prob
			index = i
		}
	}
	return dna[index: index + k]
}

func getRandom(s string, k int) (string) {
	start := rand.Int() % (len(s) - k + 1)
	return s[start:start + k]
}

func getRandomMotifs(s []string, k int) ([]string) {
	res := make([]string, len(s))
	for i, c := range s {
		res[i] = getRandom(c, k)
	}
	return res
}


func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	k, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	t, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])
	matrix := make([]string, 0)
	for i := 0; i < t; i++ {
		scanner.Scan()
		matrix = append(matrix, scanner.Text())
	}

	cur := getRandomMotifs(matrix, k)
	bestMotifs := cur
	bestScore := score(bestMotifs)
	for it := 0; it < 10000; it++ {
		profile := getProfile(cur)
		cur = make([]string, 0, len(matrix))
		for j := 0; j < t; j++ {
			cur = append(cur, mostProbable(matrix[j], k, profile))
		}
		curScore := score(cur)
		if curScore < bestScore {
			bestMotifs = cur
			bestScore = curScore
		} else {
			cur = getRandomMotifs(matrix, k)
		}
	}

	for _, s := range bestMotifs {
		fmt.Println(s)
	}
}

