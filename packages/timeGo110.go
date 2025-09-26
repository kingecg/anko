//go:build go1.10
// +build go1.10

package packages

import (
	"reflect"
	"time"

	"github.com/kingecg/anko/env"
)

func timeGo110() {
	env.Packages["time"]["LoadLocationFromTZData"] = reflect.ValueOf(time.LoadLocationFromTZData)
}
