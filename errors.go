package onelogin

// ErrNotImplemented is returned when a method is not implemented
type ErrNotImplemented struct{}

func (e ErrNotImplemented) Error() string {
	return "not implemented"
}

// ErrMissingField is returned when a required field is missing
type ErrMissingField struct {
	Field string
}

func (e ErrMissingField) Error() string {
	return "missing field: " + e.Field
}
