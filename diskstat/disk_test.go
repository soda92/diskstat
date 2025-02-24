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

func TestArrayEqual(t *testing.T) {
	arr1 := []string{"a", "b"}
	arr2 := []string{"a"}
	if ArrayEqual(arr1, arr2) {
		t.Fatal("array should not equal")
	}
	arr2 = append(arr2, "c")
	if ArrayEqual(arr1, arr2) {
		t.Fatal("array should not equal")
	}
	arr2[1] = "b"
	if !ArrayEqual(arr1, arr2) {
		t.Fatal("array should equal")
	}
}

func TestInsert(t *testing.T) {
	arr := []string{"a", "c", "d"}
	elem := "b"
	index := FindIndex(arr, elem)
	if index != 1 {
		t.Fatal("FindIndex failed")
	}
	new := InsertOrdered(arr, elem, index)
	expected := []string{"a", "b", "c", "d"}
	if !ArrayEqual(expected, new) {
		t.Fatal("insert failed")
	}
}
