package errors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	//Test Cases
	scenarios := []struct {
		//Parameters
		msg string
		//Result
		err error
	}{
		{"MOCK MSG", &Error{msg: "MOCK MSG"}},
		{"", &Error{msg: ""}},
		{"1234567890", &Error{msg: "1234567890"}},
		{"!@#$%^&*()", &Error{msg: "!@#$%^&*()"}},
	}

	//Run Tests
	for index, scenario := range scenarios {
		t.Run(fmt.Sprintf("Case%d", index+1), func(t *testing.T) {
			err := New(scenario.msg)
			if !reflect.DeepEqual(err, scenario.err) {
				t.Errorf("Actual '%v' | Expected '%v'", err, scenario.err)
			}
		})
	}
}

func TestError(t *testing.T) {
	//Test Cases
	scenarios := []struct {
		//Parameters
		msg string
		//Result
		errMsg string
	}{
		{"MOCK_MSG", "MOCK_MSG"},
		{"", ""},
		{"123456789", "123456789"},
		{"!@#$%^&*()", "!@#$%^&*()"},
	}

	//Run Tests
	for index, scenario := range scenarios {
		t.Run(fmt.Sprintf("Case%d", index+1), func(t *testing.T) {
			err := New(scenario.msg)
			if scenario.errMsg != err.Error() {
				t.Errorf("Actual '%v' | Expected '%v'", err.Error(), scenario.errMsg)
			}
		})
	}
}
