package exception

type BadRequestErr struct {
	msg string
}

func (error *BadRequestErr) Error() string {
	return error.msg
}

func BadRequestError(msg string) error {
	return &BadRequestErr{
		msg: msg,
	}
}
