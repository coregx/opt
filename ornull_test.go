package opt

import (
	"testing"
	"time"
)

func TestOrNull_Generic(t *testing.T) {
	v := OrNull(42)
	if v.V != 42 || !v.Valid {
		t.Error("OrNull(42) should be valid")
	}

	zero := OrNull(0)
	if zero.Valid {
		t.Error("OrNull(0) should be null")
	}

	s := OrNull("hello")
	if s.V != "hello" || !s.Valid {
		t.Error("OrNull(\"hello\") should be valid")
	}

	empty := OrNull("")
	if empty.Valid {
		t.Error("OrNull(\"\") should be null")
	}
}

func TestStringOrNull(t *testing.T) {
	s := StringOrNull("hello")
	if s.V != "hello" || !s.Valid {
		t.Errorf("StringOrNull(\"hello\"): got {%v, %v}", s.V, s.Valid)
	}

	null := StringOrNull("")
	if null.Valid {
		t.Error("StringOrNull(\"\") should be null")
	}
}

func TestIntOrNull(t *testing.T) {
	i := IntOrNull(42)
	if i.V != 42 || !i.Valid {
		t.Error("IntOrNull(42) should be valid")
	}

	null := IntOrNull(0)
	if null.Valid {
		t.Error("IntOrNull(0) should be null")
	}

	neg := IntOrNull(-1)
	if !neg.Valid {
		t.Error("IntOrNull(-1) should be valid")
	}
}

func TestInt32OrNull(t *testing.T) {
	i := Int32OrNull(42)
	if i.V != 42 || !i.Valid {
		t.Error("Int32OrNull(42) should be valid")
	}

	null := Int32OrNull(0)
	if null.Valid {
		t.Error("Int32OrNull(0) should be null")
	}
}

func TestInt16OrNull(t *testing.T) {
	i := Int16OrNull(42)
	if i.V != 42 || !i.Valid {
		t.Error("Int16OrNull(42) should be valid")
	}

	null := Int16OrNull(0)
	if null.Valid {
		t.Error("Int16OrNull(0) should be null")
	}
}

func TestFloatOrNull(t *testing.T) {
	f := FloatOrNull(3.14)
	if f.V != 3.14 || !f.Valid {
		t.Error("FloatOrNull(3.14) should be valid")
	}

	null := FloatOrNull(0)
	if null.Valid {
		t.Error("FloatOrNull(0) should be null")
	}
}

func TestBoolOrNull(t *testing.T) {
	b := BoolOrNull(true)
	if !b.V || !b.Valid {
		t.Error("BoolOrNull(true) should be valid true")
	}

	null := BoolOrNull(false)
	if null.Valid {
		t.Error("BoolOrNull(false) should be null")
	}
}

func TestByteOrNull(t *testing.T) {
	b := ByteOrNull(42)
	if b.V != 42 || !b.Valid {
		t.Error("ByteOrNull(42) should be valid")
	}

	null := ByteOrNull(0)
	if null.Valid {
		t.Error("ByteOrNull(0) should be null")
	}
}

func TestTimeOrNull(t *testing.T) {
	now := time.Now()
	tv := TimeOrNull(now)
	if !tv.Valid {
		t.Error("TimeOrNull(now) should be valid")
	}

	null := TimeOrNull(time.Time{})
	if null.Valid {
		t.Error("TimeOrNull(zero) should be null")
	}
}
