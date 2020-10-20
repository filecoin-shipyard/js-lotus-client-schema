package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// typeMap maps lotus types to Typescript type names (and a parameter name)
// Cid is { '/': string }
var typeMap = map[reflect.Type]param{
	reflect.TypeOf(""):                             {"str", "string"},
	reflect.TypeOf(42):                             {"num", "number"},
	reflect.TypeOf(uint64(42)):                     {"num", "number"},
	reflect.TypeOf(true):                           {"bool", "boolean"},
	reflect.TypeOf(byte(7)):                        {"byte", "number"},
	reflect.TypeOf(crypto.DomainSeparationTag(42)): {"tag", "number"},
	reflect.TypeOf(abi.ChainEpoch(42)):             {"epoch", "number"},
	reflect.TypeOf([]byte{}):                       {"bytes", "string"}, // base64 encoded
	reflect.TypeOf(cid.Cid{}):                      {"cid", "Cid"},
	reflect.TypeOf(address.Address{}):              {"addr", "string"},
	reflect.TypeOf(types.TipSetKey{}):              {"tipSetKey", "Cid[]"},
	reflect.TypeOf((*types.BeaconEntry)(nil)):      {"beaconEntry", "{ Round: number, Data: string }"},
	reflect.TypeOf((*types.BlockHeader)(nil)):      {"blockHeader", "any"}, // TODO
	reflect.TypeOf((*api.BlockMessages)(nil)):      {"blockMsgs", "any"},   // TODO
	reflect.TypeOf((*api.IpldObject)(nil)):         {"ipldObj", "{ Cid: Cid, Obj: any }"},
	reflect.TypeOf((*types.TipSet)(nil)):           {"tipSet", "any"}, // TODO
	reflect.TypeOf((*types.Message)(nil)):          {"msg", "any"},    // TODO
	reflect.TypeOf(abi.Randomness{}):               {"randomness", "string"},
	reflect.TypeOf([]auth.Permission{}):            {"permissions", "string[]"},
	reflect.TypeOf([]api.Message{}):                {"msg", "Array<{ Cid: Cid, Message: any }>"}, // TODO
	reflect.TypeOf([]*types.MessageReceipt{}):      {"msgReceipt", "Array<{ ExitCode: number, Return: string, GasUsed: number }>"},
	reflect.TypeOf([]*api.HeadChange{}):            {"headChange", "Array<{ Type: string, Val: any }>"}, // TODO
	reflect.TypeOf(make(<-chan []byte)):            {"bytes", "AsyncIterable<string>"},                  // What it should be...?
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// toParam returns a Typescript type param or an "any" param if unknown.
func toParam(t reflect.Type) param {
	p, ok := typeMap[t]
	if ok {
		return p
	}
	name := t.Name()
	if name == "" {
		name = "_"
	} else {
		if isUpper(name) {
			name = strings.ToLower(name)
		} else {
			name = strings.ToLower(name[0:1]) + name[1:]
		}
	}
	return param{name, "any"}
}

// param is the definition of a parameter to a method.
type param struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

// method is an API method definition.
type method struct {
	Params       []param `json:"params"`
	Returns      string  `json:"returns"`
	Subscription bool    `json:"subscription"`
}

// appendParam appends a parameter to the list of method parameters ensuring it
// has a unique name within the parameter set.
func (m *method) appendParam(p param) {
	n := 0
	for {
		var name string
		if n == 0 {
			name = p.Name
		} else {
			name = fmt.Sprintf("%s%d", p.Name, n)
		}
		var exists bool
		for _, ep := range m.Params {
			if ep.Name == name {
				exists = true
				break
			}
		}
		if !exists {
			p.Name = name
			break
		}
		n++
	}
	m.Params = append(m.Params, p)
}

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
	case "GatewayAPI":
		var api struct{ api.GatewayAPI }
		t = reflect.TypeOf(api)
	case "WalletAPI":
		var api struct{ api.WalletAPI }
		t = reflect.TypeOf(api)
	case "WorkerAPI":
		var api struct{ api.WorkerAPI }
		t = reflect.TypeOf(api)
	default:
		panic("Unknown API")
	}
	methods := make(map[string]method)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		var record method
		if m.Type.NumOut() >= 2 {
			firstOut := m.Type.Out(0)
			if firstOut.Kind() == reflect.Chan {
				record.Subscription = true
				// first param is the data handler function
				record.appendParam(handlerParam(firstOut))
				record.Returns = "[() => {}, Promise<void>]" // this is a weird API
			} else {
				record.Returns = fmt.Sprintf("Promise<%s>", toParam(firstOut).Type)
			}
		} else {
			record.Returns = "Promise<void>"
		}

		// first argument is receiver, so skip over this
		for i := 1; i < m.Type.NumIn(); i++ {
			t := m.Type.In(i)
			// skip context if method takes one
			if t.Name() == "Context" {
				continue
			}
			record.appendParam(toParam(t))
		}
		methods[m.Name] = record
	}

	output, err := json.Marshal(methods)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", output)
}

// handlerParam creates a handler function param based on the passed channel type.
func handlerParam(chanTyp reflect.Type) param {
	p := toParam(chanTyp)
	t := strings.Replace(strings.Replace(p.Type, "AsyncIterable<", "", 1), ">", "", 1)
	return param{Name: "handler", Type: fmt.Sprintf("(%s: %s) => {}", p.Name, t)}
}
