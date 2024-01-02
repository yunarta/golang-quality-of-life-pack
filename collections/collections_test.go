package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelta(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3}
		expectedAdds := []int{}
		expectedRemoves := []int{4, 5}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	// Add test cases here
	t.Run("test with empty slices", func(t *testing.T) {
		emptySlice := []int{}
		expectedAdds := []int{1, 2, 3}
		expectedRemoves := []int{}
		actualAdds, actualRemoves := Delta(emptySlice, []int{1, 2, 3})
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with strings", func(t *testing.T) {
		a := []string{"apple", "banana", "orange"}
		b := []string{"banana", "orange", "grape"}
		expectedAdds := []string{"grape"}
		expectedRemoves := []string{"apple"}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with empty string slices", func(t *testing.T) {
		emptySlice := []string{}
		expectedAdds := []string{"apple", "banana", "cherry"}
		expectedRemoves := []string{}
		actualAdds, actualRemoves := Delta(emptySlice, []string{"apple", "banana", "cherry"})
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with string slice and empty slice", func(t *testing.T) {
		emptySlice := []string{}
		expectedAdds := []string{}
		expectedRemoves := []string{"apple", "banana", "cherry"}
		actualAdds, actualRemoves := Delta([]string{"apple", "banana", "cherry"}, emptySlice)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with duplicate elements", func(t *testing.T) {
		// Test case with duplicate elements in both slices
		a := []int{1, 2, 2, 3, 4, 5}
		b := []int{2, 3, 3, 4, 5, 5}
		expectedAdds := []int{}
		expectedRemoves := []int{1}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	// New test cases
	t.Run("test with nil slices", func(t *testing.T) {
		var a, b []int
		expectedAdds, expectedRemoves := []int{}, []int{}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with nil and non-nil slices", func(t *testing.T) {
		var a []int
		b := []int{1, 2, 3, 4, 5}
		expectedAdds := []int{1, 2, 3, 4, 5}
		expectedRemoves := []int{}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with completely different integer slices", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{-1, -2, -3, -4, -5}
		expectedAdds := []int{-1, -2, -3, -4, -5}
		expectedRemoves := []int{1, 2, 3, 4, 5}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with partially different string slices", func(t *testing.T) {
		a := []string{"a", "b", "c", "d", "e"}
		b := []string{"d", "e", "f", "g", "h"}
		expectedAdds := []string{"f", "g", "h"}
		expectedRemoves := []string{"a", "b", "c"}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with same string slices", func(t *testing.T) {
		a := []string{"alea", "iacta", "est"}
		b := []string{"alea", "iacta", "est"}
		expectedAdds := []string{}
		expectedRemoves := []string{}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with different order of elements in slices", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{5, 4, 3, 2, 1}
		expectedAdds := []int{}
		expectedRemoves := []int{}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})

	t.Run("test with one empty string in one of the slices", func(t *testing.T) {
		a := []string{"alea", "iacta", ""}
		b := []string{"alea", "iacta", "est"}
		expectedAdds := []string{"est"}
		expectedRemoves := []string{""}
		actualAdds, actualRemoves := Delta(a, b)
		assert.ElementsMatch(t, expectedAdds, actualAdds)
		assert.ElementsMatch(t, expectedRemoves, actualRemoves)
	})
}

func TestContains(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}

		assert.True(t, Contains(a, 1))
		assert.False(t, Contains(a, 10))

	})

	t.Run("test with empty slice", func(t *testing.T) {
		emptySlice := []int{}
		assert.False(t, Contains(emptySlice, 1))
	})

}

func TestContainsFunc(t *testing.T) {
	t.Run("test with integers function", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		f := func(i int) bool { return i%2 == 0 }

		assert.True(t, ContainsFunc(a, f))
	})

	t.Run("test with function that does not have elements satisfying the condition", func(t *testing.T) {
		a := []int{1, 3, 5, 7, 9}
		f := func(i int) bool { return i%2 == 0 }

		assert.False(t, ContainsFunc(a, f))
	})

	t.Run("test with empty slice", func(t *testing.T) {
		emptySlice := []int{}
		f := func(i int) bool { return i%2 == 0 }

		assert.False(t, ContainsFunc(emptySlice, f))
	})

	t.Run("test with string func", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		f := func(s string) bool { return s == "banana" }

		assert.True(t, ContainsFunc(a, f))
	})

	t.Run("test with string function that does not have elements satisfying the condition", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		f := func(s string) bool { return s == "grape" }

		assert.False(t, ContainsFunc(a, f))
	})

	t.Run("test with nil slice", func(t *testing.T) {
		var a []int
		f := func(i int) bool { return i%2 == 0 }

		assert.False(t, ContainsFunc(a, f))
	})
}

func TestEqualsIgnoreOrder(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{5, 4, 3, 2, 1}

		assert.True(t, EqualsIgnoreOrder(a, b))
	})

	t.Run("test with empty slices", func(t *testing.T) {
		emptySlice := []int{}
		assert.True(t, EqualsIgnoreOrder(emptySlice, emptySlice))
	})

	t.Run("test with different integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{6, 7, 8, 9, 10}

		assert.False(t, EqualsIgnoreOrder(a, b))
	})

	t.Run("test with different number of integers", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3, 4}

		assert.False(t, EqualsIgnoreOrder(a, b))
	})

	t.Run("test with same integers but with duplicates in both slice", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 2, 3, 4, 5}

		assert.False(t, EqualsIgnoreOrder(a, b))
	})

	t.Run("test with same number of integers but different content", func(t *testing.T) {
		a := []int{11, 11, 11, 12}
		b := []int{11, 11, 12, 12}

		assert.False(t, EqualsIgnoreOrder(a, b))
	})
}

func TestUnique(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 2, 3, 3, 3}
		expected := []int{1, 2, 3}

		assert.ElementsMatch(t, expected, Unique(a))
	})

	t.Run("test with duplicate strings", func(t *testing.T) {
		// Test case with duplicate strings
		a := []string{"apple", "banana", "cherry", "banana", "apple"}
		expected := []string{"apple", "banana", "cherry"}
		actual := Unique(a)

		assert.ElementsMatch(t, actual, expected)
	})
}

func TestIntersect(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3}
		expected := []int{1, 2, 3}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with same list of integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with no intersect integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{6, 7, 8, 9, 10}
		expected := []int{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2, 3, 4, 5}
		expected := []int{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with strings", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"banana", "cherry", "grape"}
		expected := []string{"banana", "cherry"}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with same list of strings", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"apple", "banana", "cherry"}
		expected := []string{"apple", "banana", "cherry"}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with no intersect strings", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"grape", "melon", "lemon"}
		expected := []string{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with a string slice and an empty string slice", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{}
		expected := []string{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with a string slice and a nil string slice", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		var b []string
		expected := []string{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})

	t.Run("test with two empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		expected := []int{}

		assert.ElementsMatch(t, expected, Intersect(a, b))
	})
}

func TestMap(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3}
		expected := []int{2, 4, 6}
		actual := Map(a, func(k int) int { return k * 2 })

		assert.ElementsMatch(t, actual, expected)
	})
}

func TestMapNotNull(t *testing.T) {
	t.Run("test with integers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3}
		actual := MapNotNull(a, func(k int) *int {
			if k > 3 {
				return nil
			}
			return &k
		})

		assert.ElementsMatch(t, actual, expected)
	})
}

func TestGetKeysOfMap(t *testing.T) {
	t.Run("test with string to int map", func(t *testing.T) {
		a := map[string]int{"one": 1, "two": 2, "three": 3}
		expected := []string{"one", "two", "three"}
		actual := GetKeysOfMap(a)

		assert.ElementsMatch(t, actual, expected)
	})
}

func TestGetValuesOfMap(t *testing.T) {
	t.Run("test with string to int map", func(t *testing.T) {
		a := map[string]int{"one": 1, "two": 2, "three": 3}
		expected := []int{1, 2, 3}
		actual := GetValuesOfMap(a)

		assert.ElementsMatch(t, actual, expected)
	})
}

func TestSortStrings(t *testing.T) {
	t.Run("test with unsorted strings", func(t *testing.T) {
		unsortedStrSlice := []string{"banana", "apple", "cherry"}
		expected := []string{"apple", "banana", "cherry"}

		actual := SortStrings(unsortedStrSlice)

		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("test with sorted strings", func(t *testing.T) {
		sortedStrSlice := []string{"apple", "banana", "cherry"}
		expected := []string{"apple", "banana", "cherry"}

		actual := SortStrings(sortedStrSlice)

		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("test with empty string slice", func(t *testing.T) {
		emptyStrSlice := []string{}
		expected := []string{}

		actual := SortStrings(emptyStrSlice)

		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("test with one element string slice", func(t *testing.T) {
		oneElementStrSlice := []string{"apple"}
		expected := []string{"apple"}

		actual := SortStrings(oneElementStrSlice)

		assert.ElementsMatch(t, expected, actual)
	})
}
