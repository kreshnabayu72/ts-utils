package functions

func MakePointer[T any](v T) *T {
	return &v
}
