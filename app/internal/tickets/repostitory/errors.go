package tickets

import "errors"

// errors
var (
	ErrLoadData  = errors.New("error loading data")
	ErrEmptyList = errors.New("empty list")
)
