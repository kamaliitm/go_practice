package main

import (
	"fmt"
)

/**
	Service start
	a -> f, g
	b -> e
	c ->
	d -> e, f
	e -> f
	f -> c
	g ->

	['c', 'g', 'f', 'a', 'e', 'd']
*
*/

type Graph struct {
	services []string
	edges    map[string][]string
}

func (g *Graph) addEdge(s1 string, s2 string) {
	_, ok := g.edges[s1]
	if !ok {
		g.edges[s1] = []string{s2}
		return
	}

	g.edges[s1] = append(g.edges[s1], s2)
}

func executeServices(services []string, dependencies [][]string) {
	graph := &Graph{
		services: services,
		edges:    map[string][]string{},
	}

	graph.addEdge("a", "f")
	graph.addEdge("a", "g")
	graph.addEdge("b", "e")
	graph.addEdge("d", "e")
	graph.addEdge("d", "f")
	graph.addEdge("e", "f")
	graph.addEdge("f", "c")

	servicesOrder := []string{}
	visited := map[string]bool{}

	for _, service := range services {
		if !visited[service] {
			servicesOrderHelperDFS(graph, service, visited, &servicesOrder)
		}
	}

	fmt.Println(servicesOrder)

}

func servicesOrderHelperDFS(graph *Graph, currService string, visited map[string]bool, servicesOrder *[]string) {
	if visited[currService] {
		return
	}

	dependencies, ok := graph.edges[currService]
	if ok {
		for _, d := range dependencies {
			servicesOrderHelperDFS(graph, d, visited, servicesOrder)
		}
	}

	visited[currService] = true
	*servicesOrder = append(*servicesOrder, currService)
}

func main() {
	executeServices([]string{"a", "b", "c", "d", "e", "f", "g"}, nil)
}

// func main() {
// 	a := []int{1, 2, 3}
// 	b := append(a[:1], 10, 11, 12)
// 	// c := append(a[:1], 11)
// 	fmt.Printf("a=%v, b=%v\n", a, b)

// 	// a --> []int
// 	// b
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for _, n := range []int{1, 5, 3} {
// 		// n = 1 5 3
// 		n := n

// 		wg.Add(1)
// 		go func() {

// 			defer wg.Done()
// 			fmt.Printf("before %d \n", n)
// 			time.Sleep(time.Duration(n) * time.Millisecond)
// 			fmt.Printf("%d ", n)
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println()

// 	// 1 3 5
// 	// 3 3 3
// }
