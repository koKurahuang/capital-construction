package ds

import (
	"fmt"
	"testing"
)

var root =&Node{}
var n1 =&Node{}
var n2 =&Node{}
var n3 =&Node{}
var n4 =&Node{}
var n5 =&Node{}

func init() {
	root.SetLeft(n1)
	root.SetRight(n2)
	n1.SetLeft(n3)
	n1.SetRight(n4)
	n2.SetLeft(n5)
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