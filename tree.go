package tree

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
)

type Tree[T any] struct {
	data     T
	parent   *Tree[T]
	children []*Tree[T]
}

func New[T any](item T) *Tree[T] {
	return &Tree[T]{data: item}
}

func (t *Tree[T]) Data() T {
	return t.data
}

func (t *Tree[T]) Parent() *Tree[T] {
	return t.parent
}

func (t *Tree[T]) ParentData() (T, bool) {
	if t.parent == nil {
		return *new(T), false
	}

	return t.parent.data, true
}

func (t *Tree[T]) Children() []*Tree[T] {
	return t.children
}

func (t *Tree[T]) ChildrenData() []T {
	var out []T
	for _, child := range t.children {
		out = append(out, child.data)
	}

	return out
}

func (t *Tree[T]) Child(fn func(T) bool) *Tree[T] {
	for _, child := range t.children {
		if fn(child.data) {
			return child
		}
	}

	return nil
}

func (t *Tree[T]) ChildData(fn func(T) bool) (T, bool) {
	for _, child := range t.children {
		if fn(child.data) {
			return child.data, true
		}
	}

	return *new(T), false
}

func (t *Tree[T]) InsertTree(child *Tree[T]) {
	child.parent = t
	t.children = append(t.children, child)
}

func (t *Tree[T]) Insert(item T) {
	child := New(item)
	child.parent = t
	t.children = append(t.children, child)
}

func (t *Tree[T]) RemoveTree(child *Tree[T]) {
	for i, c := range t.children {
		if c == child {
			t.children = append(t.children[:i], t.children[i+1:]...)
		}
	}
}

func (t *Tree[T]) Remove(item T) {
	for i, c := range t.children {
		if reflect.DeepEqual(c.data, item) {
			t.children = append(t.children[:i], t.children[i+1:]...)
		}
	}
}

func (t *Tree[T]) SortChildren(less func(i, j int) bool) {
	sort.Slice(t.children, less)
}

func (t *Tree[T]) DepthFirstSearch(fn func(T) bool) *Tree[T] {
	if fn(t.data) {
		return t
	}

	for _, child := range t.children {
		v := child.DepthFirstSearch(fn)
		if v != nil {
			return v
		}
	}

	return nil
}

func (t *Tree[T]) DepthFirstSearchData(fn func(T) bool) (T, bool) {
	v := t.DepthFirstSearch(fn)
	if v == nil {
		return *new(T), false
	}

	return v.data, true
}

func (t *Tree[T]) Print() {
	t.fprint(os.Stdout, "", "", 0)
}

func (t *Tree[T]) PrintCustom(prefix, indent string) {
	t.fprint(os.Stdout, prefix, indent, 0)
}

func (t *Tree[T]) Fprint(w io.Writer) {
	t.fprint(w, "", "", 0)
}

func (t *Tree[T]) FprintCustom(w io.Writer, prefix, indent string) {
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
