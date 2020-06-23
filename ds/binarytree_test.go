package ds

import (
	"fmt"
	"testing"
)

var root =&Node{value:"000"}
var n1 =&Node{value:"001"}
var n2 =&Node{value:"002"}
var n3 =&Node{value:"003"}
var n4 =&Node{value:"004"}
var n5 =&Node{value:"005"}

func init() {
	root.SetLeft(n1)
	root.SetRight(n2)
	n2.SetRight(n3)
	n3.SetRight(n4)
	n4.SetRight(n5)
}

func TestNode_Dfs(t *testing.T) {
	fmt.Print(root.Dfs())
}

func TestNode_GetFloor(t *testing.T) {
	fmt.Println(root.GetFloor())
	fmt.Println(n2.GetFloor())
	fmt.Println(n5.GetFloor())
}

func TestNode_GetRoot(t *testing.T) {
	fmt.Println(n4.GetRoot().IsRootNode())
}

func TestGetList(t *testing.T) {
	fmt.Println(root.getList())
}


func TestAppendNil(t *testing.T) {
	var a []*Node
	a = append(a, nil)
	a = append(a, &Node{value:1})
	fmt.Println(len(a))
	fmt.Println(a)
}

func TestNode_TreePrint(t *testing.T) {
	ret := root.TreePrint()
	for k, _ := range ret{
		for kk, _ := range ret[k] {
			fmt.Print(ret[k][kk])
		}
		fmt.Println()
	}
}

func TestNode_GetLeaves(t *testing.T) {
	fmt.Println(root.GetLeaves())
}