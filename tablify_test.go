package tablify

import (
	"testing"
)

var s = []struct {
	Name   string
	Height int
}{
	{Name: "Thomas", Height: 72},
	{Name: "Lucy", Height: 65},
}

func TestCreateInstance(t *testing.T) {
	tb := New()
	tb.Struct(s)
}

func TestCreateLengthStruct(t *testing.T) {
	vals := interfaceToInterfaceSlice(s)
	l := lengths(vals)

	nameLen := l.FieldByName("Name").Int()
	if nameLen != 8 {
		t.Errorf("Expected `Name` to have length of 8. Got %d", nameLen)
	}

	heightLen := l.FieldByName("Height").Int()
	if heightLen != 8 {
		t.Errorf("Expected `Height` to have length of 8. Got %d", heightLen)
	}
}
