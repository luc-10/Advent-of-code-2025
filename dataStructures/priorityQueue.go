package dataStructures

type PriorityItem[T any] struct {
	Value    T
	Priority int
	pos      int
}

type PriorityQueue[T any] struct {
	heap []*PriorityItem[T]
	less func(a, b *PriorityItem[T]) bool
}

func NewPriorityQueue[T any](less func(a, b *PriorityItem[T]) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{heap: []*PriorityItem[T]{}, less: less}
}

func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.heap) == 0
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
	pq.heap[i].pos = i
	pq.heap[j].pos = j
}

func (pq *PriorityQueue[T]) Push(value T, priority int) *PriorityItem[T] {
	item := &PriorityItem[T]{Value: value, Priority: priority, pos: len(pq.heap)}
	pq.heap = append(pq.heap, item)

	i := item.pos
	for i > 0 {
		parent := (i - 1) / 2
		if pq.less(pq.heap[i], pq.heap[parent]) {
			pq.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
	return item
}

func (pq *PriorityQueue[T]) Top() (T, bool) {
	if pq.Empty() {
		var zero T
		return zero, false
	}
	return pq.heap[0].Value, true
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if pq.Empty() {
		var zero T
		return zero, false
	}
	pq.swap(0, len(pq.heap)-1)
	ret := pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	pq.heapRestore(0)
	return ret.Value, true
}

func (pq *PriorityQueue[T]) heapRestore(i int) {
	left := 2*i + 1
	right := 2*i + 2
	min := i
	if left < len(pq.heap) && pq.less(pq.heap[left], pq.heap[min]) {
		min = left
	}
	if right < len(pq.heap) && pq.less(pq.heap[right], pq.heap[min]) {
		min = right
	}
	if min != i {
		pq.swap(i, min)
		pq.heapRestore(min)
	}
}
