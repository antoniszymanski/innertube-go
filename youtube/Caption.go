// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

type Caption struct {
	// caption content
	Text string `js:"text"`
	// caption start time in milliseconds
	Start int64 `js:"start"`
	// caption duration in milliseconds
	Duration int64 `js:"duration"`
	// transcript end time in milliseconds
	End int64 `js:"end"`
}

type CaptionLanguage struct {
	// Caption language name
	Name string `js:"name"`
	// Caption language code
	Code string `js:"code"`
	// Whether this language is translatable
	IsTranslatable bool `js:"isTranslatable"`
	// Caption language url
	URL string `js:"url"`
}
