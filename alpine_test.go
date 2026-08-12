package alpine_test

import (
	"fmt"

	nodx "github.com/varavelio/nodxgo"
	alpine "github.com/varavelio/nodxgo-alpine"
)

func ExampleX() {
	node := nodx.Div(
		alpine.X("if", "true"),
	)
	fmt.Println(node)
	// Output: <div x-if="true"></div>
}

func ExampleXData() {
	node := nodx.Div(
		alpine.XData(`{}`),
	)
	fmt.Println(node)
	// Output: <div x-data="{}"></div>
}

func ExampleXInit() {
	node := nodx.Div(
		alpine.XInit("console.log(true)"),
	)
	fmt.Println(node)
	// Output: <div x-init="console.log(true)"></div>
}

func ExampleXShow() {
	node := nodx.Div(
		alpine.XShow("true"),
	)
	fmt.Println(node)
	// Output: <div x-show="true"></div>
}

func ExampleXBind() {
	node := nodx.Div(
		alpine.XBind("attr", "value"),
	)
	fmt.Println(node)
	// Output: <div x-bind:attr="value"></div>
}

func ExampleXOn() {
	node := nodx.Div(
		alpine.XOn("click", "console.log(true)"),
	)
	fmt.Println(node)
	// Output: <div x-on:click="console.log(true)"></div>
}

func ExampleXText() {
	node := nodx.Div(
		alpine.XText("value"),
	)
	fmt.Println(node)
	// Output: <div x-text="value"></div>
}

func ExampleXHTML() {
	node := nodx.Div(
		alpine.XHTML("value"),
	)
	fmt.Println(node)
	// Output: <div x-html="value"></div>
}

func ExampleXModel() {
	node := nodx.Div(
		alpine.XModel("value"),
	)
	fmt.Println(node)
	// Output: <div x-model="value"></div>
}

func ExampleXModelable() {
	node := nodx.Div(
		alpine.XModelable("value"),
	)
	fmt.Println(node)
	// Output: <div x-modelable="value"></div>
}

func ExampleXFor() {
	node := nodx.Div(
		alpine.XFor("color in colors"),
	)
	fmt.Println(node)
	// Output: <div x-for="color in colors"></div>
}

func ExampleXTransition() {
	node := nodx.Div(
		alpine.XTransition(),
	)
	fmt.Println(node)
	// Output: <div x-transition></div>
}

func ExampleXEffect() {
	node := nodx.Div(
		alpine.XEffect("console.log(label)"),
	)
	fmt.Println(node)
	// Output: <div x-effect="console.log(label)"></div>
}

func ExampleXIgnore() {
	node := nodx.Div(
		alpine.XIgnore(),
	)
	fmt.Println(node)
	// Output: <div x-ignore></div>
}

func ExampleXRef() {
	node := nodx.Div(
		alpine.XRef("value"),
	)
	fmt.Println(node)
	// Output: <div x-ref="value"></div>
}

func ExampleXCloak() {
	node := nodx.Div(
		alpine.XCloak(),
	)
	fmt.Println(node)
	// Output: <div x-cloak></div>
}

func ExampleXTeleport() {
	node := nodx.Div(
		alpine.XTeleport("body"),
	)
	fmt.Println(node)
	// Output: <div x-teleport="body"></div>
}

func ExampleXIf() {
	node := nodx.Div(
		alpine.XIf("true"),
	)
	fmt.Println(node)
	// Output: <div x-if="true"></div>
}

func ExampleXId() {
	node := nodx.Div(
		alpine.XId("value"),
	)
	fmt.Println(node)
	// Output: <div x-id="value"></div>
}
