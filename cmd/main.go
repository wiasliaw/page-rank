package main

import "page-rank/pagerank"

func main() {
	pr := pagerank.New()
	pr.AddEdge("B", "A")
	pr.AddEdge("C", "A")
	pr.AddEdge("D", "A")
	pr.Calculate(0.85)
	pr.Display()
}
