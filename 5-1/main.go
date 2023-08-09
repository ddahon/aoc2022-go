package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type move struct {
    n int
    src int
    dst int
}

func main() {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    sc := bufio.NewScanner(f)
    crates_def := true
    var crates []string
    var moves []string
    for sc.Scan() {
        line := sc.Text()
        if crates_def {
            if len(line) != 0 {
                crates = append(crates, line)
            } else {
                crates_def = false
            }
        } else {
            moves = append(moves, line)
        }
    }

    res := execute_moves(parse_crates(crates), parse_moves(moves))
    fmt.Println(first_row(res))
}

func parse_crates(crates []string) []string {
    var res []string

    for i := 1; i < 35; i += 4 {
        s := ""
        for j := 0; j < len(crates) - 1; j++ {
            x := crates[j][i]
            if string(x) != " " {
                s += string(crates[j][i]) 
            }
        }
        res = append(res, s)
    }
    return res
}

func parse_moves(moves []string) []move {
    var res []move

    re := regexp.MustCompile(`move (\d+) from ([0-9]) to ([0-9])`)
    for _, e := range moves {
        matches := re.FindStringSubmatch(e)
        n, err := strconv.Atoi(matches[1])
        if err != nil {
            panic(err)
        }
        src, err := strconv.Atoi(matches[2])
        if err != nil {
            panic(err)
        }
        dst, err := strconv.Atoi(matches[3])
        if err != nil {
            panic(err)
        }
        res = append(res, move{n, src, dst})
    }
    return res
}

func execute_moves(crates []string, moves []move) []string {
    fmt.Println(crates)
    for _, m := range moves {
        x := crates[m.src - 1][:m.n]
        crates[m.dst - 1] = reverse(x) + crates[m.dst - 1]
        crates[m.src - 1] = crates[m.src - 1][m.n:]
    }
    return crates
}

func first_row(crates []string) string {
    res := ""
    for _, e := range crates {
        if len(e) > 0 {
            res += string(e[0])
        } else {
            res += " "
        }
    }
    return res
}

func reverse(s string) string {
    res := ""
    for _, c := range s {
        res = string(c) + res
    }
    return res
}
