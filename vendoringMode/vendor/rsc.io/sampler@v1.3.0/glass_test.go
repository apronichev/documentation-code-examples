// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sampler

import (
	"testing"

	"golang.org/x/text/language"
)

var glassTests = []struct {
	prefs []language.Tag
	text  string
}{
	{
		[]language.Tag{language.Make("en-US"), language.Make("fr")},
		"I can eat glass and it doesn't hurt me.",
	},
	{
		[]language.Tag{language.Make("fr"), language.Make("en-US")},
		"Je peux manger du verre, ça ne me fait pas mal.",
	},
}

func TestGlass(t *testing.T) {
	for _, tt := range glassTests {
		text := Glass(tt.prefs...)
		if text != tt.text {
			t.Errorf("Glass(%v) = %q, want %q", tt.prefs, text, tt.text)
		}
	}
}
