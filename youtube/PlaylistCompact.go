// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/antoniszymanski/innertube-go/shared"
	"github.com/dop251/goja"
)

type PlaylistCompact struct {
	ID string `js:"id"`
	// The playlist's title
	Title string `js:"title"`
	// Thumbnails of the playlist with different sizes
	Thumbnails shared.Thumbnails
	// The channel that made this playlist
	Channel BaseChannel `js:"channel"`
	// How many videos in this playlist
	VideoCount int64 `js:"videoCount"`
}

func (x *PlaylistCompact) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("thumbnails"), &x.Thumbnails); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("channel"), &x.Channel); err != nil {
		return err
	}
	return nil
}
