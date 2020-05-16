package drive

type ErrorWithContext struct {
	ServiceAccountFile string
	err                error
}

func (e *ErrorWithContext) Error() string {
	return e.err.Error()
}

func NewErrorWithContext(err error, serviceAccountFile string) *ErrorWithContext {
	return &ErrorWithContext{
		ServiceAccountFile: serviceAccountFile,
		err:                err,
	}
}
