package check

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/three-plus-three/modules/as"
)

func intCheck(argValue int64) func(value interface{}) (int, error) {
	return nil
}

func uintCheck(argValue uint64) func(value interface{}) (int, error) {
	return nil
}

func floatCheck(argValue float64) func(value interface{}) (int, error) {
	return nil
}

func stringAsNumberCheck(argValue string) (func(value interface{}) (int, error), error) {
	return nil, errors.New("not implement")
}

func numberCheck(argValue interface{}) (func(value interface{}) (int, error), error) {
	switch v := argValue.(type) {
	case int:
		return intCheck(int64(v)), nil
	case int8:
		return intCheck(int64(v)), nil
	case int16:
		return intCheck(int64(v)), nil
	case int32:
		return intCheck(int64(v)), nil
	case int64:
		return intCheck(v), nil
	case uint:
		return uintCheck(uint64(v)), nil
	case uint8:
		return uintCheck(uint64(v)), nil
	case uint16:
		return uintCheck(uint64(v)), nil
	case uint32:
		return uintCheck(uint64(v)), nil
	case uint64:
		return uintCheck(v), nil
	case float32:
		return floatCheck(float64(v)), nil
	case float64:
		return floatCheck(v), nil
	case []byte:
		return stringAsNumberCheck(string(v))
	case string:
		return stringAsNumberCheck(v)
	case json.Number:
		return stringAsNumberCheck(v.String())
	}
	if nil == argValue {
		return nil, ErrValueNull
	}
	return nil, errType(argValue, "int64")
}

func init() {
	AddCheckFunc(">", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType(">", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType(">", "integer", value)
			}
			return r < 0, nil
		}), nil
	}))

	AddCheckFunc(">=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType(">=", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType(">=", "integer", value)
			}
			return r <= 0, nil
		}), nil
	}))

	AddCheckFunc("<", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("<", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("<", "integer", value)
			}
			return r > 0, nil
		}), nil
	}))

	AddCheckFunc("<=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("<=", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("<=", "integer", value)
			}
			return r >= 0, nil
		}), nil
	}))

	AddCheckFunc("=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("=", "integer", value)
			}
			return r == 0, nil
		}), nil
	}))

	AddCheckFunc("!=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := numberCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "integer", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("!=", "integer", value)
			}
			return r != 0, nil
		}), nil
	}))

	AddCheckFunc("in", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toInt64Array(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "Int64Array", argValue)
			}
			exceptedArray, err = as.Int64Array(strings.Split(svalue, ","))
			if err != nil {
				return nil, ErrArgumentType("in", "Int64Array", argValue)
			}
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toInt64(value)
			if err != nil {
				return false, ErrActualType("in", "integer", value)
			}
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return true, nil
				}
			}
			return false, nil
		}), nil
	}))

	AddCheckFunc("nin", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toInt64Array(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "Int64Array", argValue)
			}
			exceptedArray, err = as.Int64Array(strings.Split(svalue, ","))
			if err != nil {
				return nil, ErrArgumentType("in", "Int64Array", argValue)
			}
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toInt64(value)
			if err != nil {
				return false, ErrActualType("nin", "integer", value)
			}
			found := false
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					found = true
					break
				}
			}
			return !found, nil
		}), nil
	}))
}

func toInt64Array(value interface{}) ([]int64, error) {
	switch a := value.(type) {
	case []interface{}:
		ints := make([]int64, len(a))
		for i := range a {
			iv, err := toInt64(a[i])
			if err != nil {
				return nil, errType(value, "int64Array")
			}
			ints[i] = iv
		}
		return ints, nil
	case []string:
		ints := make([]int64, len(a))
		for i := range a {
			iv, err := strconv.ParseInt(a[i], 10, 64)
			if err != nil {
				return nil, errType(value, "int64Array")
			}
			ints[i] = iv
		}
		return ints, nil
	case []int64:
		return a, nil
	case []int:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return ints, nil
	case []int32:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return ints, nil
	case []uint64:
		ints := make([]int64, len(a))
		for i := range a {
			if a[i] > math.MaxInt64 {
				return nil, errType(value, "int64Array")
			}
			ints[i] = int64(a[i])
		}
		return ints, nil
	case []uint:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return ints, nil
	case []uint32:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return ints, nil
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.Kind() == reflect.Slice {
			aLen := rv.Len()
			ints := make([]int64, aLen)
			for i := 0; i < aLen; i++ {
				iv, err := toInt64(rv.Index(i).Interface())
				if err != nil {
					return nil, errType(value, "int64Array")
				}
				ints[i] = iv
			}
			return ints, nil
		}
	}
	return nil, errType(value, "int64Array")
}

func toUint64Array(value interface{}) ([]uint64, error) {
	switch a := value.(type) {
	case []interface{}:
		uints := make([]uint64, len(a))
		for i := range a {
			iv, err := toUint64(a[i])
			if err != nil {
				return nil, errType(value, "uint64Array")
			}
			uints[i] = iv
		}
		return uints, nil
	case []string:
		uints := make([]uint64, len(a))
		for i := range a {
			iv, err := strconv.ParseUint(a[i], 10, 64)
			if err != nil {
				return nil, errType(value, "uint64Array")
			}
			uints[i] = iv
		}
		return uints, nil
	case []int64:
		uints := make([]uint64, len(a))
		for i := range a {
			if a[i] < 0 {
				return nil, errType(value, "uint64Array")
			}
			uints[i] = uint64(a[i])
		}
		return uints, nil
	case []int:
		uints := make([]uint64, len(a))
		for i := range a {
			if a[i] < 0 {
				return nil, errType(value, "uint64Array")
			}
			uints[i] = uint64(a[i])
		}
		return uints, nil
	case []int32:
		uints := make([]uint64, len(a))
		for i := range a {
			if a[i] < 0 {
				return nil, errType(value, "uint64Array")
			}
			uints[i] = uint64(a[i])
		}
		return uints, nil
	case []uint64:
		return a, nil
	case []uint:
		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return uints, nil
	case []uint32:
		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return uints, nil
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.Kind() == reflect.Slice {
			aLen := rv.Len()
			uints := make([]uint64, aLen)
			for i := 0; i < aLen; i++ {
				iv, err := toUint64(rv.Index(i).Interface())
				if err != nil {
					return nil, errType(value, "uint64Array")
				}
				uints[i] = iv
			}
			return uints, nil
		}
	}
	return nil, errType(value, "uint64Array")
}

// Int type AsSerts to `float64` then converts to `int`
func toInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		if 9223372036854775807 >= int64(v) {
			return int64(v), nil
		}
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if 9223372036854775807 >= v {
			return int64(v), nil
		}
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case []byte:
		i64, err := strconv.ParseInt(string(v), 10, 64)
		if nil == err {
			return i64, nil
		}
	case string:
		i64, err := strconv.ParseInt(v, 10, 64)
		if nil == err {
			return i64, nil
		}
	case json.Number:
		i64, err := v.Int64()
		if nil == err {
			return i64, nil
		}
		f64, err := v.Float64()
		if nil == err {
			return int64(f64), nil
		}
	}
	if nil == value {
		return 0, ErrValueNull
	}
	return 0, errType(value, "int64")
}

func toUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case []byte:
		i64, err := strconv.ParseUint(string(v), 10, 64)
		if nil == err {
			return i64, nil
		}
	case string:
		i64, err := strconv.ParseUint(v, 10, 64)
		if nil == err {
			return i64, nil
		}
		return i64, errType(value, "uint64")

	case json.Number:
		i64, err := strconv.ParseUint(v.String(), 10, 64)
		if nil == err {
			return i64, nil
		}
		f64, err := v.Float64()
		if nil == err {
			if f64 >= 0 {
				return uint64(f64), nil
			}
			if f64 < 0 {
				if math.IsNaN(f64) {
					return 0, nil
				}
				if int64(f64) == 0 {
					return 0, nil
				}
			}
		} else {
			return 0, errType(value, "uint64")
		}
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case int:
		if v >= 0 {
			return uint64(v), nil
		}
	case int8:
		if v >= 0 {
			return uint64(v), nil
		}
	case int16:
		if v >= 0 {
			return uint64(v), nil
		}
	case int32:
		if v >= 0 {
			return uint64(v), nil
		}
	case int64:
		if v >= 0 {
			return uint64(v), nil
		}
	case float32:
		if v >= 0 && 18446744073709551615 >= v {
			return uint64(v), nil
		}

		if v < 0 {
			if math.IsNaN(float64(v)) {
				return 0, nil
			}
			if int64(v) == 0 {
				return 0, nil
			}
		}
	case float64:
		if v >= 0 && 18446744073709551615 >= v {
			return uint64(v), nil
		}

		if v < 0 {
			if math.IsNaN(v) {
				return 0, nil
			}
			if int64(v) == 0 {
				return 0, nil
			}
		}
	}
	if nil == value {
		return 0, ErrValueNull
	}
	return 0, errType(value, "uint64")
}
