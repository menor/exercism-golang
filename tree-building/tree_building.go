package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

const rootId = 0

func Build(records []Record) (*Node, error) {

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	tree := make(map[int]*Node)

	for i, r := range records {
		noRootNode := i != rootId
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

	return tree[rootId], nil
}
