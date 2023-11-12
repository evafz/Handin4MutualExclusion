package handin4mutualexclusion

import (
	"fmt"
	"sync"
)

type Node struct {
	id       int
	port     int
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

	//start up the nodes

	for i := 0; i < 3; i++ {
		go func(node *Node) {
			node.enterCriticalSection(nodes)
		}(nodes[i%len(nodes)])
	}
}

func (node *Node) enterCriticalSection(nodes []*Node) {
	node.Mutex.Lock()

	fmt.Println("A node is now in the critical section")
	//Sleep
	fmt.Println("The node has now left the critical section")

	nextNodeID := (node.id + 1) % len(nodes)
	node.passToken(nextNodeID, nodes)

}

func (node *Node) passToken(nextID int, nodes []*Node) {
	nextNode := nodes[nextID]
	nextNode.Mutex.Lock()

	node.hasToken = false
	node.Mutex.Unlock()

	nextNode.hasToken = true

	nextNode.Mutex.Unlock()

}

/*
func EnterRequest(Id int) {
	//wait
	//ask if others are in critsect
	//if one is true, wait and ask again over and over until they're not
	//if all are false, set yourself to true and go in babyh
	//call exit once done
}

func Exit(Id int) {
	//set inCriticalSection to false
	//start over with EnterRequest
}
*/
