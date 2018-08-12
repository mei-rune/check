package check

import (
	"errors"
	"fmt"

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
	return fmt.Sprintf("%s cannot convert to %s", spew.Sprint(e.value), e.typ)
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
