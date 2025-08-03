// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package innertube

import (
	"strconv"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func atoi(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func appendInt[T constraints.Signed](dst []byte, i T) []byte {
	return strconv.AppendInt(dst, int64(i), 10)
}

func appendQuote(dst []byte, s string) []byte {
	if strconv.CanBackquote(s) {
		dst = append(dst, '`')
		dst = append(dst, s...)
		dst = append(dst, '`')
		return dst
	} else {
		return strconv.AppendQuoteToGraphic(dst, s)
	}
}

func quotedLen(s string) int {
	return 1 + len(s) + 1 // best-case scenario
}

func bytes2string(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
