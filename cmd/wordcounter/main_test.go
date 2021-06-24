package main

import (
	"fmt"
	"testing"
)

func test_string() string {
	return `
**The Project Gutenberg Etext of Moby Dick, by Herman Melville**
#3 in our series by Herman Melville

This Project Gutenberg version of Moby Dick is based on a combination
of the etext from the ERIS project at Virginia Tech and another from
Project Gutenberg's archives, as compared to a public-domain hard copy.

Copyright laws are changing all over the world, be sure to check
the copyright laws for your country before posting these files!!

Please take a look at the important information in this header.
We encourage you to keep this file on your own disk, keeping an
electronic path open for the next readers.  Do not remove this.


**Welcome To The World of Free Plain Vanilla Electronic Texts**

**Etexts Readable By Both Humans and By Computers, Since 1971**

*These Etexts Prepared By Hundreds of Volunteers and Donations*

Information on contacting Project Gutenberg to get Etexts, and
further information is included below.  We need your donations.


Title:  Moby Dick; or The Whale

Author:  Herman Melville

June, 2001  [Etext #2701]

**The Project Gutenberg Etext of Moby Dick, by Herman Melville**
******This file should be named moby10b.txt or moby10b.zip******
    `
}

func TestLongInput(t *testing.T) {

	topwords := NewTopwords(10)
	var start Node
	test_string := test_string()
	fmt.Printf("\n%v\n", test_string)

	var inserter = topwords.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}

	inserter(' ') // ensure the last word is flushed
	if topwords.uniquewordcount != 107 {
		t.Errorf("topwords.uniquewordcount not 107 but %v", topwords.uniquewordcount)
	}

	if len(topwords.start.children) != 23 {
		t.Errorf("start.children not 23 but %v", len(start.children))
	}
	if topwords.start.children[0].char != 't' {
		t.Errorf("first child is not 't' but '%v'", string(start.children[0].char))
	}
	if topwords.start.children[1].char != 'p' {
		t.Errorf("next edge is not 'p' but '%v'", string(start.children[1].char))
	}

	fmt.Printf("\n")
	start.dump()

}

func TestSortedOut(t *testing.T) {
    //    3
    //  2   2
    //2   1

	fmt.Println("\n--------------------------------TestSortedOut\n")
	//test_string := "this is one word this is another word and another word"
	test_string := "one this this one one two two three four four"

	topwords := NewTopwords(6)
	var inserter = topwords.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	fmt.Printf("\ncalculating most frequent---------------\n")
	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.wordheap))

	oldcount := uint(99999999999999999)
	for wc, ok := topwords.readword(); ok != false; wc, ok = topwords.readword() {
		fmt.Printf("%v\t%v\n", wc.count, wc.word)
		if oldcount < wc.count {
			fmt.Printf("ERROR: unordered item\n")
		}
		oldcount = wc.count
	}
}

func TestSortedOut2(t *testing.T) {
    // heap of this shape:
    //    3
    //  3   2
    //1

	fmt.Println("\n--------------------------------TestSortedOut2\n")
	test_string := "one this this this one two three three three"

	topwords := NewTopwords(6)
	var inserter = topwords.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	fmt.Printf("\ncalculating most frequent---------------\n")
	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.wordheap))

	oldcount := uint(99999999999999999)
	for wc, ok := topwords.readword(); ok != false; wc, ok = topwords.readword() {
		fmt.Printf("%v\t%v\n", wc.count, wc.word)
		if oldcount < wc.count {
			fmt.Printf("ERROR: unordered item\n")
		}
		oldcount = wc.count
	}
}

func TestSortedOut3(t *testing.T) {
	// A heap with this shape:
	//      4
	//  4      3
	//3   3  1   

	fmt.Println("\n--------------------------------TestSortedOut3\n")
	test_string := "one this this this one two three three three four five five five five four four one t t  t  t"

	topwords := NewTopwords(8)
	var inserter = topwords.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	fmt.Printf("\ncalculating most frequent---------------\n")
	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.wordheap))

	oldcount := uint(99999999999999999)
	for wc, ok := topwords.readword(); ok != false; wc, ok = topwords.readword() {
		fmt.Printf("%v\t%v\n", wc.count, wc.word)
		if oldcount < wc.count {
			fmt.Printf("ERROR: unordered item\n")
		}
		oldcount = wc.count
	}
}

func TestSortedOut4(t *testing.T) {
	// A heap with this shape:
	//      4
	//  4      3
	//3   3  1  1 

	fmt.Println("\n--------------------------------TestSortedOut4\n")
	test_string := "one this x this this one two three three three four five five five five four four one t t  t  t"

	topwords := NewTopwords(8)
	var inserter = topwords.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	fmt.Printf("\ncalculating most frequent---------------\n")
	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.wordheap))

	oldcount := uint(99999999999999999)
	for wc, ok := topwords.readword(); ok != false; wc, ok = topwords.readword() {
		fmt.Printf("%v\t%v\n", wc.count, wc.word)
		if oldcount < wc.count {
			fmt.Printf("ERROR: unordered item\n")
		}
		oldcount = wc.count
	}
}
