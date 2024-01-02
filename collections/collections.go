package collections

import "sort"

// Delta function is used to calculate the difference between two slices.
// It compares 'source' and 'target' slices and determines which elements are to be added or removed from the 'source'
// slice to match the 'target' slice.
func Delta[T comparable](source []T, target []T) (adding []T, removing []T) {

	// Prepare slices to store elements that need to be added or removed
	var (
		elementsToBeAdded   = make([]T, 0)
		elementsToBeRemoved = make([]T, 0)
	)

	// Loop over 'source' slice to identify items missing in the 'target'. These will be added.
	for _, missingItem := range target {
		// If a 'source item' doesn't exist in 'target', it needs to be added
		if !Contains(source, missingItem) {
			elementsToBeAdded = append(elementsToBeAdded, missingItem)
		}
	}

	// Similar to previous loop, iterate over 'target' slice to determine which items are not present in 'source'. These will be removed.
	for _, extraItem := range source {
		// If a 'target item' doesn't exist in 'source', it needs to be removed
		if !Contains(target, extraItem) {
			elementsToBeRemoved = append(elementsToBeRemoved, extraItem)
		}
	}

	// Return both slice of elements that need to be added and removed to sync 'target' with 'source'
	return elementsToBeAdded, elementsToBeRemoved
}

// Contains is a helper function that checks whether a specific element ('check') exists in the provided slice ('source')
func Contains[T comparable](source []T, check T) bool {
	// Loop over each item in the slice
	for _, v := range source {
		// If the current item matches the item we're checking for, return 'true'
		if v == check {
			return true
		}
	}
	// If we've looped over all the items and haven't found a match, return 'false'
	return false
}

// ContainsFunc is a helper function that checks whether a specific element ('check') exists in the provided slice ('source')
func ContainsFunc[T comparable](source []T, check func(e T) bool) bool {
	// Loop over each item in the slice
	for _, v := range source {
		// If the current item matches the item we're checking for, return 'true'
		if check(v) {
			return true
		}
	}
	// If we've looped over all the items and haven't found a match, return 'false'
	return false
}

// EqualsIgnoreOrder compares two slices and returns 'true' if they contain the same elements, regardless of their order
func EqualsIgnoreOrder[T comparable](source, other []T) bool {
	// If the lengths of both slices aren't equal, they can't be the same
	if len(source) != len(other) {
		return false
	}

	// Otherwise, start by counting the occurrences of each element in 'source'
	counts := make(map[T]int)
	for _, item := range source {
		counts[item]++
	}

	// Then, reduce the count for each 'other' item in the map
	for _, item := range other {
		// If an element doesn't exist in 'source', the slices aren't the same
		if counts[item] == 0 {
			return false
		}
		counts[item]--
	}

	// If all counts are zero, the slices are equal, return 'true'
	return true
}

// Unique function goes through a slice and eliminates the duplicate elements
func Unique[T comparable](collection []T) []T {
	// Prepare a map to hold unique elements
	uniqueMap := make(map[T]bool)
	uniqueSlice := make([]T, 0)

	// Iterate over the collection and add each unique item to both our map and our result slice
	for _, value := range collection {
		if _, exists := uniqueMap[value]; !exists {
			uniqueMap[value] = true
			uniqueSlice = append(uniqueSlice, value)
		}
	}

	// Return slice of unique elements
	return uniqueSlice
}

// Intersect function returns a slice containing common elements between 'source' and 'other'.
func Intersect[T comparable](source, other []T) []T {
	// Map to keep track of elements in 'source'
	elemMap := make(map[T]bool)
	// Resulting slice to store common elements
	result := make([]T, 0)

	// Loop over 'source' and add each element to the map
	for _, v := range source {
		elemMap[v] = true
	}

	// Loop over 'other'. If an element also exists in our map, add it to the result
	for _, v := range other {
		if _, ok := elemMap[v]; ok {
			result = append(result, v)
			// Avoid duplicates in the result by deleting the element from the map
			delete(elemMap, v)
		}
	}

	// Return the slice of common elements
	return result
}

// Map function applies a provided function 'mapper' on each element of a given slice.
// It's a generalized function to perform any operation 'mapper' on all elements
func Map[K any](collection []K, mapper func(e K) K) []K {
	slice := make([]K, 0)

	// For each element in 'collection', apply the 'mapper' function and store the result in 'slice'
	for _, item := range collection {
		slice = append(slice, mapper(item))
	}

	// Return the slice containing results of 'mapper' function applied on each item of 'collection'
	return slice
}

// MapNotNull function applies a function 'mapper' on each element of a slice, but only considers non-nil results.
func MapNotNull[K any](collection []K, mapper func(k K) *K) []K {
	slice := make([]K, 0)

	// Loop over 'collection'. If 'mapper' returns a non-nil value, add it to 'slice'
	for _, item := range collection {
		k := mapper(item)
		if k != nil {
			slice = append(slice, *k)
		}
	}

	// Return the slice containing non-nil results
	return slice
}

// SpliceMapToKeyValue function converts a map into two slices - one containing the keys and another containing the values.
func SpliceMapToKeyValue[K comparable, V any](source map[K]V) ([]K, []V) {
	keys := make([]K, 0)
	values := make([]V, 0)

	// Loop over the source map and collect its values
	for key, value := range source {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

// GetKeysOfMap function fetches all keys from the supplied map and returns them as a slice.
func GetKeysOfMap[K comparable, V any](source map[K]V) []K {
	keys, _ := SpliceMapToKeyValue(source)
	return keys
}

// GetValuesOfMap function fetches all values from the supplied map and returns them as a slice.
func GetValuesOfMap[K comparable, V any](source map[K]V) []V {
	_, values := SpliceMapToKeyValue(source)
	return values
}

// SortStrings function sorts a slice of strings in alphabetical order.
func SortStrings(source []string) []string {
	sort.Strings(source)
	// Return sorted slice
	return source
}
