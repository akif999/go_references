# go\_references

A library to get variable references from Go source code.

## Usage

You can use this package like following code.

```main.go
package main

import (
	"fmt"
	"log"

	references "github.com/akif999/go_references"
)

func main() {
	references := references.New()
	err := references.ParseFile("../testdata/a.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(references.File)
	for _, ref := range references.Refs {
		fmt.Printf("%s: %d\n", ref.Name, ref.Row)
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
* Add identifier map(under consideration)
    * e.g Hoge {{Line:11, Context: Assign}...}
* add tests (to `testdata/a.go`)

## License

MIT

## Author

Akifumi Kitabatake(user name is "akif" or "akif999")
