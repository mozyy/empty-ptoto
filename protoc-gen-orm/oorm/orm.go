package oorm

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type fieldSort []*descriptor.FieldDescriptorProto

// Len is the number of elements in the collection.
func (f fieldSort) Len() int {
	return len(f)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (f fieldSort) Less(i int, j int) bool {
	sortStrs := []string{"CreatedAt", "UpdatedAt", "DeletedAt"}
	ii := generator.CamelCase(f[i].GetName())
	jj := generator.CamelCase(f[j].GetName())
	if ii == "ID" {
		return true
	}
	if jj == "ID" {
		return false
	}
	return SliceIndex(sortStrs, ii) < SliceIndex(sortStrs, jj)
}

// Swap swaps the elements with indexes i and j.
func (f fieldSort) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

func SliceIndex(strs []string, str string) int {
	for i, s := range strs {
		if s == str {
			return i
		}
	}
	return -1
}
