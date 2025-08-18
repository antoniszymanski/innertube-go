// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"fmt"
	"reflect"

	"github.com/dop251/goja"
)

func Call(vm *goja.Runtime, this *goja.Object, name string, args ...goja.Value) (goja.Value, error) {
	val, err := call(this, name)
	if err != nil {
		return nil, err
	}
	if err = extractError(vm, val); err != nil {
		return nil, err
	}
	return val, nil
}

func CallAsync(vm *goja.Runtime, this *goja.Object, name string, args ...goja.Value) (goja.Value, error) {
	val, err := call(this, name, args...)
	if err != nil {
		return nil, err
	}
	p, ok := val.Export().(*goja.Promise)
	if !ok {
		return nil, ErrNotPromise{name, val}
	}
	switch p.State() {
	case goja.PromiseStateFulfilled:
		if err = extractError(vm, p.Result()); err != nil {
			return nil, err
		}
		return p.Result(), nil
	case goja.PromiseStateRejected:
		if err = extractError(vm, p.Result()); err != nil {
			return nil, err
		}
		return nil, ErrPromiseRejected{p.Result()}
	default:
		return nil, ErrPromisePending{}
	}
}

type ErrNotPromise struct {
	FunctionName string
	ActualValue  goja.Value
}

func (e ErrNotPromise) Error() string {
	return fmt.Sprintf("function %q didn't return a Promise, got: %s", e.FunctionName, e.ActualValue)
}

type ErrPromiseRejected struct {
	PromiseResult goja.Value
}

func (e ErrPromiseRejected) Error() string {
	return "promise rejected: " + e.PromiseResult.String()
}

type ErrPromisePending struct{}

func (e ErrPromisePending) Error() string {
	return "unexpected promise state: pending"
}

func call(this *goja.Object, name string, args ...goja.Value) (goja.Value, error) {
	val := this.Get(name)
	if val == nil {
		return nil, ErrPropertyNotExist{name}
	}
	fn, ok := goja.AssertFunction(val)
	if !ok {
		return nil, ErrNotFunction{name}
	}
	return fn(this, args...)
}

type ErrPropertyNotExist struct {
	PropertyName string
}

func (e ErrPropertyNotExist) Error() string {
	return fmt.Sprintf("property %q does not exist", e.PropertyName)
}

type ErrNotFunction struct {
	PropertyName string
}

func (e ErrNotFunction) Error() string {
	return fmt.Sprintf("property %q is not a function", e.PropertyName)
}

func extractError(vm *goja.Runtime, val goja.Value) error {
	obj, err := ToObject(vm, val)
	if err != nil {
		return nil
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
	if v := reflect.ValueOf(obj).Elem().FieldByName("self"); v.IsValid() {
		if v = v.Elem().Elem().FieldByName("stack"); v.IsValid() {
			e.Stack = *(*[]goja.StackFrame)(v.Addr().UnsafePointer())
		}
	}
	return &e
}

type Error struct {
	Name      string
	Message   string
	Cause     any
	Stack     []goja.StackFrame
	StackText string
}

func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	} else {
		return e.Message + ": " + fmt.Sprint(e.Cause)
	}
}
