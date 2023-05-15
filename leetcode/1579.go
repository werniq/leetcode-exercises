func maxNumEdgesToRemove(n int, edges [][]int) int {
	aliceEdges := make([][]int, 0) // Edges accessible only by Alice
	bobEdges := make([][]int, 0)   // Edges accessible only by Bob
	sharedEdges := make([][]int, 0) // Edges accessible by both Alice and Bob

	for _, edge := range edges {
		edgeType, node1, node2 := edge[0], edge[1], edge[2]
		if edgeType == 1 {
			aliceEdges = append(aliceEdges, []int{node1, node2})
		} else if edgeType == 2 {
			bobEdges = append(bobEdges, []int{node1, node2})
		} else {
			sharedEdges = append(sharedEdges, []int{node1, node2})
		}
	}

	aliceComponents := make([]int, n+1) // Component IDs of nodes reachable by Alice
	bobComponents := make([]int, n+1)   // Component IDs of nodes reachable by Bob

	// Function to find the component ID of a node
	findComponentID := func(components []int, node int) int {
		for components[node] != node {
			components[node] = components[components[node]] // Path compression
			node = components[node]
		}
		return node
	}

	// Function to union two components
	union := func(components []int, node1, node2 int) {
		components[findComponentID(components, node1)] = findComponentID(components, node2)
	}

	removeCount := 0 // Counter for the removed edges

	// Assign initial component IDs to each node
	for i := 1; i <= n; i++ {
		aliceComponents[i] = i
		bobComponents[i] = i
	}

	// Process shared edges
	for _, edge := range sharedEdges {
		node1, node2 := edge[0], edge[1]
		aliceID := findComponentID(aliceComponents, node1)
		bobID := findComponentID(bobComponents, node1)
		// If Alice and Bob can already reach the nodes, this edge is redundant
		if aliceID == findComponentID(aliceComponents, node2) && bobID == findComponentID(bobComponents, node2) {
			removeCount++
			continue
		}
		union(aliceComponents, node1, node2)
		union(bobComponents, node1, node2)
	}

	aliceCopy := make([]int, n+1)
	copy(aliceCopy, aliceComponents)

	// Process Alice's edges
	for _, edge := range aliceEdges {
		node1, node2 := edge[0], edge[1]
		aliceID := findComponentID(aliceComponents, node1)
		if aliceID == findComponentID(aliceComponents, node2) {
			removeCount++
			continue
		}
		union(aliceComponents, node1, node2)
	}

	// Process Bob's edges
	for _, edge := range bobEdges {
		node1, node2 := edge[0], edge[1]
		bobID := findComponentID(bobComponents, node1)
		if bobID == findComponentID(bobComponents, node2) {
			removeCount++
			continue
		}
		union(bobComponents, node1, node2)
	}

	// If Alice or Bob cannot fully traverse the graph, return -1
	// ...

	// If Alice or Bob cannot fully traverse the graph, return -1
	// If Alice or Bob cannot fully traverse the graph, return -1
  for i := 2; i <= n; i++ {
	  if findComponentID(aliceComponents, i) != findComponentID(aliceComponents, 1) {
		  return -1
	  }
	  if findComponentID(bobComponents, i) != findComponentID(bobComponents, 1) {
		  return -1
	  }
  }

  aliceEdgesToRemove := 0
  bobEdgesToRemove := 0

  // Count the number of edges that can be removed by Alice
  for _, edge := range aliceEdges {
	  node1, node2 := edge[0], edge[1]
	  aliceID := findComponentID(aliceComponents, node1)
	    if aliceID != findComponentID(aliceComponents, node2) {
		    aliceEdgesToRemove++
		    union(aliceComponents, node1, node2)
	    }
  }

  // Count the number of edges that can be removed by Bob
  for _, edge := range bobEdges {
	  node1, node2 := edge[0], edge[1]
	  bobID := findComponentID(bobComponents, node1)
	  if bobID != findComponentID(bobComponents, node2) {
		    bobEdgesToRemove++
		    union(bobComponents, node1, node2)
	    }
    }

  return removeCount + aliceEdgesToRemove + bobEdgesToRemove
}

