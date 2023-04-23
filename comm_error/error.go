package comm_error

import "fmt"

type CommonError struct {
	Code uint32
	Msg  string
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("CommonError: code=%d, msg=%s", e.Code, e.Msg)
}

var (
	// business error code: [500000, 600000)
	ErrWrongType      = &CommonError{500000, "wrong type"}
	ErrGetColumnIndex = &CommonError{500001, "get column index failed"}
	ErrNoData         = &CommonError{500002, "no data"}
	ErrWrongColumnNum = &CommonError{500003, "column number is wrong"}
)
