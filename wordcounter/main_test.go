package main

import (
	"fmt"
	"testing"
)

func TestBasicInput1(t *testing.T) {
	var start Node

	test_string := "a aa b"
	fmt.Printf("\n%v\n", test_string)

	var inserter = start.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	if len(start.children) != 2 {
		t.Errorf("start.children not 2")
	}
	if start.children[0].char != 'a' {
		t.Errorf("first edge is not 'a' but '%v'", string(start.children[0].char))
	}
	if start.children[1].char != 'b' {
		t.Errorf("next edge is not 'b' but '%v'", string(start.children[1].char))
	}

	//expected_dump := "{0 0 [{97 1 [{97 1 []}]} {98 1 []}]}"
	//if dump:=fmt.Sprintf("%v", start); dump != expected_dump {
	//    t.Errorf("struct was '%v' but should be '%v'", dump, expected_dump)
	//}
	fmt.Printf("\n")
	start.dump()

}

func TestBasicInputRepeats(t *testing.T) {
	var start Node

	test_string := "a aa b aa b a"
	fmt.Printf("\n%v\n", test_string)

	var inserter = start.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	if len(start.children) != 2 {
		t.Errorf("start.children not 2")
	}
	if start.children[0].char != 'a' {
		t.Errorf("first edge is not 'a' but '%v'", string(start.children[0].char))
	}
	if start.children[1].char != 'b' {
		t.Errorf("next edge is not 'b' but '%v'", string(start.children[1].char))
	}

	//expected_dump := "{0 0 [{97 1 [{97 1 []}]} {98 1 []}]}"
	//if dump:=fmt.Sprintf("%v", start); dump != expected_dump {
	//    t.Errorf("struct was '%v' but should be '%v'", dump, expected_dump)
	//}
	fmt.Printf("\n")
	start.dump()

}

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
	var start Node
	test_string := test_string()
	fmt.Printf("\n%v\n", test_string)

	var inserter = start.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	if len(start.children) != 23 {
		t.Errorf("start.children not 23 but %v", len(start.children))
	}
	if start.children[0].char != 't' {
		t.Errorf("first child is not 't' but '%v'", string(start.children[0].char))
	}
	if start.children[1].char != 'p' {
		t.Errorf("next edge is not 'p' but '%v'", string(start.children[1].char))
	}

	//expected_dump := "{0 0 [{97 1 [{97 1 []}]} {98 1 []}]}"
	//if dump:=fmt.Sprintf("%v", start); dump != expected_dump {
	//    t.Errorf("struct was '%v' but should be '%v'", dump, expected_dump)
	//}
	fmt.Printf("\n")
	start.dump()

}

func TestSortedOut(t *testing.T) {

	fmt.Println("\n--------------------------------TestSortedOut\n")
	test_string := "this is one word this is another word and another word"

	topwords := NewTopwords(6)
	var inserter = topwords.start.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.sortedwords))

	for t := range topwords.sortedwords {
		fmt.Printf("%v %v\n", topwords.sortedwords[t].count, topwords.sortedwords[t].word)
	}

}

func TestSortedOutLong(t *testing.T) {

	fmt.Println("\n--------------------------------TestSortedOutLong\n")
	test_string := test_string()

	topwords := NewTopwords(6)
	var inserter = topwords.start.get_inserter()
	for b := range test_string {
		inserter(test_string[b])
	}
	inserter(' ') // ensure the last word is flushed

	fmt.Printf("\ndump---------------\n")
	topwords.start.dump()

	topwords.most_frequent()

	fmt.Printf("Topwords returned %v words\n", len(topwords.sortedwords))

	for t := range topwords.sortedwords {
		fmt.Printf("%v %v\n", topwords.sortedwords[t].count, topwords.sortedwords[t].word)
	}

}
