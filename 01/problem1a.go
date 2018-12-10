package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    freq := 0

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        flux, _ := strconv.Atoi(scanner.Text())

        freq += flux
    }

    fmt.Printf("Val: %v\n", freq)
}
