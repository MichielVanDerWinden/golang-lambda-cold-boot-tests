package exception

type InternalServerErr struct {
	msg string
}

func (error *InternalServerErr) Error() string {
	return error.msg
}

func InternalServerError(msg string) error {
	return &InternalServerErr{
		msg: msg,
	}
}
