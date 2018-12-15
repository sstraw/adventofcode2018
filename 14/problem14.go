package main

import (
    "flag"
    "fmt"
    "container/list"
)

func main () {
    nRecipes := flag.Int("r", 556061, "Number of recipes to go to")
    flag.Parse()

    recipes := list.New()

    recipes.PushBack(3)
    recipes.PushBack(7)

    elf1 := recipes.Front()
    elf2 := recipes.Back()

    out := make([]int, 0, 10)

    for recipes.Len() < *nRecipes + 10{
        recipe1, recipe2 := elf1.Value.(int), elf2.Value.(int)
        n := recipe1 + recipe2
        newRecipe1, newRecipe2 := n/10, n%10

        if newRecipe1 != 0 {
            recipes.PushBack(newRecipe1)
            if recipes.Len() > *nRecipes {
                out = append(out, newRecipe1)
            }
        }
        recipes.PushBack(newRecipe2)
        if recipes.Len() > *nRecipes {
            out = append(out, newRecipe2)
        }
        for i := 0; i < 1 + recipe1; i++ {
            elf1 = elf1.Next()
            if elf1 == nil {elf1 = recipes.Front()}
        }
        for i := 0; i < 1 + recipe2; i++ {
            elf2 = elf2.Next()
            if elf2 == nil {elf2 = recipes.Front()}
        }
    }


    fmt.Print("String: ")
    for _, i := range out {
        fmt.Print(i)
    }
    fmt.Println()
}
