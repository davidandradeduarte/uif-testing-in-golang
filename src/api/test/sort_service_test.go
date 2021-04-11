package test

import (
	"testing"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/services"
)

// black box testing, since we are testing outside the package
// no access to private consts or functions
// (prefer white box over black box testing)

func TestConstants(t *testing.T) {
	if services.PublicConst != "public" {
		t.Error("PublicConst should be 'public'")
	}

	// can't test privateConst here
}

func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 7, 0}

	services.Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}
