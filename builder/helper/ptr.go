package helper

// ToPtr converts a given var T into *T
func ToPtr[T any](v T) *T {
	return &v
}
