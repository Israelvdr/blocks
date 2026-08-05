package golfers

type RoundwiseBlockDesign[T comparable] struct {
	blockDesign[T]
}

func (bd *RoundwiseBlockDesign[T]) Solve() {}

func (bd *RoundwiseBlockDesign[T]) Optimise() {}
