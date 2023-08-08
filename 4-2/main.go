package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    sc := bufio.NewScanner(f)
    res := 0
    for sc.Scan() {
        line := sc.Text()
        res += isOverlap(line)
    }
    fmt.Println(res)
}

func isOverlap(line string) int {
    s := strings.Split(line, ",")
    p := [][]string {
        strings.Split(s[0], "-"),
        strings.Split(s[1], "-"),
    }

    pairs := [2][2]int{} 
    for i := 0; i < len(p); i++ {
        for j := 0; j < len(p[i]); j++ {
            n, err := strconv.Atoi(p[i][j])
            if err != nil {
                panic(err)
            }
            pairs[i][j] = n
        }
    }
     
    res := 0
    if pairs[0][0] <= pairs[1][0] && pairs[0][1] >= pairs[1][0] ||
       pairs[1][0] <= pairs[0][0] && pairs[1][1] >= pairs[0][0] {
        res = 1
    }
    return res
}
