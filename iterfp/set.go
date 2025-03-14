package main

// Set holds a set of elements.
type Set[E comparable] struct {
    m map[E]struct{}
}

// New returns a new [Set].
func New[E comparable]() *Set[E] {
    return &Set[E]{m: make(map[E]struct{})}
}

// Add adds an element to a set.
func (s *Set[E]) Add(v E) {
    s.m[v] = struct{}{}
}

// Contains reports whether an element is in a set.
func (s *Set[E]) Contains(v E) bool {
    _, ok := s.m[v]
    return ok
}

// Union returns the union of two sets.
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
    r := New[E]()
    // Note for/range over internal Set field m.
    // We are looping over the maps in s1 and s2.
    for v := range s1.m {
        r.Add(v)
    }
    for v := range s2.m {
        r.Add(v)
    }
    return r
}