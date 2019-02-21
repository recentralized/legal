// Package src provides the markdown documents. This file marks it
// as a go package so that it's included in the distribution.
package src

import (
	"path"
	"runtime"
)

var pkgPath string

// GetPath returns the path to this package.
func GetPath() string {
	if pkgPath == "" {
		_, f, _, ok := runtime.Caller(0)
		if !ok {
			panic("failed to get `legal` files path")
		}
		pkgPath = path.Dir(f)
	}
	return pkgPath
}
