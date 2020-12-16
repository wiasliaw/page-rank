package pagerank

import "fmt"

type node struct {
	Prob float32
	PR   float32
	Link int
}

type pagerank struct {
	node map[interface{}]*node
	edge map[interface{}][]interface{}
	size int
}

func New() *pagerank {
	pr := new(pagerank)
	pr.Clear()
	return pr
}

func (pr *pagerank) AddEdge(from interface{}, to interface{}) {
	// insert `from` and `to` if not existed
	_, exist := pr.node[from]
	if !exist {
		pr.node[from] = &node{}
	}
	_, exist = pr.node[to]
	if !exist {
		pr.node[from] = &node{}
	}
	pr.node[to] = &node{}
	// update `pr`'s value
	pr.size = len(pr.node)
	pr.edge[from] = append(pr.edge[from], to)
	pr.node[from].Link++
}

func (pr pagerank) Display() {
	fmt.Printf("graph size: %d\n", pr.size)
	fmt.Println("[Node]")
	for k, n := range pr.node {
		fmt.Printf("<%s> Prob:%f, PR:%f\n", k, n.Prob, n.PR)
	}
	fmt.Println("[Edge]")
	for k, a := range pr.edge {
		for _, e := range a {
			fmt.Printf("%s -> %s\n", k, e)
		}
	}
}

func (pr *pagerank) Clear() {
	pr.node = make(map[interface{}]*node)
	pr.edge = map[interface{}][]interface{}{}
}

func (pr *pagerank) Calculate(d float32) {
	// init
	for _, n := range pr.node {
		n.Prob = float32(1) / float32(pr.size)
		n.PR = (float32(1) - d) / float32(pr.size)
	}
	for toKey, toNode := range pr.node {
		tmp := float32(0)
		for fromKey, fromNode := range pr.node {
			if toKey == fromKey {
				continue
			}
			// check edge existed or not
			if isExist(pr.edge[fromKey], toKey) {
				tmp += fromNode.Prob / float32(fromNode.Link)
			}
		}
		toNode.PR += tmp * d
	}
}
