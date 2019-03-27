package fcnv

import (
	"strconv"
)

func sign(i string, o *string) (s int) {
	s = 1
	if len(i) > 0 {
		switch i[0] {
		case '-':
			s = -1
			fallthrough
		case '+':
			*o = i[1:]
		default:
			*o = i
		}
	}
	return
}

func isdigit(b byte) bool {
	return b^0x30 < 0xA
}

// atoi8 returns the result of ParseInt(s, 10, 8) converted to type int8.
func atoi8(s string) (ret int8, err error) {
	const (
		fnName = "Atoi8"
		cutoff = 1 << (8 - 1)
		limit  = 4 // int8 = [-128, 127]
	)
	s0 := s
	mp := sign(s, &s)
	var buf uint
	for i := range s {
		if i >= limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf < cutoff || mp < 0 && buf <= cutoff {
		return int8(buf) * int8(mp), nil
	} else {
		goto syntaxError
	}

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoi16 returns the result of ParseInt(s, 10, 16) converted to type int16.
func atoi16(s string) (ret int16, err error) {
	const (
		fnName = "Atoi16"
		cutoff = 1 << (16 - 1)
		limit  = 5 // int16 = [-32768, 32767]
	)
	s0 := s
	mp := sign(s, &s)
	var buf uint
	for i := range s {
		if i > limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf < cutoff || mp < 0 && buf <= cutoff {
		return int16(buf) * int16(mp), nil
	} else {
		goto syntaxError
	}

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoi32 returns the result of ParseInt(s, 10, 32) converted to type int32.
func atoi32(s string) (ret int32, err error) {
	const (
		fnName = "Atoi32"
		cutoff = 1 << (32 - 1)
		limit  = 10 // int32 = [-2147483648, 2147483647]
	)
	s0 := s
	mp := sign(s, &s)
	var buf uint
	for i := range s {
		if i > limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf < cutoff || mp < 0 && buf <= cutoff {
		return int32(buf) * int32(mp), nil
	} else {
		goto syntaxError
	}

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoi64 returns the result of ParseInt(s, 10, 64) converted to type int64.
func atoi64(s string) (ret int64, err error) {
	const (
		fnName = "Atoi64"
		cutoff = 1 << (64 - 1)
		limit  = 19 // int64 = [-9223372036854775808, 9223372036854775807]
	)
	s0 := s
	mp := sign(s, &s)
	var buf uint
	for i := range s {
		if i > limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf < cutoff || mp < 0 && buf <= cutoff {
		return int64(buf) * int64(mp), nil
	} else {
		goto syntaxError
	}

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoui returns the result of ParseUint(s, 10, 0) converted to type uint.
func atoui(s string) (ret uint, err error) {
	const (
		fnName = "Atoui"

		// uint32 = [0, 4294967295]
		// uint64 = [0, 18446744073709551615]
		limit = 10 * strconv.IntSize / 32
	)

	// Fast path for small integers that fit int type.
	if sLen := len(s); 0 < sLen && sLen < limit {
		s0 := s
		if sign(s, &s) < 0 {
			return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
		}
		for _, ch := range []byte(s) {
			ch ^= '0'
			if ch > 9 {
				return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
			}
			ret = ret*10 + uint(ch)
		}
		return ret, nil
	}

	// Slow path for invalid or big integers.
	ui64, err := strconv.ParseUint(s, 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnName
		ui64 = 0
	}
	return uint(ui64), err
}

// atoui8 returns the result of ParseUint(s, 10, 8) converted to type uint8.
func atoui8(s string) (ret uint8, err error) {
	const (
		fnName = "Atoui8"
		cutoff = 1<<8 - 1
		limit  = 3 // uint8 = [0, 255]
	)
	s0 := s
	if sign(s, &s) < 0 {
		return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
	}
	buf := uint(0)
	for i := range s {
		if i >= limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf > cutoff {
		goto syntaxError
	}
	return uint8(buf), nil

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoui16 returns the result of ParseUint(s, 10, 16) converted to type uint16.
func atoui16(s string) (ret uint16, err error) {
	const (
		fnName = "Atoui16"
		cutoff = 1<<16 - 1
		limit  = 5 // uint16 = [0, 65535]
	)
	s0 := s
	if sign(s, &s) < 0 {
		return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
	}
	buf := uint(0)
	for i := range s {
		if i >= limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf > cutoff {
		goto syntaxError
	}
	return uint16(buf), nil

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoui32 returns the result of ParseUint(s, 10, 32) converted to type uint32.
func atoui32(s string) (ret uint32, err error) {
	const (
		fnName = "Atoui32"
		cutoff = uint64(1<<32 - 1)
		limit  = 10 // uint32 = [0, 4294967295]
	)
	s0 := s
	if sign(s, &s) < 0 {
		return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
	}
	buf := uint64(0)
	for i := range s {
		if i >= limit {
			goto syntaxError
		}
		if isdigit(s[i]) {
			buf = buf*10 + uint64(s[i]^0x30)
		} else {
			goto syntaxError
		}
	}
	if buf > cutoff {
		goto syntaxError
	}
	return uint32(buf), nil

syntaxError:
	return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
}

// atoui64 returns the result of ParseUint(s, 10, 64) converted to type uint64.
func atoui64(s string) (ret uint64, err error) {
	const (
		fnName = "Atoui64"
		limit  = 20 // uint64 = [0, 18446744073709551615]
	)

	// Fast path for small integers that fit int type.
	if sLen := len(s); 0 < sLen && sLen < limit {
		s0 := s
		if sign(s, &s) < 0 {
			return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
		}
		for _, ch := range []byte(s) {
			ch ^= '0'
			if ch > 9 {
				return 0, &strconv.NumError{fnName, s0, strconv.ErrSyntax}
			}
			ret = ret*10 + uint64(ch)
		}
		return ret, nil
	}

	// Slow path for invalid or big integers.
	ui64, err := strconv.ParseUint(s, 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnName
		ui64 = 0
	}
	return ui64, err
}
