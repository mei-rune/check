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
	return fmt.Errorf("%s cannot convert to %s", spew.Sprint(value), typ)
}

func ErrActualType(op, typ string, value interface{}) error {
	return fmt.Errorf("%s cannot convert to %s", spew.Sprint(value), typ)
}

type ErrUnexpectedType struct {
	value interface{}
	typ   string
}

func (e *ErrUnexpectedType) Error() string {
	return fmt.Sprintf("%s cannot convert to %s", spew.Sprint(e.value), e.typ)
}

func errType(value interface{}, typ string) error {
	return &ErrUnexpectedType{value: value, typ: typ}
}

var (
	TypeAlias = map[string]string{
		"bigInteger": "integer",
		"biginteger": "integer",
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
	DispatchTypes  = map[string]string{}
	CheckFactories = map[string]map[string]CheckFactory{
		">":   map[string]CheckFactory{},
		">=":  map[string]CheckFactory{},
		"<":   map[string]CheckFactory{},
		"<=":  map[string]CheckFactory{},
		"=":   map[string]CheckFactory{},
		"!=":  map[string]CheckFactory{},
		"in":  map[string]CheckFactory{},
		"nin": map[string]CheckFactory{},
	}
)

func AddCheckFunc(op, typ string, f CheckFactoryFunc) {
	byOp := CheckFactories[op]
	if byOp == nil {
		byOp = map[string]CheckFactory{}
		CheckFactories[op] = byOp
	}
	byOp[typ] = f
}

func AddDispatchType(op, typ, realType string) {
	DispatchTypes[op+"-"+typ] = realType
}

func UnsupportedCheckFunc(op, typ string) {
	AddCheckFunc(op, typ, notSupport(ErrUnsupportedFunc(op, typ)))
}

func init() {
	// UnsupportedCheckFunc(">", "password")
	// UnsupportedCheckFunc(">=", "password")
	// UnsupportedCheckFunc("<", "password")
	// UnsupportedCheckFunc("<=", "password")
	// UnsupportedCheckFunc("=", "password")
	// UnsupportedCheckFunc("!=", "password")
	// UnsupportedCheckFunc("in", "password")
	// UnsupportedCheckFunc("nin", "password")

	AddDispatchType(">", "dynamic", "string")
	AddDispatchType(">=", "dynamic", "string")
	AddDispatchType("<", "dynamic", "string")
	AddDispatchType("<=", "dynamic", "string")
	AddDispatchType("=", "dynamic", "string")
	AddDispatchType("!=", "dynamic", "string")

	UnsupportedCheckFunc("in", "dynamic")
	UnsupportedCheckFunc("nin", "dynamic")
}

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
	factories := CheckFactories[operator]
	if factories == nil {
		alias, ok := OpAlias[operator]
		if !ok {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
		factories = CheckFactories[alias]
		if factories == nil {
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
	}

	if typ == "" {
		typ = guessType(value)
	}

	creator := factories[typ]
	if creator != nil {
		return creator.Create(value)
	}
	if alias, ok := TypeAlias[typ]; ok {
		creator = factories[alias]
		if creator != nil {
			return creator.Create(value)
		}
	}

	opAlias, ok := OpAlias[operator]
	if !ok {
		opAlias = operator
	}
	if alias, ok := TypeAlias[typ]; ok {
		creator = factories[DispatchTypes[opAlias+"-"+alias]]
	} else {
		creator = factories[DispatchTypes[opAlias+"-"+typ]]
	}
	if creator != nil {
		return creator.Create(value)
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
