// Package policies provides the policy markdown documents. This file marks it
// as a go package so that it's included in the distribution.
package policies

import (
	"path"
	"runtime"
)

// Path is the full path to this package.
var Path string

func init() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get `legal` files path")
	}
	Path = path.Dir(f)
}
