// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package testdata

import "time"

// User define an user of our test.
type User struct {
	ID       uint64
	Name     string
	Birthday *time.Time
}

// Test is a test struct.
type Test struct {
	Field      string
	IntegerPtr *int64
}
