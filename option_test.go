package opt

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	v := New(42, true)
	if v.V != 42 || !v.Valid {
		t.Errorf("New(42, true) = {%v, %v}, want {42, true}", v.V, v.Valid)
	}

	null := New(0, false)
	if null.Valid {
		t.Error("New(0, false) should be invalid")
	}
}

func TestFrom(t *testing.T) {
	v := From("hello")
	if v.V != "hello" || !v.Valid {
		t.Errorf("From(\"hello\") = {%v, %v}, want {\"hello\", true}", v.V, v.Valid)
	}
}

func TestFromPtr(t *testing.T) {
	s := "hello"
	v := FromPtr(&s)
	if v.V != "hello" || !v.Valid {
		t.Errorf("FromPtr(&\"hello\") = {%v, %v}, want {\"hello\", true}", v.V, v.Valid)
	}

	null := FromPtr[string](nil)
	if null.Valid {
		t.Error("FromPtr(nil) should be invalid")
	}
}

func TestOr(t *testing.T) {
	v := From(42)
	if v.Or(99) != 42 {
		t.Errorf("From(42).Or(99) = %v, want 42", v.Or(99))
	}

	null := New(0, false)
	if null.Or(99) != 99 {
		t.Errorf("null.Or(99) = %v, want 99", null.Or(99))
	}
}

func TestOrZero(t *testing.T) {
	v := From("hello")
	if v.OrZero() != "hello" {
		t.Errorf("From(\"hello\").OrZero() = %v, want \"hello\"", v.OrZero())
	}

	null := New("", false)
	if null.OrZero() != "" {
		t.Errorf("null.OrZero() = %v, want \"\"", null.OrZero())
	}
}

func TestOrElse(t *testing.T) {
	v := From(42)
	if v.OrElse(func() int { return 99 }) != 42 {
		t.Error("OrElse should return value when valid")
	}

	called := false
	null := New(0, false)
	result := null.OrElse(func() int {
		called = true
		return 99
	})
	if !called || result != 99 {
		t.Error("OrElse should call fn and return its result when invalid")
	}
}

func TestSetValid(t *testing.T) {
	v := New(0, false)
	v.SetValid(42)
	if v.V != 42 || !v.Valid {
		t.Errorf("after SetValid(42): {%v, %v}, want {42, true}", v.V, v.Valid)
	}
}

func TestPtr(t *testing.T) {
	v := From(42)
	ptr := v.Ptr()
	if ptr == nil || *ptr != 42 {
		t.Error("Ptr() should return pointer to value")
	}

	null := New(0, false)
	if null.Ptr() != nil {
		t.Error("null.Ptr() should return nil")
	}
}

func TestIsZero(t *testing.T) {
	v := From(42)
	if v.IsZero() {
		t.Error("valid value should not be zero")
	}

	null := New(0, false)
	if !null.IsZero() {
		t.Error("null value should be zero")
	}

	zero := From(0)
	if zero.IsZero() {
		t.Error("valid zero-value should not be IsZero")
	}
}

func TestOptionMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Option[int]
		want string
	}{
		{"valid", From(42), "42"},
		{"zero", From(0), "0"},
		{"null", New(0, false), "null"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.val)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != tt.want {
				t.Errorf("got %s, want %s", data, tt.want)
			}
		})
	}
}

func TestOptionUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Option[int]
	}{
		{"number", "42", From(42)},
		{"zero", "0", From(0)},
		{"null", "null", New(0, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v Option[int]
			if err := json.Unmarshal([]byte(tt.input), &v); err != nil {
				t.Fatal(err)
			}
			if v.Valid != tt.want.Valid || (v.Valid && v.V != tt.want.V) {
				t.Errorf("got {%v, %v}, want {%v, %v}", v.V, v.Valid, tt.want.V, tt.want.Valid)
			}
		})
	}
}

func TestOptionJSONRoundtrip(t *testing.T) {
	type payload struct {
		Name  Option[string] `json:"name"`
		Count Option[int]    `json:"count"`
	}

	original := payload{
		Name:  From("test"),
		Count: New(0, false),
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"name":"test","count":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded payload
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Name.V != "test" || !decoded.Name.Valid {
		t.Errorf("Name: got {%v, %v}, want {test, true}", decoded.Name.V, decoded.Name.Valid)
	}
	if decoded.Count.Valid {
		t.Error("Count should be invalid after roundtrip")
	}
}
