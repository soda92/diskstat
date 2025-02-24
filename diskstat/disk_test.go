package main

import (
	"testing"
)

func TestMapAccess(t *testing.T) {
	m := make(map[int]string)
	m[1] = "aaa"

	n := make(map[int]string)
	n[2] = "bbb"
	
	if !MapLeftNew(m, n, 1) {
		t.Fatalf("%s failed", "MapLeftNew")
	}

	if !MapRightNew(m, n, 2) {
		t.Fatalf("%s failed", "MapRightNew")
	}
}
