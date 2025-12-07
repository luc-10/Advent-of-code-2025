package dataStructures

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Empty() bool {
	return len(q.arr) == 0
}

func (q *Queue[T]) Push(val T) {
	q.arr = append(q.arr, val)
}

func (q *Queue[T]) Top() (T, bool) {
	if q.Empty() {
		var ret T
		return ret, false
	}
	return q.arr[0], true
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Empty() {
		var ret T
		return ret, false
	}
	ret := q.arr[0]
	q.arr = q.arr[1:]
	return ret, true
}

func (q *Queue[T]) Length() int {
	return len(q.arr)
}
