// +build !go1.9

// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"github.com/mafredri/cdp/protocol"
	"github.com/mafredri/cdp/protocol/dom"
)

// StyleSheetHeader CSS stylesheet metainformation.
type StyleSheetHeader struct {
	StyleSheetID StyleSheetID         `json:"styleSheetId"`           // The stylesheet identifier.
	FrameID      protocol.PageFrameID `json:"frameId"`                // Owner frame identifier.
	SourceURL    string               `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL *string              `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       StyleSheetOrigin     `json:"origin"`                 // Stylesheet origin.
	Title        string               `json:"title"`                  // Stylesheet title.
	OwnerNode    *dom.BackendNodeID   `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool                 `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL *bool                `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool                 `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64              `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64              `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length       float64              `json:"length"`                 // Size of the content (in characters).
}