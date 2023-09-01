package check

import (
	"bytes"
	"net"
	"strings"
)

func init() {
	AddCheckFunc("=", "ipAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
			exceptedArray := splitStrings(s, true, true)
			exceptedAddresses, err := toIPAddresses(exceptedArray)
			if err != nil {
				return nil, ErrArgumentType("=", "ipAddresses", argValue)
			}
			return InIPAddressCheck(exceptedArray, exceptedAddresses)
		}

		exceptedValue, err := toIPString(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "ipAddress", argValue)
		}
		exceptedAddr := net.ParseIP(exceptedValue)

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				return actualValue == exceptedValue, nil
			case net.IP:
				return actualValue.Equal(exceptedAddr), nil
			case *net.IP:
				if actualValue == nil {
					return false, nil
				}
				return actualValue.Equal(exceptedAddr), nil
			}
			return false, ErrActualType("=", "ipAddress", value)
		}), nil
	}))

	AddCheckFunc("!=", "ipAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
			exceptedArray := splitStrings(s, true, true)
			exceptedAddresses, err := toIPAddresses(exceptedArray)
			if err != nil {
				return nil, ErrArgumentType("!=", "ipAddresses", argValue)
			}
			return NotInIPAddressCheck(exceptedArray, exceptedAddresses)
		}

		exceptedValue, err := toIPString(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "ipAddress", argValue)
		}
		exceptedAddr := net.ParseIP(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				return actualValue != exceptedValue, nil
			case net.IP:
				return !actualValue.Equal(exceptedAddr), nil
			case *net.IP:
				if actualValue == nil {
					return true, nil
				}
				return !actualValue.Equal(exceptedAddr), nil
			}
			return false, ErrActualType("!=", "ipAddress", value)
		}), nil
	}))

	AddCheckFunc("in", "ipAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toIPStrings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "ipAddresses", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
		}
		exceptedAddresses, err := toIPAddresses(exceptedArray)
		if err != nil {
			return nil, ErrArgumentType("in", "ipAddresses", argValue)
		}
		return InIPAddressCheck(exceptedArray, exceptedAddresses)
	}))

	AddCheckFunc("nin", "ipAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toIPStrings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("nin", "ipAddresses", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
		}

		exceptedAddresses, err := toIPAddresses(exceptedArray)
		if err != nil {
			return nil, ErrArgumentType("nin", "ipAddresses", argValue)
		}
		return NotInIPAddressCheck(exceptedArray, exceptedAddresses)
	}))

	AddCheckFunc("=", "physicalAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
			exceptedArray := splitStrings(s, true, true)
			exceptedAddresses, err := parseMacStrings(argValue, exceptedArray)
			if err != nil {
				return nil, ErrArgumentType("=", "physicalAddress", argValue)
			}
			return InMacAddressCheck(exceptedArray, exceptedAddresses)
		}

		exceptedValue, err := toMacString(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "physicalAddress", argValue)
		}

		exceptedAddr, err := net.ParseMAC(exceptedValue)
		if err != nil {
			return nil, ErrArgumentType("=", "physicalAddress", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				return actualValue == exceptedValue, nil
			case net.HardwareAddr:
				return bytes.Equal(actualValue, exceptedAddr), nil
			case *net.HardwareAddr:
				if actualValue == nil {
					return false, nil
				}
				return bytes.Equal(*actualValue, exceptedAddr), nil
			}
			return false, ErrActualType("=", "physicalAddress", value)
		}), nil
	}))

	AddCheckFunc("!=", "physicalAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		if s, ok := argValue.(string); ok && strings.Contains(s, ",") {
			exceptedArray := splitStrings(s, true, true)
			exceptedAddresses, err := parseMacStrings(argValue, exceptedArray)
			if err != nil {
				return nil, ErrArgumentType("!=", "physicalAddress", argValue)
			}
			return NotInMacAddressCheck(exceptedArray, exceptedAddresses)
		}

		exceptedValue, err := toMacString(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "physicalAddress", argValue)
		}

		exceptedAddr, err := net.ParseMAC(exceptedValue)
		if err != nil {
			return nil, ErrArgumentType("=", "physicalAddress", argValue)
		}

		return CheckFunc(func(value interface{}) (bool, error) {
			switch actualValue := value.(type) {
			case string:
				return actualValue != exceptedValue, nil
			case net.HardwareAddr:
				return !bytes.Equal(actualValue, exceptedAddr), nil
			case *net.HardwareAddr:
				if actualValue == nil {
					return true, nil
				}
				return !bytes.Equal(*actualValue, exceptedAddr), nil
			}
			return false, ErrActualType("!=", "physicalAddress", value)
		}), nil
	}))

	AddCheckFunc("in", "physicalAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toMacStrings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "physicalAddresses", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
		}

		exceptedAddresses, err := toHardwareAddresses(exceptedArray)
		if err != nil {
			return nil, ErrArgumentType("in", "physicalAddresses", argValue)
		}
		return InMacAddressCheck(exceptedArray, exceptedAddresses)
	}))

	AddCheckFunc("nin", "physicalAddress", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := toMacStrings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "physicalAddresses", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
		}

		exceptedAddresses, err := toHardwareAddresses(exceptedArray)
		if err != nil {
			return nil, ErrArgumentType("nin", "physicalAddresses", argValue)
		}
		return NotInMacAddressCheck(exceptedArray, exceptedAddresses)
	}))
}

func InIPAddressCheck(exceptedArray []string, exceptedAddresses []net.IP) (Checker, error) {
	return CheckFunc(func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case string:
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return true, nil
				}
			}
			return false, nil
		case net.IP:
			for idx := range exceptedAddresses {
				if actualValue.Equal(exceptedAddresses[idx]) {
					return true, nil
				}
			}
			return false, nil
		case *net.IP:
			if actualValue == nil {
				return true, nil
			}
			for idx := range exceptedAddresses {
				if actualValue.Equal(exceptedAddresses[idx]) {
					return true, nil
				}
			}
			return false, nil
		}
		return false, ErrActualType("in", "ipAddress", value)
	}), nil
}

func NotInIPAddressCheck(exceptedArray []string, exceptedAddresses []net.IP) (Checker, error) {
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
			return !found, nil
		case net.IP:
			found := false
			for idx := range exceptedAddresses {
				if actualValue.Equal(exceptedAddresses[idx]) {
					found = true
					break
				}
			}
			return !found, nil
		case *net.IP:
			if actualValue == nil {
				return true, nil
			}
			found := false
			for idx := range exceptedAddresses {
				if actualValue.Equal(exceptedAddresses[idx]) {
					found = true
					break
				}
			}
			return !found, nil
		}
		return false, ErrActualType("nin", "ipAddress", value)
	}), nil
}

func InMacAddressCheck(exceptedArray []string, exceptedAddresses []net.HardwareAddr) (Checker, error) {
	return CheckFunc(func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case string:
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return true, nil
				}
			}
			return false, nil
		case net.HardwareAddr:
			for idx := range exceptedAddresses {
				if bytes.Equal(actualValue, exceptedAddresses[idx]) {
					return true, nil
				}
			}
			return false, nil
		case *net.HardwareAddr:
			if actualValue == nil {
				return true, nil
			}
			for idx := range exceptedAddresses {
				if bytes.Equal(*actualValue, exceptedAddresses[idx]) {
					return true, nil
				}
			}
			return false, nil
		}
		return false, ErrActualType("in", "physicalAddress", value)
	}), nil
}

func NotInMacAddressCheck(exceptedArray []string, exceptedAddresses []net.HardwareAddr) (Checker, error) {
	return CheckFunc(func(value interface{}) (bool, error) {
		switch actualValue := value.(type) {
		case string:
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return false, nil
				}
			}
			return true, nil
		case net.HardwareAddr:
			for idx := range exceptedAddresses {
				if bytes.Equal(actualValue, exceptedAddresses[idx]) {
					return false, nil
				}
			}
			return true, nil
		case *net.HardwareAddr:
			if actualValue == nil {
				return true, nil
			}
			for idx := range exceptedAddresses {
				if bytes.Equal(*actualValue, exceptedAddresses[idx]) {
					return false, nil
				}
			}
			return true, nil
		}

		return false, ErrActualType("nin", "physicalAddress", value)
	}), nil
}

func toIPString(value interface{}) (string, error) {
	switch svalue := value.(type) {
	case string:
		return svalue, nil
	case net.IP:
		return svalue.String(), nil
	case *net.IP:
		return svalue.String(), nil
	}
	return "", errType(value, "IP")
}

func toIPStrings(value interface{}) ([]string, error) {
	switch svalue := value.(type) {
	case []string:
		return svalue, nil
	case []net.IP:
		results := make([]string, len(svalue))
		for idx := range svalue {
			results[idx] = svalue[idx].String()
		}
		return results, nil
	case []*net.IP:
		results := make([]string, len(svalue))
		for idx := range svalue {
			results[idx] = svalue[idx].String()
		}
		return results, nil
	case []interface{}:
		results := make([]string, len(svalue))
		for idx := range svalue {
			addr, err := toIPString(svalue[idx])
			if err != nil {
				return nil, err
			}
			results[idx] = addr
		}
		return results, nil
	}
	return nil, errType(value, "IPStringArray")
}

func toIPAddresses(value interface{}) ([]net.IP, error) {
	switch svalue := value.(type) {
	case []string:
		results := make([]net.IP, 0, len(svalue))
		for idx := range svalue {
			if svalue[idx] == "" {
				continue
			}

			addr := net.ParseIP(svalue[idx])
			if addr == nil {
				return nil, errType(value, "IPArray")
			}
			results = append(results, addr)
		}
		return results, nil
	case []net.IP:
		return svalue, nil
	case []*net.IP:
		results := make([]net.IP, 0, len(svalue))
		for idx := range svalue {
			if svalue[idx] != nil {
				results = append(results, *svalue[idx])
			}
		}
		return results, nil
	case []interface{}:
		results := make([]net.IP, 0, len(svalue))
		for idx := range svalue {
			if svalue[idx] == nil {
				continue
			}

			switch tvalue := svalue[idx].(type) {
			case string:
				addr := net.ParseIP(tvalue)
				if addr == nil {
					return nil, errType(value, "IPArray")
				}
				results = append(results, addr)
			case net.IP:
				results = append(results, tvalue)
			case *net.IP:
				if tvalue != nil {
					results = append(results, *tvalue)
				}
			default:
				return nil, errType(value, "IPArray")
			}
		}
		return results, nil
	}
	return nil, errType(value, "IPArray")
}

func toMacString(value interface{}) (string, error) {
	switch svalue := value.(type) {
	case string:
		return svalue, nil
	case net.HardwareAddr:
		return svalue.String(), nil
	case *net.HardwareAddr:
		return svalue.String(), nil
	}
	return "", errType(value, "HardwareAddr")
}

func toMacStrings(value interface{}) ([]string, error) {
	switch svalue := value.(type) {
	case []string:
		return svalue, nil
	case []net.HardwareAddr:
		results := make([]string, len(svalue))
		for idx := range svalue {
			results[idx] = svalue[idx].String()
		}
		return results, nil
	case []*net.HardwareAddr:
		results := make([]string, len(svalue))
		for idx := range svalue {
			results[idx] = svalue[idx].String()
		}
		return results, nil
	case []interface{}:
		results := make([]string, len(svalue))
		for idx := range svalue {
			addr, err := toMacString(svalue[idx])
			if err != nil {
				return nil, err
			}
			results[idx] = addr
		}
		return results, nil
	}
	return nil, errType(value, "MacStringArray")
}

func parseMacStrings(value interface{}, ss []string) ([]net.HardwareAddr, error) {
	results := make([]net.HardwareAddr, 0, len(ss))
	for idx := range ss {
		if ss[idx] == "" {
			continue
		}

		addr, err := net.ParseMAC(ss[idx])
		if err != nil {
			return nil, errType(value, "MacArray")
		}
		results = append(results, addr)
	}
	return results, nil
}

func toHardwareAddresses(value interface{}) ([]net.HardwareAddr, error) {
	switch svalue := value.(type) {
	case []string:
		return parseMacStrings(value, svalue)
	case []net.HardwareAddr:
		return svalue, nil
	case []*net.HardwareAddr:
		results := make([]net.HardwareAddr, 0, len(svalue))
		for idx := range svalue {
			if svalue[idx] != nil {
				results = append(results, *svalue[idx])
			}
		}
		return results, nil
	case []interface{}:
		results := make([]net.HardwareAddr, 0, len(svalue))
		for idx := range svalue {
			if svalue[idx] == nil {
				continue
			}

			switch tvalue := svalue[idx].(type) {
			case string:
				addr, err := net.ParseMAC(tvalue)
				if err != nil {
					return nil, errType(value, "MacArray")
				}
				results = append(results, addr)
			case net.HardwareAddr:
				results = append(results, tvalue)
			case *net.HardwareAddr:
				if tvalue != nil {
					results = append(results, *tvalue)
				}
			default:
				return nil, errType(value, "MacArray")
			}
		}
		return results, nil
	}
	return nil, errType(value, "MacArray")
}
