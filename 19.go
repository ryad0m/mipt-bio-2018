package main

import (
	"sort"
	"strconv"
	"os"
	"bufio"
	"strings"
	"fmt"
)

var AminList = []rune{'G', 'A', 'S', 'P', 'V', 'T', 'C', 'L', 'N', 'D', 'Q', 'E', 'M', 'H', 'F', 'R', 'Y', 'W'}
var MassList = []int{57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186}
var Amin2Mass map[rune]int
var Best = ""

func Init() {
	Amin2Mass = make(map[rune]int)
	for i, ch := range AminList {
		Amin2Mass[ch] = MassList[i]
	}
}

func countMassString(peptide string) int {
	sum := 0
	for _, acid := range peptide {
		sum += Amin2Mass[acid]
	}
	return sum
}


func cyclospectrum(peptide string) []int {
	spectrum := make([]int, 1)
	prefix_sum := make([]int, 1)
	for _, amin := range peptide {
		prefix_sum = append(prefix_sum, prefix_sum[len(prefix_sum) - 1] + Amin2Mass[amin])
	}
	n := len(peptide)
	for i := range peptide {
		for j := i + 1; j < n + 1; j++ {
			spectrum = append(spectrum, prefix_sum[j] - prefix_sum[i])
			if i > 0 && j < n {
				spectrum = append(spectrum, prefix_sum[n] - prefix_sum[j] + prefix_sum[i])
			}
		}
	}
	sort.Sort(sort.IntSlice(spectrum))
	return spectrum
}


func extend(peptids []string) []string {
	result := make([]string, 0)
	for _, p := range peptids {
		for _, amin := range AminList {
			result = append(result, p + string(amin))
		}
	}
	return result
}

func Counter(spectrum []int) map[int]int {
	res := make(map[int]int)
	for _, i := range spectrum {
		res[i]++
	}
	return res
}

func score(givspectrum map[int]int, peptide string) int {
	spectrum := Counter(cyclospectrum(peptide))
	score := 0
	for amin, cnt := range spectrum {
		if cnt < givspectrum[amin] {
			score += cnt
		} else {
			score += givspectrum[amin]
		}
	}
	return score
}

var givspectrum map[int]int

type byScore []string

func (s byScore) Len() int {
	return len(s)
}
func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byScore) Less(i, j int) bool {
	return score(givspectrum, s[i]) > score(givspectrum, s[j])
}

func cutLeaderBoard(parentMass int, peptids []string, n int) []string {
	best := make([]string, 0)
	for _, p := range peptids {
		if countMassString(p) <= parentMass {
			best = append(best, p)
			if countMassString(p) == parentMass && score(givspectrum, Best) < score(givspectrum, p) {
				Best = p
			}
		}
	}
	sort.Sort(byScore(best))
	if len(best) <= n {
		return best
	}
	s := score(givspectrum, best[n - 1])
	for n < len(best) && score(givspectrum, best[n]) == s {
		n++
	}
	return best[:n]
}

func aminToStr(s string) string {
	res := ""
	for i, c := range s {
		if i > 0 {
			res += "-"
		}
		res += strconv.Itoa(Amin2Mass[c])
	}
	return res
}

func main() {
	Init()
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	givspectrum_list := make([]int, 0)
	parent_mass := 0
	for _, item := range strings.Split(scanner.Text(), " ") {
		i, _ := strconv.Atoi(item)
		givspectrum_list = append(givspectrum_list, i)
		if i > parent_mass {
			parent_mass = i
		}
	}
	givspectrum = Counter(givspectrum_list)
	peptids := make([]string, 1)
	for len(peptids) > 0 {
		peptids = cutLeaderBoard(parent_mass, extend(peptids), n)
	}
	fmt.Println(aminToStr(Best))
}

