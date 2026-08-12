<h1 align="center">Alpine.js for NodX Go</h1>

<p align="center">
	<a href="https://github.com/varavelio/nodxgo-alpine/actions">
		<img src="https://github.com/varavelio/nodxgo-alpine/actions/workflows/ci.yaml/badge.svg" alt="CI status"/>
	</a>
	<a href="https://pkg.go.dev/github.com/varavelio/nodxgo-alpine">
		<img src="https://pkg.go.dev/badge/github.com/varavelio/nodxgo-alpine" alt="Go Reference"/>
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
</p>

<p align="center">
  <a href="https://varavel.com">
    <img src="https://cdn.jsdelivr.net/gh/varavelio/brand@1.0.0/dist/badges/project.svg" alt="A Varavel project"/>
  </a>
</p>

---

nodxgo-alpine provides a Go-based integration for
[Alpine.js](https://alpinejs.dev) within the
[NodX Go](https://github.com/varavelio/nodxgo) template engine. It simplifies
the generation of dynamic HTML elements and attributes using Alpine.js
directives.

## Installation

To install the package, run:

```sh
# Tested on Go 1.22 and later
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

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to
improve the project.

---

For more details, check out the [Alpine.js documentation](https://alpinejs.dev)
and [NodX Go](https://github.com/varavelio/nodxgo).
