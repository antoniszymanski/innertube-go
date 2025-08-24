// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"reflect"

	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type BaseVideo struct {
	ID string `js:"id"`
	// The title of this video
	Title string `js:"title"`
	// Thumbnails of the video with different sizes
	Thumbnails shared.Thumbnails `js:"thumbnails"`
	// The description of this video
	Description string `js:"description"`
	// The channel that uploaded this video
	Channel BaseChannel `js:"channel"`
	// The date this video is uploaded at
	UploadDate string `js:"uploadDate"`
	// How many views does this video have, None if the view count is hidden
	ViewCount utils.Option[int64]
	// How many likes does this video have, None if the like count is hidden
	LikeCount utils.Option[int64]
	// Whether this video is a live content or not
	IsLiveContent bool `js:"isLiveContent"`
	// The tags of this video
	Tags []string `js:"tags"`
	// Continuable of videos / playlists related to this video
	Related Continuable[VideoRelated] `js:"related"`
	// Captions helper class of this video (if caption exists in this video)
	Captions VideoCaptions
}

func (x *BaseVideo) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("thumbnails"), &x.Thumbnails); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("channel"), &x.Channel); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("viewCount"), &x.ViewCount); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("likeCount"), &x.LikeCount); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("related"), &x.Related); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("captions"), &x.Captions); err != nil {
		return err
	}
	return nil
}

type VideoCaptions struct {
	vm   *goja.Runtime
	this *goja.Object
	// List of available languages for this video
	Languages []CaptionLanguage `js:"languages"`
}

func (x *VideoCaptions) IsZero() bool {
	if x == nil {
		return true
	}
	return reflect.ValueOf(x).Elem().IsZero()
}

func (x *VideoCaptions) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	x.vm = vm
	x.this = obj
	return vm.ExportTo(obj, x)
}

// Get captions of a specific language or a translation of a specific language
func (x *VideoCaptions) Get(languageCode string, translationLanguageCode string) ([]Caption, error) {
	args := make([]goja.Value, 2)
	if languageCode != "" {
		args[0] = x.vm.ToValue(languageCode)
	} else {
		args[0] = goja.Undefined()
	}
	if translationLanguageCode != "" {
		args[1] = x.vm.ToValue(translationLanguageCode)
	} else {
		args[1] = goja.Undefined()
	}

	var result []Caption
	val, err := utils.CallAsync(x.vm, x.this, "get", args...)
	if err != nil {
		return nil, err
	}
	if goja.IsUndefined(val) {
		return nil, ErrCaptionsNotFound{}
	}

	if err := x.vm.ExportTo(val, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ErrCaptionsNotFound struct{}

func (e ErrCaptionsNotFound) Error() string {
	return "captions not found for the specified language"
}
