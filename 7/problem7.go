package main

import (
    "os"
    "bufio"
    "fmt"
    "regexp"
)

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")

    steps := make([][]byte, 26)
    for scanner.Scan() {
        m := re.FindStringSubmatch(scanner.Text())

        if m != nil {
            prereq := m[1][0]
            v := m[2][0]

            steps[v-0x41] = append(steps[v-0x41], prereq)
            if steps[prereq-0x41] == nil{
                steps[prereq-0x41] = make([]byte, 0)
            }
        } else {
            fmt.Println("Regex is wrong")
            return
        }
    }

    for {
        done := true
        for i, v := range steps {
            if v != nil {
                prereqs_done := true
                for _, prereq := range v {
                    if steps[prereq-0x41] != nil {
                        prereqs_done = false
                    }
                }
                if prereqs_done {
                    fmt.Printf("%c", i+0x41)
                    steps[i] = nil
                    done = false
                    break
                }
            } else {
                //Empty slice
                continue
            }
        }
        if done {
            fmt.Println("")
            break
        }
    }
}
