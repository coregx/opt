package opt

import (
	"encoding/json"
	"testing"
)

func TestStringFrom(t *testing.T) {
	s := StringFrom("hello")
	if s.V != "hello" || !s.Valid {
		t.Errorf("StringFrom: got {%v, %v}, want {\"hello\", true}", s.V, s.Valid)
	}
}

func TestStringFromPtr(t *testing.T) {
	str := "hello"
	s := StringFromPtr(&str)
	if s.V != "hello" || !s.Valid {
		t.Errorf("StringFromPtr: got {%v, %v}, want {\"hello\", true}", s.V, s.Valid)
	}

	null := StringFromPtr(nil)
	if null.Valid {
		t.Error("StringFromPtr(nil) should be invalid")
	}
}

func TestStringMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  String
		want string
	}{
		{"valid", StringFrom("hello"), `"hello"`},
		{"empty", StringFrom(""), `""`},
		{"null", NewString("", false), "null"},
		{"special chars", StringFrom(`"quotes" & <tags>`), `"\"quotes\" \u0026 \u003ctags\u003e"`},
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

func TestStringUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    String
		wantErr bool
	}{
		{"string", `"hello"`, StringFrom("hello"), false},
		{"empty", `""`, StringFrom(""), false},
		{"null", "null", NewString("", false), false},
		{"number", "42", String{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s String
			err := json.Unmarshal([]byte(tt.input), &s)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if s.Valid != tt.want.Valid || (s.Valid && s.V != tt.want.V) {
					t.Errorf("got {%v, %v}, want {%v, %v}", s.V, s.Valid, tt.want.V, tt.want.Valid)
				}
			}
		})
	}
}

func TestStringMarshalText(t *testing.T) {
	s := StringFrom("hello")
	data, err := s.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Errorf("got %s, want hello", data)
	}

	null := NewString("", false)
	data, err = null.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 0 {
		t.Errorf("null MarshalText should return empty, got %s", data)
	}
}

func TestStringUnmarshalText(t *testing.T) {
	var s String
	if err := s.UnmarshalText([]byte("hello")); err != nil {
		t.Fatal(err)
	}
	if s.V != "hello" || !s.Valid {
		t.Errorf("got {%v, %v}, want {\"hello\", true}", s.V, s.Valid)
	}

	var empty String
	if err := empty.UnmarshalText([]byte("")); err != nil {
		t.Fatal(err)
	}
	if empty.Valid {
		t.Error("empty text should produce invalid String")
	}
}

func TestStringEqual(t *testing.T) {
	a := StringFrom("hello")
	b := StringFrom("hello")
	if !a.Equal(b) {
		t.Error("equal strings should be Equal")
	}

	c := StringFrom("world")
	if a.Equal(c) {
		t.Error("different strings should not be Equal")
	}

	null1 := NewString("", false)
	null2 := NewString("", false)
	if !null1.Equal(null2) {
		t.Error("two null strings should be Equal")
	}

	if a.Equal(null1) {
		t.Error("valid and null should not be Equal")
	}
}

func TestStringOr(t *testing.T) {
	s := StringFrom("hello")
	if s.Or("default") != "hello" {
		t.Error("valid Or should return value")
	}

	null := NewString("", false)
	if null.Or("default") != "default" {
		t.Error("null Or should return fallback")
	}
}

func TestStringIsZero(t *testing.T) {
	s := StringFrom("hello")
	if s.IsZero() {
		t.Error("valid string should not be IsZero")
	}

	empty := StringFrom("")
	if empty.IsZero() {
		t.Error("valid empty string should not be IsZero")
	}

	null := NewString("", false)
	if !null.IsZero() {
		t.Error("null string should be IsZero")
	}
}

func TestStringJSONRoundtrip(t *testing.T) {
	type user struct {
		Name  String `json:"name"`
		Email String `json:"email"`
	}

	original := user{
		Name:  StringFrom("John"),
		Email: NewString("", false),
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"name":"John","email":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded user
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Name.Equal(original.Name) || !decoded.Email.Equal(original.Email) {
		t.Error("roundtrip failed")
	}
}
