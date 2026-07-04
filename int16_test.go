package opt

import (
	"encoding/json"
	"testing"
)

func TestInt16From(t *testing.T) {
	i := Int16From(42)
	if i.V != 42 || !i.Valid {
		t.Errorf("Int16From(42): got {%v, %v}", i.V, i.Valid)
	}
}

func TestInt16MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Int16
		want string
	}{
		{"positive", Int16From(42), "42"},
		{"null", NewInt16(0, false), "null"},
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

func TestInt16UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int16
		valid   bool
		wantErr bool
	}{
		{"number", "42", 42, true, false},
		{"string number", `"42"`, 42, true, false},
		{"null", "null", 0, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Int16
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

func TestInt16Equal(t *testing.T) {
	if !Int16From(42).Equal(Int16From(42)) {
		t.Error("equal values should be Equal")
	}
}
