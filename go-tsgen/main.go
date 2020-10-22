package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alanshaw/go2ts"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func main() {
	var api struct{ api.FullNode }

	t := reflect.TypeOf(api)
	c := go2ts.NewConverter()

	// weird/exceptional types
	c.AddTypes(map[reflect.Type]string{
		reflect.TypeOf(cid.Cid{}):         "Cid",
		reflect.TypeOf([]byte{}):          "string", // base64 encoded
		reflect.TypeOf(abi.Randomness{}):  "string", // base64 encoded
		reflect.TypeOf(types.BigInt{}):    "string",
		reflect.TypeOf(address.Address{}): "string",
		reflect.TypeOf(types.TipSetKey{}): "Cid[]",
		// workaround for recursive Subcalls property on ExecutionTrace
		reflect.TypeOf([]types.ExecutionTrace{}): "ExecutionTrace[]",
	})

	c.AddParamNames(map[reflect.Type]string{
		reflect.TypeOf([]byte{}):       "bytes",
		reflect.TypeOf(types.BigInt{}): "bigInt",
	})

	decls := []string{"declare type Cid = { '/': string }"}
	exports := []string{"LotusRPC"}

	// for when the exported struct is not the struct used
	exportOverrides := map[reflect.Type]interface{}{
		reflect.TypeOf(types.TipSet{}): types.ExpTipSet{},
	}

	typeNames := map[string]reflect.Type{}
	uniqueTypeName := func(t reflect.Type) (name string) {
		n := 0
		for {
			if n == 0 {
				name = t.Name()
			} else {
				name = fmt.Sprintf("%s%d", t.Name(), n)
			}
			ct, ok := typeNames[name]
			if ok && ct == t {
				break
			} else if !ok {
				typeNames[name] = t
				break
			}
			n++
		}
		return
	}

	c.OnConvert = func(t reflect.Type, ts string) {
		if t.Kind() != reflect.Struct || t.Name() == "" {
			return
		}

		ori, ok := exportOverrides[t]
		if ok {
			ts = c.Convert(reflect.TypeOf(ori))
		}

		name := uniqueTypeName(t)
		decls = append(decls, fmt.Sprintf("declare type %s = %s", name, ts))
		exports = append(exports, name)
		c.AddTypes(map[reflect.Type]string{t: name})
	}

	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		ts := c.ConvertMethod(m)
		if strings.HasPrefix(ts, "ID ") {
			ts = strings.Replace(ts, "ID ", "id ", 1)
		} else {
			ts = strings.ToLower(ts[0:1]) + ts[1:]
		}
		methods = append(methods, ts)
	}

	decls = append(decls, fmt.Sprintf("declare class LotusRPC {\n  %s\n}", strings.Join(methods, "\n  ")))
	decls = append(decls, fmt.Sprintf("export { %s }", strings.Join(exports, ", ")))

	fmt.Println(strings.Join(decls, "\n"))
}
