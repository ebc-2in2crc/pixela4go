package pixela

import "testing"

func TestString(t *testing.T) {
	s := "Hello Pixela!"

	result := String(s)

	if StringValue(result) != s {
		t.Errorf("got: %v\nwant: %v", StringValue(result), s)
	}
}

func TestStringValue(t *testing.T) {
	a := []*string{String("Hello Pixela!"), nil}

	for _, p := range a {
		result := StringValue(p)
		if p == nil && result != "" {
			t.Errorf("got: %v\nwant: %v", result, "\"\"")
		}
		if p != nil && result != *p {
			t.Errorf("got: %v\nwant: %v", result, *p)
		}
	}
}

func TestBool(t *testing.T) {
	b := true

	result := Bool(b)

	if BoolValue(result) != b {
		t.Errorf("got: %v\nwant: %v", BoolValue(result), b)
	}
}

func TestBoolValue(t *testing.T) {
	a := []*bool{Bool(true), nil}

	for _, p := range a {
		result := BoolValue(p)
		if p == nil && result != false {
			t.Errorf("got: %v\nwant: %v", result, false)
		}
		if p != nil && result != *p {
			t.Errorf("got: %v\nwant: %v", result, *p)
		}
	}
}
