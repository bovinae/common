package trie

import (
	"sort"
	"testing"

	"github.com/bovinae/common/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSortedTrieNode(t *testing.T) {
	var root *SortedTrieNode[rune]
	tmp := [][]rune{
		util.ReverseSlice([]rune("1158万元人民币")),
		util.ReverseSlice([]rune("3000万元人民币")),
		util.ReverseSlice([]rune("100万元人民币")),
	}
	sort.Slice(tmp, func(i, j int) bool {
		return util.RunesCompare(tmp[i], tmp[j]) <= 0
	})
	root = root.Add(tmp[0])
	root = root.Add(tmp[1])
	root = root.Add(tmp[2])
	root.Dump()
	Convey("TestSortedTrieNode", t, func() {
		Convey("TestSortedTrieNode", func() {
			suffix := string(util.ReverseSlice(root.GetMaxPrefix(nil, nil, 2)))
			So(suffix, ShouldEqual, "00万元人民币")
		})
	})
}
