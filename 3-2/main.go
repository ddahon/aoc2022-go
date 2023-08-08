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
	group := make([]string, 3)
	i := 0
	for sc.Scan() {
		group[i%3] = sc.Text()
		i++
		if i%3 == 0 {
			result += itemPriority(findBadge(group))
		}
	}
	fmt.Println(result)
}

func findBadge(group []string) rune {
    var badge rune
    for _, item := range group[0] {
        if strings.ContainsRune(group[1], item) && strings.ContainsRune(group[2], item) {
            badge = item
        }

    }
	return badge
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
