package check

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	spew "github.com/davecgh/go-spew/spew"
)

var (
	ErrValueNull = errors.New("value is nil")
)

type Checker interface {
	Check(value interface{}) (bool, error)
}

type CheckFunc func(value interface{}) (bool, error)

func (c CheckFunc) Check(value interface{}) (bool, error) {
	return c(value)
}

type CheckFactory interface {
	Create(value interface{}) (Checker, error)
}

type CheckFactoryFunc func(value interface{}) (Checker, error)

func (f CheckFactoryFunc) Create(value interface{}) (Checker, error) {
	return f(value)
}

func notSupport(err error) CheckFactoryFunc {
	return CheckFactoryFunc(func(value interface{}) (Checker, error) {
		return nil, err
	})
}

func ErrUnsupportedFunc(op, typ string) error {
	return errors.New("'" + op + "' is unsupported for the type - " + typ)
}

func ErrArgumentValue(op string, value interface{}) error {
	return fmt.Errorf("%s is invalid for '%s'", spew.Sprint(value), op)
}

func ErrArgumentType(op, typ string, value interface{}) error {
	if typ == "" {
		typ = "dynamic"
	}
	return fmt.Errorf("(%T) %s cannot convert to %s", value, spew.Sprint(value), typ)
}

func ErrActualType(op, typ string, value interface{}) error {
	if typ == "" {
		typ = "dynamic"
	}
	return fmt.Errorf("(%T) %s cannot convert to %s", value, spew.Sprint(value), typ)
}

type ErrUnexpectedType struct {
	value interface{}
	typ   string
}

func (e *ErrUnexpectedType) Error() string {
	return fmt.Sprintf("(%T) %s cannot convert to %s", e.value, spew.Sprint(e.value), e.typ)
}

func errType(value interface{}, typ string) error {
	return &ErrUnexpectedType{value: value, typ: typ}
}

type TypedCheckFactory struct {
	Default CheckFactory
	Types   map[string]CheckFactory
}

var (
	TypeAlias = map[string]string{
		"bigInteger": "integer",
		"biginteger": "integer",
		"objectid":   "integer",
		"objectId":   "integer",
		"objectID":   "integer",
	}
	OpAlias = map[string]string{
		"<>":         "!=",
		"==":         "=",
		"eq":         "=",
		"equal":      "=",
		"equals":     "=",
		"ne":         "!=",
		"neq":        "!=",
		"not_equal":  "!=",
		"not_equals": "!=",
		"gt":         ">",
		"gte":        ">=",
		"lt":         "<",
		"lte":        "<=",
		"not_in":     "nin",

		"startsWith": "startWith",
		"startwith":  "startWith",
		"start_with": "startWith",

		"noStartsWith":     "noStartWith",
		"nostartwith":      "noStartWith",
		"donot_start_with": "noStartWith",
		"not_start_with":   "noStartWith",

		"startsWithIgnorecase":       "startWithIgnorecase",
		"startwithIgnorecase":        "startWithIgnorecase",
		"start_with_ignorecase":      "startWithIgnorecase",
		"start_with_and_ignore_case": "startWithIgnorecase",

		"endsWith": "endWith",
		"endwith":  "endWith",
		"end_with": "endWith",

		"endsWithIgnorecase":       "endWithIgnorecase",
		"endwithIgnorecase":        "endWithIgnorecase",
		"end_with_ignore_case":     "endWithIgnorecase",
		"end_with_and_ignore_case": "endWithIgnorecase",

		"noEndsWith":   "noEndWith",
		"no_end_with":  "noEndWith",
		"not_end_with": "noEndWith",
	}

	// DispatchTypes  = map[string]string{}
	CheckFactories = map[string]TypedCheckFactory{}
)

func AddCheckFunc(op, typ string, f CheckFactoryFunc) {
	byOp := CheckFactories[op]
	if byOp.Types == nil {
		byOp.Types = map[string]CheckFactory{}
	}
	if typ == "" {
		byOp.Default = f
	} else {
		byOp.Types[typ] = f
	}
	CheckFactories[op] = byOp
}

// func AddDispatchType(op, typ, realType string) {
// 	DispatchTypes[op+"-"+typ] = realType
// }

func UnsupportedCheckFunc(op, typ string) {
	AddCheckFunc(op, typ, notSupport(ErrUnsupportedFunc(op, typ)))
}

// func init() {
// 	// UnsupportedCheckFunc(">", "password")
// 	// UnsupportedCheckFunc(">=", "password")
// 	// UnsupportedCheckFunc("<", "password")
// 	// UnsupportedCheckFunc("<=", "password")
// 	// UnsupportedCheckFunc("=", "password")
// 	// UnsupportedCheckFunc("!=", "password")
// 	// UnsupportedCheckFunc("in", "password")
// 	// UnsupportedCheckFunc("nin", "password")

// 	AddDispatchType(">", "dynamic", "string")
// 	AddDispatchType(">=", "dynamic", "string")
// 	AddDispatchType("<", "dynamic", "string")
// 	AddDispatchType("<=", "dynamic", "string")
// 	AddDispatchType("=", "dynamic", "string")
// 	AddDispatchType("!=", "dynamic", "string")

// 	UnsupportedCheckFunc("in", "dynamic")
// 	UnsupportedCheckFunc("nin", "dynamic")
// }

func getCheckedType(typ string, value interface{}) string {
	if typ != "" {
		return typ
	}
	if value == nil {
		return "unknownType"
	}
	return fmt.Sprintf("%T", value)
}

func MakeChecker(typ, operator string, value interface{}) (Checker, error) {
	factories, ok := CheckFactories[operator]
	if !ok {
		alias, ok := OpAlias[operator]
		if !ok {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
		factories, ok = CheckFactories[alias]
		if !ok {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
	}

	if len(factories.Types) == 0 {
		if factories.Default == nil {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
		return factories.Default.Create(value)
	}

	if typ == "" {
		if factories.Default == nil {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
		return factories.Default.Create(value)
	}

	creator := factories.Types[typ]
	if creator != nil {
		return creator.Create(value)
	}
	if alias, ok := TypeAlias[typ]; ok {
		creator = factories.Types[alias]
		if creator != nil {
			return creator.Create(value)
		}
	}

	if factories.Default != nil {
		return factories.Default.Create(value)
	}
	return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
}

func guessType(value interface{}) string {
	switch tvalue := value.(type) {
	case int32:
		return "integer"
	case int:
		return "integer"
	case int64:
		return "integer"
	case uint32:
		return "biginteger"
	case uint:
		return "biginteger"
	case uint64:
		return "biginteger"
	case float32:
		return "decimal"
	case string:
		return "string"
	case net.IP:
		return "ipAddress"
	case *net.IP:
		return "ipAddress"
	case net.HardwareAddr:
		return "physicalAddress"
	case *net.HardwareAddr:
		return "physicalAddress"
	case time.Time:
		return "datetime"
	case *time.Time:
		return "datetime"
	case json.Number:
		if _, err := tvalue.Int64(); err == nil {
			return "integer"
		} else {
			return "decimal"
		}
	case *json.Number:
		if _, err := tvalue.Int64(); err == nil {
			return "integer"
		} else {
			return "decimal"
		}
	case []string:
		return "string"
	case []int32, []uint32, []int, []uint, []int64, []uint64:
		return "biginteger"
	case []interface{}:
		return guessType(tvalue[0])
	default:
		return ""
	}
}
