// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils

import "iter"

func IndicesSeq(n uint64) iter.Seq2[uint64, string] {
	return func(yield func(uint64, string) bool) {
		b := []byte("00000000000000000000") // 20 is uint64's max string length
		begin := len(b) - 1
		for i := range n {
			if !yield(i, string(b[begin:])) {
				return
			}
			for j := len(b) - 1; ; j-- {
				if b[j] < '9' {
					b[j]++
					break
				}
				b[j] = '0'
				if j == begin {
					if begin <= 0 {
						panic("unreachable") // overflow
					}
					begin--
					b[begin] = '1'
					break
				}
			}
		}
	}
}
