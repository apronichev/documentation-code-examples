// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sampler

import (
	"testing"

	"golang.org/x/text/language"
)

var helloTests = []struct {
	prefs []language.Tag
	text  string
}{
	{
		[]language.Tag{language.Make("en-US"), language.Make("fr")},
		"Hello, world.",
	},
	{
		[]language.Tag{language.Make("fr"), language.Make("en-US")},
		"Bonjour le monde.",
	},
}

func TestHello(t *testing.T) {
	for _, tt := range helloTests {
		text := Hello(tt.prefs...)
		if text != tt.text {
			t.Errorf("Hello(%v) = %q, want %q", tt.prefs, text, tt.text)
		}
	}
}
