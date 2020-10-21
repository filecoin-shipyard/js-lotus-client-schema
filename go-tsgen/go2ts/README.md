# go2ts

Convert golang to Typescript declarations.

Note:
* `chan T` is converted to `AsyncIterable<T>`.
* Assumes functions/methods are async so return values are all `Promise<T>` and errors assumed to be thrown not returned.
* If a function returns multiple values they are returned as an array.
* Context in function params is ignored.
* Recursion is NOT supported.

## Install

```
go get github.com/alanshaw/go2ts
```

## Usage

### Example

```go
package main

import (
  "reflect"
  "github.com/alanshaw/go2ts"
)

type User struct {
  Name string
}

func main () {
  c := go2ts.NewConverter()

  c.Convert(reflect.TypeOf("")) // string
  c.Convert(reflect.TypeOf(User{})) // { Name: string }
  c.Convert(reflect.TypeOf(func(string, int, bool) User { return nil })
  // (str: string, int: number, bool: boolean) => Promise<{ Name: string }>

  // Add custom type declarations
  c.AddTypes(map[reflect.Type]string{
    reflect.TypeOf(User{}): "User"
  })

  // Output now includes "User" instead of { Name: string }
  c.Convert(reflect.TypeOf(map[string]User{})) // { [k: string]: User }
}
```