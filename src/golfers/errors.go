package golfers

import "fmt"

// TODO: Do i even need this interface?
type blockError interface {
	Error() string
	Is(error) bool
	IsType(error) bool
}

type ErrBlockSwapIndexOOB struct {
	i            int
	elementCount int
}

func (e ErrBlockSwapIndexOOB) Error() string {
	return fmt.Sprintf(
		"swap element index out of bounds, received index %d, block only contains %d elements",
		e.i, e.elementCount,
	)
}
func (e ErrBlockSwapIndexOOB) Is(target error) bool {
	t, ok := target.(ErrBlockSwapIndexOOB)
	return ok && e.i == t.i && e.elementCount == t.elementCount
}
func (e ErrBlockSwapIndexOOB) IsType(target error) bool {
	_, ok := target.(ErrBlockSwapIndexOOB)
	return ok
}

type ErrBlockDesignUninitialised struct {
}

func (e ErrBlockDesignUninitialised) Error() string {
	return "block design has not been initialised; cannot solve or optimise"
}
func (e ErrBlockDesignUninitialised) Is(target error) bool {
	_, ok := target.(ErrBlockDesignUninitialised)
	return ok
}
func (e ErrBlockDesignUninitialised) IsType(target error) bool {
	_, ok := target.(ErrBlockDesignUninitialised)
	return ok
}

type ErrBlockDesignUnsolved struct {
}

func (e ErrBlockDesignUnsolved) Error() string {
	return "block design has not been solved; cannot evaluate"
}
func (e ErrBlockDesignUnsolved) Is(target error) bool {
	_, ok := target.(ErrBlockDesignUnsolved)
	return ok
}
func (e ErrBlockDesignUnsolved) IsType(target error) bool {
	_, ok := target.(ErrBlockDesignUnsolved)
	return ok
}
