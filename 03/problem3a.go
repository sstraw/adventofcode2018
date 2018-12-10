package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "log"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    re      := regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")

    var claims []Claim
    lim_x, lim_y := 0,0
    for scanner.Scan() {
        m := re.FindStringSubmatch(scanner.Text())
        if m == nil {
            log.Fatal("No match:", scanner.Text())
        }
        c := ClaimFromMatch(m)

        if lim_x < c.MaxX() {
            lim_x = c.MaxX()
        }
        if lim_y < c.MaxY() {
            lim_y = c.MaxY()
        }
        claims = append(claims, c)
    }

    n_overlap := 0
    max       := lim_x * lim_y
    for i := 0; i < max; i++ {
        x, y, n_claims:= i / lim_x, i % lim_x, 0
        for _, c := range (claims) {
            if c.Contains(x, y) {
                n_claims++
            }
            if n_claims > 1 {
                n_overlap++
                break
            }
        }
    }
    fmt.Println("Overlap:", n_overlap)
}

type Claim struct {
    Id, X, Y, W, H int
}

func ClaimFromMatch(m []string) Claim{
    c := Claim{}
    c.Id, _ = strconv.Atoi(m[1])
    c.X,  _ = strconv.Atoi(m[2])
    c.Y,  _ = strconv.Atoi(m[3])
    c.W,  _ = strconv.Atoi(m[4])
    c.H,  _ = strconv.Atoi(m[5])
    return c
}

func (c Claim) String() string {
    return fmt.Sprintf("#%v @ %v,%v: %vx%v", c.Id, c.X, c.Y, c.W, c.H)
}

// Whether a coordinate falls inside the claim
func (c Claim) Contains(x, y int) bool {
    return x > c.X && x <= (c.X+c.W) && y > c.Y && y <= (c.Y+c.H)
}

// Returns largest X value possible
func (c Claim) MaxX() int {
    return c.X + c.W
}

// Returns largest Y value possible
func (c Claim) MaxY() int {
    return c.Y + c.H
}
