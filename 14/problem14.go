package main

import (
    "flag"
    "fmt"
    "strconv"
)

func main () {
    rawInput := flag.String("r", "556061", "Number of recipes to go to")
    flag.Parse()

    recipes := make([]int, 2)
    recipes[0], recipes[1] = 3, 7

    input, _ := strconv.Atoi(*rawInput)
    inputArray := make([]int, 0)
    for i, _ := range *rawInput {
        inputArray = append(inputArray, int((*rawInput)[i] - 0x30))
    }
    nInputArray := len(inputArray)

    solveda, solvedb := false, false
    elf1, elf2 := 0, 1
    for !solveda || !solvedb {
        recipe1, recipe2 := recipes[elf1], recipes[elf2]
        n := recipe1 + recipe2
        newRecipe1, newRecipe2 := n/10, n%10
        if newRecipe1 != 0 {
            recipes = append(recipes, newRecipe1)
        }
        recipes = append(recipes, newRecipe2)

        elf1 = (elf1 + recipe1 + 1) % len(recipes)
        elf2 = (elf2 + recipe2 + 1) % len(recipes)

        if !solveda && len(recipes) >= input + 10 {
            fmt.Print("Part a: ")
            for i := input; i < input+10; i++ { fmt.Print(recipes[i]) }
            fmt.Println()
            solveda = true
        }

        if !solvedb && len(recipes) >= len(inputArray) + 5{
            end := len(recipes) - 1
            for o := 0; o < 2; o++ {
                solvedb = true
                for i, _ := range inputArray {
                    if inputArray[nInputArray-i-1] != recipes[end-o-i] {
                        solvedb = false
                        break
                    }
                }
                if solvedb {
                    fmt.Println("Part b:", end - o - len(inputArray) + 1)
                    break
                }
            }
        }
    }
}
