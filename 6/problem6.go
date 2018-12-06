package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strconv"
    "math"
)

type Coord struct {
    X, Y, A int
}

func main () {

    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("(\\d+), (\\d+)")

    //These will form the boundaries of our search for area.
    //Anything on one of the edges will necessarily be infinite
    north, west, south, east := 99999999, 99999999, 0, 0

    coords := make([]*Coord, 0)

    for scanner.Scan() {
        m := re.FindStringSubmatch(scanner.Text())
        if m != nil {
            x, _ := strconv.Atoi(m[1])
            y, _ := strconv.Atoi(m[2])
            if y < north {
                north = y
            }
            if y > south {
                south = y
            }
            if x < west {
                west = x
            }
            if x > east {
                east = x
            }
            coords = append(coords, &Coord{X: x, Y: y})
        }
    }

    for x := west + 1; x < east; x++ {
        for y := north + 1; y < south; y++ {
            c := getClosest(coords, x, y)
            if c != nil{
                c.A += 1
            }
        }
    }
    for y := north; y <= south; y++ {
        c := getClosest(coords, west, y)
        if c != nil {
            c.A = -1
        }
        c = getClosest(coords, east, y)
        if c != nil {
            c.A = -1
        }
    }
    for x := west; x <= east; x++ {
        c := getClosest(coords, x, north)
        if c != nil {
            c.A = -1
        }
        c = getClosest(coords, x, south)
        if c != nil {
            c.A = -1
        }
    }

    max := 0
    for _, c := range coords {
        if c.A > max {
            max = c.A
        }
    }
    fmt.Println("Max area:", max)
 }

func getClosest(coords []*Coord, x, y  int) *Coord {
    d := 999999999
    r := coords[0]
    tied := false
    for _, c := range coords {
        dd := getDistance(x, y, c.X, c.Y)
        if dd < d {
            d = dd
            r = c
            tied = false
        } else if dd == d {
            tied = true
        }
    }
    if tied {
        return nil
    }
    return r
}

func getDistance(x1, y1, x2, y2 int) int {
    return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func (c Coord) String() string {
    return fmt.Sprintf("X: %v Y: %v A: %v", c.X, c.Y, c.A)
}
