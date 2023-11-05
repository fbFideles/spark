package utils

func GetPointer[T any](v T) *T {
	return &v
}
