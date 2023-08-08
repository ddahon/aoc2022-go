package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	result := 0
	for sc.Scan() {
		line := sc.Text()
		result += itemPriority(findDuplicateItem(line))
	}
	fmt.Println(result)
}

func findDuplicateItem(rucksack string) rune {
	compartment_size := len(rucksack) / 2
	c1, c2 := rucksack[:compartment_size], rucksack[compartment_size:]
	var duplicate rune
	for _, item := range c1 {
		if strings.ContainsRune(c2, item) {
			duplicate = item
			break
		}
	}
	return duplicate
}

func itemPriority(item rune) int {
	priority := 0
	if unicode.IsUpper(item) {
		priority = int(item) - 65 + 27
	} else {
		priority = int(item) - 96
	}
	return priority
}
