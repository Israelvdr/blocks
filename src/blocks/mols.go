package blocks

type MOLSBlockDesign[T comparable] struct {
	blockDesign[T]
}

func (bd *MOLSBlockDesign[T]) Solve() {}

func (bd *MOLSBlockDesign[T]) Optimise() {}
