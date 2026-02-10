package main

import "sync"

type span struct {
	size      int
	allocated bool
}

type mheap struct {
	spans []*span
	lock  sync.Mutex
}

type mcentral struct {
	sizeSpans []*span
	lock      sync.Mutex
}

type mcache struct {
	localSpan []*span
}

func newHeap(size int) *mheap {
	h := &mheap{}
	for i := range size {
		h.spans = append(h.spans, &span{size: i + 1})
	}
	return h
}

func (h *mheap) getSpan(size int) *span {
	h.lock.Lock()
	defer h.lock.Unlock()

	for _, s := range h.spans {
		if s.size >= size && !s.allocated {
			s.allocated = true
			return s
		}
	}
	return nil
}

func (mc *mcentral) getSpanFromCentral(size int) *span {
	mc.lock.Lock()
	defer mc.lock.Unlock()

	for _, s := range mc.sizeSpans {
		if s.size >= size && !s.allocated {
			s.allocated = true
			return s
		}
	}
	return nil
}

func (mc *mcache) getSpanFromCache(size int) *span {
	for _, s := range mc.localSpan {
		if s.size >= size && !s.allocated {
			s.allocated = true
			return s
		}
	}
	return nil
}

func main() {
	heap := newHeap(10)
	mcentral := &mcentral{sizeSpans: heap.spans}
	mcache := &mcache{}

	requestSpan := mcache.getSpanFromCache(5)
	if requestSpan == nil {
		requestSpan = mcentral.getSpanFromCentral(5)
	}
	if requestSpan == nil {
		requestSpan = heap.getSpan(5)
	}

	if requestSpan == nil {
		panic("out of memory")
	}

	println("Allocated span of size:", requestSpan.size)
}
