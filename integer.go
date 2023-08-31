package check

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
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
	case time.Duration:
		return durationCheck(v), nil
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
	AddCheckFunc("=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
					ss := splitStrings(s, true, true)
					ints, err := toInt64s(argValue, true, ss)
					if err == nil {
						cmp := inIntArrayCheck(ints)
						return CheckFunc(func(value interface{}) (bool, error) {
										return cmp(value)
									}), nil
					}
					uints, err := toUint64s(argValue, true, ss)
					if err == nil {
						cmp := inUintArrayCheck(uints)
						return CheckFunc(func(value interface{}) (bool, error) {
										return cmp(value)
									}), nil
					}
					return nil, errType(argValue, "intArray")
		}
		return anyEquals(argValue)
	}))

	AddCheckFunc("!=", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
					ss := splitStrings(s, true, true)
					ints, err := toInt64s(argValue, true, ss)
					if err == nil {
						cmp := inIntArrayCheck(ints)
						return CheckFunc(func(value interface{}) (bool, error) {
										ret, err := cmp(value)
										return !ret, err
									}), nil
					}
					uints, err := toUint64s(argValue, true, ss)
					if err == nil {
						cmp := inUintArrayCheck(uints)
						return CheckFunc(func(value interface{}) (bool, error) {
										ret, err := cmp(value)
										return !ret, err
									}), nil
					}
					return nil, errType(argValue, "intArray")
		}
		return anyNotEquals(argValue)
	}))

	AddCheckFunc("in", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := inCheck(argValue, true)
		if err != nil {
			return nil, ErrArgumentType("in", "intArray", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			return cmp(value)
		}), nil
	}))

	AddCheckFunc("nin", "integer", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := inCheck(argValue, true)
		if err != nil {
			return nil, ErrArgumentType("nin", "intArray", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			ret, err := cmp(value)
			return !ret, err
		}), nil
	}))
}
