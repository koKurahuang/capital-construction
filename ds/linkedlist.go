package ds

type ListPoint struct {
	Value interface{}
	Next  *ListPoint
}

func (lp *ListPoint) Equal(one *ListPoint) bool {
	return lp == one
}

// 单链表，不是队列
type LinkedList struct {
	Head *ListPoint
}

// 从0开始
func (q *LinkedList) DeletePosition(pos uint) {
	if int(pos) > q.Len()-1 {
		return
	}
	if pos == 0 {
		q.Head = q.Head.Next
		return
	}
	cnt := 1
	var previous = q.Head
	for {
		if cnt == int(pos) {
			previous.Next = previous.Next.Next
			return
		}
		previous = previous.Next
		cnt++
	}
}

func (q *LinkedList) Delete(lp *ListPoint) {
	if lp.Equal(q.Head) {
		q.Head = q.Head.Next
		return
	}
	var pre = q.Head
	var one = q.Head.Next
	for one != nil {
		if one.Equal(lp) {
			pre.Next = one.Next
			return
		}
		pre = one
		one = one.Next

	}
}

func (q *LinkedList) GetHead() *ListPoint {
	return q.Head
}

func (q *LinkedList) GetTail() *ListPoint {
	if q.Head == nil {
		return nil
	}
	var one = q.Head
	for one.Next != nil {
		one = one.Next
	}

	return one
}

// 从0开始
func (q *LinkedList) GetPosition(pos int) *ListPoint {
	if q.Head == nil {
		return nil
	}
	var cnt = 0
	one := q.Head
	for one.Next != nil {
		if cnt == pos {
			break
		} else {
			cnt++
			one = one.Next
		}
	}
	return one
}

func (q *LinkedList) Contains(lp *ListPoint) bool {
	if q.Head == nil {
		return false
	}
	if q.Head.Equal(lp) {
		return true
	}
	one := q.Head
	var ret bool
	for one.Next != nil {
		if one.Next.Equal(lp) {
			return true
		} else {
			one = one.Next
		}
	}
	return ret
}

func (q *LinkedList) IsEmpty() bool {
	return q.Head == nil
}

func (q *LinkedList) SetAtHead(lp *ListPoint) {
	lp.Next = q.Head
	q.Head = lp
}

func (q *LinkedList) SetAtTail(lp *ListPoint) {
	if q.IsEmpty() {
		q.SetAtHead(lp)
	} else {
		q.GetTail().Next = lp
	}
}

func (q *LinkedList) Len() int {

	var cnt = 0
	node := q.Head

	for node != nil {
		cnt ++
		node = node.Next
	}

	return cnt
}

// 从0开始
func (q *LinkedList) SetPosition(pos int, lp *ListPoint) {
	if q.Head == nil {
		q.Head = lp
		return
	}
	if pos == 0 {
		lp.Next = q.Head
		q.Head = lp
	}
	previous := q.Head
	var one = q.Head.Next
	cnt := 1
	for one != nil {
		if cnt == pos {
			lp.Next = one
			previous.Next = lp
			return
		}
		cnt ++
		previous = one
		one = one.Next
	}
}
