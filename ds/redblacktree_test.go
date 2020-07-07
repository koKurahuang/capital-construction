package ds

import (
	"fmt"
	"testing"
)

var rb_root = &RedBlackNode{Black:false, value: 0}
var rb_n1 = &RedBlackNode{Black:false, value: 1}
var rb_n2 = &RedBlackNode{Black:true, value:2}
var rb_n3 = &RedBlackNode{Black:false, value:3}
var rb_n4 = &RedBlackNode{Black:true, value: 4}
var rb_n5 = &RedBlackNode{Black:false, value: 5}
var rb_n6 = &RedBlackNode{Black:false, value:6}
var rb_n7 = &RedBlackNode{Black:false, value: 7}
var rb_n8 = &RedBlackNode{Black:true, value: 8}
var rb_n9 = &RedBlackNode{Black:false, value: 9}
var rb_n10 = &RedBlackNode{Black:true, value: 10}
var rb_n11 = &RedBlackNode{Black:true, value: 11}
var rb_n12 = &RedBlackNode{Black:false, value: 12}
var rb_n13 = &RedBlackNode{Black:false, value: 13}
var rb_n14 = &RedBlackNode{Black:true, value: 14}

func init() {
	rb_root.SetLeft(rb_n1)
	rb_root.SetRight(rb_n2)
	//rb_n1.SetLeft(rb_n3)
	//rb_n1.SetRight(rb_n4)
	rb_n2.SetLeft(rb_n5)
	rb_n2.SetRight(rb_n6)
	/*rb_n3.SetLeft(rb_n7)
	rb_n3.SetRight(rb_n8)
	rb_n4.SetLeft(rb_n9)
	rb_n4.SetRight(rb_n10)
	rb_n5.SetLeft(rb_n11)
	rb_n5.SetRight(rb_n12)
	rb_n6.SetLeft(rb_n13)
	rb_n6.SetRight(rb_n14)*/
}

func TestRedBlackNode_TreePrint(t *testing.T) {
	ret := rb_root.TreePrint()
	for k, _ := range ret {
		for kk, _ := range ret[k] {
			fmt.Print(ret[k][kk])
		}
		fmt.Println()
	}
}

func TestRedBlackNode_leftRotate(t *testing.T) {
	rb_root.TreePrint()

	rb_n2.leftRotate(nil)

}

func TestRedBlackTree_Insert(t *testing.T) {
	var tree RedBlackTree
	tree.Insert(100)
	tree.Insert(90)
	tree.Insert(91)
	tree.Insert(92)
	tree.Insert(93)
	tree.Insert(94)

/*	for i := 90;i<=110;i ++ {
		if i == 100 {
			continue
		}
		tree.Insert(i)
	}*/
	ret := tree.root.TreePrint()
	for k, _ := range ret {
		for kk, _ := range ret[k] {
			fmt.Print(ret[k][kk])
		}
		fmt.Println()
	}
}