package check

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func floatEquals(exceptedValue float64) func(value interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case []byte:
			s := string(actualValue)
			if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
				return exceptedValue == float64(u64), nil
			}
			if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
				return exceptedValue == float64(i64), nil
			}
			if f64, e := strconv.ParseFloat(s, 64); e == nil {
				return exceptedValue == f64, nil
			}
			return false, nil
		case string:
			if u64, e := strconv.ParseUint(actualValue, 10, 64); e == nil {
				return exceptedValue == float64(u64), nil
			}
			if i64, e := strconv.ParseInt(actualValue, 10, 64); e == nil {
				return exceptedValue == float64(i64), nil
			}
			if f64, e := strconv.ParseFloat(actualValue, 64); e == nil {
				return exceptedValue == f64, nil
			}
			return false, nil
		case json.Number:
			s := actualValue.String()
			if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
				return exceptedValue == float64(u64), nil
			}
			if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
				return exceptedValue == float64(i64), nil
			}
			if f64, e := strconv.ParseFloat(s, 64); e == nil {
				return exceptedValue == f64, nil
			}
			return false, nil
		case *json.Number:
			s := actualValue.String()
			if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
				return exceptedValue == float64(u64), nil
			}
			if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
				return exceptedValue == float64(i64), nil
			}
			if f64, e := strconv.ParseFloat(s, 64); e == nil {
				return exceptedValue == f64, nil
			}
			return false, nil
		case uint:
			return exceptedValue == float64(actualValue), nil
		case uint8:
			return exceptedValue == float64(actualValue), nil
		case uint16:
			return exceptedValue == float64(actualValue), nil
		case uint32:
			return exceptedValue == float64(actualValue), nil
		case uint64:
			return exceptedValue == float64(actualValue), nil
		case int:
			return exceptedValue == float64(actualValue), nil
		case int8:
			return exceptedValue == float64(actualValue), nil
		case int16:
			return exceptedValue == float64(actualValue), nil
		case int32:
			return exceptedValue == float64(actualValue), nil
		case int64:
			return exceptedValue == float64(actualValue), nil
		case float32:
			return exceptedValue == float64(actualValue), nil
		case float64:
			return exceptedValue == actualValue, nil
		}
		if nil == value {
			return false, ErrValueNull
		}
		return false, errType(value, "uint64")
	}
}

func uintEquals(exceptedValue uint64) func(value interface{}) (bool, error) {
	exceptedStr := strconv.FormatUint(exceptedValue, 10)
	return func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case []byte:
			actual := string(actualValue)
			return exceptedStr == actual, nil
		case string:
			return exceptedStr == actualValue, nil
		case json.Number:
			if u64, e := strconv.ParseUint(actualValue.String(), 10, 64); e == nil {
				return u64 == exceptedValue, nil
			}
			return exceptedStr == actualValue.String(), nil
		case *json.Number:
			if u64, e := strconv.ParseUint(actualValue.String(), 10, 64); e == nil {
				return u64 == exceptedValue, nil
			}
			return exceptedStr == actualValue.String(), nil
		case uint:
			return uint64(actualValue) == exceptedValue, nil
		case uint8:
			return uint64(actualValue) == exceptedValue, nil
		case uint16:
			return uint64(actualValue) == exceptedValue, nil
		case uint32:
			return uint64(actualValue) == exceptedValue, nil
		case uint64:
			return actualValue == exceptedValue, nil
		case int:
			if actualValue < 0 {
				return false, nil
			}
			return exceptedValue == uint64(actualValue), nil
		case int8:
			if actualValue < 0 {
				return false, nil
			}
			return exceptedValue == uint64(actualValue), nil
		case int16:
			if actualValue < 0 {
				return false, nil
			}
			return exceptedValue == uint64(actualValue), nil
		case int32:
			if actualValue < 0 {
				return false, nil
			}
			return exceptedValue == uint64(actualValue), nil
		case int64:
			if actualValue < 0 {
				return false, nil
			}
			return exceptedValue == uint64(actualValue), nil
		case float32:
			if float64(exceptedValue) == float64(actualValue) {
				return true, nil
			}
			return false, nil
		case float64:
			if float64(exceptedValue) == actualValue {
				return true, nil
			}
			return false, nil
		}
		if nil == value {
			return false, ErrValueNull
		}
		return false, errType(value, "uint64")
	}
}

func intEquals(exceptedValue int64) func(value interface{}) (bool, error) {
	exceptedStr := strconv.FormatInt(exceptedValue, 10)
	return func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case []byte:
			actual := string(actualValue)
			return exceptedStr == actual, nil
		case string:
			return exceptedStr == actualValue, nil
		case json.Number:
			if i64, e := actualValue.Int64(); e == nil {
				return i64 == exceptedValue, nil
			}
			return exceptedStr == actualValue.String(), nil
		case *json.Number:
			if i64, e := actualValue.Int64(); e == nil {
				return i64 == exceptedValue, nil
			}
			return exceptedStr == actualValue.String(), nil
		case uint:
			if exceptedValue < 0 {
				return false, nil
			}
			return uint64(actualValue) == uint64(exceptedValue), nil
		case uint8:
			if exceptedValue < 0 {
				return false, nil
			}
			return uint64(actualValue) == uint64(exceptedValue), nil
		case uint16:
			if exceptedValue < 0 {
				return false, nil
			}
			return uint64(actualValue) == uint64(exceptedValue), nil
		case uint32:
			if exceptedValue < 0 {
				return false, nil
			}
			return uint64(actualValue) == uint64(exceptedValue), nil
		case uint64:
			if exceptedValue < 0 {
				return false, nil
			}
			return actualValue == uint64(exceptedValue), nil
		case int:
			return exceptedValue == int64(actualValue), nil
		case int8:
			return exceptedValue == int64(actualValue), nil
		case int16:
			return exceptedValue == int64(actualValue), nil
		case int32:
			return exceptedValue == int64(actualValue), nil
		case int64:
			return exceptedValue == actualValue, nil
		case float32:
			if float64(exceptedValue) == float64(actualValue) {
				return true, nil
			}
			return false, nil
		case float64:
			if float64(exceptedValue) == actualValue {
				return true, nil
			}
			return false, nil
		}
		if nil == value {
			return false, ErrValueNull
		}
		return false, errType(value, "int64")
	}
}

func stringEquals(s string) func(value interface{}) (bool, error) {
	number := s
	if strings.HasSuffix(s, ".0") {
		number = strings.TrimSuffix(s, ".0")
	}
	if u64, e := strconv.ParseUint(number, 10, 64); e == nil {
		return uintEquals(u64)
	}
	if i64, e := strconv.ParseInt(number, 10, 64); e == nil {
		return intEquals(i64)
	}
	return func(value interface{}) (bool, error) {
		if bs, ok := value.([]byte); ok {
			return s == string(bs), nil
		}
		actualValue := fmt.Sprint(value)
		return s == actualValue, nil
	}
}

func boolEquals(b bool) func(value interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		if actualBool, ok := value.(bool); ok {
			return actualBool == b, nil
		}
		if actualString, ok := value.(string); ok {
			switch actualString {
			case "true", "True", "TRUE":
				return b == true, nil
			default:
				return b == false, nil
			}
		}
		if actualBytes, ok := value.([]byte); ok {
			switch string(actualBytes) {
			case "true", "True", "TRUE":
				return b == true, nil
			default:
				return b == false, nil
			}
		}
		return false, nil
	}
}

func dynamicEquals(argValue interface{}) (func(value interface{}) (bool, error), error) {
	switch v := argValue.(type) {
	case bool:
		return boolEquals(v), nil
	case int:
		return intEquals(int64(v)), nil
	case int8:
		return intEquals(int64(v)), nil
	case int16:
		return intEquals(int64(v)), nil
	case int32:
		return intEquals(int64(v)), nil
	case int64:
		return intEquals(v), nil
	case uint:
		return uintEquals(uint64(v)), nil
	case uint8:
		return uintEquals(uint64(v)), nil
	case uint16:
		return uintEquals(uint64(v)), nil
	case uint32:
		return uintEquals(uint64(v)), nil
	case uint64:
		return uintEquals(v), nil
	case float32:
		return floatEquals(float64(v)), nil
	case float64:
		return floatEquals(v), nil
	case []byte:
		return stringEquals(string(v)), nil
	case string:
		return stringEquals(v), nil
	case json.Number:
		s := v.String()
		if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
			return uintEquals(u64), nil
		}
		if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
			return intEquals(i64), nil
		}

		if f64, e := strconv.ParseFloat(s, 64); e == nil {
			return floatEquals(f64), nil
		}
		return nil, errType(argValue, "number")
	case *json.Number:
		s := v.String()
		if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
			return uintEquals(u64), nil
		}
		if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
			return intEquals(i64), nil
		}
		if f64, e := strconv.ParseFloat(s, 64); e == nil {
			return floatEquals(f64), nil
		}
		return nil, errType(argValue, "number")
	}
	if nil == argValue {
		return nil, ErrValueNull
	}
	return nil, errType(argValue, "int64")
}

func init() {
	AddCheckFunc("=", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := dynamicEquals(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("=", "", value)
			}
			return r, nil
		}), nil
	}))

	AddCheckFunc("!=", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := dynamicEquals(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("!=", "", value)
			}
			return !r, nil
		}), nil
	}))
}
