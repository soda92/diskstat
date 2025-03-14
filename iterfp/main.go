package main

var v *Set[int]

func init() {
	v = New[int]()
	v.Add(1)
	v.Add(2)
	v.Add(3)

	v.Add(2)
}

func main() {
	PrintAllElementsPush(v)
}
