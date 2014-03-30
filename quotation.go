// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package quotation provides a string slicer for space separated words and
quoted phrases.
*/
package quotation

import "strings"

type splitter struct {
	doubled, singled, escaped bool
}

// This is like strings.Fields() but retains whitespace within quoted phrases.
func Fields(s string) []string {
	a := strings.FieldsFunc(s, new(splitter).Fields)
	for i, s := range a {
		n := len(s) - 1
		r0, rn := s[0], s[n]
		if (r0 == '"' && rn == '"') || (r0 == '\'' && rn == '\'') {
			a[i] = s[1:n]
		}
	}
	return a
}

func (x *splitter) Fields(r rune) bool {
	switch r {
	case '\\':
		if x.escaped {
			x.escaped = false
		} else {
			x.escaped = true
		}
	case '"':
		if x.escaped {
			x.escaped = false
		} else if x.doubled {
			x.doubled = false
		} else {
			x.doubled = true
		}
	case '\'':
		if x.escaped {
			x.escaped = false
		} else if x.singled {
			x.singled = false
		} else {
			x.singled = true
		}
	case ' ':
		if x.escaped {
			x.escaped = false
		} else if !x.doubled && !x.singled {
			return true
		}
	default:
		if x.escaped {
			x.escaped = false
		}
	}
	return false
}
