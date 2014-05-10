package llrb

type ItemIterator func(i Item) bool

func (t *LLRB) Ascend(iterator ItemIterator) {
	ascend(t.root, iterator)
}

func ascend(h *Node, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if !ascend(h.Left, iterator) {
		return false
	}
	if !iterator(h.Item) {
		return false
	}
	return ascend(h.Right, iterator)
}

func (t *LLRB) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	ascendRange(t.root, greaterOrEqual, lessThan, iterator)
}

func ascendRange(h *Node, inf, sup Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if !less(h.Item, sup) {
		return ascendRange(h.Left, inf, sup, iterator)
	}
	if less(h.Item, inf) {
		return ascendRange(h.Right, inf, sup, iterator)
	}

	if !ascendRange(h.Left, inf, sup, iterator) {
		return false
	}
	if !iterator(h.Item) {
		return false
	}
	return ascendRange(h.Right, inf, sup, iterator)
}

// AscendGreaterOrEqual will call iterator once for each element greater or equal to
// pivot in ascending order. It will stop whenever the iterator returns false.
func (t *LLRB) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	ascendGreaterOrEqual(t.root, pivot, iterator)
}

func ascendGreaterOrEqual(h *Node, pivot Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if !less(h.Item, pivot) {
		if !ascendGreaterOrEqual(h.Left, pivot, iterator) {
			return false
		}
		if !iterator(h.Item) {
			return false
		}
	}
	return ascendGreaterOrEqual(h.Right, pivot, iterator)
}

func (t *LLRB) AscendLessThan(pivot Item, iterator ItemIterator) {
	ascendLessThan(t.root, pivot, iterator)
}

func ascendLessThan(h *Node, pivot Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if !ascendLessThan(h.Left, pivot, iterator) {
		return false
	}
	if !iterator(h.Item) {
		return false
	}
	if less(h.Item, pivot) {
		return ascendLessThan(h.Left, pivot, iterator)
	}
	return true
}

// DescendLessOrEqual will call iterator once for each element less than the
// pivot in descending order. It will stop whenever the iterator returns false.
func (t *LLRB) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
	descendLessOrEqual(t.root, pivot, iterator)
}

func descendLessOrEqual(h *Node, pivot Item, iterator ItemIterator) bool {
	if h == nil {
		return true
	}
	if less(h.Item, pivot) || !less(pivot, h.Item) {
		if !descendLessOrEqual(h.Right, pivot, iterator) {
			return false
		}
		if !iterator(h.Item) {
			return false
		}
	}
	return descendLessOrEqual(h.Left, pivot, iterator)
}
