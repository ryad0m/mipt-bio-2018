package main

import (
	"strings"
	"sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var AminTable = []rune{'G', 'A', 'S', 'P', 'V', 'T', 'C', 'L', 'N', 'D', 'Q', 'E', 'M', 'H', 'F', 'R', 'Y', 'W'}
var MassTable = []int{57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186}
var AminMass map[rune]int
var AminString string


func Init() {
	AminMass = make(map[rune]int)
	AminString = ""
	for i, ch := range AminTable {
		AminMass[ch] = MassTable[i]
		AminString += string(ch)
	}
}

func dropUnusedAmins(spectrum []int) {
	NewAminTable := make([]rune, 0)
	NewMassTable := make([]int, 0)
	for i, amin := range AminTable {
		found := false
		for _, spec := range spectrum {
			if spec == MassTable[i] {
				found = true
				break
			}
		}
		if found {
			NewAminTable = append(NewAminTable, amin)
			NewMassTable = append(NewMassTable, MassTable[i])
		}
	}
	AminTable = NewAminTable
	MassTable = NewMassTable
	Init()
}

func countMass(peptide []int) int {
	sum := 0
	for _, acid := range peptide {
		sum += MassTable[acid]
	}
	return sum
}

func countMassString(peptide string) int {
	sum := 0
	for _, acid := range peptide {
		sum += AminMass[acid]
	}
	return sum
}

func extend(peptids []string, parent_mass int) []string {
	result := make([]string, 0)
	for _, p := range peptids {
		ind := 0
		if len(p) > 0 {
			ind = strings.IndexByte(AminString, p[0]) + 1
		}
		p_mass := countMassString(p)
		for _, amin := range AminTable[ind:] {
			if AminMass[amin] + p_mass > parent_mass {
				break
			}
			result = append(result, p + string(amin))
		}
	}
	return result
}


func cyclospectrum(peptide string) []int {
	spectrum := make([]int, 1)
	prefix_sum := make([]int, 1)
	for _, amin := range peptide {
		prefix_sum = append(prefix_sum, prefix_sum[len(prefix_sum) - 1] + AminMass[amin])
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

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func aminToStr(s string) string {
	res := ""
	for i, c := range s {
		if i > 0 {
			res += "-"
		}
		res += strconv.Itoa(AminMass[c])
	}
	return res
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	givspectrum := make([]int, 0)
	for _, item := range strings.Split(scanner.Text(), " ") {
		i, _ := strconv.Atoi(item)
		givspectrum = append(givspectrum, i)
	}
	dropUnusedAmins(givspectrum)
	peptids := make([]string, 1)
	result := make([]string, 0)
	for len(peptids) > 0 {
		selected_peptides := make([]string, 0)
		for _, p := range peptids {
			spectrum := cyclospectrum(p)
			if spectrum[len(spectrum) - 1] > givspectrum[len(givspectrum) - 1] {
				continue
			}
			if len(spectrum) > len(givspectrum) {
				continue
			}
			if spectrum[len(spectrum) - 1] < givspectrum[len(givspectrum) - 1] {
				selected_peptides = append(selected_peptides, p)
				continue
			}
			if cmp(spectrum, givspectrum) {
				result = append(result, p)
			}
		}
		peptids = extend(selected_peptides, givspectrum[len(givspectrum) - 1])
	}
	for _, s := range result {
		for i := range s {
			fmt.Print(aminToStr(s[i:] + s[:i]), " ")
		}
	}
}

