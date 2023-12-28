package collections

import "sort"

func Delta[T comparable](target []T, source []T) (adding []T, removing []T) {
	var (
		elementsToBeAdded   = make([]T, 0)
		elementsToBeRemoved = make([]T, 0)
	)

	for _, missingItem := range source {
		if !Contains(target, missingItem) {
			elementsToBeAdded = append(elementsToBeAdded, missingItem)
		}
	}

	for _, extraItem := range target {
		if !Contains(source, extraItem) {
			elementsToBeRemoved = append(elementsToBeRemoved, extraItem)
		}
	}

	return elementsToBeAdded, elementsToBeRemoved
}

func Contains[T comparable](source []T, check T) bool {
	for _, v := range source {
		if v == check {
			return true
		}
	}
	return false
}

func EqualsIgnoreOrder[T comparable](source, other []T) bool {
	if len(source) != len(other) {
		return false
	}

	counts := make(map[T]int)
	for _, item := range source {
		counts[item]++
	}

	for _, item := range other {
		if counts[item] == 0 {
			return false
		}
		counts[item]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func Unique[T comparable](collection []T) []T {
	uniqueMap := make(map[T]bool)
	var uniqueSlice = make([]T, 0)

	for _, value := range collection {
		if _, exists := uniqueMap[value]; !exists {
			uniqueMap[value] = true
			uniqueSlice = append(uniqueSlice, value)
		}
	}

	return uniqueSlice
}

func Intersect[T comparable](source, other []T) []T {
	elemMap := make(map[T]bool)
	var result = make([]T, 0)

	for _, v := range source {
		elemMap[v] = true
	}

	for _, v := range other {
		if _, ok := elemMap[v]; ok {
			result = append(result, v)
			delete(elemMap, v)
		}
	}

	return result
}

func Map[K any](collection []K, mapper func(k K) K) []K {
	var slice []K
	for _, item := range collection {
		slice = append(slice, mapper(item))
	}

	return slice
}

func MapNotNull[K any](collection []K, mapper func(k K) *K) []K {
	var slice []K
	for _, item := range collection {
		k := mapper(item)
		if k != nil {
			slice = append(slice, *k)
		}
	}

	return slice
}

func MapValues[K comparable, V any](source map[K]V) []V {
	var values []V

	for _, value := range source {
		values = append(values, value)
	}

	return values
}

func SortStrings(source []string) []string {
	sort.Strings(source)
	return source
}
