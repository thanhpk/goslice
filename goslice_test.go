package slice_test

import (
	"testing"
	. "github.com/thanhpk/goslice"
)

func TestEquals(t *testing.T) {
	if !Equal([]string{"1", "2"}, []string{"2", "1"}) {
		t.Fatal("should be true")
	}

	if Equal([]string{"1", "2"}, []string{"1"}) {
		t.Fatal("should not equal")
	}
}

func TestSubstract(t *testing.T) {
	s := Substract([]string{"1", "2"}, []string{"1"})
	if len(s) != 1 || s[0] != "2" {
		t.Fatal("wrong substract")
	}

	s = Substract([]string{"1"}, []string{"1", "2"})
	if len(s) != 0 {
		t.Fatal("wrong substract")
	}
}

func TestFind(t *testing.T) {
	i := Find([]string{"1", "2", "3"}, "2")
	if i != 1 {
		t.Fatal("wrong found")
	}

	i := Find([]string{"1", "2", "3"}, "4")
	if i != -1 {
		t.Fatal("wrong found")
	}
}
