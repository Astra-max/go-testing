package variables

import (
	"testing"
	"reflect"
)

func TestPrimitiveStr(t *testing.T) {
	expected := "astra"
	val := UserName

	if val != expected {
		t.Fatalf("Expected %s but got %s\n", expected, val)
	}
	
}

func TestPrimitiveInt(t *testing.T) {
	expected := 27
	val := Age

	if reflect.TypeOf(val).Kind() != reflect.Int {
		t.Fatalf("Expected an interger value got %v\n", val)
	}
	if val != expected {
		t.Fatalf("Expected %v but got %v\n", expected, val)
	}
}

func TestPrimitiveBool(t *testing.T) {
	expected := true
	val := IsMarried

	if reflect.TypeOf(val).Kind() != reflect.Bool {
		t.Fatalf("Expected an interger value got %v\n", val)
	}
	if val != expected {
		t.Fatalf("Expected %v but got %v\n", expected, val)
	}
}

func TestPrimitiveByte(t *testing.T) {
	expected := byte(77)
	val := GenderByte

	if reflect.TypeOf(val).Kind() != reflect.Uint8 {
		t.Fatalf("Expected an interger value got %v\n", val)
	}
	if val != expected {
		t.Fatalf("Expected %v but got %v\n", expected, val)
	}
}

func TestPrimitiveRune(t *testing.T) {
	expected := 'M'
	val := Gender

	if reflect.TypeOf(val).Kind() != reflect.Int32 {
		t.Fatalf("Expected an interger value got %v\n", val)
	}
	if val != expected {
		t.Fatalf("Expected %v but got %v\n", expected, val)
	}
}

func TestPrimitiveFloat(t *testing.T) {
	var expected float32 = 2.50
	val := Height

	if reflect.TypeOf(val).Kind() != reflect.Float32 {
		t.Fatalf("Expected an interger value got %v\n", val)
	}
	if val != expected {
		t.Fatalf("Expected %v but got %v\n", expected, val)
	}
}