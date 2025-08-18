// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils_test

import (
	"strconv"
	"testing"

	"github.com/antoniszymanski/innertube-go/common/utils"
)

const n = uint64(100_000)

func TestIndicesSeq(t *testing.T) {
	i := uint64(0)
	for j, actual := range utils.IndicesSeq(n) {
		if actual != strconv.FormatUint(i, 10) {
			t.FailNow()
		}
		if i != j {
			t.FailNow()
		}
		i++
	}
}

func BenchmarkIndicesSeq(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		for range utils.IndicesSeq(n) {
		}
	}
}

func BenchmarkFormatUint(b *testing.B) {
	b.ReportAllocs()
	for range b.N {
		for i := range n {
			strconv.FormatUint(i, 10)
		}
	}
}

func BenchmarkAppendUint(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	for range b.N {
		for i := range n {
			buf = strconv.AppendUint(buf[:0], i, 10)
		}
	}
}
