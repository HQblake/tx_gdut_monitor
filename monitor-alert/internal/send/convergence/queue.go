package convergence

import (
"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
)

// define a node
type Node struct {
	Target string
	Message model.Info
	Metric string
}

// queue
type Queue struct {
	Data Node
	Next *Queue
}

func InitQueue() *Queue {
	q := &Queue{}
	q.Next = nil
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.Next == nil
}

// push
func (q *Queue)EnQueue(data Node) {
	p := &Queue{}
	p.Data = data
	p.Next = nil

	if q.Next == nil {
		q.Next = p
	} else {
		s := q

		for s.Next != nil {

			s = s.Next
		}
		s.Next = p
	}

}

// pop
func (q *Queue) DeQueue() Node {
	if q.IsEmpty() {
		return Node{}
	} else {
		s := q.Next
		if s.Next == nil {
			q.Next = nil
		} else {
			q.Next = s.Next

		}
		return s.Data
	}
}

// size
func (q *Queue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	s := q
	var size int
	for s.Next != nil {
		size++
		s = s.Next
	}
	return size
}
// print
func Print(q *Queue) {
	if q.IsEmpty() {
		return
	}
	s := q
	for s.Next != nil {
		fmt.Println(s.Data)
		s = s.Next
	}
}

// front
func (q *Queue) Header() Node {
	if q.IsEmpty() {
		return Node{}
	}
	return q.Next.Data
}

// rear
func (q *Queue) Tail() Node {
	if q.IsEmpty() {
		return Node{}
	}
	s := q
	for s.Next != nil {
		s = s.Next
	}
	return s.Next.Data
}
