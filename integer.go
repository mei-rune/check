package check

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func int64strCmp(a int64, bvalue string) (int, error) {
	if bvalue[0] == '-' {
		b64, err := strconv.ParseInt(bvalue, 10, 64)
		if nil == err {
			return int(a - b64), nil
		}
	} else {
		b64, err := strconv.ParseUint(bvalue, 10, 64)
		if nil == err {
			return int64uint64Cmp(a, b64), nil
		}
	}

	f64, err := strconv.ParseFloat(bvalue, 64)
	if nil != err {
		return 0, errType(bvalue, "int64")
	}
	if float64(a) < f64 {
		return -1, nil
	}
	if float64(a) == f64 {
		return 0, nil
	}
	return 1, nil
}

func int64uint64Cmp(a int64, b uint64) int {
	if a < 0 || uint64(a) < b {
		return -1
	}
	if uint64(a) == b {
		return 0
	}
	return 1
}

func intCheck(exceptedValue int64) func(value interface{}) (int, error) {
	return func(value interface{}) (int, error) {
		switch actualValue := value.(type) {
		case []byte:
			return int64strCmp(exceptedValue, string(actualValue))
		case string:
			return int64strCmp(exceptedValue, actualValue)
		case json.Number:
			return int64strCmp(exceptedValue, actualValue.String())
		case *json.Number:
			return int64strCmp(exceptedValue, actualValue.String())
		case uint:
			return int64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint8:
			return int64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint16:
			return int64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint32:
			return int64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint64:
			return int64uint64Cmp(exceptedValue, actualValue), nil
		case int:
			return int(exceptedValue - int64(actualValue)), nil
		case int8:
			return int(exceptedValue - int64(actualValue)), nil
		case int16:
			return int(exceptedValue - int64(actualValue)), nil
		case int32:
			return int(exceptedValue - int64(actualValue)), nil
		case int64:
			return int(exceptedValue - actualValue), nil
		case float32:
			if float64(exceptedValue) < float64(actualValue) {
				return -1, nil
			}
			if float64(exceptedValue) == float64(actualValue) {
				return 0, nil
			}
			return 1, nil
		case float64:
			if float64(exceptedValue) < actualValue {
				return -1, nil
			}
			if float64(exceptedValue) == actualValue {
				return 0, nil
			}
			return 1, nil
		}
		if nil == value {
			return 0, ErrValueNull
		}
		return 0, errType(value, "int64")
	}
}

func uint64strCmp(a uint64, bvalue string) (int, error) {
	if bvalue[0] == '-' {
		_, err := strconv.ParseInt(bvalue, 10, 64)
		if nil == err {
			return -1, nil
		}
	} else {
		b64, err := strconv.ParseUint(bvalue, 10, 64)
		if nil == err {
			return uint64uint64Cmp(a, b64), nil
		}
	}

	f64, err := strconv.ParseFloat(bvalue, 64)
	if nil != err {
		return 0, errType(bvalue, "uint64")
	}
	if float64(a) < f64 {
		return -1, nil
	}
	if float64(a) == f64 {
		return 0, nil
	}
	return 1, nil
}

func uint64uint64Cmp(a, b uint64) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}
func uint64int64Cmp(a uint64, b int64) int {
	if b < 0 || uint64(b) < a {
		return 1
	}
	if uint64(b) == a {
		return 0
	}
	return -1
}

func uintCheck(exceptedValue uint64) func(value interface{}) (int, error) {
	return func(value interface{}) (int, error) {
		switch actualValue := value.(type) {
		case []byte:
			return uint64strCmp(exceptedValue, string(actualValue))
		case string:
			return uint64strCmp(exceptedValue, actualValue)
		case json.Number:
			return uint64strCmp(exceptedValue, actualValue.String())
		case *json.Number:
			return uint64strCmp(exceptedValue, actualValue.String())
		case uint:
			return uint64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint8:
			return uint64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint16:
			return uint64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint32:
			return uint64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint64:
			return uint64uint64Cmp(exceptedValue, actualValue), nil
		case int:
			return uint64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int8:
			return uint64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int16:
			return uint64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int32:
			return uint64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int64:
			return uint64int64Cmp(exceptedValue, int64(actualValue)), nil
		case float32:
			if actualValue < 0 {
				return 1, nil
			}

			if float64(exceptedValue) > float64(actualValue) {
				return 1, nil
			}

			if float64(exceptedValue) == float64(actualValue) {
				return 0, nil
			}
			return -1, nil
		case float64:
			if actualValue < 0 {
				return 1, nil
			}

			if float64(exceptedValue) > float64(actualValue) {
				return 1, nil
			}

			if float64(exceptedValue) == float64(actualValue) {
				return 0, nil
			}
			return -1, nil
		}
		if nil == value {
			return 0, ErrValueNull
		}
		return 0, errType(value, "uint64")
	}
}

func float64strCmp(a float64, bvalue string) (int, error) {
	f64, err := strconv.ParseFloat(bvalue, 64)
	if nil != err {
		return 0, errType(bvalue, "uint64")
	}
	if a < f64 {
		return -1, nil
	}
	if a == f64 {
		return 0, nil
	}
	return 1, nil
}

func float64uint64Cmp(a float64, b uint64) int {
	if a < 0 {
		return -1
	}

	fb := float64(b)
	if a < fb {
		return -1
	}
	if a == fb {
		return 0
	}
	return 1
}
func float64int64Cmp(a float64, b int64) int {
	if b < 0 {
		if a > 0 {
			return 1
		}
	}
	fb := float64(b)
	if fb < a {
		return 1
	}
	if fb == a {
		return 0
	}
	return -1
}

func floatCheck(exceptedValue float64) func(value interface{}) (int, error) {
	return func(value interface{}) (int, error) {
		switch actualValue := value.(type) {
		case []byte:
			return float64strCmp(exceptedValue, string(actualValue))
		case string:
			return float64strCmp(exceptedValue, actualValue)
		case json.Number:
			return float64strCmp(exceptedValue, actualValue.String())
		case *json.Number:
			return float64strCmp(exceptedValue, actualValue.String())
		case uint:
			return float64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint8:
			return float64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint16:
			return float64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint32:
			return float64uint64Cmp(exceptedValue, uint64(actualValue)), nil
		case uint64:
			return float64uint64Cmp(exceptedValue, actualValue), nil
		case int:
			return float64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int8:
			return float64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int16:
			return float64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int32:
			return float64int64Cmp(exceptedValue, int64(actualValue)), nil
		case int64:
			return float64int64Cmp(exceptedValue, int64(actualValue)), nil
		case float32:
			if exceptedValue > float64(actualValue) {
				return 1, nil
			}

			if exceptedValue == float64(actualValue) {
				return 0, nil
			}
			return -1, nil
		case float64:
			if exceptedValue > float64(actualValue) {
				return 1, nil
			}

			if exceptedValue == actualValue {
				return 0, nil
			}
			return -1, nil
		}
		if nil == value {
			return 0, ErrValueNull
		}
		return 0, errType(value, "float64")
	}
}

func stringAsNumberCheck(argValue string) (func(value interface{}) (int, error), error) {
	if strings.HasSuffix(argValue, ".0") {
		argValue = strings.TrimSuffix(argValue, ".0")
	}
	if argValue == "" {
		return nil, errType(argValue, "number")
	}
	if argValue[0] == '-' {
		i64, err := strconv.ParseInt(argValue, 10, 64)
		if err == nil {
			return intCheck(i64), nil
		}
	} else {
		u64, err := strconv.ParseUint(argValue, 10, 64)
		if err == nil {
			return uintCheck(u64), nil
		}
	}

	f64, err := strconv.ParseFloat(argValue, 64)
	if nil != err {
		return nil, errType(argValue, "number")
	}
	return floatCheck(f64), nil
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
	case *json.Number:
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
			//fmt.Printf("1(%T) %v > (%T) %v   = %v\r\n", argValue, argValue, value, value, r)
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

			//fmt.Printf("2(%T) %v >= (%T) %v   = %v\r\n", argValue, argValue, value, value, r)
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
			exceptedArray, err = toInt64Array(strings.Split(svalue, ","))
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
			exceptedArray, err = toInt64Array(strings.Split(svalue, ","))
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
		// f64, err := v.Float64()
		// if nil == err {
		// 	return int64(f64), nil
		// }

	case *json.Number:
		i64, err := v.Int64()
		if nil == err {
			return i64, nil
		}
		// f64, err := v.Float64()
		// if nil == err {
		// 	return int64(f64), nil
		// }
	}
	if nil == value {
		return 0, ErrValueNull
	}
	return 0, errType(value, "int64")
}

// func toUint64Array(value interface{}) ([]uint64, error) {
// 	switch a := value.(type) {
// 	case []interface{}:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			iv, err := toUint64(a[i])
// 			if err != nil {
// 				return nil, errType(value, "uint64Array")
// 			}
// 			uints[i] = iv
// 		}
// 		return uints, nil
// 	case []string:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			iv, err := strconv.ParseUint(a[i], 10, 64)
// 			if err != nil {
// 				return nil, errType(value, "uint64Array")
// 			}
// 			uints[i] = iv
// 		}
// 		return uints, nil
// 	case []int64:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			if a[i] < 0 {
// 				return nil, errType(value, "uint64Array")
// 			}
// 			uints[i] = uint64(a[i])
// 		}
// 		return uints, nil
// 	case []int:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			if a[i] < 0 {
// 				return nil, errType(value, "uint64Array")
// 			}
// 			uints[i] = uint64(a[i])
// 		}
// 		return uints, nil
// 	case []int32:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			if a[i] < 0 {
// 				return nil, errType(value, "uint64Array")
// 			}
// 			uints[i] = uint64(a[i])
// 		}
// 		return uints, nil
// 	case []uint64:
// 		return a, nil
// 	case []uint:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			uints[i] = uint64(a[i])
// 		}
// 		return uints, nil
// 	case []uint32:
// 		uints := make([]uint64, len(a))
// 		for i := range a {
// 			uints[i] = uint64(a[i])
// 		}
// 		return uints, nil
// 	default:
// 		rv := reflect.ValueOf(value)
// 		if rv.Kind() == reflect.Ptr {
// 			rv = rv.Elem()
// 		}

// 		if rv.Kind() == reflect.Slice {
// 			aLen := rv.Len()
// 			uints := make([]uint64, aLen)
// 			for i := 0; i < aLen; i++ {
// 				iv, err := toUint64(rv.Index(i).Interface())
// 				if err != nil {
// 					return nil, errType(value, "uint64Array")
// 				}
// 				uints[i] = iv
// 			}
// 			return uints, nil
// 		}
// 	}
// 	return nil, errType(value, "uint64Array")
// }

// func toUint64(value interface{}) (uint64, error) {
// 	switch v := value.(type) {
// 	case []byte:
// 		i64, err := strconv.ParseUint(string(v), 10, 64)
// 		if nil == err {
// 			return i64, nil
// 		}
// 	case string:
// 		i64, err := strconv.ParseUint(v, 10, 64)
// 		if nil == err {
// 			return i64, nil
// 		}
// 		return i64, errType(value, "uint64")

// 	case json.Number:
// 		i64, err := strconv.ParseUint(v.String(), 10, 64)
// 		if nil == err {
// 			return i64, nil
// 		}
// 		f64, err := v.Float64()
// 		if nil == err {
// 			if f64 >= 0 {
// 				return uint64(f64), nil
// 			}
// 			if f64 < 0 {
// 				if math.IsNaN(f64) {
// 					return 0, nil
// 				}
// 				if int64(f64) == 0 {
// 					return 0, nil
// 				}
// 			}
// 		} else {
// 			return 0, errType(value, "uint64")
// 		}
// 	case uint:
// 		return uint64(v), nil
// 	case uint8:
// 		return uint64(v), nil
// 	case uint16:
// 		return uint64(v), nil
// 	case uint32:
// 		return uint64(v), nil
// 	case uint64:
// 		return v, nil
// 	case int:
// 		if v >= 0 {
// 			return uint64(v), nil
// 		}
// 	case int8:
// 		if v >= 0 {
// 			return uint64(v), nil
// 		}
// 	case int16:
// 		if v >= 0 {
// 			return uint64(v), nil
// 		}
// 	case int32:
// 		if v >= 0 {
// 			return uint64(v), nil
// 		}
// 	case int64:
// 		if v >= 0 {
// 			return uint64(v), nil
// 		}
// 	case float32:
// 		if v >= 0 && 18446744073709551615 >= v {
// 			return uint64(v), nil
// 		}

// 		if v < 0 {
// 			if math.IsNaN(float64(v)) {
// 				return 0, nil
// 			}
// 			if int64(v) == 0 {
// 				return 0, nil
// 			}
// 		}
// 	case float64:
// 		if v >= 0 && 18446744073709551615 >= v {
// 			return uint64(v), nil
// 		}

// 		if v < 0 {
// 			if math.IsNaN(v) {
// 				return 0, nil
// 			}
// 			if int64(v) == 0 {
// 				return 0, nil
// 			}
// 		}
// 	}
// 	if nil == value {
// 		return 0, ErrValueNull
// 	}
// 	return 0, errType(value, "uint64")
// }
