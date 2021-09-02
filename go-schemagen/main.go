//go:build !windows
// +build !windows

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/filecoin-project/lotus/api"
)

func main() {
	var t reflect.Type
	switch apiArg := os.Args[1]; apiArg {
	case "Common":
		var api struct{ api.Common }
		t = reflect.TypeOf(api)
	case "FullNode":
		var api struct{ api.FullNode }
		t = reflect.TypeOf(api)
	case "StorageMiner":
		var api struct{ api.StorageMiner }
		t = reflect.TypeOf(api)
	case "Gateway":
		var api struct{ api.Gateway }
		t = reflect.TypeOf(api)
	case "Wallet":
		var api struct{ api.Wallet }
		t = reflect.TypeOf(api)
	case "Worker":
		var api struct{ api.Worker }
		t = reflect.TypeOf(api)
	default:
		panic("Unknown API")
	}
	methods := make(map[string]interface{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		record := make(map[string]interface{})
		if m.Type.NumOut() >= 2 {
			firstOut := m.Type.Out(0)
			kind := firstOut.Kind()
			if kind.String() == "chan" {
				record["subscription"] = true
			}
		}
		methods[m.Name] = record
	}

	output, err := json.MarshalIndent(methods, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", output)
}
