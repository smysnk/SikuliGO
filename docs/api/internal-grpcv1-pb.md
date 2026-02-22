# API: `internal/grpcv1/pb`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package sikuliv1 // import "github.com/sikulix/portgo/internal/grpcv1/pb"`

## Symbol Index

### Types

- <span class="api-type">[`ActionResponse`](#type-actionresponse)</span>
- <span class="api-type">[`AppActionRequest`](#type-appactionrequest)</span>
- <span class="api-type">[`AppOptions`](#type-appoptions)</span>
- <span class="api-type">[`ClickRequest`](#type-clickrequest)</span>
- <span class="api-type">[`FindAllResponse`](#type-findallresponse)</span>
- <span class="api-type">[`FindRequest`](#type-findrequest)</span>
- <span class="api-type">[`FindResponse`](#type-findresponse)</span>
- <span class="api-type">[`FindTextRequest`](#type-findtextrequest)</span>
- <span class="api-type">[`FindTextResponse`](#type-findtextresponse)</span>
- <span class="api-type">[`GrayImage`](#type-grayimage)</span>
- <span class="api-type">[`HotkeyRequest`](#type-hotkeyrequest)</span>
- <span class="api-type">[`InputOptions`](#type-inputoptions)</span>
- <span class="api-type">[`IsAppRunningResponse`](#type-isapprunningresponse)</span>
- <span class="api-type">[`ListWindowsResponse`](#type-listwindowsresponse)</span>
- <span class="api-type">[`Match`](#type-match)</span>
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
- <span class="api-type">[`SikuliServiceClient`](#type-sikuliserviceclient)</span>
- <span class="api-type">[`SikuliServiceServer`](#type-sikuliserviceserver)</span>
- <span class="api-type">[`TextMatch`](#type-textmatch)</span>
- <span class="api-type">[`TypeTextRequest`](#type-typetextrequest)</span>
- <span class="api-type">[`UnimplementedSikuliServiceServer`](#type-unimplementedsikuliserviceserver)</span>
- <span class="api-type">[`UnsafeSikuliServiceServer`](#type-unsafesikuliserviceserver)</span>
- <span class="api-type">[`Window`](#type-window)</span>

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
- <span class="api-method">[`ClickRequest.Descriptor`](#method-clickrequest-descriptor)</span>
- <span class="api-method">[`ClickRequest.GetOpts`](#method-clickrequest-getopts)</span>
- <span class="api-method">[`ClickRequest.GetX`](#method-clickrequest-getx)</span>
- <span class="api-method">[`ClickRequest.GetY`](#method-clickrequest-gety)</span>
- <span class="api-method">[`ClickRequest.ProtoMessage`](#method-clickrequest-protomessage)</span>
- <span class="api-method">[`ClickRequest.ProtoReflect`](#method-clickrequest-protoreflect)</span>
- <span class="api-method">[`ClickRequest.Reset`](#method-clickrequest-reset)</span>
- <span class="api-method">[`ClickRequest.String`](#method-clickrequest-string)</span>
- <span class="api-method">[`FindAllResponse.Descriptor`](#method-findallresponse-descriptor)</span>
- <span class="api-method">[`FindAllResponse.GetMatches`](#method-findallresponse-getmatches)</span>
- <span class="api-method">[`FindAllResponse.ProtoMessage`](#method-findallresponse-protomessage)</span>
- <span class="api-method">[`FindAllResponse.ProtoReflect`](#method-findallresponse-protoreflect)</span>
- <span class="api-method">[`FindAllResponse.Reset`](#method-findallresponse-reset)</span>
- <span class="api-method">[`FindAllResponse.String`](#method-findallresponse-string)</span>
- <span class="api-method">[`FindRequest.Descriptor`](#method-findrequest-descriptor)</span>
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
- <span class="api-method">[`UnimplementedSikuliServiceServer.Click`](#method-unimplementedsikuliserviceserver-click)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.CloseApp`](#method-unimplementedsikuliserviceserver-closeapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.Find`](#method-unimplementedsikuliserviceserver-find)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindAll`](#method-unimplementedsikuliserviceserver-findall)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FindText`](#method-unimplementedsikuliserviceserver-findtext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.FocusApp`](#method-unimplementedsikuliserviceserver-focusapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.Hotkey`](#method-unimplementedsikuliserviceserver-hotkey)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.IsAppRunning`](#method-unimplementedsikuliserviceserver-isapprunning)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ListWindows`](#method-unimplementedsikuliserviceserver-listwindows)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.MoveMouse`](#method-unimplementedsikuliserviceserver-movemouse)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveAppear`](#method-unimplementedsikuliserviceserver-observeappear)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveChange`](#method-unimplementedsikuliserviceserver-observechange)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ObserveVanish`](#method-unimplementedsikuliserviceserver-observevanish)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.OpenApp`](#method-unimplementedsikuliserviceserver-openapp)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.ReadText`](#method-unimplementedsikuliserviceserver-readtext)</span>
- <span class="api-method">[`UnimplementedSikuliServiceServer.TypeText`](#method-unimplementedsikuliserviceserver-typetext)</span>
- <span class="api-method">[`Window.Descriptor`](#method-window-descriptor)</span>
- <span class="api-method">[`Window.GetBounds`](#method-window-getbounds)</span>
- <span class="api-method">[`Window.GetFocused`](#method-window-getfocused)</span>
- <span class="api-method">[`Window.GetTitle`](#method-window-gettitle)</span>
- <span class="api-method">[`Window.ProtoMessage`](#method-window-protomessage)</span>
- <span class="api-method">[`Window.ProtoReflect`](#method-window-protoreflect)</span>
- <span class="api-method">[`Window.Reset`](#method-window-reset)</span>
- <span class="api-method">[`Window.String`](#method-window-string)</span>

## Declarations

### Types

#### <a id="type-actionresponse"></a><span class="api-type">Type</span> `ActionResponse`

- Signature: <span class="api-signature">`type ActionResponse struct {`</span>

#### <a id="type-appactionrequest"></a><span class="api-type">Type</span> `AppActionRequest`

- Signature: <span class="api-signature">`type AppActionRequest struct {`</span>

#### <a id="type-appoptions"></a><span class="api-type">Type</span> `AppOptions`

- Signature: <span class="api-signature">`type AppOptions struct {`</span>

#### <a id="type-clickrequest"></a><span class="api-type">Type</span> `ClickRequest`

- Signature: <span class="api-signature">`type ClickRequest struct {`</span>

#### <a id="type-findallresponse"></a><span class="api-type">Type</span> `FindAllResponse`

- Signature: <span class="api-signature">`type FindAllResponse struct {`</span>

#### <a id="type-findrequest"></a><span class="api-type">Type</span> `FindRequest`

- Signature: <span class="api-signature">`type FindRequest struct {`</span>

#### <a id="type-findresponse"></a><span class="api-type">Type</span> `FindResponse`

- Signature: <span class="api-signature">`type FindResponse struct {`</span>

#### <a id="type-findtextrequest"></a><span class="api-type">Type</span> `FindTextRequest`

- Signature: <span class="api-signature">`type FindTextRequest struct {`</span>

#### <a id="type-findtextresponse"></a><span class="api-type">Type</span> `FindTextResponse`

- Signature: <span class="api-signature">`type FindTextResponse struct {`</span>

#### <a id="type-grayimage"></a><span class="api-type">Type</span> `GrayImage`

- Signature: <span class="api-signature">`type GrayImage struct {`</span>

#### <a id="type-hotkeyrequest"></a><span class="api-type">Type</span> `HotkeyRequest`

- Signature: <span class="api-signature">`type HotkeyRequest struct {`</span>

#### <a id="type-inputoptions"></a><span class="api-type">Type</span> `InputOptions`

- Signature: <span class="api-signature">`type InputOptions struct {`</span>

#### <a id="type-isapprunningresponse"></a><span class="api-type">Type</span> `IsAppRunningResponse`

- Signature: <span class="api-signature">`type IsAppRunningResponse struct {`</span>

#### <a id="type-listwindowsresponse"></a><span class="api-type">Type</span> `ListWindowsResponse`

- Signature: <span class="api-signature">`type ListWindowsResponse struct {`</span>

#### <a id="type-match"></a><span class="api-type">Type</span> `Match`

- Signature: <span class="api-signature">`type Match struct {`</span>

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

#### <a id="type-window"></a><span class="api-type">Type</span> `Window`

- Signature: <span class="api-signature">`type Window struct {`</span>

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

#### <a id="method-findrequest-descriptor"></a><span class="api-method">Method</span> `FindRequest.Descriptor`

- Signature: <span class="api-signature">`func (*FindRequest) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.

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

#### <a id="method-unimplementedsikuliserviceserver-click"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Click`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Click(context.Context, *ClickRequest) (*ActionResponse, error)`</span>
- Uses: [`ClickRequest`](#type-clickrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-closeapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.CloseApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-find"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Find`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Find(context.Context, *FindRequest) (*FindResponse, error)`</span>
- Uses: [`FindRequest`](#type-findrequest), [`FindResponse`](#type-findresponse)

#### <a id="method-unimplementedsikuliserviceserver-findall"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindAll`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindAll(context.Context, *FindRequest) (*FindAllResponse, error)`</span>
- Uses: [`FindRequest`](#type-findrequest), [`FindAllResponse`](#type-findallresponse)

#### <a id="method-unimplementedsikuliserviceserver-findtext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FindText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)`</span>
- Uses: [`FindTextResponse`](#type-findtextresponse), [`FindTextRequest`](#type-findtextrequest)

#### <a id="method-unimplementedsikuliserviceserver-focusapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.FocusApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-hotkey"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.Hotkey`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)`</span>
- Uses: [`HotkeyRequest`](#type-hotkeyrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-isapprunning"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.IsAppRunning`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`IsAppRunningResponse`](#type-isapprunningresponse)

#### <a id="method-unimplementedsikuliserviceserver-listwindows"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ListWindows`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)`</span>
- Uses: [`ListWindowsResponse`](#type-listwindowsresponse), [`AppActionRequest`](#type-appactionrequest)

#### <a id="method-unimplementedsikuliserviceserver-movemouse"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.MoveMouse`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)`</span>
- Uses: [`MoveMouseRequest`](#type-movemouserequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-observeappear"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveAppear`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveResponse`](#type-observeresponse), [`ObserveRequest`](#type-observerequest)

#### <a id="method-unimplementedsikuliserviceserver-observechange"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveChange`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveChangeRequest`](#type-observechangerequest), [`ObserveResponse`](#type-observeresponse)

#### <a id="method-unimplementedsikuliserviceserver-observevanish"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ObserveVanish`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)`</span>
- Uses: [`ObserveResponse`](#type-observeresponse), [`ObserveRequest`](#type-observerequest)

#### <a id="method-unimplementedsikuliserviceserver-openapp"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.OpenApp`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)`</span>
- Uses: [`AppActionRequest`](#type-appactionrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-unimplementedsikuliserviceserver-readtext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.ReadText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)`</span>
- Uses: [`ReadTextResponse`](#type-readtextresponse), [`ReadTextRequest`](#type-readtextrequest)

#### <a id="method-unimplementedsikuliserviceserver-typetext"></a><span class="api-method">Method</span> `UnimplementedSikuliServiceServer.TypeText`

- Signature: <span class="api-signature">`func (UnimplementedSikuliServiceServer) TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)`</span>
- Uses: [`TypeTextRequest`](#type-typetextrequest), [`ActionResponse`](#type-actionresponse)

#### <a id="method-window-descriptor"></a><span class="api-method">Method</span> `Window.Descriptor`

- Signature: <span class="api-signature">`func (*Window) Descriptor() ([]byte, []int)`</span>
- Notes: Deprecated: Use Window.ProtoReflect.Descriptor instead.

#### <a id="method-window-getbounds"></a><span class="api-method">Method</span> `Window.GetBounds`

- Signature: <span class="api-signature">`func (x *Window) GetBounds() *Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-window-getfocused"></a><span class="api-method">Method</span> `Window.GetFocused`

- Signature: <span class="api-signature">`func (x *Window) GetFocused() bool`</span>

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

## Raw Package Doc

```text
package sikuliv1 // import "github.com/sikulix/portgo/internal/grpcv1/pb"


CONSTANTS

const (
	SikuliService_Find_FullMethodName          = "/sikuli.v1.SikuliService/Find"
	SikuliService_FindAll_FullMethodName       = "/sikuli.v1.SikuliService/FindAll"
	SikuliService_ReadText_FullMethodName      = "/sikuli.v1.SikuliService/ReadText"
	SikuliService_FindText_FullMethodName      = "/sikuli.v1.SikuliService/FindText"
	SikuliService_MoveMouse_FullMethodName     = "/sikuli.v1.SikuliService/MoveMouse"
	SikuliService_Click_FullMethodName         = "/sikuli.v1.SikuliService/Click"
	SikuliService_TypeText_FullMethodName      = "/sikuli.v1.SikuliService/TypeText"
	SikuliService_Hotkey_FullMethodName        = "/sikuli.v1.SikuliService/Hotkey"
	SikuliService_ObserveAppear_FullMethodName = "/sikuli.v1.SikuliService/ObserveAppear"
	SikuliService_ObserveVanish_FullMethodName = "/sikuli.v1.SikuliService/ObserveVanish"
	SikuliService_ObserveChange_FullMethodName = "/sikuli.v1.SikuliService/ObserveChange"
	SikuliService_OpenApp_FullMethodName       = "/sikuli.v1.SikuliService/OpenApp"
	SikuliService_FocusApp_FullMethodName      = "/sikuli.v1.SikuliService/FocusApp"
	SikuliService_CloseApp_FullMethodName      = "/sikuli.v1.SikuliService/CloseApp"
	SikuliService_IsAppRunning_FullMethodName  = "/sikuli.v1.SikuliService/IsAppRunning"
	SikuliService_ListWindows_FullMethodName   = "/sikuli.v1.SikuliService/ListWindows"
)

VARIABLES

var File_sikuli_v1_sikuli_proto protoreflect.FileDescriptor
var SikuliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sikuli.v1.SikuliService",
	HandlerType: (*SikuliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _SikuliService_Find_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _SikuliService_FindAll_Handler,
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
			MethodName: "Hotkey",
			Handler:    _SikuliService_Hotkey_Handler,
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

type FindRequest struct {
	Source  *GrayImage `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Pattern *Pattern   `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`

	// Has unexported fields.
}

func (*FindRequest) Descriptor() ([]byte, []int)
    Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.

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

type SikuliServiceClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	FindAll(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	ReadText(ctx context.Context, in *ReadTextRequest, opts ...grpc.CallOption) (*ReadTextResponse, error)
	FindText(ctx context.Context, in *FindTextRequest, opts ...grpc.CallOption) (*FindTextResponse, error)
	MoveMouse(ctx context.Context, in *MoveMouseRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	Click(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	TypeText(ctx context.Context, in *TypeTextRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	Hotkey(ctx context.Context, in *HotkeyRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	ObserveAppear(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	ObserveVanish(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	ObserveChange(ctx context.Context, in *ObserveChangeRequest, opts ...grpc.CallOption) (*ObserveResponse, error)
	OpenApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	FocusApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	CloseApp(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	IsAppRunning(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*IsAppRunningResponse, error)
	ListWindows(ctx context.Context, in *AppActionRequest, opts ...grpc.CallOption) (*ListWindowsResponse, error)
}
    SikuliServiceClient is the client API for SikuliService service.

    For semantics around ctx use and closing/ending streaming RPCs, please refer
    to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.

func NewSikuliServiceClient(cc grpc.ClientConnInterface) SikuliServiceClient

type SikuliServiceServer interface {
	Find(context.Context, *FindRequest) (*FindResponse, error)
	FindAll(context.Context, *FindRequest) (*FindAllResponse, error)
	ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)
	FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)
	MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)
	Click(context.Context, *ClickRequest) (*ActionResponse, error)
	TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)
	Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)
	ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)
	ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)
	ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)
	OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)
	IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)
	ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)
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

func (UnimplementedSikuliServiceServer) Click(context.Context, *ClickRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) CloseApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) Find(context.Context, *FindRequest) (*FindResponse, error)

func (UnimplementedSikuliServiceServer) FindAll(context.Context, *FindRequest) (*FindAllResponse, error)

func (UnimplementedSikuliServiceServer) FindText(context.Context, *FindTextRequest) (*FindTextResponse, error)

func (UnimplementedSikuliServiceServer) FocusApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) Hotkey(context.Context, *HotkeyRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) IsAppRunning(context.Context, *AppActionRequest) (*IsAppRunningResponse, error)

func (UnimplementedSikuliServiceServer) ListWindows(context.Context, *AppActionRequest) (*ListWindowsResponse, error)

func (UnimplementedSikuliServiceServer) MoveMouse(context.Context, *MoveMouseRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ObserveAppear(context.Context, *ObserveRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) ObserveChange(context.Context, *ObserveChangeRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) ObserveVanish(context.Context, *ObserveRequest) (*ObserveResponse, error)

func (UnimplementedSikuliServiceServer) OpenApp(context.Context, *AppActionRequest) (*ActionResponse, error)

func (UnimplementedSikuliServiceServer) ReadText(context.Context, *ReadTextRequest) (*ReadTextResponse, error)

func (UnimplementedSikuliServiceServer) TypeText(context.Context, *TypeTextRequest) (*ActionResponse, error)

type UnsafeSikuliServiceServer interface {
	// Has unexported methods.
}
    UnsafeSikuliServiceServer may be embedded to opt out of forward
    compatibility for this service. Use of this interface is not recommended,
    as added methods to SikuliServiceServer will result in compilation errors.

type Window struct {
	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Bounds  *Rect  `protobuf:"bytes,2,opt,name=bounds,proto3" json:"bounds,omitempty"`
	Focused bool   `protobuf:"varint,3,opt,name=focused,proto3" json:"focused,omitempty"`

	// Has unexported fields.
}

func (*Window) Descriptor() ([]byte, []int)
    Deprecated: Use Window.ProtoReflect.Descriptor instead.

func (x *Window) GetBounds() *Rect

func (x *Window) GetFocused() bool

func (x *Window) GetTitle() string

func (*Window) ProtoMessage()

func (x *Window) ProtoReflect() protoreflect.Message

func (x *Window) Reset()

func (x *Window) String() string

```
