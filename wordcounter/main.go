package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// A tree node, the character it represents, how many words
// ended a this character, any characters that have followed
// on from this one as part of longer words. Not quite a
// Radix tree because sequences of letters never get compressed
// into one node
type Node struct {
	char       uint8
	word_count uint
	children   []*Node
}

// A word and a count of how many.
type Wordcount struct {
	count uint
	word  string
}

// A structure representing the top N words.
// After we find the words and counts we insert then into
// a heap, wordheap, and pull them out one by one in order.
type Topwords struct {
	wordheap        []*Wordcount
	uniquewordcount uint
	countmax        uint
	start           Node
}

// Filter characters and translate upper to lower
// case at the same time. Non-english characters
// are convered to spaces. Obviously this would be
// unsuitable for many other languages and something
// much more sophisticated could be implemented.
func filterToLower(b byte) byte {

	if b >= 'A' && b <= 'Z' {
		return 'a' + (b - 'A')
	} else if b >= 'a' && b <= 'z' {
		return b
	} else {
		return ' ' // binary etc converted to space
	}
}

// Returns a closure that can insert characters into
// the tree. Usually one would call this on the
// start member of Topwords.
func (tw *Topwords) get_inserter() func(uint8) {
	current := &tw.start
	inword := false

	return func(b uint8) {
		b = filterToLower(b)

		if b == ' ' {
			//fmt.Printf("_")
			if inword {
				if current.word_count == 0 {
					tw.uniquewordcount += 1
				}
				current.word_count += 1
				current = &tw.start // Start at the beginning again
			}
			inword = false
			return
		}
		//fmt.Printf("%v", string(b))
		inword = true
		children := &(current.children)
		for e := range *children {
			if (*children)[e].char == b {
				//fmt.Printf("<")
				current = (*children)[e]
				return
			}
		}
		//fmt.Printf("^")
		new_node := &Node{b, 0, nil}
		current.children = append(current.children, new_node)
		current = new_node
		// fmt.Printf("\n+%v\n", current.children)
		return

	}

}

// This adds a single wordcount struct into the Topwords heap
// and rebalances the heap.
func (tw *Topwords) addwordcount(w *Wordcount) {
	max := cap(tw.wordheap)
	if len(tw.wordheap) < max {
		tw.wordheap = append(tw.wordheap, w)
	} else {
		if tw.wordheap[max-1].count >= w.count {
			return // No need to consider this word
		}
		tw.wordheap[max-1] = w
		// So what happens to the item at the end of the list of topwords?
		// we replaced it without a comparison! Well, we use the end of the
		// slice as a sort of "entrance/exit hall" and don't intend to show
		// what's in there finally anyhow - so whatever's "in" there is not
		// important as it won't make the final "cut" anyhow.
	}

	// Rebalance the heap

	//fmt.Printf("rebalancing heap after adding %v\n", w)
	//tw.dumpheap()
	for f := len(tw.wordheap) - 1; f > 0; {
		parent := (f - 1) >> 1
		//fmt.Printf("rebalancing: parent %v f %v\n", parent, f)
		if tw.wordheap[f].count > tw.wordheap[parent].count {
			tmp := tw.wordheap[f]
			tw.wordheap[f] = tw.wordheap[parent]
			tw.wordheap[parent] = tmp
		}
		f = parent
	}
	//fmt.Printf("---after rebalance--\n")
	//tw.dumpheap()
}
func (tw *Topwords) dumpheap() {
	for i := range tw.wordheap {
		fmt.Printf("%v %v\n", i, *tw.wordheap[i])
	}
}

// Returns a word from the top of the wordheap. This is a destructive
// operation and the heap is empty after the last word is read since
// taking an item out of the top forces us to shift others up to
// maintain the tree.
func (tw *Topwords) readword() (w *Wordcount, ok bool) {
	ok = true
	if len(tw.wordheap) == 0 {
		return nil, false
	}
	w = tw.wordheap[0]
	//fmt.Printf("heap: top= %v\n", tw.wordheap[0])
	// tw.dumpheap()

	heaplen := len(tw.wordheap) - 1
	if heaplen >= 1 {
		// put the wc from the back of the heap at the head
		tw.wordheap[0] = tw.wordheap[heaplen]

		// If there is only 1 wc remaining then we've already finished.
		if heaplen > 1 {
			// ... the tree has at least 2 wc's so it may be unordered.
			left_child := 1
			right_child := 2

			// fmt.Printf("heap: len %v\n", heaplen)
			parent := 0
			for parent < heaplen-1 {
				// Handle the case of a leaf node or a node
				// with only a left child

				//fmt.Printf("heap: loop top\n")
				if right_child >= heaplen {
					//fmt.Printf("heap: right child > heaplen  %v > %v\n", right_child, heaplen)
					if left_child < heaplen { // Not a leaf node
						// No right child exists so compare left one, swap(?) and stop
						//fmt.Printf("heap: noright, moving left len %v l%v P%v\n", heaplen, left_child, parent)

						if tw.wordheap[parent].count < tw.wordheap[left_child].count {
							tmp := tw.wordheap[parent]
							tw.wordheap[parent] = tw.wordheap[left_child]
							tw.wordheap[left_child] = tmp
						}
					}
					break
				}

				//fmt.Printf("heap: parent: %v=%v left: =%v right: =%v\n", parent, tw.wordheap[parent].count, tw.wordheap[left_child].count, tw.wordheap[right_child].count)
				if tw.wordheap[left_child].count > tw.wordheap[right_child].count {
					// the left child is bigger so swap it up if it is bigger than the parent.
					//fmt.Printf("heap: left bigger than right\n")
					if tw.wordheap[parent].count < tw.wordheap[left_child].count {
						//fmt.Printf("heap: swap parent and left %v\n",tw.wordheap[parent].count)
						tmp := tw.wordheap[parent]
						tw.wordheap[parent] = tw.wordheap[left_child]
						tw.wordheap[left_child] = tmp
					}
					parent = left_child
				} else {
					// fmt.Printf("heap: left not bigger than right so choose right\n")
					if tw.wordheap[right_child].count > tw.wordheap[parent].count {
						// left child moves up if it's bigger
						// fmt.Printf("heap: swap parent and right %v\n",tw.wordheap[parent].count)
						tmp := tw.wordheap[parent]
						tw.wordheap[parent] = tw.wordheap[right_child]
						tw.wordheap[right_child] = tmp
					}
					parent = right_child
				}

				left_child = parent<<1 + 1
				right_child = left_child + 1
				//fmt.Printf("heap: parent now %v=%v heaplen-1 is %v - parent < heaplen -1=%v\n", parent, tw.wordheap[parent].count, heaplen - 1, parent < heaplen -1)
			}
			// fmt.Printf("heap: end parent is at %v heaplen-1 is %v\n", parent, heaplen - 1)
			// tw.dumpheap()
		}

	}

	tw.wordheap = tw.wordheap[:heaplen]
	// fmt.Printf("heap: resizing %v\n", tw.wordheap)

	return w, ok
}

func NewTopwords(countmax uint) *Topwords {
	var tw Topwords
	tw.countmax = countmax
	return &tw
}

func (n *Node) preorder_walk(prefix []uint8, action func(n *Node, prefix *[]uint8)) {
	if n.char == 0 { // step over the first node if it's a start node
		// fmt.Printf("\nSTART depth %:v prefix '%v'\n", depth, prefix)
	} else {
		prefix = append(prefix, n.char)
		if n.word_count > 0 {
			action(n, &prefix)
		}
	}
	for e := range n.children {
		n.children[e].preorder_walk(prefix, action)
	}
}

// Walks the tree and adds words into the heap
func (tw *Topwords) most_frequent() {

	// fmt.Printf("Unique word count is: %v\n", tw.uniquewordcount)

	// This probably could be sized as approximately log2(tw.countmax)+1 and save more time and memory.
	// It needs to be deep enough that it won't lose high wordcounts that are at the end of the heap but
	// it doesn't really need to be big enough to hold all the unique words.
	tw.wordheap = make([]*Wordcount, 0, tw.uniquewordcount+1)

	action := func(n *Node, prefix *[]uint8) {
		w := new(Wordcount)
		w.count = n.word_count
		w.word = string(*prefix)

		tw.addwordcount(w)

	}

	tw.start.preorder_walk([]uint8{}, action)
}

func (n *Node) dump() {
	action := func(n *Node, prefix *[]uint8) {
		fmt.Printf("%v\t'%v'\n", n.word_count, string(*prefix))
	}
	n.preorder_walk([]uint8{}, action)
}

// Reads input from all files specified as arguments on the commmandline
// or if none are specified then it reads from the stdin. 
func main() {
	topwords := NewTopwords(20)

	var inserter = topwords.get_inserter()

	countwords := func(f io.Reader) {
		reader := bufio.NewReader(f)
		for char, err := reader.ReadByte(); err == nil; char, err = reader.ReadByte() {
			inserter(char)
		}
		inserter(' ') // ensure we flush the last word.
	}

	if len(os.Args) < 2 { // read from stdin if there are no files on the commandline
		countwords(os.Stdin)
	} else {
		for i := 1; i <  len(os.Args); i++ {
			f, err := os.Open(os.Args[i])
			if err != nil {
				log.Fatal("Couldn't open %v\n", os.Args[i])
			}
			countwords(f)
			f.Close()
		}
	}

	// Find the most common words for all files specified on the commandline.
	topwords.most_frequent()

	var count = 0
	for wc, ok := topwords.readword(); ok != false; wc, ok = topwords.readword() {
		fmt.Printf("%7v\t%v\n", wc.count, wc.word)
		count += 1
		if count == 20 { // Hardcoded, not nice!
			break
		}
	}

}
