package main

import (
    "os"
    "bufio"
    "fmt"
    "bytes"
)

// Compares two strings and returns a string of overlapping
// values and the number of character differences.
// Max serves as a maximum number of differences allowed
// so the loop can shortcut. Nil assumes all diffs allowed
// -1 is returned for the diffs if max is exceeded
func diff (s1, s2 string, max int) (string, int) {
    diffs := 0

    overlap := bytes.Buffer{}

    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i]{
            diffs++
            if diffs > max {
                return overlap.String(), -1
            }
        } else {
            overlap.WriteByte(s1[i])
        }
    }

    return overlap.String(), diffs
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var codes []string

    for scanner.Scan(){
        codes = append (codes, scanner.Text())
    }

    for i, s1 := range(codes[:len(codes)-1]) {
        for _, s2 := range(codes[i+1:]) {
            overlap, diffs := diff (s1, s2, 1)
            if diffs == 1 {
                fmt.Printf("%s\n%s\n%s\n", s1, s2, overlap)
                return
            }
        }
    }

    fmt.Println("None found")
}
