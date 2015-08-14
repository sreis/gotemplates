# gotemplates
Go template repository to use with [go template](https://github.com/ncw/gotemplate) tool.

## Install
``` go get github.com/sreis/gotemplates ```

## Provides

* Concurrent Map implementation in package `concurrentmap`
  * Methods
    * New()
    * Set(Key, Value)
    * Get(Key)
    * Remove(Key)

## Example

Create a file mapstringint.go with:

```

package main

//go:generate gotemplate "github.com/sreis/gotemplates/concurrentmap" "MapStringInt(string, int)"

```

And use the new MapStringInt like so:

```
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
    go generate
    go build
```
