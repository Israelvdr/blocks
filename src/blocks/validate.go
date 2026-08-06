package blocks

import (
	"fmt"
)

const (
	INVALID_BLOCK_DESIGN_NAME        string = "Invalid Block Design"
	INVALID_BLOCK_DESIGN_DESCRIPTION string = "Block design is invalid; check initial paramters."
)

func (pbd *blockDesignBase) ValidateInputParams() error {
	// Validate params

	// blockSize >= 2
	if pbd.blockSize < MIN_BLOCK_SIZE || pbd.blockSize > maxBlockSize {
		return ErrBlockSizeOOB{blockSize: pbd.blockSize}
	}

	// numElements >= 2
	// technically redundant due to below check, but error type provides detail
	if pbd.numElements < MIN_NUM_ELEMENTS || pbd.numElements > maxNumElements {
		return ErrNumElementsOOB{numElements: pbd.blockSize}
	}

	// blockSize < numElements
	if pbd.blockSize < pbd.numElements {
		return ErrBlockSizeGreaterThanNumElements{
			blockSize:   pbd.blockSize,
			numElements: pbd.numElements,
		}
	}

	return nil
}

func (pbd *blockDesignBase) ValidateBlocks() error {
	return fmt.Errorf("not implemented")
}

type invalidBlockDesign struct {
	blockDesignBase
	err error
}

func (ibd *invalidBlockDesign) designFamilyName() string {
	return INVALID_BLOCK_DESIGN_NAME
}

func (ibd *invalidBlockDesign) designFamilyDescription() string {
	return INVALID_BLOCK_DESIGN_DESCRIPTION
}

func (ibd *invalidBlockDesign) solve() blockDesign {
	return ibd
}

func (ibd *invalidBlockDesign) optimise() blockDesign {
	return ibd
}
