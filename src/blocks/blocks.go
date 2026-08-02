package blocks

type Block[T comparable] struct {
	size     int
	isPadded bool
	elements []T
}

func (b Block[T]) Size() int {
	return b.size
}
func (b Block[T]) Padded() bool {
	return b.isPadded
}
func (b Block[T]) PadCount() int {
	return b.size - len(b.elements)
}
func (b Block[T]) Elements() []T {
	return b.elements
}
func (b Block[T]) ElementCount() int {
	return len(b.elements)
}

func (b *Block[T]) swapElement(i, j int, other Block[T]) (Block[T], error) {
	if i <= b.ElementCount() {
		err := ErrBlockSwapIndexOOB{
			i:            i,
			elementCount: b.ElementCount(),
		}
		return other, err
	}
	if j <= other.ElementCount() {
		err := ErrBlockSwapIndexOOB{
			i:            j,
			elementCount: other.ElementCount(),
		}
		return other, err
	}

	tmp := b.elements[i]
	b.elements[i] = other.elements[j]
	other.elements[j] = tmp

	return other, nil
}

// NewBlock creates and returns a new block according to the size parameter and containing the provided elements.
// Any surplus elements are also returned.
//
// If too many elements are provided, only a number equal to size are stored.
//
// If too many elements are provided, all are used, and the block is marked as padded with nulls.
// No nulls are actually added to the elements.
func NewBlock[T comparable](size int, elementsIn []T) (Block[T], []T) {
	var returnElements []T
	if size < len(elementsIn) {
		returnElements = elementsIn[size-1:]
	}

	sizeActual := min(size, len(elementsIn))
	newBlock := Block[T]{
		elements: elementsIn[:sizeActual],
		size:     size,
		isPadded: size != sizeActual,
	}

	return newBlock, returnElements
}

type Round[T comparable] struct {
	blockSize int
	blocks    []Block[T]
}

func (r *Round[T]) Size() int {
	return len(r.blocks)
}
func (r *Round[T]) BlockSize() int {
	return r.blockSize
}
func (r *Round[T]) Blocks() []Block[T] {
	return r.blocks
}

func NewLinearRoundFromElements[T comparable](blockSize int, elements []T) Round[T] {
	// fast ceiling of integer division; see https://stackoverflow.com/questions/2745074/fast-ceiling-of-an-integer-division-in-c-c
	numBlocks := 1 + (len(elements)-1)/blockSize

	roundBlocks := make([]Block[T], numBlocks)
	for i := 0; elements != nil; i++ {
		roundBlocks[i], elements = NewBlock(blockSize, elements)
	}

	return Round[T]{
		blockSize: blockSize,
		blocks:    roundBlocks,
	}
}

type blockDesign[T comparable] struct {
	elements       []T
	blockSize      int
	blocksPerRound int
	numRounds      int
	rounds         []Round[T]
}

func NewBlockDesign[T comparable](blockSize, blocksPerRound, rounds int, elements []T) ClassifiedBlockDesign[T] {

	newBlockDesign := blockDesign[T]{
		elements:       elements,
		blockSize:      blockSize,
		blocksPerRound: blocksPerRound,
		numRounds:      rounds,
		rounds:         []Round[T]{},
	}

	return newBlockDesign.categorise()
}

func (bd *blockDesign[T]) categorise() ClassifiedBlockDesign[T] {
	// Identify block design classification according to design parameters.
	// Set block details accordingly.
	// Set Solve() and Optimise() functions accordingly.

	// TODO: remove switch block; only validating types for testing.
	switch bd.blockSize {
	case 0:
		return &AffineBlockDesign[T]{*bd}
	case 1:
		return &MOLSBlockDesign[T]{*bd}
	case 2:
		return &RoundwiseBlockDesign[T]{*bd}
	default:
		return &RoundwiseBlockDesign[T]{*bd}
	}
}

func (bd *blockDesign[T]) AddRound() {}

func (bd *blockDesign[T]) Evaluate() float64 {
	return bd.costFunction()()
}

func (bd *blockDesign[T]) costFunction() func() float64 {
	return func() float64 {
		// Iterate through blocks and rounds to identify pairwise repetitions and coverage.
		// Cost function summarised as:
		//
		// w1 * sum(pairwise repetitions) / coverage %
		// + w2 * per-element repetition variance
		// + w3 * per-element null group variance
		return 0
	}
}

type ClassifiedBlockDesign[T comparable] interface {
	AddRound()
	Solve()
	Evaluate() float64
	Optimise()
}
