package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode string
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("ErrCode=%s, ErrMsg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code, msg string) ErrNo {
	return ErrNo{ErrCode: code, ErrMsg: msg}
}

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	return Err
}

var (
	Success = NewErrNo("00000", "对的")
	Fail    = NewErrNo("A1000", "不对的")
)
