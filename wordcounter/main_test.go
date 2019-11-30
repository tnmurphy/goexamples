package main

import (
        "testing"
        "fmt"
)

func TestBasicInput1(t *testing.T) {
    var start Node;

    //test_string := "This is a test with repeated words - words like 'test' and 'This' and 'repeated'."
    test_string := "a aa b"
    fmt.Printf("\n%v\n",test_string)

    var inserter = start.get_inserter()
    for b := range(test_string) {
        inserter(test_string[b])
    }
    inserter(' ') // ensure the last word is flushed

    fmt.Printf("%v", )

    if len (start.edges) != 2 {
        t.Errorf("start.edges not 2")
    }
    if start.edges[0].char != 'a' {
        t.Errorf("first edge is not 'a' but '%v'", string(start.edges[0].char))
    }
    if start.edges[1].char != 'b' {
        t.Errorf("next edge is not 'b' but '%v'", string(start.edges[1].char))
    }

    //expected_dump := "{0 0 [{97 1 [{97 1 []}]} {98 1 []}]}"
    //if dump:=fmt.Sprintf("%v", start); dump != expected_dump {
    //    t.Errorf("struct was '%v' but should be '%v'", dump, expected_dump)
    //}
    fmt.Printf("\n")
    start.dump([]uint8{}, 0)

   }


func TestBasicInputRepeats(t *testing.T) {
    var start Node;

    //test_string := "This is a test with repeated words - words like 'test' and 'This' and 'repeated'."
    test_string := "a aa b aa b a"
    fmt.Printf("\n%v\n",test_string)

    var inserter = start.get_inserter()
    for b := range(test_string) {
        inserter(test_string[b])
    }
    inserter(' ') // ensure the last word is flushed

    fmt.Printf("%v", )

    if len (start.edges) != 2 {
        t.Errorf("start.edges not 2")
    }
    if start.edges[0].char != 'a' {
        t.Errorf("first edge is not 'a' but '%v'", string(start.edges[0].char))
    }
    if start.edges[1].char != 'b' {
        t.Errorf("next edge is not 'b' but '%v'", string(start.edges[1].char))
    }

    //expected_dump := "{0 0 [{97 1 [{97 1 []}]} {98 1 []}]}"
    //if dump:=fmt.Sprintf("%v", start); dump != expected_dump {
    //    t.Errorf("struct was '%v' but should be '%v'", dump, expected_dump)
    //}
    fmt.Printf("\n")
    start.dump([]uint8{}, 0)

   }
