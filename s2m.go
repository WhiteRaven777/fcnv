package fcnv

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	inputValueTypeError = "only the structure is assumed for input values"
	structureTypeError  = "structure must be a pointer to the structure to deserialize"
)

func serializeTime(t any) (ret string) {
	var v time.Time
	if rv := reflect.ValueOf(t); rv.Kind() == reflect.Pointer {
		v = rv.Elem().Interface().(time.Time)
	} else {
		v = rv.Interface().(time.Time)
	}
	return v.UTC().Format(time.RFC3339Nano)
}

func serialize(rv reflect.Value) (ret string, err error) {
	switch rv.Kind() {
	case reflect.Bool:
		ret = Bool2Str(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = Itoa(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret = Uitoa(rv.Uint())
	case reflect.Float32, reflect.Float64:
		ret = Ftoa(rv.Float())
	case reflect.Complex64, reflect.Complex128:
		ret = Ctoa(rv.Complex())
	case reflect.String:
		ret = rv.String()
	default:
		switch v := rv.Interface().(type) {
		case time.Time, *time.Time:
			ret = serializeTime(v)
		default:
			ret, err = Struct2Json(v)
		}
	}
	return
}

func deserializeTime(v string, k reflect.Kind) (ret reflect.Value, err error) {
	var buf time.Time
	if buf, err = time.Parse(time.RFC3339Nano, v); err == nil {
		buf = buf.UTC()
		ret = reflect.ValueOf(&buf)
		if k != reflect.Pointer {
			ret = ret.Elem()
		}
	}
	return
}

func deserialize(value string, rv reflect.Value) (err error) {
	switch rv.Kind() {
	case reflect.Bool:
		var buf bool
		if buf, err = Str2Bool(value); err == nil {
			rv.SetBool(buf)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var buf int64
		if buf, err = Atoi64(value); err == nil {
			rv.SetInt(buf)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var buf uint64
		if buf, err = Atoui64(value); err == nil {
			rv.SetUint(buf)
		}
	case reflect.Float32, reflect.Float64:
		var buf float64
		if buf, err = Atof64(value); err == nil {
			rv.SetFloat(buf)
		}
	case reflect.Complex64, reflect.Complex128:
		var buf complex128
		if buf, err = Atoc128(value); err == nil {
			rv.SetComplex(buf)
		}
	case reflect.String:
		rv.SetString(value)
	default:
		switch rv.Interface().(type) {
		case time.Time, *time.Time:
			var v reflect.Value
			if v, err = deserializeTime(value, rv.Type().Kind()); err == nil {
				rv.Set(v)
			}
		default:
			var dst reflect.Value
			switch rv.Type().Kind() {
			case reflect.Pointer:
				elm := reflect.New(rv.Type().Elem()).Elem()
				err = json.Unmarshal([]byte(value), elm.Addr().Interface())
				dst = elm.Addr()
			default:
				elm := reflect.New(rv.Type()).Elem()
				err = json.Unmarshal([]byte(value), elm.Addr().Interface())
				dst = elm
			}
			rv.Set(dst)
		}
	}
	return
}

func showFlatMapKey(sf reflect.StructField, keyDic ...map[string]string) (key string) {
	if key = strings.Split(sf.Tag.Get("json"), ",")[0]; key == "-" || len(key) == 0 {
		for j := range keyDic {
			if v, ok := keyDic[j][sf.Name]; ok {
				key = v
				break
			}
		}
		if len(key) == 0 {
			key = sf.Name
		}
	}
	return
}

// Struct2FlatMap serializes a structure into a one-level map[string]string.
// The one-level map[string]string uses the `json` tag of the structure entered
// for the key. If undefined, it substitutes the field name of the structure.
// If you want to specify an arbitrary key for an arbitrary field, you can
// specify map[string]string with the field name as key in keyDic.
// keyDic is used in preference to the `json` tag.
func Struct2FlatMap(structure any, keyDic ...map[string]string) (ret map[string]string, err error) {
	rv := reflect.ValueOf(structure)
	if rv.Type().Kind() == reflect.Struct {
		ret = make(map[string]string)
		var key string
		for i, l := 0, rv.NumField(); i < l; i++ {
			key = showFlatMapKey(rv.Type().Field(i), keyDic...)
			if ret[key], err = serialize(rv.Field(i)); err != nil {
				err = errors.New(fmt.Sprintf(err.Error(), key, rv.Type().Field(i).Type.String()))
			}
		}
	} else {
		err = errors.New(inputValueTypeError)
	}
	return
}

// FlatMap2Struct deserializes a one-level map[string]string into an arbitrary
// structure.
// The one-level map[string]string uses the `json` tag of the structure entered
// for the key. If undefined, it substitutes the field name of the structure.
// If you want to specify an arbitrary key for an arbitrary field, you can
// specify map[string]string with the field name as key in keyDic.
// keyDic is used in preference to the `json` tag.
func FlatMap2Struct(flatMap map[string]string, structure any, keyDic ...map[string]string) (err error) {
	rv := reflect.ValueOf(structure)
	if rv.Type().Kind() == reflect.Pointer && rv.Elem().Kind() == reflect.Struct {
		var key string
		for i, l := 0, rv.Elem().NumField(); i < l; i++ {
			key = showFlatMapKey(rv.Elem().Type().Field(i), keyDic...)

			if v, ok := flatMap[key]; ok {
				err = deserialize(v, rv.Elem().Field(i))
			}
		}
	} else {
		err = errors.New(structureTypeError)
	}
	return
}
