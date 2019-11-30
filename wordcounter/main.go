package main

import (
    "fmt"
    )


type Node struct {
    char uint8
    word_count uint
    edges []*Node
}


func (n *Node) get_inserter() func(uint8) {
    current := n
    inword := false

    return func(b uint8) {
        if b == ' ' {
            fmt.Printf("_")
            if (inword) {
                fmt.Printf("|%v", current.word_count)
                current.word_count += 1
                current = n // Start at the beginning again
            }
            inword = false
            return
        }
        fmt.Printf("%v",string(b))
        inword = true
        edges := &(current.edges)
        for e := 0; e < len(*edges); e++ {
            if (*edges)[e].char == b {
                fmt.Printf("<")
                current = (*edges)[e]
                return
            }
        }
        fmt.Printf("^")
        new_node := &Node{b, 0, nil}
        current.edges = append(current.edges, new_node)
        current = new_node
        // fmt.Printf("\n+%v\n", current.edges)

    }

}

func (n *Node) dump(prefix []uint8, depth int) {
        if n.char == 0 { // step over the first node if it's a start node
            // fmt.Printf("\nSTART depth %v prefix '%v'\n", depth, prefix)
        } else {
            prefix = append(prefix, n.char)
            if n.word_count > 0 {
                fmt.Printf("=%v: '%v'\n", n.word_count, string(prefix))
            }
        }
        for e := range(n.edges) {
            n.edges[e].dump(prefix, depth+1)
        }
}


func main() {
    fmt.Println("\nWordCount")
}
