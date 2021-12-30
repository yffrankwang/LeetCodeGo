package leetcode

import (
	"io"

	"github.com/yffrankwang/LeetCodeGo/structs"
)

// ListNode define
type ListNode = structs.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists2(lists)
}

//-----------------------------------------------------
func mergeKLists1(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists1(left, right)
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}

//-------------------------------------------------------------
func mergeKLists2(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	nk := newListNodeKmerge(lists)
	km := NewKmerge(length, nk)

	km.Merge()

	return nk.root.Next
}

type listNodeKmerge struct {
	lists []*ListNode
	heads []*ListNode

	root *ListNode
	last *ListNode
}

func newListNodeKmerge(lists []*ListNode) *listNodeKmerge {
	nk := &listNodeKmerge{lists: lists}
	nk.heads = make([]*ListNode, len(lists))
	nk.root = &ListNode{}
	nk.last = nk.root
	return nk
}

func (nk *listNodeKmerge) Less(a, b int) bool {
	return nk.heads[a].Val < nk.heads[b].Val
}

func (nk *listNodeKmerge) Read(i int) error {
	ln := nk.lists[i]
	if ln == nil {
		return io.EOF
	}
	nk.heads[i] = ln
	nk.lists[i] = ln.Next
	return nil
}

func (nk *listNodeKmerge) Write(i int) error {
	nk.last.Next = nk.heads[i]
	nk.last = nk.last.Next
	return nil
}

// Less less function for loser tree
// Should return a bool:
//    true , if a < b
//    false, if a >= b
type Less func(a, b int) bool

// LoserTree array based loser tree
type LoserTree struct {
	nodes []int
	less  Less
}

// NewLoserTree create a loser tree instance
func NewLoserTree(size int, less Less) *LoserTree {
	return &LoserTree{nodes: make([]int, size), less: less}
}

// Construct construct the loser tree
func (lt *LoserTree) Construct() {
	for i := len(lt.nodes) - 1; i >= 0; i-- {
		lt.nodes[i] = -1
	}
	for i := len(lt.nodes) - 1; i >= 0; i-- {
		lt.adjust(i)
	}
}

// Root get the root node (winner)
func (lt *LoserTree) Root() int {
	return lt.nodes[0]
}

func (lt *LoserTree) adjust(q int) {
	t := (q + len(lt.nodes)) >> 1
	p := lt.nodes[t]

	// nodes[t] = -1 --> minimum
	for t > 0 {
		if p == -1 || (q != -1 && lt.less(p, q)) {
			q, lt.nodes[t] = lt.nodes[t], q
		}
		t = t >> 1
		p = lt.nodes[t]
	}
	lt.nodes[0] = q
}

// Adjust adjust the loser tree from the root
func (lt *LoserTree) Adjust() {
	q := lt.nodes[0]
	t := (q + len(lt.nodes)) >> 1
	p := lt.nodes[t]

	for t > 0 {
		if lt.less(p, q) {
			q, lt.nodes[t] = lt.nodes[t], q
		}
		t = t >> 1
		p = lt.nodes[t]
	}
	lt.nodes[0] = q
}

// KmergeIO K-Merge IO interface
type KmergeIO interface {
	Less(a, b int) bool
	Read(i int) error
	Write(i int) error
}

// Kmerge a K-Merge struct
type Kmerge struct {
	lotr *LoserTree // loser tree
	eofs []bool     // End-Of-File slice
	kmio KmergeIO   // IO interface
}

// NewKmerge create a K-Merge instance
func NewKmerge(k int, kmio KmergeIO) *Kmerge {
	km := &Kmerge{eofs: make([]bool, k), kmio: kmio}

	km.lotr = NewLoserTree(k, func(a, b int) bool {
		if km.eofs[a] {
			return false
		}
		if km.eofs[b] {
			return true
		}
		return km.kmio.Less(a, b)
	})
	return km
}

func (km *Kmerge) read(i int) error {
	err := km.kmio.Read(i)
	if err == io.EOF {
		km.eofs[i] = true
	} else if err != nil {
		return err
	}
	return nil
}

// Merge do k-merge
func (km *Kmerge) Merge() error {
	// init read
	for i := 0; i < len(km.eofs); i++ {
		if err := km.read(i); err != nil {
			return err
		}
	}

	// construct the loser tree
	km.lotr.Construct()

	// do k-merge
	for !km.eofs[km.lotr.Root()] {
		if err := km.kmio.Write(km.lotr.Root()); err != nil {
			return err
		}

		if err := km.read(km.lotr.Root()); err != nil {
			return err
		}

		km.lotr.Adjust()
	}

	return nil
}
