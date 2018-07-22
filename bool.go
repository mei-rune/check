package check

func init() {
	AddCheckFunc("=", "boolean", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toBool(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "boolean", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toBool(value)
			if err != nil {
				return false, ErrActualType("=", "boolean", value)
			}
			return actualValue == exceptedValue, nil
		}), nil
	}))

	AddCheckFunc("!=", "boolean", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toBool(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "boolean", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toBool(value)
			if err != nil {
				return false, ErrActualType("!=", "boolean", value)
			}
			return actualValue != exceptedValue, nil
		}), nil
	}))
}

func toBool(value interface{}) (bool, error) {
	if b, ok := value.(bool); ok {
		return b, nil
	}
	if s, ok := value.(string); ok {
		switch s {
		case "TRUE", "True", "true", "YES", "Yes", "yes", "on", "enabled":
			return true, nil
		case "FALSE", "False", "false", "NO", "No", "no":
			return false, nil
		}
	}
	if nil == value {
		return false, ErrValueNull
	}
	return false, errType(value, "boolean")
}
