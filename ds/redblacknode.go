package ds

import (
	color "github.com/koKurahuang/capital-construction/screencolor"
	"math"
	"reflect"
	"strconv"
)

//https://www.cnblogs.com/ailumiyana/p/10963658.html

type RedBlackNode struct {
	Black  bool
	index  int
	left   *RedBlackNode
	right  *RedBlackNode
	parent *RedBlackNode
	value  int
}

func (n *RedBlackNode) GetLeft() Tree {
	return n.left
}

func (n *RedBlackNode) GetRight() Tree {
	return n.right
}

func (n *RedBlackNode) GetRoot() Tree {
	if n.GetParent() == nil {
		return n
	} else {
		return n.GetParent().GetRoot()
	}
}

func (n *RedBlackNode) IsRootNode() bool {
	return n.parent == nil
}

func (n *RedBlackNode) GetParent() *RedBlackNode {
	return n.parent
}

func (n *RedBlackNode) GetValue() interface{} {
	return n.value
}

func (n *RedBlackNode) SetLeft(l Tree) {
	old_left := n.GetLeft()
	if !reflect.ValueOf(old_left).IsNil() {
		old_left.(*RedBlackNode).setParent(nil)
	}
	n.left = l.(*RedBlackNode)
	l.(*RedBlackNode).setParent(n)
}

func (n *RedBlackNode) SetRight(r Tree) {
	old_right := n.GetLeft()
	if !reflect.ValueOf(old_right).IsNil() {
		old_right.(*RedBlackNode).setParent(nil)
	}
	n.right = r.(*RedBlackNode)
	r.(*RedBlackNode).setParent(n)
}

func (n *RedBlackNode) setParent(p *RedBlackNode) {
	n.parent = p
}

func (n *RedBlackNode) Dfs() int {
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

func (n *RedBlackNode) Bfs() []Tree {
	var que = NewQueue()
	que.Push(n)

	var retList []Tree
	for !que.IsEmpty() {
		get := que.Pop().(*RedBlackNode)
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

func (n *RedBlackNode) GetFloor() int {
	if n.IsRootNode() {
		return 1
	} else {
		return n.GetParent().GetFloor() + 1
	}
}

func (n *RedBlackNode) TreePrint() [][]interface{} {
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

func (n *RedBlackNode) getList() [][]interface{} {
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
			one := que.Pop().(*RedBlackNode)
			if one.value == math.MaxInt64 {
				curFloor = append(curFloor, printEmpty)
				que.Push(&RedBlackNode{value: math.MaxInt64})
				que.Push(&RedBlackNode{value: math.MaxInt64})
			} else {
				if one.Black {
					curFloor = append(curFloor, color.ColorPrint(itoA(one.value), color.GRAY))
				} else {
					curFloor = append(curFloor, color.ColorPrint(itoA(one.value), color.RED))
				}
				if one.left == nil {
					que.Push(&RedBlackNode{value: math.MaxInt64})
				} else {
					que.Push(one.left)
				}
				if one.right == nil {
					que.Push(&RedBlackNode{value: math.MaxInt64})
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

func (n *RedBlackNode) GetLeaves() []Tree {
	var que = NewQueue()
	que.Push(n)

	var retList []Tree
	for !que.IsEmpty() {
		get := que.Pop().(*RedBlackNode)
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

//因为红黑树有默认的空黑叶，所以不做检测子节点非空检测了
func (rbn *RedBlackNode) leftRotate(rbt *RedBlackTree) {
	//fmt.Println("----------------------------左----------------旋------------------------------------")
	if rbn == nil {
		return
	}

	parent := rbn.GetParent()
	if parent.parent == nil {
		rbt.root = rbn
		rbn.parent = nil
	}else{
		if parent == parent.parent.left {
			parent.parent.left = rbn
			rbn.parent = rbn.parent.parent
		}else {
			parent.parent.right = rbn
			rbn.parent = rbn.parent.parent
		}
	}

	old_left := rbn.left
	parent.right = old_left
	rbn.left = parent
	//rbt.Print()
}

//因为红黑树有默认的空黑叶，所以不做检测子节点非空检测了
func (rbn *RedBlackNode) rightRotate(rbt *RedBlackTree) {
	//fmt.Println("----------------------------右----------------旋------------------------------------")
	if rbn == nil {
		return
	}

	parent := rbn.GetParent()
	if parent.parent == nil {
		rbt.root = rbn
		rbn.parent = nil
	}else{
		if parent == parent.parent.left {
			parent.parent.left = rbn
			rbn.parent = rbn.parent.parent
		}else {
			parent.parent.right = rbn
			rbn.parent = rbn.parent.parent
		}
	}
	old_right := rbn.right
	parent.left = old_right
	rbn.right = parent
	//rbt.Print()
}

func (rbn *RedBlackNode) getUncle() *RedBlackNode {
	if rbn == nil {
		return nil
	}
	if rbn.IsRootNode() {
		return nil
	}
	if rbn.parent.IsRootNode() {
		return nil
	}
	if rbn == rbn.parent.left {
		return rbn.parent.right
	} else {
		return rbn.parent.left
	}
}

// 为了避免叔父节点为空时读取颜色会空指针
func (rbn *RedBlackNode) getColor() bool {
	if rbn == nil {
		return true
	}
	return rbn.Black
}

func (rbn *RedBlackNode) balance(rbt *RedBlackTree) {

	spawnPoint := rbn
	for !spawnPoint.parent.getColor() {
		// 情况1,空树，直接变个色
		if spawnPoint.IsRootNode() {
			spawnPoint.Black = true
			return
		}


		// 情况2不用做变化

		// 情况345， 从插入的节点开始往上遍历检测（不一定是一层层往上的，叔父都是红的话平衡完直接跳到爷检测的，因为少了一层，父下去了）
		// 主要判断依据就是父节点，父父和叔节点的颜色
		// 只要父是黑色的就没事了

		// 父红，父父不可能红，在上一次就被平衡掉了
		if spawnPoint.parent == spawnPoint.parent.parent.left {
			uncle := spawnPoint.parent.parent.right
			// 情况4，叔父都是红的，一起变色，爷也变色，然后跳到爷开始检测
			if !uncle.getColor() {
				spawnPoint.parent.Black = true
				uncle.Black = true
				spawnPoint.parent.parent.Black = false

				spawnPoint = spawnPoint.parent.parent
				//continue
			} else {
				// 左左， 情况3， 父爷变色， 父右旋
				if spawnPoint == spawnPoint.parent.left {
					spawnPoint.parent.Black = true
					spawnPoint.parent.parent.Black = false

					spawnPoint.parent.rightRotate(rbt)

					spawnPoint = spawnPoint.parent
					// 左右，  情况5， 本身左旋， 变情况3
				} else {
					spawnPoint.leftRotate(rbt)
					spawnPoint.Black = true
					spawnPoint.parent.Black = false

					spawnPoint.rightRotate(rbt)

					spawnPoint = spawnPoint.parent
				}
			}
		} else {
			uncle := spawnPoint.parent.parent.left
			// 情况4，叔父都是红的，一起变色，爷也变色，然后跳到爷开始检测
			if !uncle.getColor() {
				spawnPoint.parent.Black = true
				uncle.Black = true
				spawnPoint.parent.parent.Black = false

				spawnPoint = spawnPoint.parent.parent
				//continue
			} else {
				// 右左， 情况5， 右旋,变情况3
				if spawnPoint == spawnPoint.parent.left {
					spawnPoint.rightRotate(rbt)
					spawnPoint.Black = true
					spawnPoint.parent.Black = false

					spawnPoint.leftRotate(rbt)
					spawnPoint = spawnPoint.parent
					// 右右，  情况3， 变色, 父爷变色， 父左旋
				}else{

					spawnPoint.parent.Black = true
					spawnPoint.parent.parent.Black = false

					spawnPoint.parent.leftRotate(rbt)
					spawnPoint = spawnPoint.parent
				}

			}

		}
	}
	//有可能会产生spawpoint被刷到根节点，这样就不循环了，要是是红的话就错了
	if spawnPoint.IsRootNode() && !spawnPoint.getColor() {
		spawnPoint.Black = true
	}
}

func itoA(i int) string {
	s := strconv.Itoa(i)
	if i > 99 {
		return s
	} else if i > 9 && i < 100 {
		return "0" + s
	} else {
		return "00" + s
	}
}
