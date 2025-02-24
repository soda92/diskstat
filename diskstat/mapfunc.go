package main

import "log"

type Val interface {
	disk_usage | string | int
}

func MapLeftNew[K comparable, V Val](left, right map[K]V, key K) bool {
	_, exists := right[key]
	return !exists
}

func MapRightNew[K comparable, V Val](left, right map[K]V, key K) bool {
	_, exists := left[key]
	return !exists
}

func IsDiskNew(arr []disk_usage, disk string) bool {
	for _, v := range arr {
		if v.disk_path == disk {
			return false
		}
	}
	return true
}

func DiskIndex(arr []disk_usage, key string) int {
	for i, v := range arr {
		if v.disk_path == key {
			return i
		}
	}
	log.Fatal("index not found")
	return -1
}

func DiskPaths(arr []disk_usage) []string {
	var ret []string
	for _, v := range arr {
		path := v.disk_path
		ret = append(ret, path)
	}
	return ret
}

func FindIndex(arr []string, elem string) int {
	if len(arr) == 0 {
		return 0
	}
	if elem > arr[0] {
		return 0
	}
	for i, v := range arr {
		if elem > v {
			continue
		} else {
			return i - 1
		}
	}
	return len(arr)
}

func InsertOrdered[V Val](arr []V, elem V, index int) []V {
	arr = append(arr, elem)
	if index == len(arr)-1 {
		return arr
	} else {
		copy(arr[index+1:], arr[index:])
		arr[index] = elem
		return arr
	}
}
