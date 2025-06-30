package server

type Node struct {
	prev *Node
	key  string
	next *Node
}

type LRUList struct {
	head       *Node
	tail       *Node
	count      int
	LRUnodeMap map[string]*Node
}

func CreateLRU() *LRUList {
	lru := &LRUList{head: nil, tail: nil, count: 0}

	return lru
}

func AddNode(lru *LRUList, dataKey string, secretKey string) bool {

	if lru.head == nil && lru.tail == nil {
		newNode := &Node{prev: nil, next: nil, key: dataKey}
		lru.head = newNode
		lru.tail = newNode
		lru.LRUnodeMap[dataKey] = newNode
		lru.count++
		return true
	} else {
		if lru.count >= 30 {
			evictedNode := lru.tail
			delete(globalStore[secretKey], evictedNode.key)
			lru.tail = evictedNode.prev
			lru.tail.next = nil
			lru.count--
			newNode := &Node{prev: nil, key: dataKey, next: nil}
			lru.head.prev = newNode
			newNode.next = lru.head
			lru.LRUnodeMap[dataKey] = newNode
			lru.head = newNode
			lru.count++
			return true
		} else {
			newNode := &Node{prev: nil, key: dataKey, next: nil}
			lru.head.prev = newNode
			newNode.next = lru.head
			lru.LRUnodeMap[dataKey] = newNode
			lru.head = newNode
			lru.count++
			return true
		}
	}

}

func RecentlyUsed(lru *LRUList, dataKey string) {

	usedNode := lru.LRUnodeMap[dataKey]

	if usedNode == nil || usedNode == lru.head {
		return
	}

	if lru.head == lru.tail {
		return
	}

	if usedNode == lru.tail {
		lru.tail = usedNode.prev
	}

	usedNode.prev.next = usedNode.next
	usedNode.next = lru.head
	lru.head.prev = usedNode
	usedNode.prev = nil
	lru.head = usedNode
}

func RemoveNode(lru *LRUList, dataKey string) {
	node := lru.LRUnodeMap[dataKey]
	if node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {

		lru.tail = node.prev
	}

	delete(lru.LRUnodeMap, dataKey)
	lru.count--
}
