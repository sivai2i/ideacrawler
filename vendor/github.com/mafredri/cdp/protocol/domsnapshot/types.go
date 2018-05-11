// Code generated by cdpgen. DO NOT EDIT.

package domsnapshot

import (
	"github.com/mafredri/cdp/protocol/dom"
)

// InlineTextBox Details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
type InlineTextBox struct {
	BoundingBox         dom.Rect `json:"boundingBox"`         // The absolute position bounding box.
	StartCharacterIndex int      `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters       int      `json:"numCharacters"`       // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

// LayoutTreeNode Details of an element in the DOM tree with a LayoutObject.
type LayoutTreeNode struct {
	DOMNodeIndex    int             `json:"domNodeIndex"`              // The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	BoundingBox     dom.Rect        `json:"boundingBox"`               // The absolute position bounding box.
	LayoutText      *string         `json:"layoutText,omitempty"`      // Contents of the LayoutText, if any.
	InlineTextNodes []InlineTextBox `json:"inlineTextNodes,omitempty"` // The post-layout inline text nodes, if any.
	StyleIndex      *int            `json:"styleIndex,omitempty"`      // Index into the `computedStyles` array returned by `getSnapshot`.
	PaintOrder      *int            `json:"paintOrder,omitempty"`      // Global paint order index, which is determined by the stacking order of the nodes. Nodes that are painted together will have the same index. Only provided if includePaintOrder in getSnapshot was true.
}

// ComputedStyle A subset of the full ComputedStyle as defined by the request
// whitelist.
type ComputedStyle struct {
	Properties []NameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// NameValue A name/value pair.
type NameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}