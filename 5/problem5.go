package main

import (
    "os"
    "bufio"
    "fmt"
    "container/list"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var input string
    scanner.Scan()
    input = scanner.Text()

    chars := list.New()
    for _, char := range input {
        chars.PushBack(char)
    }

    for e := chars.Front(); e != nil; {
        c := e.Value.(int32)

        ePrev := e.Prev()
        if ePrev == nil {
            e = e.Next()
            continue
        }

        cPrev := ePrev.Value.(int32)
        cLower := c | 0x20
        cPrevLower := cPrev | 0x20

        if (cLower == cPrevLower && c != cPrev){
            ePPrev := ePrev.Prev()
            tmp := e
            if ePPrev == nil {
                e = e.Next()
            } else {
                e = ePPrev
            }
            _ = chars.Remove(tmp)
            _ = chars.Remove(ePrev)
        } else {
            e = e.Next()
        }
    }

    fmt.Println("Remaining units:", chars.Len())
}
