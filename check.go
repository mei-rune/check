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
	}
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

	UnsupportedCheckFunc(">", "dynamic")
	UnsupportedCheckFunc(">=", "dynamic")
	UnsupportedCheckFunc("<", "dynamic")
	UnsupportedCheckFunc("<=", "dynamic")
	UnsupportedCheckFunc("=", "dynamic")
	UnsupportedCheckFunc("!=", "dynamic")
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
		switch tvalue := value.(type) {
		case int32:
			typ = "integer"
		case int:
			typ = "integer"
		case int64:
			typ = "integer"
		case uint32:
			typ = "biginteger"
		case uint:
			typ = "biginteger"
		case uint64:
			typ = "biginteger"
		case float32:
			typ = "decimal"
		case string:
			typ = "string"
		case net.IP:
			typ = "ipAddress"
		case *net.IP:
			typ = "ipAddress"
		case net.HardwareAddr:
			typ = "physicalAddress"
		case *net.HardwareAddr:
			typ = "physicalAddress"
		case time.Time:
			typ = "datetime"
		case *time.Time:
			typ = "datetime"
		case json.Number:
			if _, err := tvalue.Int64(); err == nil {
				typ = "integer"
			} else {
				typ = "decimal"
			}
		case *json.Number:
			if _, err := tvalue.Int64(); err == nil {
				typ = "integer"
			} else {
				typ = "decimal"
			}
		default:
			return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
		}
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
	return nil, ErrUnsupportedFunc(operator, getCheckedType(typ, value))
}
