package fcnv

import (
	"bytes"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

const (
	ok = "[OK]"
	ng = "[NG]"

	s = "100"
)

var (
	ssInt8 = []string{
		"128",  // NG
		"127",  // OK 1
		"27",   // OK 2
		"7",    // OK 3
		"0",    // OK 4
		"-7",   // OK 5
		"-27",  // OK 6
		"-127", // OK 7
		"-128", // OK 8
		"-129", // NG
	}
	ssInt16 = []string{
		"32768",  // NG
		"32767",  // OK 1
		"2767",   // OK 2
		"767",    // OK 3
		"67",     // OK 4
		"7",      // OK 5
		"0",      // OK 6
		"-7",     // OK 7
		"-67",    // OK 8
		"-767",   // OK 9
		"-2767",  // OK 10
		"-32767", // OK 11
		"-32768", // OK 12
		"-32769", // NG
	}
	ssInt32 = []string{
		"2147483648",  // NG
		"2147483647",  // OK 1
		"147483647",   // OK 2
		"47483647",    // OK 3
		"7483647",     // OK 4
		"483647",      // OK 5
		"83647",       // OK 6
		"3647",        // OK 7
		"647",         // OK 8
		"47",          // OK 9
		"7",           // OK 10
		"0",           // OK 11
		"-7",          // OK 12
		"-47",         // OK 13
		"-647",        // OK 14
		"-3647",       // OK 15
		"-83647",      // OK 16
		"-483647",     // OK 17
		"-7483647",    // OK 18
		"-47483647",   // OK 19
		"-147483647",  // OK 20
		"-2147483647", // OK 21
		"-2147483648", // OK 22
		"-2147483649", // NG
	}
	ssInt64 = []string{
		"9223372036854775808",  // NG
		"9223372036854775807",  // OK 1
		"223372036854775807",   // OK 2
		"23372036854775807",    // OK 3
		"3372036854775807",     // OK 4
		"372036854775807",      // OK 5
		"72036854775807",       // OK 6
		"2036854775807",        // OK 7
		"036854775807",         // OK 8
		"36854775807",          // OK 9
		"6854775807",           // OK 10
		"854775807",            // OK 11
		"54775807",             // OK 12
		"4775807",              // OK 13
		"775807",               // OK 14
		"75807",                // OK 15
		"5807",                 // OK 16
		"807",                  // OK 17
		"07",                   // OK 18
		"7",                    // OK 19
		"0",                    // OK 20
		"-7",                   // OK 21
		"-07",                  // OK 22
		"-807",                 // OK 23
		"-5807",                // OK 24
		"-75807",               // OK 25
		"-775807",              // OK 26
		"-4775807",             // OK 27
		"-54775807",            // OK 28
		"-854775807",           // OK 29
		"-6854775807",          // OK 30
		"-36854775807",         // OK 31
		"-036854775807",        // OK 32
		"-2036854775807",       // OK 33
		"-72036854775807",      // OK 34
		"-372036854775807",     // OK 35
		"-3372036854775807",    // OK 36
		"-23372036854775807",   // OK 37
		"-223372036854775807",  // OK 38
		"-9223372036854775807", // OK 39
		"-9223372036854775808", // OK 40
		"-9223372036854775809", // NG
	}
	ssUint8 = []string{
		"256", // NG
		"255", // OK 1
		"55",  // OK 2
		"5",   // OK 3
		"0",   // OK 4
		"-2",  // NG
	}
	ssUint16 = []string{
		"65536", // NG
		"65535", // OK 1
		"5535",  // OK 2
		"535",   // OK 3
		"35",    // OK 4
		"5",     // OK 5
		"0",     // OK 6
		"-6",    // NG
	}
	ssUint32 = []string{
		"4294967296", // NG
		"4294967295", // OK 1
		"294967295",  // OK 2
		"94967295",   // OK 3
		"4967295",    // OK 4
		"967295",     // OK 5
		"67295",      // OK 6
		"7295",       // OK 7
		"295",        // OK 8
		"95",         // OK 9
		"5",          // OK 10
		"0",          // OK 11
		"-4",         // NG
	}
	ssUint64 = []string{
		"18446744073709551616", // NG
		"18446744073709551615", // OK 1
		"8446744073709551615",  // OK 2
		"46744073709551615",    // OK 3
		"6744073709551615",     // OK 4
		"744073709551615",      // OK 5
		"44073709551615",       // OK 6
		"4073709551615",        // OK 7
		"073709551615",         // OK 8
		"73709551615",          // OK 9
		"3709551615",           // OK 10
		"709551615",            // OK 11
		"09551615",             // OK 12
		"9551615",              // OK 13
		"551615",               // OK 14
		"51615",                // OK 15
		"1615",                 // OK 16
		"615",                  // OK 17
		"15",                   // OK 18
		"5",                    // OK 19
		"0",                    // OK 20
		"-1",                   // NG
	}

	ssInt8Size   = len(ssInt8)
	ssInt16Size  = len(ssInt16)
	ssInt32Size  = len(ssInt32)
	ssInt64Size  = len(ssInt64)
	ssUint8Size  = len(ssUint8)
	ssUint16Size = len(ssUint16)
	ssUint32Size = len(ssUint32)
	ssUint64Size = len(ssUint64)
)

func TestSign(t *testing.T) {
	var s int
	var i string
	o := new(string)

	i = "+1234567890"
	s = sign(i, o)
	if s > 0 && *o == "1234567890" {
		t.Log(ok, i, "->", *o)
	} else {
		t.Error(ng, i, "->", *o)
	}

	i = "1234567890"
	s = sign(i, o)
	if s > 0 && *o == "1234567890" {
		t.Log(ok, i, "->", *o)
	} else {
		t.Error(ng, i, "->", *o)
	}

	i = "-1234567890"
	s = sign(i, o)
	if s < 0 && *o == "1234567890" {
		t.Log(ok, i, "->", *o)
	} else {
		t.Error(ng, i, "->", *o)
	}
}

func TestIsdigit(t *testing.T) {
	for i := byte(0x21); i < 0x7f; i++ {
		if 0x30 <= i && i <= 0x39 {
			if isdigit(i) != true {
				t.Errorf("%s (0x%x, %08b) is digit", string(i), i, i)
			}
		} else {
			if isdigit(i) != false {
				t.Errorf("%s (0x%x, %08b) is not digit", string(i), i, i)
			}
		}
	}
}

func TestAtoi(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssInt64 {
		v, err := Atoi(ssInt64[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40:
			if reflect.TypeOf(v).Kind() == reflect.Int {
				t.Log(ok, "Atoi(", ssInt64[i], ") => ", v)
			} else {
				t.Error(ng, "Atoi")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoi(", ssInt64[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoi(", ssInt64[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoi(ssInt64[i%ssInt64Size])
	}
}

func BenchmarkAtoiByParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64, _ := strconv.ParseInt(ssInt64[i%ssInt64Size], 10, 0)
		_ = int(i64)
	}
}

func TestAtoi8(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssInt8 {
		v, err := Atoi8(ssInt8[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8:
			if reflect.TypeOf(v).Kind() == reflect.Int8 {
				t.Log(ok, "Atoi8(", ssInt8[i], ") => ", v)
			} else {
				t.Error(ng, "Atoi8")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoi8(", ssInt8[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoi8(", ssInt8[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoi8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoi8(ssInt8[i%ssInt8Size])
	}
}

func BenchmarkAtoi8ByParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64, _ := strconv.ParseInt(ssInt8[i%ssInt8Size], 10, 8)
		_ = int8(i64)
	}
}

func TestAtoi16(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssInt16 {
		v, err := Atoi16(ssInt16[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
			if reflect.TypeOf(v).Kind() == reflect.Int16 {
				t.Log(ok, "Atoi16(", ssInt16[i], ") => ", v)
			} else {
				t.Error(ng, "Atoi16")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoi16(", ssInt16[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoi16(", ssInt16[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoi16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoi16(ssInt16[i%ssInt16Size])
	}
}

func BenchmarkAtoi16ByParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64, _ := strconv.ParseInt(ssInt16[i%ssInt16Size], 10, 16)
		_ = int16(i64)
	}
}

func TestAtoi32(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssInt32 {
		v, err := Atoi32(ssInt32[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22:
			if reflect.TypeOf(v).Kind() == reflect.Int32 {
				t.Log(ok, "Atoi32(", ssInt32[i], ") => ", v)
			} else {
				t.Error(ng, "Atoi32")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoi32(", ssInt32[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoi32(", ssInt32[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoi32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoi32(ssInt32[i%ssInt32Size])
	}
}

func BenchmarkAtoi32ByParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64, _ := strconv.ParseInt(ssInt32[i%ssInt32Size], 10, 32)
		_ = int32(i64)
	}
}

func TestAtoi64(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssInt64 {
		v, err := Atoi64(ssInt64[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40:
			if reflect.TypeOf(v).Kind() == reflect.Int64 {
				t.Log(ok, "Atoi64(", ssInt64[i], ") => ", v)
			} else {
				t.Error(ng, "Atoi64")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoi64(", ssInt64[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoi64(", ssInt64[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoi64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoi64(ssInt64[i%ssInt64Size])
	}
}

func BenchmarkAtoi64ByParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64, _ := strconv.ParseInt(ssInt64[i%ssInt64Size], 10, 64)
		_ = int64(i64)
	}
}

func TestAtoui(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssUint64 {
		v, err := Atoui(ssUint64[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
			if reflect.TypeOf(v).Kind() == reflect.Uint {
				t.Log(ok, "Atoui(", ssUint64[i], ") => ", v)
			} else {
				t.Error(ng, "Atoui")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoui(", ssUint64[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoui(", ssUint64[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoui(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoui(ssUint64[i%ssUint8Size])
	}
}

func BenchmarkAtouiByParseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64, _ := strconv.ParseUint(ssUint64[i%ssUint8Size], 10, 0)
		_ = uint(ui64)
	}
}

func TestAtoui8(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssUint8 {
		v, err := Atoui8(ssUint8[i])
		switch i {
		case 1, 2, 3, 4:
			if reflect.TypeOf(v).Kind() == reflect.Uint8 {
				t.Log(ok, "Atoui8(", ssUint8[i], ") => ", v)
			} else {
				t.Error(ng, "Atoui8")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoui8(", ssUint8[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoui8(", ssUint8[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoui8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoui8(ssUint8[i%ssUint8Size])
	}
}

func BenchmarkAtoui8ByParseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64, _ := strconv.ParseUint(ssUint8[i%ssUint8Size], 10, 8)
		_ = uint8(ui64)
	}
}

func TestAtoui16(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssUint16 {
		v, err := Atoui16(ssUint16[i])
		switch i {
		case 1, 2, 3, 4, 5, 6:
			if reflect.TypeOf(v).Kind() == reflect.Uint16 {
				t.Log(ok, "Atoui16(", ssUint16[i], ") => ", v)
			} else {
				t.Error(ng, "Atoui16")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoui16(", ssUint16[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoui16(", ssUint16[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoui16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoui16(ssUint16[i%ssUint16Size])
	}
}

func BenchmarkAtoui16ByParseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64, _ := strconv.ParseUint(ssUint16[i%ssUint16Size], 10, 16)
		_ = uint16(ui64)
	}
}

func TestAtoui32(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssUint32 {
		v, err := Atoui32(ssUint32[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
			if reflect.TypeOf(v).Kind() == reflect.Uint32 {
				t.Log(ok, "Atoui32(", ssUint32[i], ") => ", v)
			} else {
				t.Error(ng, "Atoui32")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoui32(", ssUint32[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoui32(", ssUint32[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoui32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoui64(ssUint32[i%ssUint32Size])
	}
}

func BenchmarkAtoui32ByParseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui32, _ := strconv.ParseUint(ssUint32[i%ssUint32Size], 10, 32)
		_ = uint32(ui32)
	}
}

func TestAtoui64(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	for i := range ssUint64 {
		v, err := Atoui64(ssUint64[i])
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
			if reflect.TypeOf(v).Kind() == reflect.Uint64 {
				t.Log(ok, "Atoui64(", ssUint64[i], ") => ", v)
			} else {
				t.Error(ng, "Atoui64")
				t.Error("     type :", reflect.TypeOf(v))
			}
		default:
			if err != nil {
				t.Log(ok, "Atoui64(", ssUint64[i], ") => ", err.Error())
			} else {
				t.Error(ng, "Atoui64(", ssUint64[i], ") => ", v)
			}
		}
	}
}

func BenchmarkNewAtoui64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Atoui64(ssUint64[i%ssUint64Size])
	}
}

func BenchmarkAtoui64ByParseUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64, _ := strconv.ParseUint(ssUint64[i%ssUint64Size], 10, 64)
		_ = uint64(ui64)
	}
}

func TestAtof32(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	f, _ := Atof32(s)
	if reflect.TypeOf(f).Kind() == reflect.Float32 {
		t.Log(ok, "Atof32(", s, ") => ", f)
	} else {
		t.Error(ng, "Atof32")
		t.Error("     type :", reflect.TypeOf(f))
	}
}

func TestAtof64(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	f, _ := Atof64(s)
	if reflect.TypeOf(f).Kind() == reflect.Float64 {
		t.Log(ok, "Atof64(", s, ") => ", f)
	} else {
		t.Error(ng, "Atof64")
		t.Error("     type :", reflect.TypeOf(f))
	}
}

func TestItoa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int(100)
	s := Itoa(i)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Itoa(", i, ") => ", s)
	} else {
		t.Error(ng, "Itoa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestI8toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int8(100)
	s := I8toa(i)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "I8toa(", i, ") => ", s)
	} else {
		t.Error(ng, "I8toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestI16toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int16(100)
	s := I16toa(i)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "I16toa(", i, ") => ", s)
	} else {
		t.Error(ng, "I16toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestI32toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int32(100)
	s := I32toa(i)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "I32toa(", i, ") => ", s)
	} else {
		t.Error(ng, "I32toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestI64toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int64(100)
	s := I64toa(i)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "I64toa(", i, ") => ", s)
	} else {
		t.Error(ng, "I64toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestUitoa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	ui := uint(100)
	s := Uitoa(ui)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Uitoa(", ui, ") => ", s)
	} else {
		t.Error(ng, "Uitoa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestUi8toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	ui := uint8(100)
	s := Ui8toa(ui)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Ui8toa(", ui, ") => ", s)
	} else {
		t.Error(ng, "Ui8toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestUi16toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	ui := uint16(100)
	s := Ui16toa(ui)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Ui16toa(", ui, ") => ", s)
	} else {
		t.Error(ng, "Ui16toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestUi32toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	ui := uint32(100)
	s := Ui32toa(ui)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Ui32toa(", ui, ") => ", s)
	} else {
		t.Error(ng, "Ui32toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestUi64toa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	ui := uint64(100)
	s := Ui64toa(ui)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Ui64toa(", ui, ") => ", s)
	} else {
		t.Error(ng, "Ui64toa")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestByte2Int(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	b := []byte(string(1))
	i, _ := Byte2Int(b)
	if reflect.TypeOf(i).Kind() == reflect.Int {
		t.Log(ok, "Byte2Int(", b, ") => ", i)
	} else {
		t.Error(ng, "Byte2Int")
		t.Error("     type :", reflect.TypeOf(i))
	}
}

func TestInt2Byte(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int(1)
	b := Int2Byte(i)
	bi, _ := Byte2Int(b)
	if Byte2Bool(b) && bi == i {
		t.Log(ok, "Int2Byte(", i, ") => ", b)
	} else {
		t.Error(ng, "Int2Byte")
		t.Error("     type :", reflect.TypeOf(b))
	}
}

func TestAtobAndBtoa(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	rvString := "abcdefghijklmnopqrstyvwvyz"
	rvBytes := []byte(rvString)

	if bytes.Compare(Atob(rvString), rvBytes) != 0 {
		t.Error(ng, "Atob")
	} else {
		t.Log(ok, "Atob(", rvString, ") => ", rvBytes)
	}

	if bytes.Compare(Atob(Btoa(rvBytes)), rvBytes) != 0 {
		t.Error(ng, "Btoa -> Atob")
	} else {
		t.Log(ok, "Atob(Btoa(", rvBytes, ")) => ", rvBytes)
	}

	if Btoa(rvBytes) != rvString {
		t.Error(ng, "Btoa")
	} else {
		t.Log(ok, "Btoa(", rvBytes, ") => ", rvString)
	}

	if Btoa(Atob(rvString)) != rvString {
		t.Error(ng, "Atob -> Btoa")
	} else {
		t.Log(ok, "Btoa(Atob(", rvString, ")) => ", rvString)
	}
}

func TestBool2Byte(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	bo := bool(true)
	by := Bool2Byte(bo)
	i1, _ := Byte2Int(by)
	i2 := Bool2Int(bo)
	if Byte2Bool(by) && i1 == i2 {
		t.Log(ok, "Bool2Byte(", bo, ") => ", by)
	} else {
		t.Error(ng, "Bool2Byte")
		t.Error("     type :", reflect.TypeOf(by))
	}
}

func TestBool2Int(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	b := bool(true)
	i := Bool2Int(b)
	if reflect.TypeOf(i).Kind() == reflect.Int {
		t.Log(ok, "Bool2Int(", b, ") => ", i)
	} else {
		t.Error(ng, "Bool2Int")
		t.Error("     type :", reflect.TypeOf(i))
	}
}

func TestInt2Bool(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := int(1)
	b := Int2Bool(i)
	if reflect.TypeOf(b).Kind() == reflect.Bool {
		t.Log(ok, "Int2Bool(", i, ") => ", b)
	} else {
		t.Error(ng, "Int2Bool")
		t.Error("     type :", reflect.TypeOf(b))
	}
}

func TestBool2Uint(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	b := bool(true)
	i := Bool2Uint(b)
	if reflect.TypeOf(i).Kind() == reflect.Uint {
		t.Log(ok, "Bool2Uint(", b, ") => ", i)
	} else {
		t.Error(ng, "Bool2Uint")
		t.Error("     type :", reflect.TypeOf(i))
	}
}

func TestUint2Bool(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	i := uint(1)
	b := Uint2Bool(i)
	if reflect.TypeOf(b).Kind() == reflect.Bool {
		t.Log(ok, "Uint2Bool(", i, ") => ", b)
	} else {
		t.Error(ng, "Uint2Bool")
		t.Error("     type :", reflect.TypeOf(b))
	}
}

func TestBool2Str(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	b := bool(true)
	s := Bool2Str(b)
	if reflect.TypeOf(s).Kind() == reflect.String {
		t.Log(ok, "Bool2Str(", b, ") => ", s)
	} else {
		t.Error(ng, "Bool2Str")
		t.Error("     type :", reflect.TypeOf(s))
	}
}

func TestStr2Bool(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	s := string("true")
	b, _ := Str2Bool(s)
	if reflect.TypeOf(b).Kind() == reflect.Bool {
		t.Log(ok, "Str2Bool(", s, ") => ", b)
	} else {
		t.Error(ng, "Str2Bool")
		t.Error("     type :", reflect.TypeOf(b))
	}
}

func TestStruct2Json(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	type s struct {
		test string
	}
	json, _ := Struct2Json(s{test: "100"})
	if reflect.TypeOf(json).Kind() == reflect.String {
		t.Log(ok, "Struct2Json(struct{}) => ", json)
	} else {
		t.Error(ng, "Struct2Json")
		t.Error("     type :", reflect.TypeOf(json))
	}
}

func TestDatetime2Date(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	t1 := time.Now().UTC()
	t2 := t1.Add(time.Hour)

	if t1.Day() < t2.Day() {
		t2 = t2.Add(-2 * time.Hour)
	}

	if Datetime2Date(t1) == Datetime2Date(t2) {
		t.Log(ok)
	} else {
		t.Error(ng)
	}
}

func TestHankaku2Zenkaku(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	if Hankaku2Zenkaku("あいうえおカキクケコｳﾞｻﾞｼﾞｽﾞｾﾞｿﾞＡＢＣＤＥABCDEａｂｃｄｅabcde１２３123") ==
		"あいうえおカキクケコヴザジズゼゾＡＢＣＤＥＡＢＣＤＥａｂｃｄｅａｂｃｄｅ１２３１２３" {
		t.Log(ok)
	} else {
		t.Error(ng)
	}
}

func TestZenkaku2Hankaku(t *testing.T) {
	// GOMAXPROCS
	runtime.GOMAXPROCS(1)
	//---
	if Zenkaku2Hankaku("あいうえおガギグゲゴｳﾞｻﾞｼﾞｽﾞｾﾞｿﾞＡＢＣＤＥABCDEａｂｃｄｅabcde１２３123") ==
		"あいうえおｶﾞｷﾞｸﾞｹﾞｺﾞｳﾞｻﾞｼﾞｽﾞｾﾞｿﾞABCDEABCDEabcdeabcde123123" {
		t.Log(ok)
	} else {
		t.Error(ng)
	}
}
