package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		indeg = make([]int, numCourses)
		res   []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		res = append(res, v)
		for _, u := range edges[v] {
			indeg[u]--
			if indeg[u] == 0 {
				q = append(q, u)
			}
		}
	}
	return len(res) == numCourses
}
