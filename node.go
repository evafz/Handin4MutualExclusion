package main

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	id       int
	port     int
	hasToken bool
	Mutex    sync.Mutex
}

var globalMutex sync.Mutex
var wg sync.WaitGroup

func main() {
	//make three nodes in an array
	nodeIDs := []int{1, 2, 3}
	nodes := make([]*Node, len(nodeIDs))
	//set the id of the nodes to the id
	for i, id := range nodeIDs {
		nodes[i] = &Node{
			id: id,
		}
	}

	//start with one node having the token:
	startWithToken := nodes[0]
	startWithToken.hasToken = true

	//start up the nodes

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(node *Node) {
			defer wg.Done()
			node.enterCriticalSection(nodes)
		}(nodes[i%len(nodes)])
	}

	wg.Wait()

	select {}
}

func (node *Node) enterCriticalSection(nodes []*Node) {
	globalMutex.Lock()

	if node.hasToken {

		fmt.Println("Node", node.id, "is now in the critical section")
		time.Sleep(time.Millisecond * 200)
		fmt.Println("Node", node.id, "has now left the critical section")

		globalMutex.Unlock()

		nextNodeID := (node.id + 1) % len(nodes)
		node.passToken(nextNodeID, nodes)

	} else {
		globalMutex.Unlock()
		time.Sleep(time.Duration(200) * time.Millisecond)
	}

}

func (node *Node) passToken(nextID int, nodes []*Node) {
	nextNode := nodes[nextID]
	globalMutex.Lock()

	node.Mutex.Lock()
	node.hasToken = false
	node.Mutex.Unlock()

	nextNode.Mutex.Lock()
	nextNode.hasToken = true
	nextNode.Mutex.Unlock()

	fmt.Println("Node", node.id, "has passed on the token to Node", nextNode.id)

	globalMutex.Unlock()

}
