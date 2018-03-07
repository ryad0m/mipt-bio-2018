package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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
		for key, value := range count {
			res[i][key] = float64(value) / float64(len(motifs))
		}
	}
	return res
}

func mostProbable(dna string, k int, profile Profile) (string) {
	max_prob := 0.
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

	bestMotifs := make([]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		bestMotifs[i] = matrix[i][:k]
	}
	bestScore := score(bestMotifs)

	for i := 0; i <= len(matrix[0]) - k; i++ {
		motifs := make([]string, 1, len(matrix))
		motifs[0] = matrix[0][i:i + k]

		for j := 1; j < t; j++ {
			motifs = append(motifs, mostProbable(matrix[j], k, getProfile(motifs)))
		}
		curScore := score(motifs)
		if curScore < bestScore {
			bestMotifs = motifs
			bestScore = curScore
		}
	}
	for _, s := range bestMotifs {
		fmt.Println(s)
	}
}

