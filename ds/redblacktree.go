package ds

import (
	"fmt"
)

var printAfterRotate = true

type RedBlackTree struct {
	root *RedBlackNode
}

func (rbt *RedBlackTree) Insert(value int) {
	if rbt.root == nil {
		root := &RedBlackNode{Black: true, value: value}
		rbt.root = root
		return
	}
	var parent *RedBlackNode
	one := rbt.root
	for one != nil {
		parent = one
		if value > one.value {
			one = one.right
		} else if value < one.value {
			one = one.left
		} else {
			panic(fmt.Sprintf("already exist value %d", value))
		}
	}
	var newNode RedBlackNode
	newNode.parent = parent
	newNode.value = value
	newNode.Black = false

	if value > parent.value {
		parent.right = &newNode
	} else {
		parent.left = &newNode
	}
	newNode.balance(rbt)

}

func (rbt *RedBlackTree) Print() [][]interface{}{
	if rbt.root == nil {
		return nil
	}
	ret := rbt.root.TreePrint()
	if printAfterRotate{
		for k, _ := range ret {
			for kk, _ := range ret[k] {
				fmt.Print(ret[k][kk])
			}
			fmt.Println()
		}
		fmt.Println("-----------------------------------------------------------------------------------")
	}

	return ret
}
