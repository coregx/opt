package opt

import (
	"encoding/json"
	"testing"
)

func TestInt32From(t *testing.T) {
	i := Int32From(42)
	if i.V != 42 || !i.Valid {
		t.Errorf("Int32From(42): got {%v, %v}", i.V, i.Valid)
	}
}

func TestInt32MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Int32
		want string
	}{
		{"positive", Int32From(42), "42"},
		{"zero", Int32From(0), "0"},
		{"null", NewInt32(0, false), "null"},
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

func TestInt32UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int32
		valid   bool
		wantErr bool
	}{
		{"number", "42", 42, true, false},
		{"string number", `"42"`, 42, true, false},
		{"empty string", `""`, 0, false, false},
		{"null", "null", 0, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Int32
			err := json.Unmarshal([]byte(tt.input), &i)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Valid != tt.valid || (i.Valid && i.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", i.V, i.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestInt32TextRoundtrip(t *testing.T) {
	i := Int32From(42)
	data, err := i.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "42" {
		t.Errorf("got %s, want 42", data)
	}

	var parsed Int32
	if err := parsed.UnmarshalText(data); err != nil {
		t.Fatal(err)
	}
	if !parsed.Equal(i) {
		t.Error("text roundtrip failed")
	}
}

func TestInt32Equal(t *testing.T) {
	if !Int32From(42).Equal(Int32From(42)) {
		t.Error("equal values should be Equal")
	}
	if Int32From(42).Equal(Int32From(99)) {
		t.Error("different values should not be Equal")
	}
}
