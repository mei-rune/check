package check

import "testing"

func TestDatetime(t *testing.T) {

	testcases := []check_data{
		{t: "datetime", operator: ">", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:13+08:00", excepted_status: true},
		{t: "datetime", operator: ">", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:11+08:00", excepted_status: false},
		{t: "datetime", operator: ">=", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		{t: "datetime", operator: "<", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:13+08:00", excepted_status: false},
		{t: "datetime", operator: "<", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:11+08:00", excepted_status: true},
		{t: "datetime", operator: "<=", operant: "2025-01-06T12:12:12+08:00", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		
		{t: "datetime", operator: ">", operant: "now() - 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		{t: "datetime", operator: ">=", operant: "now() - 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},


		{t: "datetime", operator: ">", operant: "now( ) - 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		{t: "datetime", operator: ">=", operant: "now( ) - 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},


		{t: "datetime", operator: "<", operant: "now() + 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		{t: "datetime", operator: "<=", operant: "now() + 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		
		{t: "datetime", operator: "<", operant: "now( ) + 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		{t: "datetime", operator: "<=", operant: "now( ) + 99999h", value: "2025-01-06T12:12:12+08:00", excepted_status: true},
		
	}
	for _, test := range testcases {
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
		if status != test.excepted_status {
			t.Errorf("(%T) %#v %v (%T) %#v :  [%#v] failed, excepted status is %v, actual status is %v",
				test.value, test.value, test.operator, test.operant, test.operant, test, test.excepted_status, status)
		}
	}
}
