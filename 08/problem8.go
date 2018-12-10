package main

import (
    "regexp"
    "os"
    "fmt"
    "bufio"
    "strconv"
)


type Node struct {
    Children []int
    Meta []int
}

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    re := regexp.MustCompile("\\d+")

    rawTree := make([]int, 0)
    for scanner.Scan() {
        m := re.FindAllString(scanner.Text(), -1)
        for _, v := range m {
            i, _ := strconv.Atoi(v)
            rawTree = append(rawTree, i)
        }
    }
    treeData := make([]*Node, 0)
    rootNode := Parse(rawTree, &treeData)

    metaSum := 0
    for _, node := range(treeData) {
        for _, i := range (node.Meta) {
            metaSum += i
        }
    }
    fmt.Println("Sum of meta:", metaSum)

    val := Value(rootNode, &treeData)
    fmt.Println("Value of root:", val)
}

func Parse (input []int, output *[]*Node) *Node {
    n, _ := helpParse(input, output, 0)
    *output = append(*output, n)
    return n
}
func helpParse (input []int, output *[]*Node, index int) (*Node, int) {
    node := &Node{}
    nChild := input[index]
    index++
    nMeta := input[index]
    index++

    for j := 0; j < nChild; j++ {
        var newChild *Node
        newChild, index = helpParse(input, output, index)
        childIndex := len(*output)
        *output = append(*output, newChild)
        node.Children = append(node.Children, childIndex)
    }
    for j := 0; j < nMeta; j, index = j+1, index+1 {
        node.Meta = append(node.Meta, input[index])
    }
    return node, index
}

func Value (n *Node, treeData *[]*Node) int {
    i := 0
    if len(n.Children) == 0 {
        for _, v := range n.Meta {
            i += v
        }
    } else {
        for _, v := range n.Meta {
            v = v - 1 //Account for offset
            if v < 0 || v >= len(n.Children) {
                continue // skip if not a child
            } else {
                i += Value((*treeData)[n.Children[v]], treeData)
            }
        }
    }
    return i
}
