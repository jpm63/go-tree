// Package tree implements a generic tree, providing access and manipulation
// functions for both the tree nodes themselves as well as the data
// contained therein.
package tree

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
)

// Tree is a generic tree.
type Tree[T any] struct {
	data     T
	parent   *Tree[T]
	children []*Tree[T]
}

// New initializes a new tree for use.
func New[T any](data T) *Tree[T] {
	return &Tree[T]{
		data: data,
	}
}

// Parent returns the parent of t.
func (t *Tree[T]) Parent() *Tree[T] {
	return t.parent
}

// Children returns the children of t.
func (t *Tree[T]) Children() []*Tree[T] {
	return t.children
}

// Child returns the child specified by selector.
func (t *Tree[T]) Child(selector func(tree *Tree[T]) bool) *Tree[T] {
	for _, child := range t.children {
		if selector(child) {
			return child
		}
	}

	return nil
}

// Insert inserts tree as a child of t.
func (t *Tree[T]) Insert(tree *Tree[T]) {
	tree.parent = t
	t.children = append(t.children, tree)
}

// Remove removes tree, if it exists, from the children of t.
func (t *Tree[T]) Remove(tree *Tree[T]) {
	for i := range t.children {
		if t.children[i] == tree {
			t.children[i].parent = nil
			t.children = append(t.children[:i], t.children[i+1:]...)
			return
		}
	}
}

// Sort sorts the children of t according to less.
func (t *Tree[T]) Sort(less func(int, int) bool) {
	sort.Slice(t.children, less)
}

// Walk executes fn on each node in tree according to
// the specified search strategy.
func (t *Tree[T]) Walk(fn func(tree *Tree[T]), search SearchStrategy[T]) {
	for next := range search(t) {
		fn(next)
	}
}

// Find returns the first match in the tree according
// to the specified search strategy.
func (t *Tree[T]) Find(selector func(tree *Tree[T]) bool, search SearchStrategy[T]) *Tree[T] {
	for next := range search(t) {
		if selector(next) {
			return next
		}
	}

	return nil
}

// FindAll returns all matches in the tree according
// to the specified search strategy.
func (t *Tree[T]) FindAll(selector func(tree *Tree[T]) bool, search SearchStrategy[T]) []*Tree[T] {
	var out []*Tree[T]
	for next := range search(t) {
		if selector(next) {
			out = append(out, next)
		}
	}

	return out
}

// Data returns the data contained in t.
func (t *Tree[T]) Data() T {
	return t.data
}

func (t *Tree[T]) ReplaceData(data T) {
	t.data = data
}

// ParentData returns the data contained in the parent of t.
func (t *Tree[T]) ParentData() T {
	if t.parent == nil {
		return *new(T)
	}

	return t.parent.data
}

// ChildrenData returns the data contained in the children of t.
func (t *Tree[T]) ChildrenData() []T {
	return dataSlice(t.children)
}

// ChildData returns the data contained in the
// child specified by selector.
func (t *Tree[T]) ChildData(selector func(tree *Tree[T]) bool) T {
	c := t.Child(selector)
	if c == nil {
		return *new(T)
	}

	return c.data
}

// InsertData inserts data as a child of t.
func (t *Tree[T]) InsertData(data T) {
	child := New(data)
	t.Insert(child)
}

// RemoveData removes the data, if it exists, from the children
// of t. Existence is checked by reflect.DeepEqual
func (t *Tree[T]) RemoveData(data T) {
	for i := range t.children {
		if reflect.DeepEqual(data, t.children[i].data) {
			t.Remove(t.children[i])
		}
	}
}

// Find returns the data contained in the first match in
// the tree according to the specified search strategy.
func (t *Tree[T]) FindData(selector func(tree *Tree[T]) bool, search SearchStrategy[T]) T {
	return t.Find(selector, search).data
}

// FindAll returns the data coantined in all matches in
// the tree according to the specified search strategy.
func (t *Tree[T]) FindAllData(selector func(tree *Tree[T]) bool, search SearchStrategy[T]) []T {
	return dataSlice(t.FindAll(selector, search))
}

// Print prints t to stdout.
func (t *Tree[T]) Print() {
	t.fprint(os.Stdout, "", "", 0)
}

// Print pretty-prints t to stdout according to prefix and indent.
func (t *Tree[T]) PrintIndent(prefix, indent string) {
	t.fprint(os.Stdout, prefix, indent, 0)
}

// Fprint prints t to the w.
func (t *Tree[T]) Fprint(w io.Writer) {
	t.fprint(w, "", "", 0)
}

// FprintIndent pretty-prints t to w according to prefix and indent.
func (t *Tree[T]) FprintIndent(w io.Writer, prefix, indent string) {
	t.fprint(w, prefix, indent, 0)
}

func (t *Tree[T]) fprint(w io.Writer, prefix, indent string, level int) {
	var p, i string
	for j := 0; j < level; j++ {
		p = prefix
		i += indent
	}

	fmt.Fprintf(w, "%s%s%v\n", p, i, t.data)
	for _, child := range t.children {
		child.fprint(w, prefix, indent, level+1)
	}
}

func dataSlice[T any](trees []*Tree[T]) []T {
	var out []T
	for _, tree := range trees {
		out = append(out, tree.data)
	}

	return out
}
