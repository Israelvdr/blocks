package pbds

import "fmt"

type pairwiseBalancedDesign struct {
	numElements      int     // v
	blockSize        int     // k
	blocksPerPair    int     // lambda
	numBlocks        int     // b
	blocksPerElement int     // r
	elements         []int   // X = { 0:v-1 } // len(elements) = v
	blocks           [][]int // B
}

func NewPBD(numElements, blockSize, blocksPerElement int) (*pairwiseBalancedDesign, error) {
	pbd := &pairwiseBalancedDesign{
		numElements:      numElements,
		blockSize:        blockSize,
		blocksPerElement: blocksPerElement,
	}

	pbd, err := pbd.calcRemainingParams()

	return pbd, err
}

func (pbd *pairwiseBalancedDesign) init() (*pairwiseBalancedDesign, error) {
	// Given v, k, & target r; calculate lambda, & b

	// Validate inputs

	// Initialise elements slice

	// Initialise empty blocks 2D slice

	return nil, fmt.Errorf("not implemented")
}

func (pbd *pairwiseBalancedDesign) calcRemainingParams() (*pairwiseBalancedDesign, error) {

	return nil, fmt.Errorf("not implemented")
}
