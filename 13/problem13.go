package main

import (
    "os"
    "fmt"
    "bufio"
    "bytes"
    "sort"
)

const NORTH = 4
const EAST  = 5
const SOUTH = 6
const WEST  = 7

type Cart struct {
    X, Y, Facing, Turn int
}

type ByPos []*Cart

func (a ByPos) Len()          int  { return len(a) }
func (a ByPos) Swap(i, j int)      {a[i], a[j] = a[j], a[i]}
func (a ByPos) Less(i, j int) bool {
    if a[i] == nil { return false }
    if a[j] == nil { return true}
    return (a[i].X < a[j].X) || (a[i].X == a[j].X && a[i].Y < a[j].Y)
}

func main () {
    scanner := bufio.NewScanner(os.Stdin)

    var buf bytes.Buffer

    scanner.Scan()
    xWidth := len(scanner.Text())
    yWidth := 1
    buf.WriteString(scanner.Text())

    for scanner.Scan() {
        buf.WriteString(scanner.Text())
        yWidth++
    }

    track := buf.String()

    carts := make([]*Cart, 0)

    for i, c := range track {
        if c == rune('>') {
            carts = append(carts, &Cart{X: i%xWidth, Y: i/xWidth, Facing: EAST})
        }
        if c == rune('^') {
            carts = append(carts, &Cart{X: i%xWidth, Y: i/xWidth, Facing: NORTH})
        }
        if c == rune('<') {
            carts = append(carts, &Cart{X: i%xWidth, Y: i/xWidth, Facing: WEST})
        }
        if c == rune('v') {
            carts = append(carts, &Cart{X: i%xWidth, Y: i/xWidth, Facing: SOUTH})
        }
    }

    noCrashes := true

    sort.Sort(ByPos(carts))
    for tick := 0; len(carts) > 1; tick++ {
        for i, c := range carts {
            if c == nil {continue}
            t := track[c.Y * xWidth + c.X]
            switch t {
            case ' ':
                fmt.Println("Got to a space. oh no.")
                return
            case '/':
                switch c.Facing {
                case NORTH:
                    c.Facing = EAST
                case SOUTH:
                    c.Facing = WEST
                case EAST:
                    c.Facing = NORTH
                case WEST:
                    c.Facing = SOUTH
                }
            case '\\':
                switch c.Facing {
                case NORTH:
                    c.Facing = WEST
                case SOUTH:
                    c.Facing = EAST
                case EAST:
                    c.Facing = SOUTH
                case WEST:
                    c.Facing = NORTH
                }
            case '+':
                c.Facing = (c.Facing + (c.Turn % 3 - 1)) % 4 + 4
                c.Turn++
            }
            // Now we're facing a direction. Update x/y
            switch c.Facing {
            case NORTH:
                c.Y--
            case SOUTH:
                c.Y++
            case EAST:
                c.X++
            case WEST:
                c.X--
            }
            // Got new position. Loop through carts and check if we've crashed
            for i2, c2 := range carts {
                if c2 != nil && c.X == c2.X && c.Y == c2.Y && c != c2 {
                    if noCrashes {
                        fmt.Printf("First crash at %v,%v\n", c.X, c.Y)
                        noCrashes = false
                    }
                    carts[i], carts[i2] = nil, nil
                }
            }
        }
        sort.Sort(ByPos(carts))
        for i, _ := range carts {
            if carts[i] == nil {
                carts = carts[:i]
                break
            }
        }
    }

    if len(carts) == 1 {
        fmt.Printf("Last cart at %v, %v\n", carts[0].X, carts[0].Y)
    } else {
        fmt.Println("No carts left")
    }
}

func (c Cart) String() string {
    return fmt.Sprintf("X: %v Y: %v F: %v Turn: %v", c.X, c.Y, c.Facing, c.Turn)
}
