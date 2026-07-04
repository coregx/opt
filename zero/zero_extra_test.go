package zero

import (
	"encoding/json"
	"testing"
	"time"
)

// --- Int32 ---

func TestInt32From(t *testing.T) {
	i := Int32From(42)
	if i.V != 42 || !i.Valid {
		t.Error("Int32From(42) should be valid")
	}
	zero := Int32From(0)
	if zero.Valid {
		t.Error("Int32From(0) should be invalid (zero = null)")
	}
}

func TestInt32MarshalJSON(t *testing.T) {
	i := Int32From(42)
	data, _ := json.Marshal(i)
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}
	null := NewInt32(0, false)
	data, _ = json.Marshal(null)
	if string(data) != "0" {
		t.Errorf("null marshal: got %s, want 0", data)
	}
}

func TestInt32UnmarshalJSON(t *testing.T) {
	var i Int32
	json.Unmarshal([]byte("42"), &i)
	if i.V != 42 || !i.Valid {
		t.Error("should be valid 42")
	}
	var zero Int32
	json.Unmarshal([]byte("0"), &zero)
	if zero.Valid {
		t.Error("0 should be invalid")
	}
	var null Int32
	json.Unmarshal([]byte("null"), &null)
	if null.Valid {
		t.Error("null should be invalid")
	}
}

func TestInt32Text(t *testing.T) {
	i := Int32From(42)
	data, _ := i.MarshalText()
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}
	var parsed Int32
	parsed.UnmarshalText([]byte("42"))
	if !parsed.Equal(i) {
		t.Error("text roundtrip failed")
	}
	var null Int32
	null.UnmarshalText([]byte(""))
	if null.Valid {
		t.Error("empty text should be invalid")
	}
}

// --- Int16 ---

func TestInt16From(t *testing.T) {
	i := Int16From(42)
	if i.V != 42 || !i.Valid {
		t.Error("Int16From(42) should be valid")
	}
	zero := Int16From(0)
	if zero.Valid {
		t.Error("Int16From(0) should be invalid")
	}
}

func TestInt16MarshalJSON(t *testing.T) {
	i := Int16From(42)
	data, _ := json.Marshal(i)
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}
}

func TestInt16UnmarshalJSON(t *testing.T) {
	var i Int16
	json.Unmarshal([]byte("42"), &i)
	if i.V != 42 || !i.Valid {
		t.Error("should be valid 42")
	}
	var zero Int16
	json.Unmarshal([]byte("0"), &zero)
	if zero.Valid {
		t.Error("0 should be invalid")
	}
}

func TestInt16Text(t *testing.T) {
	i := Int16From(42)
	data, _ := i.MarshalText()
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}
	var null Int16
	null.UnmarshalText([]byte("null"))
	if null.Valid {
		t.Error("null text should be invalid")
	}
}

// --- Byte ---

func TestByteFrom(t *testing.T) {
	b := ByteFrom(65) // 'A'
	if b.V != 65 || !b.Valid {
		t.Error("ByteFrom(65) should be valid")
	}
	zero := ByteFrom(0)
	if zero.Valid {
		t.Error("ByteFrom(0) should be invalid")
	}
}

func TestByteMarshalJSON(t *testing.T) {
	b := ByteFrom(42)
	data, _ := json.Marshal(b)
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}
	null := NewByte(0, false)
	data, _ = json.Marshal(null)
	if string(data) != "0" {
		t.Errorf("null marshal: got %s, want 0", data)
	}
}

func TestByteUnmarshalJSON(t *testing.T) {
	var b Byte
	json.Unmarshal([]byte("42"), &b)
	if b.V != 42 || !b.Valid {
		t.Error("should be valid 42")
	}
	var null Byte
	json.Unmarshal([]byte("null"), &null)
	if null.Valid {
		t.Error("null should be invalid")
	}
}

func TestByteText(t *testing.T) {
	b := ByteFrom(65) // 'A'
	data, _ := b.MarshalText()
	if string(data) != "A" {
		t.Errorf("got %s, want A", data)
	}
	var null Byte
	null.UnmarshalText([]byte(""))
	if null.Valid {
		t.Error("empty text should be invalid")
	}
}

// --- String text ---

func TestStringText(t *testing.T) {
	s := StringFrom("hello")
	data, _ := s.MarshalText()
	if string(data) != "hello" {
		t.Errorf("got %s", data)
	}
	var parsed String
	parsed.UnmarshalText([]byte("hello"))
	if parsed.V != "hello" || !parsed.Valid {
		t.Error("text unmarshal failed")
	}
	var empty String
	empty.UnmarshalText([]byte(""))
	if empty.Valid {
		t.Error("empty text should be invalid")
	}
}

// --- Bool text ---

func TestBoolText(t *testing.T) {
	b := BoolFrom(true)
	data, _ := b.MarshalText()
	if string(data) != "true" {
		t.Errorf("got %s", data)
	}
	null := NewBool(false, false)
	data, _ = null.MarshalText()
	if string(data) != "false" {
		t.Errorf("null bool marshal text: got %s, want false", data)
	}
	var parsed Bool
	parsed.UnmarshalText([]byte("true"))
	if !parsed.V || !parsed.Valid {
		t.Error("true text should be valid")
	}
	var f Bool
	f.UnmarshalText([]byte("false"))
	if f.Valid {
		t.Error("false text should be invalid (zero = null)")
	}
}

// --- Float text ---

func TestFloatText(t *testing.T) {
	f := FloatFrom(3.14)
	data, _ := f.MarshalText()
	if string(data) != "3.14" {
		t.Errorf("got %s", data)
	}
	null := NewFloat(0, false)
	data, _ = null.MarshalText()
	if string(data) != "0" {
		t.Errorf("null float text: got %s, want 0", data)
	}
	var parsed Float
	parsed.UnmarshalText([]byte("3.14"))
	if parsed.V != 3.14 || !parsed.Valid {
		t.Error("text unmarshal failed")
	}
	var nullText Float
	nullText.UnmarshalText([]byte(""))
	if nullText.Valid {
		t.Error("empty text should be invalid")
	}
}

// --- Time text ---

func TestTimeText(t *testing.T) {
	null := NewTime(time.Time{}, false)
	data, _ := null.MarshalText()
	if len(data) == 0 {
		t.Error("zero/Time MarshalText should return zero time, not empty")
	}
	var parsed Time
	parsed.UnmarshalText([]byte(""))
	if parsed.Valid {
		t.Error("empty text should be invalid")
	}
	var nullText Time
	nullText.UnmarshalText([]byte("null"))
	if nullText.Valid {
		t.Error("null text should be invalid")
	}
}

// --- Int string JSON ---

func TestIntStringJSON(t *testing.T) {
	var i Int
	json.Unmarshal([]byte(`"42"`), &i)
	if i.V != 42 || !i.Valid {
		t.Error("string 42 should be valid")
	}
}

// --- Float string JSON ---

func TestFloatStringJSON(t *testing.T) {
	var f Float
	json.Unmarshal([]byte(`"3.14"`), &f)
	if f.V != 3.14 || !f.Valid {
		t.Error("string 3.14 should be valid")
	}
	var empty Float
	json.Unmarshal([]byte(`""`), &empty)
	if empty.Valid {
		t.Error("empty string should be invalid")
	}
}

// --- Value generic ---

func TestValueOr(t *testing.T) {
	v := From(42)
	if v.Or(99) != 42 {
		t.Error("valid Or should return value")
	}
	null := New(0, false)
	if null.Or(99) != 99 {
		t.Error("null Or should return fallback")
	}
}

func TestValueSetValid(t *testing.T) {
	v := New(0, false)
	v.SetValid(42)
	if v.V != 42 || !v.Valid {
		t.Error("SetValid should set value and valid")
	}
}

func TestValuePtr(t *testing.T) {
	v := From(42)
	ptr := v.Ptr()
	if ptr == nil || *ptr != 42 {
		t.Error("Ptr should return pointer to value")
	}
	null := New(0, false)
	if null.Ptr() != nil {
		t.Error("null Ptr should be nil")
	}
}
