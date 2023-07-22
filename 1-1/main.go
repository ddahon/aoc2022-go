package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
   if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := os.ReadFile("input")
    check(err)
    lines := strings.Split(string(dat), "\n\n")
    max := 0
    for i := 0; i < len(lines); i++ {
        inventory := strings.Split(lines[i], "\n") 
        fmt.Println(inventory)
        sum := 0
        for j := 0; j < len(inventory); j++ {
            calories, err := strconv.Atoi(inventory[j]) 
            if err == nil {
                sum += calories
            }
       }
       if sum > max {
           max = sum
        }
    }
    fmt.Println(max)
}
