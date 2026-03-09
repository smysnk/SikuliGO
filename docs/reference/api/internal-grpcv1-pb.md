# API: `internal/grpcv1/pb`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package sikuliv1 // import "github.com/smysnk/sikuligo/internal/grpcv1/pb"`

## Symbol Index

### Types

- <span class="api-type">[`ActionResponse`](#type-actionresponse)</span>
- <span class="api-type">[`AppActionRequest`](#type-appactionrequest)</span>
- <span class="api-type">[`AppOptions`](#type-appoptions)</span>
- <span class="api-type">[`CaptureScreenRequest`](#type-capturescreenrequest)</span>
- <span class="api-type">[`CaptureScreenResponse`](#type-capturescreenresponse)</span>
- <span class="api-type">[`ClickOnScreenRequest`](#type-clickonscreenrequest)</span>
- <span class="api-type">[`ClickRequest`](#type-clickrequest)</span>
- <span class="api-type">[`ExistsOnScreenRequest`](#type-existsonscreenrequest)</span>
- <span class="api-type">[`ExistsOnScreenResponse`](#type-existsonscreenresponse)</span>
- <span class="api-type">[`FindAllResponse`](#type-findallresponse)</span>
- <span class="api-type">[`FindOnScreenRequest`](#type-findonscreenrequest)</span>
- <span class="api-type">[`FindRequest`](#type-findrequest)</span>
- <span class="api-type">[`FindResponse`](#type-findresponse)</span>
- <span class="api-type">[`FindTextRequest`](#type-findtextrequest)</span>
- <span class="api-type">[`FindTextResponse`](#type-findtextresponse)</span>
- <span class="api-type">[`GetPrimaryScreenRequest`](#type-getprimaryscreenrequest)</span>
- <span class="api-type">[`GetPrimaryScreenResponse`](#type-getprimaryscreenresponse)</span>
- <span class="api-type">[`GetWindowResponse`](#type-getwindowresponse)</span>
- <span class="api-type">[`GrayImage`](#type-grayimage)</span>
- <span class="api-type">[`HotkeyRequest`](#type-hotkeyrequest)</span>
- <span class="api-type">[`InputOptions`](#type-inputoptions)</span>
- <span class="api-type">[`IsAppRunningResponse`](#type-isapprunningresponse)</span>
- <span class="api-type">[`ListScreensRequest`](#type-listscreensrequest)</span>
- <span class="api-type">[`ListScreensResponse`](#type-listscreensresponse)</span>
- <span class="api-type">[`ListWindowsResponse`](#type-listwindowsresponse)</span>
- <span class="api-type">[`Match`](#type-match)</span>
- <span class="api-type">[`MatcherEngine`](#type-matcherengine)</span>
- <span class="api-type">[`MoveMouseRequest`](#type-movemouserequest)</span>
- <span class="api-type">[`OCRParams`](#type-ocrparams)</span>
- <span class="api-type">[`ObserveChangeRequest`](#type-observechangerequest)</span>
- <span class="api-type">[`ObserveEvent`](#type-observeevent)</span>
- <span class="api-type">[`ObserveOptions`](#type-observeoptions)</span>
- <span class="api-type">[`ObserveRequest`](#type-observerequest)</span>
- <span class="api-type">[`ObserveResponse`](#type-observeresponse)</span>
- <span class="api-type">[`Pattern`](#type-pattern)</span>
- <span class="api-type">[`Point`](#type-point)</span>
- <span class="api-type">[`ReadTextRequest`](#type-readtextrequest)</span>
- <span class="api-type">[`ReadTextResponse`](#type-readtextresponse)</span>
- <span class="api-type">[`Rect`](#type-rect)</span>
- <span class="api-type">[`ScreenDescriptor`](#type-screendescriptor)</span>
- <span class="api-type">[`ScreenQueryOptions`](#type-screenqueryoptions)</span>
- <span class="api-type">[`ScrollWheelRequest`](#type-scrollwheelrequest)</span>
- <span class="api-type">[`SikuliServiceClient`](#type-sikuliserviceclient)</span>
- <span class="api-type">[`SikuliServiceServer`](#type-sikuliserviceserver)</span>
- <span class="api-type">[`TextMatch`](#type-textmatch)</span>
- <span class="api-type">[`TypeTextRequest`](#type-typetextrequest)</span>
- <span class="api-type">[`UnimplementedSikuliServiceServer`](#type-unimplementedsikuliserviceserver)</span>
- <span class="api-type">[`UnsafeSikuliServiceServer`](#type-unsafesikuliserviceserver)</span>
- <span class="api-type">[`WaitOnScreenRequest`](#type-waitonscreenrequest)</span>
- <span class="api-type">[`Window`](#type-window)</span>
- <span class="api-type">[`WindowQuery`](#type-windowquery)</span>
- <span class="api-type">[`WindowQueryRequest`](#type-windowqueryrequest)</span>

### Functions

- <span class="api-func">[`RegisterSikuliServiceServer`](#func-registersikuliserviceserver)</span>
- <span class="api-func">[`NewSikuliServiceClient`](#func-newsikuliserviceclient)</span>

### Methods

- <span class="api-method">[`ActionResponse.Descriptor`](#method-actionresponse-descriptor)</span>
- <span class="api-method">[`ActionResponse.ProtoMessage`](#method-actionresponse-protomessage)</span>
- <span class="api-method">[`ActionResponse.ProtoReflect`](#method-actionresponse-protoreflect)</span>
- <span class="api-method">[`ActionResponse.Reset`](#method-actionresponse-reset)</span>
- <span class="api-method">[`ActionResponse.String`](#method-actionresponse-string)</span>
- <span class="api-method">[`AppActionRequest.Descriptor`](#method-appactionrequest-descriptor)</span>
- <span class="api-method">[`AppActionRequest.GetArgs`](#method-appactionrequest-getargs)</span>
- <span class="api-method">[`AppActionRequest.GetName`](#method-appactionrequest-getname)</span>
- <span class="api-method">[`AppActionRequest.GetOpts`](#method-appactionrequest-getopts)</span>
- <span class="api-method">[`AppActionRequest.ProtoMessage`](#method-appactionrequest-protomessage)</span>
- <span class="api-method">[`AppActionRequest.ProtoReflect`](#method-appactionrequest-protoreflect)</span>
- <span class="api-method">[`AppActionRequest.Reset`](#method-appactionrequest-reset)</span>
- <span class="api-method">[`AppActionRequest.String`](#method-appactionrequest-string)</span>
- <span class="api-method">[`AppOptions.Descriptor`](#method-appoptions-descriptor)</span>
- <span class="api-method">[`AppOptions.GetTimeoutMillis`](#method-appoptions-gettimeoutmillis)</span>
- <span class="api-method">[`AppOptions.ProtoMessage`](#method-appoptions-protomessage)</span>
- <span class="api-method">[`AppOptions.ProtoReflect`](#method-appoptions-protoreflect)</span>
- <span class="api-method">[`AppOptions.Reset`](#method-appoptions-reset)</span>
- <span class="api-method">[`AppOptions.String`](#method-appoptions-string)</span>
- <span class="api-method">[`CaptureScreenRequest.Descriptor`](#method-capturescreenrequest-descriptor)</span>
- <span class="api-method">[`CaptureScreenRequest.GetRegion`](#method-capturescreenrequest-getregion)</span>
- <span class="api-method">[`CaptureScreenRequest.GetScreenId`](#method-capturescreenrequest-getscreenid)</span>
- <span class="api-method">[`CaptureScreenRequest.ProtoMessage`](#method-capturescreenrequest-protomessage)</span>
- <span class="api-method">[`CaptureScreenRequest.ProtoReflect`](#method-capturescreenrequest-protoreflect)</span>
- <span class="api-method">[`CaptureScreenRequest.Reset`](#method-capturescreenrequest-reset)</span>
- <span class="api-method">[`CaptureScreenRequest.String`](#method-capturescreenrequest-string)</span>
- <span class="api-method">[`CaptureScreenResponse.Descriptor`](#method-capturescreenresponse-descriptor)</span>
- <span class="api-method">[`CaptureScreenResponse.GetImage`](#method-capturescreenresponse-getimage)</span>
- <span class="api-method">[`CaptureScreenResponse.GetScreen`](#method-capturescreenresponse-getscreen)</span>
- <span class="api-method">[`CaptureScreenResponse.ProtoMessage`](#method-capturescreenresponse-protomessage)</span>
- <span class="api-method">[`CaptureScreenResponse.ProtoReflect`](#method-capturescreenresponse-protoreflect)</span>
- <span class="api-method">[`CaptureScreenResponse.Reset`](#method-capturescreenresponse-reset)</span>
- <span class="api-method">[`CaptureScreenResponse.String`](#method-capturescreenresponse-string)</span>
- <span class="api-method">[`ClickOnScreenRequest.Descriptor`](#method-clickonscreenrequest-descriptor)</span>
- <span class="api-method">[`ClickOnScreenRequest.GetClickOpts`](#method-clickonscreenrequest-getclickopts)</span>
- <span class="api-method">[`ClickOnScreenRequest.GetOpts`](#method-clickonscreenrequest-getopts)</span>
- <span class="api-method">[`ClickOnScreenRequest.GetPattern`](#method-clickonscreenrequest-getpattern)</span>
- <span class="api-method">[`ClickOnScreenRequest.ProtoMessage`](#method-clickonscreenrequest-protomessage)</span>
- <span class="api-method">[`ClickOnScreenRequest.ProtoReflect`](#method-clickonscreenrequest-protoreflect)</span>
- <span class="api-method">[`ClickOnScreenRequest.Reset`](#method-clickonscreenrequest-reset)</span>
- <span class="api-method">[`ClickOnScreenRequest.String`](#method-clickonscreenrequest-string)</span>
- <span class="api-method">[`ClickRequest.Descriptor`](#method-clickrequest-descriptor)</span>
- <span class="api-method">[`ClickRequest.GetOpts`](#method-clickrequest-getopts)</span>
- <span class="api-method">[`ClickRequest.GetX`](#method-clickrequest-getx)</span>
- <span class="api-method">[`ClickRequest.GetY`](#method-clickrequest-gety)</span>
- <span class="api-method">[`ClickRequest.ProtoMessage`](#method-clickrequest-protomessage)</span>
- <span class="api-method">[`ClickRequest.ProtoReflect`](#method-clickrequest-protoreflect)</span>
- <span class="api-method">[`ClickRequest.Reset`](#method-clickrequest-reset)</span>
- <span class="api-method">[`ClickRequest.String`](#method-clickrequest-string)</span>
- <span class="api-method">[`ExistsOnScreenRequest.Descriptor`](#method-existsonscreenrequest-descriptor)</span>
- <span class="api-method">[`ExistsOnScreenRequest.GetOpts`](#method-existsonscreenrequest-getopts)</span>
- <span class="api-method">[`ExistsOnScreenRequest.GetPattern`](#method-existsonscreenrequest-getpattern)</span>
- <span class="api-method">[`ExistsOnScreenRequest.ProtoMessage`](#method-existsonscreenrequest-protomessage)</span>
- <span class="api-method">[`ExistsOnScreenRequest.ProtoReflect`](#method-existsonscreenrequest-protoreflect)</span>
- <span class="api-method">[`ExistsOnScreenRequest.Reset`](#method-existsonscreenrequest-reset)</span>
- <span class="api-method">[`ExistsOnScreenRequest.String`](#method-existsonscreenrequest-string)</span>
- <span class="api-method">[`ExistsOnScreenResponse.Descriptor`](#method-existsonscreenresponse-descriptor)</span>
- <span class="api-method">[`ExistsOnScreenResponse.GetExists`](#method-existsonscreenresponse-getexists)</span>
- <span class="api-method">[`ExistsOnScreenResponse.GetMatch`](#method-existsonscreenresponse-getmatch)</span>
- <span class="api-method">[`ExistsOnScreenResponse.ProtoMessage`](#method-existsonscreenresponse-protomessage)</span>
- <span class="api-method">[`ExistsOnScreenResponse.ProtoReflect`](#method-existsonscreenresponse-protoreflect)</span>
- <span class="api-method">[`ExistsOnScreenResponse.Reset`](#method-existsonscreenresponse-reset)</span>
- <span class="api-method">[`ExistsOnScreenResponse.String`](#method-existsonscreenresponse-string)</span>
- <span class="api-method">[`FindAllResponse.Descriptor`](#method-findallresponse-descriptor)</span>
- <span class="api-method">[`FindAllResponse.GetMatches`](#method-findallresponse-getmatches)</span>
- <span class="api-method">[`FindAllResponse.ProtoMessage`](#method-findallresponse-protomessage)</span>
- <span class="api-method">[`FindAllResponse.ProtoReflect`](#method-findallresponse-protoreflect)</span>
- <span class="api-method">[`FindAllResponse.Reset`](#method-findallresponse-reset)</span>
- <span class="api-method">[`FindAllResponse.String`](#method-findallresponse-string)</span>
- <span class="api-method">[`FindOnScreenRequest.Descriptor`](#method-findonscreenrequest-descriptor)</span>
- <span class="api-method">[`FindOnScreenRequest.GetOpts`](#method-findonscreenrequest-getopts)</span>
- <span class="api-method">[`FindOnScreenRequest.GetPattern`](#method-findonscreenrequest-getpattern)</span>
- <span class="api-method">[`FindOnScreenRequest.ProtoMessage`](#method-findonscreenrequest-protomessage)</span>
- <span class="api-method">[`FindOnScreenRequest.ProtoReflect`](#method-findonscreenrequest-protoreflect)</span>
- <span class="api-method">[`FindOnScreenRequest.Reset`](#method-findonscreenrequest-reset)</span>
- <span class="api-method">[`FindOnScreenRequest.String`](#method-findonscreenrequest-string)</span>
- <span class="api-method">[`FindRequest.Descriptor`](#method-findrequest-descriptor)</span>
- <span class="api-method">[`FindRequest.GetMatcherEngine`](#method-findrequest-getmatcherengine)</span>
- <span class="api-method">[`FindRequest.GetPattern`](#method-findrequest-getpattern)</span>
- <span class="api-method">[`FindRequest.GetSource`](#method-findrequest-getsource)</span>
- <span class="api-method">[`FindRequest.ProtoMessage`](#method-findrequest-protomessage)</span>
- <span class="api-method">[`FindRequest.ProtoReflect`](#method-findrequest-protoreflect)</span>
- <span class="api-method">[`FindRequest.Reset`](#method-findrequest-reset)</span>
- <span class="api-method">[`FindRequest.String`](#method-findrequest-string)</span>
- <span class="api-method">[`FindResponse.Descriptor`](#method-findresponse-descriptor)</span>
- <span class="api-method">[`FindResponse.GetMatch`](#method-findresponse-getmatch)</span>
- <span class="api-method">[`FindResponse.ProtoMessage`](#method-findresponse-protomessage)</span>
- <span class="api-method">[`FindResponse.ProtoReflect`](#method-findresponse-protoreflect)</span>
- <span class="api-method">[`FindResponse.Reset`](#method-findresponse-reset)</span>
- <span class="api-method">[`FindResponse.String`](#method-findresponse-string)</span>
- <span class="api-method">[`FindTextRequest.Descriptor`](#method-findtextrequest-descriptor)</span>
- <span class="api-method">[`FindTextRequest.GetParams`](#method-findtextrequest-getparams)</span>
- <span class="api-method">[`FindTextRequest.GetQuery`](#method-findtextrequest-getquery)</span>
- <span class="api-method">[`FindTextRequest.GetSource`](#method-findtextrequest-getsource)</span>
- <span class="api-method">[`FindTextRequest.ProtoMessage`](#method-findtextrequest-protomessage)</span>
- <span class="api-method">[`FindTextRequest.ProtoReflect`](#method-findtextrequest-protoreflect)</span>
- <span class="api-method">[`FindTextRequest.Reset`](#method-findtextrequest-reset)</span>
- <span class="api-method">[`FindTextRequest.String`](#method-findtextrequest-string)</span>
- <span class="api-method">[`FindTextResponse.Descriptor`](#method-findtextresponse-descriptor)</span>
- <span class="api-method">[`FindTextResponse.GetMatches`](#method-findtextresponse-getmatches)</span>
- <span class="api-method">[`FindTextResponse.ProtoMessage`](#method-findtextresponse-protomessage)</span>
- <span class="api-method">[`FindTextResponse.ProtoReflect`](#method-findtextresponse-protoreflect)</span>
- <span class="api-method">[`FindTextResponse.Reset`](#method-findtextresponse-reset)</span>
- <span class="api-method">[`FindTextResponse.String`](#method-findtextresponse-string)</span>
- <span class="api-method">[`GetPrimaryScreenRequest.Descriptor`](#method-getprimaryscreenrequest-descriptor)</span>
- <span class="api-method">[`GetPrimaryScreenRequest.ProtoMessage`](#method-getprimaryscreenrequest-protomessage)</span>
- <span class="api-method">[`GetPrimaryScreenRequest.ProtoReflect`](#method-getprimaryscreenrequest-protoreflect)</span>
- <span class="api-method">[`GetPrimaryScreenRequest.Reset`](#method-getprimaryscreenrequest-reset)</span>
- <span class="api-method">[`GetPrimaryScreenRequest.String`](#method-getprimaryscreenrequest-string)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.Descriptor`](#method-getprimaryscreenresponse-descriptor)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.GetScreen`](#method-getprimaryscreenresponse-getscreen)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.ProtoMessage`](#method-getprimaryscreenresponse-protomessage)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.ProtoReflect`](#method-getprimaryscreenresponse-protoreflect)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.Reset`](#method-getprimaryscreenresponse-reset)</span>
- <span class="api-method">[`GetPrimaryScreenResponse.String`](#method-getprimaryscreenresponse-string)</span>
- <span class="api-method">[`GetWindowResponse.Descriptor`](#method-getwindowresponse-descriptor)</span>
- <span class="api-method">[`GetWindowResponse.GetFound`](#method-getwindowresponse-getfound)</span>
- <span class="api-method">[`GetWindowResponse.GetWindow`](#method-getwindowresponse-getwindow)</span>
- <span class="api-method">[`GetWindowResponse.ProtoMessage`](#method-getwindowresponse-protomessage)</span>
- <span class="api-method">[`GetWindowResponse.ProtoReflect`](#method-getwindowresponse-protoreflect)</span>
- <span class="api-method">[`GetWindowResponse.Reset`](#method-getwindowresponse-reset)</span>
- <span class="api-method">[`GetWindowResponse.String`](#method-getwindowresponse-string)</span>
- <span class="api-method">[`GrayImage.Descriptor`](#method-grayimage-descriptor)</span>
- <span class="api-method">[`GrayImage.GetHeight`](#method-grayimage-getheight)</span>
- <span class="api-method">[`GrayImage.GetName`](#method-grayimage-getname)</span>
- <span class="api-method">[`GrayImage.GetPix`](#method-grayimage-getpix)</span>
- <span class="api-method">[`GrayImage.GetWidth`](#method-grayimage-getwidth)</span>
- <span class="api-method">[`GrayImage.ProtoMessage`](#method-grayimage-protomessage)</span>
- <span class="api-method">[`GrayImage.ProtoReflect`](#method-grayimage-protoreflect)</span>
- <span class="api-method">[`GrayImage.Reset`](#method-grayimage-reset)</span>
- <span class="api-method">[`GrayImage.String`](#method-grayimage-string)</span>
- <span class="api-method">[`HotkeyRequest.Descriptor`](#method-hotkeyrequest-descriptor)</span>
- <span class="api-method">[`HotkeyRequest.GetKeys`](#method-hotkeyrequest-getkeys)</span>
- <span class="api-method">[`HotkeyRequest.ProtoMessage`](#method-hotkeyrequest-protomessage)</span>
- <span class="api-method">[`HotkeyRequest.ProtoReflect`](#method-hotkeyrequest-protoreflect)</span>
- <span class="api-method">[`HotkeyRequest.Reset`](#method-hotkeyrequest-reset)</span>
- <span class="api-method">[`HotkeyRequest.String`](#method-hotkeyrequest-string)</span>
- <span class="api-method">[`InputOptions.Descriptor`](#method-inputoptions-descriptor)</span>
- <span class="api-method">[`InputOptions.GetButton`](#method-inputoptions-getbutton)</span>
- <span class="api-method">[`InputOptions.GetDelayMillis`](#method-inputoptions-getdelaymillis)</span>
- <span class="api-method">[`InputOptions.ProtoMessage`](#method-inputoptions-protomessage)</span>
- <span class="api-method">[`InputOptions.ProtoReflect`](#method-inputoptions-protoreflect)</span>
- <span class="api-method">[`InputOptions.Reset`](#method-inputoptions-reset)</span>
- <span class="api-method">[`InputOptions.String`](#method-inputoptions-string)</span>
- <span class="api-method">[`IsAppRunningResponse.Descriptor`](#method-isapprunningresponse-descriptor)</span>
- <span class="api-method">[`IsAppRunningResponse.GetRunning`](#method-isapprunningresponse-getrunning)</span>
- <span class="api-method">[`IsAppRunningResponse.ProtoMessage`](#method-isapprunningresponse-protomessage)</span>
- <span class="api-method">[`IsAppRunningResponse.ProtoReflect`](#method-isapprunningresponse-protoreflect)</span>
- <span class="api-method">[`IsAppRunningResponse.Reset`](#method-isapprunningresponse-reset)</span>
- <span class="api-method">[`IsAppRunningResponse.String`](#method-isapprunningresponse-string)</span>
- <span class="api-method">[`ListScreensRequest.Descriptor`](#method-listscreensrequest-descriptor)</span>
- <span class="api-method">[`ListScreensRequest.ProtoMessage`](#method-listscreensrequest-protomessage)</span>
- <span class="api-method">[`ListScreensRequest.ProtoReflect`](#method-listscreensrequest-protoreflect)</span>
- <span class="api-method">[`ListScreensRequest.Reset`](#method-listscreensrequest-reset)</span>
- <span class="api-method">[`ListScreensRequest.String`](#method-listscreensrequest-string)</span>
- <span class="api-method">[`ListScreensResponse.Descriptor`](#method-listscreensresponse-descriptor)</span>
- <span class="api-method">[`ListScreensResponse.GetScreens`](#method-listscreensresponse-getscreens)</span>
- <span class="api-method">[`ListScreensResponse.ProtoMessage`](#method-listscreensresponse-protomessage)</span>
- <span class="api-method">[`ListScreensResponse.ProtoReflect`](#method-listscreensresponse-protoreflect)</span>
- <span class="api-method">[`ListScreensResponse.Reset`](#method-listscreensresponse-reset)</span>
- <span class="api-method">[`ListScreensResponse.String`](#method-listscreensresponse-string)</span>
- <span class="api-method">[`ListWindowsResponse.Descriptor`](#method-listwindowsresponse-descriptor)</span>
- <span class="api-method">[`ListWindowsResponse.GetWindows`](#method-listwindowsresponse-getwindows)</span>
- <span class="api-method">[`ListWindowsResponse.ProtoMessage`](#method-listwindowsresponse-protomessage)</span>
- <span class="api-method">[`ListWindowsResponse.ProtoReflect`](#method-listwindowsresponse-protoreflect)</span>
- <span class="api-method">[`ListWindowsResponse.Reset`](#method-listwindowsresponse-reset)</span>
- <span class="api-method">[`ListWindowsResponse.String`](#method-listwindowsresponse-string)</span>
- <span class="api-method">[`Match.Descriptor`](#method-match-descriptor)</span>
- <span class="api-method">[`Match.GetIndex`](#method-match-getindex)</span>
- <span class="api-method">[`Match.GetRect`](#method-match-getrect)</span>
- <span class="api-method">[`Match.GetScore`](#method-match-getscore)</span>
- <span class="api-method">[`Match.GetTarget`](#method-match-gettarget)</span>
- <span class="api-method">[`Match.ProtoMessage`](#method-match-protomessage)</span>
- <span class="api-method">[`Match.ProtoReflect`](#method-match-protoreflect)</span>
- <span class="api-method">[`Match.Reset`](#method-match-reset)</span>
- <span class="api-method">[`Match.String`](#method-match-string)</span>
- <span class="api-method">[`MatcherEngine.Descriptor`](#method-matcherengine-descriptor)</span>
- <span class="api-method">[`MatcherEngine.Enum`](#method-matcherengine-enum)</span>
- <span class="api-method">[`MatcherEngine.EnumDescriptor`](#method-matcherengine-enumdescriptor)</span>
- <span class="api-method">[`MatcherEngine.Number`](#method-matcherengine-number)</span>
- <span class="api-method">[`MatcherEngine.String`](#method-matcherengine-string)</span>
- <span class="api-method">[`MatcherEngine.Type`](#method-matcherengine-type)</span>
- <span class="api-method">[`MoveMouseRequest.Descriptor`](#method-movemouserequest-descriptor)</span>
- <span class="api-method">[`MoveMouseRequest.GetOpts`](#method-movemouserequest-getopts)</span>
- <span class="api-method">[`MoveMouseRequest.GetX`](#method-movemouserequest-getx)</span>
- <span class="api-method">[`MoveMouseRequest.GetY`](#method-movemouserequest-gety)</span>
- <span class="api-method">[`MoveMouseRequest.ProtoMessage`](#method-movemouserequest-protomessage)</span>
- <span class="api-method">[`MoveMouseRequest.ProtoReflect`](#method-movemouserequest-protoreflect)</span>
- <span class="api-method">[`MoveMouseRequest.Reset`](#method-movemouserequest-reset)</span>
- <span class="api-method">[`MoveMouseRequest.String`](#method-movemouserequest-string)</span>
- <span class="api-method">[`OCRParams.Descriptor`](#method-ocrparams-descriptor)</span>
- <span class="api-method">[`OCRParams.GetCaseSensitive`](#method-ocrparams-getcasesensitive)</span>
- <span class="api-method">[`OCRParams.GetLanguage`](#method-ocrparams-getlanguage)</span>
- <span class="api-method">[`OCRParams.GetMinConfidence`](#method-ocrparams-getminconfidence)</span>
- <span class="api-method">[`OCRParams.GetTimeoutMillis`](#method-ocrparams-gettimeoutmillis)</span>
- <span class="api-method">[`OCRParams.GetTrainingDataPath`](#method-ocrparams-gettrainingdatapath)</span>
- <span class="api-method">[`OCRParams.ProtoMessage`](#method-ocrparams-protomessage)</span>
- <span class="api-method">[`OCRParams.ProtoReflect`](#method-ocrparams-protoreflect)</span>
- <span class="api-method">[`OCRParams.Reset`](#method-ocrparams-reset)</span>
- <span class="api-method">[`OCRParams.String`](#method-ocrparams-string)</span>
- <span class="api-method">[`ObserveChangeRequest.Descriptor`](#method-observechangerequest-descriptor)</span>
- <span class="api-method">[`ObserveChangeRequest.GetOpts`](#method-observechangerequest-getopts)</span>
- <span class="api-method">[`ObserveChangeRequest.GetRegion`](#method-observechangerequest-getregion)</span>
- <span class="api-method">[`ObserveChangeRequest.GetSource`](#method-observechangerequest-getsource)</span>
- <span class="api-method">[`ObserveChangeRequest.ProtoMessage`](#method-observechangerequest-protomessage)</span>
- <span class="api-method">[`ObserveChangeRequest.ProtoReflect`](#method-observechangerequest-protoreflect)</span>
- <span class="api-method">[`ObserveChangeRequest.Reset`](#method-observechangerequest-reset)</span>
- <span class="api-method">[`ObserveChangeRequest.String`](#method-observechangerequest-string)</span>
- <span class="api-method">[`ObserveEvent.Descriptor`](#method-observeevent-descriptor)</span>
- <span class="api-method">[`ObserveEvent.GetMatch`](#method-observeevent-getmatch)</span>
- <span class="api-method">[`ObserveEvent.GetTimestampUnixMillis`](#method-observeevent-gettimestampunixmillis)</span>
- <span class="api-method">[`ObserveEvent.GetType`](#method-observeevent-gettype)</span>
- <span class="api-method">[`ObserveEvent.ProtoMessage`](#method-observeevent-protomessage)</span>
- <span class="api-method">[`ObserveEvent.ProtoReflect`](#method-observeevent-protoreflect)</span>
- <span class="api-method">[`ObserveEvent.Reset`](#method-observeevent-reset)</span>
- <span class="api-method">[`ObserveEvent.String`](#method-observeevent-string)</span>
- <span class="api-method">[`ObserveOptions.Descriptor`](#method-observeoptions-descriptor)</span>
- <span class="api-method">[`ObserveOptions.GetIntervalMillis`](#method-observeoptions-getintervalmillis)</span>
- <span class="api-method">[`ObserveOptions.GetTimeoutMillis`](#method-observeoptions-gettimeoutmillis)</span>
- <span class="api-method">[`ObserveOptions.ProtoMessage`](#method-observeoptions-protomessage)</span>
- <span class="api-method">[`ObserveOptions.ProtoReflect`](#method-observeoptions-protoreflect)</span>
- <span class="api-method">[`ObserveOptions.Reset`](#method-observeoptions-reset)</span>
- <span class="api-method">[`ObserveOptions.String`](#method-observeoptions-string)</span>
- <span class="api-method">[`ObserveRequest.Descriptor`](#method-observerequest-descriptor)</span>
- <span class="api-method">[`ObserveRequest.GetOpts`](#method-observerequest-getopts)</span>
- <span class="api-method">[`ObserveRequest.GetPattern`](#method-observerequest-getpattern)</span>
- <span class="api-method">[`ObserveRequest.GetRegion`](#method-observerequest-getregion)</span>
- <span class="api-method">[`ObserveRequest.GetSource`](#method-observerequest-getsource)</span>
- <span class="api-method">[`ObserveRequest.ProtoMessage`](#method-observerequest-protomessage)</span>
- <span class="api-method">[`ObserveRequest.ProtoReflect`](#method-observerequest-protoreflect)</span>
- <span class="api-method">[`ObserveRequest.Reset`](#method-observerequest-reset)</span>
- <span class="api-method">[`ObserveRequest.String`](#method-observerequest-string)</span>
- <span class="api-method">[`ObserveResponse.Descriptor`](#method-observeresponse-descriptor)</span>
- <span class="api-method">[`ObserveResponse.GetEvents`](#method-observeresponse-getevents)</span>
- <span class="api-method">[`ObserveResponse.ProtoMessage`](#method-observeresponse-protomessage)</span>
- <span class="api-method">[`ObserveResponse.ProtoReflect`](#method-observeresponse-protoreflect)</span>
- <span class="api-method">[`ObserveResponse.Reset`](#method-observeresponse-reset)</span>
- <span class="api-method">[`ObserveResponse.String`](#method-observeresponse-string)</span>
- <span class="api-method">[`Pattern.Descriptor`](#method-pattern-descriptor)</span>
- <span class="api-method">[`Pattern.GetExact`](#method-pattern-getexact)</span>
- <span class="api-method">[`Pattern.GetImage`](#method-pattern-getimage)</span>
- <span class="api-method">[`Pattern.GetMask`](#method-pattern-getmask)</span>
- <span class="api-method">[`Pattern.GetResizeFactor`](#method-pattern-getresizefactor)</span>
- <span class="api-method">[`Pattern.GetSimilarity`](#method-pattern-getsimilarity)</span>
- <span class="api-method">[`Pattern.GetTargetOffset`](#method-pattern-gettargetoffset)</span>
- <span class="api-method">[`Pattern.ProtoMessage`](#method-pattern-protomessage)</span>
- <span class="api-method">[`Pattern.ProtoReflect`](#method-pattern-protoreflect)</span>
- <span class="api-method">[`Pattern.Reset`](#method-pattern-reset)</span>
- <span class="api-method">[`Pattern.String`](#method-pattern-string)</span>
- <span class="api-method">[`Point.Descriptor`](#method-point-descriptor)</span>
- <span class="api-method">[`Point.GetX`](#method-point-getx)</span>
- <span class="api-method">[`Point.GetY`](#method-point-gety)</span>
- <span class="api-method">[`Point.ProtoMessage`](#method-point-protomessage)</span>
- <span class="api-method">[`Point.ProtoReflect`](#method-point-protoreflect)</span>
- <span class="api-method">[`Point.Reset`](#method-point-reset)</span>
- <span class="api-method">[`Point.String`](#method-point-string)</span>
- <span class="api-method">[`ReadTextRequest.Descriptor`](#method-readtextrequest-descriptor)</span>
- <span class="api-method">[`ReadTextRequest.GetParams`](#method-readtextrequest-getparams)</span>
- <span class="api-method">[`ReadTextRequest.GetSource`](#method-readtextrequest-getsource)</span>
- <span class="api-method">[`ReadTextRequest.ProtoMessage`](#method-readtextrequest-protomessage)</span>
- <span class="api-method">[`ReadTextRequest.ProtoReflect`](#method-readtextrequest-protoreflect)</span>
- <span class="api-method">[`ReadTextRequest.Reset`](#method-readtextrequest-reset)</span>
- <span class="api-method">[`ReadTextRequest.String`](#method-readtextrequest-string)</span>
- <span class="api-method">[`ReadTextResponse.Descriptor`](#method-readtextresponse-descriptor)</span>
- <span class="api-method">[`ReadTextResponse.GetText`](#method-readtextresponse-gettext)</span>
- <span class="api-method">[`ReadTextResponse.ProtoMessage`](#method-readtextresponse-protomessage)</span>
- <span class="api-method">[`ReadTextResponse.ProtoReflect`](#method-readtextresponse-protoreflect)</span>
- <span class="api-method">[`ReadTextResponse.Reset`](#method-readtextresponse-reset)</span>
- <span class="api-method">[`ReadTextResponse.String`](#method-readtextresponse-string)</span>
- <span class="api-method">[`Rect.Descriptor`](#method-rect-descriptor)</span>
- <span class="api-method">[`Rect.GetH`](#method-rect-geth)</span>
- <span class="api-method">[`Rect.GetW`](#method-rect-getw)</span>
- <span class="api-method">[`Rect.GetX`](#method-rect-getx)</span>
- <span class="api-method">[`Rect.GetY`](#method-rect-gety)</span>
- <span class="api-method">[`Rect.ProtoMessage`](#method-rect-protomessage)</span>
- <span class="api-method">[`Rect.ProtoReflect`](#method-rect-protoreflect)</span>
- <span class="api-method">[`Rect.Reset`](#method-rect-reset)</span>
- <span class="api-method">[`Rect.String`](#method-rect-string)</span>
- <span class="api-method">[`ScreenDescriptor.Descriptor`](#method-screendescriptor-descriptor)</span>
- <span class="api-method">[`ScreenDescriptor.GetBounds`](#method-screendescriptor-getbounds)</span>
- <span class="api-method">[`ScreenDescriptor.GetId`](#method-screendescriptor-getid)</span>
- <span class="api-method">[`ScreenDescriptor.GetName`](#method-screendescriptor-getname)</span>
- <span class="api-method">[`ScreenDescriptor.GetPrimary`](#method-screendescriptor-getprimary)</span>
- <span class="api-method">[`ScreenDescriptor.ProtoMessage`](#method-screendescriptor-protomessage)</span>
- <span class="api-method">[`ScreenDescriptor.ProtoReflect`](#method-screendescriptor-protoreflect)</span>
- <span class="api-method">[`ScreenDescriptor.Reset`](#method-screendescriptor-reset)</span>
- <span class="api-method">[`ScreenDescriptor.String`](#method-screendescriptor-string)</span>
- <span class="api-method">[`ScreenQueryOptions.Descriptor`](#method-screenqueryoptions-descriptor)</span>
- <span class="api-method">[`ScreenQueryOptions.GetIntervalMillis`](#method-screenqueryoptions-getintervalmillis)</span>
- <span class="api-method">[`ScreenQueryOptions.GetMatcherEngine`](#method-screenqueryoptions-getmatcherengine)</span>
- <span class="api-method">[`ScreenQueryOptions.GetRegion`](#method-screenqueryoptions-getregion)</span>
- <span class="api-method">[`ScreenQueryOptions.GetScreenId`](#method-screenqueryoptions-getscreenid)</span>
- <span class="api-method">[`ScreenQueryOptions.GetTimeoutMillis`](#method-screenqueryoptions-gettimeoutmillis)</span>
- <span class="api-method">[`ScreenQueryOptions.ProtoMessage`](#method-screenqueryoptions-protomessage)</span>
- <span class="api-method">[`ScreenQueryOptions.ProtoReflect`](#method-screenqueryoptions-protoreflect)</span>
- <span class="api-method">[`ScreenQueryOptions.Reset`](#method-screenqueryoptions-reset)</span>
- <span class="api-method">[`ScreenQueryOptions.String`](#method-screenqueryoptions-string)</span>
- <span class="api-method">[`ScrollWheelRequest.Descriptor`](#method-scrollwheelrequest-descriptor)</span>
- <span class="api-method">[`ScrollWheelRequest.GetDirection`](#method-scrollwheelrequest-getdirection)</span>
- <span class="api-method">[`ScrollWheelRequest.GetOpts`](#method-scrollwheelrequest-getopts)</span>
- <span class="api-method">[`ScrollWheelRequest.GetSteps`](#method-scrollwheelrequest-getsteps)</span>
- <span class="api-method">[`ScrollWheelRequest.GetX`](#method-scrollwheelrequest-getx)</span>
- <span class="api-method">[`ScrollWheelRequest.GetY`](#method-scrollwheelrequest-gety)</span>
- <span class="api-method">[`ScrollWheelRequest.ProtoMessage`](#method-scrollwheelrequest-protomessage)</span>
- <span class="api-method">[`ScrollWheelRequest.ProtoReflect`](#method-scrollwheelrequest-protoreflect)</span>
- <span class="api-method">[`ScrollWheelRequest.Reset`](#method-scrollwheelrequest-reset)</span>
- <span class="api-method">[`ScrollWheelRequest.String`](#method-scrollwheelrequest-string)</span>
- <span class="api-method">[`TextMatch.Descriptor`](#method-textmatch-descriptor)</span>
- <span class="api-method">[`TextMatch.GetConfidence`](#method-textmatch-getconfidence)</span>
- <span class="api-method">[`TextMatch.GetIndex`](#method-textmatch-getindex)</span>
- <span class="api-method">[`TextMatch.GetRect`](#method-textmatch-getrect)</span>
- <span class="api-method">[`TextMatch.GetText`](#method-textmatch-gettext)</span>
- <span class="api-method">[`TextMatch.ProtoMessage`](#method-textmatch-protomessage)</span>
- <span class="api-method">[`TextMatch.ProtoReflect`](#method-textmatch-protoreflect)</span>
- <span class="api-method">[`TextMatch.Reset`](#method-textmatch-reset)</span>
- <span class="api-method">[`TextMatch.String`](#method-textmatch-string)</span>
- <span class="api-method">[`TypeTextRequest.Descriptor`](#method-typetextrequest-descriptor)</span>
- <span class="api-method">[`TypeTextRequest.GetOpts`](#method-typetextrequest-getopts)</span>
- <span class="api-method">[`TypeTextRequest.GetText`](#method-typetextrequest-gettext)</span>
- <span class="api-method">[`TypeTextRequest.ProtoMessage`](#method-typetextrequest-protomessage)</span>
- <span class="api-method">[`TypeTextRequest.ProtoReflect`](#method-typetextrequest-protoreflect)</span>
- <span class="api-method">[`TypeTextRequest.Reset`](#method-typetextrequest-reset)</span>
- <span class="api-method">[`TypeTextRequest.String`](#method-typetextrequest-string)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.CaptureScreen`](#method-unimplementedsikuliserviceserver-capturescreen)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.Click`](#method-unimplementedsikuliserviceserver-click)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ClickOnScreen`](#method-unimplementedsikuliserviceserver-clickonscreen)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.CloseApp`](#method-unimplementedsikuliserviceserver-closeapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ExistsOnScreen`](#method-unimplementedsikuliserviceserver-existsonscreen)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.Find`](#method-unimplementedsikuliserviceserver-find)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindAll`](#method-unimplementedsikuliserviceserver-findall)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindOnScreen`](#method-unimplementedsikuliserviceserver-findonscreen)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindText`](#method-unimplementedsikuliserviceserver-findtext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindWindows`](#method-unimplementedsikuliserviceserver-findwindows)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FocusApp`](#method-unimplementedsikuliserviceserver-focusapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.GetFocusedWindow`](#method-unimplementedsikuliserviceserver-getfocusedwindow)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.GetPrimaryScreen`](#method-unimplementedsikuliserviceserver-getprimaryscreen)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.GetWindow`](#method-unimplementedsikuliserviceserver-getwindow)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.Hotkey`](#method-unimplementedsikuliserviceserver-hotkey)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.IsAppRunning`](#method-unimplementedsikuliserviceserver-isapprunning)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.KeyDown`](#method-unimplementedsikuliserviceserver-keydown)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.KeyUp`](#method-unimplementedsikuliserviceserver-keyup)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ListScreens`](#method-unimplementedsikuliserviceserver-listscreens)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ListWindows`](#method-unimplementedsikuliserviceserver-listwindows)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.MouseDown`](#method-unimplementedsikuliserviceserver-mousedown)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.MouseUp`](#method-unimplementedsikuliserviceserver-mouseup)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.MoveMouse`](#method-unimplementedsikuliserviceserver-movemouse)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveAppear`](#method-unimplementedsikuliserviceserver-observeappear)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveChange`](#method-unimplementedsikuliserviceserver-observechange)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveVanish`](#method-unimplementedsikuliserviceserver-observevanish)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.OpenApp`](#method-unimplementedsikuliserviceserver-openapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.PasteText`](#method-unimplementedsikuliserviceserver-pastetext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ReadText`](#method-unimplementedsikuliserviceserver-readtext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ScrollWheel`](#method-unimplementedsikuliserviceserver-scrollwheel)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.TypeText`](#method-unimplementedsikuliserviceserver-typetext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.WaitOnScreen`](#method-unimplementedsikuliserviceserver-waitonscreen)</span>
- <span class="api-method">[`WaitOnScreenRequest.Descriptor`](#method-waitonscreenrequest-descriptor)</span>
- <span class="api-method">[`WaitOnScreenRequest.GetOpts`](#method-waitonscreenrequest-getopts)</span>
- <span class="api-method">[`WaitOnScreenRequest.GetPattern`](#method-waitonscreenrequest-getpattern)</span>
- <span class="api-method">[`WaitOnScreenRequest.ProtoMessage`](#method-waitonscreenrequest-protomessage)</span>
- <span class="api-method">[`WaitOnScreenRequest.ProtoReflect`](#method-waitonscreenrequest-protoreflect)</span>
- <span class="api-method">[`WaitOnScreenRequest.Reset`](#method-waitonscreenrequest-reset)</span>
- <span class="api-method">[`WaitOnScreenRequest.String`](#method-waitonscreenrequest-string)</span>
- <span class="api-method">[`Window.Descriptor`](#method-window-descriptor)</span>
- <span class="api-method">[`Window.GetApp`](#method-window-getapp)</span>
- <span class="api-method">[`Window.GetBounds`](#method-window-getbounds)</span>
- <span class="api-method">[`Window.GetFocused`](#method-window-getfocused)</span>
- <span class="api-method">[`Window.GetId`](#method-window-getid)</span>
- <span class="api-method">[`Window.GetPid`](#method-window-getpid)</span>
- <span class="api-method">[`Window.GetTitle`](#method-window-gettitle)</span>
- <span class="api-method">[`Window.ProtoMessage`](#method-window-protomessage)</span>
- <span class="api-method">[`Window.ProtoReflect`](#method-window-protoreflect)</span>
- <span class="api-method">[`Window.Reset`](#method-window-reset)</span>
- <span class="api-method">[`Window.String`](#method-window-string)</span>
- <span class="api-method">[`WindowQuery.Descriptor`](#method-windowquery-descriptor)</span>
- <span class="api-method">[`WindowQuery.GetFocusedOnly`](#method-windowquery-getfocusedonly)</span>
- <span class="api-method">[`WindowQuery.GetId`](#method-windowquery-getid)</span>
- <span class="api-method">[`WindowQuery.GetIndex`](#method-windowquery-getindex)</span>
- <span class="api-method">[`WindowQuery.GetTitleContains`](#method-windowquery-gettitlecontains)</span>
- <span class="api-method">[`WindowQuery.GetTitleExact`](#method-windowquery-gettitleexact)</span>
- <span class="api-method">[`WindowQuery.ProtoMessage`](#method-windowquery-protomessage)</span>
- <span class="api-method">[`WindowQuery.ProtoReflect`](#method-windowquery-protoreflect)</span>
- <span class="api-method">[`WindowQuery.Reset`](#method-windowquery-reset)</span>
- <span class="api-method">[`WindowQuery.String`](#method-windowquery-string)</span>
- <span class="api-method">[`WindowQueryRequest.Descriptor`](#method-windowqueryrequest-descriptor)</span>
- <span class="api-method">[`WindowQueryRequest.GetName`](#method-windowqueryrequest-getname)</span>
- <span class="api-method">[`WindowQueryRequest.GetOpts`](#method-windowqueryrequest-getopts)</span>
- <span class="api-method">[`WindowQueryRequest.GetQuery`](#method-windowqueryrequest-getquery)</span>
- <span class="api-method">[`WindowQueryRequest.ProtoMessage`](#method-windowqueryrequest-protomessage)</span>
- <span class="api-method">[`WindowQueryRequest.ProtoReflect`](#method-windowqueryrequest-protoreflect)</span>
- <span class="api-method">[`WindowQueryRequest.Reset`](#method-windowqueryrequest-reset)</span>
- <span class="api-method">[`WindowQueryRequest.String`](#method-windowqueryrequest-string)</span>

## Declarations

### Types

#### <a id="type-actionresponse"></a><span class="api-type">Type</span> `ActionResponse`

- Signature: <span class="api-signature">`type ActionResponse struct {`</span>

#### <a id="type-appactionrequest"></a><span class="api-type">Type</span> `AppActionRequest`

- Signature: <span class="api-signature">`type AppActionRequest struct {`</span>

#### <a id="type-appoptions"></a><span class="api-type">Type</span> `AppOptions`

- Signature: <span class="api-signature">`type AppOptions struct {`</span>

#### <a id="type-capturescreenrequest"></a><span class="api-type">Type</span> `CaptureScreenRequest`

- Signature: <span class="api-signature">`type CaptureScreenRequest struct {`</span>

#### <a id="type-capturescreenresponse"></a><span class="api-type">Type</span> `CaptureScreenResponse`

- Signature: <span class="api-signature">`type CaptureScreenResponse struct {`</span>

#### <a id="type-clickonscreenrequest"></a><span class="api-type">Type</span> `ClickOnScreenRequest`

- Signature: <span class="api-signature">`type ClickOnScreenRequest struct {`</span>

#### <a id="type-clickrequest"></a><span class="api-type">Type</span> `ClickRequest`

- Signature: <span class="api-signature">`type ClickRequest struct {`</span>

#### <a id="type-existsonscreenrequest"></a><span class="api-type">Type</span> `ExistsOnScreenRequest`

- Signature: <span class="api-signature">`type ExistsOnScreenRequest struct {`</span>

#### <a id="type-existsonscreenresponse"></a><span class="api-type">Type</span> `ExistsOnScreenResponse`

- Signature: <span class="api-signature">`type ExistsOnScreenResponse struct {`</span>

#### <a id="type-findallresponse"></a><span class="api-type">Type</span> `FindAllResponse`

- Signature: <span class="api-signature">`type FindAllResponse struct {`</span>

#### <a id="type-findonscreenrequest"></a><span class="api-type">Type</span> `FindOnScreenRequest`

- Signature: <span class="api-signature">`type FindOnScreenRequest struct {`</span>

#### <a id="type-findrequest"></a><span class="api-type">Type</span> `FindRequest`

- Signature: <span class="api-signature">`type FindRequest struct {`</span>

#### <a id="type-findresponse"></a><span class="api-type">Type</span> `FindResponse`

- Signature: <span class="api-signature">`type FindResponse struct {`</span>

#### <a id="type-findtextrequest"></a><span class="api-type">Type</span> `FindTextRequest`

- Signature: <span class="api-signature">`type FindTextRequest struct {`</span>

#### <a id="type-findtextresponse"></a><span class="api-type">Type</span> `FindTextResponse`

- Signature: <span class="api-signature">`type FindTextResponse struct {`</span>

#### <a id="type-getprimaryscreenrequest"></a><span class="api-type">Type</span> `GetPrimaryScreenRequest`

- Signature: <span class="api-signature">`type GetPrimaryScreenRequest struct {`</span>

#### <a id="type-getprimaryscreenresponse"></a><span class="api-type">Type</span> `GetPrimaryScreenResponse`

- Signature: <span class="api-signature">`type GetPrimaryScreenResponse struct {`</span>

#### <a id="type-getwindowresponse"></a><span class="api-type">Type</span> `GetWindowResponse`

- Signature: <span class="api-signature">`type GetWindowResponse struct {`</span>

#### <a id="type-grayimage"></a><span class="api-type">Type</span> `GrayImage`

- Signature: <span class="api-signature">`type GrayImage struct {`</span>

#### <a id="type-hotkeyrequest"></a><span class="api-type">Type</span> `HotkeyRequest`

- Signature: <span class="api-signature">`type HotkeyRequest struct {`</span>

#### <a id="type-inputoptions"></a><span class="api-type">Type</span> `InputOptions`

- Signature: <span class="api-signature">`type InputOptions struct {`</span>

#### <a id="type-isapprunningresponse"></a><span class="api-type">Type</span> `IsAppRunningResponse`

- Signature: <span class="api-signature">`type IsAppRunningResponse struct {`</span>

#### <a id="type-listscreensrequest"></a><span class="api-type">Type</span> `ListScreensRequest`

- Signature: <span class="api-signature">`type ListScreensRequest struct {`</span>

#### <a id="type-listscreensresponse"></a><span class="api-type">Type</span> `ListScreensResponse`

- Signature: <span class="api-signature">`type ListScreensResponse struct {`</span>

#### <a id="type-listwindowsresponse"></a><span class="api-type">Type</span> `ListWindowsResponse`

- Signature: <span class="api-signature">`type ListWindowsResponse struct {`</span>

#### <a id="type-match"></a><span class="api-type">Type</span> `Match`

- Signature: <span class="api-signature">`type Match struct {`</span>

#### <a id="type-matcherengine"></a><span class="api-type">Type</span> `MatcherEngine`

- Signature: <span class="api-signature">`type MatcherEngine int32`</span>

#### <a id="type-movemouserequest"></a><span class="api-type">Type</span> `MoveMouseRequest`

- Signature: <span class="api-signature">`type MoveMouseRequest struct {`</span>

#### <a id="type-ocrparams"></a><span class="api-type">Type</span> `OCRParams`

- Signature: <span class="api-signature">`type OCRParams struct {`</span>

#### <a id="type-observechangerequest"></a><span class="api-type">Type</span> `ObserveChangeRequest`

- Signature: <span class="api-signature">`type ObserveChangeRequest struct {`</span>

#### <a id="type-observeevent"></a><span class="api-type">Type</span> `ObserveEvent`

- Signature: <span class="api-signature">`type ObserveEvent struct {`</span>

#### <a id="type-observeoptions"></a><span class="api-type">Type</span> `ObserveOptions`

- Signature: <span class="api-signature">`type ObserveOptions struct {`</span>

#### <a id="type-observerequest"></a><span class="api-type">Type</span> `ObserveRequest`

- Signature: <span class="api-signature">`type ObserveRequest struct {`</span>

#### <a id="type-observeresponse"></a><span class="api-type">Type</span> `ObserveResponse`

- Signature: <span class="api-signature">`type ObserveResponse struct {`</span>

#### <a id="type-pattern"></a><span class="api-type">Type</span> `Pattern`

- Signature: <span class="api-signature">`type Pattern struct {`</span>

#### <a id="type-point"></a><span class="api-type">Type</span> `Point`

- Signature: <span class="api-signature">`type Point struct {`</span>

#### <a id="type-readtextrequest"></a><span class="api-type">Type</span> `ReadTextRequest`

- Signature: <span class="api-signature">`type ReadTextRequest struct {`</span>

#### <a id="type-readtextresponse"></a><span class="api-type">Type</span> `ReadTextResponse`

- Signature: <span class="api-signature">`type ReadTextResponse struct {`</span>

#### <a id="type-rect"></a><span class="api-type">Type</span> `Rect`

- Signature: <span class="api-signature">`type Rect struct {`</span>

#### <a id="type-screendescriptor"></a><span class="api-type">Type</span> `ScreenDescriptor`

- Signature: <span class="api-signature">`type ScreenDescriptor struct {`</span>

#### <a id="type-screenqueryoptions"></a><span class="api-type">Type</span> `ScreenQueryOptions`

- Signature: <span class="api-signature">`type ScreenQueryOptions struct {`</span>

#### <a id="type-scrollwheelrequest"></a><span class="api-type">Type</span> `ScrollWheelRequest`

- Signature: <span class="api-signature">`type ScrollWheelRequest struct {`</span>

#### <a id="type-sikuliserviceclient"></a><span class="api-type">Type</span> `SikuliServiceClient`

- Signature: <span class="api-signature">`type SikuliServiceClient interface {`</span>

#### <a id="type-sikuliserviceserver"></a><span class="api-type">Type</span> `SikuliServiceServer`

- Signature: <span class="api-signature">`type SikuliServiceServer interface {`</span>

#### <a id="type-textmatch"></a><span class="api-type">Type</span> `TextMatch`

- Signature: <span class="api-signature">`type TextMatch struct {`</span>

#### <a id="type-typetextrequest"></a><span class="api-type">Type</span> `TypeTextRequest`

- Signature: <span class="api-signature">`type TypeTextRequest struct {`</span>

#### <a id="type-unimplementedsikuliserviceserver"></a><span class="api-type">Type</span> `UnimplementedSikuliServiceServer`

- Signature: <span class="api-signature">`type UnimplementedSikuliServiceServer struct{}`</span>
- Notes: UnimplementedSikuliServiceServer must be embedded to have forward compatible implementations.

#### <a id="type-unsafesikuliserviceserver"></a><span class="api-type">Type</span> `UnsafeSikuliServiceServer`

- Signature: <span class="api-signature">`type UnsafeSikuliServiceServer interface {`</span>

#### <a id="type-waitonscreenrequest"></a><span class="api-type">Type</span> `WaitOnScreenRequest`

- Signature: <span class="api-signature">`type WaitOnScreenRequest struct {`</span>

#### <a id="type-window"></a><span class="api-type">Type</span> `Window`

- Signature: <span class="api-signature">`type Window struct {`</span>

#### <a id="type-windowquery"></a><span class="api-type">Type</span> `WindowQuery`

- Signature: <span class="api-signature">`type WindowQuery struct {`</span>

#### <a id="type-windowqueryrequest"></a><span class="api-type">Type</span> `WindowQueryRequest`

- Signature: <span class="api-signature">`type WindowQueryRequest struct {`</span>

### Functions

#### <a id="func-registersikuliserviceserver"></a><span class="api-func">Function</span> `RegisterSikuliServiceServer`

- Signature: <span class="api-signature">`func RegisterSikuliServiceServer(s grpc.ServiceRegistrar, srv SikuliServiceServer)`</span>
- Uses: [`SikuliServiceServer`](#type-sikuliserviceserver)

#### <a id="func-newsikuliserviceclient"></a><span class="api-func">Function</span> `NewSikuliServiceClient`

- Signature: <span class="api-signature">`func NewSikuliServiceClient(cc grpc.ClientConnInterface) SikuliServiceClient`</span>
- Uses: [`SikuliServiceClient`](#type-sikuliserviceclient)

### Methods

#### <a id="method-actionresponse-descriptor"></a><span class="api-method">Method</span> `ActionResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ActionResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.

#### <a id="method-actionresponse-protomessage"></a><span class="api-method">Method</span> `ActionResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ActionResponse) ProtoMessage()`</span>

#### <a id="method-actionresponse-protoreflect"></a><span class="api-method">Method</span> `ActionResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ActionResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-actionresponse-reset"></a><span class="api-method">Method</span> `ActionResponse.Reset`

- Signature: <span class="api-signature">`func (x *ActionResponse) Reset()`</span>

#### <a id="method-actionresponse-string"></a><span class="api-method">Method</span> `ActionResponse.String`

- Signature: <span class="api-signature">`func (x *ActionResponse) String() string`</span>

#### <a id="method-appactionrequest-descriptor"></a><span class="api-method">Method</span> `AppActionRequest.Descriptor`

- Signature: <span class="api-signature">`func (*AppActionRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use AppActionRequest.ProtoReflect.Descriptor instead.

#### <a id="method-appactionrequest-getargs"></a><span class="api-method">Method</span> `AppActionRequest.GetArgs`

- Signature: <span class="api-signature">`func (x *AppActionRequest) GetArgs() []string`</span>

#### <a id="method-appactionrequest-getname"></a><span class="api-method">Method</span> `AppActionRequest.GetName`

- Signature: <span class="api-signature">`func (x *AppActionRequest) GetName() string`</span>

#### <a id="method-appactionrequest-getopts"></a><span class="api-method">Method</span> `AppActionRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *AppActionRequest) GetOpts() *AppOptions`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appactionrequest-protomessage"></a><span class="api-method">Method</span> `AppActionRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*AppActionRequest) ProtoMessage()`</span>

#### <a id="method-appactionrequest-protoreflect"></a><span class="api-method">Method</span> `AppActionRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *AppActionRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-appactionrequest-reset"></a><span class="api-method">Method</span> `AppActionRequest.Reset`

- Signature: <span class="api-signature">`func (x *AppActionRequest) Reset()`</span>

#### <a id="method-appactionrequest-string"></a><span class="api-method">Method</span> `AppActionRequest.String`

- Signature: <span class="api-signature">`func (x *AppActionRequest) String() string`</span>

#### <a id="method-appoptions-descriptor"></a><span class="api-method">Method</span> `AppOptions.Descriptor`

- Signature: <span class="api-signature">`func (*AppOptions) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use AppOptions.ProtoReflect.Descriptor instead.

#### <a id="method-appoptions-gettimeoutmillis"></a><span class="api-method">Method</span> `AppOptions.GetTimeoutMillis`

- Signature: <span class="api-signature">`func (x *AppOptions) GetTimeoutMillis() int64`</span>

#### <a id="method-appoptions-protomessage"></a><span class="api-method">Method</span> `AppOptions.ProtoMessage`

- Signature: <span class="api-signature">`func (*AppOptions) ProtoMessage()`</span>

#### <a id="method-appoptions-protoreflect"></a><span class="api-method">Method</span> `AppOptions.ProtoReflect`

- Signature: <span class="api-signature">`func (x *AppOptions) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-appoptions-reset"></a><span class="api-method">Method</span> `AppOptions.Reset`

- Signature: <span class="api-signature">`func (x *AppOptions) Reset()`</span>

#### <a id="method-appoptions-string"></a><span class="api-method">Method</span> `AppOptions.String`

- Signature: <span class="api-signature">`func (x *AppOptions) String() string`</span>

#### <a id="method-capturescreenrequest-descriptor"></a><span class="api-method">Method</span> `CaptureScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*CaptureScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use CaptureScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-capturescreenrequest-getregion"></a><span class="api-method">Method</span> `CaptureScreenRequest.GetRegion`

- Signature: <span class="api-signature">`func (x *CaptureScreenRequest) GetRegion() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-capturescreenrequest-getscreenid"></a><span class="api-method">Method</span> `CaptureScreenRequest.GetScreenId`

- Signature: <span class="api-signature">`func (x *CaptureScreenRequest) GetScreenId() int32`</span>

#### <a id="method-capturescreenrequest-protomessage"></a><span class="api-method">Method</span> `CaptureScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*CaptureScreenRequest) ProtoMessage()`</span>

#### <a id="method-capturescreenrequest-protoreflect"></a><span class="api-method">Method</span> `CaptureScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *CaptureScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-capturescreenrequest-reset"></a><span class="api-method">Method</span> `CaptureScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *CaptureScreenRequest) Reset()`</span>

#### <a id="method-capturescreenrequest-string"></a><span class="api-method">Method</span> `CaptureScreenRequest.String`

- Signature: <span class="api-signature">`func (x *CaptureScreenRequest) String() string`</span>

#### <a id="method-capturescreenresponse-descriptor"></a><span class="api-method">Method</span> `CaptureScreenResponse.Descriptor`

- Signature: <span class="api-signature">`func (*CaptureScreenResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use CaptureScreenResponse.ProtoReflect.Descriptor instead.

#### <a id="method-capturescreenresponse-getimage"></a><span class="api-method">Method</span> `CaptureScreenResponse.GetImage`

- Signature: <span class="api-signature">`func (x *CaptureScreenResponse) GetImage() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-capturescreenresponse-getscreen"></a><span class="api-method">Method</span> `CaptureScreenResponse.GetScreen`

- Signature: <span class="api-signature">`func (x *CaptureScreenResponse) GetScreen() *ScreenDescriptor`</span>
- Uses: [`ScreenDescriptor`](#type-screendescriptor)

#### <a id="method-capturescreenresponse-protomessage"></a><span class="api-method">Method</span> `CaptureScreenResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*CaptureScreenResponse) ProtoMessage()`</span>

#### <a id="method-capturescreenresponse-protoreflect"></a><span class="api-method">Method</span> `CaptureScreenResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *CaptureScreenResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-capturescreenresponse-reset"></a><span class="api-method">Method</span> `CaptureScreenResponse.Reset`

- Signature: <span class="api-signature">`func (x *CaptureScreenResponse) Reset()`</span>

#### <a id="method-capturescreenresponse-string"></a><span class="api-method">Method</span> `CaptureScreenResponse.String`

- Signature: <span class="api-signature">`func (x *CaptureScreenResponse) String() string`</span>

#### <a id="method-clickonscreenrequest-descriptor"></a><span class="api-method">Method</span> `ClickOnScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ClickOnScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ClickOnScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-clickonscreenrequest-getclickopts"></a><span class="api-method">Method</span> `ClickOnScreenRequest.GetClickOpts`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) GetClickOpts() *InputOptions`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-clickonscreenrequest-getopts"></a><span class="api-method">Method</span> `ClickOnScreenRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) GetOpts() *ScreenQueryOptions`</span>
- Uses: [`ScreenQueryOptions`](#type-screenqueryoptions)

#### <a id="method-clickonscreenrequest-getpattern"></a><span class="api-method">Method</span> `ClickOnScreenRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-clickonscreenrequest-protomessage"></a><span class="api-method">Method</span> `ClickOnScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ClickOnScreenRequest) ProtoMessage()`</span>

#### <a id="method-clickonscreenrequest-protoreflect"></a><span class="api-method">Method</span> `ClickOnScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-clickonscreenrequest-reset"></a><span class="api-method">Method</span> `ClickOnScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) Reset()`</span>

#### <a id="method-clickonscreenrequest-string"></a><span class="api-method">Method</span> `ClickOnScreenRequest.String`

- Signature: <span class="api-signature">`func (x *ClickOnScreenRequest) String() string`</span>

#### <a id="method-clickrequest-descriptor"></a><span class="api-method">Method</span> `ClickRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ClickRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ClickRequest.ProtoReflect.Descriptor instead.

#### <a id="method-clickrequest-getopts"></a><span class="api-method">Method</span> `ClickRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ClickRequest) GetOpts() *InputOptions`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-clickrequest-getx"></a><span class="api-method">Method</span> `ClickRequest.GetX`

- Signature: <span class="api-signature">`func (x *ClickRequest) GetX() int32`</span>

#### <a id="method-clickrequest-gety"></a><span class="api-method">Method</span> `ClickRequest.GetY`

- Signature: <span class="api-signature">`func (x *ClickRequest) GetY() int32`</span>

#### <a id="method-clickrequest-protomessage"></a><span class="api-method">Method</span> `ClickRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ClickRequest) ProtoMessage()`</span>

#### <a id="method-clickrequest-protoreflect"></a><span class="api-method">Method</span> `ClickRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ClickRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-clickrequest-reset"></a><span class="api-method">Method</span> `ClickRequest.Reset`

- Signature: <span class="api-signature">`func (x *ClickRequest) Reset()`</span>

#### <a id="method-clickrequest-string"></a><span class="api-method">Method</span> `ClickRequest.String`

- Signature: <span class="api-signature">`func (x *ClickRequest) String() string`</span>

#### <a id="method-existsonscreenrequest-descriptor"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ExistsOnScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ExistsOnScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-existsonscreenrequest-getopts"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenRequest) GetOpts() *ScreenQueryOptions`</span>
- Uses: [`ScreenQueryOptions`](#type-screenqueryoptions)

#### <a id="method-existsonscreenrequest-getpattern"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-existsonscreenrequest-protomessage"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ExistsOnScreenRequest) ProtoMessage()`</span>

#### <a id="method-existsonscreenrequest-protoreflect"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-existsonscreenrequest-reset"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenRequest) Reset()`</span>

#### <a id="method-existsonscreenrequest-string"></a><span class="api-method">Method</span> `ExistsOnScreenRequest.String`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenRequest) String() string`</span>

#### <a id="method-existsonscreenresponse-descriptor"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ExistsOnScreenResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ExistsOnScreenResponse.ProtoReflect.Descriptor instead.

#### <a id="method-existsonscreenresponse-getexists"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.GetExists`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenResponse) GetExists() bool`</span>

#### <a id="method-existsonscreenresponse-getmatch"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.GetMatch`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenResponse) GetMatch() *Match`</span>
- Uses: [`Match`](#type-match)

#### <a id="method-existsonscreenresponse-protomessage"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ExistsOnScreenResponse) ProtoMessage()`</span>

#### <a id="method-existsonscreenresponse-protoreflect"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-existsonscreenresponse-reset"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.Reset`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenResponse) Reset()`</span>

#### <a id="method-existsonscreenresponse-string"></a><span class="api-method">Method</span> `ExistsOnScreenResponse.String`

- Signature: <span class="api-signature">`func (x *ExistsOnScreenResponse) String() string`</span>

#### <a id="method-findallresponse-descriptor"></a><span class="api-method">Method</span> `FindAllResponse.Descriptor`

- Signature: <span class="api-signature">`func (*FindAllResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.

#### <a id="method-findallresponse-getmatches"></a><span class="api-method">Method</span> `FindAllResponse.GetMatches`

- Signature: <span class="api-signature">`func (x *FindAllResponse) GetMatches() []*Match`</span>
- Uses: [`Match`](#type-match)

#### <a id="method-findallresponse-protomessage"></a><span class="api-method">Method</span> `FindAllResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindAllResponse) ProtoMessage()`</span>

#### <a id="method-findallresponse-protoreflect"></a><span class="api-method">Method</span> `FindAllResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindAllResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findallresponse-reset"></a><span class="api-method">Method</span> `FindAllResponse.Reset`

- Signature: <span class="api-signature">`func (x *FindAllResponse) Reset()`</span>

#### <a id="method-findallresponse-string"></a><span class="api-method">Method</span> `FindAllResponse.String`

- Signature: <span class="api-signature">`func (x *FindAllResponse) String() string`</span>

#### <a id="method-findonscreenrequest-descriptor"></a><span class="api-method">Method</span> `FindOnScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*FindOnScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindOnScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-findonscreenrequest-getopts"></a><span class="api-method">Method</span> `FindOnScreenRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *FindOnScreenRequest) GetOpts() *ScreenQueryOptions`</span>
- Uses: [`ScreenQueryOptions`](#type-screenqueryoptions)

#### <a id="method-findonscreenrequest-getpattern"></a><span class="api-method">Method</span> `FindOnScreenRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *FindOnScreenRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-findonscreenrequest-protomessage"></a><span class="api-method">Method</span> `FindOnScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindOnScreenRequest) ProtoMessage()`</span>

#### <a id="method-findonscreenrequest-protoreflect"></a><span class="api-method">Method</span> `FindOnScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindOnScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findonscreenrequest-reset"></a><span class="api-method">Method</span> `FindOnScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *FindOnScreenRequest) Reset()`</span>

#### <a id="method-findonscreenrequest-string"></a><span class="api-method">Method</span> `FindOnScreenRequest.String`

- Signature: <span class="api-signature">`func (x *FindOnScreenRequest) String() string`</span>

#### <a id="method-findrequest-descriptor"></a><span class="api-method">Method</span> `FindRequest.Descriptor`

- Signature: <span class="api-signature">`func (*FindRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.

#### <a id="method-findrequest-getmatcherengine"></a><span class="api-method">Method</span> `FindRequest.GetMatcherEngine`

- Signature: <span class="api-signature">`func (x *FindRequest) GetMatcherEngine() MatcherEngine`</span>
- Uses: [`MatcherEngine`](#type-matcherengine)

#### <a id="method-findrequest-getpattern"></a><span class="api-method">Method</span> `FindRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *FindRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-findrequest-getsource"></a><span class="api-method">Method</span> `FindRequest.GetSource`

- Signature: <span class="api-signature">`func (x *FindRequest) GetSource() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-findrequest-protomessage"></a><span class="api-method">Method</span> `FindRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindRequest) ProtoMessage()`</span>

#### <a id="method-findrequest-protoreflect"></a><span class="api-method">Method</span> `FindRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findrequest-reset"></a><span class="api-method">Method</span> `FindRequest.Reset`

- Signature: <span class="api-signature">`func (x *FindRequest) Reset()`</span>

#### <a id="method-findrequest-string"></a><span class="api-method">Method</span> `FindRequest.String`

- Signature: <span class="api-signature">`func (x *FindRequest) String() string`</span>

#### <a id="method-findresponse-descriptor"></a><span class="api-method">Method</span> `FindResponse.Descriptor`

- Signature: <span class="api-signature">`func (*FindResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.

#### <a id="method-findresponse-getmatch"></a><span class="api-method">Method</span> `FindResponse.GetMatch`

- Signature: <span class="api-signature">`func (x *FindResponse) GetMatch() *Match`</span>
- Uses: [`Match`](#type-match)

#### <a id="method-findresponse-protomessage"></a><span class="api-method">Method</span> `FindResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindResponse) ProtoMessage()`</span>

#### <a id="method-findresponse-protoreflect"></a><span class="api-method">Method</span> `FindResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findresponse-reset"></a><span class="api-method">Method</span> `FindResponse.Reset`

- Signature: <span class="api-signature">`func (x *FindResponse) Reset()`</span>

#### <a id="method-findresponse-string"></a><span class="api-method">Method</span> `FindResponse.String`

- Signature: <span class="api-signature">`func (x *FindResponse) String() string`</span>

#### <a id="method-findtextrequest-descriptor"></a><span class="api-method">Method</span> `FindTextRequest.Descriptor`

- Signature: <span class="api-signature">`func (*FindTextRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindTextRequest.ProtoReflect.Descriptor instead.

#### <a id="method-findtextrequest-getparams"></a><span class="api-method">Method</span> `FindTextRequest.GetParams`

- Signature: <span class="api-signature">`func (x *FindTextRequest) GetParams() *OCRParams`</span>
- Uses: [`OCRParams`](#type-ocrparams)

#### <a id="method-findtextrequest-getquery"></a><span class="api-method">Method</span> `FindTextRequest.GetQuery`

- Signature: <span class="api-signature">`func (x *FindTextRequest) GetQuery() string`</span>

#### <a id="method-findtextrequest-getsource"></a><span class="api-method">Method</span> `FindTextRequest.GetSource`

- Signature: <span class="api-signature">`func (x *FindTextRequest) GetSource() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-findtextrequest-protomessage"></a><span class="api-method">Method</span> `FindTextRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindTextRequest) ProtoMessage()`</span>

#### <a id="method-findtextrequest-protoreflect"></a><span class="api-method">Method</span> `FindTextRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindTextRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findtextrequest-reset"></a><span class="api-method">Method</span> `FindTextRequest.Reset`

- Signature: <span class="api-signature">`func (x *FindTextRequest) Reset()`</span>

#### <a id="method-findtextrequest-string"></a><span class="api-method">Method</span> `FindTextRequest.String`

- Signature: <span class="api-signature">`func (x *FindTextRequest) String() string`</span>

#### <a id="method-findtextresponse-descriptor"></a><span class="api-method">Method</span> `FindTextResponse.Descriptor`

- Signature: <span class="api-signature">`func (*FindTextResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindTextResponse.ProtoReflect.Descriptor instead.

#### <a id="method-findtextresponse-getmatches"></a><span class="api-method">Method</span> `FindTextResponse.GetMatches`

- Signature: <span class="api-signature">`func (x *FindTextResponse) GetMatches() []*TextMatch`</span>
- Uses: [`TextMatch`](#type-textmatch)

#### <a id="method-findtextresponse-protomessage"></a><span class="api-method">Method</span> `FindTextResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*FindTextResponse) ProtoMessage()`</span>

#### <a id="method-findtextresponse-protoreflect"></a><span class="api-method">Method</span> `FindTextResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *FindTextResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-findtextresponse-reset"></a><span class="api-method">Method</span> `FindTextResponse.Reset`

- Signature: <span class="api-signature">`func (x *FindTextResponse) Reset()`</span>

#### <a id="method-findtextresponse-string"></a><span class="api-method">Method</span> `FindTextResponse.String`

- Signature: <span class="api-signature">`func (x *FindTextResponse) String() string`</span>

#### <a id="method-getprimaryscreenrequest-descriptor"></a><span class="api-method">Method</span> `GetPrimaryScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*GetPrimaryScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use GetPrimaryScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-getprimaryscreenrequest-protomessage"></a><span class="api-method">Method</span> `GetPrimaryScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*GetPrimaryScreenRequest) ProtoMessage()`</span>

#### <a id="method-getprimaryscreenrequest-protoreflect"></a><span class="api-method">Method</span> `GetPrimaryScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-getprimaryscreenrequest-reset"></a><span class="api-method">Method</span> `GetPrimaryScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenRequest) Reset()`</span>

#### <a id="method-getprimaryscreenrequest-string"></a><span class="api-method">Method</span> `GetPrimaryScreenRequest.String`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenRequest) String() string`</span>

#### <a id="method-getprimaryscreenresponse-descriptor"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.Descriptor`

- Signature: <span class="api-signature">`func (*GetPrimaryScreenResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use GetPrimaryScreenResponse.ProtoReflect.Descriptor instead.

#### <a id="method-getprimaryscreenresponse-getscreen"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.GetScreen`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenResponse) GetScreen() *ScreenDescriptor`</span>
- Uses: [`ScreenDescriptor`](#type-screendescriptor)

#### <a id="method-getprimaryscreenresponse-protomessage"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*GetPrimaryScreenResponse) ProtoMessage()`</span>

#### <a id="method-getprimaryscreenresponse-protoreflect"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-getprimaryscreenresponse-reset"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.Reset`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenResponse) Reset()`</span>

#### <a id="method-getprimaryscreenresponse-string"></a><span class="api-method">Method</span> `GetPrimaryScreenResponse.String`

- Signature: <span class="api-signature">`func (x *GetPrimaryScreenResponse) String() string`</span>

#### <a id="method-getwindowresponse-descriptor"></a><span class="api-method">Method</span> `GetWindowResponse.Descriptor`

- Signature: <span class="api-signature">`func (*GetWindowResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use GetWindowResponse.ProtoReflect.Descriptor instead.

#### <a id="method-getwindowresponse-getfound"></a><span class="api-method">Method</span> `GetWindowResponse.GetFound`

- Signature: <span class="api-signature">`func (x *GetWindowResponse) GetFound() bool`</span>

#### <a id="method-getwindowresponse-getwindow"></a><span class="api-method">Method</span> `GetWindowResponse.GetWindow`

- Signature: <span class="api-signature">`func (x *GetWindowResponse) GetWindow() *Window`</span>
- Uses: [`Window`](#type-window)

#### <a id="method-getwindowresponse-protomessage"></a><span class="api-method">Method</span> `GetWindowResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*GetWindowResponse) ProtoMessage()`</span>

#### <a id="method-getwindowresponse-protoreflect"></a><span class="api-method">Method</span> `GetWindowResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *GetWindowResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-getwindowresponse-reset"></a><span class="api-method">Method</span> `GetWindowResponse.Reset`

- Signature: <span class="api-signature">`func (x *GetWindowResponse) Reset()`</span>

#### <a id="method-getwindowresponse-string"></a><span class="api-method">Method</span> `GetWindowResponse.String`

- Signature: <span class="api-signature">`func (x *GetWindowResponse) String() string`</span>

#### <a id="method-grayimage-descriptor"></a><span class="api-method">Method</span> `GrayImage.Descriptor`

- Signature: <span class="api-signature">`func (*GrayImage) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use GrayImage.ProtoReflect.Descriptor instead.

#### <a id="method-grayimage-getheight"></a><span class="api-method">Method</span> `GrayImage.GetHeight`

- Signature: <span class="api-signature">`func (x *GrayImage) GetHeight() int32`</span>

#### <a id="method-grayimage-getname"></a><span class="api-method">Method</span> `GrayImage.GetName`

- Signature: <span class="api-signature">`func (x *GrayImage) GetName() string`</span>

#### <a id="method-grayimage-getpix"></a><span class="api-method">Method</span> `GrayImage.GetPix`

- Signature: <span class="api-signature">`func (x *GrayImage) GetPix() []byte`</span>

#### <a id="method-grayimage-getwidth"></a><span class="api-method">Method</span> `GrayImage.GetWidth`

- Signature: <span class="api-signature">`func (x *GrayImage) GetWidth() int32`</span>

#### <a id="method-grayimage-protomessage"></a><span class="api-method">Method</span> `GrayImage.ProtoMessage`

- Signature: <span class="api-signature">`func (*GrayImage) ProtoMessage()`</span>

#### <a id="method-grayimage-protoreflect"></a><span class="api-method">Method</span> `GrayImage.ProtoReflect`

- Signature: <span class="api-signature">`func (x *GrayImage) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-grayimage-reset"></a><span class="api-method">Method</span> `GrayImage.Reset`

- Signature: <span class="api-signature">`func (x *GrayImage) Reset()`</span>

#### <a id="method-grayimage-string"></a><span class="api-method">Method</span> `GrayImage.String`

- Signature: <span class="api-signature">`func (x *GrayImage) String() string`</span>

#### <a id="method-hotkeyrequest-descriptor"></a><span class="api-method">Method</span> `HotkeyRequest.Descriptor`

- Signature: <span class="api-signature">`func (*HotkeyRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use HotkeyRequest.ProtoReflect.Descriptor instead.

#### <a id="method-hotkeyrequest-getkeys"></a><span class="api-method">Method</span> `HotkeyRequest.GetKeys`

- Signature: <span class="api-signature">`func (x *HotkeyRequest) GetKeys() []string`</span>

#### <a id="method-hotkeyrequest-protomessage"></a><span class="api-method">Method</span> `HotkeyRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*HotkeyRequest) ProtoMessage()`</span>

#### <a id="method-hotkeyrequest-protoreflect"></a><span class="api-method">Method</span> `HotkeyRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *HotkeyRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-hotkeyrequest-reset"></a><span class="api-method">Method</span> `HotkeyRequest.Reset`

- Signature: <span class="api-signature">`func (x *HotkeyRequest) Reset()`</span>

#### <a id="method-hotkeyrequest-string"></a><span class="api-method">Method</span> `HotkeyRequest.String`

- Signature: <span class="api-signature">`func (x *HotkeyRequest) String() string`</span>

#### <a id="method-inputoptions-descriptor"></a><span class="api-method">Method</span> `InputOptions.Descriptor`

- Signature: <span class="api-signature">`func (*InputOptions) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use InputOptions.ProtoReflect.Descriptor instead.

#### <a id="method-inputoptions-getbutton"></a><span class="api-method">Method</span> `InputOptions.GetButton`

- Signature: <span class="api-signature">`func (x *InputOptions) GetButton() string`</span>

#### <a id="method-inputoptions-getdelaymillis"></a><span class="api-method">Method</span> `InputOptions.GetDelayMillis`

- Signature: <span class="api-signature">`func (x *InputOptions) GetDelayMillis() int64`</span>

#### <a id="method-inputoptions-protomessage"></a><span class="api-method">Method</span> `InputOptions.ProtoMessage`

- Signature: <span class="api-signature">`func (*InputOptions) ProtoMessage()`</span>

#### <a id="method-inputoptions-protoreflect"></a><span class="api-method">Method</span> `InputOptions.ProtoReflect`

- Signature: <span class="api-signature">`func (x *InputOptions) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-inputoptions-reset"></a><span class="api-method">Method</span> `InputOptions.Reset`

- Signature: <span class="api-signature">`func (x *InputOptions) Reset()`</span>

#### <a id="method-inputoptions-string"></a><span class="api-method">Method</span> `InputOptions.String`

- Signature: <span class="api-signature">`func (x *InputOptions) String() string`</span>

#### <a id="method-isapprunningresponse-descriptor"></a><span class="api-method">Method</span> `IsAppRunningResponse.Descriptor`

- Signature: <span class="api-signature">`func (*IsAppRunningResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use IsAppRunningResponse.ProtoReflect.Descriptor instead.

#### <a id="method-isapprunningresponse-getrunning"></a><span class="api-method">Method</span> `IsAppRunningResponse.GetRunning`

- Signature: <span class="api-signature">`func (x *IsAppRunningResponse) GetRunning() bool`</span>

#### <a id="method-isapprunningresponse-protomessage"></a><span class="api-method">Method</span> `IsAppRunningResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*IsAppRunningResponse) ProtoMessage()`</span>

#### <a id="method-isapprunningresponse-protoreflect"></a><span class="api-method">Method</span> `IsAppRunningResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *IsAppRunningResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-isapprunningresponse-reset"></a><span class="api-method">Method</span> `IsAppRunningResponse.Reset`

- Signature: <span class="api-signature">`func (x *IsAppRunningResponse) Reset()`</span>

#### <a id="method-isapprunningresponse-string"></a><span class="api-method">Method</span> `IsAppRunningResponse.String`

- Signature: <span class="api-signature">`func (x *IsAppRunningResponse) String() string`</span>

#### <a id="method-listscreensrequest-descriptor"></a><span class="api-method">Method</span> `ListScreensRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ListScreensRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ListScreensRequest.ProtoReflect.Descriptor instead.

#### <a id="method-listscreensrequest-protomessage"></a><span class="api-method">Method</span> `ListScreensRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ListScreensRequest) ProtoMessage()`</span>

#### <a id="method-listscreensrequest-protoreflect"></a><span class="api-method">Method</span> `ListScreensRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ListScreensRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-listscreensrequest-reset"></a><span class="api-method">Method</span> `ListScreensRequest.Reset`

- Signature: <span class="api-signature">`func (x *ListScreensRequest) Reset()`</span>

#### <a id="method-listscreensrequest-string"></a><span class="api-method">Method</span> `ListScreensRequest.String`

- Signature: <span class="api-signature">`func (x *ListScreensRequest) String() string`</span>

#### <a id="method-listscreensresponse-descriptor"></a><span class="api-method">Method</span> `ListScreensResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ListScreensResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ListScreensResponse.ProtoReflect.Descriptor instead.

#### <a id="method-listscreensresponse-getscreens"></a><span class="api-method">Method</span> `ListScreensResponse.GetScreens`

- Signature: <span class="api-signature">`func (x *ListScreensResponse) GetScreens() []*ScreenDescriptor`</span>
- Uses: [`ScreenDescriptor`](#type-screendescriptor)

#### <a id="method-listscreensresponse-protomessage"></a><span class="api-method">Method</span> `ListScreensResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ListScreensResponse) ProtoMessage()`</span>

#### <a id="method-listscreensresponse-protoreflect"></a><span class="api-method">Method</span> `ListScreensResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ListScreensResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-listscreensresponse-reset"></a><span class="api-method">Method</span> `ListScreensResponse.Reset`

- Signature: <span class="api-signature">`func (x *ListScreensResponse) Reset()`</span>

#### <a id="method-listscreensresponse-string"></a><span class="api-method">Method</span> `ListScreensResponse.String`

- Signature: <span class="api-signature">`func (x *ListScreensResponse) String() string`</span>

#### <a id="method-listwindowsresponse-descriptor"></a><span class="api-method">Method</span> `ListWindowsResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ListWindowsResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ListWindowsResponse.ProtoReflect.Descriptor instead.

#### <a id="method-listwindowsresponse-getwindows"></a><span class="api-method">Method</span> `ListWindowsResponse.GetWindows`

- Signature: <span class="api-signature">`func (x *ListWindowsResponse) GetWindows() []*Window`</span>
- Uses: [`Window`](#type-window)

#### <a id="method-listwindowsresponse-protomessage"></a><span class="api-method">Method</span> `ListWindowsResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ListWindowsResponse) ProtoMessage()`</span>

#### <a id="method-listwindowsresponse-protoreflect"></a><span class="api-method">Method</span> `ListWindowsResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ListWindowsResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-listwindowsresponse-reset"></a><span class="api-method">Method</span> `ListWindowsResponse.Reset`

- Signature: <span class="api-signature">`func (x *ListWindowsResponse) Reset()`</span>

#### <a id="method-listwindowsresponse-string"></a><span class="api-method">Method</span> `ListWindowsResponse.String`

- Signature: <span class="api-signature">`func (x *ListWindowsResponse) String() string`</span>

#### <a id="method-match-descriptor"></a><span class="api-method">Method</span> `Match.Descriptor`

- Signature: <span class="api-signature">`func (*Match) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Match.ProtoReflect.Descriptor instead.

#### <a id="method-match-getindex"></a><span class="api-method">Method</span> `Match.GetIndex`

- Signature: <span class="api-signature">`func (x *Match) GetIndex() int32`</span>

#### <a id="method-match-getrect"></a><span class="api-method">Method</span> `Match.GetRect`

- Signature: <span class="api-signature">`func (x *Match) GetRect() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-match-getscore"></a><span class="api-method">Method</span> `Match.GetScore`

- Signature: <span class="api-signature">`func (x *Match) GetScore() float64`</span>

#### <a id="method-match-gettarget"></a><span class="api-method">Method</span> `Match.GetTarget`

- Signature: <span class="api-signature">`func (x *Match) GetTarget() *Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-match-protomessage"></a><span class="api-method">Method</span> `Match.ProtoMessage`

- Signature: <span class="api-signature">`func (*Match) ProtoMessage()`</span>

#### <a id="method-match-protoreflect"></a><span class="api-method">Method</span> `Match.ProtoReflect`

- Signature: <span class="api-signature">`func (x *Match) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-match-reset"></a><span class="api-method">Method</span> `Match.Reset`

- Signature: <span class="api-signature">`func (x *Match) Reset()`</span>

#### <a id="method-match-string"></a><span class="api-method">Method</span> `Match.String`

- Signature: <span class="api-signature">`func (x *Match) String() string`</span>

#### <a id="method-matcherengine-descriptor"></a><span class="api-method">Method</span> `MatcherEngine.Descriptor`

- Signature: <span class="api-signature">`func (MatcherEngine) Descriptor() protoreflect.EnumDescriptor`</span>

#### <a id="method-matcherengine-enum"></a><span class="api-method">Method</span> `MatcherEngine.Enum`

- Signature: <span class="api-signature">`func (x MatcherEngine) Enum() *MatcherEngine`</span>

#### <a id="method-matcherengine-enumdescriptor"></a><span class="api-method">Method</span> `MatcherEngine.EnumDescriptor`

- Signature: <span class="api-signature">`func (MatcherEngine) EnumDescriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use MatcherEngine.Descriptor instead.

#### <a id="method-matcherengine-number"></a><span class="api-method">Method</span> `MatcherEngine.Number`

- Signature: <span class="api-signature">`func (x MatcherEngine) Number() protoreflect.EnumNumber`</span>

#### <a id="method-matcherengine-string"></a><span class="api-method">Method</span> `MatcherEngine.String`

- Signature: <span class="api-signature">`func (x MatcherEngine) String() string`</span>

#### <a id="method-matcherengine-type"></a><span class="api-method">Method</span> `MatcherEngine.Type`

- Signature: <span class="api-signature">`func (MatcherEngine) Type() protoreflect.EnumType`</span>

#### <a id="method-movemouserequest-descriptor"></a><span class="api-method">Method</span> `MoveMouseRequest.Descriptor`

- Signature: <span class="api-signature">`func (*MoveMouseRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use MoveMouseRequest.ProtoReflect.Descriptor instead.

#### <a id="method-movemouserequest-getopts"></a><span class="api-method">Method</span> `MoveMouseRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) GetOpts() *InputOptions`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-movemouserequest-getx"></a><span class="api-method">Method</span> `MoveMouseRequest.GetX`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) GetX() int32`</span>

#### <a id="method-movemouserequest-gety"></a><span class="api-method">Method</span> `MoveMouseRequest.GetY`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) GetY() int32`</span>

#### <a id="method-movemouserequest-protomessage"></a><span class="api-method">Method</span> `MoveMouseRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*MoveMouseRequest) ProtoMessage()`</span>

#### <a id="method-movemouserequest-protoreflect"></a><span class="api-method">Method</span> `MoveMouseRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-movemouserequest-reset"></a><span class="api-method">Method</span> `MoveMouseRequest.Reset`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) Reset()`</span>

#### <a id="method-movemouserequest-string"></a><span class="api-method">Method</span> `MoveMouseRequest.String`

- Signature: <span class="api-signature">`func (x *MoveMouseRequest) String() string`</span>

#### <a id="method-ocrparams-descriptor"></a><span class="api-method">Method</span> `OCRParams.Descriptor`

- Signature: <span class="api-signature">`func (*OCRParams) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use OCRParams.ProtoReflect.Descriptor instead.

#### <a id="method-ocrparams-getcasesensitive"></a><span class="api-method">Method</span> `OCRParams.GetCaseSensitive`

- Signature: <span class="api-signature">`func (x *OCRParams) GetCaseSensitive() bool`</span>

#### <a id="method-ocrparams-getlanguage"></a><span class="api-method">Method</span> `OCRParams.GetLanguage`

- Signature: <span class="api-signature">`func (x *OCRParams) GetLanguage() string`</span>

#### <a id="method-ocrparams-getminconfidence"></a><span class="api-method">Method</span> `OCRParams.GetMinConfidence`

- Signature: <span class="api-signature">`func (x *OCRParams) GetMinConfidence() float64`</span>

#### <a id="method-ocrparams-gettimeoutmillis"></a><span class="api-method">Method</span> `OCRParams.GetTimeoutMillis`

- Signature: <span class="api-signature">`func (x *OCRParams) GetTimeoutMillis() int64`</span>

#### <a id="method-ocrparams-gettrainingdatapath"></a><span class="api-method">Method</span> `OCRParams.GetTrainingDataPath`

- Signature: <span class="api-signature">`func (x *OCRParams) GetTrainingDataPath() string`</span>

#### <a id="method-ocrparams-protomessage"></a><span class="api-method">Method</span> `OCRParams.ProtoMessage`

- Signature: <span class="api-signature">`func (*OCRParams) ProtoMessage()`</span>

#### <a id="method-ocrparams-protoreflect"></a><span class="api-method">Method</span> `OCRParams.ProtoReflect`

- Signature: <span class="api-signature">`func (x *OCRParams) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-ocrparams-reset"></a><span class="api-method">Method</span> `OCRParams.Reset`

- Signature: <span class="api-signature">`func (x *OCRParams) Reset()`</span>

#### <a id="method-ocrparams-string"></a><span class="api-method">Method</span> `OCRParams.String`

- Signature: <span class="api-signature">`func (x *OCRParams) String() string`</span>

#### <a id="method-observechangerequest-descriptor"></a><span class="api-method">Method</span> `ObserveChangeRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ObserveChangeRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ObserveChangeRequest.ProtoReflect.Descriptor instead.

#### <a id="method-observechangerequest-getopts"></a><span class="api-method">Method</span> `ObserveChangeRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) GetOpts() *ObserveOptions`</span>
- Uses: [`ObserveOptions`](#type-observeoptions)

#### <a id="method-observechangerequest-getregion"></a><span class="api-method">Method</span> `ObserveChangeRequest.GetRegion`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) GetRegion() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-observechangerequest-getsource"></a><span class="api-method">Method</span> `ObserveChangeRequest.GetSource`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) GetSource() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-observechangerequest-protomessage"></a><span class="api-method">Method</span> `ObserveChangeRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ObserveChangeRequest) ProtoMessage()`</span>

#### <a id="method-observechangerequest-protoreflect"></a><span class="api-method">Method</span> `ObserveChangeRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-observechangerequest-reset"></a><span class="api-method">Method</span> `ObserveChangeRequest.Reset`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) Reset()`</span>

#### <a id="method-observechangerequest-string"></a><span class="api-method">Method</span> `ObserveChangeRequest.String`

- Signature: <span class="api-signature">`func (x *ObserveChangeRequest) String() string`</span>

#### <a id="method-observeevent-descriptor"></a><span class="api-method">Method</span> `ObserveEvent.Descriptor`

- Signature: <span class="api-signature">`func (*ObserveEvent) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ObserveEvent.ProtoReflect.Descriptor instead.

#### <a id="method-observeevent-getmatch"></a><span class="api-method">Method</span> `ObserveEvent.GetMatch`

- Signature: <span class="api-signature">`func (x *ObserveEvent) GetMatch() *Match`</span>
- Uses: [`Match`](#type-match)

#### <a id="method-observeevent-gettimestampunixmillis"></a><span class="api-method">Method</span> `ObserveEvent.GetTimestampUnixMillis`

- Signature: <span class="api-signature">`func (x *ObserveEvent) GetTimestampUnixMillis() int64`</span>

#### <a id="method-observeevent-gettype"></a><span class="api-method">Method</span> `ObserveEvent.GetType`

- Signature: <span class="api-signature">`func (x *ObserveEvent) GetType() string`</span>

#### <a id="method-observeevent-protomessage"></a><span class="api-method">Method</span> `ObserveEvent.ProtoMessage`

- Signature: <span class="api-signature">`func (*ObserveEvent) ProtoMessage()`</span>

#### <a id="method-observeevent-protoreflect"></a><span class="api-method">Method</span> `ObserveEvent.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ObserveEvent) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-observeevent-reset"></a><span class="api-method">Method</span> `ObserveEvent.Reset`

- Signature: <span class="api-signature">`func (x *ObserveEvent) Reset()`</span>

#### <a id="method-observeevent-string"></a><span class="api-method">Method</span> `ObserveEvent.String`

- Signature: <span class="api-signature">`func (x *ObserveEvent) String() string`</span>

#### <a id="method-observeoptions-descriptor"></a><span class="api-method">Method</span> `ObserveOptions.Descriptor`

- Signature: <span class="api-signature">`func (*ObserveOptions) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ObserveOptions.ProtoReflect.Descriptor instead.

#### <a id="method-observeoptions-getintervalmillis"></a><span class="api-method">Method</span> `ObserveOptions.GetIntervalMillis`

- Signature: <span class="api-signature">`func (x *ObserveOptions) GetIntervalMillis() int64`</span>

#### <a id="method-observeoptions-gettimeoutmillis"></a><span class="api-method">Method</span> `ObserveOptions.GetTimeoutMillis`

- Signature: <span class="api-signature">`func (x *ObserveOptions) GetTimeoutMillis() int64`</span>

#### <a id="method-observeoptions-protomessage"></a><span class="api-method">Method</span> `ObserveOptions.ProtoMessage`

- Signature: <span class="api-signature">`func (*ObserveOptions) ProtoMessage()`</span>

#### <a id="method-observeoptions-protoreflect"></a><span class="api-method">Method</span> `ObserveOptions.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ObserveOptions) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-observeoptions-reset"></a><span class="api-method">Method</span> `ObserveOptions.Reset`

- Signature: <span class="api-signature">`func (x *ObserveOptions) Reset()`</span>

#### <a id="method-observeoptions-string"></a><span class="api-method">Method</span> `ObserveOptions.String`

- Signature: <span class="api-signature">`func (x *ObserveOptions) String() string`</span>

#### <a id="method-observerequest-descriptor"></a><span class="api-method">Method</span> `ObserveRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ObserveRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ObserveRequest.ProtoReflect.Descriptor instead.

#### <a id="method-observerequest-getopts"></a><span class="api-method">Method</span> `ObserveRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ObserveRequest) GetOpts() *ObserveOptions`</span>
- Uses: [`ObserveOptions`](#type-observeoptions)

#### <a id="method-observerequest-getpattern"></a><span class="api-method">Method</span> `ObserveRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *ObserveRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-observerequest-getregion"></a><span class="api-method">Method</span> `ObserveRequest.GetRegion`

- Signature: <span class="api-signature">`func (x *ObserveRequest) GetRegion() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-observerequest-getsource"></a><span class="api-method">Method</span> `ObserveRequest.GetSource`

- Signature: <span class="api-signature">`func (x *ObserveRequest) GetSource() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-observerequest-protomessage"></a><span class="api-method">Method</span> `ObserveRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ObserveRequest) ProtoMessage()`</span>

#### <a id="method-observerequest-protoreflect"></a><span class="api-method">Method</span> `ObserveRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ObserveRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-observerequest-reset"></a><span class="api-method">Method</span> `ObserveRequest.Reset`

- Signature: <span class="api-signature">`func (x *ObserveRequest) Reset()`</span>

#### <a id="method-observerequest-string"></a><span class="api-method">Method</span> `ObserveRequest.String`

- Signature: <span class="api-signature">`func (x *ObserveRequest) String() string`</span>

#### <a id="method-observeresponse-descriptor"></a><span class="api-method">Method</span> `ObserveResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ObserveResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ObserveResponse.ProtoReflect.Descriptor instead.

#### <a id="method-observeresponse-getevents"></a><span class="api-method">Method</span> `ObserveResponse.GetEvents`

- Signature: <span class="api-signature">`func (x *ObserveResponse) GetEvents() []*ObserveEvent`</span>
- Uses: [`ObserveEvent`](#type-observeevent)

#### <a id="method-observeresponse-protomessage"></a><span class="api-method">Method</span> `ObserveResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ObserveResponse) ProtoMessage()`</span>

#### <a id="method-observeresponse-protoreflect"></a><span class="api-method">Method</span> `ObserveResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ObserveResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-observeresponse-reset"></a><span class="api-method">Method</span> `ObserveResponse.Reset`

- Signature: <span class="api-signature">`func (x *ObserveResponse) Reset()`</span>

#### <a id="method-observeresponse-string"></a><span class="api-method">Method</span> `ObserveResponse.String`

- Signature: <span class="api-signature">`func (x *ObserveResponse) String() string`</span>

#### <a id="method-pattern-descriptor"></a><span class="api-method">Method</span> `Pattern.Descriptor`

- Signature: <span class="api-signature">`func (*Pattern) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Pattern.ProtoReflect.Descriptor instead.

#### <a id="method-pattern-getexact"></a><span class="api-method">Method</span> `Pattern.GetExact`

- Signature: <span class="api-signature">`func (x *Pattern) GetExact() bool`</span>

#### <a id="method-pattern-getimage"></a><span class="api-method">Method</span> `Pattern.GetImage`

- Signature: <span class="api-signature">`func (x *Pattern) GetImage() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-pattern-getmask"></a><span class="api-method">Method</span> `Pattern.GetMask`

- Signature: <span class="api-signature">`func (x *Pattern) GetMask() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-pattern-getresizefactor"></a><span class="api-method">Method</span> `Pattern.GetResizeFactor`

- Signature: <span class="api-signature">`func (x *Pattern) GetResizeFactor() float64`</span>

#### <a id="method-pattern-getsimilarity"></a><span class="api-method">Method</span> `Pattern.GetSimilarity`

- Signature: <span class="api-signature">`func (x *Pattern) GetSimilarity() float64`</span>

#### <a id="method-pattern-gettargetoffset"></a><span class="api-method">Method</span> `Pattern.GetTargetOffset`

- Signature: <span class="api-signature">`func (x *Pattern) GetTargetOffset() *Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-pattern-protomessage"></a><span class="api-method">Method</span> `Pattern.ProtoMessage`

- Signature: <span class="api-signature">`func (*Pattern) ProtoMessage()`</span>

#### <a id="method-pattern-protoreflect"></a><span class="api-method">Method</span> `Pattern.ProtoReflect`

- Signature: <span class="api-signature">`func (x *Pattern) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-pattern-reset"></a><span class="api-method">Method</span> `Pattern.Reset`

- Signature: <span class="api-signature">`func (x *Pattern) Reset()`</span>

#### <a id="method-pattern-string"></a><span class="api-method">Method</span> `Pattern.String`

- Signature: <span class="api-signature">`func (x *Pattern) String() string`</span>

#### <a id="method-point-descriptor"></a><span class="api-method">Method</span> `Point.Descriptor`

- Signature: <span class="api-signature">`func (*Point) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Point.ProtoReflect.Descriptor instead.

#### <a id="method-point-getx"></a><span class="api-method">Method</span> `Point.GetX`

- Signature: <span class="api-signature">`func (x *Point) GetX() int32`</span>

#### <a id="method-point-gety"></a><span class="api-method">Method</span> `Point.GetY`

- Signature: <span class="api-signature">`func (x *Point) GetY() int32`</span>

#### <a id="method-point-protomessage"></a><span class="api-method">Method</span> `Point.ProtoMessage`

- Signature: <span class="api-signature">`func (*Point) ProtoMessage()`</span>

#### <a id="method-point-protoreflect"></a><span class="api-method">Method</span> `Point.ProtoReflect`

- Signature: <span class="api-signature">`func (x *Point) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-point-reset"></a><span class="api-method">Method</span> `Point.Reset`

- Signature: <span class="api-signature">`func (x *Point) Reset()`</span>

#### <a id="method-point-string"></a><span class="api-method">Method</span> `Point.String`

- Signature: <span class="api-signature">`func (x *Point) String() string`</span>

#### <a id="method-readtextrequest-descriptor"></a><span class="api-method">Method</span> `ReadTextRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ReadTextRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ReadTextRequest.ProtoReflect.Descriptor instead.

#### <a id="method-readtextrequest-getparams"></a><span class="api-method">Method</span> `ReadTextRequest.GetParams`

- Signature: <span class="api-signature">`func (x *ReadTextRequest) GetParams() *OCRParams`</span>
- Uses: [`OCRParams`](#type-ocrparams)

#### <a id="method-readtextrequest-getsource"></a><span class="api-method">Method</span> `ReadTextRequest.GetSource`

- Signature: <span class="api-signature">`func (x *ReadTextRequest) GetSource() *GrayImage`</span>
- Uses: [`GrayImage`](#type-grayimage)

#### <a id="method-readtextrequest-protomessage"></a><span class="api-method">Method</span> `ReadTextRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ReadTextRequest) ProtoMessage()`</span>

#### <a id="method-readtextrequest-protoreflect"></a><span class="api-method">Method</span> `ReadTextRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ReadTextRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-readtextrequest-reset"></a><span class="api-method">Method</span> `ReadTextRequest.Reset`

- Signature: <span class="api-signature">`func (x *ReadTextRequest) Reset()`</span>

#### <a id="method-readtextrequest-string"></a><span class="api-method">Method</span> `ReadTextRequest.String`

- Signature: <span class="api-signature">`func (x *ReadTextRequest) String() string`</span>

#### <a id="method-readtextresponse-descriptor"></a><span class="api-method">Method</span> `ReadTextResponse.Descriptor`

- Signature: <span class="api-signature">`func (*ReadTextResponse) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ReadTextResponse.ProtoReflect.Descriptor instead.

#### <a id="method-readtextresponse-gettext"></a><span class="api-method">Method</span> `ReadTextResponse.GetText`

- Signature: <span class="api-signature">`func (x *ReadTextResponse) GetText() string`</span>

#### <a id="method-readtextresponse-protomessage"></a><span class="api-method">Method</span> `ReadTextResponse.ProtoMessage`

- Signature: <span class="api-signature">`func (*ReadTextResponse) ProtoMessage()`</span>

#### <a id="method-readtextresponse-protoreflect"></a><span class="api-method">Method</span> `ReadTextResponse.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ReadTextResponse) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-readtextresponse-reset"></a><span class="api-method">Method</span> `ReadTextResponse.Reset`

- Signature: <span class="api-signature">`func (x *ReadTextResponse) Reset()`</span>

#### <a id="method-readtextresponse-string"></a><span class="api-method">Method</span> `ReadTextResponse.String`

- Signature: <span class="api-signature">`func (x *ReadTextResponse) String() string`</span>

#### <a id="method-rect-descriptor"></a><span class="api-method">Method</span> `Rect.Descriptor`

- Signature: <span class="api-signature">`func (*Rect) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Rect.ProtoReflect.Descriptor instead.

#### <a id="method-rect-geth"></a><span class="api-method">Method</span> `Rect.GetH`

- Signature: <span class="api-signature">`func (x *Rect) GetH() int32`</span>

#### <a id="method-rect-getw"></a><span class="api-method">Method</span> `Rect.GetW`

- Signature: <span class="api-signature">`func (x *Rect) GetW() int32`</span>

#### <a id="method-rect-getx"></a><span class="api-method">Method</span> `Rect.GetX`

- Signature: <span class="api-signature">`func (x *Rect) GetX() int32`</span>

#### <a id="method-rect-gety"></a><span class="api-method">Method</span> `Rect.GetY`

- Signature: <span class="api-signature">`func (x *Rect) GetY() int32`</span>

#### <a id="method-rect-protomessage"></a><span class="api-method">Method</span> `Rect.ProtoMessage`

- Signature: <span class="api-signature">`func (*Rect) ProtoMessage()`</span>

#### <a id="method-rect-protoreflect"></a><span class="api-method">Method</span> `Rect.ProtoReflect`

- Signature: <span class="api-signature">`func (x *Rect) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-rect-reset"></a><span class="api-method">Method</span> `Rect.Reset`

- Signature: <span class="api-signature">`func (x *Rect) Reset()`</span>

#### <a id="method-rect-string"></a><span class="api-method">Method</span> `Rect.String`

- Signature: <span class="api-signature">`func (x *Rect) String() string`</span>

#### <a id="method-screendescriptor-descriptor"></a><span class="api-method">Method</span> `ScreenDescriptor.Descriptor`

- Signature: <span class="api-signature">`func (*ScreenDescriptor) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ScreenDescriptor.ProtoReflect.Descriptor instead.

#### <a id="method-screendescriptor-getbounds"></a><span class="api-method">Method</span> `ScreenDescriptor.GetBounds`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) GetBounds() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-screendescriptor-getid"></a><span class="api-method">Method</span> `ScreenDescriptor.GetId`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) GetId() int32`</span>

#### <a id="method-screendescriptor-getname"></a><span class="api-method">Method</span> `ScreenDescriptor.GetName`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) GetName() string`</span>

#### <a id="method-screendescriptor-getprimary"></a><span class="api-method">Method</span> `ScreenDescriptor.GetPrimary`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) GetPrimary() bool`</span>

#### <a id="method-screendescriptor-protomessage"></a><span class="api-method">Method</span> `ScreenDescriptor.ProtoMessage`

- Signature: <span class="api-signature">`func (*ScreenDescriptor) ProtoMessage()`</span>

#### <a id="method-screendescriptor-protoreflect"></a><span class="api-method">Method</span> `ScreenDescriptor.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-screendescriptor-reset"></a><span class="api-method">Method</span> `ScreenDescriptor.Reset`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) Reset()`</span>

#### <a id="method-screendescriptor-string"></a><span class="api-method">Method</span> `ScreenDescriptor.String`

- Signature: <span class="api-signature">`func (x *ScreenDescriptor) String() string`</span>

#### <a id="method-screenqueryoptions-descriptor"></a><span class="api-method">Method</span> `ScreenQueryOptions.Descriptor`

- Signature: <span class="api-signature">`func (*ScreenQueryOptions) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ScreenQueryOptions.ProtoReflect.Descriptor instead.

#### <a id="method-screenqueryoptions-getintervalmillis"></a><span class="api-method">Method</span> `ScreenQueryOptions.GetIntervalMillis`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) GetIntervalMillis() int64`</span>

#### <a id="method-screenqueryoptions-getmatcherengine"></a><span class="api-method">Method</span> `ScreenQueryOptions.GetMatcherEngine`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) GetMatcherEngine() MatcherEngine`</span>
- Uses: [`MatcherEngine`](#type-matcherengine)

#### <a id="method-screenqueryoptions-getregion"></a><span class="api-method">Method</span> `ScreenQueryOptions.GetRegion`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) GetRegion() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-screenqueryoptions-getscreenid"></a><span class="api-method">Method</span> `ScreenQueryOptions.GetScreenId`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) GetScreenId() int32`</span>

#### <a id="method-screenqueryoptions-gettimeoutmillis"></a><span class="api-method">Method</span> `ScreenQueryOptions.GetTimeoutMillis`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) GetTimeoutMillis() int64`</span>

#### <a id="method-screenqueryoptions-protomessage"></a><span class="api-method">Method</span> `ScreenQueryOptions.ProtoMessage`

- Signature: <span class="api-signature">`func (*ScreenQueryOptions) ProtoMessage()`</span>

#### <a id="method-screenqueryoptions-protoreflect"></a><span class="api-method">Method</span> `ScreenQueryOptions.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-screenqueryoptions-reset"></a><span class="api-method">Method</span> `ScreenQueryOptions.Reset`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) Reset()`</span>

#### <a id="method-screenqueryoptions-string"></a><span class="api-method">Method</span> `ScreenQueryOptions.String`

- Signature: <span class="api-signature">`func (x *ScreenQueryOptions) String() string`</span>

#### <a id="method-scrollwheelrequest-descriptor"></a><span class="api-method">Method</span> `ScrollWheelRequest.Descriptor`

- Signature: <span class="api-signature">`func (*ScrollWheelRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use ScrollWheelRequest.ProtoReflect.Descriptor instead.

#### <a id="method-scrollwheelrequest-getdirection"></a><span class="api-method">Method</span> `ScrollWheelRequest.GetDirection`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) GetDirection() string`</span>

#### <a id="method-scrollwheelrequest-getopts"></a><span class="api-method">Method</span> `ScrollWheelRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) GetOpts() *InputOptions`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-scrollwheelrequest-getsteps"></a><span class="api-method">Method</span> `ScrollWheelRequest.GetSteps`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) GetSteps() int32`</span>

#### <a id="method-scrollwheelrequest-getx"></a><span class="api-method">Method</span> `ScrollWheelRequest.GetX`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) GetX() int32`</span>

#### <a id="method-scrollwheelrequest-gety"></a><span class="api-method">Method</span> `ScrollWheelRequest.GetY`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) GetY() int32`</span>

#### <a id="method-scrollwheelrequest-protomessage"></a><span class="api-method">Method</span> `ScrollWheelRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*ScrollWheelRequest) ProtoMessage()`</span>

#### <a id="method-scrollwheelrequest-protoreflect"></a><span class="api-method">Method</span> `ScrollWheelRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-scrollwheelrequest-reset"></a><span class="api-method">Method</span> `ScrollWheelRequest.Reset`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) Reset()`</span>

#### <a id="method-scrollwheelrequest-string"></a><span class="api-method">Method</span> `ScrollWheelRequest.String`

- Signature: <span class="api-signature">`func (x *ScrollWheelRequest) String() string`</span>

#### <a id="method-textmatch-descriptor"></a><span class="api-method">Method</span> `TextMatch.Descriptor`

- Signature: <span class="api-signature">`func (*TextMatch) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use TextMatch.ProtoReflect.Descriptor instead.

#### <a id="method-textmatch-getconfidence"></a><span class="api-method">Method</span> `TextMatch.GetConfidence`

- Signature: <span class="api-signature">`func (x *TextMatch) GetConfidence() float64`</span>

#### <a id="method-textmatch-getindex"></a><span class="api-method">Method</span> `TextMatch.GetIndex`

- Signature: <span class="api-signature">`func (x *TextMatch) GetIndex() int32`</span>

#### <a id="method-textmatch-getrect"></a><span class="api-method">Method</span> `TextMatch.GetRect`

- Signature: <span class="api-signature">`func (x *TextMatch) GetRect() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-textmatch-gettext"></a><span class="api-method">Method</span> `TextMatch.GetText`

- Signature: <span class="api-signature">`func (x *TextMatch) GetText() string`</span>

#### <a id="method-textmatch-protomessage"></a><span class="api-method">Method</span> `TextMatch.ProtoMessage`

- Signature: <span class="api-signature">`func (*TextMatch) ProtoMessage()`</span>

#### <a id="method-textmatch-protoreflect"></a><span class="api-method">Method</span> `TextMatch.ProtoReflect`

- Signature: <span class="api-signature">`func (x *TextMatch) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-textmatch-reset"></a><span class="api-method">Method</span> `TextMatch.Reset`

- Signature: <span class="api-signature">`func (x *TextMatch) Reset()`</span>

#### <a id="method-textmatch-string"></a><span class="api-method">Method</span> `TextMatch.String`

- Signature: <span class="api-signature">`func (x *TextMatch) String() string`</span>

#### <a id="method-typetextrequest-descriptor"></a><span class="api-method">Method</span> `TypeTextRequest.Descriptor`

- Signature: <span class="api-signature">`func (*TypeTextRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use TypeTextRequest.ProtoReflect.Descriptor instead.

#### <a id="method-typetextrequest-getopts"></a><span class="api-method">Method</span> `TypeTextRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *TypeTextRequest) GetOpts() *InputOptions`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-typetextrequest-gettext"></a><span class="api-method">Method</span> `TypeTextRequest.GetText`

- Signature: <span class="api-signature">`func (x *TypeTextRequest) GetText() string`</span>

#### <a id="method-typetextrequest-protomessage"></a><span class="api-method">Method</span> `TypeTextRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*TypeTextRequest) ProtoMessage()`</span>

#### <a id="method-typetextrequest-protoreflect"></a><span class="api-method">Method</span> `TypeTextRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *TypeTextRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-typetextrequest-reset"></a><span class="api-method">Method</span> `TypeTextRequest.Reset`

- Signature: <span class="api-signature">`func (x *TypeTextRequest) Reset()`</span>

#### <a id="method-typetextrequest-string"></a><span class="api-method">Method</span> `TypeTextRequest.String`

- Signature: <span class="api-signature">`func (x *TypeTextRequest) String() string`</span>

#### <a id="method-unimplementedsikuliserviceserver-capturescreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.CaptureScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) CaptureScreen(context.Context, *CaptureScreenRequest) (*CaptureScreenResponse, error)`</span>
- Uses: [`CaptureScreenRequest`](#type-capturescreenrequest), [`CaptureScreenResponse`](#type-capturescreenresponse)

#### <a id="method-unimplementedsikuliserviceserver-click"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Click`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Click(context.Context, *ClickRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`ClickRequest`](#type-clickrequest)

#### <a id="method-unimplementedsikuliserviceserver-clickonscreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ClickOnScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ClickOnScreen(context.Context, *ClickOnScreenRequest) (*FindResponse, error)`</span>
- Uses: [`ClickOnScreenRequest`](#type-clickonscreenrequest), [`FindResponse`](#type-findresponse)

#### <a id="method-unimplementedsikuliserviceserver-closeapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.CloseApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`AppActionRequest`](#type-appactionrequest)

#### <a id="method-unimplementedsikuliserviceserver-existsonscreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ExistsOnScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ExistsOnScreen(context.Context, *ExistsOnScreenRequest) (*ExistsOnScreenResponse, error)`</span>
- Uses: [`ExistsOnScreenRequest`](#type-existsonscreenrequest), [`ExistsOnScreenResponse`](#type-existsonscreenresponse)

#### <a id="method-unimplementedsikuliserviceserver-find"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Find`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Find(context.Context, *FindRequest) (*FindResponse, error)`</span>
- Uses: [`FindRequest`](#type-findrequest), [`FindResponse`](#type-findresponse)

#### <a id="method-unimplementedsikuliserviceserver-findall"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindAll`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindAll(context.Context, *FindRequest) (*FindAllResponse, error)`</span>
- Uses: [`FindAllResponse`](#type-findallresponse), [`FindRequest`](#type-findrequest)

#### <a id="method-unimplementedsikuliserviceserver-findonscreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindOnScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindOnScreen(context.Context, *FindOnScreenRequest) (*FindResponse, error)`</span>
- Uses: [`FindOnScreenRequest`](#type-findonscreenrequest), [`FindResponse`](#type-findresponse)

#### <a id="method-unimplementedsikuliserviceserver-findtext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)`</span>
- Uses: [`FindTextRequest`](#type-findtextrequest), [`FindTextResponse`](#type-findtextresponse)

#### <a id="method-unimplementedsikuliserviceserver-findwindows"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindWindows`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindWindows(context.Context, *WindowQueryRequest) (*ListWindowsResponse, error)`</span>
- Uses: [`ListWindowsResponse`](#type-listwindowsresponse), [`WindowQueryRequest`](#type-windowqueryrequest)

#### <a id="method-unimplementedsikuliserviceserver-focusapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FocusApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`AppActionRequest`](#type-appactionrequest)

#### <a id="method-unimplementedsikuliserviceserver-getfocusedwindow"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.GetFocusedWindow`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) GetFocusedWindow(context.Context, *AppActionRequest) (*GetWindowResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`GetWindowResponse`](#type-getwindowresponse)

#### <a id="method-unimplementedsikuliserviceserver-getprimaryscreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.GetPrimaryScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) GetPrimaryScreen(context.Context, *GetPrimaryScreenRequest) (*GetPrimaryScreenResponse, error)`</span>
- Uses: [`GetPrimaryScreenRequest`](#type-getprimaryscreenrequest), [`GetPrimaryScreenResponse`](#type-getprimaryscreenresponse)

#### <a id="method-unimplementedsikuliserviceserver-getwindow"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.GetWindow`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) GetWindow(context.Context, *WindowQueryRequest) (*GetWindowResponse, error)`</span>
- Uses: [`GetWindowResponse`](#type-getwindowresponse), [`WindowQueryRequest`](#type-windowqueryrequest)

#### <a id="method-unimplementedsikuliserviceserver-hotkey"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Hotkey`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`HotkeyRequest`](#type-hotkeyrequest)

#### <a id="method-unimplementedsikuliserviceserver-isapprunning"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.IsAppRunning`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`IsAppRunningResponse`](#type-isapprunningresponse)

#### <a id="method-unimplementedsikuliserviceserver-keydown"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.KeyDown`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) KeyDown(context.Context, *HotkeyRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`HotkeyRequest`](#type-hotkeyrequest)

#### <a id="method-unimplementedsikuliserviceserver-keyup"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.KeyUp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) KeyUp(context.Context, *HotkeyRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`HotkeyRequest`](#type-hotkeyrequest)

#### <a id="method-unimplementedsikuliserviceserver-listscreens"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ListScreens`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ListScreens(context.Context, *ListScreensRequest) (*ListScreensResponse, error)`</span>
- Uses: [`ListScreensRequest`](#type-listscreensrequest), [`ListScreensResponse`](#type-listscreensresponse)

#### <a id="method-unimplementedsikuliserviceserver-listwindows"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ListWindows`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`ListWindowsResponse`](#type-listwindowsresponse)

#### <a id="method-unimplementedsikuliserviceserver-mousedown"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.MouseDown`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) MouseDown(context.Context, *ClickRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`ClickRequest`](#type-clickrequest)

#### <a id="method-unimplementedsikuliserviceserver-mouseup"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.MouseUp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) MouseUp(context.Context, *ClickRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`ClickRequest`](#type-clickrequest)

#### <a id="method-unimplementedsikuliserviceserver-movemouse"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.MoveMouse`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`MoveMouseRequest`](#type-movemouserequest)

#### <a id="method-unimplementedsikuliserviceserver-observeappear"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveAppear`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveRequest`](#type-observerequest), [`ObserveResponse`](#type-observeresponse)

#### <a id="method-unimplementedsikuliserviceserver-observechange"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveChange`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveChangeRequest`](#type-observechangerequest), [`ObserveResponse`](#type-observeresponse)

#### <a id="method-unimplementedsikuliserviceserver-observevanish"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveVanish`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveRequest`](#type-observerequest), [`ObserveResponse`](#type-observeresponse)

#### <a id="method-unimplementedsikuliserviceserver-openapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.OpenApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`AppActionRequest`](#type-appactionrequest)

#### <a id="method-unimplementedsikuliserviceserver-pastetext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.PasteText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) PasteText(context.Context, *TypeTextRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`TypeTextRequest`](#type-typetextrequest)

#### <a id="method-unimplementedsikuliserviceserver-readtext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ReadText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)`</span>
- Uses: [`ReadTextRequest`](#type-readtextrequest), [`ReadTextResponse`](#type-readtextresponse)

#### <a id="method-unimplementedsikuliserviceserver-scrollwheel"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ScrollWheel`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ScrollWheel(context.Context, *ScrollWheelRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`ScrollWheelRequest`](#type-scrollwheelrequest)

#### <a id="method-unimplementedsikuliserviceserver-typetext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.TypeText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)`</span>
- Uses: [`ActionResponse`](#type-actionresponse), [`TypeTextRequest`](#type-typetextrequest)

#### <a id="method-unimplementedsikuliserviceserver-waitonscreen"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.WaitOnScreen`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) WaitOnScreen(context.Context, *WaitOnScreenRequest) (*FindResponse, error)`</span>
- Uses: [`FindResponse`](#type-findresponse), [`WaitOnScreenRequest`](#type-waitonscreenrequest)

#### <a id="method-waitonscreenrequest-descriptor"></a><span class="api-method">Method</span> `WaitOnScreenRequest.Descriptor`

- Signature: <span class="api-signature">`func (*WaitOnScreenRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use WaitOnScreenRequest.ProtoReflect.Descriptor instead.

#### <a id="method-waitonscreenrequest-getopts"></a><span class="api-method">Method</span> `WaitOnScreenRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *WaitOnScreenRequest) GetOpts() *ScreenQueryOptions`</span>
- Uses: [`ScreenQueryOptions`](#type-screenqueryoptions)

#### <a id="method-waitonscreenrequest-getpattern"></a><span class="api-method">Method</span> `WaitOnScreenRequest.GetPattern`

- Signature: <span class="api-signature">`func (x *WaitOnScreenRequest) GetPattern() *Pattern`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-waitonscreenrequest-protomessage"></a><span class="api-method">Method</span> `WaitOnScreenRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*WaitOnScreenRequest) ProtoMessage()`</span>

#### <a id="method-waitonscreenrequest-protoreflect"></a><span class="api-method">Method</span> `WaitOnScreenRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *WaitOnScreenRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-waitonscreenrequest-reset"></a><span class="api-method">Method</span> `WaitOnScreenRequest.Reset`

- Signature: <span class="api-signature">`func (x *WaitOnScreenRequest) Reset()`</span>

#### <a id="method-waitonscreenrequest-string"></a><span class="api-method">Method</span> `WaitOnScreenRequest.String`

- Signature: <span class="api-signature">`func (x *WaitOnScreenRequest) String() string`</span>

#### <a id="method-window-descriptor"></a><span class="api-method">Method</span> `Window.Descriptor`

- Signature: <span class="api-signature">`func (*Window) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Window.ProtoReflect.Descriptor instead.

#### <a id="method-window-getapp"></a><span class="api-method">Method</span> `Window.GetApp`

- Signature: <span class="api-signature">`func (x *Window) GetApp() string`</span>

#### <a id="method-window-getbounds"></a><span class="api-method">Method</span> `Window.GetBounds`

- Signature: <span class="api-signature">`func (x *Window) GetBounds() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-window-getfocused"></a><span class="api-method">Method</span> `Window.GetFocused`

- Signature: <span class="api-signature">`func (x *Window) GetFocused() bool`</span>

#### <a id="method-window-getid"></a><span class="api-method">Method</span> `Window.GetId`

- Signature: <span class="api-signature">`func (x *Window) GetId() string`</span>

#### <a id="method-window-getpid"></a><span class="api-method">Method</span> `Window.GetPid`

- Signature: <span class="api-signature">`func (x *Window) GetPid() int32`</span>

#### <a id="method-window-gettitle"></a><span class="api-method">Method</span> `Window.GetTitle`

- Signature: <span class="api-signature">`func (x *Window) GetTitle() string`</span>

#### <a id="method-window-protomessage"></a><span class="api-method">Method</span> `Window.ProtoMessage`

- Signature: <span class="api-signature">`func (*Window) ProtoMessage()`</span>

#### <a id="method-window-protoreflect"></a><span class="api-method">Method</span> `Window.ProtoReflect`

- Signature: <span class="api-signature">`func (x *Window) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-window-reset"></a><span class="api-method">Method</span> `Window.Reset`

- Signature: <span class="api-signature">`func (x *Window) Reset()`</span>

#### <a id="method-window-string"></a><span class="api-method">Method</span> `Window.String`

- Signature: <span class="api-signature">`func (x *Window) String() string`</span>

#### <a id="method-windowquery-descriptor"></a><span class="api-method">Method</span> `WindowQuery.Descriptor`

- Signature: <span class="api-signature">`func (*WindowQuery) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use WindowQuery.ProtoReflect.Descriptor instead.

#### <a id="method-windowquery-getfocusedonly"></a><span class="api-method">Method</span> `WindowQuery.GetFocusedOnly`

- Signature: <span class="api-signature">`func (x *WindowQuery) GetFocusedOnly() bool`</span>

#### <a id="method-windowquery-getid"></a><span class="api-method">Method</span> `WindowQuery.GetId`

- Signature: <span class="api-signature">`func (x *WindowQuery) GetId() string`</span>

#### <a id="method-windowquery-getindex"></a><span class="api-method">Method</span> `WindowQuery.GetIndex`

- Signature: <span class="api-signature">`func (x *WindowQuery) GetIndex() int32`</span>

#### <a id="method-windowquery-gettitlecontains"></a><span class="api-method">Method</span> `WindowQuery.GetTitleContains`

- Signature: <span class="api-signature">`func (x *WindowQuery) GetTitleContains() string`</span>

#### <a id="method-windowquery-gettitleexact"></a><span class="api-method">Method</span> `WindowQuery.GetTitleExact`

- Signature: <span class="api-signature">`func (x *WindowQuery) GetTitleExact() string`</span>

#### <a id="method-windowquery-protomessage"></a><span class="api-method">Method</span> `WindowQuery.ProtoMessage`

- Signature: <span class="api-signature">`func (*WindowQuery) ProtoMessage()`</span>

#### <a id="method-windowquery-protoreflect"></a><span class="api-method">Method</span> `WindowQuery.ProtoReflect`

- Signature: <span class="api-signature">`func (x *WindowQuery) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-windowquery-reset"></a><span class="api-method">Method</span> `WindowQuery.Reset`

- Signature: <span class="api-signature">`func (x *WindowQuery) Reset()`</span>

#### <a id="method-windowquery-string"></a><span class="api-method">Method</span> `WindowQuery.String`

- Signature: <span class="api-signature">`func (x *WindowQuery) String() string`</span>

#### <a id="method-windowqueryrequest-descriptor"></a><span class="api-method">Method</span> `WindowQueryRequest.Descriptor`

- Signature: <span class="api-signature">`func (*WindowQueryRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use WindowQueryRequest.ProtoReflect.Descriptor instead.

#### <a id="method-windowqueryrequest-getname"></a><span class="api-method">Method</span> `WindowQueryRequest.GetName`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) GetName() string`</span>

#### <a id="method-windowqueryrequest-getopts"></a><span class="api-method">Method</span> `WindowQueryRequest.GetOpts`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) GetOpts() *AppOptions`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-windowqueryrequest-getquery"></a><span class="api-method">Method</span> `WindowQueryRequest.GetQuery`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) GetQuery() *WindowQuery`</span>
- Uses: [`WindowQuery`](#type-windowquery)

#### <a id="method-windowqueryrequest-protomessage"></a><span class="api-method">Method</span> `WindowQueryRequest.ProtoMessage`

- Signature: <span class="api-signature">`func (*WindowQueryRequest) ProtoMessage()`</span>

#### <a id="method-windowqueryrequest-protoreflect"></a><span class="api-method">Method</span> `WindowQueryRequest.ProtoReflect`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) ProtoReflect() protoreflect.Message`</span>

#### <a id="method-windowqueryrequest-reset"></a><span class="api-method">Method</span> `WindowQueryRequest.Reset`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) Reset()`</span>

#### <a id="method-windowqueryrequest-string"></a><span class="api-method">Method</span> `WindowQueryRequest.String`

- Signature: <span class="api-signature">`func (x *WindowQueryRequest) String() string`</span>

## Raw Package Doc

```text
package sikuliv1 // import "github.com/smysnk/sikuligo/internal/grpcv1/pb"


CONSTANTS

const (
	SikuliService_ListScreens_FullMethodName      = "/sikuli.v1.SikuliService/ListScreens"
	SikuliService_GetPrimaryScreen_FullMethodName = "/sikuli.v1.SikuliService/GetPrimaryScreen"
	SikuliService_CaptureScreen_FullMethodName    = "/sikuli.v1.SikuliService/CaptureScreen"
	SikuliService_Find_FullMethodName             = "/sikuli.v1.SikuliService/Find"
	SikuliService_FindAll_FullMethodName          = "/sikuli.v1.SikuliService/FindAll"
	SikuliService_FindOnScreen_FullMethodName     = "/sikuli.v1.SikuliService/FindOnScreen"
	SikuliService_ExistsOnScreen_FullMethodName   = "/sikuli.v1.SikuliService/ExistsOnScreen"
	SikuliService_WaitOnScreen_FullMethodName     = "/sikuli.v1.SikuliService/WaitOnScreen"
	SikuliService_ClickOnScreen_FullMethodName    = "/sikuli.v1.SikuliService/ClickOnScreen"
	SikuliService_ReadText_FullMethodName         = "/sikuli.v1.SikuliService/ReadText"
	SikuliService_FindText_FullMethodName         = "/sikuli.v1.SikuliService/FindText"
	SikuliService_MoveMouse_FullMethodName        = "/sikuli.v1.SikuliService/MoveMouse"
	SikuliService_Click_FullMethodName            = "/sikuli.v1.SikuliService/Click"
	SikuliService_TypeText_FullMethodName         = "/sikuli.v1.SikuliService/TypeText"
	SikuliService_PasteText_FullMethodName        = "/sikuli.v1.SikuliService/PasteText"
	SikuliService_Hotkey_FullMethodName           = "/sikuli.v1.SikuliService/Hotkey"
	SikuliService_MouseDown_FullMethodName        = "/sikuli.v1.SikuliService/MouseDown"
	SikuliService_MouseUp_FullMethodName          = "/sikuli.v1.SikuliService/MouseUp"
	SikuliService_KeyDown_FullMethodName          = "/sikuli.v1.SikuliService/KeyDown"
	SikuliService_KeyUp_FullMethodName            = "/sikuli.v1.SikuliService/KeyUp"
	SikuliService_ScrollWheel_FullMethodName      = "/sikuli.v1.SikuliService/ScrollWheel"
	SikuliService_ObserveAppear_FullMethodName    = "/sikuli.v1.SikuliService/ObserveAppear"
	SikuliService_ObserveVanish_FullMethodName    = "/sikuli.v1.SikuliService/ObserveVanish"
	SikuliService_ObserveChange_FullMethodName    = "/sikuli.v1.SikuliService/ObserveChange"
	SikuliService_OpenApp_FullMethodName          = "/sikuli.v1.SikuliService/OpenApp"
	SikuliService_FocusApp_FullMethodName         = "/sikuli.v1.SikuliService/FocusApp"
	SikuliService_CloseApp_FullMethodName         = "/sikuli.v1.SikuliService/CloseApp"
	SikuliService_IsAppRunning_FullMethodName     = "/sikuli.v1.SikuliService/IsAppRunning"
	SikuliService_ListWindows_FullMethodName      = "/sikuli.v1.SikuliService/ListWindows"
	SikuliService_FindWindows_FullMethodName      = "/sikuli.v1.SikuliService/FindWindows"
	SikuliService_GetWindow_FullMethodName        = "/sikuli.v1.SikuliService/GetWindow"
	SikuliService_GetFocusedWindow_FullMethodName = "/sikuli.v1.SikuliService/GetFocusedWindow"
)

VARIABLES

var (
	MatcherEngine_name = map[int32]string{
		0: "MATCHER_ENGINE_UNSPECIFIED",
		1: "MATCHER_ENGINE_TEMPLATE",
		2: "MATCHER_ENGINE_ORB",
		3: "MATCHER_ENGINE_HYBRID",
		4: "MATCHER_ENGINE_AKAZE",
		5: "MATCHER_ENGINE_BRISK",
		6: "MATCHER_ENGINE_KAZE",
		7: "MATCHER_ENGINE_SIFT",
	}
	MatcherEngine_value = map[string]int32{
		"MATCHER_ENGINE_UNSPECIFIED": 0,
		"MATCHER_ENGINE_TEMPLATE":    1,
		"MATCHER_ENGINE_ORB":         2,
		"MATCHER_ENGINE_HYBRID":      3,
		"MATCHER_ENGINE_AKAZE":       4,
		"MATCHER_ENGINE_BRISK":       5,
		"MATCHER_ENGINE_KAZE":        6,
		"MATCHER_ENGINE_SIFT":        7,
	}
)
    Enum value maps for MatcherEngine.

var File_sikuli_v1_sikuli_proto protoreflect.FileDescriptor
var SikuliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sikuli.v1.SikuliService",
	HandlerType: (*SikuliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListScreens",
			Handler:    _SikuliService_ListScreens_Handler,
		},
		{
			MethodName: "GetPrimaryScreen",
			Handler:    _SikuliService_GetPrimaryScreen_Handler,
		},
		{
			MethodName: "CaptureScreen",
			Handler:    _SikuliService_CaptureScreen_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _SikuliService_Find_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _SikuliService_FindAll_Handler,
		},
		{
			MethodName: "FindOnScreen",
			Handler:    _SikuliService_FindOnScreen_Handler,
		},
		{
			MethodName: "ExistsOnScreen",
			Handler:    _SikuliService_ExistsOnScreen_Handler,
		},
		{
			MethodName: "WaitOnScreen",
			Handler:    _SikuliService_WaitOnScreen_Handler,
		},
		{
			MethodName: "ClickOnScreen",
			Handler:    _SikuliService_ClickOnScreen_Handler,
		},
		{
			MethodName: "ReadText",
			Handler:    _SikuliService_ReadText_Handler,
		},
		{
			MethodName: "FindText",
			Handler:    _SikuliService_FindText_Handler,
		},
		{
			MethodName: "MoveMouse",
			Handler:    _SikuliService_MoveMouse_Handler,
		},
		{
			MethodName: "Click",
			Handler:    _SikuliService_Click_Handler,
		},
		{
			MethodName: "TypeText",
			Handler:    _SikuliService_TypeText_Handler,
		},
		{
			MethodName: "PasteText",
			Handler:    _SikuliService_PasteText_Handler,
		},
		{
			MethodName: "Hotkey",
			Handler:    _SikuliService_Hotkey_Handler,
		},
		{
			MethodName: "MouseDown",
			Handler:    _SikuliService_MouseDown_Handler,
		},
		{
			MethodName: "MouseUp",
			Handler:    _SikuliService_MouseUp_Handler,
		},
		{
			MethodName: "KeyDown",
			Handler:    _SikuliService_KeyDown_Handler,
		},
		{
			MethodName: "KeyUp",
			Handler:    _SikuliService_KeyUp_Handler,
		},
		{
			MethodName: "ScrollWheel",
			Handler:    _SikuliService_ScrollWheel_Handler,
		},
		{
			MethodName: "ObserveAppear",
			Handler:    _SikuliService_ObserveAppear_Handler,
		},
		{
			MethodName: "ObserveVanish",
			Handler:    _SikuliService_ObserveVanish_Handler,
		},
		{
			MethodName: "ObserveChange",
			Handler:    _SikuliService_ObserveChange_Handler,
		},
		{
			MethodName: "OpenApp",
			Handler:    _SikuliService_OpenApp_Handler,
		},
		{
			MethodName: "FocusApp",
			Handler:    _SikuliService_FocusApp_Handler,
		},
		{
			MethodName: "CloseApp",
			Handler:    _SikuliService_CloseApp_Handler,
		},
		{
			MethodName: "IsAppRunning",
			Handler:    _SikuliService_IsAppRunning_Handler,
		},
		{
			MethodName: "ListWindows",
			Handler:    _SikuliService_ListWindows_Handler,
		},
		{
			MethodName: "FindWindows",
			Handler:    _SikuliService_FindWindows_Handler,
		},
		{
			MethodName: "GetWindow",
			Handler:    _SikuliService_GetWindow_Handler,
		},
		{
			MethodName: "GetFocusedWindow",
			Handler:    _SikuliService_GetFocusedWindow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sikuli/v1/sikuli.proto",
}
    SikuliService_ServiceDesc is the grpc.ServiceDesc for SikuliService service.
    It's only intended for direct use with grpc.RegisterService, and not to be
    introspected or modified (even as a copy)


FUNCTIONS

func RegisterSikuliServiceServer(s grpc.ServiceRegistrar, srv SikuliServiceServer)

TYPES

type ActionResponse struct {
	// Has unexported fields.
}

func (*ActionResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.

func (*ActionResponse) ProtoMessage()

func (x *ActionResponse) ProtoReflect() protoreflect.Message

func (x *ActionResponse) Reset()

func (x *ActionResponse) String() string

type AppActionRequest struct {
	Name string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []string    `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Opts *AppOptions `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*AppActionRequest) Descriptor() ([]byte, []int)
    Deprecated: Use AppActionRequest.ProtoReflect.Descriptor instead.

func (x *AppActionRequest) GetArgs() []string

func (x *AppActionRequest) GetName() string

func (x *AppActionRequest) GetOpts() *AppOptions

func (*AppActionRequest) ProtoMessage()

func (x *AppActionRequest) ProtoReflect() protoreflect.Message

func (x *AppActionRequest) Reset()

func (x *AppActionRequest) String() string

type AppOptions struct {
	TimeoutMillis *int64 `protobuf:"varint,1,opt,name=timeout_millis,json=timeoutMillis,proto3,oneof" json:"timeout_millis,omitempty"`

	// Has unexported fields.
}

func (*AppOptions) Descriptor() ([]byte, []int)
    Deprecated: Use AppOptions.ProtoReflect.Descriptor instead.

func (x *AppOptions) GetTimeoutMillis() int64

func (*AppOptions) ProtoMessage()

func (x *AppOptions) ProtoReflect() protoreflect.Message

func (x *AppOptions) Reset()

func (x *AppOptions) String() string

type CaptureScreenRequest struct {
	ScreenId *int32 `protobuf:"varint,1,opt,name=screen_id,json=screenId,proto3,oneof" json:"screen_id,omitempty"`
	Region   *Rect  `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`

	// Has unexported fields.
}

func (*CaptureScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use CaptureScreenRequest.ProtoReflect.Descriptor instead.

func (x *CaptureScreenRequest) GetRegion() *Rect

func (x *CaptureScreenRequest) GetScreenId() int32

func (*CaptureScreenRequest) ProtoMessage()

func (x *CaptureScreenRequest) ProtoReflect() protoreflect.Message

func (x *CaptureScreenRequest) Reset()

func (x *CaptureScreenRequest) String() string

type CaptureScreenResponse struct {
	Image  *GrayImage        `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Screen *ScreenDescriptor `protobuf:"bytes,2,opt,name=screen,proto3" json:"screen,omitempty"`

	// Has unexported fields.
}

func (*CaptureScreenResponse) Descriptor() ([]byte, []int)
    Deprecated: Use CaptureScreenResponse.ProtoReflect.Descriptor instead.

func (x *CaptureScreenResponse) GetImage() *GrayImage

func (x *CaptureScreenResponse) GetScreen() *ScreenDescriptor

func (*CaptureScreenResponse) ProtoMessage()

func (x *CaptureScreenResponse) ProtoReflect() protoreflect.Message

func (x *CaptureScreenResponse) Reset()

func (x *CaptureScreenResponse) String() string

type ClickOnScreenRequest struct {
	Pattern   *Pattern            `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Opts      *ScreenQueryOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`
	ClickOpts *InputOptions       `protobuf:"bytes,3,opt,name=click_opts,json=clickOpts,proto3" json:"click_opts,omitempty"`

	// Has unexported fields.
}

func (*ClickOnScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ClickOnScreenRequest.ProtoReflect.Descriptor instead.

func (x *ClickOnScreenRequest) GetClickOpts() *InputOptions

func (x *ClickOnScreenRequest) GetOpts() *ScreenQueryOptions

func (x *ClickOnScreenRequest) GetPattern() *Pattern

func (*ClickOnScreenRequest) ProtoMessage()

func (x *ClickOnScreenRequest) ProtoReflect() protoreflect.Message

func (x *ClickOnScreenRequest) Reset()

func (x *ClickOnScreenRequest) String() string

type ClickRequest struct {
	X    int32         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y    int32         `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Opts *InputOptions `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*ClickRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ClickRequest.ProtoReflect.Descriptor instead.

func (x *ClickRequest) GetOpts() *InputOptions

func (x *ClickRequest) GetX() int32

func (x *ClickRequest) GetY() int32

func (*ClickRequest) ProtoMessage()

func (x *ClickRequest) ProtoReflect() protoreflect.Message

func (x *ClickRequest) Reset()

func (x *ClickRequest) String() string

type ExistsOnScreenRequest struct {
	Pattern *Pattern            `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Opts    *ScreenQueryOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*ExistsOnScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ExistsOnScreenRequest.ProtoReflect.Descriptor instead.

func (x *ExistsOnScreenRequest) GetOpts() *ScreenQueryOptions

func (x *ExistsOnScreenRequest) GetPattern() *Pattern

func (*ExistsOnScreenRequest) ProtoMessage()

func (x *ExistsOnScreenRequest) ProtoReflect() protoreflect.Message

func (x *ExistsOnScreenRequest) Reset()

func (x *ExistsOnScreenRequest) String() string

type ExistsOnScreenResponse struct {
	Exists bool   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	Match  *Match `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`

	// Has unexported fields.
}

func (*ExistsOnScreenResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ExistsOnScreenResponse.ProtoReflect.Descriptor instead.

func (x *ExistsOnScreenResponse) GetExists() bool

func (x *ExistsOnScreenResponse) GetMatch() *Match

func (*ExistsOnScreenResponse) ProtoMessage()

func (x *ExistsOnScreenResponse) ProtoReflect() protoreflect.Message

func (x *ExistsOnScreenResponse) Reset()

func (x *ExistsOnScreenResponse) String() string

type FindAllResponse struct {
	Matches []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`

	// Has unexported fields.
}

func (*FindAllResponse) Descriptor() ([]byte, []int)
    Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.

func (x *FindAllResponse) GetMatches() []*Match

func (*FindAllResponse) ProtoMessage()

func (x *FindAllResponse) ProtoReflect() protoreflect.Message

func (x *FindAllResponse) Reset()

func (x *FindAllResponse) String() string

type FindOnScreenRequest struct {
	Pattern *Pattern            `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Opts    *ScreenQueryOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*FindOnScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use FindOnScreenRequest.ProtoReflect.Descriptor instead.

func (x *FindOnScreenRequest) GetOpts() *ScreenQueryOptions

func (x *FindOnScreenRequest) GetPattern() *Pattern

func (*FindOnScreenRequest) ProtoMessage()

func (x *FindOnScreenRequest) ProtoReflect() protoreflect.Message

func (x *FindOnScreenRequest) Reset()

func (x *FindOnScreenRequest) String() string

type FindRequest struct {
	Source        *GrayImage    `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Pattern       *Pattern      `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	MatcherEngine MatcherEngine `protobuf:"varint,3,opt,name=matcher_engine,json=matcherEngine,proto3,enum=sikuli.v1.MatcherEngine" json:"matcher_engine,omitempty"`

	// Has unexported fields.
}

func (*FindRequest) Descriptor() ([]byte, []int)
    Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.

func (x *FindRequest) GetMatcherEngine() MatcherEngine

func (x *FindRequest) GetPattern() *Pattern

func (x *FindRequest) GetSource() *GrayImage

func (*FindRequest) ProtoMessage()

func (x *FindRequest) ProtoReflect() protoreflect.Message

func (x *FindRequest) Reset()

func (x *FindRequest) String() string

type FindResponse struct {
	Match *Match `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`

	// Has unexported fields.
}

func (*FindResponse) Descriptor() ([]byte, []int)
    Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.

func (x *FindResponse) GetMatch() *Match

func (*FindResponse) ProtoMessage()

func (x *FindResponse) ProtoReflect() protoreflect.Message

func (x *FindResponse) Reset()

func (x *FindResponse) String() string

type FindTextRequest struct {
	Source *GrayImage `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Query  string     `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Params *OCRParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`

	// Has unexported fields.
}

func (*FindTextRequest) Descriptor() ([]byte, []int)
    Deprecated: Use FindTextRequest.ProtoReflect.Descriptor instead.

func (x *FindTextRequest) GetParams() *OCRParams

func (x *FindTextRequest) GetQuery() string

func (x *FindTextRequest) GetSource() *GrayImage

func (*FindTextRequest) ProtoMessage()

func (x *FindTextRequest) ProtoReflect() protoreflect.Message

func (x *FindTextRequest) Reset()

func (x *FindTextRequest) String() string

type FindTextResponse struct {
	Matches []*TextMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`

	// Has unexported fields.
}

func (*FindTextResponse) Descriptor() ([]byte, []int)
    Deprecated: Use FindTextResponse.ProtoReflect.Descriptor instead.

func (x *FindTextResponse) GetMatches() []*TextMatch

func (*FindTextResponse) ProtoMessage()

func (x *FindTextResponse) ProtoReflect() protoreflect.Message

func (x *FindTextResponse) Reset()

func (x *FindTextResponse) String() string

type GetPrimaryScreenRequest struct {
	// Has unexported fields.
}

func (*GetPrimaryScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use GetPrimaryScreenRequest.ProtoReflect.Descriptor instead.

func (*GetPrimaryScreenRequest) ProtoMessage()

func (x *GetPrimaryScreenRequest) ProtoReflect() protoreflect.Message

func (x *GetPrimaryScreenRequest) Reset()

func (x *GetPrimaryScreenRequest) String() string

type GetPrimaryScreenResponse struct {
	Screen *ScreenDescriptor `protobuf:"bytes,1,opt,name=screen,proto3" json:"screen,omitempty"`

	// Has unexported fields.
}

func (*GetPrimaryScreenResponse) Descriptor() ([]byte, []int)
    Deprecated: Use GetPrimaryScreenResponse.ProtoReflect.Descriptor instead.

func (x *GetPrimaryScreenResponse) GetScreen() *ScreenDescriptor

func (*GetPrimaryScreenResponse) ProtoMessage()

func (x *GetPrimaryScreenResponse) ProtoReflect() protoreflect.Message

func (x *GetPrimaryScreenResponse) Reset()

func (x *GetPrimaryScreenResponse) String() string

type GetWindowResponse struct {
	Found  bool    `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Window *Window `protobuf:"bytes,2,opt,name=window,proto3" json:"window,omitempty"`

	// Has unexported fields.
}

func (*GetWindowResponse) Descriptor() ([]byte, []int)
    Deprecated: Use GetWindowResponse.ProtoReflect.Descriptor instead.

func (x *GetWindowResponse) GetFound() bool

func (x *GetWindowResponse) GetWindow() *Window

func (*GetWindowResponse) ProtoMessage()

func (x *GetWindowResponse) ProtoReflect() protoreflect.Message

func (x *GetWindowResponse) Reset()

func (x *GetWindowResponse) String() string

type GrayImage struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Pix    []byte `protobuf:"bytes,4,opt,name=pix,proto3" json:"pix,omitempty"`

	// Has unexported fields.
}

func (*GrayImage) Descriptor() ([]byte, []int)
    Deprecated: Use GrayImage.ProtoReflect.Descriptor instead.

func (x *GrayImage) GetHeight() int32

func (x *GrayImage) GetName() string

func (x *GrayImage) GetPix() []byte

func (x *GrayImage) GetWidth() int32

func (*GrayImage) ProtoMessage()

func (x *GrayImage) ProtoReflect() protoreflect.Message

func (x *GrayImage) Reset()

func (x *GrayImage) String() string

type HotkeyRequest struct {
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`

	// Has unexported fields.
}

func (*HotkeyRequest) Descriptor() ([]byte, []int)
    Deprecated: Use HotkeyRequest.ProtoReflect.Descriptor instead.

func (x *HotkeyRequest) GetKeys() []string

func (*HotkeyRequest) ProtoMessage()

func (x *HotkeyRequest) ProtoReflect() protoreflect.Message

func (x *HotkeyRequest) Reset()

func (x *HotkeyRequest) String() string

type InputOptions struct {
	DelayMillis *int64 `protobuf:"varint,1,opt,name=delay_millis,json=delayMillis,proto3,oneof" json:"delay_millis,omitempty"`
	Button      string `protobuf:"bytes,2,opt,name=button,proto3" json:"button,omitempty"`

	// Has unexported fields.
}

func (*InputOptions) Descriptor() ([]byte, []int)
    Deprecated: Use InputOptions.ProtoReflect.Descriptor instead.

func (x *InputOptions) GetButton() string

func (x *InputOptions) GetDelayMillis() int64

func (*InputOptions) ProtoMessage()

func (x *InputOptions) ProtoReflect() protoreflect.Message

func (x *InputOptions) Reset()

func (x *InputOptions) String() string

type IsAppRunningResponse struct {
	Running bool `protobuf:"varint,1,opt,name=running,proto3" json:"running,omitempty"`

	// Has unexported fields.
}

func (*IsAppRunningResponse) Descriptor() ([]byte, []int)
    Deprecated: Use IsAppRunningResponse.ProtoReflect.Descriptor instead.

func (x *IsAppRunningResponse) GetRunning() bool

func (*IsAppRunningResponse) ProtoMessage()

func (x *IsAppRunningResponse) ProtoReflect() protoreflect.Message

func (x *IsAppRunningResponse) Reset()

func (x *IsAppRunningResponse) String() string

type ListScreensRequest struct {
	// Has unexported fields.
}

func (*ListScreensRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ListScreensRequest.ProtoReflect.Descriptor instead.

func (*ListScreensRequest) ProtoMessage()

func (x *ListScreensRequest) ProtoReflect() protoreflect.Message

func (x *ListScreensRequest) Reset()

func (x *ListScreensRequest) String() string

type ListScreensResponse struct {
	Screens []*ScreenDescriptor `protobuf:"bytes,1,rep,name=screens,proto3" json:"screens,omitempty"`

	// Has unexported fields.
}

func (*ListScreensResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ListScreensResponse.ProtoReflect.Descriptor instead.

func (x *ListScreensResponse) GetScreens() []*ScreenDescriptor

func (*ListScreensResponse) ProtoMessage()

func (x *ListScreensResponse) ProtoReflect() protoreflect.Message

func (x *ListScreensResponse) Reset()

func (x *ListScreensResponse) String() string

type ListWindowsResponse struct {
	Windows []*Window `protobuf:"bytes,1,rep,name=windows,proto3" json:"windows,omitempty"`

	// Has unexported fields.
}

func (*ListWindowsResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ListWindowsResponse.ProtoReflect.Descriptor instead.

func (x *ListWindowsResponse) GetWindows() []*Window

func (*ListWindowsResponse) ProtoMessage()

func (x *ListWindowsResponse) ProtoReflect() protoreflect.Message

func (x *ListWindowsResponse) Reset()

func (x *ListWindowsResponse) String() string

type Match struct {
	Rect   *Rect   `protobuf:"bytes,1,opt,name=rect,proto3" json:"rect,omitempty"`
	Score  float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	Target *Point  `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Index  int32   `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`

	// Has unexported fields.
}

func (*Match) Descriptor() ([]byte, []int)
    Deprecated: Use Match.ProtoReflect.Descriptor instead.

func (x *Match) GetIndex() int32

func (x *Match) GetRect() *Rect

func (x *Match) GetScore() float64

func (x *Match) GetTarget() *Point

func (*Match) ProtoMessage()

func (x *Match) ProtoReflect() protoreflect.Message

func (x *Match) Reset()

func (x *Match) String() string

type MatcherEngine int32

const (
	MatcherEngine_MATCHER_ENGINE_UNSPECIFIED MatcherEngine = 0
	MatcherEngine_MATCHER_ENGINE_TEMPLATE    MatcherEngine = 1
	MatcherEngine_MATCHER_ENGINE_ORB         MatcherEngine = 2
	MatcherEngine_MATCHER_ENGINE_HYBRID      MatcherEngine = 3
	MatcherEngine_MATCHER_ENGINE_AKAZE       MatcherEngine = 4
	MatcherEngine_MATCHER_ENGINE_BRISK       MatcherEngine = 5
	MatcherEngine_MATCHER_ENGINE_KAZE        MatcherEngine = 6
	MatcherEngine_MATCHER_ENGINE_SIFT        MatcherEngine = 7
)
func (MatcherEngine) Descriptor() protoreflect.EnumDescriptor

func (x MatcherEngine) Enum() *MatcherEngine

func (MatcherEngine) EnumDescriptor() ([]byte, []int)
    Deprecated: Use MatcherEngine.Descriptor instead.

func (x MatcherEngine) Number() protoreflect.EnumNumber

func (x MatcherEngine) String() string

func (MatcherEngine) Type() protoreflect.EnumType

type MoveMouseRequest struct {
	X    int32         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y    int32         `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Opts *InputOptions `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*MoveMouseRequest) Descriptor() ([]byte, []int)
    Deprecated: Use MoveMouseRequest.ProtoReflect.Descriptor instead.

func (x *MoveMouseRequest) GetOpts() *InputOptions

func (x *MoveMouseRequest) GetX() int32

func (x *MoveMouseRequest) GetY() int32

func (*MoveMouseRequest) ProtoMessage()

func (x *MoveMouseRequest) ProtoReflect() protoreflect.Message

func (x *MoveMouseRequest) Reset()

func (x *MoveMouseRequest) String() string

type OCRParams struct {
	Language         string   `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	TrainingDataPath string   `protobuf:"bytes,2,opt,name=training_data_path,json=trainingDataPath,proto3" json:"training_data_path,omitempty"`
	MinConfidence    *float64 `protobuf:"fixed64,3,opt,name=min_confidence,json=minConfidence,proto3,oneof" json:"min_confidence,omitempty"`
	TimeoutMillis    *int64   `protobuf:"varint,4,opt,name=timeout_millis,json=timeoutMillis,proto3,oneof" json:"timeout_millis,omitempty"`
	CaseSensitive    *bool    `protobuf:"varint,5,opt,name=case_sensitive,json=caseSensitive,proto3,oneof" json:"case_sensitive,omitempty"`

	// Has unexported fields.
}

func (*OCRParams) Descriptor() ([]byte, []int)
    Deprecated: Use OCRParams.ProtoReflect.Descriptor instead.

func (x *OCRParams) GetCaseSensitive() bool

func (x *OCRParams) GetLanguage() string

func (x *OCRParams) GetMinConfidence() float64

func (x *OCRParams) GetTimeoutMillis() int64

func (x *OCRParams) GetTrainingDataPath() string

func (*OCRParams) ProtoMessage()

func (x *OCRParams) ProtoReflect() protoreflect.Message

func (x *OCRParams) Reset()

func (x *OCRParams) String() string

type ObserveChangeRequest struct {
	Source *GrayImage      `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Region *Rect           `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Opts   *ObserveOptions `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*ObserveChangeRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ObserveChangeRequest.ProtoReflect.Descriptor instead.

func (x *ObserveChangeRequest) GetOpts() *ObserveOptions

func (x *ObserveChangeRequest) GetRegion() *Rect

func (x *ObserveChangeRequest) GetSource() *GrayImage

func (*ObserveChangeRequest) ProtoMessage()

func (x *ObserveChangeRequest) ProtoReflect() protoreflect.Message

func (x *ObserveChangeRequest) Reset()

func (x *ObserveChangeRequest) String() string

type ObserveEvent struct {
	Type                string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Match               *Match `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	TimestampUnixMillis int64  `protobuf:"varint,3,opt,name=timestamp_unix_millis,json=timestampUnixMillis,proto3" json:"timestamp_unix_millis,omitempty"`

	// Has unexported fields.
}

func (*ObserveEvent) Descriptor() ([]byte, []int)
    Deprecated: Use ObserveEvent.ProtoReflect.Descriptor instead.

func (x *ObserveEvent) GetMatch() *Match

func (x *ObserveEvent) GetTimestampUnixMillis() int64

func (x *ObserveEvent) GetType() string

func (*ObserveEvent) ProtoMessage()

func (x *ObserveEvent) ProtoReflect() protoreflect.Message

func (x *ObserveEvent) Reset()

func (x *ObserveEvent) String() string

type ObserveOptions struct {
	IntervalMillis *int64 `protobuf:"varint,1,opt,name=interval_millis,json=intervalMillis,proto3,oneof" json:"interval_millis,omitempty"`
	TimeoutMillis  *int64 `protobuf:"varint,2,opt,name=timeout_millis,json=timeoutMillis,proto3,oneof" json:"timeout_millis,omitempty"`

	// Has unexported fields.
}

func (*ObserveOptions) Descriptor() ([]byte, []int)
    Deprecated: Use ObserveOptions.ProtoReflect.Descriptor instead.

func (x *ObserveOptions) GetIntervalMillis() int64

func (x *ObserveOptions) GetTimeoutMillis() int64

func (*ObserveOptions) ProtoMessage()

func (x *ObserveOptions) ProtoReflect() protoreflect.Message

func (x *ObserveOptions) Reset()

func (x *ObserveOptions) String() string

type ObserveRequest struct {
	Source  *GrayImage      `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Region  *Rect           `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Pattern *Pattern        `protobuf:"bytes,3,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Opts    *ObserveOptions `protobuf:"bytes,4,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*ObserveRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ObserveRequest.ProtoReflect.Descriptor instead.

func (x *ObserveRequest) GetOpts() *ObserveOptions

func (x *ObserveRequest) GetPattern() *Pattern

func (x *ObserveRequest) GetRegion() *Rect

func (x *ObserveRequest) GetSource() *GrayImage

func (*ObserveRequest) ProtoMessage()

func (x *ObserveRequest) ProtoReflect() protoreflect.Message

func (x *ObserveRequest) Reset()

func (x *ObserveRequest) String() string

type ObserveResponse struct {
	Events []*ObserveEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`

	// Has unexported fields.
}

func (*ObserveResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ObserveResponse.ProtoReflect.Descriptor instead.

func (x *ObserveResponse) GetEvents() []*ObserveEvent

func (*ObserveResponse) ProtoMessage()

func (x *ObserveResponse) ProtoReflect() protoreflect.Message

func (x *ObserveResponse) Reset()

func (x *ObserveResponse) String() string

type Pattern struct {
	Image        *GrayImage `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Similarity   *float64   `protobuf:"fixed64,2,opt,name=similarity,proto3,oneof" json:"similarity,omitempty"`
	Exact        *bool      `protobuf:"varint,3,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
	TargetOffset *Point     `protobuf:"bytes,4,opt,name=target_offset,json=targetOffset,proto3" json:"target_offset,omitempty"`
	ResizeFactor *float64   `protobuf:"fixed64,5,opt,name=resize_factor,json=resizeFactor,proto3,oneof" json:"resize_factor,omitempty"`
	Mask         *GrayImage `protobuf:"bytes,6,opt,name=mask,proto3" json:"mask,omitempty"`

	// Has unexported fields.
}

func (*Pattern) Descriptor() ([]byte, []int)
    Deprecated: Use Pattern.ProtoReflect.Descriptor instead.

func (x *Pattern) GetExact() bool

func (x *Pattern) GetImage() *GrayImage

func (x *Pattern) GetMask() *GrayImage

func (x *Pattern) GetResizeFactor() float64

func (x *Pattern) GetSimilarity() float64

func (x *Pattern) GetTargetOffset() *Point

func (*Pattern) ProtoMessage()

func (x *Pattern) ProtoReflect() protoreflect.Message

func (x *Pattern) Reset()

func (x *Pattern) String() string

type Point struct {
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`

	// Has unexported fields.
}

func (*Point) Descriptor() ([]byte, []int)
    Deprecated: Use Point.ProtoReflect.Descriptor instead.

func (x *Point) GetX() int32

func (x *Point) GetY() int32

func (*Point) ProtoMessage()

func (x *Point) ProtoReflect() protoreflect.Message

func (x *Point) Reset()

func (x *Point) String() string

type ReadTextRequest struct {
	Source *GrayImage `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Params *OCRParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`

	// Has unexported fields.
}

func (*ReadTextRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ReadTextRequest.ProtoReflect.Descriptor instead.

func (x *ReadTextRequest) GetParams() *OCRParams

func (x *ReadTextRequest) GetSource() *GrayImage

func (*ReadTextRequest) ProtoMessage()

func (x *ReadTextRequest) ProtoReflect() protoreflect.Message

func (x *ReadTextRequest) Reset()

func (x *ReadTextRequest) String() string

type ReadTextResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`

	// Has unexported fields.
}

func (*ReadTextResponse) Descriptor() ([]byte, []int)
    Deprecated: Use ReadTextResponse.ProtoReflect.Descriptor instead.

func (x *ReadTextResponse) GetText() string

func (*ReadTextResponse) ProtoMessage()

func (x *ReadTextResponse) ProtoReflect() protoreflect.Message

func (x *ReadTextResponse) Reset()

func (x *ReadTextResponse) String() string

type Rect struct {
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W int32 `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H int32 `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`

	// Has unexported fields.
}

func (*Rect) Descriptor() ([]byte, []int)
    Deprecated: Use Rect.ProtoReflect.Descriptor instead.

func (x *Rect) GetH() int32

func (x *Rect) GetW() int32

func (x *Rect) GetX() int32

func (x *Rect) GetY() int32

func (*Rect) ProtoMessage()

func (x *Rect) ProtoReflect() protoreflect.Message

func (x *Rect) Reset()

func (x *Rect) String() string

type ScreenDescriptor struct {
	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bounds  *Rect  `protobuf:"bytes,3,opt,name=bounds,proto3" json:"bounds,omitempty"`
	Primary bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`

	// Has unexported fields.
}

func (*ScreenDescriptor) Descriptor() ([]byte, []int)
    Deprecated: Use ScreenDescriptor.ProtoReflect.Descriptor instead.

func (x *ScreenDescriptor) GetBounds() *Rect

func (x *ScreenDescriptor) GetId() int32

func (x *ScreenDescriptor) GetName() string

func (x *ScreenDescriptor) GetPrimary() bool

func (*ScreenDescriptor) ProtoMessage()

func (x *ScreenDescriptor) ProtoReflect() protoreflect.Message

func (x *ScreenDescriptor) Reset()

func (x *ScreenDescriptor) String() string

type ScreenQueryOptions struct {
	Region         *Rect         `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	TimeoutMillis  *int64        `protobuf:"varint,2,opt,name=timeout_millis,json=timeoutMillis,proto3,oneof" json:"timeout_millis,omitempty"`
	IntervalMillis *int64        `protobuf:"varint,3,opt,name=interval_millis,json=intervalMillis,proto3,oneof" json:"interval_millis,omitempty"`
	MatcherEngine  MatcherEngine `protobuf:"varint,4,opt,name=matcher_engine,json=matcherEngine,proto3,enum=sikuli.v1.MatcherEngine" json:"matcher_engine,omitempty"`
	ScreenId       *int32        `protobuf:"varint,5,opt,name=screen_id,json=screenId,proto3,oneof" json:"screen_id,omitempty"`

	// Has unexported fields.
}

func (*ScreenQueryOptions) Descriptor() ([]byte, []int)
    Deprecated: Use ScreenQueryOptions.ProtoReflect.Descriptor instead.

func (x *ScreenQueryOptions) GetIntervalMillis() int64

func (x *ScreenQueryOptions) GetMatcherEngine() MatcherEngine

func (x *ScreenQueryOptions) GetRegion() *Rect

func (x *ScreenQueryOptions) GetScreenId() int32

func (x *ScreenQueryOptions) GetTimeoutMillis() int64

func (*ScreenQueryOptions) ProtoMessage()

func (x *ScreenQueryOptions) ProtoReflect() protoreflect.Message

func (x *ScreenQueryOptions) Reset()

func (x *ScreenQueryOptions) String() string

type ScrollWheelRequest struct {
	X         int32         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y         int32         `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Direction string        `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Steps     int32         `protobuf:"varint,4,opt,name=steps,proto3" json:"steps,omitempty"`
	Opts      *InputOptions `protobuf:"bytes,5,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*ScrollWheelRequest) Descriptor() ([]byte, []int)
    Deprecated: Use ScrollWheelRequest.ProtoReflect.Descriptor instead.

func (x *ScrollWheelRequest) GetDirection() string

func (x *ScrollWheelRequest) GetOpts() *InputOptions

func (x *ScrollWheelRequest) GetSteps() int32

func (x *ScrollWheelRequest) GetX() int32

func (x *ScrollWheelRequest) GetY() int32

func (*ScrollWheelRequest) ProtoMessage()

func (x *ScrollWheelRequest) ProtoReflect() protoreflect.Message

func (x *ScrollWheelRequest) Reset()

func (x *ScrollWheelRequest) String() string

type SikuliServiceClient interface {
	ListScreens(ctx context.Context, in *ListScreensRequest, opts ...grpc.CallOption) (*ListScreensResponse, error)
	GetPrimaryScreen(ctx context.Context, in *GetPrimaryScreenRequest, opts ...grpc.CallOption) (*GetPrimaryScreenResponse, error)
	CaptureScreen(ctx context.Context, in *CaptureScreenRequest, opts ...grpc.CallOption) (*CaptureScreenResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	FindAll(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	FindOnScreen(ctx context.Context, in *FindOnScreenRequest, opts ...grpc.CallOption) (*FindResponse, error)
	ExistsOnScreen(ctx context.Context, in *ExistsOnScreenRequest, opts ...grpc.CallOption) (*ExistsOnScreenResponse, error)
	WaitOnScreen(ctx context.Context, in *WaitOnScreenRequest, opts ...grpc.CallOption) (*FindResponse, error)
	ClickOnScreen(ctx context.Context, in *ClickOnScreenRequest, opts ...grpc.CallOption) (*FindResponse, error)
	ReadText(ctx context.Context, in *ReadTextRequest, opts ...grpc.CallOption) (*ReadTextResponse, error)
	FindText(ctx context.Context, in *FindTextRequest, opts ...grpc.CallOption) (*FindTextResponse, error)
	MoveMouse(ctx context.Context, in *MoveMouseRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	Click(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	TypeText(ctx context.Context, in *TypeTextRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	PasteText(ctx context.Context, in *TypeTextRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	Hotkey(ctx context.Context, in *HotkeyRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	MouseDown(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	MouseUp(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	KeyDown(ctx context.Context, in *HotkeyRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	KeyUp(ctx context.Context, in *HotkeyRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	ScrollWheel(ctx context.Context, in *ScrollWheelRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	ObserveAppear(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	ObserveVanish(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	ObserveChange(ctx context.Context, in *ObserveChangeRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	OpenApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	FocusApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	CloseApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	IsAppRunning(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*IsAppRunningResponse, error)
	ListWindows(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ListWindowsResponse, error)
	FindWindows(ctx context.Context, in *WindowQueryRequest, opts ...grpc.CallOption) (*ListWindowsResponse, error)
	GetWindow(ctx context.Context, in *WindowQueryRequest, opts ...grpc.CallOption) (*GetWindowResponse, error)
	GetFocusedWindow(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*GetWindowResponse, error)
}
    SikuliServiceClient is the client API for SikuliService service.

    For semantics around ctx use and closing/ending streaming RPCs, please refer
    to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.

func NewSikuliServiceClient(cc grpc.ClientConnInterface) SikuliServiceClient

type SikuliServiceServer interface {
	ListScreens(context.Context, *ListScreensRequest) (*ListScreensResponse, error)
	GetPrimaryScreen(context.Context, *GetPrimaryScreenRequest) (*GetPrimaryScreenResponse, error)
	CaptureScreen(context.Context, *CaptureScreenRequest) (*CaptureScreenResponse, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	FindAll(context.Context, *FindRequest) (*FindAllResponse, error)
	FindOnScreen(context.Context, *FindOnScreenRequest) (*FindResponse, error)
	ExistsOnScreen(context.Context, *ExistsOnScreenRequest) (*ExistsOnScreenResponse, error)
	WaitOnScreen(context.Context, *WaitOnScreenRequest) (*FindResponse, error)
	ClickOnScreen(context.Context, *ClickOnScreenRequest) (*FindResponse, error)
	ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)
	FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)
	MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)
	Click(context.Context, *ClickRequest) (*ActionResponse, error)
	TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)
	PasteText(context.Context, *TypeTextRequest) (*ActionResponse, error)
	Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)
	MouseDown(context.Context, *ClickRequest) (*ActionResponse, error)
	MouseUp(context.Context, *ClickRequest) (*ActionResponse, error)
	KeyDown(context.Context, *HotkeyRequest) (*ActionResponse, error)
	KeyUp(context.Context, *HotkeyRequest) (*ActionResponse, error)
	ScrollWheel(context.Context, *ScrollWheelRequest) (*ActionResponse, error)
	ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)
	ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)
	ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)
	OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)
	ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)
	FindWindows(context.Context, *WindowQueryRequest) (*ListWindowsResponse, error)
	GetWindow(context.Context, *WindowQueryRequest) (*GetWindowResponse, error)
	GetFocusedWindow(context.Context, *AppActionRequest) (*GetWindowResponse, error)
	// Has unexported methods.
}
    SikuliServiceServer is the server API for SikuliService service.
    All implementations must embed UnimplementedSikuliServiceServer for forward
    compatibility.

type TextMatch struct {
	Rect       *Rect   `protobuf:"bytes,1,opt,name=rect,proto3" json:"rect,omitempty"`
	Text       string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Confidence float64 `protobuf:"fixed64,3,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Index      int32   `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`

	// Has unexported fields.
}

func (*TextMatch) Descriptor() ([]byte, []int)
    Deprecated: Use TextMatch.ProtoReflect.Descriptor instead.

func (x *TextMatch) GetConfidence() float64

func (x *TextMatch) GetIndex() int32

func (x *TextMatch) GetRect() *Rect

func (x *TextMatch) GetText() string

func (*TextMatch) ProtoMessage()

func (x *TextMatch) ProtoReflect() protoreflect.Message

func (x *TextMatch) Reset()

func (x *TextMatch) String() string

type TypeTextRequest struct {
	Text string        `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Opts *InputOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*TypeTextRequest) Descriptor() ([]byte, []int)
    Deprecated: Use TypeTextRequest.ProtoReflect.Descriptor instead.

func (x *TypeTextRequest) GetOpts() *InputOptions

func (x *TypeTextRequest) GetText() string

func (*TypeTextRequest) ProtoMessage()

func (x *TypeTextRequest) ProtoReflect() protoreflect.Message

func (x *TypeTextRequest) Reset()

func (x *TypeTextRequest) String() string

type UnimplementedSikuliServiceServer struct{}
    UnimplementedSikuliServiceServer must be embedded to have forward compatible
    implementations.

    NOTE: this should be embedded by value instead of pointer to avoid a nil
    pointer dereference when methods are called.

func (UnimplementedSikuliServiceServer) CaptureScreen(context.Context, *CaptureScreenRequest) (*CaptureScreenResponse, error)

func (UnimplementedSikuliServiceServer) Click(context.Context, *ClickRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ClickOnScreen(context.Context, *ClickOnScreenRequest) (*FindResponse, error)

func (UnimplementedSikuliServiceServer) CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ExistsOnScreen(context.Context, *ExistsOnScreenRequest) (*ExistsOnScreenResponse, error)

func (UnimplementedSikuliServiceServer) Find(context.Context, *FindRequest) (*FindResponse, error)

func (UnimplementedSikuliServiceServer) FindAll(context.Context, *FindRequest) (*FindAllResponse, error)

func (UnimplementedSikuliServiceServer) FindOnScreen(context.Context, *FindOnScreenRequest) (*FindResponse, error)

func (UnimplementedSikuliServiceServer) FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)

func (UnimplementedSikuliServiceServer) FindWindows(context.Context, *WindowQueryRequest) (*ListWindowsResponse, error)

func (UnimplementedSikuliServiceServer) FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) GetFocusedWindow(context.Context, *AppActionRequest) (*GetWindowResponse, error)

func (UnimplementedSikuliServiceServer) GetPrimaryScreen(context.Context, *GetPrimaryScreenRequest) (*GetPrimaryScreenResponse, error)

func (UnimplementedSikuliServiceServer) GetWindow(context.Context, *WindowQueryRequest) (*GetWindowResponse, error)

func (UnimplementedSikuliServiceServer) Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)

func (UnimplementedSikuliServiceServer) KeyDown(context.Context, *HotkeyRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) KeyUp(context.Context, *HotkeyRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ListScreens(context.Context, *ListScreensRequest) (*ListScreensResponse, error)

func (UnimplementedSikuliServiceServer) ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)

func (UnimplementedSikuliServiceServer) MouseDown(context.Context, *ClickRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) MouseUp(context.Context, *ClickRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) PasteText(context.Context, *TypeTextRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)

func (UnimplementedSikuliServiceServer) ScrollWheel(context.Context, *ScrollWheelRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) WaitOnScreen(context.Context, *WaitOnScreenRequest) (*FindResponse, error)

type UnsafeSikuliServiceServer interface {
	// Has unexported methods.
}
    UnsafeSikuliServiceServer may be embedded to opt out of forward
    compatibility for this service. Use of this interface is not recommended,
    as added methods to SikuliServiceServer will result in compilation errors.

type WaitOnScreenRequest struct {
	Pattern *Pattern            `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Opts    *ScreenQueryOptions `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`

	// Has unexported fields.
}

func (*WaitOnScreenRequest) Descriptor() ([]byte, []int)
    Deprecated: Use WaitOnScreenRequest.ProtoReflect.Descriptor instead.

func (x *WaitOnScreenRequest) GetOpts() *ScreenQueryOptions

func (x *WaitOnScreenRequest) GetPattern() *Pattern

func (*WaitOnScreenRequest) ProtoMessage()

func (x *WaitOnScreenRequest) ProtoReflect() protoreflect.Message

func (x *WaitOnScreenRequest) Reset()

func (x *WaitOnScreenRequest) String() string

type Window struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	App     string `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Pid     int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Bounds  *Rect  `protobuf:"bytes,5,opt,name=bounds,proto3" json:"bounds,omitempty"`
	Focused bool   `protobuf:"varint,6,opt,name=focused,proto3" json:"focused,omitempty"`

	// Has unexported fields.
}

func (*Window) Descriptor() ([]byte, []int)
    Deprecated: Use Window.ProtoReflect.Descriptor instead.

func (x *Window) GetApp() string

func (x *Window) GetBounds() *Rect

func (x *Window) GetFocused() bool

func (x *Window) GetId() string

func (x *Window) GetPid() int32

func (x *Window) GetTitle() string

func (*Window) ProtoMessage()

func (x *Window) ProtoReflect() protoreflect.Message

func (x *Window) Reset()

func (x *Window) String() string

type WindowQuery struct {
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TitleExact    string `protobuf:"bytes,2,opt,name=title_exact,json=titleExact,proto3" json:"title_exact,omitempty"`
	TitleContains string `protobuf:"bytes,3,opt,name=title_contains,json=titleContains,proto3" json:"title_contains,omitempty"`
	FocusedOnly   bool   `protobuf:"varint,4,opt,name=focused_only,json=focusedOnly,proto3" json:"focused_only,omitempty"`
	Index         *int32 `protobuf:"varint,5,opt,name=index,proto3,oneof" json:"index,omitempty"`

	// Has unexported fields.
}

func (*WindowQuery) Descriptor() ([]byte, []int)
    Deprecated: Use WindowQuery.ProtoReflect.Descriptor instead.

func (x *WindowQuery) GetFocusedOnly() bool

func (x *WindowQuery) GetId() string

func (x *WindowQuery) GetIndex() int32

func (x *WindowQuery) GetTitleContains() string

func (x *WindowQuery) GetTitleExact() string

func (*WindowQuery) ProtoMessage()

func (x *WindowQuery) ProtoReflect() protoreflect.Message

func (x *WindowQuery) Reset()

func (x *WindowQuery) String() string

type WindowQueryRequest struct {
	Name  string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Opts  *AppOptions  `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`
	Query *WindowQuery `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`

	// Has unexported fields.
}

func (*WindowQueryRequest) Descriptor() ([]byte, []int)
    Deprecated: Use WindowQueryRequest.ProtoReflect.Descriptor instead.

func (x *WindowQueryRequest) GetName() string

func (x *WindowQueryRequest) GetOpts() *AppOptions

func (x *WindowQueryRequest) GetQuery() *WindowQuery

func (*WindowQueryRequest) ProtoMessage()

func (x *WindowQueryRequest) ProtoReflect() protoreflect.Message

func (x *WindowQueryRequest) Reset()

func (x *WindowQueryRequest) String() string

```
