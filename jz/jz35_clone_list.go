package main

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func Clone(head *RandomListNode) *RandomListNode {
	// Init the link list.
	if head == nil {
		return nil
	}

	hash := make(map[*RandomListNode]*RandomListNode)

	newHead := new(RandomListNode)
	newCur := newHead
	cur := head
	for cur != nil {
		temp := new(RandomListNode)
		temp.Label = cur.Label
		// Set the relation.
		hash[cur] = temp
		newCur.Next = temp
		newCur = temp
		cur = cur.Next
	}

	for k, v := range hash {
		v.Random = selectNode(k, hash)
	}

	return newHead.Next
}

func selectNode(node *RandomListNode, hash map[*RandomListNode]*RandomListNode) *RandomListNode {
	if node.Random == nil {
		return nil
	} else {
		return hash[node.Random]
	}
}

// func Clone( head *RandomListNode ) *RandomListNode {
//     if head == nil {
//         return head
//     }

//     p1, p2 := head, head
//     m := map[*RandomListNode]*RandomListNode{}

//     // p1指针扫描旧链表，建立结点map
//     for p1 != nil {
//         m[p1] = &RandomListNode{ Label: p1.Label }
//         p1 = p1.Next
//     }

//     // p2指针扫描旧链表，对map进行结点链接
//     for p2 != nil {
//         m[p2].Next = m[p2.Next]
//         m[p2].Random = m[p2.Random]
//         p2 = p2.Next
//     }

//     // head指针一直指向链表头
//     return m[head]
// }
