package tablify

import (
	"testing"
)

func TestCreateLengthStruct(t *testing.T) {
	s := []struct {
		Name   string
		Height int
	}{
		{Name: "Thomas", Height: 72},
		{Name: "Lucy", Height: 65},
	}

	// An []slice of something doesn't satisfy a slice of interfaces,
	// so we need to explicity create one here before passing to lengths.
	g := make([]interface{}, len(s))
	for i, v := range s {
		g[i] = v
	}

	l := lengths(g)

	nameLen := l.FieldByName("Name").Int()
	if nameLen != 8 {
		t.Errorf("Expected `Name` to have length of 8. Got %d", nameLen)
	}

	heightLen := l.FieldByName("Height").Int()
	if heightLen != 8 {
		t.Errorf("Expected `Height` to have length of 8. Got %d", heightLen)
	}
}
