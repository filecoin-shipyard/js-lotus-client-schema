package go2ts

// StructInfo is exported information about a golang func.
type StructInfo struct {
	Name   string
	Fields []Field
	// TODO: methods?
}

// Field is a struct field.
type Field struct {
	Name string
	Type string
}
