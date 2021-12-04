package structs

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		s := messageFromMsgAndArgs(msgAndArgs...)
		t.Error(s, "\n  Expect: ", expected, "\n  Actual: ", actual)
	}
}

func assertNil(t *testing.T, actual interface{}, msgAndArgs ...interface{}) {
	if !isNil(actual) {
		s := messageFromMsgAndArgs(msgAndArgs...)
		t.Error(s, "\n  Expect: ", nil, "\n  Actual: ", actual)
	}
}

func assertTrue(t *testing.T, actual interface{}, msgAndArgs ...interface{}) {
	if true != actual {
		s := messageFromMsgAndArgs(msgAndArgs...)
		t.Error(s, "\n  Expect: ", true, "\n  Actual: ", actual)
	}
}

func assertPanics(t *testing.T, f func(), msgAndArgs ...interface{}) {
	if funcDidPanic, panicValue := didPanic(f); !funcDidPanic {
		s := messageFromMsgAndArgs(msgAndArgs...)
		t.Error(s, fmt.Sprint("\n  func should panic\n\tPanic value:", panicValue))
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

func didPanic(f func()) (bool, interface{}) {
	didPanic := false
	var message interface{}
	func() {

		defer func() {
			if message = recover(); message != nil {
				didPanic = true
			}
		}()

		// call the target function
		f()

	}()

	return didPanic, message
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
