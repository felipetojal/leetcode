package main

import (
	"fmt"
	"strings"
)

type occurence struct {
	character string
	count     int
	positions []int
}

func longestCommonPrefix(strs []string) string {

	prefixes := make([]string, 0)

	for i := 0; i < len(strs); i++ {

	}

	return ""
}

func createStrOccurrences(str string) map[string]occurence {
	m := make(map[string]occurence)
	split := strings.Split(str, "")
	for pos, char := range split {
		o, ok := m[char]
		if !ok {
			m[char] = occurence{
				character: char,
				count:     1,
				positions: []int{pos},
			}
		} else {
			o.count++
			o.positions = append(o.positions, pos)
			m[char] = o
		}
	}

	return m
}

/*
Ideia: fazer por ordem sequencial. Teríamos O(n).

Pega a primeira string, vê se ela tem algum prefixo em comum com a segunda.
Se elas não tiverem, ja podemos finalizar a função pois não há nada em
comum entre todas as strings.

Se tiverem algo em comum, separamos somente a parte em comum e comparamos
com a string seguinte, repetindo o processo até o final do slice.
*/

func printOccurenceMap(m map[string]occurence) {
	for s, o := range m {
		fmt.Println("--------")
		fmt.Printf("string: %v\n", s)
		fmt.Printf("occurence character: %v\n", o.character)
		fmt.Printf("occurrence count: %v\n", o.count)
		fmt.Printf("occurence positions: %v\n", o.positions)
	}
}

func main() {
	m := createStrOccurrences("paralelepipedo")
	printOccurenceMap(m)
}
