package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
)

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("[#\\.]+")

    scanner.Scan()
    input := re.FindAllString(scanner.Text(), -1)[0]
    scanner.Scan() // Blank line

    buffLen := (2 * 2000) + 2 // Can only spread 2 * gen per gen, each side

    pots := make([]byte, len(input) + (2 * buffLen))

    for i := 0; i < buffLen; i++ {
        pots[i] = byte('.')
    }
    for i := buffLen; i < buffLen + len(input); i++ {
        pots[i] = byte(input[i-buffLen])
    }
    for i := buffLen + len(input); i < (2 * buffLen) + len(input); i++ {
        pots[i] = byte('.')
    }

    patterns := make(map[string]byte)
    for scanner.Scan() {
        m := re.FindAllString(scanner.Text(), -1)
        patterns[m[0]] = byte(m[1][0])
    }

    for gen := 0; gen < 2000; gen++ {
        newPots := make([]byte, len(pots))
        newPots[0], newPots[1], newPots[len(pots)-1], newPots[len(pots)-2] = 46, 46, 46, 46
        sum := 0
        for i := 2; i < (len(pots)-2); i++ {
            newPots[i] = patterns[string(pots[i-2:i+3])]
            if newPots[i] == byte('#') {
                sum += i - buffLen
            }
        }

        pots = newPots

        fmt.Printf("Generation %4v: %9v\n", gen, sum)
    }

}
