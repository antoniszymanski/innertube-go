// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Reply struct {
	ID string `js:"id"`
	// The comment this reply belongs to
	Comment Comment
	// The video this reply belongs to
	Video Video `js:"video"`
	// The comment's author
	Author BaseChannel
	// The content of this comment
	Content string `js:"content"`
	// The publish date of the comment
	PublishDate string `js:"publishDate"`
	// How many likes does this comment have
	LikeCount int64 `js:"likeCount"`
	// Whether the comment is posted by the video uploader / owner
	IsAuthorChannelOwner bool `js:"isAuthorChannelOwner"`
}

func (x *Reply) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("comment"), &x.Comment); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("author"), &x.Author); err != nil {
		return err
	}
	return nil
}
