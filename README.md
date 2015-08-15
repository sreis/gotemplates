# gotemplates
Go template repository to use with [go template](https://github.com/ncw/gotemplate) tool.

**TODO** add godoc

## Install

```
go get -u github.com/ncw/gotemplate
go get -u github.com/sreis/gotemplates
```

## Provides

* Concurrent Map implementation in package `concurrentmap`

## Example

Create a file mapstringint.go with:

```go
package main

//go:generate gotemplate "github.com/sreis/gotemplates/concurrentmap" "MapStringInt(string, int)"
```

And use the new `MapStringInt` like so:

```go
package main

import "fmt"

func main() {
    cmap := NewMapStringInt()
    key := "foobar"
    value := 1337
    cmap.Set(key, value)
    v, ok :=  cmap.Get(key)
    fmt.Println("value ", v, ok)
}

```

Build with:

```
make
make generate
make test
```
