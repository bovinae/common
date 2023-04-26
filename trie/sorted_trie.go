package trie

import (
	"fmt"
)

type SortedTrieNode struct {
	Character rune
	Count     int
	Next      []*SortedTrieNode
}

func NewSortedTrieNode(c rune, cnt int) *SortedTrieNode {
	return &SortedTrieNode{
		Character: c,
		Count:     cnt,
	}
}

func (root *SortedTrieNode) Add(input []rune) *SortedTrieNode {
	if len(input) == 0 {
		return root
	}

	if root == nil || root.Character != input[0] {
		root = NewSortedTrieNode(input[0], 1)
		curr := root
		for i := 1; i < len(input); i++ {
			curr.Next = append(curr.Next, NewSortedTrieNode(input[i], 1))
			curr = curr.Next[0]
		}
		return root
	}

	var pre *SortedTrieNode
	curr := root
	for i := 0; i < len(input); i++ {
		if curr == nil {
			curr = NewSortedTrieNode(input[i], 1)
			pre.Next = append(pre.Next, curr)
			pre = curr
			curr = nil
			continue
		}
		if input[i] == curr.Character {
			curr.Count++
			pre = curr
			if len(curr.Next) == 0 {
				curr = nil
			} else {
				curr = curr.Next[len(curr.Next)-1]
			}
			continue
		}
		pre.Next = append(pre.Next, NewSortedTrieNode(input[i], 1))
		pre = curr
		curr = nil
	}

	return root
}

func (root *SortedTrieNode) GetMaxPrefix(prefix, maxPrefix []rune, threshold int) []rune {
	if root == nil {
		return maxPrefix
	}

	curr := root
	prefix = append(prefix, curr.Character)
	if curr.Count >= threshold && len(prefix) > len(maxPrefix) {
		maxPrefix = make([]rune, len(prefix))
		copy(maxPrefix, prefix)
	}

	for _, curr = range curr.Next {
		maxPrefix = curr.GetMaxPrefix(prefix, maxPrefix, threshold)
	}

	return maxPrefix
}

func (root *SortedTrieNode) Dump() {
	if root == nil {
		return
	}

	var queue []*SortedTrieNode
	queue = append(queue, root)
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			fmt.Printf("[%v,%v] ", queue[i].Character, queue[i].Count)
			queue = append(queue, queue[i].Next...)
		}
		fmt.Println("")
		queue = queue[qLen:]
	}
}
