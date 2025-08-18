// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Comment struct {
	ID string `js:"id"`
	// The video this comment belongs to
	Video Video
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
	// Whether the comment is pinned
	IsPinned bool `js:"isPinned"`
	// Reply count of this comment
	ReplyCount int64 `js:"replyCount"`
	// Continuable of replies in this comment
	Replies CommentReplies
	// URL to the video with this comment being highlighted (appears on top of the comment section)
	URL string `js:"url"`
}

type CommentReplies = Continuable[Reply]

func (x *Comment) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("video"), &x.Video); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("author"), &x.Author); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("replies"), &x.Replies); err != nil {
		return err
	}
	return nil
}
