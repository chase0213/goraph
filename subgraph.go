package goraph

// GenerageSubgraph generates and returns a sugraph from given ids.
func GenerateSubgraph(g Graph, ids []ID) (Graph, error) {
	subgraph := NewGraph()

	// adjacency matrix
	adj := make(map[ID]map[ID]bool)

	// First of all, create a graph consisted of nodes without edges,
	// because AddEdge function will return error if
	// given node ID is not registered(unsafe).
	for _, id := range ids {
		node, err := g.GetNode(id)
		if err != nil {
			return nil, err
		}

		// even if node has edges or not, add node to the subgraph
		subgraph.AddNode(node)

		targets, err := g.GetTargets(id)
		if err != nil {
			continue
		}

		for target := range targets {
			if _, ok := adj[id]; ok {
				adj[id][target] = true
			} else {
				// must create an empty map before assignment
				adj[id] = map[ID]bool{}
				adj[id][target] = true
			}
		}
	}

	// Finally, add edges according to adjacency matrix
	for s, ts := range adj {
		for t, adjacent := range ts {
			if adjacent {
				weight, err := g.GetWeight(s, t)
				if err != nil {
					subgraph.AddEdge(s, t, 0.0)
				} else {
					subgraph.AddEdge(s, t, weight)
				}
			}
		}
	}

	return subgraph, nil
}
