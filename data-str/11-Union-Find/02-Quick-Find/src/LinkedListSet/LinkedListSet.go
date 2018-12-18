package LinkedListSet

import (
	"mygo/data-str/11-Union-Find/02-Quick-Find/src/LinkedList"
	"mygo/data-str/11-Union-Find/02-Quick-Find/src/Set"
)

type LinkedListSet struct {
	LinkedList *LinkedList.LinkedList
}

func Constructor() *LinkedListSet {
	list := LinkedList.GetLinkedList()
	return &LinkedListSet{list}
}

func (s *LinkedListSet) Add(e Set.E) {
	if !s.LinkedList.Contains(e) {
		s.LinkedList.AddFirst(e)
	}
}

func (s *LinkedListSet) Remove(e Set.E) {
	s.LinkedList.RemoveElement(e)
}

func (s *LinkedListSet) Contains(e Set.E) bool {
	return s.LinkedList.Contains(e)
}

func (s *LinkedListSet) GetSize() int {
	return s.LinkedList.GetSize()
}

func (s *LinkedListSet) IsEmpty() bool {
	return s.LinkedList.IsEmpty()
}
