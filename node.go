package main

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	id       int
	hasToken bool
	Mutex    sync.Mutex
}

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

	//Create a WaitGroup
	var wg sync.WaitGroup

	//start up the nodes
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(node *Node) {
			defer wg.Done()
			node.enterCriticalSection(nodes)
		}(nodes[i%len(nodes)])
	}
	//wait for all goroutines to finish
	wg.Wait()

	//wait for a short duration to allow time for print statements
	time.Sleep(time.Second)
}

func (node *Node) enterCriticalSection(nodes []*Node) {
	for{ 
		//lock the node's mutex
		node.Mutex.Lock()

		if node.hasToken {

			fmt.Println("Node", node.id, "is now in the critical section")
			time.Sleep(time.Millisecond * 200)
			fmt.Println("Node", node.id, "has now left the critical section")

			nextNodeID := (node.id + 1) % len(nodes)

			//unlock the current node's Mutex before passing the token
			node.Mutex.Unlock()

			node.passToken(nextNodeID, nodes)
			return
		}	

		//if the node doesn't have the token, unlock the Mutex and sleep for a short time
		node.Mutex.Unlock()
		time.Sleep(time.Duration(200) * time.Millisecond)
	}

}

func (node *Node) passToken(nextID int, nodes []*Node) {
	nextNode := nodes[nextID]
	
	//lock the next node's Mutex
	nextNode.Mutex.Lock()

	//pass the token
	node.hasToken = false	
	nextNode.hasToken = true

	fmt.Println("Node", node.id, "has passed on the token to Node", nextNode.id)

	//Unlock the next node's Mutex
	nextNode.Mutex.Unlock()
}
	


