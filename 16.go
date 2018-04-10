package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

var TranslationMatrix = map[string]string{
"AAA": "K", "AAC": "N", "AAG": "K", "AAU": "N",
"ACA": "T", "ACC": "T", "ACG": "T", "ACU": "T",
"AGA": "R", "AGC": "S", "AGG": "R", "AGU": "S",
"AUA": "I", "AUC": "I", "AUG": "M", "AUU": "I",
"CAA": "Q", "CAC": "H", "CAG": "Q", "CAU": "H",
"CCA": "P", "CCC": "P", "CCG": "P", "CCU": "P",
"CGA": "R", "CGC": "R", "CGG": "R", "CGU": "R",
"CUA": "L", "CUC": "L", "CUG": "L", "CUU": "L",
"GAA": "E", "GAC": "D", "GAG": "E", "GAU": "D",
"GCA": "A", "GCC": "A", "GCG": "A", "GCU": "A",
"GGA": "G", "GGC": "G", "GGG": "G", "GGU": "G",
"GUA": "V", "GUC": "V", "GUG": "V", "GUU": "V",
"UAA": "X", "UAC": "Y", "UAG": "X", "UAU": "Y",
"UCA": "S", "UCC": "S", "UCG": "S", "UCU": "S",
"UGA": "X", "UGC": "C", "UGG": "W", "UGU": "C",
"UUA": "L", "UUC": "F", "UUG": "L", "UUU": "F"}

var ComplementMatrix = map[string]string{"A": "T", "C": "G", "G": "C", "T": "A"}


func translate(rna string) string {
	res := ""
	for i := 0; i < len(rna) - 2; i += 3 {
		codon := rna[i:i + 3]
		if TranslationMatrix[codon] == "X" {
			break
		}
		res += TranslationMatrix[codon]
	}
	return res
}

func transcript(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func reverseComplement(dna string) string {
	res := ""
	for i := len(dna) - 1; i >= 0; i-- {
		res += ComplementMatrix[dna[i:i + 1]]
	}
	return res
}

func main() {
	fin, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(fin)
	scanner.Scan()
	dna := scanner.Text()
	scanner.Scan()
	peptide := scanner.Text()

	k := len(peptide) * 3
	res := make([]string, 0)

	for i := 0; i < len(dna) - k + 1; i++ {
		if translate(transcript(dna[i:i + k])) == peptide {
			res = append(res, dna[i:i + k])
		}
		if translate(transcript(reverseComplement(dna[i:i + k]))) == peptide {
			res = append(res, dna[i:i + k])
		}
	}

	for _, e := range res {
		fmt.Println(e)
	}
}

