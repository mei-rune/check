package check

func init() {
	AddCheckFunc(">", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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

	AddCheckFunc(">=", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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

	AddCheckFunc("<", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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

	AddCheckFunc("<=", "", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
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
}
