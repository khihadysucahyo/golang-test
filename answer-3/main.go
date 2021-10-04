package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	indexAfterFirstBracketFound := indexFirstBracketFound + 1

	indexClosingBracketFound := strings.Index(str[indexAfterFirstBracketFound:], ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	return str[indexAfterFirstBracketFound:indexClosingBracketFound]

}

func main() {
	fmt.Println(findFirstStringInBracket("(tes)"))
}
