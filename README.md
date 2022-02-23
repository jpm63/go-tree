# Overview
Provides an implementation of a generic tree. Additional tree implementations to come over time.

# Example
Requires Go 1.18

``` go
import "github.com/jpm63/go-tree"

func main() {
    root := tree.New(0)
    child1 := tree.New(1)
    child2 := tree.New(2)
    grandChild1 := tree.New(3)
    grandChild2 := tree.New(4)

    child1.Insert(grandChild1)
    child1.Insert(grandChild2)
    root.Insert(child1)
    root.Insert(child2)

    root.Children() // Returns [child1, child2]
    root.Sort(func(i, j int) bool {
        return root.Children()[i].Data() > root.Children()[j].Data()
    })

    root.ChildrenData() // Returns [2, 1]

    f := func(t *tree.Tree[int]) bool {
        return t.Data() % 2 == 0
    }
    root.Find(f, tree.BreadthFirstStrategy[int]) // Returns child2
    root.FindData(f, tree.DepthFirstStrategy[int]) // Returns 4 (grandChild2)
}
```

# Contributions
All contributions are welcome and appreciated.