package main

import "fmt"

func (s *Set[E]) Push(f func(E) bool) {
    for v := range s.m {
        if !f(v) {
            return
        }
    }
}

func PrintAllElementsPush[E comparable](s *Set[E]) {
    s.Push(func(v E) bool {
        fmt.Println(v)
        return true
    })
}