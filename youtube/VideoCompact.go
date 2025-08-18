// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type VideoCompact struct {
	ID string `js:"id"`
	// The title of the video
	Title string `js:"title"`
	// Thumbnails of the video with different sizes
	Thumbnails shared.Thumbnails
	// Description of the video (not a full description, rather a preview / snippet)
	Description string `js:"description"`
	// The duration of this video in second, None if the video is live
	Duration utils.Option[int64]
	// Whether this video is a live now or not
	IsLive bool `js:"isLive"`
	// Whether this video is a shorts or not
	IsShort bool `js:"isShort"`
	// The channel who uploads this video
	Channel BaseChannel
	// The date this video is uploaded at
	UploadDate string `js:"uploadDate"`
	// How many views does this video have, None if the view count is hidden
	ViewCount utils.Option[int64]
	// Whether this video is private / deleted or not, only useful in playlist's videos
	IsPrivateOrDeleted bool `js:"isPrivateOrDeleted"`
}

func (x *VideoCompact) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("thumbnails"), &x.Thumbnails); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("duration"), &x.Duration); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("channel"), &x.Channel); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("viewCount"), &x.ViewCount); err != nil {
		return err
	}
	return nil
}
