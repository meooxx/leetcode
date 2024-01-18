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
type Iterator struct {
}
type PeekingIterator struct {
	NextVal int
	Peeked  bool
	Iter    *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		NextVal: 0,
		Peeked:  false,
		Iter:    iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.Peeked {
		return true
	}
	return this.Iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.Peeked {
		this.Peeked = false
		return this.NextVal
	}
	return this.Iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.Peeked {
		return this.NextVal
	}
	this.Peeked = true
	this.NextVal = this.Iter.next()
	return this.NextVal
}
