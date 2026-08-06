package blocks

type blockDesignBase struct {
	numElements      int     // v
	blockSize        int     // k
	blocksPerPair    int     // lambda
	numBlocks        int     // b
	blocksPerElement int     // r
	blocks           [][]int // B
	family           blockDesign
}

func NewPBD(numElements, blockSize, blocksPerElement int) (blockDesign, error) {
	pbd := &blockDesignBase{
		numElements:      numElements,
		blockSize:        blockSize,
		blocksPerElement: blocksPerElement,
	}

	return pbd.init()
}

func (pbd *blockDesignBase) init() (blockDesign, error) {
	// Validate inputs
	err := pbd.ValidateInputParams()
	if err != nil {
		return &invalidBlockDesign{
			blockDesignBase: *pbd,
			err:             err,
		}, err
	}

	// Below comments are only for balanced, complete designs and should be moved to a specific solver
	// Given v, k, & target r; calculate lambda, & b
	// bk = vr
	//lambda(v-1) = r(k-1)

	// Initialise empty blocks 2D slice
	pbd.blocks = make([][]int, pbd.numBlocks)
	for i := range pbd.numBlocks {
		pbd.blocks[i] = make([]int, pbd.blockSize)
	}

	return pbd.classify(), nil
}

func (pbd *blockDesignBase) classify() blockDesign {
	switch {
	default:
		return &invalidBlockDesign{
			blockDesignBase: *pbd,
			err: &ErrUnknownDesignFamily{
				blockBase: pbd,
			},
		}
	}
}

type blockDesign interface {
	designFamilyName() string
	designFamilyDescription() string
	solve() blockDesign
	optimise() blockDesign
}

// type projectionBlockDesign struct{}
// type steinerBlockDesign struct{}
// type completeBlockDesign struct{}
// type balancedIncompleteBlockDesign struct{}
// type partiallyBalancedIncompleteBlockDesign struct{}
