package blocks

import (
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
