package variables

import (
	"testing"
)

func TestPrimitive(t *testing.T) {
	expected := "astra"
	val := UserName

	if val != expected {
		t.Fatalf("Expected %s but got %s\n", expected, val)
	}
}