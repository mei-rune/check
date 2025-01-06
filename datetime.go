package check

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

func init() {
	AddCheckFunc(">", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue time.Time
		switch aValue := argValue.(type) {
		case string:
			if containNowFunc(aValue) {
				readTime, err := ParseTime(aValue)
				if err != nil {
					return nil, err
				}

				return CheckFunc(func(value interface{}) (bool, error) {
					exceptedValue := readTime()
					// if err != nil {
					// 	return false, ErrArgumentType(">", "datetime", argValue)
					// }
					switch actualValue := value.(type) {
					case string:
						aValue, err := toTime(actualValue)
						if err != nil {
							return false, ErrActualType(">", "datetime", value)
						}
						return aValue.After(exceptedValue), nil
					case time.Time:
						return actualValue.After(exceptedValue), nil
					case *time.Time:
						if actualValue == nil {
							return false, nil
						}
						return actualValue.After(exceptedValue), nil
					}
					return false, ErrActualType(">", "datetime", value)
				}), nil
			}
			tt, err := toTime(aValue)
			if err != nil {
				return nil, ErrArgumentType(">", "datetime", argValue)
			}
			exceptedValue = tt
		case time.Time:
			exceptedValue = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType(">", "datetime", argValue)
			}
			exceptedValue = *aValue
		default:
			return nil, ErrArgumentType(">", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				aValue, err := toTime(actualValue)
				if err != nil {
					return false, ErrActualType(">", "datetime", value)
				}
				return aValue.After(exceptedValue), nil
			case time.Time:
				return actualValue.After(exceptedValue), nil
			case *time.Time:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.After(exceptedValue), nil
			}
			return false, ErrActualType(">", "datetime", value)
		}), nil
	}))

	AddCheckFunc(">=", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue time.Time
		switch aValue := argValue.(type) {
		case string:
			if containNowFunc(aValue) {
				readTime, err := ParseTime(aValue)
				if err != nil {
					return nil, err
				}
				return CheckFunc(func(value interface{}) (bool, error) {
					exceptedValue := readTime()
					// if err != nil {
					// 	return false, ErrArgumentType(">=", "datetime", argValue)
					// }
					switch actualValue := value.(type) {
					case string:
						aValue, err := toTime(actualValue)
						if err != nil {
							return false, ErrActualType(">=", "datetime", value)
						}
						return aValue.After(exceptedValue) || aValue.Equal(exceptedValue), nil
					case time.Time:
						return actualValue.After(exceptedValue), nil
					case *time.Time:
						if actualValue == nil {
							return false, nil
						}
						return actualValue.After(exceptedValue) || actualValue.Equal(exceptedValue), nil
					}
					return false, ErrActualType(">=", "datetime", value)
				}), nil
			}

			tt, err := toTime(aValue)
			if err != nil {
				return nil, ErrArgumentType(">=", "datetime", argValue)
			}
			exceptedValue = tt
		case time.Time:
			exceptedValue = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType(">=", "datetime", argValue)
			}
			exceptedValue = *aValue
		default:
			return nil, ErrArgumentType(">=", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				aValue, err := toTime(actualValue)
				if err != nil {
					return false, ErrActualType(">=", "datetime", value)
				}
				return aValue.After(exceptedValue) || aValue.Equal(exceptedValue), nil
			case time.Time:
				return actualValue.After(exceptedValue), nil
			case *time.Time:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.After(exceptedValue) || actualValue.Equal(exceptedValue), nil
			}
			return false, ErrActualType(">=", "datetime", value)
		}), nil
	}))

	AddCheckFunc("<", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue time.Time
		switch aValue := argValue.(type) {
		case string:
			if containNowFunc(aValue) {
				readTime, err := ParseTime(aValue)
				if err != nil {
					return nil, err
				}
				return CheckFunc(func(value interface{}) (bool, error) {
					exceptedValue := readTime()
					// if err != nil {
					// 	return false, ErrArgumentType(">", "datetime", argValue)
					// }
					switch actualValue := value.(type) {
					case string:
						aValue, err := toTime(actualValue)
						if err != nil {
							return false, ErrActualType("<", "datetime", value)
						}
						return aValue.Before(exceptedValue), nil
					case time.Time:
						return actualValue.Before(exceptedValue), nil
					case *time.Time:
						if actualValue == nil {
							return false, nil
						}
						return actualValue.Before(exceptedValue), nil
					}
					return false, ErrActualType("<", "datetime", value)
				}), nil
			}

			tt, err := toTime(aValue)
			if err != nil {
				return nil, ErrArgumentType("<", "datetime", argValue)
			}
			exceptedValue = tt
		case time.Time:
			exceptedValue = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType("<", "datetime", argValue)
			}
			exceptedValue = *aValue
		default:
			return nil, ErrArgumentType("<", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				aValue, err := toTime(actualValue)
				if err != nil {
					return false, ErrActualType("<", "datetime", value)
				}
				return aValue.Before(exceptedValue), nil
			case time.Time:
				return actualValue.Before(exceptedValue), nil
			case *time.Time:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.Before(exceptedValue), nil
			}
			return false, ErrActualType("<", "datetime", value)
		}), nil
	}))

	AddCheckFunc("<=", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue time.Time
		switch aValue := argValue.(type) {
		case string:
			if containNowFunc(aValue) {
				readTime, err := ParseTime(aValue)
				if err != nil {
					return nil, err
				}
				return CheckFunc(func(value interface{}) (bool, error) {
					exceptedValue := readTime()
					// if err != nil {
					// 	return false, ErrArgumentType("<=", "datetime", argValue)
					// }
					switch actualValue := value.(type) {
					case string:
						aValue, err := toTime(actualValue)
						if err != nil {
							return false, ErrActualType("<=", "datetime", value)
						}
						return aValue.Before(exceptedValue) || aValue.Equal(exceptedValue), nil
					case time.Time:
						return actualValue.Before(exceptedValue), nil
					case *time.Time:
						if actualValue == nil {
							return false, nil
						}
						return actualValue.Before(exceptedValue) || actualValue.Equal(exceptedValue), nil
					}
					return false, ErrActualType("<=", "datetime", value)
				}), nil
			}

			tt, err := toTime(aValue)
			if err != nil {
				return nil, ErrArgumentType("<=", "datetime", argValue)
			}
			exceptedValue = tt
		case time.Time:
			exceptedValue = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType("<=", "datetime", argValue)
			}
			exceptedValue = *aValue
		default:
			return nil, ErrArgumentType("<=", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				aValue, err := toTime(actualValue)
				if err != nil {
					return false, ErrActualType("<", "datetime", value)
				}
				return aValue.Before(exceptedValue) || aValue.Equal(exceptedValue), nil
			case time.Time:
				return actualValue.Before(exceptedValue), nil
			case *time.Time:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.Before(exceptedValue) || actualValue.Equal(exceptedValue), nil
			}
			return false, ErrActualType("<=", "datetime", value)
		}), nil
	}))

	AddCheckFunc("=", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue string
		var exceptedTime time.Time

		switch aValue := argValue.(type) {
		case string:
			exceptedValue = aValue
			if tt, err := toTime(aValue); err != nil {
				return nil, ErrArgumentType("=", "datetime", argValue)
			} else {
				exceptedTime = tt
			}
		case time.Time:
			exceptedValue = aValue.Format(time.RFC3339Nano)
			exceptedTime = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType("=", "datetime", argValue)
			}
			exceptedValue = aValue.Format(time.RFC3339Nano)
			exceptedTime = *aValue
		default:
			return nil, ErrArgumentType("=", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				if actualValue == exceptedValue {
					return true, nil
				}

				aValue, err := toTime(actualValue)
				if err != nil {
					return false, nil
				}
				return aValue.Equal(exceptedTime), nil
			case time.Time:
				return actualValue.Equal(exceptedTime), nil
			case *time.Time:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.Equal(exceptedTime), nil
			}
			return false, ErrActualType("=", "datetime", value)
		}), nil
	}))

	AddCheckFunc("!=", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedValue string
		var exceptedTime time.Time

		switch aValue := argValue.(type) {
		case string:
			exceptedValue = aValue
			if tt, err := toTime(aValue); err != nil {
				return nil, ErrArgumentType("=", "datetime", argValue)
			} else {
				exceptedTime = tt
			}
		case time.Time:
			exceptedValue = aValue.Format(time.RFC3339Nano)
			exceptedTime = aValue
		case *time.Time:
			if aValue == nil {
				return nil, ErrArgumentType("!=", "datetime", argValue)
			}
			exceptedValue = aValue.Format(time.RFC3339Nano)
			exceptedTime = *aValue
		default:
			return nil, ErrArgumentType("!=", "datetime", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				if actualValue == exceptedValue {
					return false, nil
				}

				aValue, err := toTime(actualValue)
				if err != nil {
					return true, nil
				}
				return !aValue.Equal(exceptedTime), nil
			case time.Time:
				return !actualValue.Equal(exceptedTime), nil
			case *time.Time:
				if actualValue == nil {
					return true, nil
				}
				return !actualValue.Equal(exceptedTime), nil
			}
			return false, ErrActualType("!=", "datetime", value)
		}), nil
	}))

	AddCheckFunc("in", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedArray, exceptedTimes, err = toTimeArray("in", argValue)
		if err != nil {
			return nil, err
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				for idx := range exceptedArray {
					if actualValue == exceptedArray[idx] {
						return true, nil
					}
				}
				return false, nil
			case time.Time:
				for idx := range exceptedTimes {
					if actualValue.Equal(exceptedTimes[idx]) {
						return true, nil
					}
				}
				return false, nil
			case *time.Time:
				if actualValue == nil {
					return true, nil
				}
				for idx := range exceptedTimes {
					if actualValue.Equal(exceptedTimes[idx]) {
						return true, nil
					}
				}
				return false, nil
			}
			return false, ErrActualType("in", "datetime", value)
		}), nil
	}))

	AddCheckFunc("nin", "datetime", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		var exceptedArray, exceptedTimes, err = toTimeArray("nin", argValue)
		if err != nil {
			return nil, err
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				found := false
				for idx := range exceptedArray {
					if actualValue == exceptedArray[idx] {
						found = true
						break
					}
				}
				if found {
					return false, nil
				}

				aTime, err := toTime(actualValue)
				if err != nil {
					return false, ErrActualType("nin", "datetime", value)
				}

				for idx := range exceptedTimes {
					if aTime.Equal(exceptedTimes[idx]) {
						found = true
						break
					}
				}
				return !found, nil
			case time.Time:
				found := false
				for idx := range exceptedTimes {
					if actualValue.Equal(exceptedTimes[idx]) {
						found = true
						break
					}
				}
				return !found, nil
			case *time.Time:
				if actualValue == nil {
					return true, nil
				}
				found := false
				for idx := range exceptedTimes {
					if actualValue.Equal(exceptedTimes[idx]) {
						found = true
						break
					}
				}
				return !found, nil
			}

			return false, ErrActualType("nin", "datetime", value)
		}), nil
	}))
}

func containNowFunc(s string) bool {
	return strings.Contains(s, "now(")
}

func ParseTime(s string) (func() time.Time, error) {
	s = strings.TrimSpace(s)
	t, err := toTime(s)
	if err == nil {
		return func() time.Time { return t }, nil
	}

	if !strings.HasPrefix(s, "now(") {
		return nil, errors.New("'" + s + "' is invalid, format must is [now()+-]xxx.")
	}

	s = strings.TrimPrefix(s, "now(")
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, ")") {
		if strings.Contains(s, ")") {
			return nil, errors.New("'" + s + "' is invalid, now() is none parameters.")
		}
		return nil, errors.New("'" + s + "' is invalid, parentheses is missing.")
	}

	s = strings.TrimPrefix(s, ")")
	s = strings.TrimSpace(s)
	if s == "" {
		return func() time.Time { return time.Now() }, nil
	}
	c := s[0]
	if c != '+' && c != '-' {
		return nil, errors.New("'" + s + "' is invalid, operator must is minus or plus.")
	}

	s = strings.TrimSpace(s[1:])
	if s == "" {
		return func() time.Time { return time.Now() }, nil
	}

	interval, e := time.ParseDuration(s)
	if nil != e {
		return nil, errors.New("'" + s + "' is invalid, " + e.Error())
	}

	if c == '+' {
		return func() time.Time { return time.Now().Add(interval) }, nil
	}
	interval = -interval
	return func() time.Time { return time.Now().Add(interval) }, nil
}

func toTime(v interface{}) (time.Time, error) {
	if t, ok := v.(time.Time); ok {
		return t, nil
	}

	s, ok := v.(string)
	if !ok {
		return time.Time{}, errType(v, "Time")
	}

	for _, layout := range []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05Z07:00",
		"2006-01-02 15:04:05",
		"2006-01-02",
	} {
		m, e := time.ParseInLocation(layout, s, time.Local)
		if nil == e {
			return m, nil
		}
	}

	return time.Time{}, errType(v, "Time")
}

var zeroTimePtr = (*time.Time)(nil)

func toTimeArray(op string, argValue interface{}) ([]string, []time.Time, error) {
	var exceptedArray []string
	var exceptedTimes []time.Time

	switch a := argValue.(type) {
	case []time.Time:
		exceptedTimes = a
		for _, t := range exceptedTimes {
			exceptedArray = append(exceptedArray, t.Format(time.RFC3339Nano))
		}
		return exceptedArray, exceptedTimes, nil
	case string:
		svalue, ok := argValue.(string)
		if !ok {
			return nil, nil, ErrArgumentType(op, "datetimeArray", argValue)
		}
		for _, s := range strings.Split(svalue, ",") {
			if s == "" {
				continue
			}
			exceptedArray = append(exceptedArray, s)
			t, err := toTime(s)
			if err != nil {
				return nil, nil, ErrArgumentType(op, "datetimeArray", argValue)
			}

			exceptedTimes = append(exceptedTimes, t)
		}
		return exceptedArray, exceptedTimes, nil
	}

	rv := reflect.ValueOf(argValue)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		return nil, nil, ErrArgumentType(op, "datetimeArray", argValue)
	}
	aLen := rv.Len()

	for i := 0; i < aLen; i++ {
		a := rv.Index(i).Interface()
		if a == nil || zeroTimePtr == a {
			continue
		}
		if s, ok := a.(string); ok {
			exceptedArray = append(exceptedArray, s)
		}
		t, err := toTime(a)
		if err != nil {
			return nil, nil, ErrArgumentType(op, "datetimeArray", argValue)
		}
		exceptedTimes = append(exceptedTimes, t)
		exceptedArray = append(exceptedArray, t.Format(time.RFC3339Nano))
	}
	return exceptedArray, exceptedTimes, nil
}
