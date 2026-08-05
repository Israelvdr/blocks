package blocks

import (
	"fmt"
	"os"
	"strconv"
)

// Fixed bounding values
const (
	MIN_BLOCK_SIZE   int = 2
	MIN_NUM_ELEMENTS int = 2
)

// Reasonable default configurable bounding values
const (
	DEFAULT_MAX_BLOCK_SIZE   int = 4
	DEFAULT_MAX_NUM_ELEMENTS int = 256
)

// Hard limits on configurabel bounding values to enforce computational solvability
const (
	HARD_LIMIT_MAX_BLOCK_SIZE int = 16
)

// Environment variables for configurable bounding values
const (
	ENV_VAR_NAME_MAX_BLOCK_SIZE   string = "MAX_BLOCK_SIZE"
	ENV_VAR_NAME_MAX_NUM_ELEMENTS string = "MAX_NUM_ELEMENTS"
)

// Global variables for configurable bounding values
// These are used for any actual validations
var (
	maxBlockSize   int = DEFAULT_MAX_BLOCK_SIZE
	maxNumElements int = DEFAULT_MAX_NUM_ELEMENTS
)

func init() {
	// Parse configurable maximum block size from environment variable and validate
	envMaxBlockSizeString := os.Getenv(ENV_VAR_NAME_MAX_BLOCK_SIZE)
	envMaxBlockSize, _ := strconv.Atoi(envMaxBlockSizeString)
	if envMaxBlockSize > MIN_BLOCK_SIZE && envMaxBlockSize < HARD_LIMIT_MAX_BLOCK_SIZE {
		maxBlockSize = envMaxBlockSize
	}

	// Parse configurable maximum number of elements from environment variable and validate
	envMaxNumElementsString := os.Getenv(ENV_VAR_NAME_MAX_NUM_ELEMENTS)
	envMaxNumElements, _ := strconv.Atoi(envMaxNumElementsString)
	if envMaxNumElements > MIN_BLOCK_SIZE && envMaxNumElements < HARD_LIMIT_MAX_BLOCK_SIZE {
		maxNumElements = envMaxNumElements
	}
}

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

type ErrBlockInvalid interface {
	ErrBlockSizeOOB |
		ErrNumElementsOOB |
		ErrBlockSizeGreaterThanNumElements
}

type ErrBlockSizeOOB struct {
	blockSize int
}

func (e ErrBlockSizeOOB) Error() string {
	return fmt.Sprintf(
		"invalid block size; expected %d <= numBlocks <= %d, got %d",
		MIN_BLOCK_SIZE, maxBlockSize, e.blockSize,
	)
}

type ErrNumElementsOOB struct {
	numElements int
}

func (e ErrNumElementsOOB) Error() string {
	return fmt.Sprintf(
		"invalid number of elements; expected %d <= numBlocks <= %d, got %d",
		MIN_NUM_ELEMENTS, maxNumElements, e.numElements,
	)
}

type ErrBlockSizeGreaterThanNumElements struct {
	blockSize   int
	numElements int
}

func (e ErrBlockSizeGreaterThanNumElements) Error() string {
	return fmt.Sprintf(
		"block size must be less than number of elements; got %d block size for %d elements",
		e.blockSize, e.numElements,
	)
}
