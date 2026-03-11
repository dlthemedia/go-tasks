package main

type Store[T any] struct {
	items []T
}

func (s *Store[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s *Store[T]) All() []T {

	result := make([]T, 0, len(s.items))
	for _, val := range s.items {

		result = append(result, val)

	}
	return result
}

func Contains[T comparable](items []T, target T) bool {

	if len(items) == 0 {
		return false
	}

	for _, val := range items {
		if val == target {
			return true
		}
	}
	return false
}
