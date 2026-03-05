// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/antoniszymanski/innertube-go/shared"
	. "github.com/antoniszymanski/option-go"
	"github.com/grafana/sobek"
)

type Video struct {
	BaseVideo
	// The duration of this video in second
	Duration int64 `js:"duration"`
	// Chapters on this video if exists
	Chapters []Chapter
	// Continuable of videos inside a Video
	Comments VideoComments
	// Music metadata (if exists)
	Music Option[MusicMetadata]
}

type VideoComments = Continuable[Comment]

func (x *Video) FromObject(vm *sobek.Runtime, obj *sobek.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj, &x.BaseVideo); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("chapters"), &x.Chapters); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("comments"), &x.Comments); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("music"), &x.Music); err != nil {
		return err
	}
	return nil
}

type Chapter struct {
	Title      string `js:"title"`
	Start      int64  `js:"start"`
	Thumbnails shared.Thumbnails
}

func (x *Chapter) FromObject(vm *sobek.Runtime, obj *sobek.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("thumbnails"), &x.Thumbnails); err != nil {
		return err
	}
	return nil
}

type MusicMetadata struct {
	ImageUrl string `js:"imageUrl"`
	Title    string `js:"title"`
	Artist   string `js:"artist"`
	Album    Option[string]
}

func (x *MusicMetadata) FromObject(vm *sobek.Runtime, obj *sobek.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("album"), &x.Album); err != nil {
		return err
	}
	return nil
}
