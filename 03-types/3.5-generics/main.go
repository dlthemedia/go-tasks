package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func echo[T any](v T) T {
	return v
}

func first[T any](items []T) (T, bool) {

	if len(items) == 0 {
		var null T
		return null, false
	}
	return items[0], true
}

func sum[T Number](items []T) T {
	var total T
	for _, v := range items {
		total += v
	}
	return total
}

func IndexOf[T comparable](items []T, target T) int {

	if len(items) == 0 {
		return -2
	}

	for val, _ := range items {
		if items[val] == target {
			return val
		}
	}
	return -1

}

func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {

	fmt.Println(echo("Say my Name"))

	nums := []int{10, 20}
	n, ok := first(nums)
	fmt.Println(n, ok)

	empty := []string{}
	s, ok := first(empty)
	fmt.Println(s, ok)

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(ints))

	floats := []float64{1.5, 2.5, 3.0}
	fmt.Println(sum(floats))

	ints = []int{1, 2, 3, 4, 5}
	fmt.Println(IndexOf(ints, 3))

	intPair := Pair[int, int]{Key: 52, Value: 42}
	strPair := Pair[string, string]{Key: "Donbass", Value: "course"}

	fmt.Println(intPair.Value)
	fmt.Println(strPair.Value)

	m := map[string]int{
		"BLACK apple": 5,
		"Big banana":  3,
		"Goida":       8,
	}

	values := Values(m)
	fmt.Println(values)

	stringStore := &Store[string]{}
	stringStore.Add("hello")
	stringStore.Add("world")

	fmt.Println(stringStore.All())

	intStore := &Store[int]{}
	intStore.Add(42)
	intStore.Add(100)

	fmt.Println(intStore.All())

	// 1. Store для строк
	stringStore = &Store[string]{}
	stringStore.Add("hello")
	stringStore.Add("world")
	stringStore.Add("golang")

	fmt.Println("String store:", stringStore.All())
	fmt.Println("Contains 'world':", Contains(stringStore.All(), "world"))
	fmt.Println("Contains 'python':", Contains(stringStore.All(), "python"))

	// 2. Store для чисел
	intStore = &Store[int]{}
	intStore.Add(10)
	intStore.Add(20)
	intStore.Add(30)

	fmt.Println("\nInt store:", intStore.All())
	fmt.Println("Contains 20:", Contains(intStore.All(), 20))
	fmt.Println("Contains 99:", Contains(intStore.All(), 99))

	// 3. Дополнительно — проверка с пустым слайсом
	emptyStore := &Store[float64]{}
	fmt.Println("\nEmpty store:", emptyStore.All())
	fmt.Println("Contains 3.14:", Contains(emptyStore.All(), 3.14))

}
