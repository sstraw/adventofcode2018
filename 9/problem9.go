package main

import (
    "flag"
    "fmt"
    "container/list"
)

func main () {
    nPlayers   := flag.Int("p", 429, "Number of players")
    lastMarble := flag.Int("m", 70901, "Score of last marble")
    flag.Parse()

    circle     := list.New()
    //Start with the zero marble as current
    currMarble := circle.PushBack(0)

    //List of scores to track
    scores := make([]int, *nPlayers)

    currPlayer := 0

    for newMarble := 1; newMarble <= *lastMarble; newMarble++ {
        if newMarble % 23 == 0 {
            scores[currPlayer] += newMarble
            for i := 0; i < 7; i++ {
                currMarble = currMarble.Prev()
                if currMarble == nil {
                    currMarble = circle.Back()
                }
            }
            tmp := currMarble
            currMarble = currMarble.Next()
            if currMarble == nil {
                currMarble = circle.Front()
            }
            scores[currPlayer] += circle.Remove(tmp).(int)
        } else {
            placement := currMarble.Next()
            if placement == nil {
                placement = circle.Front()
            }
            currMarble = circle.InsertAfter(newMarble, placement)
        }
        currPlayer++
        if !(currPlayer < *nPlayers) {
            currPlayer = 0
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
