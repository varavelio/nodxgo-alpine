# nodxgo-alpine

<a href="https://pkg.go.dev/github.com/varavelio/nodxgo-alpine">
  <img src="https://pkg.go.dev/badge/github.com/varavelio/nodxgo-alpine" alt="Go Reference"/>
</a>
<a href="https://goreportcard.com/report/varavelio/nodxgo-alpine">
  <img src="https://goreportcard.com/badge/varavelio/nodxgo-alpine" alt="Go Report Card"/>
</a>
<a href="https://github.com/varavelio/nodxgo-alpine/releases/latest">
  <img src="https://img.shields.io/github/release/varavelio/nodxgo-alpine.svg" alt="Release Version"/>
</a>
<a href="LICENSE">
  <img src="https://img.shields.io/github/license/varavelio/nodxgo-alpine.svg" alt="License"/>
</a>
<a href="https://github.com/varavelio/nodxgo-alpine">
  <img src="https://img.shields.io/github/stars/varavelio/nodxgo-alpine?style=flat&label=github+stars"/>
</a>

---

nodxgo-alpine provides a Go-based integration for
[Alpine.js](https://alpinejs.dev) within the
[NodX Go](https://github.com/varavelio/nodxgo) template engine. It simplifies
the generation of dynamic HTML elements and attributes using Alpine.js
directives.

## Installation

To install the package, run:

```sh
# Go 1.22 or later is required
go get github.com/varavelio/nodxgo-alpine
```

## Usage

### Example: Using Alpine.js Directives

#### `x-data` Directive

```go
node := nodx.Div(
	alpine.XData("{ count: 0 }"),
)
fmt.Println(node)
```

**Output:**

```html
<div x-data="{ count: 0 }"></div>
```

#### `x-bind` Directive

```go
node := nodx.Button(
	alpine.XBind("disabled", "isDisabled"),
	nodx.Text("Click Me"),
)
fmt.Println(node)
```

**Output:**

```html
<button x-bind:disabled="isDisabled">Click Me</button>
```

### Example: Rendering a Template

```go
package main

import (
	"fmt"
	nodx "github.com/varavelio/nodxgo"
	alpine "github.com/varavelio/nodxgo-alpine"
)

func main() {
	node := alpine.Template(
		nodx.Div(nodx.Text("Hello, World!")),
	)
	fmt.Println(node)
}
```

**Output:**

```html
<template>
	<div>Hello, World!</div>
</template>
```

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to
improve the project.

---

For more details, check out the [Alpine.js documentation](https://alpinejs.dev)
and [NodX Go](https://github.com/varavelio/nodxgo).
