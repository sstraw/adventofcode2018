package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "strconv"
)


type Point struct {
    X, Y, Xv, Yv int
}

type Points []*Point
func (p Points) Len() int { return len(p) }
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Points) Less(i, j int) bool {
    if p[i].X < p[j].X { return true }
    if p[i].X > p[j].X { return false }
    if p[i].Y < p[j].Y { return true }
    return false
}

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("-?\\d+")

    points := make([]*Point, 0)

    for scanner.Scan() {
        m := re.FindAllString(scanner.Text(), -1)
        p := &Point{}
        p.X , _ = strconv.Atoi(m[0])
        p.Y , _ = strconv.Atoi(m[1])
        p.Xv, _ = strconv.Atoi(m[2])
        p.Yv, _ = strconv.Atoi(m[3])
        points = append(points, p)
    }

    for i := 0; ;i++ {
        minX, minY, maxX, maxY := 99999999, 99999999, 0, 0
        for _, p := range points {
            if p.X < minX { minX = p.X }
            if p.Y < minY { minY = p.Y }
            if p.X > maxX { maxX = p.X }
            if p.Y > maxY { maxY = p.Y }
        }

        area := (maxX-minX)*(maxY-minY)

        if area < 600 {
            fmt.Println("Found possible message", i)
            for y := minY; y <= maxY; y++ {
                for x := minX; x <=maxX; x++ {
                    var p *Point
                    for _, v := range points {
                        if v.X == x && v.Y == y {
                            p = v
                            break
                        }
                    }
                    if p == nil {
                        fmt.Print(".")
                    } else {
                        fmt.Print("#")
                    }
                }
                fmt.Println()
            }
        }
        for _, p := range points {
            p.Move()
        }
    }
}

func (p Point) String() string {
    return fmt.Sprintf("X: %3v Y: %3v Xv: %3v Yv: %3v", p.X, p.Y, p.Xv, p.Yv)
}

func (p *Point) Move () {
    p.X += p.Xv
    p.Y += p.Yv
}
