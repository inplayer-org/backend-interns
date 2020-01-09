package tree

import (
	"errors"
	"sort"
)

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(r []Record) (*Node, error) {
	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})
	nodes := make(map[int]*Node)
	for i, r := range r {
		if r.ID != i || r.ID < r.Parent || (r.ID != 0 && r.ID == r.Parent) {
			return nil, errors.New("Invalid record")
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
