package main

import (
    "os"
    "bufio"
    "sort"
    "fmt"
    "regexp"
    "strconv"
    "bytes"
)

type Guard struct {
    Id int
    MinSleep [60]int
}

func main () {
    scanner := bufio.NewScanner(os.Stdin)

    input := make([]string, 0)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    sort.Strings(input)

    reShift := regexp.MustCompile("^\\[\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}\\] Guard #(\\d+) begins shift")
    reSleep := regexp.MustCompile("^\\[\\d{4}-\\d{2}-\\d{2} \\d{2}:(\\d{2})\\] falls asleep$")
    reWakes := regexp.MustCompile("^\\[\\d{4}-\\d{2}-\\d{2} \\d{2}:(\\d{2})\\] wakes up$")

    var gId int
    guards := make(map[int]*Guard)
    for i := 0; i < len(input); i++{
        t     := input[i]
        shift := reShift.FindStringSubmatch(t)
        sleep := reSleep.FindStringSubmatch(t)

        if shift != nil {
            gId, _  = strconv.Atoi(shift[1])
            _, prs := guards[gId]
            if !prs{
                guards[gId] = &Guard{Id: gId}
            }
        } else if sleep != nil {
            start, _ := strconv.Atoi(sleep[1])
            i++
            t = input[i]
            wakes := reWakes.FindStringSubmatch(t)
            if wakes == nil {
                fmt.Println("Got to 2")
                return
            }
            stop, _ := strconv.Atoi(wakes[1])
            guards[gId].Slept(start, stop)
        } else {
            fmt.Println("Got to 1")
            return
        }
    }

    g := guards[gId]
    for _, v := range guards {
        if !g.SleptMore(v){
            g = v
        }
    }

    m := g.HighestMinute()
    fmt.Printf("Guard %v slept most in minutes %v\n%vx%v=%v\n",
                g.Id, m, g.Id, m, g.Id*m)
}

func (g *Guard) Slept (start, stop int) {
    for i := start; i < stop; i++ {
        g.MinSleep[i]++
    }
}

func (g *Guard) TimeSlept () int {
    result := 0
    for _, v := range(g.MinSleep) {
        result += v
    }
    return result
}

func (g Guard) String () string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintf("Guard: %v\n", g.Id))
    for i := 0; i < 6; i++ {
        for j := 0; j < 10; j++ {
            s := fmt.Sprintf("-%3v", g.MinSleep[i*10+j])
            buf.WriteString(s)
        }
        buf.WriteString("\n")
    }

    buf.WriteString("\n")

    return buf.String()
}

func (g *Guard) HighestMinute () int {
    index, max := 0, g.MinSleep[0]
    for i, v := range g.MinSleep {
        if v > max {
            index, max = i, v
        }
    }
    return index
}

func (g *Guard) SleptMore (g2 *Guard) bool {
    return g.TimeSlept() > g2.TimeSlept()
}
