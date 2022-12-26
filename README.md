# go\_references

A library to get variable references from Go source code.

## Usage

You can use this package like following code.

```main.go
package main

import (
	"fmt"

	references "github.com/akif999/go_references"
)

func main() {
	references := references.New()
	references.ParseFile("./testdata/a.go")
	for _, ref := range references.Refs {
		fmt.Printf("  %s: %d\n", ref.Name, ref.Row)
	}
}
```

## Installation

`go install .github.com/akif999/go_references`

## TODO

* Implementing parse processing
    * parse process (DONE)
    * exclude Keywords (e.g int, string, make...)
        * builtins (int, string, panic...) (DONE)
        * others
* Explore APIs (DONE)
* Add Context(under consideration)
    * ex. AssignStmt, StarExpr..
* add tests (to `testdata/a.go`)

## Author

Akifumi Kitabatake(user name is "akif" or "akif999")
