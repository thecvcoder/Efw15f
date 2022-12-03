package utils

const (
	SystemErr = 500 + iota
	MysqlErr
	Validator
)

type RspError struct {
	code int
	err  error
}

func (r *RspError) Error() string {
	return r.err.Error()
}

func (r *RspError) Code() int {
	return r.code
}

// NewRspError New
func NewRspError(code int, err error) *RspError {
	return &RspError{
		code: code,
		err:  err,
	}
}
