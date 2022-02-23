package tree_test

import (
	"reflect"
	"testing"

	"github.com/jpm63/go-tree"
)

func TestDepthFirstSearch(t *testing.T) {
	root := tree.New(0)
	c1 := tree.New(1)
	c2 := tree.New(2)
	c3 := tree.New(3)

	c1.Insert(c3)
	root.Insert(c1)
	root.Insert(c2)

	want := []*tree.Tree[int]{root, c1, c3, c2}

	var got []*tree.Tree[int]
	for a := range tree.DepthFirstStrategy(root) {
		got = append(got, a)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	root := tree.New(0)
	c1 := tree.New(1)
	c2 := tree.New(2)
	c3 := tree.New(3)

	c1.Insert(c3)
	root.Insert(c1)
	root.Insert(c2)

	want := []*tree.Tree[int]{root, c1, c2, c3}

	var got []*tree.Tree[int]
	for a := range tree.BreadthFirstStrategy(root) {
		got = append(got, a)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
