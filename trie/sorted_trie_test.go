package trie

import (
	"fmt"
	"sort"
	"testing"

	"github.com/bovinae/common/util"
)

func TestSortedTrieNode(t *testing.T) {
	var root *SortedTrieNode
	tmp := [][]rune{
		util.ReverseRuneSlice([]rune("1158万元人民币")),
		util.ReverseRuneSlice([]rune("3000万元人民币")),
		util.ReverseRuneSlice([]rune("100万元人民币")),
	}
	sort.Slice(tmp, func(i, j int) bool {
		return util.RunesCompare(tmp[i], tmp[j]) <= 0
	})
	root = root.Add(tmp[0])
	root = root.Add(tmp[1])
	root = root.Add(tmp[2])
	root.Dump()
	fmt.Println(string(util.ReverseRuneSlice(root.GetMaxPrefix(nil, nil, 2))))
}
