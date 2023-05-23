package function

import (
	"testing"
)

func TestCeil(t *testing.T) {
	var nums int
	var f64, result float64
	f64 = 3.14
	result = 4
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
	}
	nums = 1
	f64 = 3.14
	result = 3.2
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.14
	result = 3.14
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.14
	result = 3.14
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 0
	f64 = 3
	result = 3
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3
	result = 3
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3
	result = 3
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 0
	f64 = 3.00
	result = 3.00
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3.00
	result = 3.00
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.00
	result = 3.00
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.00
	result = 3.00
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 0
	f64 = 3.10
	result = 4
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3.10
	result = 3.1
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.10
	result = 3.1
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.10
	result = 3.1
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 0
	f64 = 3.01
	result = 4
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3.01
	result = 3.1
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.01
	result = 3.01
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.01
	result = 3.01
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3.91
	result = 4.0
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.91
	result = 3.91
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.91
	result = 3.91
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 0
	f64 = 3.01990099
	result = 4
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 1
	f64 = 3.01990099
	result = 3.1
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 2
	f64 = 3.01990099
	result = 3.02
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 3
	f64 = 3.01990099
	result = 3.02
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 4
	f64 = 3.01990099
	result = 3.02
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 5
	f64 = 3.01990099
	result = 3.01991
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 6
	f64 = 3.01990099
	result = 3.019901
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 7
	f64 = 3.01990099
	result = 3.019901
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 8
	f64 = 3.01990099
	result = 3.01990099
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
	nums = 9
	f64 = 3.01990099
	result = 3.01990099
	if result != Ceil(f64, nums) {
		t.Error(result, " !=  Ceil(", f64, ", ", nums, ")")
		return
	}
}
