//go:build go1.7
// +build go1.7

package packages

import (
	"bytes"
	"reflect"

	"github.com/kingecg/anko/env"
)

func bytesGo17() {
	env.Packages["bytes"]["ContainsRune"] = reflect.ValueOf(bytes.ContainsRune)
}
