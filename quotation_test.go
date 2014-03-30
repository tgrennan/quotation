// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quotation_test

import (
	"github.com/tgrennan/quotation"
	"testing"
)

// Returns index of first mismatched string,
// or -1 if all match.
func mismatch(got []string, want ...string) int {
	for i, s := range want {
		if got[i] != s {
			return i
		}
	}
	return -1
}

func Test(t *testing.T) {
	got := quotation.Fields(`foo "abc 123" bar 'hello world'`)
	if i := mismatch(got, "foo", "abc 123", "bar", "hello world"); i >= 0 {
		t.Fatalf("mismatch at %d of %q", i, got)
	}
}
