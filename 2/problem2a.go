package main

import (
    "fmt"
    "os"
    "bufio"
)

// Given a string, returns whether there are exactly two
// and/or exactly three occurances of any character in it
func getNums (s string) (int, int) {
    twos, threes := 0, 0

    alphabet := make([]int, 26, 26)

    for _, c := range (s) {
        alphabet[c-97] += 1
    }

    for _, c := range (alphabet) {
        switch (c) {
            case 2 :
                twos = 1
            case 3:
                threes = 1
        }
    }

    return twos, threes
}

func main () {
    scanner := bufio.NewScanner(os.Stdin)

    twos, threes := 0, 0

    for scanner.Scan() {
        n_twos, n_threes := getNums(scanner.Text())
        twos += n_twos
        threes += n_threes
    }

    checksum := twos * threes
    fmt.Println("Checksum:", checksum)
}
