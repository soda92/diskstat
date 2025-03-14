package main

import (
	"testing"
)

func Test1(t *testing.T) {
	s := New[int]()
	s.Add(1)

	if !s.Contains(1) {
		t.Fatal("set add failed")
	}
}

func Test2(t *testing.T) {
	s := New[int]()
	s.Add(1)

	s2 := New[int]()
	s2.Add(1)

	s3 := Union(s, s2)

	if !s3.Contains(1) {
		t.Fatal("set union failed")
	}
}
