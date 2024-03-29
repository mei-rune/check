package check

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func init() {
	AddCheckFunc("between", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := betweenCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("between", "range", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			return cmp(value)
		}), nil
	}))

	AddCheckFunc("not_between", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := betweenCheck(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_between", "range", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			ret, err := cmp(value)
			return !ret, err
		}), nil
	}))
}

func betweenCheck(value interface{}) (func(interface{}) (bool, error), error) {
	switch ss := value.(type) {
	case string:
		ssArr := strings.Split(ss, ",")
		if len(ssArr) != 2 {
			return nil, errType(value, "range")
		}
		return rangeCheck(ssArr[0], ssArr[1])
	case []string:
		if len(ss) != 2 {
			return nil, errType(value, "range")
		}
		return rangeCheck(ss[0], ss[1])
	case []int64:
		if len(ss) != 2 {
			return nil, errType(value, "int_range")
		}
		return intRangeCheck(ss[0], ss[1]), nil
	case []int:
		if len(ss) != 2 {
			return nil, errType(value, "int_range")
		}
		return intRangeCheck(int64(ss[0]), int64(ss[1])), nil
	case []int8:
		if len(ss) != 2 {
			return nil, errType(value, "int_range")
		}
		return intRangeCheck(int64(ss[0]), int64(ss[1])), nil
	case []int16:
		if len(ss) != 2 {
			return nil, errType(value, "int_range")
		}
		return intRangeCheck(int64(ss[0]), int64(ss[1])), nil
	case []int32:
		if len(ss) != 2 {
			return nil, errType(value, "int_range")
		}
		return intRangeCheck(int64(ss[0]), int64(ss[1])), nil
	case []uint64:
		if len(ss) != 2 {
			return nil, errType(value, "uint_range")
		}
		return uintRangeCheck(ss[0], ss[1]), nil
	case []uint:
		if len(ss) != 2 {
			return nil, errType(value, "uint_range")
		}
		return uintRangeCheck(uint64(ss[0]), uint64(ss[1])), nil
	case []uint8:
		if len(ss) != 2 {
			return nil, errType(value, "uint_range")
		}
		return uintRangeCheck(uint64(ss[0]), uint64(ss[1])), nil
	case []uint16:
		if len(ss) != 2 {
			return nil, errType(value, "uint_range")
		}
		return uintRangeCheck(uint64(ss[0]), uint64(ss[1])), nil
	case []uint32:
		if len(ss) != 2 {
			return nil, errType(value, "uint_range")
		}
		return uintRangeCheck(uint64(ss[0]), uint64(ss[1])), nil
	case []float32:
		if len(ss) != 2 {
			return nil, errType(value, "float_range")
		}
		return floatRangeCheck(float64(ss[0]), float64(ss[1])), nil
	case []float64:
		if len(ss) != 2 {
			return nil, errType(value, "float_range")
		}
		return floatRangeCheck(ss[0], ss[1]), nil
	case []interface{}:
		if len(ss) != 2 {
			return nil, errType(value, "range")
		}
		min64, ok := asInt64(ss[0], true)
		if ok {
			max64, ok := asInt64(ss[1], true)
			if ok {
				return intRangeCheck(min64, max64), nil
			}

			umax64, ok := asUint64(ss[1], true)
			if ok {
				if min64 < 0 {
					return floatRangeCheck(float64(min64), float64(umax64)), nil
				}

				return uintRangeCheck(uint64(min64), umax64), nil
			}
		}

		if a, ok := ss[0].(string); ok {
			if b, ok := ss[1].(string); ok {
				return rangeCheck(a, b)
			}
		}

		if a, ok := ss[0].(json.Number); ok {
			if b, ok := ss[1].(json.Number); ok {
				return rangeCheck(a.String(), b.String())
			}
		}

		if a, ok := ss[0].(time.Time); ok {
			if b, ok := ss[1].(time.Time); ok {
				return datetimeRangeCheck(a, b), nil
			}
		}
	}
	return nil, errType(value, "range")
}

func rangeCheck(a, b string) (func(interface{}) (bool, error), error) {
	if strings.HasPrefix(a, "-") || strings.HasPrefix(b, "-") {
		min, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, errType(a+","+b, "range")
		}
		max, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			return nil, errType(a+","+b, "range")
		}
		return intRangeCheck(min, max), nil
	}
	min, err := strconv.ParseUint(a, 10, 64)
	if err == nil {
		max, err := strconv.ParseUint(b, 10, 64)
		if err == nil {
			return uintRangeCheck(min, max), nil
		}

		maxf, err := strconv.ParseFloat(b, 64)
		if err == nil {
			return floatRangeCheck(float64(min), maxf), nil
		}
	} else {
		minf, err := strconv.ParseFloat(a, 64)
		if err == nil {
			maxf, err := strconv.ParseFloat(b, 64)
			if err == nil {
				return floatRangeCheck(minf, maxf), nil
			}
		}
	}

	mint, err := toTime(a)
	if err == nil {
		maxt, err := toTime(b)
		if err == nil {
			return datetimeRangeCheck(mint, maxt), nil
		}
	}

	mind, err := time.ParseDuration(a)
	if err == nil {
		maxd, err := time.ParseDuration(b)
		if err == nil {
			return durationRangeCheck(mind, maxd), nil
		}
	}

	return nil, errType(a+","+b, "range")
}

func intRangeCheck(a, b int64) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch i := value.(type) {
		case uint64:
			if b < 0 {
				return false, nil
			}
			if a < 0 {
				return i < uint64(b), nil
			}
			return uint64(a) <= i && i <= uint64(b), nil
		case uint:
			if b < 0 {
				return false, nil
			}
			if a < 0 {
				return i < uint(b), nil
			}
			return uint(a) <= i && i <= uint(b), nil
		case uint8:
			current := int64(i)
			return a <= current && current <= b, nil
		case uint16:
			current := int64(i)
			return a <= current && current <= b, nil
		case uint32:
			current := int64(i)
			return a <= current && current <= b, nil
		case int:
			current := int64(i)
			return a <= current && current <= b, nil
		case int64:
			return a <= i && i <= b, nil
		case int8:
			current := int64(i)
			return a <= current && current <= b, nil
		case int16:
			current := int64(i)
			return a <= current && current <= b, nil
		case int32:
			current := int64(i)
			return a <= current && current <= b, nil
		case float32:
			return float32(a) <= i && i <= float32(b), nil
		case float64:
			return float64(a) <= i && i <= float64(b), nil
		case json.Number:
			i64, err := i.Int64()
			if err == nil {
				return a <= i64 && i64 <= b, nil
			}

			f64, err := i.Float64()
			if err == nil {
				return float64(a) <= f64 && f64 <= float64(b), nil
			}
		case *json.Number:
			i64, err := i.Int64()
			if err == nil {
				return a <= i64 && i64 <= b, nil
			}

			f64, err := i.Float64()
			if err == nil {
				return float64(a) <= f64 && f64 <= float64(b), nil
			}
		// case []byte:
		// 	if !mustInt {
		// 		if len(i) == 0 {
		// 			return 0, false
		// 		}

		// 		u64, err := strconv.ParseUint(string(i), 10, 64)
		// 		if err == nil {
		// 			return u64, true
		// 		}
		// 	}
		case string:
			i64, err := strconv.ParseInt(i, 10, 64)
			if err == nil {
				return a <= i64 && i64 <= b, nil
			}

			f64, err := strconv.ParseFloat(i, 64)
			if err == nil {
				return float64(a) <= f64 && f64 <= float64(b), nil
			}
		}
		return false, ErrArgumentType("between", "int", value)
	}
}

func uintRangeCheck(a, b uint64) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch i := value.(type) {
		case uint64:
			return a <= i && i <= b, nil
		case uint:
			current := uint64(i)
			return a <= current && current <= b, nil
		case uint8:
			current := uint64(i)
			return a <= current && current <= b, nil
		case uint16:
			current := uint64(i)
			return a <= current && current <= b, nil
		case uint32:
			current := uint64(i)
			return a <= current && current <= b, nil
		case int:
			if i < 0 {
				return false, nil
			}
			current := uint64(i)
			return a <= current && current <= b, nil
		case int64:
			if i < 0 {
				return false, nil
			}
			current := uint64(i)
			return a <= current && current <= b, nil
		case int8:
			if i < 0 {
				return false, nil
			}
			current := uint64(i)
			return a <= current && current <= b, nil
		case int16:
			if i < 0 {
				return false, nil
			}
			current := uint64(i)
			return a <= current && current <= b, nil
		case int32:
			if i < 0 {
				return false, nil
			}
			current := uint64(i)
			return a <= current && current <= b, nil

		case float32:
			if i < 0 {
				return false, nil
			}
			return float32(a) <= i && i <= float32(b), nil
		case float64:
			if i < 0 {
				return false, nil
			}
			return float64(a) <= i && i <= float64(b), nil
		case json.Number:
			i64, err := i.Int64()
			if err == nil {
				if i64 < 0 {
					return false, nil
				}
				current := uint64(i64)
				return a <= current && current <= b, nil
			}

			f64, err := i.Float64()
			if err == nil {
				if f64 < 0 {
					return false, nil
				}
				return float64(a) <= f64 && f64 <= float64(b), nil
			}

		case *json.Number:
			i64, err := i.Int64()
			if err == nil {
				if i64 < 0 {
					return false, nil
				}
				current := uint64(i64)
				return a <= current && current <= b, nil
			}

			f64, err := i.Float64()
			if err == nil {
				if f64 < 0 {
					return false, nil
				}
				return float64(a) <= f64 && f64 <= float64(b), nil
			}
		// case []byte:
		// 	if !mustInt {
		// 		if len(i) == 0 {
		// 			return 0, false
		// 		}

		// 		u64, err := strconv.ParseUint(string(i), 10, 64)
		// 		if err == nil {
		// 			return u64, true
		// 		}
		// 	}
		case string:
			i64, err := strconv.ParseInt(i, 10, 64)
			if err == nil {
				if i64 < 0 {
					return false, nil
				}
				current := uint64(i64)
				return a <= current && current <= b, nil
			}

			f64, err := strconv.ParseFloat(i, 64)
			if err == nil {
				if f64 < 0 {
					return false, nil
				}
				return float64(a) <= f64 && f64 <= float64(b), nil
			}
		}
		return false, ErrArgumentType("between", "uint", value)
	}
}

func floatRangeCheck(a, b float64) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		current, ok := asFloat64(value)
		if !ok {
			return false, ErrArgumentType("between", "float", value)
		}

		return a <= current && current <= b, nil
	}
}

func datetimeRangeCheck(a, b time.Time) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		currentValue, err := toTime(value)
		if err != nil {
			return false, ErrArgumentType("between", "datetime", value)
		}
		return a.Before(currentValue) && a.After(currentValue), nil
	}
}

func durationRangeCheck(a, b time.Duration) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		switch v := value.(type) {
		case time.Duration:
			return a <= v && v <= b, nil
		case string:
			interval, err := time.ParseDuration(v)
			if err != nil {
				return false, ErrArgumentType("between", "duration", value)
			}

			return a <= interval && interval <= b, nil
		case json.Number:
			interval, err := time.ParseDuration(v.String())
			if err != nil {
				return false, ErrArgumentType("between", "duration", value)
			}

			return a <= interval && interval <= b, nil
		case float64:
			interval := time.Duration(v)
			return a <= interval && interval <= b, nil
		}
		return false, ErrArgumentType("between", "duration", value)
	}
}
