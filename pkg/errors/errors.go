package errors

// Error собственный тип на базе встроенного типа string
type Error string

func (e Error) Error() string {
	return string(e)
}
