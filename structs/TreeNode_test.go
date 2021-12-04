package structs

import (
	"reflect"
	"testing"
)

var (
	// 同一个 TreeNode 的不同表达方式
	//            1
	//      	/  \
	//         2    3
	//        / \  /  \
	//       4  5  6   7
	LeetCodeOrder = []int{1, 2, 3, 4, 5, 6, 7}
	preOrder      = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder       = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder     = []int{4, 5, 2, 6, 7, 3, 1}
)

func Test_Ints2TreeNode(t *testing.T) {
	expected := PreIn2Tree(preOrder, inOrder)
	actual := Ints2TreeNode(LeetCodeOrder)
	assertEqual(t, expected, actual)

	actual = Ints2TreeNode([]int{})
	assertNil(t, actual)
}

func Test_preIn2Tree(t *testing.T) {
	actual := Tree2Postorder(PreIn2Tree(preOrder, inOrder))
	expected := postOrder
	assertEqual(t, expected, actual)

	assertPanics(t, func() { PreIn2Tree([]int{1}, []int{}) })

	assertNil(t, PreIn2Tree([]int{}, []int{}))
}

func Test_inPost2Tree(t *testing.T) {
	actual := Tree2Preorder(InPost2Tree(inOrder, postOrder))
	expected := preOrder
	assertEqual(t, expected, actual)

	assertPanics(t, func() { InPost2Tree([]int{1}, []int{}) })

	assertNil(t, InPost2Tree([]int{}, []int{}))
}

func Test_tree2Ints(t *testing.T) {
	root := PreIn2Tree(preOrder, inOrder)

	assertEqual(t, preOrder, Tree2Preorder(root))
	assertNil(t, Tree2Preorder(nil))

	assertEqual(t, inOrder, Tree2Inorder(root))
	assertNil(t, Tree2Inorder(nil))

	assertEqual(t, postOrder, Tree2Postorder(root))
	assertNil(t, Tree2Postorder(nil))
}

func Test_indexOf(t *testing.T) {
	assertEqual(t, 1, indexOf(1, []int{0, 1, 2, 3}))

	assertPanics(t, func() { indexOf(0, []int{1, 2, 3}) })
}

func Test_TreeNode_Equal(t *testing.T) {
	type args struct {
		a *TreeNode
	}
	tests := []struct {
		name   string
		fields args
		args   args
		want   bool
	}{

		{
			"相等",
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			true,
		},

		{
			"不相等",
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			args{Ints2TreeNode([]int{1, 2, 3, NULL, 5})},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := tt.fields.a
			if got := tn.Equal(tt.args.a); got != tt.want {
				t.Errorf("TreeNode.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTargetNode(t *testing.T) {
	ints := []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}
	root := Ints2TreeNode(ints)

	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{

		{
			"找到 root.Right.Right",
			args{
				root:   root,
				target: 8,
			},
			root.Right.Right,
		},

		{
			"找到 root.Left.Left",
			args{
				root:   root,
				target: 6,
			},
			root.Left.Left,
		},

		//
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTargetNode(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTargetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tree2ints(t *testing.T) {
	root := PreIn2Tree(preOrder, inOrder)
	actual := LeetCodeOrder
	expected := Tree2ints(root)
	assertEqual(t, expected, actual)
}
