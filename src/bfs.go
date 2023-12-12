package main

type GQueue struct {
	data []int32
	size int32
}

func (q *GQueue) enqueue(e int32) {
	if q.data == nil {
		q.data = []int32{}
	}

	q.data = append(q.data, e)
	q.size += 1
}

func (q *GQueue) dequeue() int32 {
	if q.size == 0 {
		return -1
	}

	e := q.data[0]
	q.data = q.data[1:q.size]
	q.size -= 1

	return e
}

func (q *GQueue) peek() int32 {
	if q.size == 0 {
		return -1
	}

	return q.data[0]
}

func distances(n int32, m int32, edges [][]int32, s int32) []int32 {

	// create graph
	graph := make([]map[int32]bool, n) // should be a map but keys are from 1 to n -- so array would suffice
	for i := int32(0); i < m; i++ {
		v1 := edges[i][0] - 1
		v2 := edges[i][1] - 1
		if graph[v1] == nil {
			graph[v1] = map[int32]bool{}
		}
		if graph[v2] == nil {
			graph[v2] = map[int32]bool{}
		}

		graph[v1][v2] = true
		graph[v2][v1] = true
	}

	// fmt.Println("graph", graph)

	distances := make([]int32, n)
	visited := make([]int32, n)
	for i := int32(0); i < n; i++ {
		distances[i] = -1
	}

	queue := GQueue{}
	queue.enqueue(s - 1) // start is s: index = s-1
	currDistance := int32(0)

	for queue.size > 0 {
		head := queue.peek()
		children := graph[head]
		if len(children) > 0 {
			currDistance += 6
			for vertex := range children {
				if visited[vertex] == 0 {
					distances[vertex] = currDistance
					queue.enqueue(vertex)
				}
			}
		}

		visited[head] = 1
		queue.dequeue()
	}

	distances = append(distances[0:s-1], distances[s:n]...)

	return distances

}
