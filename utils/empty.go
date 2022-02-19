package utils

import "reflect"

// IsEmptySlice 是否是空切片
func IsEmptySlice[T any](slice []T) bool {
	if len(slice) == 0 {
		return true
	}
	return false
}
func IsSlide[T comparable](arr []T, msg string) {
	t := reflect.TypeOf(arr)
	if t.Kind() != reflect.Slice {
		panic(msg)
	}
}
