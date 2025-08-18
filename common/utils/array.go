// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils

import "github.com/dop251/goja"

type Array[T any] []T

func (x *Array[T]) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	val := obj.Get("length")
	if val == nil {
		return ErrPropertyNotExist{"length"}
	}
	length := val.ToInteger()
	if length < 0 {
		return ErrNegativeArrayLength{}
	}
	for _, name := range IndicesSeq(uint64(length)) {
		val := obj.Get(name)
		if val == nil {
			return ErrPropertyNotExist{name}
		}
		var elem T
		err := ExportTo(vm, val, &elem)
		if err != nil {
			return err
		}
		*x = append(*x, elem)
	}
	return nil
}

type ErrNegativeArrayLength struct{}

func (e ErrNegativeArrayLength) Error() string {
	return "array length must not be negative"
}
