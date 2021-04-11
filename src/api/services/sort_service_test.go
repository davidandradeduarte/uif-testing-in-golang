package services

import (
	"testing"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/utils/sort"
)

// white box testing, since we are testing inside the same package
// we have full access to private consts or functions

func TestConstants(t *testing.T) {
	if PublicConst != "public" {
		t.Error("PublicConst should be 'public'")
	}

	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	//elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 7, 0}
	elements := sort.GetElements(10)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}

func TestSortMoreThan1000(t *testing.T) {
	//elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 7, 0}
	elements := sort.GetElements(1001)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 1000 {
		t.Error("last element should be 1000")
	}
}

func BenchmarkBubbleSort1K(b *testing.B) {
	elements := sort.GetElements(1000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort10K(b *testing.B) {
	elements := sort.GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
