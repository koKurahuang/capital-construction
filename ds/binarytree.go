package ds

import (
	"strconv"
)

type Node struct {
	id     string
	left   *Node
	right  *Node
	parent *Node
	value  interface{}
}

func NewNode(rootNode bool) *Node {
	var n Node
	if rootNode {
		n.id = "000"
	}
	return &n
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) GetRoot() *Node {
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

func (n *Node) SetLeft(l *Node) {
	old_left := n.GetLeft()
	if old_left != nil {
		old_left.setParent(nil)
	}
	n.left = l
	l.setParent(n)
	num, _ := strconv.Atoi(n.id)
	l.id = transferInt(2 * num)
}

func (n *Node) SetRight(r *Node) {
	old_right := n.GetLeft()
	if old_right != nil {
		old_right.setParent(nil)
	}
	n.right = r
	r.setParent(n)
	num, _ := strconv.Atoi(n.id)
	r.id = transferInt(2*num + 1)
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

func (n *Node) Bfs() []*Node{
	var que = NewQueue()
	que.Push(n)

	var retList []*Node
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

func transferInt(n int) string {
	num := strconv.Itoa(int(n))
	if n < 10 {
		return "00" + num
	} else if 10 <= n || n < 100 {
		return "0" + num
	} else {
		return num
	}
}

func (n *Node) TreePrint() {
/*	var blank = "   "
	height := n.Dfs()
	maxLength := math.Pow(2, float64(height)) - 1
	for i:=1; i<= height; i++ {
		var sideBlank string
		if height - 1 != 0 {
			blankNum := 2*(height - i) + 1

			for j:=1;j<= blankNum; j ++ {
				sideBlank += blank
			}
		}
		fmt.Print(sideBlank)
		for j:=1;j<=

	}
*/
}
