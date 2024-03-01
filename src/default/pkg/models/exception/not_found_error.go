package exception

import (
	"fmt"
	"strings"
)

type NotFoundErr struct {
	id       string
	resource string
	msg      string
}

func (error *NotFoundErr) Error() string {
	return strings.TrimSpace(fmt.Sprintf("%s with identifier %s not found. %s", error.resource, error.id, error.msg))
}

func NotFoundError(id string, resource string, msg string) error {
	return &NotFoundErr{
		id:       id,
		resource: resource,
		msg:      msg,
	}
}
