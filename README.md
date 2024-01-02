# GoLang Quality of Life Pack

Collection of functions for improving the quality of life writing Golang code.

This package provides utility functions to perform operations on slices and maps. The operations include:

- `Delta` : This function calculates the difference between two slices. It returns the elements that must be added to or removed from the source slice to make it match the target slice.

- `Contains` : This function checks whether a specified element exists in a provided slice.

- `ContainsFunc` : This function is similar to `Contains` but takes in a function to perform a custom check on the elements.

- `EqualsIgnoreOrder` : This function compares two slices, and returns true if they contain the same elements, regardless of their order.

- `Unique` : This function removes duplicate elements from a slice.

- `Intersect` : This function returns a slice containing common elements between two slices.

- `Map` : This function applies a provided function on each element of a given slice.

- `MapNotNull` : Similar to the `Map` function, but this function only considers non-nil results.

- `SpliceMapToKeyValue` : This function converts a map into two slices - one containing the keys and another containing the values.

- `GetKeysOfMap` : This function fetches all keys from a provided map and returns them as a slice.

- `GetValuesOfMap` : This function fetches all values from a given map, and returns them as a slice.

- `SortStrings` : This function sorts a slice of strings in alphabetical order.

The package utilizes the powerful features available in Go programming language, such as interfaces, functions as first-class citizens, and the 'sort' package.