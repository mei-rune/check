package check

import (
	"encoding/json"
	"testing"
)

func TestInChecker(t *testing.T) {
	var all_check_data []check_data

	n11 := json.Number("11")
	n12 := json.Number("12")
	//n13 := json.Number("13")

	// n11__001 := json.Number("11.001")
	// n12__001 := json.Number("12.001")
	// n13__001 := json.Number("13.001")

	// n_11 := json.Number("-11")
	// n_12 := json.Number("-12")
	// n_13 := json.Number("-13")

	//n_11__001 := json.Number("-11.001")
	// n_12__001 := json.Number("-12.001")
	//n_13__001 := json.Number("-13.001")

	for _, class := range []string{"integer", "biginteger", "decimal", ""} {
		actual := []interface{}{n12, &n12, "12", uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12)}
		if class != "" {
			actual = append(actual, []byte("12"))
			actual = append(actual, float32(12), float64(12))
		}
		for _, v := range actual {

			aa := []interface{}{n11, &n11, "11", uint(11), uint8(11), uint16(11), uint32(11), uint64(11), int(11), int8(11), int16(11), int32(11), int64(11)}
			if class != "" {
				aa = append(aa, []byte("11"))
				aa = append(aa, float32(11), float64(11))
			}
			for _, operant := range aa {
				all_check_data = append(all_check_data, check_data{t: class, operator: "in", operant: []interface{}{operant}, value: v, excepted_status: false})
				all_check_data = append(all_check_data, check_data{t: class, operator: "nin", operant: []interface{}{operant}, value: v, excepted_status: true})

				all_check_data = append(all_check_data, check_data{t: class, operator: "in", operant: operant, value: v, excepted_status: false})
				all_check_data = append(all_check_data, check_data{t: class, operator: "nin", operant: operant, value: v, excepted_status: true})

			}
			aa = []interface{}{n12, &n12, "12", uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12)}
			if class != "" {
				aa = append(aa, []byte("12"))
				aa = append(aa, float32(12), float64(12))
			}
			for _, operant := range aa {
				all_check_data = append(all_check_data, check_data{t: class, operator: "in", operant: []interface{}{operant}, value: v, excepted_status: true})
				all_check_data = append(all_check_data, check_data{t: class, operator: "nin", operant: []interface{}{operant}, value: v, excepted_status: false})
			}

			ss := [2]string{"11", "12"}
			aa = []interface{}{
				[]uint{uint(11), uint(12)},
				[]uint16{uint16(11), uint16(12)},
				[]uint32{uint32(11), uint32(12)},
				[]uint64{uint64(11), uint64(12)},
				[]int{int(11), int(12)},
				[]int8{int8(11), int8(12)},
				[]int16{int16(11), int16(12)},
				[]int32{int32(11), int32(12)},
				[]int64{int64(11), int64(12)},

				[]interface{}{n11, n12},
				[]interface{}{&n11, &n12},
				[]interface{}{"11", "12"},
				[]interface{}{"11", 12},
				[]interface{}{"11", int64(12)},
				"11,12",
				[]string{"11", "12"},
				[2]string{"11", "12"},
				[2]string{"11", "12"},
				&ss,
			}
			if class != "" {
				aa = append(aa, []uint8{uint8(11), uint8(12)})
				aa = append(aa, []interface{}{[]byte("11"), []byte("12")})
			}
			for _, operant := range aa {
				all_check_data = append(all_check_data, check_data{t: class, operator: "in", operant: operant, value: v, excepted_status: true})
				all_check_data = append(all_check_data, check_data{t: class, operator: "nin", operant: operant, value: v, excepted_status: false})
			}
		}
	}

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
		if status != test.excepted_status {
			t.Errorf("(%T) %#v %v (%T) %#v :  [%#v] failed, excepted status is %v, actual status is %v",
				test.value, test.value, test.operator, test.operant, test.operant, test, test.excepted_status, status)
		}
	}
}
