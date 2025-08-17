package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) dfsUtil(node int, visited map[int]bool) {

	visited[node] = true
	fmt.Print(node, " ")

	for _, neighbour := range g.adjList[node] {
		if !visited[neighbour] {
			g.dfsUtil(neighbour, visited)
		}
	}
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Print("DFS: ")
	g.dfsUtil(start, visited)
	fmt.Println()
}

func (g *Graph) BFS(start int) {

	visited := map[int]bool{}
	queue := []int{start}
	visited[start] = true

	fmt.Print("BFS : ")

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")

		for _, neighbour := range g.adjList[node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	fmt.Println()
}

func main() {

	g := Graph{adjList: make(map[int][]int)}

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)

	g.BFS(0)
	g.DFS(0)
}
