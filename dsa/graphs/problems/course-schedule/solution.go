// Package courseschedule solves: determine whether all courses can be
// finished given a list of prerequisite pairs.
package courseschedule

// CanFinish reports whether all numCourses courses can be completed given
// prerequisites, where each pair [a, b] means "b must be taken before a".
// Runs in O(V + E) time and space using Kahn's algorithm: courses can all
// be finished if and only if the prerequisite graph has no cycle, which
// Kahn's algorithm detects by checking whether every node can eventually
// reach in-degree zero.
func CanFinish(numCourses int, prerequisites [][2]int) bool {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		course, prereq := p[0], p[1]
		adj[prereq] = append(adj[prereq], course)
		indegree[course]++
	}

	queue := make([]int, 0, numCourses)
	for course := 0; course < numCourses; course++ {
		if indegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	processed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		processed++

		for _, next := range adj[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return processed == numCourses
}
