package check

import (
	"encoding/json"
	"math"
	"strconv"
)

func sub(a, b interface{}) (interface{}, error) {
	switch aValue := a.(type) {
	case float64:
		return subFloat64(aValue, b)
	case int64:
		return subInt64(aValue, b)
	case int:
		return subInt64(int64(aValue), b)
	case uint64:
		return subUint64(aValue, b)
	case uint:
		return subUint64(uint64(aValue), b)
	case string:
		return subStr(aValue, b)
	case json.Number:
		return subStr(aValue.String(), b)
	case uint8:
		return subUint64(uint64(aValue), b)
	case uint16:
		return subUint64(uint64(aValue), b)
	case uint32:
		return subUint64(uint64(aValue), b)
	case int8:
		return subInt64(int64(aValue), b)
	case int16:
		return subInt64(int64(aValue), b)
	case int32:
		return subInt64(int64(aValue), b)
	case float32:
		return subFloat64(float64(aValue), b)
	case []byte:
		return subStr(string(aValue), b)
	case *json.Number:
		return subStr(aValue.String(), b)
	}
	if nil == a {
		return false, ErrValueNull
	}
	return false, errType("sub", "uint64")
}

func subStr(a string, b interface{}) (interface{}, error) {
	if u64, e := strconv.ParseUint(a, 10, 64); e == nil {
		return subUint64(u64, b)
	}
	if i64, e := strconv.ParseInt(a, 10, 64); e == nil {
		return subInt64(i64, b)
	}
	if f64, e := strconv.ParseFloat(a, 64); e == nil {
		return subFloat64(f64, b)
	}
	return nil, ErrActualType("sub", "", a)
}

func subFloat64(a float64, b interface{}) (interface{}, error) {
	switch bValue := b.(type) {
	case float64:
		return a - bValue, nil
	case int64:
		return a - float64(bValue), nil
	case int:
		return a - float64(bValue), nil
	case uint64:
		return a - float64(bValue), nil
	case uint:
		return a - float64(bValue), nil
	case json.Number:
		s := bValue.String()
		if f64, e := strconv.ParseFloat(s, 64); e == nil {
			return a - f64, nil
		}
		return nil, ErrActualType("sub", "", s)
	case string:
		if f64, e := strconv.ParseFloat(bValue, 64); e == nil {
			return a - f64, nil
		}
		return nil, ErrActualType("sub", "", bValue)
	case uint8:
		return a - float64(bValue), nil
	case uint16:
		return a - float64(bValue), nil
	case uint32:
		return a - float64(bValue), nil
	case int8:
		return a - float64(bValue), nil
	case int16:
		return a - float64(bValue), nil
	case int32:
		return a - float64(bValue), nil
	case float32:
		return a - float64(bValue), nil
	case []byte:
		s := string(bValue)
		if f64, e := strconv.ParseFloat(s, 64); e == nil {
			return a - f64, nil
		}
		return nil, ErrActualType("sub", "", s)
	case *json.Number:
		s := bValue.String()
		if f64, e := strconv.ParseFloat(s, 64); e == nil {
			return a - f64, nil
		}
		return nil, ErrActualType("sub", "", s)
	}
	if nil == b {
		return nil, ErrValueNull
	}
	return nil, errType("sub", "uint64")
}

func subInt64(a int64, b interface{}) (interface{}, error) {
	switch bValue := b.(type) {
	case float64:
		return float64(a) - bValue, nil
	case int64:
		return a - bValue, nil
	case int:
		return a - int64(bValue), nil
	case uint64:
		return int64SubUint64(a, bValue)
	case uint:
		return int64SubUint64(a, uint64(bValue))
	case json.Number:
		return int64SubStr(a, bValue.String())
	case string:
		return int64SubStr(a, bValue)
	case uint8:
		return a - int64(bValue), nil
	case uint16:
		return a - int64(bValue), nil
	case uint32:
		return a - int64(bValue), nil
	case int8:
		return a - int64(bValue), nil
	case int16:
		return a - int64(bValue), nil
	case int32:
		return a - int64(bValue), nil
	case float32:
		return float64(a) - float64(bValue), nil
	case []byte:
		return int64SubStr(a, string(bValue))
	case *json.Number:
		return int64SubStr(a, bValue.String())
	}
	if nil == b {
		return nil, ErrValueNull
	}
	return nil, errType("sub", "uint64")
}

func subUint64(a uint64, b interface{}) (interface{}, error) {
	switch bValue := b.(type) {
	case float64:
		return float64(a) - bValue, nil
	case int64:
		return uint64SubInt64(a, bValue)
	case int:
		return uint64SubInt64(a, int64(bValue))
	case uint64:
		if a > bValue {
			return a - bValue, nil
		}
		return -int64(bValue - a), nil
	case uint:
		if a > uint64(bValue) {
			return a - uint64(bValue), nil
		}
		return -int64(uint64(bValue) - a), nil
	case json.Number:
		return uint64SubStr(a, bValue.String())
	case string:
		return uint64SubStr(a, bValue)
	case uint8:
		return a - uint64(bValue), nil
	case uint16:
		return a - uint64(bValue), nil
	case uint32:
		return a - uint64(bValue), nil
	case int8:
		if bValue < 0 {
			return a + uint64(-bValue), nil
		}
		return a - uint64(bValue), nil
	case int16:
		if bValue < 0 {
			return a + uint64(-bValue), nil
		}
		return a - uint64(bValue), nil
	case int32:
		if bValue < 0 {
			return a + uint64(-bValue), nil
		}
		return a - uint64(bValue), nil
	case float32:
		return float64(a) - float64(bValue), nil
	case []byte:
		return uint64SubStr(a, string(bValue))
	case *json.Number:
		return uint64SubStr(a, bValue.String())
	}
	if nil == b {
		return nil, ErrValueNull
	}
	return nil, errType("sub", "uint64")
}

func uint64SubInt64(a uint64, b int64) (interface{}, error) {
	if b == 0 {
		return a, nil
	}
	if b > 0 {
		return a - uint64(b), nil
	}
	f64 := float64(a) + float64(-b)
	if f64 > math.MaxUint64 {
		return f64, nil
	}
	return uint64(f64), nil
}

func int64SubUint64(a int64, b uint64) (interface{}, error) {
	if b == 0 {
		return a, nil
	}

	if a >= 0 {
		if uint64(a) > b {
			return a - int64(b), nil
		}

		aValue := b - uint64(a)
		if a > math.MaxInt64 {
			return -float64(aValue), nil
		}
		return -int64(aValue), nil
	}

	f64 := float64(-a) + float64(b)
	if f64 < math.MinInt64 {
		return f64, nil
	}
	return int64(f64), nil
}

func int64SubStr(a int64, s string) (interface{}, error) {
	if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
		return int64SubUint64(a, u64)
	}
	if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
		return a - i64, nil
	}
	if f64, e := strconv.ParseFloat(s, 64); e == nil {
		return float64(a) - float64(f64), nil
	}
	return nil, ErrActualType("sub", "", s)
}

func uint64SubStr(a uint64, s string) (interface{}, error) {
	if u64, e := strconv.ParseUint(s, 10, 64); e == nil {
		if a > u64 {
			return a - u64, nil
		}
		return -int64(u64 - a), nil
	}
	if i64, e := strconv.ParseInt(s, 10, 64); e == nil {
		return uint64SubInt64(a, i64)
	}
	if f64, e := strconv.ParseFloat(s, 64); e == nil {
		return float64(a) - f64, nil
	}
	return nil, ErrActualType("sub", "", s)
}

// func subUint64(a uint64, b interface{}) interface{} {
//   switch bv := b.(type) {
//   case float64:
//     return float64(a) - bv
//   case float32:
//     return float64(a) - float64(bv)
//   case uint64:
//     return a - bv
//   case uint:
//     return a - uint64(bv)
//   case json.Number:
//     bb, e := bv.Int64()
//     if e == nil {
//       if bb > 0 {
//         return a - uint64(bb)
//       }
//       return float64(a) - float64(bb)
//     }

//     f64, e := bv.Float64()
//     if e != nil {
//       panic(e)
//     }
//     return float64(a) - f64
//   }
//   return float64(a) - toFloat(b)
// }

// func subWithFloat(a interface{}, b float64) interface{} {
//   switch av := a.(type) {
//   case float64:
//     return av - b
//   case float32:
//     return float64(av) - b
//   case uint64:
//     return float64(av) - b
//   case uint:
//     return float64(av) - b
//   case json.Number:
//     f64, e := av.Float64()
//     if e != nil {
//       panic(e)
//     }
//     return f64 - b
//   }
//   return toFloat(a) - b
// }

// func subWithUint64(a interface{}, b uint64) interface{} {
//   switch av := a.(type) {
//   case float64:
//     return av - float64(b)
//   case float32:
//     return float64(av) - float64(b)
//   case uint64:
//     return av - b
//   case uint:
//     return uint64(av) - b
//   case json.Number:
//     aa, e := av.Int64()
//     if e == nil {
//       if aa > 0 {
//         return uint64(aa) - b
//       }
//       return float64(aa) - b
//     }

//     f64, e := av.Float64()
//     if e != nil {
//       panic(e)
//     }
//     return f64 - float64(b)
//   }
//   return toFloat(a) - float64(b)
// }
