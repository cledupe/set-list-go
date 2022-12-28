# set-list-go
Collection Struct that contains no duplicate element


![example branch parameter](https://github.com/cledupe/set-list-go/actions/workflows/test.yml/badge.svg)


### Install
``go get github.com/cledupe/set-list-go``

### How to Use
```go
package main

import (
	"fmt"

	SetList "github.com/cledupe/set-list-go"
)

type Test struct {
	value int
}

func main() {
	setList := SetList.New[Test]()
	setList.InsertSlice([]Test{{value: 2}, {value: 1}, {value: 3}, {value: 2}})
	fmt.Println(setList.AsSlice())
}

```

#### Result

```sh
[{2} {1} {3}]
```