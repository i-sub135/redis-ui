package helpers

import (
	"log"
)

// DerefOrDefault converts pointer to value, returns zero value if nil
func DerefOrDefault[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}
	var zero T
	return zero
}

// DerefOrValue converts pointer to value, returns fallback if nil
func DerefOrValue[T any](ptr *T, fallback T) T {
	if ptr != nil {
		return *ptr
	}
	return fallback
}

func AsPtr[T any](v T) *T {
	return &v
}

func IgnorErr[T any](fn T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return fn
}
