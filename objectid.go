package check

import "strings"

func init() {
	AddCheckFunc("=", "objectId", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toInt64(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "objectId", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toInt64(value)
			if err != nil {
				return false, ErrActualType("=", "objectId", value)
			}
			return actualValue == exceptedValue, nil
		}), nil
	}))

	AddCheckFunc("!=", "objectId", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toInt64(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "objectId", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toInt64(value)
			if err != nil {
				return false, ErrActualType("!=", "objectId", value)
			}
			return actualValue != exceptedValue, nil
		}), nil
	}))

	AddCheckFunc("in", "objectId", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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
				return false, ErrActualType("in", "objectId", value)
			}
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return true, nil
				}
			}
			return false, nil
		}), nil
	}))

	AddCheckFunc("nin", "objectId", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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
				return false, ErrActualType("nin", "objectId", value)
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
