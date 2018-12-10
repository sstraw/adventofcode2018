package main

import (
    "os"
    "bufio"
    "fmt"
    "container/list"
    "regexp"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var input string
    scanner.Scan()
    input = scanner.Text()

    nPoly := react(input)
    fmt.Println("Reaction with sub:", nPoly)

    for i := 65; i < 91; i++ {
        re := regexp.MustCompile(fmt.Sprintf("[%c%c]", i, i | 0x20))
        new_input := re.ReplaceAllLiteralString(input, "")
        newReact := react(new_input)
        if newReact < nPoly {
            nPoly = newReact
        }
    }

    fmt.Println("Minimal reaction:", nPoly)
}

func react(input string) int {
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

    return chars.Len()
}
