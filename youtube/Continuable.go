// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/grafana/sobek"
)

type Continuable[T any] struct {
	vm   *sobek.Runtime
	this *sobek.Object
}

func (x *Continuable[T]) Items() ([]T, error) {
	var result []T
	if err := internal.ExportTo(x.vm, x.this.Get("items"), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (x *Continuable[T]) Next(count int64) ([]T, error) {
	val, err := internal.CallAsync(x.vm, x.this, "next", x.vm.ToValue(count))
	if err != nil {
		return nil, err
	}
	var result []T
	if err = internal.ExportTo(x.vm, val, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (x *Continuable[T]) FromObject(vm *sobek.Runtime, obj *sobek.Object) error {
	x.vm = vm
	x.this = obj
	return nil
}
