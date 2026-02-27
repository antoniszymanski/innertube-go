// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/dop251/goja"
)

type Post struct {
	ID string `js:"id"`
	// The channel who posted this post
	Channel BaseChannel
	// The content of this post
	Content string `js:"content"`
	// Timestamp
	Timestamp string `js:"timestamp"`
	// Vote count like '1.2K likes'
	VoteCount string `js:"voteCount"`
}

func (x *Post) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("channel"), &x.Channel); err != nil {
		return err
	}
	return nil
}
