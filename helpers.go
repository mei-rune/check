package check

import (
	"encoding/json"
	"strconv"
)

// func toUint64(op string, argValue interface{}) (uint64, error) {
// 	exceptedValue, err := as.Uint64(argValue)
// 	if err == nil {
// 		return exceptedValue, nil
// 	}

// 	if s, ok := argValue.(string); ok {
// 		if strings.HasSuffix(s, ".0") {
// 			s = strings.TrimSuffix(s, ".0")
// 			exceptedValue, err := strconv.ParseUint(s, 10, 64)
// 			if err == nil {
// 				return exceptedValue, nil
// 			}
// 			return 0, ErrArgumentType(op, "unsigned integer", argValue)
// 		}
// 	}
// 	return 0, ErrArgumentType(op, "unsigned integer", argValue)
// }

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

func int64Cmp(exceptedValue int64, value interface{}) (int, error) {
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

func uint64Cmp(exceptedValue uint64, value interface{}) (int, error) {
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
		if float64(exceptedValue) == float64(actualValue) {
			return 0, nil
		}
		return -1, nil
	case float64:
		if actualValue < 0 {
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
