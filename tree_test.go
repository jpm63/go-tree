package tree_test

import (
	"reflect"
	"testing"

	"github.com/jpm63/go-tree"
)

func TestNew(t *testing.T) {
	want := 5
	root := tree.New(want)
	if root == nil {
		t.Errorf("got nil, want tree")
	}

	got := root.Data()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParent(t *testing.T) {
	want := tree.New(5)
	child := tree.New(10)
	want.Insert(child)

	got := child.Parent()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChildren(t *testing.T) {
	want1 := tree.New(5)
	want2 := tree.New(10)
	want := []*tree.Tree[int]{want1, want2}

	root := tree.New(0)
	root.Insert(want1)
	root.Insert(want2)

	got := root.Children()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChild(t *testing.T) {
	root := tree.New(5)
	want := tree.New(10)
	root.Insert(want)

	t.Run("match", func(t *testing.T) {
		got := root.Child(func(z *tree.Tree[int]) bool {
			return z.Data() == 10
		})

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		got := root.Child(func(z *tree.Tree[int]) bool {
			return z.Data() == 0
		})

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}

func TestInsert(t *testing.T) {
	root := tree.New(5)
	want := tree.New(10)
	root.Insert(want)
	got := root.Children()[0]
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	root := tree.New(5)
	child := tree.New(10)
	root.Insert(child)
	root.Remove(child)

	if got := child.Parent(); got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

	if got := len(root.Children()); got != 0 {
		t.Errorf("got %v, want %v", got, 0)
	}
}

func TestSort(t *testing.T) {
	c1 := tree.New(5)
	c2 := tree.New(10)
	c3 := tree.New(15)

	root := tree.New(0)
	root.Insert(c3)
	root.Insert(c2)
	root.Insert(c1)

	want := []*tree.Tree[int]{c1, c2, c3}
	root.Sort(func(i, j int) bool {
		return root.Children()[i].Data() < root.Children()[j].Data()
	})

	got := root.Children()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWalk(t *testing.T) {
	root := tree.New(0)
	c1 := tree.New(5)
	root.Insert(c1)

	var got []*tree.Tree[int]
	root.Walk(func(tree *tree.Tree[int]) {
		got = append(got, tree)
	}, tree.BreadthFirstStrategy[int])

	want := []*tree.Tree[int]{root, c1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFind(t *testing.T) {
	want := tree.New(0)
	c1 := tree.New(5)
	want.Insert(c1)

	t.Run("match", func(t *testing.T) {
		got := want.Find(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 0
		}, tree.BreadthFirstStrategy[int])

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		got := want.Find(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 7
		}, tree.BreadthFirstStrategy[int])

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}

func TestFindAll(t *testing.T) {
	root := tree.New(0)
	c1 := tree.New(5)
	c2 := tree.New(5)
	root.Insert(c1)
	root.Insert(c2)

	t.Run("match", func(t *testing.T) {
		want := []*tree.Tree[int]{c1, c2}
		got := root.FindAll(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 5
		}, tree.BreadthFirstStrategy[int])

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		var want []*tree.Tree[int]
		got := root.FindAll(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 7
		}, tree.BreadthFirstStrategy[int])

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestData(t *testing.T) {
	want := 5
	root := tree.New(want)
	got := root.Data()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReplaceData(t *testing.T) {
	tests := []struct {
		want any
		init any
	}{
		{5, 0},
		{"asdf", "fdsa"},
		{[]int{1, 2, 3}, []int{0}},
		{[]int{1, 2, 3}, nil},
		{struct{ s string }{"asdf"}, struct{ s string }{"fdsa"}},
	}
	for _, tt := range tests {
		root := tree.New(tt.init)
		root.ReplaceData(tt.want)
		got := root.Data()

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}

func TestParentData(t *testing.T) {
	t.Run("parent", func(t *testing.T) {
		want := 5
		root := tree.New(want)
		child := tree.New(10)
		root.Insert(child)

		got := child.ParentData()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noParent", func(t *testing.T) {
		root := tree.New(5)
		want := 0
		got := root.ParentData()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestChildrenData(t *testing.T) {
	want1 := 5
	want2 := 10
	r1 := tree.New(want1)
	r2 := tree.New(want2)
	want := []int{want1, want2}

	root := tree.New(0)
	root.Insert(r1)
	root.Insert(r2)

	got := root.ChildrenData()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChildData(t *testing.T) {
	want := 10
	root := tree.New(5)
	child := tree.New(want)
	root.Insert(child)

	t.Run("match", func(t *testing.T) {
		got := root.ChildData(func(z *tree.Tree[int]) bool {
			return z.Data() == 10
		})

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		got := root.ChildData(func(z *tree.Tree[int]) bool {
			return z.Data() == 0
		})

		if got != 0 {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}

func TestInsertData(t *testing.T) {
	want := 10
	root := tree.New(5)
	root.InsertData(want)
	got := root.ChildrenData()[0]
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemoveData(t *testing.T) {
	want := 10
	root := tree.New(5)
	child := tree.New(want)
	root.Insert(child)
	root.RemoveData(want)

	if got := child.Parent(); got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

	if got := len(root.Children()); got != 0 {
		t.Errorf("got %v, want %v", got, 0)
	}
}

func TestFindData(t *testing.T) {
	want := 0
	root := tree.New(want)
	c1 := tree.New(5)
	root.Insert(c1)

	t.Run("match", func(t *testing.T) {
		got := root.FindData(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 0
		}, tree.BreadthFirstStrategy[int])

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		got := root.Find(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 7
		}, tree.BreadthFirstStrategy[int])

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}

func TestFindAllData(t *testing.T) {
	root := tree.New(0)
	c1 := tree.New(5)
	c2 := tree.New(5)
	root.Insert(c1)
	root.Insert(c2)

	t.Run("match", func(t *testing.T) {
		want := []int{5, 5}
		got := root.FindAllData(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 5
		}, tree.BreadthFirstStrategy[int])

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("noMatch", func(t *testing.T) {
		var want []int
		got := root.FindAllData(func(tree *tree.Tree[int]) bool {
			return tree.Data() == 7
		}, tree.BreadthFirstStrategy[int])

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
