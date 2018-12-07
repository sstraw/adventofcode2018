package main

import (
    "os"
    "bufio"
    "fmt"
    "regexp"
    "flag"
)

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")

    nWorkers := flag.Int("w", 5, "Number of workers for part 2. " +
                         "2 workers for test input")
    weight   := flag.Int("t", 61, "Time for step A to complete. " +
                         "For test input, weight is 1")
    test     := flag.Bool("test", false, "Set to true for test worker/weight")

    flag.Parse()
    if *test {
        *nWorkers = 2
        *weight = 1
    }

    fmt.Println(*nWorkers, "workers, A =", *weight, "second.")

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

    fmt.Print("7a: ")
    stepsA := steps
    for !allDone(stepsA) {
        p := getJob(stepsA)
        fmt.Printf("%c", p)
        stepsA[p-0x41] = nil
    }
    fmt.Println()



    workers := make([]*Worker, *nWorkers)

    for i, _ := range workers {
        workers[i] = &Worker{}
    }

    i := 0
    for ;;i++ {
        for _, w := range workers {
            if w.Project != 0 {
                w.Time--
                if w.Time == 0 {
                    steps[w.Project-0x41] = nil
                    w.Project = 0
                }
            }
        }
        for _, w := range workers {
            if w.Project == 0 {
                p := getJob(steps)
                if p == 0 {
                    continue
                }
                steps[p-0x41].Working = true
                w.Project = p
                w.Time = *weight + int(p) - 0x41
            }
        }
        if allDone(steps){
            break
        }
    }
    fmt.Println("7b:", i, "seconds")
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
