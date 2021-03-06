// Code generated by "stringer -type=StatusCode -linecomment -output stringer_string.go"; DO NOT EDIT.

package practice

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Success-1]
	_ = x[Failed-2]
}

const _StatusCode_name = "未知成功失败"

var _StatusCode_index = [...]uint8{0, 6, 12, 18}

func (i StatusCode) String() string {
	if i < 0 || i >= StatusCode(len(_StatusCode_index)-1) {
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatusCode_name[_StatusCode_index[i]:_StatusCode_index[i+1]]
}
