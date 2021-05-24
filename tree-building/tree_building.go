package tree

import (
	"fmt"
	"sort"
)

// Record is an unsorted post
type Record struct {
	ID     int
	Parent int
}

// Node represents a a record once inserted in the tree
type Node struct {
	ID       int
	Children []*Node
}

const rootID = 0

// Build takes a slice of records an constructs a tree based
// on the information those records contain
func Build(records []Record) (*Node, error) {

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	tree := make(map[int]*Node)

	for i, r := range records {
		noRootNode := i != rootID
		rootNodeHasParent := r.Parent > i
		noPreviousNode := i != r.ID
		nodeIsItsOwnParent := (noRootNode && r.Parent == i)

		if noPreviousNode || rootNodeHasParent || nodeIsItsOwnParent {
			return nil, fmt.Errorf("invalid record: %+v", r)
		}
		n := &Node{ID: r.ID}
		if r.ID > 0 {
			tree[r.Parent].Children = append(tree[r.Parent].Children, n)
		}
		tree[i] = n
	}

	return tree[rootID], nil
}
