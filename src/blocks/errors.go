package blocks

import "fmt"

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

type ErrUnknownDesignFamily struct {
	blockBase *blockDesignBase
}

func (e *ErrUnknownDesignFamily) Error() string {
	return "failed to classify block family"
}
