package main

// LRU缓存
type LRUCache struct {
	CntMap   map[int]*Node
	size     int
	capacity int
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{CntMap: map[int]*Node{}, size: 0, capacity: capacity, head: InitNode(0, 0), tail: InitNode(0, 0)}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.CntMap[key]; !ok {
		return -1
	}
	node := this.CntMap[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) removeTailNode() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.CntMap[key]; ok {
		node := this.CntMap[key]
		node.value = value
		this.moveToHead(node)
	} else {
		node := &Node{key: key, value: value}
		if this.size < this.capacity {
			this.addToHead(node)
			this.CntMap[key] = node
			this.size++
		} else {
			tailNode := this.removeTailNode()
			delete(this.CntMap, tailNode.key)
			this.addToHead(node)
			this.CntMap[key] = node
		}
	}

}

func InitNode(key int, value int) *Node {
	return &Node{key: key, value: value}
}
