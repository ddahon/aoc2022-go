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

func round_score(opponent string, outcome string) int {
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
    move_score, outcome_score := 0, 0
    if outcome == "X" { //lose
        move_score = winsAgainst[moves[opponent]]
        outcome_score = 0
    }
    if outcome == "Y" { //draw
        move_score = moves[opponent]
        outcome_score = 3
    }
    if outcome == "Z" { //win
        for k, v := range winsAgainst {
            if v == moves[opponent] {
                move_score = k
            }
        }
        outcome_score = 6
    }
    return move_score + outcome_score 
}
