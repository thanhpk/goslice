package slice_test

import (
	"testing"
	. "github.com/thanhpk/goslice"
)

func TestEquals(t *testing.T) {
	if !Equals([]string{"1", "2"}, []string{"2", "1"}) {
		t.Fatal("should be true")
	}

	if Equals([]string{"1", "2"}, []string{"1"}) {
		t.Fatal("should not equal")
	}
}
