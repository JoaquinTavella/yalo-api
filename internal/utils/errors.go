package utils

import (
	"fmt"
)

type Error struct {
	message string
	parent  error
	domain  string
}

func NewError(pError error, msg string, domain interface{}, p ...any) Error {
	return Error{
		message: fmt.Sprintf(msg, p...),
		parent:  pError,
		domain:  fmt.Sprintf("%T", domain),
	}
}

func (rErr Error) Error() string {
	if rErr.parent == nil {
		return rErr.message
	}

	return fmt.Sprintf("%s -> %s", rErr.message, rErr.parent.Error())
}
