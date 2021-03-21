package main

import "fmt"

func dfs(V int, adj [][]int, visited []int, result []int) []int {
	visited[V] = 1
	result = append(result, V)

	for _, x := range adj[V] {
		if visited[x] == 0 {
			res := dfs(x, adj, visited, result)
			result = res
		}
	}

	return result
}

func dfsOfGraph(V int, adj [][]int) []int {
	visited := make([]int, V)
	result := []int{}

	i := 0
	for i < V {
		if visited[i] == 0 {
			res := dfs(i, adj, visited, result)
			result = append(result, res...)
		}
		i++
	}

	return result
}

func main() {
	V := 5
	adj := [][]int{{1, 2, 3}, {0}, {0, 4}, {0}, {2}}

	ans := dfsOfGraph(V, adj)
	fmt.Println(ans)

	//0 1 2 4 3
}
