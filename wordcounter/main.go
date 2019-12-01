package main

import (
    "fmt"
    )


type Node struct {
    char uint8
    word_count uint
    children []*Node
}

func filterToLower(b byte) byte {

        if b >= 'A' && b <= 'Z' {
            return 'a' + (b - 'A')
        } else if b >= 'a' && b <= 'z' {
           return b
        } else {
           return ' ' // binary etc converted to space
        }
}


func (n *Node) get_inserter() func(uint8) {
    current := n
    inword := false

    return func(b uint8) {
        b = filterToLower(b)

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
        children := &(current.children)
        for e := range(*children) {
            if (*children)[e].char == b {
                fmt.Printf("<")
                current = (*children)[e]
                return
            }
        }
        fmt.Printf("^")
        new_node := &Node{b, 0, nil}
        current.children = append(current.children, new_node)
        current = new_node
        // fmt.Printf("\n+%v\n", current.children)

    }

}

type Wordcount struct {
    count uint
    word string
}

type Topwords struct {
    sortedwords []*Wordcount
    start Node
}

func NewTopwords(countmax uint) *Topwords {
        var tw Topwords
        tw.sortedwords = make([]*Wordcount,0, countmax+1)

        return &tw
}


func (n *Node) preorder_walk(prefix []uint8, action func (n *Node, prefix *[]uint8)) {
        if n.char == 0 { // step over the first node if it's a start node
                // fmt.Printf("\nSTART depth %:v prefix '%v'\n", depth, prefix)
        } else {
            prefix = append(prefix, n.char)
            if n.word_count > 0 {
                action(n, &prefix)
            }
        }
        for e := range(n.children) {
            n.children[e].preorder_walk(prefix, action)
        }
}


func (tw *Topwords) most_frequent() {

    action := func(n *Node, prefix *[]uint8) {
        w := new(Wordcount)
        w.count = n.word_count
        w.word = string(*prefix)

        max := cap(tw.sortedwords)
        if len(tw.sortedwords) < max {
            tw.sortedwords = append(tw.sortedwords, w)
        }  else {
            tw.sortedwords[max-1] = w
            // So what happens to the item at the end of the list of topwords? 
            // we replaced it without a comparison! Well, we use the end of the
            // slice as a sort of "entrance/exit hall" and don't intend to show
            // what's in there finally anyhow - so whatever's "in" there is not
            // important as it won't make the final "cut" anyhow.
        }

        // Insertion sort to swap the new word up to its
        // position.
        fmt.Printf("---%v\n", len(tw.sortedwords))
        for f := len(tw.sortedwords)-1; f >= 1 ; f-- {
            if  tw.sortedwords[f].count > tw.sortedwords[f-1].count {
                tmp := tw.sortedwords[f]
                tw.sortedwords[f] = tw.sortedwords[f-1]
                tw.sortedwords[f-1] = tmp
            }
        }
    }

    tw.start.preorder_walk([]uint8{}, action)
    tw.sortedwords = tw.sortedwords[:len(tw.sortedwords)-1]
}

func (n *Node) dump() {
    action := func(n *Node, prefix *[]uint8) {
        fmt.Printf("=%v: '%v'\n", n.word_count, string(*prefix))
    }
    n.preorder_walk([]uint8{}, action)
}


func main() {
    fmt.Println("\nWordCount")
}
