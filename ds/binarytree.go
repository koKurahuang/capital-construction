package ds

import (
	"math"
	"reflect"
)

const printBlank = "   "
const printEmpty = "   "

type Node struct {
	index  int
	left   *Node
	right  *Node
	parent *Node
	value  interface{}
}

func (n *Node) GetLeft() Tree {
	return n.left
}

func (n *Node) GetRight() Tree {
	return n.right
}

func (n *Node) GetRoot() Tree {
	if n.GetParent() == nil {
		return n
	} else {
		return n.GetParent().GetRoot()
	}
}

func (n *Node) IsRootNode() bool {
	return n.parent == nil
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func (n *Node) SetLeft(l Tree) {
	old_left := n.GetLeft()
	if !reflect.ValueOf(old_left).IsNil() {
		old_left.(*Node).setParent(nil)
	}
	n.left = l.(*Node)
	l.(*Node).setParent(n)
}

func (n *Node) SetRight(r Tree) {
	old_right := n.GetLeft()
	if !reflect.ValueOf(old_right).IsNil() {
		old_right.(*Node).setParent(nil)
	}
	n.right = r.(*Node)
	r.(*Node).setParent(n)
}

func (n *Node) setParent(p *Node) {
	n.parent = p
}

func (n *Node) Dfs() int {
	if n == nil {
		return 0
	}
	ldfs := n.left.Dfs()
	rdfs := n.right.Dfs()
	var max = 0
	if ldfs > rdfs {
		max = ldfs + 1
	} else {
		max = rdfs + 1
	}
	return max
}

func (n *Node) Bfs() []Tree {
	var que = NewQueue()
	que.Push(n)

	var retList []Tree
	for !que.IsEmpty() {
		get := que.Pop().(*Node)
		if get.left != nil {
			que.Push(get.left)
		}
		if get.right != nil {
			que.Push(get.right)
		}
		retList = append(retList, get)
	}

	return retList
}

func (n *Node) GetFloor() int {
	if n.IsRootNode() {
		return 1
	} else {
		return n.GetParent().GetFloor() + 1
	}
}

func (n *Node) TreePrint() [][]interface{} {
	list := n.getList()
	height := n.Dfs()

	var ret [][]interface{}

	for k, _ := range list {
		var one []interface{}
		lineNo := k + 1
		blankNum := int(math.Pow(2, float64(height-lineNo)) - 1)
		if blankNum < 0 {
			blankNum = 0
		}
		blank := ""
		for i := 0; i < blankNum; i++ {
			blank += printBlank
		}
		//fmt.Print(blank)
		one = append(one, blank)

		blankBetweenNum := int(math.Pow(2, float64(height-k)) - 1)
		blankBetween := ""
		for i := 0; i < blankBetweenNum; i++ {
			blankBetween += printBlank
		}
		for kk, _ := range list[k] {
			//fmt.Print(list[k][kk])
			one = append(one, list[k][kk])
			if kk < len(list[k])-1 {
				//fmt.Print(blankBetween)
				one = append(one, blankBetween)
			}
		}
		//fmt.Println(blank)
		one = append(one, blank)
		ret = append(ret, one)
	}

	return ret
}

func (n *Node) getList() [][]interface{} {
	var que = NewQueue()
	que.Push(n)
	height := n.Dfs()
	curHeight := 1
	var allFloor [][]interface{}
	for !que.IsEmpty() {
		if curHeight > height {
			break
		}
		var curFloor []interface{}
		var curFloorLen = que.Len()
		for i := 0; i < curFloorLen; i++ {
			one := que.Pop().(*Node)
			if one.value.(string) == printEmpty {
				curFloor = append(curFloor, printEmpty)
				que.Push(&Node{value: printEmpty})
				que.Push(&Node{value: printEmpty})
			} else {
				curFloor = append(curFloor, one.value)
				if one.left == nil {
					que.Push(&Node{value: printEmpty})
				} else {
					que.Push(one.left)
				}
				if one.right == nil {
					que.Push(&Node{value: printEmpty})
				} else {
					que.Push(one.right)
				}
			}
		}
		//fmt.Println(curFloor)
		allFloor = append(allFloor, curFloor)
		curHeight++
	}

	return allFloor
}

func (n*Node) GetLeaves() []Tree {
	var que = NewQueue()
	que.Push(n)

	var retList []Tree
	for !que.IsEmpty() {
		get := que.Pop().(*Node)
		if get.left != nil {
			que.Push(get.left)
		}
		if get.right != nil {
			que.Push(get.right)
		}
		if get.left == nil && get.right == nil {
			retList = append(retList, get)
		}
	}

	return retList
}