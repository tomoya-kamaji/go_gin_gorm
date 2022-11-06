package pointerLib


func ToPointer[T any](param T) *T {
	return &param
}
