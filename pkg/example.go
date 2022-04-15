// Copyright 2022 bonzai-example Authors
// SPDX-License-Identifier: Apache-2.0

// Package example provides high-level functions that are called from
// the Go Bonzai branch of the same name providing universal access to
// the core functionality.
package example

import "fmt"

func Foo(a string) { fmt.Println("foo " + a) }
