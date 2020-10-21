package go2ts

import (
	"fmt"
	"reflect"
	"strings"
)

var primitives = map[reflect.Type]string{
	reflect.TypeOf(false):      "boolean",
	reflect.TypeOf(0):          "number",
	reflect.TypeOf(int8(0)):    "number",
	reflect.TypeOf(int16(0)):   "number",
	reflect.TypeOf(int32(0)):   "number",
	reflect.TypeOf(int64(0)):   "number",
	reflect.TypeOf(uint(0)):    "number",
	reflect.TypeOf(uint8(0)):   "number",
	reflect.TypeOf(uint16(0)):  "number",
	reflect.TypeOf(uint32(0)):  "number",
	reflect.TypeOf(uint64(0)):  "number",
	reflect.TypeOf(float32(0)): "number",
	reflect.TypeOf(float64(0)): "number",
	reflect.TypeOf(uintptr(0)): "number",
	reflect.TypeOf(""):         "string",
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Converter will convert a golang reflect.Type to Typescript type string.
type Converter struct {
	types      map[reflect.Type]string
	paramNames map[reflect.Type]string
	// OnConvert is called when a type is converted but NOT present in the types
	// table. It is safe (and expected) that Converter.AddTypes is called from
	// this handler so that discovered types can be included in a converted type.
	OnConvert func(reflect.Type, string)
}

// NewConverter creates a new converter instance with primatives added.
func NewConverter() *Converter {
	c := Converter{
		types:      make(map[reflect.Type]string),
		paramNames: make(map[reflect.Type]string),
		OnConvert:  func(reflect.Type, string) {},
	}
	c.AddTypes(primitives)
	c.AddParamNames(paramNames)
	return &c
}

// AddTypes adds custom types.
func (c *Converter) AddTypes(customTypes map[reflect.Type]string) {
	for k, v := range customTypes {
		c.types[k] = v
	}
}

// AddParamNames adds custom function parameter names for types.
func (c *Converter) AddParamNames(customParamNames map[reflect.Type]string) {
	for k, v := range customParamNames {
		c.paramNames[k] = v
	}
}

// Convert takes a golang reflect.Type and returns a Typescript type string.
//
// Notes:
//
// chan is converted to AsyncIterable.
//
// Assumes functions/methods are async so return values are all Promise<T>
// and errors assumed to be thrown not returned.
//
// If a function returns multiple values they are returned as an array.
//
// Context in function params is ignored.
//
// Recursion is NOT supported.
func (c *Converter) Convert(t reflect.Type) (ts string) {
	ts, ok := c.types[t]
	if ok {
		return
	}

	kind := t.Kind()

	// Handle type aliases
	for t, s := range primitives {
		if t.Kind() == kind {
			ts = s
			return
		}
	}

	defer func() { c.OnConvert(t, ts) }()

	if kind == reflect.Ptr {
		ts = c.convert(t.Elem())
	} else if kind == reflect.Chan {
		ts = fmt.Sprintf("AsyncIterable<%s>", c.convert(t.Elem()))
	} else if kind == reflect.Func {
		ts = c.ConvertFunc(t)
	} else if kind == reflect.Struct {
		ts = c.ConvertStruct(t)
	} else if kind == reflect.Slice {
		ts = fmt.Sprintf("Array<%s>", c.convert(t.Elem()))
	} else if kind == reflect.Map {
		ts = fmt.Sprintf("{ [k: string]: %s }", c.convert(t.Elem()))
	} else if kind == reflect.Interface {
		ts = "any"
	} else {
		panic(fmt.Errorf("unhandled type: %v (%s)", t, t.Kind()))
	}
	return
}

func (c *Converter) convert(t reflect.Type) string {
	ts := c.Convert(t)
	// re-check against types: OnConvert may have called AddTypes
	uts, ok := c.types[t]
	if ok {
		return uts
	}
	return ts
}

// extractFunc extracts type inforamtion about a function.
func (c *Converter) extractFunc(t reflect.Type, isMethod bool) *FuncInfo {
	finfo := FuncInfo{Name: t.Name(), Returns: "Promise<void>"}
	if t.NumOut() > 0 {
		var rets []string
		for i := 0; i < t.NumOut(); i++ {
			out := t.Out(i)
			if i == t.NumOut()-1 && out.Implements(errorType) {
				break // skip last param if error
			}
			rets = append(rets, c.convert(out))
		}

		// If only 1 value just return it, if more than 1 we need to wrap in array.
		if len(rets) == 1 {
			if t.Out(0).Kind() == reflect.Chan {
				finfo.Returns = rets[0]
			} else {
				finfo.Returns = fmt.Sprintf("Promise<%s>", rets[0])
			}
		} else if len(rets) > 1 {
			finfo.Returns = fmt.Sprintf("Promise<[%s]>", strings.Join(rets, ", "))
		}
	}

	start := 0
	if isMethod {
		start = 1 // first argument is receiver, so skip over this
	}
	for i := start; i < t.NumIn(); i++ {
		in := t.In(i)
		// skip context if method takes one
		if in.Name() == "Context" {
			continue
		}
		p := Param{c.paramName(in), c.convert(in)}
		finfo.appendParam(p)
	}

	return &finfo
}

// paramName attempts to find a name for a function parameter.
func (c *Converter) paramName(t reflect.Type) string {
	name, ok := c.paramNames[t]
	if ok {
		return name
	}
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		return c.paramName(t.Elem())
	}
	name = t.Name()
	if name == "" {
		name = "_"
	} else {
		if isUpper(name) {
			name = strings.ToLower(name)
		} else {
			name = strings.ToLower(name[0:1]) + name[1:]
		}
	}
	return name
}

// ConvertMethod converts a method type to a typescript declaration.
func (c *Converter) ConvertMethod(m reflect.Method) string {
	finfo := c.extractFunc(m.Type, true)
	var params []string
	for _, p := range finfo.Params {
		params = append(params, fmt.Sprintf("%s: %s", p.Name, p.Type))
	}
	return fmt.Sprintf("%s (%s): %s", m.Name, strings.Join(params, ", "), finfo.Returns)
}

// ConvertFunc converts a function type to a typescript declaration.
func (c *Converter) ConvertFunc(t reflect.Type) string {
	finfo := c.extractFunc(t, false)
	var params []string
	for _, p := range finfo.Params {
		params = append(params, fmt.Sprintf("%s: %s", p.Name, p.Type))
	}
	return fmt.Sprintf("(%s) => %s", strings.Join(params, ", "), finfo.Returns)
}

// extractStruct extracts typescript type information about a struct.
func (c *Converter) extractStruct(t reflect.Type) *StructInfo {
	sinfo := StructInfo{Name: t.Name()}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !isUpper(f.Name[0:1]) {
			continue
		}
		sinfo.Fields = append(sinfo.Fields, Field{Name: f.Name, Type: c.convert(f.Type)})
	}
	return &sinfo
}

// ConvertStruct converts a struct to a typescript declaration.
func (c *Converter) ConvertStruct(t reflect.Type) string {
	sinfo := c.extractStruct(t)
	if len(sinfo.Fields) == 0 {
		return "{}"
	}
	var fields []string
	for _, f := range sinfo.Fields {
		fields = append(fields, fmt.Sprintf("%s: %s", f.Name, f.Type))
	}
	return fmt.Sprintf("{ %s }", strings.Join(fields, ", "))
}
