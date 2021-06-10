package check

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	AddCheckFunc("in", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := inCheck(argValue, false)
		if err != nil {
			return nil, ErrArgumentType("in", "integerArray", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			return cmp(value)
		}), nil
	}))

	AddCheckFunc("nin", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := inCheck(argValue, false)
		if err != nil {
			return nil, ErrArgumentType("nin", "integerArray", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			ret, err := cmp(value)
			return !ret, err
		}), nil
	}))
}

func inIntArray(exceptedArray []int64, actualValue int64) bool {
	for idx := range exceptedArray {
		if actualValue == exceptedArray[idx] {
			return true
		}
	}
	return false
}

func inIntArrayU(exceptedArray []int64, actualValue uint64) bool {
	for idx := range exceptedArray {
		if exceptedArray[idx] < 0 {
			continue
		}

		if actualValue == uint64(exceptedArray[idx]) {
			return true
		}
	}
	return false
}

func inIntArrayF(exceptedArray []int64, actualValue float64) bool {
	for idx := range exceptedArray {
		if actualValue == float64(exceptedArray[idx]) {
			return true
		}
	}
	return false
}

func inIntArrayCheck(exceptedArray []int64) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case int:
			return inIntArray(exceptedArray, int64(actualValue)), nil
		case int8:
			return inIntArray(exceptedArray, int64(actualValue)), nil
		case int16:
			return inIntArray(exceptedArray, int64(actualValue)), nil
		case int32:
			return inIntArray(exceptedArray, int64(actualValue)), nil
		case int64:
			return inIntArray(exceptedArray, actualValue), nil
		case uint:
			return inIntArrayU(exceptedArray, uint64(actualValue)), nil
		case uint8:
			return inIntArrayU(exceptedArray, uint64(actualValue)), nil
		case uint16:
			return inIntArrayU(exceptedArray, uint64(actualValue)), nil
		case uint32:
			return inIntArrayU(exceptedArray, uint64(actualValue)), nil
		case uint64:
			return inIntArrayU(exceptedArray, actualValue), nil
		case float32:
			return inIntArrayF(exceptedArray, float64(actualValue)), nil
		case float64:
			return inIntArrayF(exceptedArray, actualValue), nil
		case []byte:
			s := string(actualValue)
			i64, err := strconv.ParseInt(s, 10, 64)
			if nil == err {
				return inIntArray(exceptedArray, i64), nil
			}
			u64, err := strconv.ParseUint(s, 10, 64)
			if nil == err {
				return inIntArrayU(exceptedArray, u64), nil
			}
		case string:
			i64, err := strconv.ParseInt(actualValue, 10, 64)
			if nil == err {
				return inIntArray(exceptedArray, i64), nil
			}
			u64, err := strconv.ParseUint(actualValue, 10, 64)
			if nil == err {
				return inIntArrayU(exceptedArray, u64), nil
			}
		case json.Number:
			i64, err := strconv.ParseInt(actualValue.String(), 10, 64)
			if nil == err {
				return inIntArray(exceptedArray, i64), nil
			}
			u64, err := strconv.ParseUint(actualValue.String(), 10, 64)
			if nil == err {
				return inIntArrayU(exceptedArray, u64), nil
			}

		case *json.Number:
			i64, err := strconv.ParseInt(actualValue.String(), 10, 64)
			if nil == err {
				return inIntArray(exceptedArray, i64), nil
			}
			u64, err := strconv.ParseUint(actualValue.String(), 10, 64)
			if nil == err {
				return inIntArrayU(exceptedArray, u64), nil
			}
		}
		if nil == value {
			return false, ErrValueNull
		}
		return false, errType(value, "int64")
	}
}

func inUintArray(exceptedArray []uint64, actualValue uint64) bool {
	for idx := range exceptedArray {
		if actualValue == exceptedArray[idx] {
			return true
		}
	}
	return false
}

func inUintArrayF(exceptedArray []uint64, actualValue float64) bool {
	for idx := range exceptedArray {
		if actualValue == float64(exceptedArray[idx]) {
			return true
		}
	}
	return false
}

func inUintArrayCheck(exceptedArray []uint64) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case int:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case int8:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case int16:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case int32:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case int64:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case uint:
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case uint8:
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case uint16:
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case uint32:
			return inUintArray(exceptedArray, uint64(actualValue)), nil
		case uint64:
			return inUintArray(exceptedArray, actualValue), nil
		case float32:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArrayF(exceptedArray, float64(actualValue)), nil
		case float64:
			if actualValue < 0 {
				return false, nil
			}
			return inUintArrayF(exceptedArray, actualValue), nil
		case []byte:
			s := string(actualValue)
			if strings.HasPrefix(s, "-") {
				return false, nil
			}
			u64, err := strconv.ParseUint(s, 10, 64)
			if nil == err {
				return inUintArray(exceptedArray, u64), nil
			}
		case string:
			if strings.HasPrefix(actualValue, "-") {
				return false, nil
			}
			u64, err := strconv.ParseUint(actualValue, 10, 64)
			if nil == err {
				return inUintArray(exceptedArray, u64), nil
			}
		case json.Number:
			if strings.HasPrefix(actualValue.String(), "-") {
				return false, nil
			}
			u64, err := strconv.ParseUint(actualValue.String(), 10, 64)
			if nil == err {
				return inUintArray(exceptedArray, u64), nil
			}
		case *json.Number:
			if strings.HasPrefix(actualValue.String(), "-") {
				return false, nil
			}
			u64, err := strconv.ParseUint(actualValue.String(), 10, 64)
			if nil == err {
				return inUintArray(exceptedArray, u64), nil
			}
		}
		if nil == value {
			return false, ErrValueNull
		}
		return false, errType(value, "int64")
	}
}

func inStringArrayCheck(exceptedArray []string) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		if value == nil {
			return false, nil
		}
		actualValue := fmt.Sprint(value)
		for idx := range exceptedArray {
			if actualValue == exceptedArray[idx] {
				return true, nil
			}
		}
		return false, nil
	}
}

func asUint64(a interface{}, mustInt bool) (uint64, bool) {
	switch i := a.(type) {
	case uint:
		return uint64(i), true
	case uint8:
		return uint64(i), true
	case uint16:
		return uint64(i), true
	case uint32:
		return uint64(i), true
	case uint64:
		return i, true
	case int:
		if i < 0 {
			return 0, false
		}
		return uint64(i), true
	case int8:
		if i < 0 {
			return 0, false
		}
		return uint64(i), true
	case int16:
		if i < 0 {
			return 0, false
		}
		return uint64(i), true
	case int32:
		if i < 0 {
			return 0, false
		}
		return uint64(i), true
	case int64:
		if i < 0 {
			return 0, false
		}
		return uint64(i), true

	case float32:
		if !mustInt {
			if i < 0 {
				return 0, false
			}
			if i > math.MaxUint64 {
				return 0, false
			}
			return uint64(i), true
		}
	case float64:
		if !mustInt {
			if i < 0 {
				return 0, false
			}
			if i > math.MaxUint64 {
				return 0, false
			}
			return uint64(i), true
		}
	case json.Number:
		i64, err := strconv.ParseUint(i.String(), 10, 64)
		if err == nil {
			return i64, true
		}
	case *json.Number:
		i64, err := strconv.ParseUint(i.String(), 10, 64)
		if err == nil {
			return i64, true
		}
	case []byte:
		if !mustInt {
			if len(i) == 0 {
				return 0, false
			}

			u64, err := strconv.ParseUint(string(i), 10, 64)
			if err == nil {
				return u64, true
			}
		}
	case string:
		if !mustInt {
			u64, err := strconv.ParseUint(i, 10, 64)
			if err == nil {
				return u64, true
			}
		}
	}
	return 0, false
}

func asInt64(a interface{}, mustInt bool) (int64, bool) {
	switch i := a.(type) {
	case uint:
		if i > math.MaxInt64 {
			return 0, false
		}
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint64:
		if i > math.MaxInt64 {
			return 0, false
		}
		return int64(i), true
	case int:
		return int64(i), true
	case int8:
		return int64(i), true
	case int16:
		return int64(i), true
	case int32:
		return int64(i), true
	case int64:
		return i, true
	case float32:
		if !mustInt {
			if i > math.MaxInt64 {
				return 0, false
			}
			return int64(i), true
		}
	case float64:
		if !mustInt {
			if i > math.MaxInt64 {
				return 0, false
			}
			return int64(i), true
		}
	case json.Number:
		i64, err := strconv.ParseInt(i.String(), 10, 64)
		if err == nil {
			return i64, true
		}
	case *json.Number:
		i64, err := strconv.ParseInt(i.String(), 10, 64)
		if err == nil {
			return i64, true
		}
	case []byte:
		if !mustInt {
			if len(i) == 0 {
				return 0, false
			}

			i64, err := strconv.ParseInt(string(i), 10, 64)
			if err == nil {
				return i64, true
			}
		}
	case string:
		if !mustInt {
			i64, err := strconv.ParseInt(i, 10, 64)
			if err == nil {
				return i64, true
			}
		}
	}
	return 0, false
}

func asFloat64(a interface{}) (float64, bool) {
	switch i := a.(type) {
	case uint:
		return float64(i), true
	case uint8:
		return float64(i), true
	case uint16:
		return float64(i), true
	case uint32:
		return float64(i), true
	case uint64:
		return float64(i), true
	case int:
		return float64(i), true
	case int8:
		return float64(i), true
	case int16:
		return float64(i), true
	case int32:
		return float64(i), true
	case int64:
		return float64(i), true
	case float32:
			return float64(i), true
	case float64:
			return i, true
	case json.Number:
		f64, err := strconv.ParseFloat(i.String(), 64)
		if err == nil {
			return f64, true
		}
	case *json.Number:
		f64, err := strconv.ParseFloat(i.String(),  64)
		if err == nil {
			return f64, true
		}
	// case []byte:
	// 	if mustInt {
	// 		if len(i) == 0 {
	// 			return 0, false
	// 		}

	// 		f64, err := strconv.ParseFloat(string(i), 10, 64)
	// 		if err == nil {
	// 			return f64, true
	// 		}
	// 	}
	case string:
		// if mustInt {
			f64, err := strconv.ParseFloat(i, 64)
			if err == nil {
				return f64, true
			}
		// }
	}
	return 0, false
}

func inArrayCheck(exceptedArray []interface{}, mustInt bool) (func(interface{}) (bool, error), error) {
	ints := make([]int64, len(exceptedArray))
	for i, a := range exceptedArray {
		iv, ok := asInt64(a, mustInt)
		if !ok {
			goto _uint64
		}
		ints[i] = iv
	}
	return inIntArrayCheck(ints), nil
_uint64:
	uints := make([]uint64, len(exceptedArray))
	for i, a := range exceptedArray {
		iv, ok := asUint64(a, mustInt)
		if !ok {
			if mustInt {
				return nil, errType(exceptedArray, "intArray")
			}
			goto _strings
		}
		uints[i] = iv
	}
	return inUintArrayCheck(uints), nil
_strings:
	ss := make([]string, len(exceptedArray))
	for i, a := range exceptedArray {
		s, ok := a.(string)
		if !ok {
			return nil, errType(exceptedArray, "strArray")
		}
		ss[i] = s
	}
	return inStringArrayCheck(ss), nil
}

func inAnyToStringArrayCheck(exceptedArray []interface{}) (func(interface{}) (bool, error), error) {
	ss := make([]string, len(exceptedArray))
	for i, a := range exceptedArray {
		ss[i] = fmt.Sprint(a)
	}
	return inStringArrayCheck(ss), nil
}

func inCheck(value interface{}, mustInt bool) (func(interface{}) (bool, error), error) {
	switch a := value.(type) {
	case string:
		ss := strings.Split(a, ",")
		ints := make([]int64, 0, len(ss))
		for _, s := range ss {
			if s == "" {
				continue
			}
			iv, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				if mustInt {
					return nil, errType(value, "intArray")
				}
				return inStringArrayCheck(ss), nil
			}
			ints = append(ints, iv)
		}
		return inIntArrayCheck(ints), nil
	case []string:
		ints := make([]int64, 0, len(a))
		for _, s := range a {
			if s == "" {
				continue
			}
			iv, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				if mustInt {
					return nil, errType(value, "intArray")
				}
				return inStringArrayCheck(a), nil
			}
			ints = append(ints, iv)
		}
		return inIntArrayCheck(ints), nil
	case []int64:
		return inIntArrayCheck(a), nil
	case []int:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return inIntArrayCheck(ints), nil
	case []int8:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return inIntArrayCheck(ints), nil
	case []int16:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return inIntArrayCheck(ints), nil
	case []int32:
		ints := make([]int64, len(a))
		for i := range a {
			ints[i] = int64(a[i])
		}
		return inIntArrayCheck(ints), nil
	case []uint64:
		return inUintArrayCheck(a), nil
	case []uint:
		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return inUintArrayCheck(uints), nil
	case []uint8:
		if mustInt && len(a) != 0 {
			u64, err := strconv.ParseUint(string(a), 10, 64)
			if err == nil {
				return uintEquals(u64), nil
			}
			i64, err := strconv.ParseInt(string(a), 10, 64)
			if err == nil {
				return intEquals(i64), nil
			}
		}

		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return inUintArrayCheck(uints), nil
	case []uint16:
		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return inUintArrayCheck(uints), nil
	case []uint32:
		uints := make([]uint64, len(a))
		for i := range a {
			uints[i] = uint64(a[i])
		}
		return inUintArrayCheck(uints), nil
	case []interface{}:
		checkFunc, err := inArrayCheck(a, mustInt)
		if err == nil {
			return checkFunc, nil
		}
		if mustInt {
			return nil, err
		}
		return inAnyToStringArrayCheck(a)
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, json.Number, *json.Number, time.Duration:
		return DynamicEquals(a)
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			aLen := rv.Len()
			values := make([]interface{}, aLen)
			for i := 0; i < aLen; i++ {
				values[i] = rv.Index(i).Interface()
			}
			return inArrayCheck(values, mustInt)
		}
	}
	return nil, errType(value, "Array")
}
