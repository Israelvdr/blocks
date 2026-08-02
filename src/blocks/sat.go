package blocks

type SATBlockDesign[T comparable] struct {
	blockDesign[T]
}

func (bd *SATBlockDesign[T]) Solve() {}

func (bd *SATBlockDesign[T]) Optimise() {}
