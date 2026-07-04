package opt

import (
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	v := From(42)
	result := Map(v, func(i int) string { return strconv.Itoa(i) })
	if result.V != "42" || !result.Valid {
		t.Errorf("Map valid: got {%v, %v}, want {\"42\", true}", result.V, result.Valid)
	}

	null := New(0, false)
	nullResult := Map(null, func(i int) string { return strconv.Itoa(i) })
	if nullResult.Valid {
		t.Error("Map null: should be invalid")
	}
}

func TestFlatMap(t *testing.T) {
	parseInt := func(s string) Value[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return New(0, false)
		}
		return From(n)
	}

	v := From("42")
	result := FlatMap(v, parseInt)
	if result.V != 42 || !result.Valid {
		t.Errorf("FlatMap valid: got {%v, %v}, want {42, true}", result.V, result.Valid)
	}

	bad := From("not_a_number")
	badResult := FlatMap(bad, parseInt)
	if badResult.Valid {
		t.Error("FlatMap bad parse: should be invalid")
	}

	null := New("", false)
	nullResult := FlatMap(null, parseInt)
	if nullResult.Valid {
		t.Error("FlatMap null: should be invalid")
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a, b Value[int]
		want bool
	}{
		{"both valid equal", From(42), From(42), true},
		{"both valid differ", From(42), From(99), false},
		{"both null", New(0, false), New(0, false), true},
		{"valid vs null", From(42), New(0, false), false},
		{"null vs valid", New(0, false), From(42), false},
		{"zero valid equal", From(0), From(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.a, tt.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapChain(t *testing.T) {
	v := From("  Hello, World!  ")
	trimmed := Map(v, strings.TrimSpace)
	upper := Map(trimmed, strings.ToUpper)

	if upper.V != "HELLO, WORLD!" || !upper.Valid {
		t.Errorf("chain: got {%v, %v}, want {\"HELLO, WORLD!\", true}", upper.V, upper.Valid)
	}

	null := New("", false)
	nullUpper := Map(Map(null, strings.TrimSpace), strings.ToUpper)
	if nullUpper.Valid {
		t.Error("chain null: should be invalid")
	}
}
