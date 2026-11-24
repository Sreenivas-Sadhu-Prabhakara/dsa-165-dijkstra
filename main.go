package main

import (
	"container/heap"
	"fmt"
)

// Edge is an undirected weighted edge.
type Edge struct {
	U, V, W int
}

type pqItem struct{ dist, v int }

type pq []pqItem

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i].dist < p[j].dist }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)        { *p = append(*p, x.(pqItem)) }
func (p *pq) Pop() any {
	old := *p
	v := old[len(old)-1]
	*p = old[:len(old)-1]
	return v
}

// Dijkstra returns the shortest distance from src to every vertex of an
// undirected graph with non-negative weights; -1 marks unreachable. Lazy
// deletion: stale heap entries are skipped when popped.
func Dijkstra(n int, edges []Edge, src int) []int {
	type arc struct{ to, w int }
	adj := make([][]arc, n)
	for _, e := range edges {
		adj[e.U] = append(adj[e.U], arc{e.V, e.W})
		adj[e.V] = append(adj[e.V], arc{e.U, e.W})
	}
	const inf = int(^uint(0) >> 1)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	h := &pq{{0, src}}
	for h.Len() > 0 {
		it := heap.Pop(h).(pqItem)
		if it.dist > dist[it.v] {
			continue // stale entry; a shorter path was already settled
		}
		for _, a := range adj[it.v] {
			if nd := it.dist + a.w; nd < dist[a.to] {
				dist[a.to] = nd
				heap.Push(h, pqItem{nd, a.to})
			}
		}
	}
	for i, d := range dist {
		if d == inf {
			dist[i] = -1
		}
	}
	return dist
}

func main() {
	edges := []Edge{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}}
	fmt.Println(Dijkstra(4, edges, 0))
}
