package copyRandomList

/**
 * Definition for a Node.
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	newHead := &Node{}
	q := newHead

	for p := head; p != nil; p = p.Next {
		newNode := &Node{p.Val, nil, nil}
		q.Next = newNode
		q = q.Next
		nodeMap[p] = newNode
	}

	for p, q := head, newHead.Next; q != nil; p, q = p.Next, q.Next {
		node, ok := nodeMap[p.Random]
		if !ok {
			q.Random = nil
		} else {
			q.Random = node
		}
	}
	return newHead.Next
}
