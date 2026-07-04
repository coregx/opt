package opt

import (
	"encoding/json"
	"testing"
)

func TestBoolFrom(t *testing.T) {
	b := BoolFrom(true)
	if !b.V || !b.Valid {
		t.Error("BoolFrom(true) should be valid true")
	}

	f := BoolFrom(false)
	if f.V || !f.Valid {
		t.Error("BoolFrom(false) should be valid false")
	}
}

func TestBoolFromPtr(t *testing.T) {
	val := true
	b := BoolFromPtr(&val)
	if !b.V || !b.Valid {
		t.Error("BoolFromPtr(&true) should be valid true")
	}

	null := BoolFromPtr(nil)
	if null.Valid {
		t.Error("BoolFromPtr(nil) should be invalid")
	}
}

func TestBoolMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Bool
		want string
	}{
		{"true", BoolFrom(true), "true"},
		{"false", BoolFrom(false), "false"},
		{"null", NewBool(false, false), "null"},
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

func TestBoolUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		valid   bool
		wantErr bool
	}{
		{"true", "true", true, true, false},
		{"false", "false", false, true, false},
		{"null", "null", false, false, false},
		{"number", "42", false, false, true},
		{"string", `"true"`, false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			err := json.Unmarshal([]byte(tt.input), &b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if b.Valid != tt.valid || (b.Valid && b.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", b.V, b.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestBoolMarshalText(t *testing.T) {
	tests := []struct {
		name string
		val  Bool
		want string
	}{
		{"true", BoolFrom(true), "true"},
		{"false", BoolFrom(false), "false"},
		{"null", NewBool(false, false), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.val.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != tt.want {
				t.Errorf("got %s, want %s", data, tt.want)
			}
		})
	}
}

func TestBoolUnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		valid   bool
		wantErr bool
	}{
		{"true", "true", true, true, false},
		{"false", "false", false, true, false},
		{"null", "null", false, false, false},
		{"empty", "", false, false, false},
		{"invalid", "yes", false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			err := b.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if b.Valid != tt.valid || (b.Valid && b.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", b.V, b.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestBoolEqual(t *testing.T) {
	if !BoolFrom(true).Equal(BoolFrom(true)) {
		t.Error("equal bools should be Equal")
	}
	if BoolFrom(true).Equal(BoolFrom(false)) {
		t.Error("different bools should not be Equal")
	}
	if !NewBool(false, false).Equal(NewBool(false, false)) {
		t.Error("two null bools should be Equal")
	}
}

func TestBoolJSONRoundtrip(t *testing.T) {
	type flags struct {
		Active  Bool `json:"active"`
		Deleted Bool `json:"deleted"`
	}

	original := flags{Active: BoolFrom(true), Deleted: NewBool(false, false)}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"active":true,"deleted":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded flags
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Active.Equal(original.Active) || !decoded.Deleted.Equal(original.Deleted) {
		t.Error("roundtrip failed")
	}
}
