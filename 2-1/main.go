package main 
import ( "bufio"
         "fmt" 
         "os"
         "strings"
)

func main() {
   f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    total_score := 0
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        moves := strings.Split(sc.Text(), " ")
        my_move := moves[1]
        opponent_move := moves[0]
        round_score := round_score(opponent_move, my_move)
        total_score += round_score 
        
    }
    fmt.Println(total_score)
}

func round_score(opponent string, mine string) int {
    moves := map[string]int {
        "A": 1,
        "X": 1,
        "B": 2,
        "Y": 2,
        "C": 3,
        "Z": 3,
    }
    winsAgainst := map[int]int {
        1: 3,
        2: 1,
        3: 2,
    }
    score := 0

    if moves[opponent] == moves[mine] {
        score = 3
    }
    if winsAgainst[moves[opponent]] == moves[mine] {
        score = 0
    }
    if winsAgainst[moves[mine]] == moves[opponent] {
        score = 6
    }
    return score + moves[mine]
}
