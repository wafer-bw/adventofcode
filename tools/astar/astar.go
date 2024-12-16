package astar

import "container/heap"

type Pather interface {
	// Return reachable neighbors of this Pather.
	Neighbors() []Pather

	// Return the cost to move from this Pather to the given Pather.
	NeighborCost(to Pather) float64

	// Return the estimated cost to reach the given Pather from this
	// non-adjacent Pather.
	EstimatedCost(to Pather) float64
}

type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

type nodeMap map[Pather]*node

func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{pather: p}
		nm[p] = n
	}
	return n
}

func Path(start, end Pather) (path []Pather, cost float64, found bool) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)
	fromNode := nm.get(start)
	fromNode.open = true
	heap.Push(nq, fromNode)

	for {
		if nq.Len() == 0 { // no path found.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(end) { // found path to end.
			p := []Pather{}
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range current.pather.Neighbors() {
			cost := current.cost + current.pather.NeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}

			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.EstimatedCost(end)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}
