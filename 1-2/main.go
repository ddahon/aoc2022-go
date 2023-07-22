package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    dat, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    inventories := strings.Split(string(dat), "\n\n")
    elves := make([]int, len(inventories))
    for i := 0; i < len(inventories); i++ {
        inventory := strings.Split(inventories[i], "\n")
        sum := 0
        for j := 0; j < len(inventory); j++ {
            calories, err := strconv.Atoi(inventory[j])
            if err == nil {
                sum += calories
            }
        }
        fmt.Println(sum)
        elves[i] = sum

    }
    sort.Sort(sort.Reverse(sort.IntSlice(elves)))
    res := elves[0] + elves[1] + elves[2]
    fmt.Println(res)

}
