package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	anagrams := make(map[string][]string)

	for _, w := range words {
		sliceOfWord := strings.Split(w, "")
		sort.Strings(sliceOfWord)
		sortedWord := strings.Join(sliceOfWord, "")

		anagrams[sortedWord] = append(anagrams[sortedWord], w)
	}

	for _, a := range anagrams {
		fmt.Println(a)
	}
}
