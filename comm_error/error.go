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
	ErrGetColumnIndex = &CommonError{500000, "get column index failed"}
	ErrNoData         = &CommonError{500001, "no data"}
)
