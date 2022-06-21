package request

import "errors"

type E error

var (
	ErrConfig      E = errors.New("config request is empty")
	ErrValidate    E = errors.New("error validate the config request")
	ErrIniRequest  E = errors.New("error initialize client request")
	ErrReadData    E = errors.New("error read response data")
	ErrContextDone E = errors.New("error context is done")
	ErrServer      E = errors.New("can not reach server")
	ErrParse       E = errors.New("error parse the response data")
	ErrMethod      E = errors.New("error method is empty")
)
