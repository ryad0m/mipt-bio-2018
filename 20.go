package main

import (
	"sort"
	"strconv"
	"os"
	"bufio"
	"strings"
	"fmt"
)


func iabs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Counter(spectrum []int) map[int]int {
	res := make(map[int]int)
	for _, i := range spectrum {
		res[i]++
	}
	return res
}

func most_freq_conv(spectrum []int, m int) []int {
	res := make([]int, 0)
	for i := range spectrum {
		for j := i + 1; j < len(spectrum); j++ {
			res = append(res, iabs(spectrum[i] - spectrum[j]))
		}
	}
	nres := make([]int, 0)
	for _, x := range res {
		if 57 <= x && x <= 200 {
			nres = append(nres, x)
		}
	}
	cnts := Counter(nres)
	res = make([]int, 0)
	for k := range cnts {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool { return cnts[res[i]] > cnts[res[j]] })

	for len(res) > m && cnts[res[len(res) - 1]] < cnts[res[m - 1]] {
		res = res[:len(res) - 1]
	}
	return res
}

func score(givspectrum map[int]int, peptide []int) int {
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

func cyclospectrum(peptide []int) []int {
	spectrum := make([]int, 1)
	prefix_sum := make([]int, 1)
	for _, amin := range peptide {
		prefix_sum = append(prefix_sum, prefix_sum[len(prefix_sum) - 1] + amin)
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
	sort.Ints(spectrum)
	return spectrum
}

var Best = make([]int, 0)

func getBest(givspectrum map[int]int, par_mass int, peptids [][]int, n int) [][]int {
	best := make([][]int, 0)
	for _, p := range peptids {
		p_mass := 0
		for _, x := range p {
			p_mass += x
		}
		if p_mass <= par_mass {
			best = append(best, p)
			if p_mass == par_mass && score(givspectrum, p) > score(givspectrum, Best) {
				Best = p
			}
		}
	}

	sort.Slice(best, func (i, j int) bool { return score(givspectrum, best[i]) > score(givspectrum, best[j]) })

	if len(best) <= n {
		return best
	}
	s := score(givspectrum, best[n - 1])
	for n < len(best) && score(givspectrum, best[n]) == s {
		n++
	}
	return best[:n]
}

func extend(peptids [][]int, Acids []int) [][]int {
	result := make([][]int, 0)
	for _, p := range peptids {
		for _, acid := range Acids {
			res := append(make([]int, 0, len(p)), p...)
			result = append(result, append(res, acid))
		}
	}
	return result
}

func aminToStr(s []int) string {
	res := ""
	for i, c := range s {
		if i > 0 {
			res += "-"
		}
		res += strconv.Itoa(c)
	}
	return res
}

func printDebug(givspectrum map[int]int, peptides [][]int) {
	for _, p := range peptides {
		fmt.Print(score(givspectrum, p), " ")
	}
	fmt.Println()
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	m--
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
	givspectrum := Counter(givspectrum_list)
	Acids := most_freq_conv(givspectrum_list, m)
	peptids := make([][]int, 0)
	for _, x := range Acids {
		peptids = append(peptids, append(make([]int, 0), x))
	}
	for len(peptids) > 0 {
		peptids = getBest(givspectrum, parent_mass, extend(peptids, Acids), n)
	}
	fmt.Println(aminToStr(Best))
}

