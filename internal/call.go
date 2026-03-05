// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"fmt"
	"reflect"

	"github.com/grafana/sobek"
)

func Call(vm *sobek.Runtime, this *sobek.Object, name string, args ...sobek.Value) (sobek.Value, error) {
	val, err := call(this, name)
	if err != nil {
		return nil, err
	}
	if err := extractError(vm, val); err != nil {
		return nil, err
	}
	return val, nil
}

func CallAsync(vm *sobek.Runtime, this *sobek.Object, name string, args ...sobek.Value) (sobek.Value, error) {
	val, err := call(this, name, args...)
	if err != nil {
		return nil, err
	}
	p, ok := val.Export().(*sobek.Promise)
	if !ok {
		return nil, NotPromiseError{name, val}
	}
	switch p.State() {
	case sobek.PromiseStateFulfilled:
		if err := extractError(vm, p.Result()); err != nil {
			return nil, err
		}
		return p.Result(), nil
	case sobek.PromiseStateRejected:
		if err := extractError(vm, p.Result()); err != nil {
			return nil, err
		}
		return nil, PromiseRejectedError{p.Result()}
	default:
		return nil, PromisePendingError{}
	}
}

type NotPromiseError struct {
	FunctionName string
	ActualValue  sobek.Value
}

func (e NotPromiseError) Error() string {
	return fmt.Sprintf("function %q didn't return a Promise, got: %s", e.FunctionName, e.ActualValue)
}

type PromiseRejectedError struct {
	PromiseResult sobek.Value
}

func (e PromiseRejectedError) Error() string {
	return "promise rejected: " + e.PromiseResult.String()
}

type PromisePendingError struct{}

func (e PromisePendingError) Error() string {
	return "unexpected promise state: pending"
}

func call(this *sobek.Object, name string, args ...sobek.Value) (sobek.Value, error) {
	val := this.Get(name)
	if val == nil {
		return nil, PropertyNotExistError{name}
	}
	fn, ok := sobek.AssertFunction(val)
	if !ok {
		return nil, NotFunctionError{name}
	}
	return fn(this, args...)
}

type PropertyNotExistError struct {
	PropertyName string
}

func (e PropertyNotExistError) Error() string {
	return fmt.Sprintf("property %q does not exist", e.PropertyName)
}

type NotFunctionError struct {
	PropertyName string
}

func (e NotFunctionError) Error() string {
	return fmt.Sprintf("property %q is not a function", e.PropertyName)
}

func extractError(vm *sobek.Runtime, val sobek.Value) *Error {
	obj, err := ToObject(vm, val)
	if err != nil {
		return nil //nolint:nilerr
	}
	if obj.ClassName() != "Error" {
		return nil
	}
	var e Error
	if name := obj.Get("name"); name != nil {
		e.Name = name.String()
	}
	if message := obj.Get("message"); message != nil {
		e.Message = message.String()
	}
	if cause := obj.Get("cause"); cause != nil {
		e.Cause = cause.Export()
	}
	if stackText := obj.Get("stack"); stackText != nil {
		e.StackText = stackText.String()
	}
	e.Stack = extractStack(obj)
	return &e
}

func extractStack(obj *sobek.Object) []sobek.StackFrame {
	v := reflect.ValueOf(obj).Elem().FieldByName("self")
	if !v.IsValid() {
		return nil
	}
	if v.Kind() != reflect.Interface {
		return nil
	}
	if v = v.Elem(); v.Kind() != reflect.Pointer {
		return nil
	}
	if v = v.Elem(); v.Kind() != reflect.Struct {
		return nil
	}
	stack, ok := reflect.TypeAssert[*[]sobek.StackFrame](field(v, "stack"))
	if !ok {
		return nil
	}
	return *stack
}

type Error struct {
	Name      string
	Message   string
	Cause     any
	Stack     []sobek.StackFrame
	StackText string
}

func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Cause)
}
