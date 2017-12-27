/*
Package runtime provides type definitions for use with the Chrome Runtime protocol

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/
*/
package runtime

/*
AwaitPromiseParams represents Runtime.awaitPromise parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
type AwaitPromiseParams struct {
	// Identifier of the promise.
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`

	// Optional. Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
AwaitPromiseResult represents the result of calls to Runtime.awaitPromise.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-awaitPromise
*/
type AwaitPromiseResult struct {
	// Promise result. Will contain rejected value if promise was rejected.
	Result RemoteObject `json:"result"`

	// Exception details if stack strace is available.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
CallFunctionOnParams represents Runtime.callFunctionOn parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
type CallFunctionOnParams struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`

	// Optional. Identifier of the object to call function on. Either objectID or executionContextID
	// should be specified.
	ObjectID RemoteObjectID `json:"objectId,omitempty"`

	// Optional. Call arguments. All call arguments must belong to the same JavaScript world as the
	// target object.
	Arguments []CallArgument `json:"arguments,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not
	// pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise
	// is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`

	// Optional. Specifies execution context which global object will be used to call function on.
	// Either executionContextID or objectID should be specified.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple objects. If objectGroup is
	// not specified and objectID is, objectGroup will be inherited from object.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

/*
CallFunctionOnResult represents the result of calls to Runtime.callFunctionOn.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-callFunctionOn
*/
type CallFunctionOnResult struct {
	// Call result.
	Result RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
CompileScriptParams represents Runtime.compileScript parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
type CompileScriptParams struct {
	// Expression to compile.
	Expression string `json:"expression"`

	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`

	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`

	// Optional. Specifies in which execution context to perform script run. If the parameter is
	// omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
CompileScriptResult represents the result of calls to Runtime.compileScript.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-compileScript
*/
type CompileScriptResult struct {
	// ID of the script.
	ScriptID ScriptID `json:"scriptId"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
EvaluateParams represents Runtime.evaluate parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
type EvaluateParams struct {
	// Expression to evaluate.
	Expression string `json:"expression"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Specifies in which execution context to perform evaluation. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ContextID ExecutionContextID `json:"contextId,omitempty"`

	// Optional. Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
EvaluateResult represents the result of calls to Runtime.evaluate.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-evaluate
*/
type EvaluateResult struct {
	// Evaluation result.
	Result RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
GetPropertiesParams represents Runtime.getProperties parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
type GetPropertiesParams struct {
	// Identifier of the object to return properties for.
	ObjectID RemoteObjectID `json:"objectId"`

	// Optional. If true, returns properties belonging only to the element itself, not to its
	// prototype chain.
	OwnProperties bool `json:"ownProperties,omitempty"`

	// Optional. If true, returns accessor properties (with getter/setter) only; internal properties
	// are not returned either. EXPERIMENTAL
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`

	// Optional. Whether preview should be generated for the results. EXPERIMENTAL
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

/*
GetPropertiesResult represents the result of calls to Runtime.getProperties.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-getProperties
*/
type GetPropertiesResult struct {
	// Object properties.
	Result []PropertyDescriptor `json:"result"`

	// Internal object properties (only of the element itself).
	InternalProperties []InternalPropertyDescriptor `json:"internalProperties"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
GlobalLexicalScopeNamesParams represents Runtime.globalLexicalScopeNames parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
type GlobalLexicalScopeNamesParams struct {
	// Optional. Specifies in which execution context to lookup global scope variables.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
GlobalLexicalScopeNamesResult represents the result of calls to Runtime.globalLexicalScopeNames.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-globalLexicalScopeNames
*/
type GlobalLexicalScopeNamesResult struct {
	Names []string `json:"names"`
}

/*
QueryObjectsParams represents Runtime.queryObjects parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
type QueryObjectsParams struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectID RemoteObjectID `json:"prototypeObjectId"`
}

/*
QueryObjectsResult represents the result of calls to Runtime.queryObjects.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-queryObjects
*/
type QueryObjectsResult struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`
}

/*
ReleaseObjectParams represents Runtime.releaseObject parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObject
*/
type ReleaseObjectParams struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`
}

/*
ReleaseObjectGroupParams represents Runtime.releaseObjectGroup parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-releaseObjectGroup
*/
type ReleaseObjectGroupParams struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

/*
RunScriptParams represents Runtime.runScript parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
type RunScriptParams struct {
	// ID of the script to run.
	ScriptID ScriptID `json:"scriptId"`

	// Optional. Specifies in which execution context to perform script run. If the parameter is
	// omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`

	// Optional. Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`

	// Optional. In silent mode exceptions thrown during evaluation are not reported and do not
	// pause execution. Overrides setPauseOnException state.
	Silent bool `json:"silent,omitempty"`

	// Optional. Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`

	// Optional. Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`

	// Optional. Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// Optional. Whether execution should await for resulting value and return once awaited promise
	// is resolved.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

/*
RunScriptResult represents the result of calls to Runtime.runScript.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-runScript
*/
type RunScriptResult struct {
	// Identifier of the object to release.
	ObjectID RemoteObjectID `json:"objectId"`
}

/*
SetCustomObjectFormatterEnabledParams represents Runtime.setCustomObjectFormatterEnabled parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#method-setCustomObjectFormatterEnabled
*/
type SetCustomObjectFormatterEnabledParams struct {
	// Run result.
	Result RemoteObject `json:"result"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
ConsoleAPICalledEvent represents Runtime.consoleAPICalled event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-consoleAPICalled
*/
type ConsoleAPICalledEvent struct {
	// Type of the call. Allowed values: log, debug, info, error, warning, dir, dirxml, table,
	// trace, clear, startGroup, startGroupCollapsed, endGroup, assert, profile, profileEnd, count,
	// timeEnd.
	Type string `json:"type"`

	// Call arguments.
	Args []RemoteObject `json:"args"`

	// Identifier of the context where the call was made.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`

	// Call timestamp.
	Timestamp Timestamp `json:"timestamp"`

	// Optional. Stack trace captured when the call was made.
	StackTrace StackTrace `json:"stackTrace,omitempty"`

	// Optional. Console context descriptor for calls on non-default console context (not
	// console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id'
	// for call on named context. EXPERIMENTAL
	Context string `json:"context,omitempty"`
}

/*
ExceptionRevokedEvent represents Runtime.exceptionRevoked event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionRevoked
*/
type ExceptionRevokedEvent struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`

	// The ID of revoked exception, as reported in exceptionThrown.
	ExceptionID int `json:"exceptionId"`
}

/*
ExceptionThrownEvent represents Runtime.exceptionThrown event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-exceptionThrown
*/
type ExceptionThrownEvent struct {
	// Timestamp of the exception.
	Timestamp Timestamp `json:"timestamp"`

	// Exception details.
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

/*
ExecutionContextCreatedEvent represents Runtime.executionContextCreated event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextCreated
*/
type ExecutionContextCreatedEvent struct {
	// A newly created execution context.
	Context ExecutionContextDescription `json:"context"`
}

/*
ExecutionContextDestroyedEvent represents Runtime.executionContextDestroyed event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextDestroyed
*/
type ExecutionContextDestroyedEvent struct {
	// ID of the destroyed context.
	ExecutionContextID ExecutionContextID `json:"executionContextId"`
}

/*
ExecutionContextsClearedEvent represents Runtime.executionContextsCleared event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-executionContextsCleared
*/
type ExecutionContextsClearedEvent struct{}

/*
InspectRequestedEvent represents Runtime.inspectRequested event data.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#event-inspectRequested
*/
type InspectRequestedEvent struct {
	// Remote object.
	Object RemoteObject `json:"object"`

	// Hints.
	Hints map[string]string `json:"hints"`
}

/*
ScriptID is a unique script identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ScriptId
*/
type ScriptID string

/*
RemoteObjectID is a unique object identifier.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObjectId
*/
type RemoteObjectID string

/*
UnserializableValue is a primitive value which cannot be JSON-stringified.

ALLOWED VALUES
	- Infinity
	- NaN
	- -Infinity
	- -0

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-UnserializableValue
*/
type UnserializableValue int

/*
RemoteObject is a mirror object referencing the original JavaScript object.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-RemoteObject
*/
type RemoteObject struct {
	// Object type. Allowed values: object, function, undefined, string, number, boolean, symbol.
	Type string `json:"type"`

	// Optional. Object subtype hint. Specified for object type values only. Allowed values: array,
	// null, node, regexp, date, map, set, weakmap, weakset, iterator, generator, error, proxy,
	// promise, typedarray.
	Subtype string `json:"subtype,omitempty"`

	// Optional. Object class (constructor) name. Specified for object type values only.
	ClassName string `json:"className,omitempty"`

	// Optional. Remote object value in case of primitive values or JSON values (if it was
	// requested).
	Value interface{} `json:"value,omitempty"`

	// Optional. Primitive value which can not be JSON-stringified does not have value, but gets
	// this property.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`

	// Optional. String representation of the object.
	Description string `json:"description,omitempty"`

	// Optional. Unique object identifier (for non-primitive values).
	ObjectID RemoteObjectID `json:"objectId,omitempty"`

	// Optional. Preview containing abbreviated property values. Specified for object type values
	// only. EXPERIMENTAL
	//
	// This expects an instance of ObjectPreview, but that doesn't omitempty correctly so it must be
	// added manually.
	Preview interface{} `json:"preview,omitempty"`

	// Optional. EXPERIMENTAL
	//
	// This expects an instance of CustomPreview, but that doesn't omitempty correctly so it must be
	// added manually.
	CustomPreview interface{} `json:"customPreview,omitempty"`
}

/*
CustomPreview is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CustomPreview
*/
type CustomPreview struct {
	Header                     string         `json:"header"`
	HasBody                    bool           `json:"hasBody"`
	FormatterObjectID          RemoteObjectID `json:"formatterObjectId"`
	BindRemoteObjectFunctionID RemoteObjectID `json:"bindRemoteObjectFunctionId"`
	// Optional.
	ConfigObjectID RemoteObjectID `json:"configObjectId,omitempty,omitempty"`
}

/*
ObjectPreview is an object containing abbreviated remote object value. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ObjectPreview
*/
type ObjectPreview struct {
	// Object type. Allowed values: object, function, undefined, string, number, boolean, symbol.
	Type string `json:"type"`

	// Object subtype hint. Specified for object type values only. Allowed values: array, null,
	// node, regexp, date, map, set, weakmap, weakset, iterator, generator, error.
	Subtype string `json:"subtype"`

	// Optional. String representation of the object.
	Description string `json:"description,omitempty"`

	// True iff some of the properties or entries of the original object did not fit.
	Overflow bool `json:"overflow"`

	// List of the properties.
	Properties []*PropertyPreview `json:"properties"`

	// Optional. List of the entries. Specified for map and set subtype values only.
	Entries []*EntryPreview `json:"entries,omitempty"`
}

/*
PropertyPreview is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyPreview
*/
type PropertyPreview struct {
	// Property name.
	Name string `json:"name"`

	// Object type. Accessor means that the property itself is an accessor property. Allowed values:
	// object, function, undefined, string, number, boolean, symbol, accessor.
	Type string `json:"type"`

	// Optional. User-friendly property value string.
	Value string `json:"value,omitempty"`

	// Optional. Nested value preview.
	//
	// This expects an instance of ObjectPreview, but that doesn't omitempty correctly so it must be
	// added manually.
	ValuePreview interface{} `json:"valuePreview,omitempty"`

	// Optional. Object subtype hint. Specified for object type values only. Allowed values: array,
	// null, node, regexp, date, map, set, weakmap, weakset, iterator, generator, error.
	Subtype string `json:"subtype,omitempty"`
}

/*
EntryPreview is EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-EntryPreview
*/
type EntryPreview struct {
	// Optional. Preview of the key. Specified for map-like collection entries.
	//
	// This expects an instance of ObjectPreview, but that doesn't omitempty correctly so it must be
	// added manually.
	Key interface{} `json:"key,omitempty"`

	// Preview of the value.
	Value ObjectPreview `json:"value"`
}

/*
PropertyDescriptor is an object property descriptor.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-PropertyDescriptor
*/
type PropertyDescriptor struct {
	// Property name or symbol description.
	Name string `json:"name"`

	// Optional. The value associated with the property.
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Value interface{} `json:"value,omitempty"`

	// Optional. True if the value associated with the property may be changed (data descriptors
	// only).
	Writable bool `json:"writable,omitempty"`

	// Optional. A function which serves as a getter for the property, or undefined if there is no
	// getter (accessor descriptors only).
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Get interface{} `json:"get,omitempty"`

	// Optional. A function which serves as a setter for the property, or undefined if there is no
	// setter (accessor descriptors only).
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Set interface{} `json:"set,omitempty"`

	// True if the type of this property descriptor may be changed and if the property may be
	// deleted from the corresponding object.
	Configurable bool `json:"configurable"`

	// True if this property shows up during enumeration of the properties on the corresponding
	// object.
	Enumerable bool `json:"enumerable"`

	// Optional. True if the result was thrown during the evaluation.
	WasThrown bool `json:"wasThrown,omitempty"`

	// Optional. True if the property is owned for the object.
	IsOwn bool `json:"isOwn,omitempty"`

	// Optional. Property symbol object, if the property is of the symbol type.
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Symbol interface{} `json:"symbol,omitempty"`
}

/*
InternalPropertyDescriptor is an object's internal property descriptor. This property isn't normally
visible in JavaScript code.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-InternalPropertyDescriptor
*/
type InternalPropertyDescriptor struct {
	// Conventional property name.
	Name string `json:"name"`

	// Optional. The value associated with the property.
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Value interface{} `json:"value,omitempty"`
}

/*
CallArgument represents a function call argument. Either remote object id objectId, primitive value,
unserializable primitive value or neither of (for undefined) them should be specified.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CallArgument
*/
type CallArgument struct {
	// Optional. Primitive value or serializable javascript object.
	Value interface{} `json:"value,omitempty"`

	// Optional. Primitive value which can not be JSON-stringified.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"`

	// Optional. Remote object handle.
	ObjectID RemoteObjectID `json:"objectId,omitempty"`
}

/*
ExecutionContextID is the ID of an execution context.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExecutionContextId
*/
type ExecutionContextID int

/*
ExecutionContextDescription is the description of an isolated world.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExecutionContextDescription
*/
type ExecutionContextDescription struct {
	// Unique ID of the execution context. It can be used to specify in which execution context
	// script evaluation should be performed.
	ID ExecutionContextID `json:"id"`

	// Execution context origin.
	Origin string `json:"origin"`

	// Human readable name describing given context.
	Name string `json:"name"`

	// Optional. Embedder-specific auxiliary data.
	AuxData map[string]string `json:"auxData,omitempty"`
}

/*
ExceptionDetails contains detailed information about exception (or error) that was thrown during
script compilation or execution.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-ExceptionDetails
*/
type ExceptionDetails struct {
	// Exception id.
	ExceptionID int `json:"exceptionId"`

	// Exception text, which should be used together with exception object when available.
	Text string `json:"text"`

	// Line number of the exception location (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number of the exception location (0-based).
	ColumnNumber int `json:"columnNumber"`

	// Optional. Script ID of the exception location.
	ScriptID ScriptID `json:"scriptId,omitempty"`

	// Optional. URL of the exception location, to be used when the script was not reported.
	URL string `json:"url,omitempty"`

	// Optional. JavaScript stack trace if available.
	//
	// This expects an instance of StackTrace, but that doesn't omitempty correctly so it must be
	// added manually.
	StackTrace interface{} `json:"stackTrace,omitempty"`

	// Optional. Exception object if available.
	//
	// This expects an instance of RemoteObject, but that doesn't omitempty correctly so it must be
	// added manually.
	Exception interface{} `json:"exception,omitempty"`

	// Optional. Identifier of the context where exception happened.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"`
}

/*
Timestamp is the number of milliseconds since epoch.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-Timestamp
*/
type Timestamp int

/*
CallFrame is a stack entry for runtime errors and assertions.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-CallFrame
*/
type CallFrame struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`

	// JavaScript script id.
	ScriptID ScriptID `json:"scriptId"`

	// JavaScript script name or url.
	URL string `json:"url"`

	// JavaScript script line number (0-based).
	LineNumber int `json:"lineNumber"`

	// JavaScript script column number (0-based).
	ColumnNumber int `json:"columnNumber"`
}

/*
StackTrace contains call frames for assertions or error messages.

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-StackTrace
*/
type StackTrace struct {
	// Optional. String label of this stack trace. For async traces this may be a name of the
	// function that initiated the async call.
	Description string `json:"description,omitempty"`

	// JavaScript function name.
	CallFrames []*CallFrame `json:"callFrames"`

	// Optional. Asynchronous JavaScript stack trace that preceded this stack, if available.
	//
	// This expects an instance of StackTrace, but that doesn't omitempty correctly so it must be
	// added manually.
	Parent interface{} `json:"parent,omitempty"`

	// Optional. Asynchronous JavaScript stack trace that preceded this stack, if available.
	// EXPERIMENTAL
	ParentID StackTraceID `json:"parentId,omitempty"`
}

/*
UniqueDebuggerID is the unique identifier of the  current debugger. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-UniqueDebuggerId
*/
type UniqueDebuggerID string

/*
StackTraceID - If debuggerID is set stack trace comes from another debugger and can be resolved
there. This allows to track cross-debugger calls. See StackTrace and Debugger.paused for usages.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Runtime/#type-StackTraceId
*/
type StackTraceID struct {
	ID string `json:"id"`
	// Optional.
	DebuggerID UniqueDebuggerID `json:"debuggerId,omitempty"`
}