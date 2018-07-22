package check

import (
	"testing"
)

type check_data struct {
	t        string
	operator string
	operant  string
	value    interface{}

	excepted_error  string
	excepted_status bool
	excepted_value  interface{}
}

func TestChecker(t *testing.T) {
	var all_check_data []check_data

	//, float32(12), float64(12)
	for _, v := range []interface{}{"12", uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12)} {
		for _, test := range []check_data{
			{t: "biginteger", operator: ">", operant: "11", value: v, excepted_status: true, excepted_value: int64(13)},
			{t: "biginteger", operator: ">", operant: "12", value: v, excepted_status: false, excepted_value: int64(12)},

			{t: "biginteger", operator: ">=", operant: "11", value: v, excepted_status: true, excepted_value: int64(13)},
			{t: "biginteger", operator: ">=", operant: "12", value: v, excepted_status: true, excepted_value: int64(12)},
			{t: "biginteger", operator: ">=", operant: "13", value: v, excepted_status: false, excepted_value: int64(11)},

			{t: "biginteger", operator: "<", operant: "14", value: v, excepted_status: true, excepted_value: int64(13)},
			{t: "biginteger", operator: "<", operant: "12", value: v, excepted_status: false, excepted_value: int64(12)},

			{t: "biginteger", operator: "<=", operant: "14", value: v, excepted_status: true, excepted_value: int64(13)},
			{t: "biginteger", operator: "<=", operant: "12", value: v, excepted_status: true, excepted_value: int64(12)},
			{t: "biginteger", operator: "<=", operant: "11", value: v, excepted_status: false, excepted_value: int64(11)},

			{t: "biginteger", operator: "=", operant: "12", value: v, excepted_status: true, excepted_value: "11"},
			{t: "biginteger", operator: "==", operant: "12", value: v, excepted_status: true, excepted_value: "11"},
			{t: "biginteger", operator: "!=", operant: "11", value: v, excepted_status: true, excepted_value: "12"},
			{t: "biginteger", operator: "<>", operant: "11", value: v, excepted_status: true, excepted_value: "12"},

			{t: "biginteger", operator: "=", operant: "11", value: v, excepted_status: false, excepted_value: "11"},
			{t: "biginteger", operator: "==", operant: "11", value: v, excepted_status: false, excepted_value: "11"},
			{t: "biginteger", operator: "!=", operant: "12", value: v, excepted_status: false, excepted_value: "12"},
			{t: "biginteger", operator: "<>", operant: "12", value: v, excepted_status: false, excepted_value: "12"},
		} {
			all_check_data = append(all_check_data, test)
		}
	}

	for _, class := range []string{"integer", "biginteger", "decimal"} {
		for _, test := range []check_data{
			{t: class, operator: ">", operant: "12", value: 13, excepted_status: true, excepted_value: int64(13)},
			{t: class, operator: ">", operant: "12", value: 12, excepted_status: false, excepted_value: int64(12)},

			{t: class, operator: ">=", operant: "12", value: 13, excepted_status: true, excepted_value: int64(13)},
			{t: class, operator: ">=", operant: "12", value: 12, excepted_status: true, excepted_value: int64(12)},
			{t: class, operator: ">=", operant: "12", value: 11, excepted_status: false, excepted_value: int64(11)},

			{t: class, operator: "<", operant: "14", value: 13, excepted_status: true, excepted_value: int64(13)},
			{t: class, operator: "<", operant: "12", value: 12, excepted_status: false, excepted_value: int64(12)},

			{t: class, operator: "<=", operant: "14", value: 13, excepted_status: true, excepted_value: int64(13)},
			{t: class, operator: "<=", operant: "12", value: 12, excepted_status: true, excepted_value: int64(12)},
			{t: class, operator: "<=", operant: "11", value: 12, excepted_status: false, excepted_value: int64(11)},

			{t: class, operator: "=", operant: "11", value: 11, excepted_status: true, excepted_value: "11"},
			{t: class, operator: "==", operant: "11", value: 11, excepted_status: true, excepted_value: "11"},
			{t: class, operator: "!=", operant: "11", value: "12", excepted_status: true, excepted_value: "12"},
			{t: class, operator: "<>", operant: "11", value: "12", excepted_status: true, excepted_value: "12"},

			{t: class, operator: "=", operant: "12", value: 11, excepted_status: false, excepted_value: "11"},
			{t: class, operator: "==", operant: "12", value: 11, excepted_status: false, excepted_value: "11"},
			{t: class, operator: "!=", operant: "12", value: "12", excepted_status: false, excepted_value: "12"},
			{t: class, operator: "<>", operant: "12", value: "12", excepted_status: false, excepted_value: "12"},
		} {
			all_check_data = append(all_check_data, test)
		}
	}

	for _, class := range []string{"ipAddress", "string"} {
		for _, test := range []struct {
			t        string
			operator string
			operant  string
			value    interface{}

			excepted_error  string
			excepted_status bool
			excepted_value  interface{}
		}{
			{t: class, operator: "=", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: true, excepted_value: "a"},
			{t: class, operator: "==", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: true, excepted_value: "a"},
			{t: class, operator: "!=", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: true, excepted_value: "abc"},
			{t: class, operator: "<>", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: true, excepted_value: "abc"},

			{t: class, operator: "=", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: false, excepted_value: "abc"},
			{t: class, operator: "==", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: false, excepted_value: "abc"},
			{t: class, operator: "!=", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: false, excepted_value: "a"},
			{t: class, operator: "<>", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: false, excepted_value: "a"},
		} {
			all_check_data = append(all_check_data, test)
		}
	}

	for _, class := range []string{"physicalAddress"} {
		for _, test := range []struct {
			t        string
			operator string
			operant  string
			value    interface{}

			excepted_error  string
			excepted_status bool
			excepted_value  interface{}
		}{
			{t: class, operator: "=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: true, excepted_value: "a"},
			{t: class, operator: "==", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: true, excepted_value: "a"},
			{t: class, operator: "!=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: true, excepted_value: "abc"},
			{t: class, operator: "<>", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: true, excepted_value: "abc"},

			{t: class, operator: "=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: false, excepted_value: "abc"},
			{t: class, operator: "==", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: false, excepted_value: "abc"},
			{t: class, operator: "!=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: false, excepted_value: "a"},
			{t: class, operator: "<>", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: false, excepted_value: "a"},
		} {
			all_check_data = append(all_check_data, test)
		}
	}
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "true", value: "true", excepted_status: true, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "true", value: "true", excepted_status: true, excepted_value: "a"})

	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "false", value: "false", excepted_status: true, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "false", value: "false", excepted_status: true, excepted_value: "a"})

	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "true", value: "false", excepted_status: true, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "true", value: "false", excepted_status: true, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "true", value: "false", excepted_status: false, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "true", value: "false", excepted_status: false, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "true", value: "true", excepted_status: false, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "true", value: "true", excepted_status: false, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "false", value: "false", excepted_status: false, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "false", value: "false", excepted_status: false, excepted_value: "a"})

	all_check_data = append(all_check_data, check_data{t: "string", operator: "equals", operant: "a", value: "a", excepted_status: true, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_equals", operant: "a", value: "a", excepted_status: false, excepted_value: "a"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "contains", operant: "a", value: "abc", excepted_status: true, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_contains", operant: "a", value: "abc", excepted_status: false, excepted_value: "abc"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "equals_with_ignore_case", operant: "a", value: "A", excepted_status: true, excepted_value: "A"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_equals_with_ignore_case", operant: "a", value: "A", excepted_status: false, excepted_value: "A"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "contains_with_ignore_case", operant: "a", value: "Abc", excepted_status: true, excepted_value: "Abc"})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_contains_with_ignore_case", operant: "a", value: "Abc", excepted_status: false, excepted_value: "Abc"})

	for _, test := range all_check_data {
		check, e := MakeChecker(test.t, test.operator, test.operant)
		if nil != e {
			if 0 == len(test.excepted_error) {
				t.Errorf("[%#v] failed, %v", test, e)
			} else if test.excepted_error != e.Error() {
				t.Errorf("[%#v] failed, excepted is '%v', actual is '%v'", test, test.excepted_error, e)
			}
			continue
		}

		status, e := check.Check(test.value)
		if nil != e {
			if 0 == len(test.excepted_error) {
				t.Errorf("[%#v] failed, %v", test, e)
			} else if test.excepted_error != e.Error() {
				t.Errorf("[%#v] failed, excepted error is '%v', actual error is '%v'", test, test.excepted_error, e)
			}
			continue
		}

		// if v != test.excepted_value {
		// 	t.Errorf("test all_checker_tests[%v] failed, excepted v is %v, actual v is %v", idx, test.excepted_value, v)
		// }

		if status != test.excepted_status {
			t.Errorf("[%#v] failed, excepted status is %v, actual status is %v", test, test.excepted_status, status)
		}
	}
}
