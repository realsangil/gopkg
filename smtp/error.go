package smtp

// Error means contants error and implements error and
type Error string

func (e Error) String() string {
	return string(e)
}

func (e Error) Error() string {
	return e.String()
}
