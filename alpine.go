// Package alpine provides Alpine.js integration for NodX Go
//
//   - https://alpinejs.dev
//   - https://github.com/varavelio/nodxgo
package alpine

import nodx "github.com/varavelio/nodxgo"

// X is an attribute that renders a x-[key]="[value]" attribute.
func X(key string, value ...string) nodx.Node {
	return nodx.Attr("x-"+key, value...)
}

// XData is an attribute that renders a x-data="[value]" attribute.
//
// https://alpinejs.dev/directives/data
func XData(value string) nodx.Node {
	return X("data", value)
}

// XInit is an attribute that renders a x-init="[value]" attribute.
//
// https://alpinejs.dev/directives/init
func XInit(value string) nodx.Node {
	return X("init", value)
}

// XShow is an attribute that renders a x-show="[vlue]" attribute.
//
// https://alpinejs.dev/directives/show
func XShow(value string) nodx.Node {
	return X("show", value)
}

// XBind is an attribute that renders a x-bind:[targetAttr]="[value]" attribute.
//
// https://alpinejs.dev/directives/bind
func XBind(targetAttr, value string) nodx.Node {
	return X("bind:"+targetAttr, value)
}

// XOn is an attribute that renders a x-on:[targetEvent]="[value]" attribute.
//
// https://alpinejs.dev/directives/on
func XOn(targetEvent, value string) nodx.Node {
	return X("on:"+targetEvent, value)
}

// XText is an attribute that renders a x-text="[value]" attribute.
//
// https://alpinejs.dev/directives/text
func XText(value string) nodx.Node {
	return X("text", value)
}

// XHTML is an attribute that renders a x-html="[value]" attribute.
//
// https://alpinejs.dev/directives/html
func XHTML(value string) nodx.Node {
	return X("html", value)
}

// XModel is an attribute that renders a x-model="[value]" attribute.
//
// https://alpinejs.dev/directives/model
func XModel(value string) nodx.Node {
	return X("model", value)
}

// XModelable is an attribute that renders a x-modelable="[value]" attribute.
//
// https://alpinejs.dev/directives/modelable
func XModelable(value string) nodx.Node {
	return X("modelable", value)
}

// XFor is an attribute that renders a x-for="[value]" attribute.
//
// https://alpinejs.dev/directives/for
func XFor(value string) nodx.Node {
	return X("for", value)
}

// XTransition is an attribute that renders a x-transition attribute.
//
// https://alpinejs.dev/directives/transition
func XTransition() nodx.Node {
	return X("transition")
}

// XEffect is an attribute that renders a x-effect="[value]" attribute.
//
// https://alpinejs.dev/directives/effect
func XEffect(value string) nodx.Node {
	return X("effect", value)
}

// XIgnore is an attribute that renders a x-ignore attribute.
//
// https://alpinejs.dev/directives/ignore
func XIgnore() nodx.Node {
	return X("ignore")
}

// XRef is an attribute that renders a x-ref="[value]" attribute.
//
// https://alpinejs.dev/directives/ref
func XRef(value string) nodx.Node {
	return X("ref", value)
}

// XCloak is an attribute that renders a x-cloak attribute.
//
// https://alpinejs.dev/directives/cloak
func XCloak() nodx.Node {
	return X("cloak")
}

// XTeleport is an attribute that renders a x-teleport="[value]" attribute.
//
// https://alpinejs.dev/directives/teleport
func XTeleport(value string) nodx.Node {
	return X("teleport", value)
}

// XIf is an attribute that renders a x-if="[value]" attribute.
//
// https://alpinejs.dev/directives/if
func XIf(value string) nodx.Node {
	return X("if", value)
}

// XId is an attribute that renders a x-id="[value]" attribute.
//
// https://alpinejs.dev/directives/id
func XId(value string) nodx.Node {
	return X("id", value)
}
