package check

import (
	"time"
)

func toDuration(v interface{}) (time.Duration, error) {
	if t, ok := v.(time.Duration); ok {
		return t, nil
	}

	if i, e := toInt64(v); nil == e {
		return time.Duration(i), nil
	}

	s, ok := v.(string)
	if !ok {
		return 0, errType(v, "duration")
	}

	return time.ParseDuration(s)
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
			return nil, ErrArgumentType(">", "duration", argValue)
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
			return nil, ErrArgumentType(">", "duration", argValue)
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
			return nil, ErrArgumentType(">", "duration", argValue)
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
			return nil, ErrArgumentType(">", "duration", argValue)
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
			return nil, ErrArgumentType(">", "duration", argValue)
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
