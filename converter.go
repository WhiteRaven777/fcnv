package fcnv

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/exp/constraints"
	"golang.org/x/text/width"
)

func Atoi(s string) (ret int, err error) {
	return strconv.Atoi(s)
}

func Atoi8(s string) (ret int8, err error) {
	return atoi8(s)
}

func Atoi16(s string) (ret int16, err error) {
	return atoi16(s)
}

func Atoi32(s string) (ret int32, err error) {
	return atoi32(s)
}

func Atoi64(s string) (ret int64, err error) {
	return atoi64(s)
}

func Atoui(s string) (ret uint, err error) {
	return atoui(s)
}

func Atoui8(s string) (ret uint8, err error) {
	return atoui8(s)
}

func Atoui16(s string) (ret uint16, err error) {
	return atoui16(s)
}

func Atoui32(s string) (ret uint32, err error) {
	return atoui32(s)
}

func Atoui64(s string) (ret uint64, err error) {
	return atoui64(s)
}

func Atof32(s string) (ret float32, err error) {
	f, e := strconv.ParseFloat(s, 32)
	return float32(f), e
}

func Atof64(s string) (ret float64, err error) {
	return strconv.ParseFloat(s, 64)
}

func Atoc64(s string) (ret complex64, err error) {
	cm, e := strconv.ParseComplex(s, 64)
	return complex64(cm), e
}

func Atoc128(s string) (ret complex128, err error) {
	return strconv.ParseComplex(s, 128)
}

func Atob(s string) (b []byte) {
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: h.Data,
		Len:  h.Len,
		Cap:  h.Len,
	}))
}

// Itoa converts an int* to a string.
func Itoa[intx constraints.Signed](i intx) (ret string) {
	return strconv.FormatInt(int64(i), 10)
}

// Deprecated: I8toa is discontinued in the future.
// Itoa should be used.
func I8toa(i int8) (ret string) { return Itoa(i) }

// Deprecated: I16toa is discontinued in the future.
// Itoa should be used.
func I16toa(i int16) (ret string) { return Itoa(i) }

// Deprecated: I32toa is discontinued in the future.
// Itoa should be used.
func I32toa(i int32) (ret string) { return Itoa(i) }

// Deprecated: I64toa is discontinued in the future.
// Itoa should be used.
func I64toa(i int64) (ret string) { return Itoa(i) }

// Uitoa converts an uint* to a string.
func Uitoa[uintx constraints.Unsigned](ui uintx) (ret string) {
	return strconv.FormatUint(uint64(ui), 10)
}

// Deprecated: Ui8toa is discontinued in the future.
// Uitoa should be used.
func Ui8toa(ui uint8) (ret string) { return Uitoa(ui) }

// Deprecated: Ui16toa is discontinued in the future.
// Uitoa should be used.
func Ui16toa(ui uint16) (ret string) { return Uitoa(ui) }

// Deprecated: Ui32toa is discontinued in the future.
// Uitoa should be used.
func Ui32toa(ui uint32) (ret string) { return Uitoa(ui) }

// Deprecated: Ui64toa is discontinued in the future.
// Uitoa should be used.
func Ui64toa(ui uint64) (ret string) { return Uitoa(ui) }

// Ctoa converts a Complex (complex64, complex128) to a string.
// The representation precision is complex64 compliant.
func Ctoa[cmx constraints.Complex](f cmx) (ret string) {
	return strconv.FormatComplex(complex128(f), 'f', -1, 128)
}

// Ftoa converts a Float (float32, float64) to a string.
// The representation precision is float64 compliant.
func Ftoa[floatx constraints.Float](f floatx) (ret string) {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

func Btoa(b []byte) (ret string) {
	h := *(*reflect.StringHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: h.Data,
		Len:  h.Len,
	}))
}

func Byte2Int(b []byte) (ret int, err error) {
	return strconv.Atoi(hex.EncodeToString(b))
}

func Int2Byte[intx constraints.Signed](i intx) (ret []byte) {
	return []byte(string(rune(i)))
}

func Byte2Bool(by []byte) (bo bool) {
	i, _ := Byte2Int(by)
	if i == 0 {
		bo = false
	} else {
		bo = true
	}
	return
}

func Bool2Byte(bo bool) (by []byte) {
	if bo {
		by = []byte{0x01}
	} else {
		by = []byte{0x00}
	}
	return
}

func Bool2Int(b bool) (i int) {
	if b {
		i = 1
	} else {
		i = 0
	}
	return
}

func Int2Bool[INT constraints.Integer](i INT) (b bool) {
	if i == 0 {
		b = false
	} else {
		b = true
	}
	return
}

// Deprecated: Bool2Uint is discontinued in the future.
// Bool2Int should be used.
func Bool2Uint(b bool) (i uint) { return uint(Bool2Int(b)) }

// Deprecated: Uint2Bool is discontinued in the future.
// Int2Bool should be used.
func Uint2Bool(i uint) (b bool) { return Int2Bool(i) }

func Bool2Str(b bool) (s string) {
	if b {
		s = "true"
	} else {
		s = "false"
	}
	return
}

func Str2Bool(s string) (ret bool, err error) {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false, errors.New("invalid syntax")
}

func Struct2Json(structure any) (jsonStr string, err error) {
	b, err := json.Marshal(structure)
	jsonStr = Btoa(b)
	return
}

func Datetime2Date(t time.Time) (ret time.Time) {
	return t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
}

func Hankaku2Zenkaku(str string) (ret string) {
	harf2full := strings.NewReplacer(
		"ｳﾞ", "ヴ",
		"ｶﾞ", "ガ",
		"ｷﾞ", "ギ",
		"ｸﾞ", "グ",
		"ｹﾞ", "ゲ",
		"ｺﾞ", "ゴ",
		"ｻﾞ", "ザ",
		"ｼﾞ", "ジ",
		"ｽﾞ", "ズ",
		"ｾﾞ", "ゼ",
		"ｿﾞ", "ゾ",
		"ﾀﾞ", "ダ",
		"ﾁﾞ", "ヂ",
		"ﾂﾞ", "ヅ",
		"ﾃﾞ", "デ",
		"ﾄﾞ", "ド",
		"ﾊﾞ", "バ",
		"ﾋﾞ", "ビ",
		"ﾌﾞ", "ブ",
		"ﾍﾞ", "ベ",
		"ﾎﾞ", "ボ",
		"ﾊﾟ", "パ",
		"ﾋﾟ", "ピ",
		"ﾌﾟ", "プ",
		"ﾍﾟ", "ペ",
		"ﾎﾟ", "ポ",
	)
	return width.Widen.String(harf2full.Replace(str))
}

func Zenkaku2Hankaku(str string) (ret string) {
	full2harf := strings.NewReplacer(
		"ヴ", "ｳﾞ",
		"ガ", "ｶﾞ",
		"ギ", "ｷﾞ",
		"グ", "ｸﾞ",
		"ゲ", "ｹﾞ",
		"ゴ", "ｺﾞ",
		"ザ", "ｻﾞ",
		"ジ", "ｼﾞ",
		"ズ", "ｽﾞ",
		"ゼ", "ｾﾞ",
		"ゾ", "ｿﾞ",
		"ダ", "ﾀﾞ",
		"ヂ", "ﾁﾞ",
		"ヅ", "ﾂﾞ",
		"デ", "ﾃﾞ",
		"ド", "ﾄﾞ",
		"バ", "ﾊﾞ",
		"ビ", "ﾋﾞ",
		"ブ", "ﾌﾞ",
		"ベ", "ﾍﾞ",
		"ボ", "ﾎﾞ",
		"パ", "ﾊﾟ",
		"ピ", "ﾋﾟ",
		"プ", "ﾌﾟ",
		"ペ", "ﾍﾟ",
		"ポ", "ﾎﾟ",
		"ヰ", "ｲ",
		"ヱ", "ｴ",
		"ヮ", "ﾜ",
	)
	return width.Narrow.String(full2harf.Replace(str))
}
