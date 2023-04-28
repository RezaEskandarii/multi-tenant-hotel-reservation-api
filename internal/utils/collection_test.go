package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := &GenericCollection[int]{}
	c.Add(1)
	if c.Size() != 1 {
		t.Errorf("Expected size to be 1, but got %d", c.Size())
	}
}

func TestRemove(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	removed := c.Remove(2)
	if !removed || c.Size() != 2 {
		t.Errorf("Expected element to be removed and size to be 2, but got size %d", c.Size())
	}
}

func TestContains(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	if !c.Contains(1) {
		t.Error("Expected collection to contain 1")
	}
}

func TestSize(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	if c.Size() != 3 {
		t.Errorf("Expected size to be 3, but got %d", c.Size())
	}
}

func TestGet(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	val, _ := c.Get(1)
	if val != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %d", val)
	}
}

func TestWhere(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	filtered := c.Where(func(v int) bool { return v > 1 })
	if filtered.Size() != 2 {
		t.Errorf("Expected size of filtered collection to be 2, but got %d", filtered.Size())
	}
}

func TestFilter(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	filtered := c.Filter(func(v int) bool { return v > 1 })
	if filtered.Size() != 2 {
		t.Errorf("Expected size of filtered collection to be 2, but got %d", filtered.Size())
	}
}

func TestMap(t *testing.T) {
	c := &GenericCollection[int]{elements: []int{1, 2, 3}}
	mapped := c.Map(func(v int) int { return v * 2 })
	expected := []int{2, 4, 6}
	for i, val := range mapped.(*GenericCollection[int]).elements {
		if val != expected[i] {
			t.Errorf("Expected mapped value at index %d to be %d, but got %d", i, expected[i], val)
		}
	}
}
