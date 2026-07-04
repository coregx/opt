// Package internal provides shared marshal/unmarshal helpers for opt and opt/zero packages.
package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// UnmarshalIntJSON parses JSON data as an integer, accepting both numbers and strings.
// Returns (value, valid, error). Empty strings return (0, false, nil).
func UnmarshalIntJSON(data []byte, bits int) (int64, bool, error) {
	if len(data) == 0 {
		return 0, false, fmt.Errorf("opt: empty JSON input")
	}
	switch data[0] {
	case 'n':
		return 0, false, nil
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		if s == "" {
			return 0, false, nil
		}
		n, err := strconv.ParseInt(s, 10, bits)
		if err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		return n, true, nil
	default:
		var n int64
		if err := json.Unmarshal(data, &n); err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		return n, true, nil
	}
}

// UnmarshalIntText parses text as an integer.
// Returns (value, valid, error). Empty and "null" strings return (0, false, nil).
func UnmarshalIntText(text []byte, bits int) (int64, bool, error) {
	str := string(text)
	if str == "" || str == "null" {
		return 0, false, nil
	}
	n, err := strconv.ParseInt(str, 10, bits)
	if err != nil {
		return 0, false, err
	}
	return n, true, nil
}

// UnmarshalFloatJSON parses JSON data as a float64, accepting both numbers and strings.
// Returns (value, valid, error). Empty strings return (0, false, nil).
func UnmarshalFloatJSON(data []byte) (float64, bool, error) {
	if len(data) == 0 {
		return 0, false, fmt.Errorf("opt: empty JSON input")
	}
	switch data[0] {
	case 'n':
		return 0, false, nil
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		if s == "" {
			return 0, false, nil
		}
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		return n, true, nil
	default:
		var n float64
		if err := json.Unmarshal(data, &n); err != nil {
			return 0, false, fmt.Errorf("opt: couldn't unmarshal JSON: %w", err)
		}
		return n, true, nil
	}
}

// UnmarshalFloatText parses text as a float64.
// Returns (value, valid, error). Empty and "null" strings return (0, false, nil).
func UnmarshalFloatText(text []byte) (float64, bool, error) {
	str := string(text)
	if str == "" || str == "null" {
		return 0, false, nil
	}
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, false, err
	}
	return n, true, nil
}
