package linked_list

import (
	"math/rand"
	"time"
)

//const (
//	maxLevel = 32
//	probability = 0.25
//)
//
//type Skiplist struct {
//
//}
//
//func Constructor() Skiplist {
//	return Skiplist{}
//}
//
//func (s *Skiplist) Search(target int) bool {
//	return false
//}
//
//func (s *Skiplist) Add(num int)  {
//
//}
//
//func (s *Skiplist) Erase(num int) bool {
//	return false
//}

const MaxLevel = 32

const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小

func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
	}
	return
}

type node struct {
	nextNodeArray []*node
	key           int
}

func newNode(key, level int) *node {
	return &node{key: key, nextNodeArray: make([]*node, level)}
}

type Skiplist struct {
	head  *node
	level int
}

func Constructor() Skiplist {
	return Skiplist{head: newNode(0, MaxLevel), level: 1}
}

func (s *Skiplist) Search(target int) bool {
	// 类似 delete
	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			if current.nextNodeArray[i].key == target {
				return true
			} else if current.nextNodeArray[i].key > target {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return false
}

func (s *Skiplist) Add(num int) {
	current := s.head
	update := make([]*node, MaxLevel) // 新节点插入以后的前驱节点
	for i := s.level - 1; i >= 0; i-- {
		if current.nextNodeArray[i] == nil || current.nextNodeArray[i].key > num {
			update[i] = current
		} else {
			for current.nextNodeArray[i] != nil && current.nextNodeArray[i].key < num {
				current = current.nextNodeArray[i] // 指针往前推进
			}
			update[i] = current
		}
	}

	level := randLevel()
	if level > s.level {
		// 新节点层数大于跳表当前层数时候, 现有层数 + 1 的 head 指向新节点
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}
	node := newNode(num, level)
	for i := 0; i < level; i++ {
		node.nextNodeArray[i] = update[i].nextNodeArray[i]
		update[i].nextNodeArray[i] = node
	}
}

func (s *Skiplist) Erase(num int) bool {
	current := s.head
	flag := false
	for i := s.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			if current.nextNodeArray[i].key == num {
				tmp := current.nextNodeArray[i]
				current.nextNodeArray[i] = tmp.nextNodeArray[i]
				tmp.nextNodeArray[i] = nil
				flag = true
				// 这里要 break, 因为leetcode 的判定是，在有两个 num 9时， erase 只能删除一个
				break
			} else if current.nextNodeArray[i].key > num {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return flag
}
