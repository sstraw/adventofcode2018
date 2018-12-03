package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "log"
    "strconv"
    "container/list"
    "bytes"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    re      := regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")

    claims := list.New()
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
        _ = claims.PushBack(c)
    }

    for e1 := claims.Front(); e1 != nil; e1 = e1.Next(){
        c1 := e1.Value.(Claim)
        no_overlap := true
        for e2 := claims.Front(); e2 != nil; e2 = e2.Next() {
            c2 := e2.Value.(Claim)
            if (c1.Overlaps(&c2)) && e1 != e2 {
                no_overlap = false
                break
            }
        }
        if no_overlap {
            fmt.Println("Non overlapping claim found:")
            fmt.Println(c1)
            return
        }
    }
    fmt.Println("No claim found :(")
}

type Claim struct {
    Id, X, Y, W, H, X2, Y2 int
}

func ClaimFromMatch(m []string) Claim{
    c := Claim{}
    c.Id, _ = strconv.Atoi(m[1])
    c.X,  _ = strconv.Atoi(m[2])
    c.X     = c.X
    c.Y,  _ = strconv.Atoi(m[3])
    c.Y     = c.Y
    c.W,  _ = strconv.Atoi(m[4])
    c.H,  _ = strconv.Atoi(m[5])
    c.X2    = c.X + c.W - 1
    c.Y2    = c.Y + c.H - 1
    return c
}

// Checks if two claims overlap
func (c *Claim) Overlaps(c2 *Claim) bool {
    // I'm not proud of how long this took me
    // https://stackoverflow.com/questions/306316/determine-if-two-rectangles-overlap-each-other
    return (c.Y <= c2.Y2 && c.Y2 >= c2.Y &&
            c.X <= c2.X2 && c.X2 >= c2.X)
}

func (c Claim) String() string {
    return fmt.Sprintf("#%v @ %v,%v: %vx%v", c.Id, c.X, c.Y, c.W, c.H)
}

func (c Claim) PrettyString() string {
    var buf bytes.Buffer

    buf.WriteString(c.String())
    buf.WriteString("\n")
    for y :=0; y < (c.Y2+5); y++ {
        for x := 0; x < (c.X2+5); x++{
            if c.Contains(x, y){
                buf.WriteString("#")
            } else {
                buf.WriteString(".")
            }
        }
        buf.WriteString("\n")
    }
    return buf.String()
}

// Whether a coordinate falls inside the claim
func (c *Claim) Contains(x, y int) bool {
    return x >= c.X && x <= c.X2 && y >= c.Y && y <= c.Y2
}

// Returns largest X value possible
func (c *Claim) MaxX() int {
    return c.X + c.W
}

// Returns largest Y value possible
func (c *Claim) MaxY() int {
    return c.Y + c.H
}
