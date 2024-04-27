package errox

type Exception struct {
	Code    int
	Message string
}

func WithMessageCode(message string, code int) (e Exception) {
	if message != "" {
		e.Message = message
	}
	e.Code = code
	return e
}
func WithMessage(message string) (e Exception) {
	if message != "" {
		e.Message = message
	}
	return e
}
