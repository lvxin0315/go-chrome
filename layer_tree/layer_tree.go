package LayerTree

import (
	"github.com/mkenney/go-chrome/dom"
)

/*
CompositingReasonsParams represents LayerTree.compositingReasons parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
type CompositingReasonsParams struct {
	// The ID of the layer for which we want to get the reasons it was composited.
	LayerID LayerID `json:"layerId"`
}

/*
CompositingReasonsResult represents the result of calls to LayerTree.compositingReasons.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-compositingReasons
*/
type CompositingReasonsResult struct {
	// A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons []string `json:"compositingReasons"`
}

/*
LoadSnapshotParams represents LayerTree.loadSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
type LoadSnapshotParams struct {
	// An array of tiles composing the snapshot.
	Tiles []PictureTile `json:"tiles"`
}

/*
LoadSnapshotResult represents the result of calls to LayerTree.loadSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-loadSnapshot
*/
type LoadSnapshotResult struct {
	// The ID of the snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
MakeSnapshotParams represents LayerTree.makeSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
type MakeSnapshotParams struct {
	// The ID of the layer.
	LayerID LayerID `json:"layerId"`
}

/*
MakeSnapshotResult represents the result of calls to LayerTree.makeSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-makeSnapshot
*/
type MakeSnapshotResult struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
ProfileSnapshotParams represents LayerTree.profileSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
type ProfileSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount int `json:"minRepeatCount,omitempty"`

	// Optional. The minimum duration (in seconds) to replay the snapshot.
	MinDuration float64 `json:"minDuration,omitempty"`

	// Optional. The clip rectangle to apply when replaying the snapshot.
	//
	// This is an instance of DOM.Rect, but that doesn't omitempty correctly so it must be added
	// manually.
	ClipRect interface{} `json:"clipRect,omitempty"`
}

/*
ProfileSnapshotResult represents the result of calls to LayerTree.profileSnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-profileSnapshot
*/
type ProfileSnapshotResult struct {
	// The array of paint profiles, one per run.
	Timings []PaintProfile `json:"timings"`
}

/*
ReleaseSnapshotParams represents LayerTree.releaseSnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-releaseSnapshot
*/
type ReleaseSnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
ReplaySnapshotParams represents LayerTree.replaySnapshot parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
type ReplaySnapshotParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`

	// Optional. The first step to replay from (replay from the very start if not specified).
	FromStep int `json:"fromStep,omitempty"`

	// Optional. The last step to replay to (replay till the end if not specified).
	ToStep int `json:"toStep,omitempty"`

	// Optional. The scale to apply while replaying (defaults to 1).
	Scale float64 `json:"scale,omitempty"`
}

/*
ReplaySnapshotResult represents the result of calls to LayerTree.replaySnapshot.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-replaySnapshot
*/
type ReplaySnapshotResult struct {
	// A data: URL for resulting image.
	DataURL string `json:"dataURL"`
}

/*
SnapshotCommandLogParams represents LayerTree.snapshotCommandLog parameters.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
type SnapshotCommandLogParams struct {
	// The ID of the layer snapshot.
	SnapshotID SnapshotID `json:"snapshotId"`
}

/*
SnapshotCommandLogResult represents the result of calls to LayerTree.snapshotCommandLog.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#method-snapshotCommandLog
*/
type SnapshotCommandLogResult struct {
	// The array of canvas function calls.
	CommandLog []map[string]string `json:"commandLog"`
}

/*
LayerPaintedEvent represents LayerTree.layerPainted event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerPainted
*/
type LayerPaintedEvent struct {
	// The ID of the painted layer.
	LayerID LayerID `json:"layerId"`

	// Clip rectangle.
	Clip DOM.Rect `json:"clip"`
}

/*
DidChangeEvent represents LayerTree.layerTreeDidChange event data.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#event-layerTreeDidChange
*/
type DidChangeEvent struct {
	// Optional. Layer tree, absent if not in the comspositing mode.
	Layers []Layer `json:"layers,omitempty"`
}

/*
LayerID is a unique Layer identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-LayerId
*/
type LayerID string

/*
SnapshotID is a unique snapshot identifier.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-SnapshotId
*/
type SnapshotID string

/*
ScrollRect is a rectangle where scrolling happens on the main thread.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-ScrollRect
*/
type ScrollRect struct {
	// Rectangle itself.
	Rect DOM.Rect `json:"rect"`

	// Reason for rectangle to force scrolling on the main thread Allowed values: RepaintsOnScroll, \
	// TouchEventHandler, WheelEventHandler.
	Type string `json:"type"`
}

/*
StickyPositionConstraint is sticky position constraints.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-StickyPositionConstraint
*/
type StickyPositionConstraint struct {
	// Layout rectangle of the sticky element before being shifted.
	StickyBoxRect DOM.Rect `json:"stickyBoxRect"`

	// Layout rectangle of the containing block of the sticky element.
	ContainingBlockRect DOM.Rect `json:"containingBlockRect"`

	// Optional. The nearest sticky layer that shifts the sticky box.
	NearestLayerShiftingStickyBox LayerID `json:"nearestLayerShiftingStickyBox,omitempty"`

	// Optional. The nearest sticky layer that shifts the containing block.
	NearestLayerShiftingContainingBlock LayerID `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

/*
PictureTile is a serialized fragment of layer picture along with its offset within the layer.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-PictureTile
*/
type PictureTile struct {
	// Offset from owning layer left boundary.
	X float64 `json:"x"`

	// Offset from owning layer top boundary.
	Y float64 `json:"y"`

	// Base64-encoded snapshot data.
	Picture string `json:"picture"`
}

/*
Layer is information about a compositing layer.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-Layer
*/
type Layer struct {
	// The unique ID for this layer.
	LayerID LayerID `json:"layerId"`

	// Optional. The ID of parent (not present for root).
	ParentLayerID LayerID `json:"parentLayerId,omitempty"`

	// Optional. The backend ID for the node associated with this layer.
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`

	// Offset from parent layer, X coordinate.
	OffsetX float64 `json:"offsetX"`

	// Offset from parent layer, Y coordinate.
	OffsetY float64 `json:"offsetY"`

	// Layer width.
	Width float64 `json:"width"`

	// Layer height.
	Height float64 `json:"height"`

	// Optional. Transformation matrix for layer, default is identity matrix.
	Transform []float64 `json:"transform,omitempty"`

	// Optional. Transform anchor point X, absent if no transform specified.
	AnchorX float64 `json:"anchorX,omitempty"`

	// Optional. Transform anchor point Y, absent if no transform specified.
	AnchorY float64 `json:"anchorY,omitempty"`

	// Optional. Transform anchor point Z, absent if no transform specified.
	AnchorZ float64 `json:"anchorZ,omitempty"`

	// Indicates how many time this layer has painted.
	PaintCount int `json:"paintCount"`

	// Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	DrawsContent bool `json:"drawsContent"`

	// Optional. Set if layer is not visible.
	Invisible bool `json:"invisible,omitempty"`

	// Optional. Rectangles scrolling on main thread only.
	ScrollRects []*ScrollRect `json:"scrollRects,omitempty"`

	// Optional. Sticky position constraint information.
	//
	// This is an instance of StickyPositionConstraint, but that doesn't omitempty correctly so it
	// must be added manually.
	StickyPositionConstraint interface{} `json:"stickyPositionConstraint,omitempty"`
}

/*
PaintProfile is an array of timings, one per paint step.

https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/#type-PaintProfile
*/
type PaintProfile []interface{}
