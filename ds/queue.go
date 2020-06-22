package ds

type Queue struct {
	queue *LinkedList
}

func NewQueue() *Queue {
	var q Queue
	q.queue = &LinkedList{}
	return &q
}

func (q *Queue) Pop() interface{} {
	ret := q.queue.GetHead()
	q.queue.DeletePosition(0)
	return ret.Value
}

func (q *Queue) Push(v interface{}) {
	var lp = &ListPoint{
		Value: v,
		Next:  nil,
	}
	q.queue.SetAtTail(lp)
}

func (q *Queue) Len() int {
	return q.queue.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
