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
var n6 =&Node{value:"006"}
var n7 =&Node{value:"007"}
var n8 =&Node{value:"008"}
var n9 =&Node{value:"009"}
var n10 =&Node{value:"010"}
var n11 =&Node{value:"011"}
var n12 =&Node{value:"012"}
var n13 =&Node{value:"013"}
var n14 =&Node{value:"014"}

func init() {
	root.SetLeft(n1)
	root.SetRight(n2)
	n1.SetLeft(n3)
	n1.SetRight(n4)
	n2.SetLeft(n5)
	n2.SetRight(n6)
	n3.SetLeft(n7)
	n3.SetRight(n8)
	n4.SetLeft(n9)
	n4.SetRight(n10)
	n5.SetLeft(n11)
	n5.SetRight(n12)
	n6.SetLeft(n13)
	n6.SetRight(n14)
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
	ret := n3.TreePrint()
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