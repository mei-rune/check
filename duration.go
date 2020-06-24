package check

import (
	"encoding/json"
	"strconv"
	"time"
)

func toDuration(v interface{}) (time.Duration, error) {
	switch i := v.(type) {
	case time.Duration:
		return i, nil
	case uint:
		return time.Duration(i), nil
	case uint8:
		return time.Duration(i), nil
	case uint16:
		return time.Duration(i), nil
	case uint32:
		return time.Duration(i), nil
	case uint64:
		return time.Duration(i), nil
	case int:
		return time.Duration(i), nil
	case int8:
		return time.Duration(i), nil
	case int16:
		return time.Duration(i), nil
	case int32:
		return time.Duration(i), nil
	case int64:
		return time.Duration(i), nil
	case float32:
		return time.Duration(i), nil
	case float64:
		return time.Duration(i), nil
	case string:
		if i64, err := strconv.ParseInt(i, 10, 64); err == nil {
			return time.Duration(i64), nil
		}

		return time.ParseDuration(i)
	case json.Number:
		if i64, err := strconv.ParseInt(i.String(), 10, 64); err == nil {
			return time.Duration(i64), nil
		}

		return time.ParseDuration(i.String())
	case *json.Number:
		if i64, err := strconv.ParseInt(i.String(), 10, 64); err == nil {
			return time.Duration(i64), nil
		}
		return time.ParseDuration(i.String())
	case []byte:
		s := string(i)
		if i64, err := strconv.ParseInt(s, 10, 64); err == nil {
			return time.Duration(i64), nil
		}
		return time.ParseDuration(s)
	}
	return 0, errType(v, "duration")
}

func durationCheck(exceptedValue time.Duration) func(value interface{}) (int, error) {
	return func(value interface{}) (int, error) {
		if nil == value {
			return 0, ErrValueNull
		}

		actualValue, err := toDuration(value)
		if err != nil {
			return 0, errType(value, "duration")
		}

		if actualValue > exceptedValue {
			return -1, nil
		}

		if actualValue < exceptedValue {
			return 1, nil
		}

		return 0, nil
	}
}

func init() {
	AddCheckFunc(">", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType(">", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType(">", "duration", value)
			}
			return r < 0, nil
		}), nil
	}))

	AddCheckFunc(">=", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType(">=", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType(">=", "duration", value)
			}

			// fmt.Printf("2(%T) %v >= (%T) %v   = %v\r\n", argValue, argValue, value, value, r)
			return r <= 0, nil
		}), nil
	}))

	AddCheckFunc("<", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType("<", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("<", "duration", value)
			}
			return r > 0, nil
		}), nil
	}))

	AddCheckFunc("<=", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType("<=", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("<=", "duration", value)
			}
			return r >= 0, nil
		}), nil
	}))

	AddCheckFunc("=", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("=", "duration", value)
			}
			return r == 0, nil
		}), nil
	}))

	AddCheckFunc("!=", "duration", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toDuration(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "duration", argValue)
		}
		cmp := durationCheck(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			r, err := cmp(value)
			if err != nil {
				return false, ErrActualType("!=", "duration", value)
			}
			return r != 0, nil
		}), nil
	}))

	UnsupportedCheckFunc("in", "duration")
	UnsupportedCheckFunc("nin", "duration")
}
