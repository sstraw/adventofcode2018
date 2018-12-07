package main

import (
    "os"
    "bufio"
    "fmt"
    "regexp"
)

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")

    var steps [26]*Step
    for scanner.Scan() {
        m := re.FindStringSubmatch(scanner.Text())

        s := steps[m[2][0]-0x41]
        if s == nil {
            steps[m[2][0]-0x41] = &Step{}
            s = steps[m[2][0]-0x41]
        }
        s.Prereqs = append(s.Prereqs, m[1][0])

        if s := steps[m[1][0]-0x41]; s == nil {
            steps[m[1][0]-0x41] = &Step{}
        }

    }

    workers := make([]*Worker, 5)

    for i, _ := range workers {
        workers[i] = &Worker{}
    }

    i := 0
    for ;;i++ {
        fmt.Printf("===%v===\n", i)
        for _, w := range workers {
            if w.Project != 0 {
                fmt.Println("Subtracting")
                w.Time--
                if w.Time == 0 {
                    steps[w.Project-0x41] = nil
                    w.Project = 0
                }
            }
            if w.Project == 0 {
                p := getJob(steps)
                if p == 0 {
                    continue
                }
                fmt.Printf("Assigning %c\n", p)
                steps[p-0x41].Working = true
                w.Project = p
                w.Time = 61 + int(p) - 0x41
            }
            fmt.Printf("P:%c , T:%v \n", w.Project, w.Time)
        }
        if allDone(steps){
            break
        }
    }
    fmt.Println("Took", i-1)
}

type Worker struct {
    Project byte
    Time int
}

type Step struct {
    Prereqs []byte
    Working bool
}

func (s Step) String() string {
    return fmt.Sprintf("Working %v prereqs %v\n", s.Working, s.Prereqs)
}


func getJob(steps [26]*Step) byte{
    for i, s := range steps {
        if s == nil || s.Working {
            continue
        }
        ready := true
        for _, s2 := range s.Prereqs {
            prereq := steps[s2-0x41]
            if prereq != nil {
                ready = false
            }
        }
        if ready {
            return byte(i + 0x41)
        }
    }
    return 0
}

func allDone(steps [26]*Step) bool {
    for _, s := range steps {
        if s != nil {
            return false
        }
    }
    return true
}
