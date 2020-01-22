package list

import (
	"math"
	"math/rand"
)

type SkipNode struct {
	value int
	left *SkipNode
	right *SkipNode
	up *SkipNode
	down *SkipNode
}

func NewSkipNode(value int) *SkipNode {
	return &SkipNode{
		value: value,
		left:  nil,
		right: nil,
		up:    nil,
		down:  nil,
	}
}
func NewSkipNodeFromLower(lowerNode *SkipNode) *SkipNode {
	return &SkipNode{
		value: lowerNode.value,
		left:  nil,
		right: nil,
		up:    nil,
		down:  lowerNode,
	}
}


type SkipList struct {
	positiveInfinity *SkipNode
	negativeInfinity *SkipNode
}

func NewSkipList() *SkipList {
	postInf := NewSkipNode(math.MaxInt32)
	negInf := NewSkipNode(math.MinInt32)
	postInf.left = negInf
	negInf.right = postInf
	return &SkipList{
		positiveInfinity: postInf,
		negativeInfinity: negInf,
	}
}

func (s *SkipList) Search(value int) *SkipNode {
	current := s.negativeInfinity
	for current != nil {
		for current.right != nil && current.right.value <= value {
			current = current.right
		}
		if current.value == value {
			break
		}
		current = current.down
	}
	return current
}

func (s *SkipList) Insert(value int) error {
	current := s.negativeInfinity
	pointersToUpdate := make([]*SkipNode, 0)
	for current != nil {
		for current.right != nil && current.right.value < value {
			current = current.right
		}
		pointersToUpdate = append(pointersToUpdate, current)
		current = current.down
	}
	level := 0
	var newNode *SkipNode = nil
	for level == 0 || FlipCoin() {
		if newNode == nil {
			newNode = NewSkipNode(value)
		} else {
			newNode = NewSkipNodeFromLower(newNode)
		}
		var nodeToUpdate *SkipNode = nil
		if len(pointersToUpdate) <= level {
			s.CreateNewLayer()
			nodeToUpdate =  s.negativeInfinity
		} else {
			nodeToUpdate = pointersToUpdate[len(pointersToUpdate) - level - 1]
		}
		newNode.right = nodeToUpdate.right
		newNode.left = nodeToUpdate

		newNode.right.left = newNode
		nodeToUpdate.right = newNode
		level ++
	}
	return  nil
}


func (s *SkipList) CreateNewLayer() {
	negativeInfinity := NewSkipNode(math.MinInt32)
	positiveInfinity := NewSkipNode(math.MaxInt32)
	negativeInfinity.right = positiveInfinity
	positiveInfinity.left = negativeInfinity
	s.negativeInfinity.up = negativeInfinity
	negativeInfinity.down = s.negativeInfinity
	s.negativeInfinity = negativeInfinity

	s.positiveInfinity.up = positiveInfinity
	positiveInfinity.down = s.positiveInfinity
	s.positiveInfinity = positiveInfinity
}

func FlipCoin() bool {
	return rand.Float32() >= 0.5
}
