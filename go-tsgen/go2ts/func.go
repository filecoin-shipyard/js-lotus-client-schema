package go2ts

import (
	"fmt"
	"reflect"
	"unicode"
)

// FuncInfo is exported information about a golang func.
type FuncInfo struct {
	Name    string
	Params  []Param
	Returns string
}

// Param is a function parameter.
type Param struct {
	Name string
	Type string
}

var paramNames = map[reflect.Type]string{
	reflect.TypeOf(false):      "bool",
	reflect.TypeOf(0):          "int",
	reflect.TypeOf(int8(0)):    "int",
	reflect.TypeOf(int16(0)):   "int",
	reflect.TypeOf(int32(0)):   "int",
	reflect.TypeOf(int64(0)):   "int",
	reflect.TypeOf(uint(0)):    "uint",
	reflect.TypeOf(uint8(0)):   "uint",
	reflect.TypeOf(uint16(0)):  "uint",
	reflect.TypeOf(uint32(0)):  "uint",
	reflect.TypeOf(uint64(0)):  "uint",
	reflect.TypeOf(float32(0)): "num",
	reflect.TypeOf(float64(0)): "num",
	reflect.TypeOf(""):         "str",
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// appendParam appends a parameter to the list of method parameters ensuring it
// has a unique name within the parameter set.
func (finfo *FuncInfo) appendParam(p Param) {
	n := 0
	for {
		var name string
		if n == 0 {
			name = p.Name
		} else {
			name = fmt.Sprintf("%s%d", p.Name, n)
		}
		var exists bool
		for _, ep := range finfo.Params {
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
	finfo.Params = append(finfo.Params, p)
}
