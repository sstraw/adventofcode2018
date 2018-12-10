package main

import (
    "flag"
    "fmt"
    "container/ring"
)

func main () {
    nPlayers   := flag.Int("p", 429, "Number of players")
    lastMarble := flag.Int("m", 70901, "Score of last marble")
    flag.Parse()

    circle      := ring.New(1)
    //Start with the zero marble as current
    circle.Value = 0

    //List of scores to track
    scores := make([]int, *nPlayers)


    for newMarble := 1; newMarble <= *lastMarble; newMarble++ {
        if newMarble % 23 == 0 {
            currPlayer := (newMarble - 1) % *nPlayers
            scores[currPlayer] += newMarble
            for i:=0; i < 7; i++ {
                circle = circle.Prev()
            }
            scores[currPlayer] += circle.Value.(int)
            circle = circle.Prev()
            circle.Unlink(1)
            circle = circle.Next()
        } else {
            circle = circle.Next()
            circle.Link(ring.New(1))
            circle = circle.Next()
            circle.Value = newMarble
        }
    }

    m := scores[0]
    for _, v := range scores {
        if v > m {
            m = v
        }
    }

    fmt.Println("Max score", m)
}
