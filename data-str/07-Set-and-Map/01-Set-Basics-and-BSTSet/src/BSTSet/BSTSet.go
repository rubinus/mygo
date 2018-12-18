package BSTSet

import (
	"mygo/data-str/07-Set-and-Map/01-Set-Basics-and-BSTSet/src/BST"
	"mygo/data-str/07-Set-and-Map/01-Set-Basics-and-BSTSet/src/Set"
)

type BSTSet struct {
	BST *BST.BST
}

func GetBSTSet() *BSTSet {
	bst := BST.GetBST()
	return &BSTSet{bst}
}

func (s *BSTSet) Add(e Set.E) {
	s.BST.Add(e)
}

func (s *BSTSet) Remove(e Set.E) {
	s.BST.Remove(e)
}

func (s *BSTSet) Contains(e Set.E) bool {
	return s.BST.Contains(e)
}

func (s *BSTSet) GetSize() int {
	return s.BST.GetSize()
}

func (s *BSTSet) IsEmpty() bool {
	return s.BST.IsEmpty()
}
