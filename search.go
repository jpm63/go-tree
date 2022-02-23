package tree

// SearchStrategy defines a manner in which to search a
// generic tree
type SearchStrategy[T any] func(root *Tree[T]) <-chan *Tree[T]

// DepthFirstStrategy searches root in a depth-first manner.
func DepthFirstStrategy[T any](root *Tree[T]) <-chan *Tree[T] {
	out := make(chan *Tree[T])
	go func() {
		depthFirst(root, out)
		close(out)
	}()

	return out
}

func depthFirst[T any](root *Tree[T], out chan *Tree[T]) {
	out <- root

	for _, child := range root.children {
		depthFirst(child, out)
	}
}

// BreadthFirstStrategy searches root in a breadth-first manner.
func BreadthFirstStrategy[T any](root *Tree[T]) <-chan *Tree[T] {
	out := make(chan *Tree[T])
	go func() {
		breadthFirst(root, out)
		close(out)
	}()
	return out
}

func breadthFirst[T any](root *Tree[T], out chan *Tree[T]) {
	out <- root
	next := root.children
	for len(next) != 0 {
		var n []*Tree[T]
		for i := range next {
			out <- next[i]
			n = append(n, next[i].children...)
		}

		next = n
	}
}
