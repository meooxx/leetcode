package main

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter *Iterator
	cur  int
	peeked bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, 0, false}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.peeked {
		this.peeked = false
		return this.cur
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.peeked {
		return this.cur
	}
	this.cur = this.iter.next()
	this.peeked = true
	return this.cur
}
