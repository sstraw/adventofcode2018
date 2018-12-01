package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    fin, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer fin.Close()

    scanner := bufio.NewScanner(fin)

    var fluxes []int 
    for scanner.Scan() {
        flux, _ := strconv.Atoi(scanner.Text())
        fluxes = append(fluxes, flux)
    }
    fluxes_len := len(fluxes)

    fin.Close()


    freq := 0
    var freqs []int
    i := 0

    for {
        flux := fluxes[i % fluxes_len]
        freq += flux

        // Check if in list
        for _, v := range freqs{
            if freq == v {
                fmt.Printf("Repeat: %v\n", v)
                return
            }
        }

        freqs = append(freqs, freq)

        i++
    }

}
