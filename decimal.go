package check

func init() {
	TypeAlias["decimal"] = "integer"
	// AddCheckFunc(">", "decimal", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		return nil, ErrArgumentType(">", "decimal", argValue)
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType(">", "decimal", value)
	// 		}
	// 		return r < 0, nil
	// 	}), nil
	// }))

	// AddCheckFunc(">=", "decimal", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		return nil, ErrArgumentType(">=", "decimal", argValue)
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType(">=", "decimal", value)
	// 		}
	// 		return r <= 0, nil
	// 	}), nil
	// }))

	// AddCheckFunc("<", "decimal", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		return nil, ErrArgumentType("<", "decimal", argValue)
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType("<", "decimal", value)
	// 		}
	// 		return r > 0, nil
	// 	}), nil
	// }))

	// AddCheckFunc("<=", "decimal", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
	// 	cmp, err := numberCheck(argValue)
	// 	if err != nil {
	// 		return nil, ErrArgumentType("<=", "decimal", argValue)
	// 	}
	// 	return CheckFunc(func(value interface{}) (bool, error) {
	// 		r, err := cmp(value)
	// 		if err != nil {
	// 			return false, ErrActualType("<=", "decimal", value)
	// 		}
	// 		return r >= 0, nil
	// 	}), nil
	// }))
}
