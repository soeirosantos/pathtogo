package main

import (
	"bufio"
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
	"log"
	"os"
	"strings"
)

var GraphPath = os.Getenv("NODES_EDGES_LOCATION")

var nodeMap = make(map[string]*graph.Node, 0)

func main() {

	handleLine := make(chan string)
	done := make(chan bool)

	go readFile(handleLine, done)

	g := graph.New(graph.Directed)

createGraph:
	for {
		select {
		case line := <-handleLine:
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				continue
			}

			err := g.MakeEdge(*getFromNodeMap(parts[0], g), *getFromNodeMap(parts[1], g))
			if err != nil {
				log.Fatalf("error creating graph %v", err)
			}

		case <-done:
			break createGraph
		}
	}

	sorted := g.TopologicalSort()

	size := len(sorted) - 1
	for i := range sorted { // for now let's just redirect the output to a file
		fmt.Printf("%d,%s\n", i+1, *sorted[size-i].Value)
	}
}

func getFromNodeMap(key string, g *graph.Graph) *graph.Node {
	if _, ok := nodeMap[key]; !ok {
		n := g.MakeNode()
		*n.Value = key
		nodeMap[key] = &n
	}
	return nodeMap[key]
}

func readFile(lineChannel chan string, done chan bool) {

	file, err := os.Open(GraphPath)

	if err != nil {
		log.Fatalf("open file error %v", err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		lineChannel <- sc.Text()
	}

	close(lineChannel)
	close(done)

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}
}
