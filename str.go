package check

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	// AddCheckFunc(">", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		// exceptedValue, err := toString(argValue)
	// 		// if err != nil {
	// 		return nil, ErrArgumentType(">", "string", argValue)
	// 		// }
	// 		// return CheckFunc(func(value interface{}) (bool, error) {
	// 		// 	actualValue, err := toString(value)
	// 		// 	if err != nil {
	// 		// 		return false, ErrActualType(">", "string", value)
	// 		// 	}
	// 		// 	return actualValue > exceptedValue, nil
	// 		// }), nil
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType(">", "stringNumber", value)
	// 		}
	// 		//fmt.Printf("1(%T) %v > (%T) %v   = %v\r\n", argValue, argValue, value, value, r)
	// 		return r < 0, nil
	// 	}), nil
	// }))
	// AddCheckFunc(">=", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		// exceptedValue, err := toString(argValue)
	// 		// if err != nil {
	// 		return nil, ErrArgumentType(">=", "string", argValue)
	// 		// }
	// 		// return CheckFunc(func(value interface{}) (bool, error) {
	// 		// 	actualValue, err := toString(value)
	// 		// 	if err != nil {
	// 		// 		return false, ErrActualType(">=", "string", value)
	// 		// 	}
	// 		// 	return actualValue >= exceptedValue, nil
	// 		// }), nil
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType(">=", "stringNumber", value)
	// 		}

	// 		//fmt.Printf("2(%T) %v >= (%T) %v   = %v\r\n", argValue, argValue, value, value, r)
	// 		return r <= 0, nil
	// 	}), nil
	// }))
	// AddCheckFunc("<", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		// exceptedValue, err := toString(argValue)
	// 		// if err != nil {
	// 		return nil, ErrArgumentType("<", "string", argValue)
	// 		// }
	// 		// return CheckFunc(func(value interface{}) (bool, error) {
	// 		// 	actualValue, err := toString(value)
	// 		// 	if err != nil {
	// 		// 		return false, ErrActualType("<", "string", value)
	// 		// 	}
	// 		// 	return actualValue < exceptedValue, nil
	// 		// }), nil
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType("<", "stringNumber", value)
	// 		}
	// 		return r > 0, nil
	// 	}), nil
	// }))
	// AddCheckFunc("<=", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		// exceptedValue, err := toString(argValue)
	// 		// if err != nil {
	// 		return nil, ErrArgumentType("<=", "string", argValue)
	// 		// }
	// 		// return CheckFunc(func(value interface{}) (bool, error) {
	// 		// 	actualValue, err := toString(value)
	// 		// 	if err != nil {
	// 		// 		return false, ErrActualType("<=", "string", value)
	// 		// 	}
	// 		// 	return actualValue <= exceptedValue, nil
	// 		// }), nil
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType("<=", "stringNumber", value)
	// 		}
	// 		return r >= 0, nil
	// 	}), nil
	// }))

	strEquals := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "string", argValue)
		}
		exceptedValueWithoutZero := exceptedValue
		hasZero := strings.HasSuffix(exceptedValue, ".0")
		if hasZero {
			exceptedValueWithoutZero = strings.TrimSuffix(exceptedValue, ".0")
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("=", "string", value)
			}
			if actualValue == exceptedValue {
				return true, nil
			}
			if hasZero {
				if actualValue == exceptedValueWithoutZero {
					return true, nil
				}
			}
			return strings.TrimSuffix(actualValue, ".0") == exceptedValueWithoutZero, nil
		}), nil
	})
	AddCheckFunc("=", "string", strEquals)
	AddCheckFunc("equals", "string", strEquals)
	AddCheckFunc("equals", "", strEquals)

	strNotEquals := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "string", argValue)
		}

		exceptedValueWithoutZero := exceptedValue
		hasZero := strings.HasSuffix(exceptedValue, ".0")
		if hasZero {
			exceptedValueWithoutZero = strings.TrimSuffix(exceptedValue, ".0")
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("!=", "string", value)
			}
			if actualValue == exceptedValue {
				return false, nil
			}
			if hasZero {
				if actualValue == exceptedValueWithoutZero {
					return false, nil
				}
			}
			return strings.TrimSuffix(actualValue, ".0") != exceptedValueWithoutZero, nil
		}), nil
	})
	AddCheckFunc("!=", "string", strNotEquals)
	AddCheckFunc("not_equals", "string", strNotEquals)
	AddCheckFunc("not_equals", "", strNotEquals)

	strContains := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("contains", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("contains", "string", value)
			}
			return strings.Contains(actualValue, exceptedValue), nil
		}), nil
	})
	AddCheckFunc("contains", "string", strContains)
	AddCheckFunc("contains", "", strContains)

	strNotContains := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_contains", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_contains", "string", value)
			}
			return !strings.Contains(actualValue, exceptedValue), nil
		}), nil
	})
	AddCheckFunc("not_contains", "string", strNotContains)
	AddCheckFunc("not_contains", "", strNotContains)

	strContainsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("contains_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("contains_with_ignore_case", "string", value)
			}
			if strings.Contains(actualValue, exceptedValue) {
				return true, nil
			}
			return strings.Contains(strings.ToLower(actualValue), exceptedValue), nil
		}), nil
	})
	AddCheckFunc("contains_with_ignore_case", "string", strContainsWithIgnorecase)
	AddCheckFunc("contains_with_ignore_case", "", strContainsWithIgnorecase)

	strNotContainsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_contains_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_contains_with_ignore_case", "string", value)
			}
			if strings.Contains(actualValue, exceptedValue) {
				return false, nil
			}
			return !strings.Contains(strings.ToLower(actualValue), exceptedValue), nil
		}), nil
	})
	AddCheckFunc("not_contains_with_ignore_case", "string", strNotContainsWithIgnorecase)
	AddCheckFunc("not_contains_with_ignore_case", "", strNotContainsWithIgnorecase)

	strEqualsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("equals_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("equals_with_ignore_case", "string", value)
			}
			if actualValue == exceptedValue {
				return true, nil
			}
			return strings.ToLower(actualValue) == exceptedValue, nil
		}), nil
	})
	AddCheckFunc("equals_with_ignore_case", "string", strEqualsWithIgnorecase)
	AddCheckFunc("equals_with_ignore_case", "", strEqualsWithIgnorecase)

	strNotEqualsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_equals_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_equals_with_ignore_case", "string", value)
			}
			if actualValue == exceptedValue {
				return false, nil
			}
			return strings.ToLower(actualValue) != exceptedValue, nil
		}), nil
	})
	AddCheckFunc("not_equals_with_ignore_case", "string", strNotEqualsWithIgnorecase)
	AddCheckFunc("not_equals_with_ignore_case", "", strNotEqualsWithIgnorecase)

	AddCheckFunc("in", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := stringIn(argValue)
		if err != nil {
			return nil, err
		}
		return CheckFunc(cmp), nil
	}))

	AddCheckFunc("nin", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		cmp, err := stringIn(argValue)
		if err != nil {
			return nil, err
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			ok, err := cmp(value)
			if err != nil {
				return false, err
			}
			return !ok, nil
		}), nil
	}))

	startsWith := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("startWith", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("startWith", "string", value)
			}
			return strings.HasPrefix(actualValue, excepted), nil
		}), nil
	})

	AddCheckFunc("startWith", "string", startsWith)
	AddCheckFunc("startWith", "", startsWith)

	startsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("startWithIgnorecase", "string", argValue)
		}
		excepted = strings.ToLower(excepted)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("startWithIgnorecase", "string", value)
			}
			if len(actualValue) < len(excepted) {
				return false, nil
			}
			return strings.ToLower(actualValue[:len(excepted)]) == excepted, nil
		}), nil
	})

	AddCheckFunc("startWithIgnorecase", "string", startsWithIgnorecase)
	AddCheckFunc("startWithIgnorecase", "", startsWithIgnorecase)

	noStartsWith := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("noStartWith", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("noStartWith", "string", value)
			}
			return !strings.HasPrefix(actualValue, excepted), nil
		}), nil
	})
	AddCheckFunc("noStartWith", "string", noStartsWith)
	AddCheckFunc("noStartWith", "", noStartsWith)

	endsWith := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("endWith", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("endWith", "string", value)
			}
			return strings.HasSuffix(actualValue, excepted), nil
		}), nil
	})
	AddCheckFunc("endWith", "string", endsWith)
	AddCheckFunc("endWith", "", endsWith)

	endsWithIgnorecase := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("endWithIgnorecase", "string", argValue)
		}
		excepted = strings.ToLower(excepted)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("endWithIgnorecase", "string", value)
			}
			if len(actualValue) < len(excepted) {
				return false, nil
			}
			return strings.ToLower(actualValue[len(actualValue)-len(excepted):]) == excepted, nil
		}), nil
	})
	AddCheckFunc("endWithIgnorecase", "string", endsWithIgnorecase)
	AddCheckFunc("endWithIgnorecase", "", endsWithIgnorecase)

	noEndsWith := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("noEndWith", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("noEndWith", "string", value)
			}
			return !strings.HasSuffix(actualValue, excepted), nil
		}), nil
	})
	AddCheckFunc("noEndWith", "string", noEndsWith)
	AddCheckFunc("noEndWith", "", noEndsWith)

	strMatch := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("match", "string", argValue)
		}

		re, err := regexp.Compile(excepted)
		if err != nil {
			return nil, ErrArgumentValue("match", excepted)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("match", "string", value)
			}
			return re.MatchString(actualValue), nil
		}), nil
	})
	AddCheckFunc("match", "string", strMatch)
	AddCheckFunc("match", "", strMatch)

	strNotMatch := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		excepted, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("notmatch", "string", argValue)
		}

		re, err := regexp.Compile(excepted)
		if err != nil {
			return nil, ErrArgumentValue("notmatch", excepted)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("notmatch", "string", value)
			}
			return !re.MatchString(actualValue), nil
		}), nil
	})
	AddCheckFunc("notmatch", "string", strNotMatch)
	AddCheckFunc("notmatch", "", strNotMatch)
}

func toString(value interface{}) (string, error) {
	if nil == value {
		return "", ErrValueNull
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case json.Number:
		s := v.String()
		return s, nil
	case *json.Number:
		return v.String(), nil
	case []byte:
		if v == nil {
			return "", nil
		}
		return string(v), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case bool:
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	}

	if st, ok := value.(fmt.Stringer); ok {
		return st.String(), nil
	}
	return "", errType(value, "string")
}

func toStrings(argValue interface{}) ([]string, error) {
	if argValue == nil {
		return nil, ErrValueNull
	}
	if ss, ok := argValue.([]string); ok {
		return ss, nil
	}

	rv := reflect.ValueOf(argValue)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		if s, ok := argValue.(string); ok {
			return strings.Split(s, ","), nil
		}
		return nil, errType(argValue, "string array")
	}
	aLen := rv.Len()
	results := make([]string, 0, aLen)
	for i := 0; i < aLen; i++ {
		v := rv.Index(i)

		if !v.IsValid() {
			continue
		}

		if v.Type().Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
		}

		s, e := toString(v.Interface())
		if e != nil {
			return nil, e
		}
		results = append(results, s)
	}
	return results, nil
}

func intstringCheckEqual(exceptedArray []string, ints []int64, s string) (bool, error) {
	for idx := range exceptedArray {
		if s == exceptedArray[idx] {
			return true, nil
		}
	}
	i64, ok := ParseInt64WithStr(s)
	if ok {
		for _, ex := range ints {
			if ex == i64 {
				return true, nil
			}
		}
	}
	return false, nil
}

func uintstringCheckEqual(exceptedArray []string, uints []uint64, s string) (bool, error) {
	for idx := range exceptedArray {
		if s == exceptedArray[idx] {
			return true, nil
		}
	}
	u64, ok := ParseUint64WithStr(s)
	if ok {
		for _, ex := range uints {
			if ex == u64 {
				return true, nil
			}
		}
	}
	return false, nil
}

func stringIn(argValue interface{}) (func(interface{}) (bool, error), error) {
	exceptedArray, err := toStrings(argValue)
	if err != nil {
		svalue, ok := argValue.(string)
		if !ok {
			return nil, ErrArgumentType("in", "stringArray", argValue)
		}
		exceptedArray = strings.Split(svalue, ",")
		exceptedArray = append(exceptedArray, svalue)
	}

	ints, err := toInt64s(argValue, true, exceptedArray)
	if err == nil {
		return func(value interface{}) (bool, error) {
			switch s := value.(type) {
			case string:
				return intstringCheckEqual(exceptedArray, ints, s)
			case json.Number:
				return intstringCheckEqual(exceptedArray, ints, s.String())
			case *json.Number:
				return intstringCheckEqual(exceptedArray, ints, s.String())
			case []byte:
				return intstringCheckEqual(exceptedArray, ints, string(s))
			case float64:
				for _, ex := range ints {
					if float64(ex) == s {
						return true, nil
					}
				}
				return false, nil
			case float32:
				for _, ex := range ints {
					if float32(ex) == s {
						return true, nil
					}
				}
				return false, nil
			}
			i64, ok := asInt64(value, true)
			if ok {
				for _, ex := range ints {
					if ex == i64 {
						return true, nil
					}
				}
				return false, nil
			}
			if st, ok := value.(fmt.Stringer); ok {
				s := st.String()
				for idx := range exceptedArray {
					if s == exceptedArray[idx] {
						return true, nil
					}
				}
				return false, nil
			}
			return false, ErrActualType("in", "string", value)
		}, nil
	}

	uints, err := toUint64s(argValue, true, exceptedArray)
	if err == nil {
		return func(value interface{}) (bool, error) {
			switch s := value.(type) {
			case string:
				return uintstringCheckEqual(exceptedArray, uints, s)
			case json.Number:
				return uintstringCheckEqual(exceptedArray, uints, s.String())
			case *json.Number:
				return uintstringCheckEqual(exceptedArray, uints, s.String())
			case []byte:
				return intstringCheckEqual(exceptedArray, ints, string(s))
			case float64:
				for _, ex := range ints {
					if float64(ex) == s {
						return true, nil
					}
				}
				return false, nil
			case float32:
				for _, ex := range ints {
					if float32(ex) == s {
						return true, nil
					}
				}
				return false, nil
			}
			u64, ok := asUint64(value, true)
			if ok {
				for _, ex := range uints {
					if ex == u64 {
						return true, nil
					}
				}
				return false, nil
			}
			if st, ok := value.(fmt.Stringer); ok {
				s := st.String()
				for idx := range exceptedArray {
					if s == exceptedArray[idx] {
						return true, nil
					}
				}
				return false, nil
			}
			return false, ErrActualType("in", "string", value)
		}, nil
	}

	return func(value interface{}) (bool, error) {
		actualValue, err := toString(value)
		if err != nil {
			return false, ErrActualType("in", "string", value)
		}
		for idx := range exceptedArray {
			if actualValue == exceptedArray[idx] {
				return true, nil
			}
		}

		return false, nil
	}, nil
}
