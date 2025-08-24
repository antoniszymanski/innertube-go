// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Video struct {
	BaseVideo
	// The duration of this video in second
	Duration int64 `js:"duration"`
	// Chapters on this video if exists
	Chapters Chapters
	// Continuable of videos inside a Video
	Comments VideoComments
}

type VideoComments = Continuable[Comment]

func (x *Video) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj, &x.BaseVideo); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("chapters"), &x.Chapters); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("comments"), &x.Comments); err != nil {
		return err
	}
	return nil
}

type Chapters []Chapter

func (x *Chapters) FromValue(vm *goja.Runtime, iterable goja.Value) error {
	return (*utils.Array[Chapter])(x).FromValue(vm, iterable)
}

type Chapter struct {
	Title      string `js:"title"`
	Start      int64  `js:"start"`
	Thumbnails shared.Thumbnails
}

func (x *Chapter) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("thumbnails"), &x.Thumbnails); err != nil {
		return err
	}
	return nil
}
