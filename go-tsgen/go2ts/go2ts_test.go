package go2ts

import (
	"context"
	"reflect"
	"testing"
)

func expect(t *testing.T, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Fatalf("expected \"%s\" to equal \"%s\"", expected, actual)
	}
}

func typ(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

type User struct{ Name string }
type Nested struct{ Owner User }

func (*Nested) Method(arg string) {}

func TestPrimitives(t *testing.T) {
	c := NewConverter()
	expect(t, c.Convert(typ("")), "string")
	expect(t, c.Convert(typ(138)), "number")
	expect(t, c.Convert(typ(int64(138))), "number")
	expect(t, c.Convert(typ(true)), "boolean")
}

func TestStructs(t *testing.T) {
	c := NewConverter()
	expect(t, c.Convert(typ(User{})), "{ Name: string }")
	expect(t, c.Convert(typ(&User{})), "{ Name: string }")
	expect(t, c.Convert(typ(Nested{})), "{ Owner: { Name: string } }")
}

func TestFuncs(t *testing.T) {
	c := NewConverter()
	// params
	expect(t, c.Convert(typ(func() {})), "() => Promise<void>")
	expect(t, c.Convert(typ(func(string, int, bool) {})), "(str: string, int: number, bool: boolean) => Promise<void>")
	expect(t, c.Convert(typ(func(struct{}) {})), "(_: {}) => Promise<void>")
	expect(t, c.Convert(typ(func(struct{ ID string }) {})), "(_: { ID: string }) => Promise<void>")
	expect(t, c.Convert(typ(func(User) {})), "(user: { Name: string }) => Promise<void>")
	expect(t, c.Convert(typ(func(*User) {})), "(_: { Name: string }) => Promise<void>")
	expect(t, c.Convert(typ(func(Nested) {})), "(nested: { Owner: { Name: string } }) => Promise<void>")
	// returns
	expect(t, c.Convert(typ(func() string { return "foo" })), "() => Promise<string>")
	expect(t, c.Convert(typ(func() User { return User{} })), "() => Promise<{ Name: string }>")
	expect(t, c.Convert(typ(func() *User { return nil })), "() => Promise<{ Name: string }>")
	expect(t, c.Convert(typ(func() error { return nil })), "() => Promise<void>")
	expect(t, c.Convert(typ(func() (string, error) { return "", nil })), "() => Promise<string>")
	expect(t, c.Convert(typ(func() (string, string, error) { return "", "", nil })), "() => Promise<[string, string]>")
	expect(t, c.Convert(typ(func() chan string { return nil })), "() => AsyncIterable<string>")
	// ignore context
	expect(t, c.Convert(typ(func(context.Context, string) {})), "(str: string) => Promise<void>")
	// methods
	n := Nested{}
	m, _ := typ(&n).MethodByName("Method")
	expect(t, c.ConvertMethod(m), "Method (str: string): Promise<void>")
}

func TestSlices(t *testing.T) {
	c := NewConverter()
	expect(t, c.Convert(typ([]string{})), "Array<string>")
	expect(t, c.Convert(typ([]*User{})), "Array<{ Name: string }>")
}

func TestMaps(t *testing.T) {
	c := NewConverter()
	expect(t, c.Convert(typ(map[string]int{})), "{ [k: string]: number }")
	expect(t, c.Convert(typ(map[string]User{})), "{ [k: string]: { Name: string } }")
}
