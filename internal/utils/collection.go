package utils

import (
	"reflect"
)

// Collection is a generic interface for a collection of elements.
// It provides methods for adding, removing, and filtering elements,
// as well as checking if an element is in the collection and retrieving
// elements by index.
type Collection[T any] interface {
	Add(T)
	Remove(T) bool
	Contains(T) bool
	Size() int
	Get(int) (T, error)
	Where(func(T) bool) Collection[T]
	Filter(func(T) bool) Collection[T]
	Map(func(T) T) Collection[T]
}

// GenericCollection is a generic implementation of the Collection interface.
type GenericCollection[T any] struct {
	elements []T
}

// Add appends an element to the collection.
func (c *GenericCollection[T]) Add(e T) {
	c.elements = append(c.elements, e)
}

// Remove removes an element from the collection if it exists and returns true.
// If the element is not in the collection, it returns false.
func (c *GenericCollection[T]) Remove(e T) bool {
	for i, v := range c.elements {
		if isEqual(v, e) {
			c.elements = append(c.elements[:i], c.elements[i+1:]...)
			return true
		}
	}
	return false
}

// Contains checks if an element is in the collection.
func (c *GenericCollection[T]) Contains(e T) bool {

	for _, v := range c.elements {
		if isEqual(v, e) {
			return true
		}
	}
	return false
}

// Size returns the number of elements in the collection.
func (c *GenericCollection[T]) Size() int {
	return len(c.elements)
}

// Get retrieves the element at the specified index.
func (c *GenericCollection[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(c.elements) {
		panic("index out of range")
	}
	return c.elements[index], nil
}

// Where filters the collection based on a provided predicate function,
// returning a new collection containing the matching elements.
func (c *GenericCollection[T]) Where(f func(T) bool) Collection[T] {
	result := &GenericCollection[T]{}
	for _, v := range c.elements {
		if f(v) {
			result.Add(v)
		}
	}
	return result
}

// Filter is an alias for the Where method.
func (c *GenericCollection[T]) Filter(f func(T) bool) Collection[T] {
	return c.Where(f)
}

// Map applies a provided function to each element in the collection
// and returns a new collection with the transformed elements.
func (c *GenericCollection[T]) Map(f func(T) T) Collection[T] {
	result := &GenericCollection[T]{}
	for _, v := range c.elements {
		result.Add(f(v))
	}
	return result
}

// isEqual is a helper function to compare two values of any type for equality.
func isEqual[T any](a, b T) bool {
	return reflect.ValueOf(a).Interface() == reflect.ValueOf(b).Interface()
}
