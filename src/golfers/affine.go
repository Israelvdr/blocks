package golfers

type AffineBlockDesign[T comparable] struct {
	blockDesign[T]
}

func (bd *blockDesign[T]) Solve() {}

func (bd *blockDesign[T]) Optimise() {}
