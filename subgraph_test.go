package goraph

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateSubgraph(t *testing.T) {
	g1 := NewGraph()
	fmt.Println("g1:", g1.String())

	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g2, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("g2:", g2.String())

	// test only if subgraph is generatable from empty graph
	sg1IDs := make([]ID, 0)
	sg1, err := GenerateSubgraph(g1, sg1IDs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("sg1:", sg1.String())

	// n := #nodes; n = 0
	sg2Nodes := g2.GetNodes()
	sg2IDs := make([]ID, 0, len(sg2Nodes))
	for k := range sg2Nodes {
		sg2IDs = append(sg2IDs, k)
	}
	sg2, err := GenerateSubgraph(g2, sg2IDs)
	if err != nil {
		t.Fatal(err)
	}

	subnodes := sg2.GetNodes()
	if len(subnodes) != len(sg2IDs) {
		t.Fatalf("the number of nodes in the subgraph must be %d but %d", len(sg2IDs), len(subnodes))
	}

	// n := #nodes; 0 < n < max(nodes)
	// n := #nodes; n = max(nodes)
	// n := #nodes; n > max(nodes)
}
