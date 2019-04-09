package template

import (
	"testing"
)

func TestIsEmployeeOrStudentNone(t *testing.T) {
	ret := isEmployeeOrStudent([]string{"None"}, nil)

	if ret {
		t.Errorf("Expected isEmployeeOrStudent([]string{'None'}) to return false. Got: %t", ret)
	}
}
func TestIsEmployeeOrStudentEmployee(t *testing.T) {
	ret := isEmployeeOrStudent([]string{"Employee"}, nil)

	if !ret {
		t.Errorf("Expected isEmployeeOrStudent([]string{'Employee'}) to return true. Got: %t", ret)
	}
}
func TestIsEmployeeOrStudentStudent(t *testing.T) {
	ret := isEmployeeOrStudent([]string{"Student"}, nil)

	if !ret {
		t.Errorf("Expected isEmployeeOrStudent([]string{'Student'}) to return true. Got: %t", ret)
	}
}
