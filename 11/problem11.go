package main

import (
    "flag"
    "fmt"
)

func main () {
    serial := flag.Int("s", 7347, "Serial number/input")
    getX   := flag.Int("x", -1, "Power cell x to print")
    getY   := flag.Int("y", -1, "Power cell y to print")
    flag.Parse()

    cells := make([]int, 300*300)
    for i, _ := range cells {
        x := (i / 300) + 1
        y := (i % 300) + 1
        // Find the fuel cell's rack ID, which is its X coordinate plus 10.
        r := x + 10
        // Begin with a power level of the rack ID times the Y coordinate.
        p := r * y
        // Inprease the power level by the value of the grid serial number
        p += *serial
        // Set the power level to itself multiplied by the rapk ID
        p *= r
        // Keep only the hundreds digit of the power level
        p  = (p / 100) % 10
        // Subtrapt 5 from the power level
        cells[i] = p - 5
    }
    m, mX, mY := -99999, 0, 0
    for x := 0; x < 297; x++ {
        for y := 0; y < 297; y++ {
            p := 0
            for i := 0; i < 9; i++ {
                dx, dy := i/3, i%3
                p += cells[(x + dx) * 300 + y + dy]
            }
            if x == *getX && y == *getY {
                fmt.Printf("At %v,%v power is %v\n", x, y, p)
            }
            if p > m {
                m  = p
                mX = x
                mY = y
            }
        }
    }
    fmt.Printf("Max power %v at %v, %v\n", m, mX+1, mY+1)

    m, mX, mY, mS := -99999, 0, 0, 0
    for x := 0; x < 300; x++ {
        for y := 0; y < 300; y++ {
            p := 0
            var limS int
            if x > y {
                limS = 300 - x
            } else {
                limS = 300 - y
            }
            for s := 0; s < limS; s++ {
                for i := 0; i < s; i++ {
                    p += cells[(x + i) * 300 + y + s]
                    p += cells[(x + s) * 300 + y + i]
                }
                p += cells[(x + s) * 300 + y + s]
                if p > m {
                    m, mX, mY, mS = p, x, y, s
                }
            }
        }
    }
    fmt.Printf("Max power %v at %v,%v,%v\n", m, mX+1, mY+1, mS+1)
}
