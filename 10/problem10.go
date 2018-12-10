package main

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
    "strconv"
    "bytes"
    "flag"
)

type Star struct {
    X, Y, Xv, Yv int
}

type Sky struct {
    Stars []*Star
    MinX, MinY, MaxX, MaxY, Area int
}

func main () {
    limArea := flag.Int("a", 600, "Max area coverage of stars for possible message")
    limMatch:= flag.Int("m", 1, "Max number of matches to look for. -1 for unlimited")
    flag.Parse()

    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("-?\\d+")

    sky := &Sky{}

    for scanner.Scan() {
        m := re.FindAllString(scanner.Text(), -1)
        s := &Star{}
        s.X , _ = strconv.Atoi(m[0])
        s.Y , _ = strconv.Atoi(m[1])
        s.Xv, _ = strconv.Atoi(m[2])
        s.Yv, _ = strconv.Atoi(m[3])
        sky.Stars = append(sky.Stars, s)
    }

    nMatches := 0
    sky.Update()
    for i := 1; nMatches != *limMatch; i++ {
        if sky.Area < *limArea {
            fmt.Println("Found possible message at second", i)
            fmt.Println(sky)
            nMatches++
        }
        sky.Update()
    }
}

func (s *Sky) Update() {
    p := s.Stars[0]
    s.MinX, s.MaxX, s.MinY, s.MaxY = p.X, p.X, p.Y, p.Y
    for _, p := range s.Stars {
        p.Move()
        if p.X < s.MinX { s.MinX = p.X }
        if p.Y < s.MinY { s.MinY = p.Y }
        if p.X > s.MaxX { s.MaxX = p.X }
        if p.Y > s.MaxY { s.MaxY = p.Y }
    }
    s.Area = (s.MaxX - s.MinX) * (s.MaxY - s.MinY)
}


func (p Star) String() string {
    return fmt.Sprintf("X: %3v Y: %3v Xv: %3v Yv: %3v", p.X, p.Y, p.Xv, p.Yv)
}

func (p *Star) Move () {
    p.X += p.Xv
    p.Y += p.Yv
}

func (s Sky) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintf("Area: %v MinX: %v MaxX: %v MinY: %v MaxY: %v", s.Area, s.MinX, s.MaxX, s.MinY, s.MaxY))
    for y := s.MinY; y <= s.MaxY; y++ {
        buf.WriteString("\n")
        for x := s.MinX; x <= s.MaxX; x++ {
            taken := false
            for _, p := range s.Stars {
                if x == p.X && y == p.Y {
                    buf.WriteString("#")
                    taken = true
                    break
                }
            }
            if ! taken {
                buf.WriteString(".")
            }
        }
    }
    return buf.String()
}

