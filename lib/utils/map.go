package utils

func SliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, v := range slice {
		set[v] = true
	}
	return set
}
