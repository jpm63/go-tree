package tree_test

import (
	"testing"

	"github.com/jpm63/go-tree"
)

func TestNew(t *testing.T) {
	v := tree.New(5)
	if v == nil {
		t.Errorf("got nil, want value")
	}
}

func TestData(t *testing.T) {
	want := 5
	v := tree.New(want)
	got := v.Data()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
