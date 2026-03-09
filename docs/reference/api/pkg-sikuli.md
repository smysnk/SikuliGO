# API: `pkg/sikuli`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package sikuli // import "github.com/smysnk/sikuligo/pkg/sikuli"`

## Symbol Index

### Types

- <span class="api-type">[`AppAPI`](#type-appapi)</span>
- <span class="api-type">[`AppController`](#type-appcontroller)</span>
- <span class="api-type">[`AppOptions`](#type-appoptions)</span>
- <span class="api-type">[`Finder`](#type-finder)</span>
- <span class="api-type">[`FinderAPI`](#type-finderapi)</span>
- <span class="api-type">[`Image`](#type-image)</span>
- <span class="api-type">[`ImageAPI`](#type-imageapi)</span>
- <span class="api-type">[`InputAPI`](#type-inputapi)</span>
- <span class="api-type">[`InputController`](#type-inputcontroller)</span>
- <span class="api-type">[`InputOptions`](#type-inputoptions)</span>
- <span class="api-type">[`LiveRegion`](#type-liveregion)</span>
- <span class="api-type">[`LiveRegionAPI`](#type-liveregionapi)</span>
- <span class="api-type">[`Location`](#type-location)</span>
- <span class="api-type">[`Match`](#type-match)</span>
- <span class="api-type">[`MatchAPI`](#type-matchapi)</span>
- <span class="api-type">[`MatcherEngine`](#type-matcherengine)</span>
- <span class="api-type">[`MouseButton`](#type-mousebutton)</span>
- <span class="api-type">[`OCRLine`](#type-ocrline)</span>
- <span class="api-type">[`OCRParams`](#type-ocrparams)</span>
- <span class="api-type">[`OCRWord`](#type-ocrword)</span>
- <span class="api-type">[`ObserveAPI`](#type-observeapi)</span>
- <span class="api-type">[`ObserveEvent`](#type-observeevent)</span>
- <span class="api-type">[`ObserveEventType`](#type-observeeventtype)</span>
- <span class="api-type">[`ObserveOptions`](#type-observeoptions)</span>
- <span class="api-type">[`ObserverController`](#type-observercontroller)</span>
- <span class="api-type">[`Offset`](#type-offset)</span>
- <span class="api-type">[`Options`](#type-options)</span>
- <span class="api-type">[`Pattern`](#type-pattern)</span>
- <span class="api-type">[`PatternAPI`](#type-patternapi)</span>
- <span class="api-type">[`Point`](#type-point)</span>
- <span class="api-type">[`Rect`](#type-rect)</span>
- <span class="api-type">[`Region`](#type-region)</span>
- <span class="api-type">[`RegionAPI`](#type-regionapi)</span>
- <span class="api-type">[`Runtime`](#type-runtime)</span>
- <span class="api-type">[`RuntimeAPI`](#type-runtimeapi)</span>
- <span class="api-type">[`RuntimeOption`](#type-runtimeoption)</span>
- <span class="api-type">[`RuntimeSettings`](#type-runtimesettings)</span>
- <span class="api-type">[`Screen`](#type-screen)</span>
- <span class="api-type">[`ScreenAPI`](#type-screenapi)</span>
- <span class="api-type">[`SearchProbe`](#type-searchprobe)</span>
- <span class="api-type">[`TargetPointProvider`](#type-targetpointprovider)</span>
- <span class="api-type">[`TextMatch`](#type-textmatch)</span>
- <span class="api-type">[`WheelDirection`](#type-wheeldirection)</span>
- <span class="api-type">[`Window`](#type-window)</span>
- <span class="api-type">[`WindowQuery`](#type-windowquery)</span>

### Functions

- <span class="api-func">[`SearchWaitVanish`](#func-searchwaitvanish)</span>
- <span class="api-func">[`SortMatchesByColumnRow`](#func-sortmatchesbycolumnrow)</span>
- <span class="api-func">[`SortMatchesByRowColumn`](#func-sortmatchesbyrowcolumn)</span>
- <span class="api-func">[`NewAppController`](#func-newappcontroller)</span>
- <span class="api-func">[`NewFinder`](#func-newfinder)</span>
- <span class="api-func">[`NewImageFromAny`](#func-newimagefromany)</span>
- <span class="api-func">[`NewImageFromGray`](#func-newimagefromgray)</span>
- <span class="api-func">[`NewImageFromMatrix`](#func-newimagefrommatrix)</span>
- <span class="api-func">[`NewInputController`](#func-newinputcontroller)</span>
- <span class="api-func">[`NewLocation`](#func-newlocation)</span>
- <span class="api-func">[`NewMatch`](#func-newmatch)</span>
- <span class="api-func">[`SearchExists`](#func-searchexists)</span>
- <span class="api-func">[`SearchWait`](#func-searchwait)</span>
- <span class="api-func">[`NewObserverController`](#func-newobservercontroller)</span>
- <span class="api-func">[`NewOffset`](#func-newoffset)</span>
- <span class="api-func">[`NewOptions`](#func-newoptions)</span>
- <span class="api-func">[`NewOptionsFromMap`](#func-newoptionsfrommap)</span>
- <span class="api-func">[`NewPattern`](#func-newpattern)</span>
- <span class="api-func">[`NewPoint`](#func-newpoint)</span>
- <span class="api-func">[`NewRect`](#func-newrect)</span>
- <span class="api-func">[`NewRegion`](#func-newregion)</span>
- <span class="api-func">[`NewRuntime`](#func-newruntime)</span>
- <span class="api-func">[`WithRuntimeAuthToken`](#func-withruntimeauthtoken)</span>
- <span class="api-func">[`WithRuntimeConn`](#func-withruntimeconn)</span>
- <span class="api-func">[`WithRuntimeContextDialer`](#func-withruntimecontextdialer)</span>
- <span class="api-func">[`WithRuntimeDialTimeout`](#func-withruntimedialtimeout)</span>
- <span class="api-func">[`WithRuntimeMatcherEngine`](#func-withruntimematcherengine)</span>
- <span class="api-func">[`WithRuntimeRPCTimeout`](#func-withruntimerpctimeout)</span>
- <span class="api-func">[`GetSettings`](#func-getsettings)</span>
- <span class="api-func">[`ResetSettings`](#func-resetsettings)</span>
- <span class="api-func">[`UpdateSettings`](#func-updatesettings)</span>
- <span class="api-func">[`NewScreen`](#func-newscreen)</span>

### Methods

- <span class="api-method">[`AppController.Close`](#method-appcontroller-close)</span>
- <span class="api-method">[`AppController.FindWindows`](#method-appcontroller-findwindows)</span>
- <span class="api-method">[`AppController.Focus`](#method-appcontroller-focus)</span>
- <span class="api-method">[`AppController.FocusedWindow`](#method-appcontroller-focusedwindow)</span>
- <span class="api-method">[`AppController.GetWindow`](#method-appcontroller-getwindow)</span>
- <span class="api-method">[`AppController.IsRunning`](#method-appcontroller-isrunning)</span>
- <span class="api-method">[`AppController.ListWindows`](#method-appcontroller-listwindows)</span>
- <span class="api-method">[`AppController.Open`](#method-appcontroller-open)</span>
- <span class="api-method">[`AppController.SetBackend`](#method-appcontroller-setbackend)</span>
- <span class="api-method">[`Finder.CollectLines`](#method-finder-collectlines)</span>
- <span class="api-method">[`Finder.CollectWords`](#method-finder-collectwords)</span>
- <span class="api-method">[`Finder.Count`](#method-finder-count)</span>
- <span class="api-method">[`Finder.Destroy`](#method-finder-destroy)</span>
- <span class="api-method">[`Finder.Exists`](#method-finder-exists)</span>
- <span class="api-method">[`Finder.Find`](#method-finder-find)</span>
- <span class="api-method">[`Finder.FindAll`](#method-finder-findall)</span>
- <span class="api-method">[`Finder.FindAllByColumn`](#method-finder-findallbycolumn)</span>
- <span class="api-method">[`Finder.FindAllByRow`](#method-finder-findallbyrow)</span>
- <span class="api-method">[`Finder.FindAnyList`](#method-finder-findanylist)</span>
- <span class="api-method">[`Finder.FindBestList`](#method-finder-findbestlist)</span>
- <span class="api-method">[`Finder.FindText`](#method-finder-findtext)</span>
- <span class="api-method">[`Finder.Has`](#method-finder-has)</span>
- <span class="api-method">[`Finder.HasNext`](#method-finder-hasnext)</span>
- <span class="api-method">[`Finder.Iterate`](#method-finder-iterate)</span>
- <span class="api-method">[`Finder.IterateAll`](#method-finder-iterateall)</span>
- <span class="api-method">[`Finder.LastMatches`](#method-finder-lastmatches)</span>
- <span class="api-method">[`Finder.Next`](#method-finder-next)</span>
- <span class="api-method">[`Finder.ReadText`](#method-finder-readtext)</span>
- <span class="api-method">[`Finder.Reset`](#method-finder-reset)</span>
- <span class="api-method">[`Finder.SetMatcher`](#method-finder-setmatcher)</span>
- <span class="api-method">[`Finder.SetOCRBackend`](#method-finder-setocrbackend)</span>
- <span class="api-method">[`Finder.Wait`](#method-finder-wait)</span>
- <span class="api-method">[`Finder.WaitAnyList`](#method-finder-waitanylist)</span>
- <span class="api-method">[`Finder.WaitBestList`](#method-finder-waitbestlist)</span>
- <span class="api-method">[`Finder.WaitVanish`](#method-finder-waitvanish)</span>
- <span class="api-method">[`Image.Clone`](#method-image-clone)</span>
- <span class="api-method">[`Image.Crop`](#method-image-crop)</span>
- <span class="api-method">[`Image.Gray`](#method-image-gray)</span>
- <span class="api-method">[`Image.Height`](#method-image-height)</span>
- <span class="api-method">[`Image.Name`](#method-image-name)</span>
- <span class="api-method">[`Image.Width`](#method-image-width)</span>
- <span class="api-method">[`InputController.Click`](#method-inputcontroller-click)</span>
- <span class="api-method">[`InputController.DoubleClick`](#method-inputcontroller-doubleclick)</span>
- <span class="api-method">[`InputController.DragDrop`](#method-inputcontroller-dragdrop)</span>
- <span class="api-method">[`InputController.Hotkey`](#method-inputcontroller-hotkey)</span>
- <span class="api-method">[`InputController.Hover`](#method-inputcontroller-hover)</span>
- <span class="api-method">[`InputController.KeyDown`](#method-inputcontroller-keydown)</span>
- <span class="api-method">[`InputController.KeyUp`](#method-inputcontroller-keyup)</span>
- <span class="api-method">[`InputController.MouseDown`](#method-inputcontroller-mousedown)</span>
- <span class="api-method">[`InputController.MouseUp`](#method-inputcontroller-mouseup)</span>
- <span class="api-method">[`InputController.MoveMouse`](#method-inputcontroller-movemouse)</span>
- <span class="api-method">[`InputController.Paste`](#method-inputcontroller-paste)</span>
- <span class="api-method">[`InputController.RightClick`](#method-inputcontroller-rightclick)</span>
- <span class="api-method">[`InputController.SetBackend`](#method-inputcontroller-setbackend)</span>
- <span class="api-method">[`InputController.TypeText`](#method-inputcontroller-typetext)</span>
- <span class="api-method">[`InputController.Wheel`](#method-inputcontroller-wheel)</span>
- <span class="api-method">[`LiveRegion.Bounds`](#method-liveregion-bounds)</span>
- <span class="api-method">[`LiveRegion.Capture`](#method-liveregion-capture)</span>
- <span class="api-method">[`LiveRegion.Center`](#method-liveregion-center)</span>
- <span class="api-method">[`LiveRegion.Click`](#method-liveregion-click)</span>
- <span class="api-method">[`LiveRegion.CollectLines`](#method-liveregion-collectlines)</span>
- <span class="api-method">[`LiveRegion.CollectWords`](#method-liveregion-collectwords)</span>
- <span class="api-method">[`LiveRegion.DoubleClick`](#method-liveregion-doubleclick)</span>
- <span class="api-method">[`LiveRegion.DragDrop`](#method-liveregion-dragdrop)</span>
- <span class="api-method">[`LiveRegion.Exists`](#method-liveregion-exists)</span>
- <span class="api-method">[`LiveRegion.Find`](#method-liveregion-find)</span>
- <span class="api-method">[`LiveRegion.FindAnyList`](#method-liveregion-findanylist)</span>
- <span class="api-method">[`LiveRegion.FindBestList`](#method-liveregion-findbestlist)</span>
- <span class="api-method">[`LiveRegion.FindText`](#method-liveregion-findtext)</span>
- <span class="api-method">[`LiveRegion.Grow`](#method-liveregion-grow)</span>
- <span class="api-method">[`LiveRegion.Has`](#method-liveregion-has)</span>
- <span class="api-method">[`LiveRegion.Hover`](#method-liveregion-hover)</span>
- <span class="api-method">[`LiveRegion.KeyDown`](#method-liveregion-keydown)</span>
- <span class="api-method">[`LiveRegion.KeyUp`](#method-liveregion-keyup)</span>
- <span class="api-method">[`LiveRegion.MouseDown`](#method-liveregion-mousedown)</span>
- <span class="api-method">[`LiveRegion.MouseUp`](#method-liveregion-mouseup)</span>
- <span class="api-method">[`LiveRegion.MoveTo`](#method-liveregion-moveto)</span>
- <span class="api-method">[`LiveRegion.Offset`](#method-liveregion-offset)</span>
- <span class="api-method">[`LiveRegion.Paste`](#method-liveregion-paste)</span>
- <span class="api-method">[`LiveRegion.ReadText`](#method-liveregion-readtext)</span>
- <span class="api-method">[`LiveRegion.RightClick`](#method-liveregion-rightclick)</span>
- <span class="api-method">[`LiveRegion.SetSize`](#method-liveregion-setsize)</span>
- <span class="api-method">[`LiveRegion.TargetPoint`](#method-liveregion-targetpoint)</span>
- <span class="api-method">[`LiveRegion.TypeText`](#method-liveregion-typetext)</span>
- <span class="api-method">[`LiveRegion.Wait`](#method-liveregion-wait)</span>
- <span class="api-method">[`LiveRegion.WaitAnyList`](#method-liveregion-waitanylist)</span>
- <span class="api-method">[`LiveRegion.WaitBestList`](#method-liveregion-waitbestlist)</span>
- <span class="api-method">[`LiveRegion.WaitVanish`](#method-liveregion-waitvanish)</span>
- <span class="api-method">[`LiveRegion.Wheel`](#method-liveregion-wheel)</span>
- <span class="api-method">[`LiveRegion.WithMatcherEngine`](#method-liveregion-withmatcherengine)</span>
- <span class="api-method">[`Location.Move`](#method-location-move)</span>
- <span class="api-method">[`Location.String`](#method-location-string)</span>
- <span class="api-method">[`Location.TargetPoint`](#method-location-targetpoint)</span>
- <span class="api-method">[`Location.ToPoint`](#method-location-topoint)</span>
- <span class="api-method">[`Match.Bounds`](#method-match-bounds)</span>
- <span class="api-method">[`Match.Capture`](#method-match-capture)</span>
- <span class="api-method">[`Match.Center`](#method-match-center)</span>
- <span class="api-method">[`Match.Click`](#method-match-click)</span>
- <span class="api-method">[`Match.CollectLines`](#method-match-collectlines)</span>
- <span class="api-method">[`Match.CollectWords`](#method-match-collectwords)</span>
- <span class="api-method">[`Match.DoubleClick`](#method-match-doubleclick)</span>
- <span class="api-method">[`Match.DragDrop`](#method-match-dragdrop)</span>
- <span class="api-method">[`Match.Exists`](#method-match-exists)</span>
- <span class="api-method">[`Match.Find`](#method-match-find)</span>
- <span class="api-method">[`Match.FindAnyList`](#method-match-findanylist)</span>
- <span class="api-method">[`Match.FindBestList`](#method-match-findbestlist)</span>
- <span class="api-method">[`Match.FindText`](#method-match-findtext)</span>
- <span class="api-method">[`Match.Has`](#method-match-has)</span>
- <span class="api-method">[`Match.Hover`](#method-match-hover)</span>
- <span class="api-method">[`Match.KeyDown`](#method-match-keydown)</span>
- <span class="api-method">[`Match.KeyUp`](#method-match-keyup)</span>
- <span class="api-method">[`Match.Live`](#method-match-live)</span>
- <span class="api-method">[`Match.MouseDown`](#method-match-mousedown)</span>
- <span class="api-method">[`Match.MouseUp`](#method-match-mouseup)</span>
- <span class="api-method">[`Match.MoveMouse`](#method-match-movemouse)</span>
- <span class="api-method">[`Match.Paste`](#method-match-paste)</span>
- <span class="api-method">[`Match.ReadText`](#method-match-readtext)</span>
- <span class="api-method">[`Match.Region`](#method-match-region)</span>
- <span class="api-method">[`Match.RightClick`](#method-match-rightclick)</span>
- <span class="api-method">[`Match.String`](#method-match-string)</span>
- <span class="api-method">[`Match.TargetPoint`](#method-match-targetpoint)</span>
- <span class="api-method">[`Match.TypeText`](#method-match-typetext)</span>
- <span class="api-method">[`Match.Wait`](#method-match-wait)</span>
- <span class="api-method">[`Match.WaitAnyList`](#method-match-waitanylist)</span>
- <span class="api-method">[`Match.WaitBestList`](#method-match-waitbestlist)</span>
- <span class="api-method">[`Match.WaitVanish`](#method-match-waitvanish)</span>
- <span class="api-method">[`Match.Wheel`](#method-match-wheel)</span>
- <span class="api-method">[`ObserverController.ObserveAppear`](#method-observercontroller-observeappear)</span>
- <span class="api-method">[`ObserverController.ObserveChange`](#method-observercontroller-observechange)</span>
- <span class="api-method">[`ObserverController.ObserveVanish`](#method-observercontroller-observevanish)</span>
- <span class="api-method">[`ObserverController.SetBackend`](#method-observercontroller-setbackend)</span>
- <span class="api-method">[`Offset.String`](#method-offset-string)</span>
- <span class="api-method">[`Offset.ToPoint`](#method-offset-topoint)</span>
- <span class="api-method">[`Options.Clone`](#method-options-clone)</span>
- <span class="api-method">[`Options.Delete`](#method-options-delete)</span>
- <span class="api-method">[`Options.Entries`](#method-options-entries)</span>
- <span class="api-method">[`Options.GetBool`](#method-options-getbool)</span>
- <span class="api-method">[`Options.GetFloat64`](#method-options-getfloat64)</span>
- <span class="api-method">[`Options.GetInt`](#method-options-getint)</span>
- <span class="api-method">[`Options.GetString`](#method-options-getstring)</span>
- <span class="api-method">[`Options.Has`](#method-options-has)</span>
- <span class="api-method">[`Options.Merge`](#method-options-merge)</span>
- <span class="api-method">[`Options.SetBool`](#method-options-setbool)</span>
- <span class="api-method">[`Options.SetFloat64`](#method-options-setfloat64)</span>
- <span class="api-method">[`Options.SetInt`](#method-options-setint)</span>
- <span class="api-method">[`Options.SetString`](#method-options-setstring)</span>
- <span class="api-method">[`Pattern.Exact`](#method-pattern-exact)</span>
- <span class="api-method">[`Pattern.Image`](#method-pattern-image)</span>
- <span class="api-method">[`Pattern.Mask`](#method-pattern-mask)</span>
- <span class="api-method">[`Pattern.Offset`](#method-pattern-offset)</span>
- <span class="api-method">[`Pattern.Resize`](#method-pattern-resize)</span>
- <span class="api-method">[`Pattern.ResizeFactor`](#method-pattern-resizefactor)</span>
- <span class="api-method">[`Pattern.Similar`](#method-pattern-similar)</span>
- <span class="api-method">[`Pattern.Similarity`](#method-pattern-similarity)</span>
- <span class="api-method">[`Pattern.TargetOffset`](#method-pattern-targetoffset)</span>
- <span class="api-method">[`Pattern.WithMask`](#method-pattern-withmask)</span>
- <span class="api-method">[`Pattern.WithMaskMatrix`](#method-pattern-withmaskmatrix)</span>
- <span class="api-method">[`Point.TargetPoint`](#method-point-targetpoint)</span>
- <span class="api-method">[`Point.ToLocation`](#method-point-tolocation)</span>
- <span class="api-method">[`Point.ToOffset`](#method-point-tooffset)</span>
- <span class="api-method">[`Rect.Contains`](#method-rect-contains)</span>
- <span class="api-method">[`Rect.Empty`](#method-rect-empty)</span>
- <span class="api-method">[`Rect.String`](#method-rect-string)</span>
- <span class="api-method">[`Region.Center`](#method-region-center)</span>
- <span class="api-method">[`Region.CollectLines`](#method-region-collectlines)</span>
- <span class="api-method">[`Region.CollectWords`](#method-region-collectwords)</span>
- <span class="api-method">[`Region.Contains`](#method-region-contains)</span>
- <span class="api-method">[`Region.ContainsLocation`](#method-region-containslocation)</span>
- <span class="api-method">[`Region.ContainsRegion`](#method-region-containsregion)</span>
- <span class="api-method">[`Region.Count`](#method-region-count)</span>
- <span class="api-method">[`Region.Exists`](#method-region-exists)</span>
- <span class="api-method">[`Region.Find`](#method-region-find)</span>
- <span class="api-method">[`Region.FindAll`](#method-region-findall)</span>
- <span class="api-method">[`Region.FindAllByColumn`](#method-region-findallbycolumn)</span>
- <span class="api-method">[`Region.FindAllByRow`](#method-region-findallbyrow)</span>
- <span class="api-method">[`Region.FindAnyList`](#method-region-findanylist)</span>
- <span class="api-method">[`Region.FindBestList`](#method-region-findbestlist)</span>
- <span class="api-method">[`Region.FindText`](#method-region-findtext)</span>
- <span class="api-method">[`Region.Grow`](#method-region-grow)</span>
- <span class="api-method">[`Region.Has`](#method-region-has)</span>
- <span class="api-method">[`Region.Intersection`](#method-region-intersection)</span>
- <span class="api-method">[`Region.MoveTo`](#method-region-moveto)</span>
- <span class="api-method">[`Region.MoveToLocation`](#method-region-movetolocation)</span>
- <span class="api-method">[`Region.Offset`](#method-region-offset)</span>
- <span class="api-method">[`Region.OffsetBy`](#method-region-offsetby)</span>
- <span class="api-method">[`Region.ReadText`](#method-region-readtext)</span>
- <span class="api-method">[`Region.ResetThrowException`](#method-region-resetthrowexception)</span>
- <span class="api-method">[`Region.SetAutoWaitTimeout`](#method-region-setautowaittimeout)</span>
- <span class="api-method">[`Region.SetObserveScanRate`](#method-region-setobservescanrate)</span>
- <span class="api-method">[`Region.SetSize`](#method-region-setsize)</span>
- <span class="api-method">[`Region.SetThrowException`](#method-region-setthrowexception)</span>
- <span class="api-method">[`Region.SetWaitScanRate`](#method-region-setwaitscanrate)</span>
- <span class="api-method">[`Region.TargetPoint`](#method-region-targetpoint)</span>
- <span class="api-method">[`Region.Union`](#method-region-union)</span>
- <span class="api-method">[`Region.Wait`](#method-region-wait)</span>
- <span class="api-method">[`Region.WaitAnyList`](#method-region-waitanylist)</span>
- <span class="api-method">[`Region.WaitBestList`](#method-region-waitbestlist)</span>
- <span class="api-method">[`Region.WaitVanish`](#method-region-waitvanish)</span>
- <span class="api-method">[`Runtime.Address`](#method-runtime-address)</span>
- <span class="api-method">[`Runtime.Capture`](#method-runtime-capture)</span>
- <span class="api-method">[`Runtime.CaptureRegion`](#method-runtime-captureregion)</span>
- <span class="api-method">[`Runtime.Close`](#method-runtime-close)</span>
- <span class="api-method">[`Runtime.PrimaryScreen`](#method-runtime-primaryscreen)</span>
- <span class="api-method">[`Runtime.Region`](#method-runtime-region)</span>
- <span class="api-method">[`Runtime.Screen`](#method-runtime-screen)</span>
- <span class="api-method">[`Runtime.Screens`](#method-runtime-screens)</span>
- <span class="api-method">[`Screen.Capture`](#method-screen-capture)</span>
- <span class="api-method">[`Screen.Click`](#method-screen-click)</span>
- <span class="api-method">[`Screen.CollectLines`](#method-screen-collectlines)</span>
- <span class="api-method">[`Screen.CollectWords`](#method-screen-collectwords)</span>
- <span class="api-method">[`Screen.DoubleClick`](#method-screen-doubleclick)</span>
- <span class="api-method">[`Screen.DragDrop`](#method-screen-dragdrop)</span>
- <span class="api-method">[`Screen.Exists`](#method-screen-exists)</span>
- <span class="api-method">[`Screen.Find`](#method-screen-find)</span>
- <span class="api-method">[`Screen.FindAnyList`](#method-screen-findanylist)</span>
- <span class="api-method">[`Screen.FindBestList`](#method-screen-findbestlist)</span>
- <span class="api-method">[`Screen.FindText`](#method-screen-findtext)</span>
- <span class="api-method">[`Screen.FullRegion`](#method-screen-fullregion)</span>
- <span class="api-method">[`Screen.Has`](#method-screen-has)</span>
- <span class="api-method">[`Screen.Hover`](#method-screen-hover)</span>
- <span class="api-method">[`Screen.KeyDown`](#method-screen-keydown)</span>
- <span class="api-method">[`Screen.KeyUp`](#method-screen-keyup)</span>
- <span class="api-method">[`Screen.Live`](#method-screen-live)</span>
- <span class="api-method">[`Screen.MouseDown`](#method-screen-mousedown)</span>
- <span class="api-method">[`Screen.MouseUp`](#method-screen-mouseup)</span>
- <span class="api-method">[`Screen.Paste`](#method-screen-paste)</span>
- <span class="api-method">[`Screen.ReadText`](#method-screen-readtext)</span>
- <span class="api-method">[`Screen.Region`](#method-screen-region)</span>
- <span class="api-method">[`Screen.RegionRect`](#method-screen-regionrect)</span>
- <span class="api-method">[`Screen.RightClick`](#method-screen-rightclick)</span>
- <span class="api-method">[`Screen.TargetPoint`](#method-screen-targetpoint)</span>
- <span class="api-method">[`Screen.TypeText`](#method-screen-typetext)</span>
- <span class="api-method">[`Screen.Wait`](#method-screen-wait)</span>
- <span class="api-method">[`Screen.WaitAnyList`](#method-screen-waitanylist)</span>
- <span class="api-method">[`Screen.WaitBestList`](#method-screen-waitbestlist)</span>
- <span class="api-method">[`Screen.WaitVanish`](#method-screen-waitvanish)</span>
- <span class="api-method">[`Screen.Wheel`](#method-screen-wheel)</span>

## Declarations

### Types

#### <a id="type-appapi"></a><span class="api-type">Type</span> `AppAPI`

- Signature: <span class="api-signature">`type AppAPI interface {`</span>

#### <a id="type-appcontroller"></a><span class="api-type">Type</span> `AppController`

- Signature: <span class="api-signature">`type AppController struct {`</span>

#### <a id="type-appoptions"></a><span class="api-type">Type</span> `AppOptions`

- Signature: <span class="api-signature">`type AppOptions struct {`</span>

#### <a id="type-finder"></a><span class="api-type">Type</span> `Finder`

- Signature: <span class="api-signature">`type Finder struct {`</span>

#### <a id="type-finderapi"></a><span class="api-type">Type</span> `FinderAPI`

- Signature: <span class="api-signature">`type FinderAPI interface {`</span>

#### <a id="type-image"></a><span class="api-type">Type</span> `Image`

- Signature: <span class="api-signature">`type Image struct {`</span>

#### <a id="type-imageapi"></a><span class="api-type">Type</span> `ImageAPI`

- Signature: <span class="api-signature">`type ImageAPI interface {`</span>

#### <a id="type-inputapi"></a><span class="api-type">Type</span> `InputAPI`

- Signature: <span class="api-signature">`type InputAPI interface {`</span>

#### <a id="type-inputcontroller"></a><span class="api-type">Type</span> `InputController`

- Signature: <span class="api-signature">`type InputController struct {`</span>

#### <a id="type-inputoptions"></a><span class="api-type">Type</span> `InputOptions`

- Signature: <span class="api-signature">`type InputOptions struct {`</span>

#### <a id="type-liveregion"></a><span class="api-type">Type</span> `LiveRegion`

- Signature: <span class="api-signature">`type LiveRegion struct {`</span>

#### <a id="type-liveregionapi"></a><span class="api-type">Type</span> `LiveRegionAPI`

- Signature: <span class="api-signature">`type LiveRegionAPI interface {`</span>

#### <a id="type-location"></a><span class="api-type">Type</span> `Location`

- Signature: <span class="api-signature">`type Location struct {`</span>

#### <a id="type-match"></a><span class="api-type">Type</span> `Match`

- Signature: <span class="api-signature">`type Match struct {`</span>

#### <a id="type-matchapi"></a><span class="api-type">Type</span> `MatchAPI`

- Signature: <span class="api-signature">`type MatchAPI interface {`</span>

#### <a id="type-matcherengine"></a><span class="api-type">Type</span> `MatcherEngine`

- Signature: <span class="api-signature">`type MatcherEngine string`</span>
- Notes: MatcherEngine selects the server-side matcher implementation used for live screen queries.

#### <a id="type-mousebutton"></a><span class="api-type">Type</span> `MouseButton`

- Signature: <span class="api-signature">`type MouseButton string`</span>

#### <a id="type-ocrline"></a><span class="api-type">Type</span> `OCRLine`

- Signature: <span class="api-signature">`type OCRLine struct {`</span>

#### <a id="type-ocrparams"></a><span class="api-type">Type</span> `OCRParams`

- Signature: <span class="api-signature">`type OCRParams struct {`</span>

#### <a id="type-ocrword"></a><span class="api-type">Type</span> `OCRWord`

- Signature: <span class="api-signature">`type OCRWord struct {`</span>

#### <a id="type-observeapi"></a><span class="api-type">Type</span> `ObserveAPI`

- Signature: <span class="api-signature">`type ObserveAPI interface {`</span>

#### <a id="type-observeevent"></a><span class="api-type">Type</span> `ObserveEvent`

- Signature: <span class="api-signature">`type ObserveEvent struct {`</span>

#### <a id="type-observeeventtype"></a><span class="api-type">Type</span> `ObserveEventType`

- Signature: <span class="api-signature">`type ObserveEventType string`</span>

#### <a id="type-observeoptions"></a><span class="api-type">Type</span> `ObserveOptions`

- Signature: <span class="api-signature">`type ObserveOptions struct {`</span>

#### <a id="type-observercontroller"></a><span class="api-type">Type</span> `ObserverController`

- Signature: <span class="api-signature">`type ObserverController struct {`</span>

#### <a id="type-offset"></a><span class="api-type">Type</span> `Offset`

- Signature: <span class="api-signature">`type Offset struct {`</span>

#### <a id="type-options"></a><span class="api-type">Type</span> `Options`

- Signature: <span class="api-signature">`type Options struct {`</span>

#### <a id="type-pattern"></a><span class="api-type">Type</span> `Pattern`

- Signature: <span class="api-signature">`type Pattern struct {`</span>

#### <a id="type-patternapi"></a><span class="api-type">Type</span> `PatternAPI`

- Signature: <span class="api-signature">`type PatternAPI interface {`</span>

#### <a id="type-point"></a><span class="api-type">Type</span> `Point`

- Signature: <span class="api-signature">`type Point struct {`</span>

#### <a id="type-rect"></a><span class="api-type">Type</span> `Rect`

- Signature: <span class="api-signature">`type Rect struct {`</span>

#### <a id="type-region"></a><span class="api-type">Type</span> `Region`

- Signature: <span class="api-signature">`type Region struct {`</span>

#### <a id="type-regionapi"></a><span class="api-type">Type</span> `RegionAPI`

- Signature: <span class="api-signature">`type RegionAPI interface {`</span>

#### <a id="type-runtime"></a><span class="api-type">Type</span> `Runtime`

- Signature: <span class="api-signature">`type Runtime struct {`</span>

#### <a id="type-runtimeapi"></a><span class="api-type">Type</span> `RuntimeAPI`

- Signature: <span class="api-signature">`type RuntimeAPI interface {`</span>

#### <a id="type-runtimeoption"></a><span class="api-type">Type</span> `RuntimeOption`

- Signature: <span class="api-signature">`type RuntimeOption func(*runtimeConfig)`</span>

#### <a id="type-runtimesettings"></a><span class="api-type">Type</span> `RuntimeSettings`

- Signature: <span class="api-signature">`type RuntimeSettings struct {`</span>

#### <a id="type-screen"></a><span class="api-type">Type</span> `Screen`

- Signature: <span class="api-signature">`type Screen struct {`</span>

#### <a id="type-screenapi"></a><span class="api-type">Type</span> `ScreenAPI`

- Signature: <span class="api-signature">`type ScreenAPI interface {`</span>

#### <a id="type-searchprobe"></a><span class="api-type">Type</span> `SearchProbe`

- Signature: <span class="api-signature">`type SearchProbe func() (Match, error)`</span>
- Uses: [`Match`](#type-match)
- Notes: SearchProbe returns the next match candidate for a parity search operation. Returning ErrFindFailed indicates a miss; any other error aborts the search.

#### <a id="type-targetpointprovider"></a><span class="api-type">Type</span> `TargetPointProvider`

- Signature: <span class="api-signature">`type TargetPointProvider interface {`</span>

#### <a id="type-textmatch"></a><span class="api-type">Type</span> `TextMatch`

- Signature: <span class="api-signature">`type TextMatch struct {`</span>

#### <a id="type-wheeldirection"></a><span class="api-type">Type</span> `WheelDirection`

- Signature: <span class="api-signature">`type WheelDirection string`</span>

#### <a id="type-window"></a><span class="api-type">Type</span> `Window`

- Signature: <span class="api-signature">`type Window struct {`</span>

#### <a id="type-windowquery"></a><span class="api-type">Type</span> `WindowQuery`

- Signature: <span class="api-signature">`type WindowQuery struct {`</span>

### Functions

#### <a id="func-searchwaitvanish"></a><span class="api-func">Function</span> `SearchWaitVanish`

- Signature: <span class="api-signature">`func SearchWaitVanish(probe SearchProbe, timeout, interval time.Duration) (bool, error)`</span>
- Uses: [`SearchProbe`](#type-searchprobe)
- Notes: SearchWaitVanish applies the canonical vanish contract to a search probe. It returns true when the target is absent and false on timeout.

#### <a id="func-sortmatchesbycolumnrow"></a><span class="api-func">Function</span> `SortMatchesByColumnRow`

- Signature: <span class="api-signature">`func SortMatchesByColumnRow(matches []Match)`</span>
- Uses: [`Match`](#type-match)
- Notes: SortMatchesByColumnRow keeps parity with Java helper behavior for "by column".

#### <a id="func-sortmatchesbyrowcolumn"></a><span class="api-func">Function</span> `SortMatchesByRowColumn`

- Signature: <span class="api-signature">`func SortMatchesByRowColumn(matches []Match)`</span>
- Uses: [`Match`](#type-match)
- Notes: SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".

#### <a id="func-newappcontroller"></a><span class="api-func">Function</span> `NewAppController`

- Signature: <span class="api-signature">`func NewAppController() *AppController`</span>
- Uses: [`AppController`](#type-appcontroller)

#### <a id="func-newfinder"></a><span class="api-func">Function</span> `NewFinder`

- Signature: <span class="api-signature">`func NewFinder(source *Image) (*Finder, error)`</span>
- Uses: [`Finder`](#type-finder), [`Image`](#type-image)
- Notes: NewFinder creates a search/OCR helper bound to a source image.

#### <a id="func-newimagefromany"></a><span class="api-func">Function</span> `NewImageFromAny`

- Signature: <span class="api-signature">`func NewImageFromAny(name string, src image.Image) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newimagefromgray"></a><span class="api-func">Function</span> `NewImageFromGray`

- Signature: <span class="api-signature">`func NewImageFromGray(name string, src *image.Gray) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newimagefrommatrix"></a><span class="api-func">Function</span> `NewImageFromMatrix`

- Signature: <span class="api-signature">`func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newinputcontroller"></a><span class="api-func">Function</span> `NewInputController`

- Signature: <span class="api-signature">`func NewInputController() *InputController`</span>
- Uses: [`InputController`](#type-inputcontroller)

#### <a id="func-newlocation"></a><span class="api-func">Function</span> `NewLocation`

- Signature: <span class="api-signature">`func NewLocation(x, y int) Location`</span>
- Uses: [`Location`](#type-location)

#### <a id="func-newmatch"></a><span class="api-func">Function</span> `NewMatch`

- Signature: <span class="api-signature">`func NewMatch(x, y, w, h int, score float64, off Point) Match`</span>
- Uses: [`Match`](#type-match), [`Point`](#type-point)

#### <a id="func-searchexists"></a><span class="api-func">Function</span> `SearchExists`

- Signature: <span class="api-signature">`func SearchExists(probe SearchProbe, timeout, interval time.Duration) (Match, bool, error)`</span>
- Uses: [`Match`](#type-match), [`SearchProbe`](#type-searchprobe)
- Notes: SearchExists applies the canonical sikuli-go parity contract to a search probe. Misses are reported as (Match{}, false, nil); timeout <= 0 performs one probe.

#### <a id="func-searchwait"></a><span class="api-func">Function</span> `SearchWait`

- Signature: <span class="api-signature">`func SearchWait(probe SearchProbe, timeout, interval time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`SearchProbe`](#type-searchprobe)
- Notes: SearchWait applies the canonical sikuli-go wait contract to a search probe. Misses are promoted to ErrTimeout once the wait budget is exhausted.

#### <a id="func-newobservercontroller"></a><span class="api-func">Function</span> `NewObserverController`

- Signature: <span class="api-signature">`func NewObserverController() *ObserverController`</span>
- Uses: [`ObserverController`](#type-observercontroller)

#### <a id="func-newoffset"></a><span class="api-func">Function</span> `NewOffset`

- Signature: <span class="api-signature">`func NewOffset(x, y int) Offset`</span>
- Uses: [`Offset`](#type-offset)

#### <a id="func-newoptions"></a><span class="api-func">Function</span> `NewOptions`

- Signature: <span class="api-signature">`func NewOptions() *Options`</span>
- Uses: [`Options`](#type-options)

#### <a id="func-newoptionsfrommap"></a><span class="api-func">Function</span> `NewOptionsFromMap`

- Signature: <span class="api-signature">`func NewOptionsFromMap(entries map[string]string) *Options`</span>
- Uses: [`Options`](#type-options)

#### <a id="func-newpattern"></a><span class="api-func">Function</span> `NewPattern`

- Signature: <span class="api-signature">`func NewPattern(img *Image) (*Pattern, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)
- Notes: NewPattern creates a match pattern from an image with default similarity settings.

#### <a id="func-newpoint"></a><span class="api-func">Function</span> `NewPoint`

- Signature: <span class="api-signature">`func NewPoint(x, y int) Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="func-newrect"></a><span class="api-func">Function</span> `NewRect`

- Signature: <span class="api-signature">`func NewRect(x, y, w, h int) Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="func-newregion"></a><span class="api-func">Function</span> `NewRegion`

- Signature: <span class="api-signature">`func NewRegion(x, y, w, h int) Region`</span>
- Uses: [`Region`](#type-region)
- Notes: NewRegion constructs a rectangular search area with default timing settings.

#### <a id="func-newruntime"></a><span class="api-func">Function</span> `NewRuntime`

- Signature: <span class="api-signature">`func NewRuntime(address string, opts ...RuntimeOption) (*Runtime, error)`</span>
- Uses: [`Runtime`](#type-runtime), [`RuntimeOption`](#type-runtimeoption)
- Notes: NewRuntime connects to a running sikuli-go API runtime and exposes live screen operations.

#### <a id="func-withruntimeauthtoken"></a><span class="api-func">Function</span> `WithRuntimeAuthToken`

- Signature: <span class="api-signature">`func WithRuntimeAuthToken(token string) RuntimeOption`</span>
- Uses: [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-withruntimeconn"></a><span class="api-func">Function</span> `WithRuntimeConn`

- Signature: <span class="api-signature">`func WithRuntimeConn(conn *grpc.ClientConn) RuntimeOption`</span>
- Uses: [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-withruntimecontextdialer"></a><span class="api-func">Function</span> `WithRuntimeContextDialer`

- Signature: <span class="api-signature">`func WithRuntimeContextDialer(dialer func(context.Context, string) (net.Conn, error)) RuntimeOption`</span>
- Uses: [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-withruntimedialtimeout"></a><span class="api-func">Function</span> `WithRuntimeDialTimeout`

- Signature: <span class="api-signature">`func WithRuntimeDialTimeout(timeout time.Duration) RuntimeOption`</span>
- Uses: [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-withruntimematcherengine"></a><span class="api-func">Function</span> `WithRuntimeMatcherEngine`

- Signature: <span class="api-signature">`func WithRuntimeMatcherEngine(engine MatcherEngine) RuntimeOption`</span>
- Uses: [`MatcherEngine`](#type-matcherengine), [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-withruntimerpctimeout"></a><span class="api-func">Function</span> `WithRuntimeRPCTimeout`

- Signature: <span class="api-signature">`func WithRuntimeRPCTimeout(timeout time.Duration) RuntimeOption`</span>
- Uses: [`RuntimeOption`](#type-runtimeoption)

#### <a id="func-getsettings"></a><span class="api-func">Function</span> `GetSettings`

- Signature: <span class="api-signature">`func GetSettings() RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-resetsettings"></a><span class="api-func">Function</span> `ResetSettings`

- Signature: <span class="api-signature">`func ResetSettings() RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-updatesettings"></a><span class="api-func">Function</span> `UpdateSettings`

- Signature: <span class="api-signature">`func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-newscreen"></a><span class="api-func">Function</span> `NewScreen`

- Signature: <span class="api-signature">`func NewScreen(id int, bounds Rect) Screen`</span>
- Uses: [`Rect`](#type-rect), [`Screen`](#type-screen)
- Notes: NewScreen constructs a logical screen descriptor.

### Methods

#### <a id="method-appcontroller-close"></a><span class="api-method">Method</span> `AppController.Close`

- Signature: <span class="api-signature">`func (c *AppController) Close(name string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-findwindows"></a><span class="api-method">Method</span> `AppController.FindWindows`

- Signature: <span class="api-signature">`func (c *AppController) FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error)`</span>
- Uses: [`AppOptions`](#type-appoptions), [`Window`](#type-window), [`WindowQuery`](#type-windowquery)

#### <a id="method-appcontroller-focus"></a><span class="api-method">Method</span> `AppController.Focus`

- Signature: <span class="api-signature">`func (c *AppController) Focus(name string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-focusedwindow"></a><span class="api-method">Method</span> `AppController.FocusedWindow`

- Signature: <span class="api-signature">`func (c *AppController) FocusedWindow(name string, opts AppOptions) (Window, bool, error)`</span>
- Uses: [`AppOptions`](#type-appoptions), [`Window`](#type-window)

#### <a id="method-appcontroller-getwindow"></a><span class="api-method">Method</span> `AppController.GetWindow`

- Signature: <span class="api-signature">`func (c *AppController) GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error)`</span>
- Uses: [`AppOptions`](#type-appoptions), [`Window`](#type-window), [`WindowQuery`](#type-windowquery)

#### <a id="method-appcontroller-isrunning"></a><span class="api-method">Method</span> `AppController.IsRunning`

- Signature: <span class="api-signature">`func (c *AppController) IsRunning(name string, opts AppOptions) (bool, error)`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-listwindows"></a><span class="api-method">Method</span> `AppController.ListWindows`

- Signature: <span class="api-signature">`func (c *AppController) ListWindows(name string, opts AppOptions) ([]Window, error)`</span>
- Uses: [`AppOptions`](#type-appoptions), [`Window`](#type-window)

#### <a id="method-appcontroller-open"></a><span class="api-method">Method</span> `AppController.Open`

- Signature: <span class="api-signature">`func (c *AppController) Open(name string, args []string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-setbackend"></a><span class="api-method">Method</span> `AppController.SetBackend`

- Signature: <span class="api-signature">`func (c *AppController) SetBackend(backend core.App)`</span>

#### <a id="method-finder-collectlines"></a><span class="api-method">Method</span> `Finder.CollectLines`

- Signature: <span class="api-signature">`func (f *Finder) CollectLines(params OCRParams) ([]OCRLine, error)`</span>
- Uses: [`OCRLine`](#type-ocrline), [`OCRParams`](#type-ocrparams)
- Notes: CollectLines runs OCR and groups word-level geometry into line-level results.

#### <a id="method-finder-collectwords"></a><span class="api-method">Method</span> `Finder.CollectWords`

- Signature: <span class="api-signature">`func (f *Finder) CollectWords(params OCRParams) ([]OCRWord, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`OCRWord`](#type-ocrword)
- Notes: CollectWords runs OCR and returns word-level geometry for the source image.

#### <a id="method-finder-count"></a><span class="api-method">Method</span> `Finder.Count`

- Signature: <span class="api-signature">`func (f *Finder) Count(pattern *Pattern) (int, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Count returns the number of matches for the given pattern.

#### <a id="method-finder-destroy"></a><span class="api-method">Method</span> `Finder.Destroy`

- Signature: <span class="api-signature">`func (f *Finder) Destroy()`</span>
- Notes: Destroy clears the compatibility iterator state and last-match cache.

#### <a id="method-finder-exists"></a><span class="api-method">Method</span> `Finder.Exists`

- Signature: <span class="api-signature">`func (f *Finder) Exists(pattern *Pattern) (Match, bool, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: Exists returns the first match when present. Missing targets return (Match{}, false, nil).

#### <a id="method-finder-find"></a><span class="api-method">Method</span> `Finder.Find`

- Signature: <span class="api-signature">`func (f *Finder) Find(pattern *Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: Find returns the best match for the pattern.

#### <a id="method-finder-findall"></a><span class="api-method">Method</span> `Finder.FindAll`

- Signature: <span class="api-signature">`func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: FindAll returns all candidate matches for the pattern.

#### <a id="method-finder-findallbycolumn"></a><span class="api-method">Method</span> `Finder.FindAllByColumn`

- Signature: <span class="api-signature">`func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: FindAllByColumn returns all matches sorted left-to-right then top-to-bottom.

#### <a id="method-finder-findallbyrow"></a><span class="api-method">Method</span> `Finder.FindAllByRow`

- Signature: <span class="api-signature">`func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-findanylist"></a><span class="api-method">Method</span> `Finder.FindAnyList`

- Signature: <span class="api-signature">`func (f *Finder) FindAnyList(patterns []*Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-findbestlist"></a><span class="api-method">Method</span> `Finder.FindBestList`

- Signature: <span class="api-signature">`func (f *Finder) FindBestList(patterns []*Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-findtext"></a><span class="api-method">Method</span> `Finder.FindText`

- Signature: <span class="api-signature">`func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)
- Notes: FindText runs OCR and returns word-level matches for the query string.

#### <a id="method-finder-has"></a><span class="api-method">Method</span> `Finder.Has`

- Signature: <span class="api-signature">`func (f *Finder) Has(pattern *Pattern) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Has reports whether the target exists and bubbles non-find errors.

#### <a id="method-finder-hasnext"></a><span class="api-method">Method</span> `Finder.HasNext`

- Signature: <span class="api-signature">`func (f *Finder) HasNext() bool`</span>
- Notes: HasNext reports whether the compatibility iterator has another match available.

#### <a id="method-finder-iterate"></a><span class="api-method">Method</span> `Finder.Iterate`

- Signature: <span class="api-signature">`func (f *Finder) Iterate(pattern *Pattern) error`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Iterate prepares a compatibility iterator over the best match. Unlike Find, a miss does not return ErrFindFailed. Call HasNext to inspect presence.

#### <a id="method-finder-iterateall"></a><span class="api-method">Method</span> `Finder.IterateAll`

- Signature: <span class="api-signature">`func (f *Finder) IterateAll(pattern *Pattern) error`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: IterateAll prepares a compatibility iterator over all matches. Unlike Java SikuliX this additive Go surface keeps LastMatches available even after iteration.

#### <a id="method-finder-lastmatches"></a><span class="api-method">Method</span> `Finder.LastMatches`

- Signature: <span class="api-signature">`func (f *Finder) LastMatches() []Match`</span>
- Uses: [`Match`](#type-match)
- Notes: LastMatches returns a copy of the full most recent match set. It does not shrink as the compatibility iterator advances.

#### <a id="method-finder-next"></a><span class="api-method">Method</span> `Finder.Next`

- Signature: <span class="api-signature">`func (f *Finder) Next() (Match, bool)`</span>
- Uses: [`Match`](#type-match)
- Notes: Next returns the next compatibility-iterator match and advances the iterator. It returns false when the iterator is empty or exhausted.

#### <a id="method-finder-readtext"></a><span class="api-method">Method</span> `Finder.ReadText`

- Signature: <span class="api-signature">`func (f *Finder) ReadText(params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams)
- Notes: ReadText runs OCR and returns normalized text.

#### <a id="method-finder-reset"></a><span class="api-method">Method</span> `Finder.Reset`

- Signature: <span class="api-signature">`func (f *Finder) Reset()`</span>
- Notes: Reset rewinds the compatibility iterator to the start of the most recent match set.

#### <a id="method-finder-setmatcher"></a><span class="api-method">Method</span> `Finder.SetMatcher`

- Signature: <span class="api-signature">`func (f *Finder) SetMatcher(m core.Matcher)`</span>
- Notes: SetMatcher overrides the matcher backend used by this finder.

#### <a id="method-finder-setocrbackend"></a><span class="api-method">Method</span> `Finder.SetOCRBackend`

- Signature: <span class="api-signature">`func (f *Finder) SetOCRBackend(ocr core.OCR)`</span>
- Notes: SetOCRBackend overrides the OCR backend used by this finder.

#### <a id="method-finder-wait"></a><span class="api-method">Method</span> `Finder.Wait`

- Signature: <span class="api-signature">`func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-waitanylist"></a><span class="api-method">Method</span> `Finder.WaitAnyList`

- Signature: <span class="api-signature">`func (f *Finder) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-waitbestlist"></a><span class="api-method">Method</span> `Finder.WaitBestList`

- Signature: <span class="api-signature">`func (f *Finder) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-finder-waitvanish"></a><span class="api-method">Method</span> `Finder.WaitVanish`

- Signature: <span class="api-signature">`func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: WaitVanish blocks until the pattern disappears or timeout expires.

#### <a id="method-image-clone"></a><span class="api-method">Method</span> `Image.Clone`

- Signature: <span class="api-signature">`func (i *Image) Clone() *Image`</span>

#### <a id="method-image-crop"></a><span class="api-method">Method</span> `Image.Crop`

- Signature: <span class="api-signature">`func (i *Image) Crop(rect Rect) (*Image, error)`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-image-gray"></a><span class="api-method">Method</span> `Image.Gray`

- Signature: <span class="api-signature">`func (i *Image) Gray() *image.Gray`</span>

#### <a id="method-image-height"></a><span class="api-method">Method</span> `Image.Height`

- Signature: <span class="api-signature">`func (i *Image) Height() int`</span>

#### <a id="method-image-name"></a><span class="api-method">Method</span> `Image.Name`

- Signature: <span class="api-signature">`func (i *Image) Name() string`</span>

#### <a id="method-image-width"></a><span class="api-method">Method</span> `Image.Width`

- Signature: <span class="api-signature">`func (i *Image) Width() int`</span>

#### <a id="method-inputcontroller-click"></a><span class="api-method">Method</span> `InputController.Click`

- Signature: <span class="api-signature">`func (c *InputController) Click(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-doubleclick"></a><span class="api-method">Method</span> `InputController.DoubleClick`

- Signature: <span class="api-signature">`func (c *InputController) DoubleClick(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-dragdrop"></a><span class="api-method">Method</span> `InputController.DragDrop`

- Signature: <span class="api-signature">`func (c *InputController) DragDrop(fromX, fromY, toX, toY int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-hotkey"></a><span class="api-method">Method</span> `InputController.Hotkey`

- Signature: <span class="api-signature">`func (c *InputController) Hotkey(keys ...string) error`</span>

#### <a id="method-inputcontroller-hover"></a><span class="api-method">Method</span> `InputController.Hover`

- Signature: <span class="api-signature">`func (c *InputController) Hover(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-keydown"></a><span class="api-method">Method</span> `InputController.KeyDown`

- Signature: <span class="api-signature">`func (c *InputController) KeyDown(keys ...string) error`</span>

#### <a id="method-inputcontroller-keyup"></a><span class="api-method">Method</span> `InputController.KeyUp`

- Signature: <span class="api-signature">`func (c *InputController) KeyUp(keys ...string) error`</span>

#### <a id="method-inputcontroller-mousedown"></a><span class="api-method">Method</span> `InputController.MouseDown`

- Signature: <span class="api-signature">`func (c *InputController) MouseDown(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-mouseup"></a><span class="api-method">Method</span> `InputController.MouseUp`

- Signature: <span class="api-signature">`func (c *InputController) MouseUp(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-movemouse"></a><span class="api-method">Method</span> `InputController.MoveMouse`

- Signature: <span class="api-signature">`func (c *InputController) MoveMouse(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-paste"></a><span class="api-method">Method</span> `InputController.Paste`

- Signature: <span class="api-signature">`func (c *InputController) Paste(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-rightclick"></a><span class="api-method">Method</span> `InputController.RightClick`

- Signature: <span class="api-signature">`func (c *InputController) RightClick(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-setbackend"></a><span class="api-method">Method</span> `InputController.SetBackend`

- Signature: <span class="api-signature">`func (c *InputController) SetBackend(backend core.Input)`</span>

#### <a id="method-inputcontroller-typetext"></a><span class="api-method">Method</span> `InputController.TypeText`

- Signature: <span class="api-signature">`func (c *InputController) TypeText(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-wheel"></a><span class="api-method">Method</span> `InputController.Wheel`

- Signature: <span class="api-signature">`func (c *InputController) Wheel(x, y int, direction WheelDirection, steps int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`WheelDirection`](#type-wheeldirection)

#### <a id="method-liveregion-bounds"></a><span class="api-method">Method</span> `LiveRegion.Bounds`

- Signature: <span class="api-signature">`func (r LiveRegion) Bounds() Region`</span>
- Uses: [`Region`](#type-region)

#### <a id="method-liveregion-capture"></a><span class="api-method">Method</span> `LiveRegion.Capture`

- Signature: <span class="api-signature">`func (r LiveRegion) Capture() (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="method-liveregion-center"></a><span class="api-method">Method</span> `LiveRegion.Center`

- Signature: <span class="api-signature">`func (r LiveRegion) Center() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-liveregion-click"></a><span class="api-method">Method</span> `LiveRegion.Click`

- Signature: <span class="api-signature">`func (r LiveRegion) Click(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-collectlines"></a><span class="api-method">Method</span> `LiveRegion.CollectLines`

- Signature: <span class="api-signature">`func (r LiveRegion) CollectLines(params OCRParams) ([]OCRLine, error)`</span>
- Uses: [`OCRLine`](#type-ocrline), [`OCRParams`](#type-ocrparams)

#### <a id="method-liveregion-collectwords"></a><span class="api-method">Method</span> `LiveRegion.CollectWords`

- Signature: <span class="api-signature">`func (r LiveRegion) CollectWords(params OCRParams) ([]OCRWord, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`OCRWord`](#type-ocrword)

#### <a id="method-liveregion-doubleclick"></a><span class="api-method">Method</span> `LiveRegion.DoubleClick`

- Signature: <span class="api-signature">`func (r LiveRegion) DoubleClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-dragdrop"></a><span class="api-method">Method</span> `LiveRegion.DragDrop`

- Signature: <span class="api-signature">`func (r LiveRegion) DragDrop(target TargetPointProvider, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`TargetPointProvider`](#type-targetpointprovider)

#### <a id="method-liveregion-exists"></a><span class="api-method">Method</span> `LiveRegion.Exists`

- Signature: <span class="api-signature">`func (r LiveRegion) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-find"></a><span class="api-method">Method</span> `LiveRegion.Find`

- Signature: <span class="api-signature">`func (r LiveRegion) Find(pattern *Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-findanylist"></a><span class="api-method">Method</span> `LiveRegion.FindAnyList`

- Signature: <span class="api-signature">`func (r LiveRegion) FindAnyList(patterns []*Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-findbestlist"></a><span class="api-method">Method</span> `LiveRegion.FindBestList`

- Signature: <span class="api-signature">`func (r LiveRegion) FindBestList(patterns []*Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-findtext"></a><span class="api-method">Method</span> `LiveRegion.FindText`

- Signature: <span class="api-signature">`func (r LiveRegion) FindText(query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)

#### <a id="method-liveregion-grow"></a><span class="api-method">Method</span> `LiveRegion.Grow`

- Signature: <span class="api-signature">`func (r LiveRegion) Grow(dx, dy int) LiveRegion`</span>

#### <a id="method-liveregion-has"></a><span class="api-method">Method</span> `LiveRegion.Has`

- Signature: <span class="api-signature">`func (r LiveRegion) Has(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-liveregion-hover"></a><span class="api-method">Method</span> `LiveRegion.Hover`

- Signature: <span class="api-signature">`func (r LiveRegion) Hover(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-keydown"></a><span class="api-method">Method</span> `LiveRegion.KeyDown`

- Signature: <span class="api-signature">`func (r LiveRegion) KeyDown(keys ...string) error`</span>

#### <a id="method-liveregion-keyup"></a><span class="api-method">Method</span> `LiveRegion.KeyUp`

- Signature: <span class="api-signature">`func (r LiveRegion) KeyUp(keys ...string) error`</span>

#### <a id="method-liveregion-mousedown"></a><span class="api-method">Method</span> `LiveRegion.MouseDown`

- Signature: <span class="api-signature">`func (r LiveRegion) MouseDown(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-mouseup"></a><span class="api-method">Method</span> `LiveRegion.MouseUp`

- Signature: <span class="api-signature">`func (r LiveRegion) MouseUp(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-moveto"></a><span class="api-method">Method</span> `LiveRegion.MoveTo`

- Signature: <span class="api-signature">`func (r LiveRegion) MoveTo(x, y int) LiveRegion`</span>

#### <a id="method-liveregion-offset"></a><span class="api-method">Method</span> `LiveRegion.Offset`

- Signature: <span class="api-signature">`func (r LiveRegion) Offset(dx, dy int) LiveRegion`</span>
- Uses: [`Offset`](#type-offset)

#### <a id="method-liveregion-paste"></a><span class="api-method">Method</span> `LiveRegion.Paste`

- Signature: <span class="api-signature">`func (r LiveRegion) Paste(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-readtext"></a><span class="api-method">Method</span> `LiveRegion.ReadText`

- Signature: <span class="api-signature">`func (r LiveRegion) ReadText(params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams)

#### <a id="method-liveregion-rightclick"></a><span class="api-method">Method</span> `LiveRegion.RightClick`

- Signature: <span class="api-signature">`func (r LiveRegion) RightClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-setsize"></a><span class="api-method">Method</span> `LiveRegion.SetSize`

- Signature: <span class="api-signature">`func (r LiveRegion) SetSize(w, h int) LiveRegion`</span>

#### <a id="method-liveregion-targetpoint"></a><span class="api-method">Method</span> `LiveRegion.TargetPoint`

- Signature: <span class="api-signature">`func (r LiveRegion) TargetPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-liveregion-typetext"></a><span class="api-method">Method</span> `LiveRegion.TypeText`

- Signature: <span class="api-signature">`func (r LiveRegion) TypeText(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-liveregion-wait"></a><span class="api-method">Method</span> `LiveRegion.Wait`

- Signature: <span class="api-signature">`func (r LiveRegion) Wait(pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-waitanylist"></a><span class="api-method">Method</span> `LiveRegion.WaitAnyList`

- Signature: <span class="api-signature">`func (r LiveRegion) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-waitbestlist"></a><span class="api-method">Method</span> `LiveRegion.WaitBestList`

- Signature: <span class="api-signature">`func (r LiveRegion) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-liveregion-waitvanish"></a><span class="api-method">Method</span> `LiveRegion.WaitVanish`

- Signature: <span class="api-signature">`func (r LiveRegion) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-liveregion-wheel"></a><span class="api-method">Method</span> `LiveRegion.Wheel`

- Signature: <span class="api-signature">`func (r LiveRegion) Wheel(direction WheelDirection, steps int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`WheelDirection`](#type-wheeldirection)

#### <a id="method-liveregion-withmatcherengine"></a><span class="api-method">Method</span> `LiveRegion.WithMatcherEngine`

- Signature: <span class="api-signature">`func (r LiveRegion) WithMatcherEngine(engine MatcherEngine) LiveRegion`</span>
- Uses: [`MatcherEngine`](#type-matcherengine)

#### <a id="method-location-move"></a><span class="api-method">Method</span> `Location.Move`

- Signature: <span class="api-signature">`func (l Location) Move(dx, dy int) Location`</span>

#### <a id="method-location-string"></a><span class="api-method">Method</span> `Location.String`

- Signature: <span class="api-signature">`func (l Location) String() string`</span>

#### <a id="method-location-targetpoint"></a><span class="api-method">Method</span> `Location.TargetPoint`

- Signature: <span class="api-signature">`func (l Location) TargetPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-location-topoint"></a><span class="api-method">Method</span> `Location.ToPoint`

- Signature: <span class="api-signature">`func (l Location) ToPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-match-bounds"></a><span class="api-method">Method</span> `Match.Bounds`

- Signature: <span class="api-signature">`func (m Match) Bounds() Region`</span>
- Uses: [`Region`](#type-region)
- Notes: Bounds returns the match as a region-like value so callers can treat a live match the same way SikuliX documents region-capable match results.

#### <a id="method-match-capture"></a><span class="api-method">Method</span> `Match.Capture`

- Signature: <span class="api-signature">`func (m Match) Capture() (*Image, error)`</span>
- Uses: [`Image`](#type-image)
- Notes: Capture captures the currently matched live region.

#### <a id="method-match-center"></a><span class="api-method">Method</span> `Match.Center`

- Signature: <span class="api-signature">`func (m Match) Center() Point`</span>
- Uses: [`Point`](#type-point)
- Notes: Center returns the geometric center of the matched rectangle.

#### <a id="method-match-click"></a><span class="api-method">Method</span> `Match.Click`

- Signature: <span class="api-signature">`func (m Match) Click(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: Click clicks the match target point.

#### <a id="method-match-collectlines"></a><span class="api-method">Method</span> `Match.CollectLines`

- Signature: <span class="api-signature">`func (m Match) CollectLines(params OCRParams) ([]OCRLine, error)`</span>
- Uses: [`OCRLine`](#type-ocrline), [`OCRParams`](#type-ocrparams)
- Notes: CollectLines returns OCR line-level results inside the matched live region.

#### <a id="method-match-collectwords"></a><span class="api-method">Method</span> `Match.CollectWords`

- Signature: <span class="api-signature">`func (m Match) CollectWords(params OCRParams) ([]OCRWord, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`OCRWord`](#type-ocrword)
- Notes: CollectWords returns OCR word-level results inside the matched live region.

#### <a id="method-match-doubleclick"></a><span class="api-method">Method</span> `Match.DoubleClick`

- Signature: <span class="api-signature">`func (m Match) DoubleClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: DoubleClick performs two click actions against the match target point.

#### <a id="method-match-dragdrop"></a><span class="api-method">Method</span> `Match.DragDrop`

- Signature: <span class="api-signature">`func (m Match) DragDrop(target TargetPointProvider, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`TargetPointProvider`](#type-targetpointprovider)
- Notes: DragDrop drags from the match target point to the target point.

#### <a id="method-match-exists"></a><span class="api-method">Method</span> `Match.Exists`

- Signature: <span class="api-signature">`func (m Match) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Exists probes within the matched live region.

#### <a id="method-match-find"></a><span class="api-method">Method</span> `Match.Find`

- Signature: <span class="api-signature">`func (m Match) Find(pattern *Pattern) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Find searches again within the matched live region.

#### <a id="method-match-findanylist"></a><span class="api-method">Method</span> `Match.FindAnyList`

- Signature: <span class="api-signature">`func (m Match) FindAnyList(patterns []*Pattern) ([]Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-match-findbestlist"></a><span class="api-method">Method</span> `Match.FindBestList`

- Signature: <span class="api-signature">`func (m Match) FindBestList(patterns []*Pattern) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-match-findtext"></a><span class="api-method">Method</span> `Match.FindText`

- Signature: <span class="api-signature">`func (m Match) FindText(query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)
- Notes: FindText searches OCR text inside the matched live region.

#### <a id="method-match-has"></a><span class="api-method">Method</span> `Match.Has`

- Signature: <span class="api-signature">`func (m Match) Has(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Has reports whether the target exists within the matched live region.

#### <a id="method-match-hover"></a><span class="api-method">Method</span> `Match.Hover`

- Signature: <span class="api-signature">`func (m Match) Hover(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: Hover is a parity-friendly alias for MoveMouse.

#### <a id="method-match-keydown"></a><span class="api-method">Method</span> `Match.KeyDown`

- Signature: <span class="api-signature">`func (m Match) KeyDown(keys ...string) error`</span>
- Notes: KeyDown holds the provided keys.

#### <a id="method-match-keyup"></a><span class="api-method">Method</span> `Match.KeyUp`

- Signature: <span class="api-signature">`func (m Match) KeyUp(keys ...string) error`</span>
- Notes: KeyUp releases the provided keys.

#### <a id="method-match-live"></a><span class="api-method">Method</span> `Match.Live`

- Signature: <span class="api-signature">`func (m Match) Live() bool`</span>
- Notes: Live reports whether this match is bound to a running sikuli-go runtime.

#### <a id="method-match-mousedown"></a><span class="api-method">Method</span> `Match.MouseDown`

- Signature: <span class="api-signature">`func (m Match) MouseDown(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: MouseDown presses and holds the mouse button at the match target point.

#### <a id="method-match-mouseup"></a><span class="api-method">Method</span> `Match.MouseUp`

- Signature: <span class="api-signature">`func (m Match) MouseUp(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: MouseUp releases the mouse button at the match target point.

#### <a id="method-match-movemouse"></a><span class="api-method">Method</span> `Match.MoveMouse`

- Signature: <span class="api-signature">`func (m Match) MoveMouse(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: MoveMouse moves the pointer to the match target point.

#### <a id="method-match-paste"></a><span class="api-method">Method</span> `Match.Paste`

- Signature: <span class="api-signature">`func (m Match) Paste(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: Paste focuses the match target point and pastes text into it.

#### <a id="method-match-readtext"></a><span class="api-method">Method</span> `Match.ReadText`

- Signature: <span class="api-signature">`func (m Match) ReadText(params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams)
- Notes: ReadText runs OCR inside the matched live region.

#### <a id="method-match-region"></a><span class="api-method">Method</span> `Match.Region`

- Signature: <span class="api-signature">`func (m Match) Region() Region`</span>
- Uses: [`Region`](#type-region)
- Notes: Region is an alias for Bounds for parity-friendly call sites.

#### <a id="method-match-rightclick"></a><span class="api-method">Method</span> `Match.RightClick`

- Signature: <span class="api-signature">`func (m Match) RightClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: RightClick clicks the match target point with the right mouse button.

#### <a id="method-match-string"></a><span class="api-method">Method</span> `Match.String`

- Signature: <span class="api-signature">`func (m Match) String() string`</span>

#### <a id="method-match-targetpoint"></a><span class="api-method">Method</span> `Match.TargetPoint`

- Signature: <span class="api-signature">`func (m Match) TargetPoint() Point`</span>
- Uses: [`Point`](#type-point)
- Notes: TargetPoint returns the resolved click target point for this match.

#### <a id="method-match-typetext"></a><span class="api-method">Method</span> `Match.TypeText`

- Signature: <span class="api-signature">`func (m Match) TypeText(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)
- Notes: TypeText focuses the match target point and types text into it.

#### <a id="method-match-wait"></a><span class="api-method">Method</span> `Match.Wait`

- Signature: <span class="api-signature">`func (m Match) Wait(pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Wait waits for the target within the matched live region.

#### <a id="method-match-waitanylist"></a><span class="api-method">Method</span> `Match.WaitAnyList`

- Signature: <span class="api-signature">`func (m Match) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-match-waitbestlist"></a><span class="api-method">Method</span> `Match.WaitBestList`

- Signature: <span class="api-signature">`func (m Match) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-match-waitvanish"></a><span class="api-method">Method</span> `Match.WaitVanish`

- Signature: <span class="api-signature">`func (m Match) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: WaitVanish waits for the target to disappear from the matched live region.

#### <a id="method-match-wheel"></a><span class="api-method">Method</span> `Match.Wheel`

- Signature: <span class="api-signature">`func (m Match) Wheel(direction WheelDirection, steps int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`WheelDirection`](#type-wheeldirection)
- Notes: Wheel scrolls at the match target point.

#### <a id="method-observercontroller-observeappear"></a><span class="api-method">Method</span> `ObserverController.ObserveAppear`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`ObserveEvent`](#type-observeevent), [`ObserveOptions`](#type-observeoptions), [`Pattern`](#type-pattern), [`Region`](#type-region)

#### <a id="method-observercontroller-observechange"></a><span class="api-method">Method</span> `ObserverController.ObserveChange`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`ObserveEvent`](#type-observeevent), [`ObserveOptions`](#type-observeoptions), [`Region`](#type-region)

#### <a id="method-observercontroller-observevanish"></a><span class="api-method">Method</span> `ObserverController.ObserveVanish`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`ObserveEvent`](#type-observeevent), [`ObserveOptions`](#type-observeoptions), [`Pattern`](#type-pattern), [`Region`](#type-region)

#### <a id="method-observercontroller-setbackend"></a><span class="api-method">Method</span> `ObserverController.SetBackend`

- Signature: <span class="api-signature">`func (c *ObserverController) SetBackend(backend core.Observer)`</span>

#### <a id="method-offset-string"></a><span class="api-method">Method</span> `Offset.String`

- Signature: <span class="api-signature">`func (o Offset) String() string`</span>

#### <a id="method-offset-topoint"></a><span class="api-method">Method</span> `Offset.ToPoint`

- Signature: <span class="api-signature">`func (o Offset) ToPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-options-clone"></a><span class="api-method">Method</span> `Options.Clone`

- Signature: <span class="api-signature">`func (o *Options) Clone() *Options`</span>

#### <a id="method-options-delete"></a><span class="api-method">Method</span> `Options.Delete`

- Signature: <span class="api-signature">`func (o *Options) Delete(key string)`</span>

#### <a id="method-options-entries"></a><span class="api-method">Method</span> `Options.Entries`

- Signature: <span class="api-signature">`func (o *Options) Entries() map[string]string`</span>

#### <a id="method-options-getbool"></a><span class="api-method">Method</span> `Options.GetBool`

- Signature: <span class="api-signature">`func (o *Options) GetBool(key string, def bool) bool`</span>

#### <a id="method-options-getfloat64"></a><span class="api-method">Method</span> `Options.GetFloat64`

- Signature: <span class="api-signature">`func (o *Options) GetFloat64(key string, def float64) float64`</span>

#### <a id="method-options-getint"></a><span class="api-method">Method</span> `Options.GetInt`

- Signature: <span class="api-signature">`func (o *Options) GetInt(key string, def int) int`</span>

#### <a id="method-options-getstring"></a><span class="api-method">Method</span> `Options.GetString`

- Signature: <span class="api-signature">`func (o *Options) GetString(key, def string) string`</span>

#### <a id="method-options-has"></a><span class="api-method">Method</span> `Options.Has`

- Signature: <span class="api-signature">`func (o *Options) Has(key string) bool`</span>

#### <a id="method-options-merge"></a><span class="api-method">Method</span> `Options.Merge`

- Signature: <span class="api-signature">`func (o *Options) Merge(other *Options)`</span>

#### <a id="method-options-setbool"></a><span class="api-method">Method</span> `Options.SetBool`

- Signature: <span class="api-signature">`func (o *Options) SetBool(key string, value bool)`</span>

#### <a id="method-options-setfloat64"></a><span class="api-method">Method</span> `Options.SetFloat64`

- Signature: <span class="api-signature">`func (o *Options) SetFloat64(key string, value float64)`</span>

#### <a id="method-options-setint"></a><span class="api-method">Method</span> `Options.SetInt`

- Signature: <span class="api-signature">`func (o *Options) SetInt(key string, value int)`</span>

#### <a id="method-options-setstring"></a><span class="api-method">Method</span> `Options.SetString`

- Signature: <span class="api-signature">`func (o *Options) SetString(key, value string)`</span>

#### <a id="method-pattern-exact"></a><span class="api-method">Method</span> `Pattern.Exact`

- Signature: <span class="api-signature">`func (p *Pattern) Exact() *Pattern`</span>
- Notes: Exact is a convenience for Similar(1.0).

#### <a id="method-pattern-image"></a><span class="api-method">Method</span> `Pattern.Image`

- Signature: <span class="api-signature">`func (p *Pattern) Image() *Image`</span>
- Uses: [`Image`](#type-image)
- Notes: Image returns the underlying pattern image.

#### <a id="method-pattern-mask"></a><span class="api-method">Method</span> `Pattern.Mask`

- Signature: <span class="api-signature">`func (p *Pattern) Mask() *image.Gray`</span>
- Notes: Mask returns the currently configured mask.

#### <a id="method-pattern-offset"></a><span class="api-method">Method</span> `Pattern.Offset`

- Signature: <span class="api-signature">`func (p *Pattern) Offset() Point`</span>
- Uses: [`Offset`](#type-offset), [`Point`](#type-point)
- Notes: Offset returns the configured click anchor offset.

#### <a id="method-pattern-resize"></a><span class="api-method">Method</span> `Pattern.Resize`

- Signature: <span class="api-signature">`func (p *Pattern) Resize(factor float64) *Pattern`</span>
- Notes: Resize scales the pattern before matching.

#### <a id="method-pattern-resizefactor"></a><span class="api-method">Method</span> `Pattern.ResizeFactor`

- Signature: <span class="api-signature">`func (p *Pattern) ResizeFactor() float64`</span>
- Notes: ResizeFactor returns the currently configured resize factor.

#### <a id="method-pattern-similar"></a><span class="api-method">Method</span> `Pattern.Similar`

- Signature: <span class="api-signature">`func (p *Pattern) Similar(sim float64) *Pattern`</span>
- Notes: Similar sets the acceptance threshold in the [0,1] range. Higher values require a closer match.

#### <a id="method-pattern-similarity"></a><span class="api-method">Method</span> `Pattern.Similarity`

- Signature: <span class="api-signature">`func (p *Pattern) Similarity() float64`</span>
- Notes: Similarity returns the current acceptance threshold.

#### <a id="method-pattern-targetoffset"></a><span class="api-method">Method</span> `Pattern.TargetOffset`

- Signature: <span class="api-signature">`func (p *Pattern) TargetOffset(dx, dy int) *Pattern`</span>
- Notes: TargetOffset sets the click anchor relative to the matched rectangle.

#### <a id="method-pattern-withmask"></a><span class="api-method">Method</span> `Pattern.WithMask`

- Signature: <span class="api-signature">`func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error)`</span>
- Notes: WithMask sets an optional per-pixel mask where 0 excludes and >0 includes pixels.

#### <a id="method-pattern-withmaskmatrix"></a><span class="api-method">Method</span> `Pattern.WithMaskMatrix`

- Signature: <span class="api-signature">`func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error)`</span>
- Notes: WithMaskMatrix sets an optional binary mask from matrix rows.

#### <a id="method-point-targetpoint"></a><span class="api-method">Method</span> `Point.TargetPoint`

- Signature: <span class="api-signature">`func (p Point) TargetPoint() Point`</span>

#### <a id="method-point-tolocation"></a><span class="api-method">Method</span> `Point.ToLocation`

- Signature: <span class="api-signature">`func (p Point) ToLocation() Location`</span>
- Uses: [`Location`](#type-location)
- Notes: ToLocation converts a point to a parity-friendly Location value.

#### <a id="method-point-tooffset"></a><span class="api-method">Method</span> `Point.ToOffset`

- Signature: <span class="api-signature">`func (p Point) ToOffset() Offset`</span>
- Uses: [`Offset`](#type-offset)
- Notes: ToOffset converts a point to a parity-friendly Offset value.

#### <a id="method-rect-contains"></a><span class="api-method">Method</span> `Rect.Contains`

- Signature: <span class="api-signature">`func (r Rect) Contains(p Point) bool`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-rect-empty"></a><span class="api-method">Method</span> `Rect.Empty`

- Signature: <span class="api-signature">`func (r Rect) Empty() bool`</span>

#### <a id="method-rect-string"></a><span class="api-method">Method</span> `Rect.String`

- Signature: <span class="api-signature">`func (r Rect) String() string`</span>

#### <a id="method-region-center"></a><span class="api-method">Method</span> `Region.Center`

- Signature: <span class="api-signature">`func (r Region) Center() Point`</span>
- Uses: [`Point`](#type-point)
- Notes: Center returns the midpoint of the region.

#### <a id="method-region-collectlines"></a><span class="api-method">Method</span> `Region.CollectLines`

- Signature: <span class="api-signature">`func (r Region) CollectLines(source *Image, params OCRParams) ([]OCRLine, error)`</span>
- Uses: [`Image`](#type-image), [`OCRLine`](#type-ocrline), [`OCRParams`](#type-ocrparams)
- Notes: CollectLines runs OCR in region and returns line-level results.

#### <a id="method-region-collectwords"></a><span class="api-method">Method</span> `Region.CollectWords`

- Signature: <span class="api-signature">`func (r Region) CollectWords(source *Image, params OCRParams) ([]OCRWord, error)`</span>
- Uses: [`Image`](#type-image), [`OCRParams`](#type-ocrparams), [`OCRWord`](#type-ocrword)
- Notes: CollectWords runs OCR in region and returns word-level results.

#### <a id="method-region-contains"></a><span class="api-method">Method</span> `Region.Contains`

- Signature: <span class="api-signature">`func (r Region) Contains(p Point) bool`</span>
- Uses: [`Point`](#type-point)
- Notes: Contains reports whether a point lies within the region.

#### <a id="method-region-containslocation"></a><span class="api-method">Method</span> `Region.ContainsLocation`

- Signature: <span class="api-signature">`func (r Region) ContainsLocation(loc Location) bool`</span>
- Uses: [`Location`](#type-location)
- Notes: ContainsLocation reports whether this region contains the given location.

#### <a id="method-region-containsregion"></a><span class="api-method">Method</span> `Region.ContainsRegion`

- Signature: <span class="api-signature">`func (r Region) ContainsRegion(other Region) bool`</span>

#### <a id="method-region-count"></a><span class="api-method">Method</span> `Region.Count`

- Signature: <span class="api-signature">`func (r Region) Count(source *Image, pattern *Pattern) (int, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)
- Notes: Count returns the number of matches found in this region.

#### <a id="method-region-exists"></a><span class="api-method">Method</span> `Region.Exists`

- Signature: <span class="api-signature">`func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: Exists checks for pattern presence within timeout and returns first match when found.

#### <a id="method-region-find"></a><span class="api-method">Method</span> `Region.Find`

- Signature: <span class="api-signature">`func (r Region) Find(source *Image, pattern *Pattern) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-findall"></a><span class="api-method">Method</span> `Region.FindAll`

- Signature: <span class="api-signature">`func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)
- Notes: FindAll returns all matches in this region.

#### <a id="method-region-findallbycolumn"></a><span class="api-method">Method</span> `Region.FindAllByColumn`

- Signature: <span class="api-signature">`func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-findallbyrow"></a><span class="api-method">Method</span> `Region.FindAllByRow`

- Signature: <span class="api-signature">`func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-findanylist"></a><span class="api-method">Method</span> `Region.FindAnyList`

- Signature: <span class="api-signature">`func (r Region) FindAnyList(source *Image, patterns []*Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-findbestlist"></a><span class="api-method">Method</span> `Region.FindBestList`

- Signature: <span class="api-signature">`func (r Region) FindBestList(source *Image, patterns []*Pattern) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-findtext"></a><span class="api-method">Method</span> `Region.FindText`

- Signature: <span class="api-signature">`func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`Image`](#type-image), [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)
- Notes: FindText runs OCR in region and returns matches for the query.

#### <a id="method-region-grow"></a><span class="api-method">Method</span> `Region.Grow`

- Signature: <span class="api-signature">`func (r Region) Grow(dx, dy int) Region`</span>
- Notes: Grow expands the region outward in both directions.

#### <a id="method-region-has"></a><span class="api-method">Method</span> `Region.Has`

- Signature: <span class="api-signature">`func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)

#### <a id="method-region-intersection"></a><span class="api-method">Method</span> `Region.Intersection`

- Signature: <span class="api-signature">`func (r Region) Intersection(other Region) Region`</span>
- Notes: Intersection returns the overlap between this region and another.

#### <a id="method-region-moveto"></a><span class="api-method">Method</span> `Region.MoveTo`

- Signature: <span class="api-signature">`func (r Region) MoveTo(x, y int) Region`</span>

#### <a id="method-region-movetolocation"></a><span class="api-method">Method</span> `Region.MoveToLocation`

- Signature: <span class="api-signature">`func (r Region) MoveToLocation(loc Location) Region`</span>
- Uses: [`Location`](#type-location)
- Notes: MoveToLocation moves this region using a Location alias.

#### <a id="method-region-offset"></a><span class="api-method">Method</span> `Region.Offset`

- Signature: <span class="api-signature">`func (r Region) Offset(dx, dy int) Region`</span>
- Uses: [`Offset`](#type-offset)
- Notes: Offset translates the region by dx and dy.

#### <a id="method-region-offsetby"></a><span class="api-method">Method</span> `Region.OffsetBy`

- Signature: <span class="api-signature">`func (r Region) OffsetBy(off Offset) Region`</span>
- Uses: [`Offset`](#type-offset)
- Notes: OffsetBy applies an Offset alias to this region position.

#### <a id="method-region-readtext"></a><span class="api-method">Method</span> `Region.ReadText`

- Signature: <span class="api-signature">`func (r Region) ReadText(source *Image, params OCRParams) (string, error)`</span>
- Uses: [`Image`](#type-image), [`OCRParams`](#type-ocrparams)

#### <a id="method-region-resetthrowexception"></a><span class="api-method">Method</span> `Region.ResetThrowException`

- Signature: <span class="api-signature">`func (r *Region) ResetThrowException()`</span>

#### <a id="method-region-setautowaittimeout"></a><span class="api-method">Method</span> `Region.SetAutoWaitTimeout`

- Signature: <span class="api-signature">`func (r *Region) SetAutoWaitTimeout(sec float64)`</span>

#### <a id="method-region-setobservescanrate"></a><span class="api-method">Method</span> `Region.SetObserveScanRate`

- Signature: <span class="api-signature">`func (r *Region) SetObserveScanRate(rate float64)`</span>

#### <a id="method-region-setsize"></a><span class="api-method">Method</span> `Region.SetSize`

- Signature: <span class="api-signature">`func (r Region) SetSize(w, h int) Region`</span>
- Notes: SetSize updates width and height while clamping negatives to zero.

#### <a id="method-region-setthrowexception"></a><span class="api-method">Method</span> `Region.SetThrowException`

- Signature: <span class="api-signature">`func (r *Region) SetThrowException(flag bool)`</span>

#### <a id="method-region-setwaitscanrate"></a><span class="api-method">Method</span> `Region.SetWaitScanRate`

- Signature: <span class="api-signature">`func (r *Region) SetWaitScanRate(rate float64)`</span>

#### <a id="method-region-targetpoint"></a><span class="api-method">Method</span> `Region.TargetPoint`

- Signature: <span class="api-signature">`func (r Region) TargetPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-region-union"></a><span class="api-method">Method</span> `Region.Union`

- Signature: <span class="api-signature">`func (r Region) Union(other Region) Region`</span>
- Notes: Union returns the smallest region containing both regions.

#### <a id="method-region-wait"></a><span class="api-method">Method</span> `Region.Wait`

- Signature: <span class="api-signature">`func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-waitanylist"></a><span class="api-method">Method</span> `Region.WaitAnyList`

- Signature: <span class="api-signature">`func (r Region) WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-waitbestlist"></a><span class="api-method">Method</span> `Region.WaitBestList`

- Signature: <span class="api-signature">`func (r Region) WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-region-waitvanish"></a><span class="api-method">Method</span> `Region.WaitVanish`

- Signature: <span class="api-signature">`func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)
- Notes: WaitVanish waits until pattern disappears or timeout expires.

#### <a id="method-runtime-address"></a><span class="api-method">Method</span> `Runtime.Address`

- Signature: <span class="api-signature">`func (r *Runtime) Address() string`</span>

#### <a id="method-runtime-capture"></a><span class="api-method">Method</span> `Runtime.Capture`

- Signature: <span class="api-signature">`func (r *Runtime) Capture() (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="method-runtime-captureregion"></a><span class="api-method">Method</span> `Runtime.CaptureRegion`

- Signature: <span class="api-signature">`func (r *Runtime) CaptureRegion(region Region) (*Image, error)`</span>
- Uses: [`Image`](#type-image), [`Region`](#type-region)

#### <a id="method-runtime-close"></a><span class="api-method">Method</span> `Runtime.Close`

- Signature: <span class="api-signature">`func (r *Runtime) Close() error`</span>

#### <a id="method-runtime-primaryscreen"></a><span class="api-method">Method</span> `Runtime.PrimaryScreen`

- Signature: <span class="api-signature">`func (r *Runtime) PrimaryScreen() (Screen, error)`</span>
- Uses: [`Screen`](#type-screen)

#### <a id="method-runtime-region"></a><span class="api-method">Method</span> `Runtime.Region`

- Signature: <span class="api-signature">`func (r *Runtime) Region(region Region) LiveRegion`</span>
- Uses: [`LiveRegion`](#type-liveregion), [`Region`](#type-region)

#### <a id="method-runtime-screen"></a><span class="api-method">Method</span> `Runtime.Screen`

- Signature: <span class="api-signature">`func (r *Runtime) Screen(id int) (Screen, error)`</span>
- Uses: [`Screen`](#type-screen)

#### <a id="method-runtime-screens"></a><span class="api-method">Method</span> `Runtime.Screens`

- Signature: <span class="api-signature">`func (r *Runtime) Screens() ([]Screen, error)`</span>
- Uses: [`Screen`](#type-screen)

#### <a id="method-screen-capture"></a><span class="api-method">Method</span> `Screen.Capture`

- Signature: <span class="api-signature">`func (s Screen) Capture() (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="method-screen-click"></a><span class="api-method">Method</span> `Screen.Click`

- Signature: <span class="api-signature">`func (s Screen) Click(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-collectlines"></a><span class="api-method">Method</span> `Screen.CollectLines`

- Signature: <span class="api-signature">`func (s Screen) CollectLines(params OCRParams) ([]OCRLine, error)`</span>
- Uses: [`OCRLine`](#type-ocrline), [`OCRParams`](#type-ocrparams)

#### <a id="method-screen-collectwords"></a><span class="api-method">Method</span> `Screen.CollectWords`

- Signature: <span class="api-signature">`func (s Screen) CollectWords(params OCRParams) ([]OCRWord, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`OCRWord`](#type-ocrword)

#### <a id="method-screen-doubleclick"></a><span class="api-method">Method</span> `Screen.DoubleClick`

- Signature: <span class="api-signature">`func (s Screen) DoubleClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-dragdrop"></a><span class="api-method">Method</span> `Screen.DragDrop`

- Signature: <span class="api-signature">`func (s Screen) DragDrop(target TargetPointProvider, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`TargetPointProvider`](#type-targetpointprovider)

#### <a id="method-screen-exists"></a><span class="api-method">Method</span> `Screen.Exists`

- Signature: <span class="api-signature">`func (s Screen) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-find"></a><span class="api-method">Method</span> `Screen.Find`

- Signature: <span class="api-signature">`func (s Screen) Find(pattern *Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-findanylist"></a><span class="api-method">Method</span> `Screen.FindAnyList`

- Signature: <span class="api-signature">`func (s Screen) FindAnyList(patterns []*Pattern) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-findbestlist"></a><span class="api-method">Method</span> `Screen.FindBestList`

- Signature: <span class="api-signature">`func (s Screen) FindBestList(patterns []*Pattern) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-findtext"></a><span class="api-method">Method</span> `Screen.FindText`

- Signature: <span class="api-signature">`func (s Screen) FindText(query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)

#### <a id="method-screen-fullregion"></a><span class="api-method">Method</span> `Screen.FullRegion`

- Signature: <span class="api-signature">`func (s Screen) FullRegion() LiveRegion`</span>
- Uses: [`LiveRegion`](#type-liveregion)

#### <a id="method-screen-has"></a><span class="api-method">Method</span> `Screen.Has`

- Signature: <span class="api-signature">`func (s Screen) Has(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-screen-hover"></a><span class="api-method">Method</span> `Screen.Hover`

- Signature: <span class="api-signature">`func (s Screen) Hover(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-keydown"></a><span class="api-method">Method</span> `Screen.KeyDown`

- Signature: <span class="api-signature">`func (s Screen) KeyDown(keys ...string) error`</span>

#### <a id="method-screen-keyup"></a><span class="api-method">Method</span> `Screen.KeyUp`

- Signature: <span class="api-signature">`func (s Screen) KeyUp(keys ...string) error`</span>

#### <a id="method-screen-live"></a><span class="api-method">Method</span> `Screen.Live`

- Signature: <span class="api-signature">`func (s Screen) Live() bool`</span>

#### <a id="method-screen-mousedown"></a><span class="api-method">Method</span> `Screen.MouseDown`

- Signature: <span class="api-signature">`func (s Screen) MouseDown(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-mouseup"></a><span class="api-method">Method</span> `Screen.MouseUp`

- Signature: <span class="api-signature">`func (s Screen) MouseUp(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-paste"></a><span class="api-method">Method</span> `Screen.Paste`

- Signature: <span class="api-signature">`func (s Screen) Paste(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-readtext"></a><span class="api-method">Method</span> `Screen.ReadText`

- Signature: <span class="api-signature">`func (s Screen) ReadText(params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams)

#### <a id="method-screen-region"></a><span class="api-method">Method</span> `Screen.Region`

- Signature: <span class="api-signature">`func (s Screen) Region(x, y, w, h int) LiveRegion`</span>
- Uses: [`LiveRegion`](#type-liveregion), [`Region`](#type-region)

#### <a id="method-screen-regionrect"></a><span class="api-method">Method</span> `Screen.RegionRect`

- Signature: <span class="api-signature">`func (s Screen) RegionRect(rect Rect) LiveRegion`</span>
- Uses: [`LiveRegion`](#type-liveregion), [`Rect`](#type-rect)

#### <a id="method-screen-rightclick"></a><span class="api-method">Method</span> `Screen.RightClick`

- Signature: <span class="api-signature">`func (s Screen) RightClick(opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-targetpoint"></a><span class="api-method">Method</span> `Screen.TargetPoint`

- Signature: <span class="api-signature">`func (s Screen) TargetPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-screen-typetext"></a><span class="api-method">Method</span> `Screen.TypeText`

- Signature: <span class="api-signature">`func (s Screen) TypeText(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-screen-wait"></a><span class="api-method">Method</span> `Screen.Wait`

- Signature: <span class="api-signature">`func (s Screen) Wait(pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-waitanylist"></a><span class="api-method">Method</span> `Screen.WaitAnyList`

- Signature: <span class="api-signature">`func (s Screen) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-waitbestlist"></a><span class="api-method">Method</span> `Screen.WaitBestList`

- Signature: <span class="api-signature">`func (s Screen) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Match`](#type-match), [`Pattern`](#type-pattern)

#### <a id="method-screen-waitvanish"></a><span class="api-method">Method</span> `Screen.WaitVanish`

- Signature: <span class="api-signature">`func (s Screen) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-screen-wheel"></a><span class="api-method">Method</span> `Screen.Wheel`

- Signature: <span class="api-signature">`func (s Screen) Wheel(direction WheelDirection, steps int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions), [`WheelDirection`](#type-wheeldirection)

## Raw Package Doc

```text
package sikuli // import "github.com/smysnk/sikuligo/pkg/sikuli"

Package sikuli provides the compatibility-facing automation API used by
sikuli-go.

The surface is intentionally aligned with common SikuliX concepts so existing
script flows can migrate with minimal rewriting:
  - Pattern and similarity tuning
  - Region scoped search and wait semantics
  - Screen level orchestration
  - Input control (click, type, hotkey)
  - OCR and observe events

Java SikuliX and sikuli-go are not byte-for-byte identical, but the exported
contracts in this package are designed to preserve the same mental model.

CONSTANTS

const (
	// DefaultSimilarity matches classic Sikuli behavior for image search.
	DefaultSimilarity = 0.70

	// ExactSimilarity is used by Pattern.Exact().
	ExactSimilarity = 0.99

	// DefaultAutoWaitTimeout is the baseline timeout for wait/find loops.
	DefaultAutoWaitTimeout = 3.0

	// DefaultWaitScanRate controls wait polling frequency.
	DefaultWaitScanRate = 3.0

	// DefaultObserveScanRate controls observe polling frequency.
	DefaultObserveScanRate = 3.0
)
const DefaultOCRLanguage = "eng"

VARIABLES

var (
	ErrFindFailed         = errors.New("sikuli: find failed")
	ErrTimeout            = errors.New("sikuli: timeout")
	ErrInvalidTarget      = errors.New("sikuli: invalid target")
	ErrBackendUnsupported = errors.New("sikuli: backend unsupported")
	ErrRuntimeUnavailable = errors.New("sikuli: live runtime unavailable")
)

FUNCTIONS

func SearchWaitVanish(probe SearchProbe, timeout, interval time.Duration) (bool, error)
    SearchWaitVanish applies the canonical vanish contract to a search probe.
    It returns true when the target is absent and false on timeout.

func SortMatchesByColumnRow(matches []Match)
    SortMatchesByColumnRow keeps parity with Java helper behavior for "by
    column".

func SortMatchesByRowColumn(matches []Match)
    SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".


TYPES

type AppAPI interface {
	Open(name string, args []string, opts AppOptions) error
	Focus(name string, opts AppOptions) error
	Close(name string, opts AppOptions) error
	IsRunning(name string, opts AppOptions) (bool, error)
	ListWindows(name string, opts AppOptions) ([]Window, error)
	FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error)
	GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error)
	FocusedWindow(name string, opts AppOptions) (Window, bool, error)
}
    AppAPI exposes lightweight app lifecycle helpers used by script flows.

type AppController struct {
	// Has unexported fields.
}

func NewAppController() *AppController

func (c *AppController) Close(name string, opts AppOptions) error

func (c *AppController) FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error)

func (c *AppController) Focus(name string, opts AppOptions) error

func (c *AppController) FocusedWindow(name string, opts AppOptions) (Window, bool, error)

func (c *AppController) GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error)

func (c *AppController) IsRunning(name string, opts AppOptions) (bool, error)

func (c *AppController) ListWindows(name string, opts AppOptions) ([]Window, error)

func (c *AppController) Open(name string, args []string, opts AppOptions) error

func (c *AppController) SetBackend(backend core.App)

type AppOptions struct {
	Timeout time.Duration
}

type Finder struct {
	// Has unexported fields.
}

func NewFinder(source *Image) (*Finder, error)
    NewFinder creates a search/OCR helper bound to a source image.

func (f *Finder) CollectLines(params OCRParams) ([]OCRLine, error)
    CollectLines runs OCR and groups word-level geometry into line-level
    results.

func (f *Finder) CollectWords(params OCRParams) ([]OCRWord, error)
    CollectWords runs OCR and returns word-level geometry for the source image.

func (f *Finder) Count(pattern *Pattern) (int, error)
    Count returns the number of matches for the given pattern.

func (f *Finder) Destroy()
    Destroy clears the compatibility iterator state and last-match cache.

func (f *Finder) Exists(pattern *Pattern) (Match, bool, error)
    Exists returns the first match when present. Missing targets return
    (Match{}, false, nil).

func (f *Finder) Find(pattern *Pattern) (Match, error)
    Find returns the best match for the pattern.

func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)
    FindAll returns all candidate matches for the pattern.

func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error)
    FindAllByColumn returns all matches sorted left-to-right then top-to-bottom.

func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error)

func (f *Finder) FindAnyList(patterns []*Pattern) ([]Match, error)

func (f *Finder) FindBestList(patterns []*Pattern) (Match, error)

func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error)
    FindText runs OCR and returns word-level matches for the query string.

func (f *Finder) Has(pattern *Pattern) (bool, error)
    Has reports whether the target exists and bubbles non-find errors.

func (f *Finder) HasNext() bool
    HasNext reports whether the compatibility iterator has another match
    available.

func (f *Finder) Iterate(pattern *Pattern) error
    Iterate prepares a compatibility iterator over the best match. Unlike Find,
    a miss does not return ErrFindFailed. Call HasNext to inspect presence.

func (f *Finder) IterateAll(pattern *Pattern) error
    IterateAll prepares a compatibility iterator over all matches. Unlike Java
    SikuliX this additive Go surface keeps LastMatches available even after
    iteration.

func (f *Finder) LastMatches() []Match
    LastMatches returns a copy of the full most recent match set. It does not
    shrink as the compatibility iterator advances.

func (f *Finder) Next() (Match, bool)
    Next returns the next compatibility-iterator match and advances the
    iterator. It returns false when the iterator is empty or exhausted.

func (f *Finder) ReadText(params OCRParams) (string, error)
    ReadText runs OCR and returns normalized text.

func (f *Finder) Reset()
    Reset rewinds the compatibility iterator to the start of the most recent
    match set.

func (f *Finder) SetMatcher(m core.Matcher)
    SetMatcher overrides the matcher backend used by this finder.

func (f *Finder) SetOCRBackend(ocr core.OCR)
    SetOCRBackend overrides the OCR backend used by this finder.

func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error)

func (f *Finder) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)

func (f *Finder) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)

func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
    WaitVanish blocks until the pattern disappears or timeout expires.

type FinderAPI interface {
	Find(pattern *Pattern) (Match, error)
	FindAll(pattern *Pattern) ([]Match, error)
	FindAllByRow(pattern *Pattern) ([]Match, error)
	FindAllByColumn(pattern *Pattern) ([]Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Iterate(pattern *Pattern) error
	IterateAll(pattern *Pattern) error
	Exists(pattern *Pattern) (Match, bool, error)
	Has(pattern *Pattern) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	HasNext() bool
	Next() (Match, bool)
	Reset()
	Destroy()
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	LastMatches() []Match
}
    FinderAPI performs match/OCR operations against a source image.
    Miss and timeout handling follows the canonical sikuli-go parity contract:
    Find returns ErrFindFailed on a miss, Exists returns (Match{}, false,
    nil), Wait returns ErrTimeout after the wait budget is exhausted, and
    WaitVanish returns false,nil on timeout. The additive compatibility iterator
    surface provides SikuliX-style traversal without removing the existing
    value/slice-oriented methods.

type Image struct {
	// Has unexported fields.
}

func NewImageFromAny(name string, src image.Image) (*Image, error)

func NewImageFromGray(name string, src *image.Gray) (*Image, error)

func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)

func (i *Image) Clone() *Image

func (i *Image) Crop(rect Rect) (*Image, error)

func (i *Image) Gray() *image.Gray

func (i *Image) Height() int

func (i *Image) Name() string

func (i *Image) Width() int

type ImageAPI interface {
	Name() string
	Width() int
	Height() int
	Gray() *image.Gray
	Clone() *Image
	Crop(rect Rect) (*Image, error)
}
    ImageAPI describes immutable image primitives used by matching and OCR. This
    aligns with the SikuliX notion of image snapshots used by Region/Finder.

type InputAPI interface {
	MoveMouse(x, y int, opts InputOptions) error
	Hover(x, y int, opts InputOptions) error
	Click(x, y int, opts InputOptions) error
	RightClick(x, y int, opts InputOptions) error
	DoubleClick(x, y int, opts InputOptions) error
	MouseDown(x, y int, opts InputOptions) error
	MouseUp(x, y int, opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	Hotkey(keys ...string) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
	Wheel(x, y int, direction WheelDirection, steps int, opts InputOptions) error
	DragDrop(fromX, fromY, toX, toY int, opts InputOptions) error
}
    InputAPI exposes desktop input actions. This is the compatibility layer for
    click/type/hotkey style operations.

type InputController struct {
	// Has unexported fields.
}

func NewInputController() *InputController

func (c *InputController) Click(x, y int, opts InputOptions) error

func (c *InputController) DoubleClick(x, y int, opts InputOptions) error

func (c *InputController) DragDrop(fromX, fromY, toX, toY int, opts InputOptions) error

func (c *InputController) Hotkey(keys ...string) error

func (c *InputController) Hover(x, y int, opts InputOptions) error

func (c *InputController) KeyDown(keys ...string) error

func (c *InputController) KeyUp(keys ...string) error

func (c *InputController) MouseDown(x, y int, opts InputOptions) error

func (c *InputController) MouseUp(x, y int, opts InputOptions) error

func (c *InputController) MoveMouse(x, y int, opts InputOptions) error

func (c *InputController) Paste(text string, opts InputOptions) error

func (c *InputController) RightClick(x, y int, opts InputOptions) error

func (c *InputController) SetBackend(backend core.Input)

func (c *InputController) TypeText(text string, opts InputOptions) error

func (c *InputController) Wheel(x, y int, direction WheelDirection, steps int, opts InputOptions) error

type InputOptions struct {
	Delay  time.Duration
	Button MouseButton
}

type LiveRegion struct {
	// Has unexported fields.
}

func (r LiveRegion) Bounds() Region

func (r LiveRegion) Capture() (*Image, error)

func (r LiveRegion) Center() Point

func (r LiveRegion) Click(opts InputOptions) error

func (r LiveRegion) CollectLines(params OCRParams) ([]OCRLine, error)

func (r LiveRegion) CollectWords(params OCRParams) ([]OCRWord, error)

func (r LiveRegion) DoubleClick(opts InputOptions) error

func (r LiveRegion) DragDrop(target TargetPointProvider, opts InputOptions) error

func (r LiveRegion) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)

func (r LiveRegion) Find(pattern *Pattern) (Match, error)

func (r LiveRegion) FindAnyList(patterns []*Pattern) ([]Match, error)

func (r LiveRegion) FindBestList(patterns []*Pattern) (Match, error)

func (r LiveRegion) FindText(query string, params OCRParams) ([]TextMatch, error)

func (r LiveRegion) Grow(dx, dy int) LiveRegion

func (r LiveRegion) Has(pattern *Pattern, timeout time.Duration) (bool, error)

func (r LiveRegion) Hover(opts InputOptions) error

func (r LiveRegion) KeyDown(keys ...string) error

func (r LiveRegion) KeyUp(keys ...string) error

func (r LiveRegion) MouseDown(opts InputOptions) error

func (r LiveRegion) MouseUp(opts InputOptions) error

func (r LiveRegion) MoveTo(x, y int) LiveRegion

func (r LiveRegion) Offset(dx, dy int) LiveRegion

func (r LiveRegion) Paste(text string, opts InputOptions) error

func (r LiveRegion) ReadText(params OCRParams) (string, error)

func (r LiveRegion) RightClick(opts InputOptions) error

func (r LiveRegion) SetSize(w, h int) LiveRegion

func (r LiveRegion) TargetPoint() Point

func (r LiveRegion) TypeText(text string, opts InputOptions) error

func (r LiveRegion) Wait(pattern *Pattern, timeout time.Duration) (Match, error)

func (r LiveRegion) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)

func (r LiveRegion) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)

func (r LiveRegion) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)

func (r LiveRegion) Wheel(direction WheelDirection, steps int, opts InputOptions) error

func (r LiveRegion) WithMatcherEngine(engine MatcherEngine) LiveRegion

type LiveRegionAPI interface {
	Bounds() Region
	Center() Point
	TargetPoint() Point
	Grow(dx, dy int) LiveRegion
	Offset(dx, dy int) LiveRegion
	MoveTo(x, y int) LiveRegion
	SetSize(w, h int) LiveRegion
	WithMatcherEngine(engine MatcherEngine) LiveRegion
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}
    LiveRegionAPI defines screen-backed region operations that route through the
    API runtime.

type Location struct {
	X int
	Y int
}

func NewLocation(x, y int) Location

func (l Location) Move(dx, dy int) Location

func (l Location) String() string

func (l Location) TargetPoint() Point

func (l Location) ToPoint() Point

type Match struct {
	Rect
	Score  float64
	Target Point
	Index  int

	// Has unexported fields.
}

func NewMatch(x, y, w, h int, score float64, off Point) Match

func SearchExists(probe SearchProbe, timeout, interval time.Duration) (Match, bool, error)
    SearchExists applies the canonical sikuli-go parity contract to a search
    probe. Misses are reported as (Match{}, false, nil); timeout <= 0 performs
    one probe.

func SearchWait(probe SearchProbe, timeout, interval time.Duration) (Match, error)
    SearchWait applies the canonical sikuli-go wait contract to a search probe.
    Misses are promoted to ErrTimeout once the wait budget is exhausted.

func (m Match) Bounds() Region
    Bounds returns the match as a region-like value so callers can treat a live
    match the same way SikuliX documents region-capable match results.

func (m Match) Capture() (*Image, error)
    Capture captures the currently matched live region.

func (m Match) Center() Point
    Center returns the geometric center of the matched rectangle.

func (m Match) Click(opts InputOptions) error
    Click clicks the match target point.

func (m Match) CollectLines(params OCRParams) ([]OCRLine, error)
    CollectLines returns OCR line-level results inside the matched live region.

func (m Match) CollectWords(params OCRParams) ([]OCRWord, error)
    CollectWords returns OCR word-level results inside the matched live region.

func (m Match) DoubleClick(opts InputOptions) error
    DoubleClick performs two click actions against the match target point.

func (m Match) DragDrop(target TargetPointProvider, opts InputOptions) error
    DragDrop drags from the match target point to the target point.

func (m Match) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
    Exists probes within the matched live region.

func (m Match) Find(pattern *Pattern) (Match, error)
    Find searches again within the matched live region.

func (m Match) FindAnyList(patterns []*Pattern) ([]Match, error)

func (m Match) FindBestList(patterns []*Pattern) (Match, error)

func (m Match) FindText(query string, params OCRParams) ([]TextMatch, error)
    FindText searches OCR text inside the matched live region.

func (m Match) Has(pattern *Pattern, timeout time.Duration) (bool, error)
    Has reports whether the target exists within the matched live region.

func (m Match) Hover(opts InputOptions) error
    Hover is a parity-friendly alias for MoveMouse.

func (m Match) KeyDown(keys ...string) error
    KeyDown holds the provided keys.

func (m Match) KeyUp(keys ...string) error
    KeyUp releases the provided keys.

func (m Match) Live() bool
    Live reports whether this match is bound to a running sikuli-go runtime.

func (m Match) MouseDown(opts InputOptions) error
    MouseDown presses and holds the mouse button at the match target point.

func (m Match) MouseUp(opts InputOptions) error
    MouseUp releases the mouse button at the match target point.

func (m Match) MoveMouse(opts InputOptions) error
    MoveMouse moves the pointer to the match target point.

func (m Match) Paste(text string, opts InputOptions) error
    Paste focuses the match target point and pastes text into it.

func (m Match) ReadText(params OCRParams) (string, error)
    ReadText runs OCR inside the matched live region.

func (m Match) Region() Region
    Region is an alias for Bounds for parity-friendly call sites.

func (m Match) RightClick(opts InputOptions) error
    RightClick clicks the match target point with the right mouse button.

func (m Match) String() string

func (m Match) TargetPoint() Point
    TargetPoint returns the resolved click target point for this match.

func (m Match) TypeText(text string, opts InputOptions) error
    TypeText focuses the match target point and types text into it.

func (m Match) Wait(pattern *Pattern, timeout time.Duration) (Match, error)
    Wait waits for the target within the matched live region.

func (m Match) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)

func (m Match) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)

func (m Match) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
    WaitVanish waits for the target to disappear from the matched live region.

func (m Match) Wheel(direction WheelDirection, steps int, opts InputOptions) error
    Wheel scrolls at the match target point.

type MatchAPI interface {
	Bounds() Region
	Region() Region
	Center() Point
	TargetPoint() Point
	Live() bool
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	MoveMouse(opts InputOptions) error
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}
    MatchAPI defines a live match result that can be used directly for follow-up
    search, capture, and current direct input verbs without manual point
    routing.

type MatcherEngine string
    MatcherEngine selects the server-side matcher implementation used for live
    screen queries.

const (
	MatcherEngineDefault  MatcherEngine = ""
	MatcherEngineTemplate MatcherEngine = "template"
	MatcherEngineORB      MatcherEngine = "orb"
	MatcherEngineHybrid   MatcherEngine = "hybrid"
	MatcherEngineAKAZE    MatcherEngine = "akaze"
	MatcherEngineBRISK    MatcherEngine = "brisk"
	MatcherEngineKAZE     MatcherEngine = "kaze"
	MatcherEngineSIFT     MatcherEngine = "sift"
)
type MouseButton string

const (
	MouseButtonLeft   MouseButton = "left"
	MouseButtonRight  MouseButton = "right"
	MouseButtonMiddle MouseButton = "middle"
)
type OCRLine struct {
	Rect
	Text       string
	Confidence float64
	Index      int
	Words      []OCRWord
}

type OCRParams struct {
	Language         string
	TrainingDataPath string
	MinConfidence    float64
	Timeout          time.Duration
	CaseSensitive    bool
}

type OCRWord struct {
	Rect
	Text       string
	Confidence float64
	Index      int
}

type ObserveAPI interface {
	ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)
}
    ObserveAPI exposes appear/vanish/change polling contracts for a region.

type ObserveEvent struct {
	Type      ObserveEventType
	Match     Match
	Timestamp time.Time
}

type ObserveEventType string

const (
	ObserveEventAppear ObserveEventType = "appear"
	ObserveEventVanish ObserveEventType = "vanish"
	ObserveEventChange ObserveEventType = "change"
)
type ObserveOptions struct {
	Interval time.Duration
	Timeout  time.Duration
}

type ObserverController struct {
	// Has unexported fields.
}

func NewObserverController() *ObserverController

func (c *ObserverController) ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) SetBackend(backend core.Observer)

type Offset struct {
	X int
	Y int
}

func NewOffset(x, y int) Offset

func (o Offset) String() string

func (o Offset) ToPoint() Point

type Options struct {
	// Has unexported fields.
}

func NewOptions() *Options

func NewOptionsFromMap(entries map[string]string) *Options

func (o *Options) Clone() *Options

func (o *Options) Delete(key string)

func (o *Options) Entries() map[string]string

func (o *Options) GetBool(key string, def bool) bool

func (o *Options) GetFloat64(key string, def float64) float64

func (o *Options) GetInt(key string, def int) int

func (o *Options) GetString(key, def string) string

func (o *Options) Has(key string) bool

func (o *Options) Merge(other *Options)

func (o *Options) SetBool(key string, value bool)

func (o *Options) SetFloat64(key string, value float64)

func (o *Options) SetInt(key string, value int)

func (o *Options) SetString(key, value string)

type Pattern struct {
	// Has unexported fields.
}

func NewPattern(img *Image) (*Pattern, error)
    NewPattern creates a match pattern from an image with default similarity
    settings.

func (p *Pattern) Exact() *Pattern
    Exact is a convenience for Similar(1.0).

func (p *Pattern) Image() *Image
    Image returns the underlying pattern image.

func (p *Pattern) Mask() *image.Gray
    Mask returns the currently configured mask.

func (p *Pattern) Offset() Point
    Offset returns the configured click anchor offset.

func (p *Pattern) Resize(factor float64) *Pattern
    Resize scales the pattern before matching.

func (p *Pattern) ResizeFactor() float64
    ResizeFactor returns the currently configured resize factor.

func (p *Pattern) Similar(sim float64) *Pattern
    Similar sets the acceptance threshold in the [0,1] range. Higher values
    require a closer match.

func (p *Pattern) Similarity() float64
    Similarity returns the current acceptance threshold.

func (p *Pattern) TargetOffset(dx, dy int) *Pattern
    TargetOffset sets the click anchor relative to the matched rectangle.

func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error)
    WithMask sets an optional per-pixel mask where 0 excludes and >0 includes
    pixels.

func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error)
    WithMaskMatrix sets an optional binary mask from matrix rows.

type PatternAPI interface {
	Image() *Image
	Similar(sim float64) *Pattern
	Similarity() float64
	Exact() *Pattern
	TargetOffset(dx, dy int) *Pattern
	Offset() Point
	Resize(factor float64) *Pattern
	ResizeFactor() float64
	Mask() *image.Gray
}
    PatternAPI configures how a target image should be matched on screen.
    It mirrors SikuliX Pattern behavior such as similar(), exact(), and
    targetOffset().

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point

func (p Point) TargetPoint() Point

func (p Point) ToLocation() Location
    ToLocation converts a point to a parity-friendly Location value.

func (p Point) ToOffset() Offset
    ToOffset converts a point to a parity-friendly Offset value.

type Rect struct {
	X int
	Y int
	W int
	H int
}

func NewRect(x, y, w, h int) Rect

func (r Rect) Contains(p Point) bool

func (r Rect) Empty() bool

func (r Rect) String() string

type Region struct {
	Rect
	// ThrowException is retained as parity metadata for SikuliX-style ports.
	// The Go API uses explicit return values for misses and timeouts regardless of this flag.
	ThrowException  bool
	AutoWaitTimeout float64
	WaitScanRate    float64
	ObserveScanRate float64
}

func NewRegion(x, y, w, h int) Region
    NewRegion constructs a rectangular search area with default timing settings.

func (r Region) Center() Point
    Center returns the midpoint of the region.

func (r Region) CollectLines(source *Image, params OCRParams) ([]OCRLine, error)
    CollectLines runs OCR in region and returns line-level results.

func (r Region) CollectWords(source *Image, params OCRParams) ([]OCRWord, error)
    CollectWords runs OCR in region and returns word-level results.

func (r Region) Contains(p Point) bool
    Contains reports whether a point lies within the region.

func (r Region) ContainsLocation(loc Location) bool
    ContainsLocation reports whether this region contains the given location.

func (r Region) ContainsRegion(other Region) bool

func (r Region) Count(source *Image, pattern *Pattern) (int, error)
    Count returns the number of matches found in this region.

func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)
    Exists checks for pattern presence within timeout and returns first match
    when found.

func (r Region) Find(source *Image, pattern *Pattern) (Match, error)

func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error)
    FindAll returns all matches in this region.

func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)

func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)

func (r Region) FindAnyList(source *Image, patterns []*Pattern) ([]Match, error)

func (r Region) FindBestList(source *Image, patterns []*Pattern) (Match, error)

func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)
    FindText runs OCR in region and returns matches for the query.

func (r Region) Grow(dx, dy int) Region
    Grow expands the region outward in both directions.

func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)

func (r Region) Intersection(other Region) Region
    Intersection returns the overlap between this region and another.

func (r Region) MoveTo(x, y int) Region

func (r Region) MoveToLocation(loc Location) Region
    MoveToLocation moves this region using a Location alias.

func (r Region) Offset(dx, dy int) Region
    Offset translates the region by dx and dy.

func (r Region) OffsetBy(off Offset) Region
    OffsetBy applies an Offset alias to this region position.

func (r Region) ReadText(source *Image, params OCRParams) (string, error)

func (r *Region) ResetThrowException()

func (r *Region) SetAutoWaitTimeout(sec float64)

func (r *Region) SetObserveScanRate(rate float64)

func (r Region) SetSize(w, h int) Region
    SetSize updates width and height while clamping negatives to zero.

func (r *Region) SetThrowException(flag bool)

func (r *Region) SetWaitScanRate(rate float64)

func (r Region) TargetPoint() Point

func (r Region) Union(other Region) Region
    Union returns the smallest region containing both regions.

func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)

func (r Region) WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error)

func (r Region) WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error)

func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
    WaitVanish waits until pattern disappears or timeout expires.

type RegionAPI interface {
	Center() Point
	Grow(dx, dy int) Region
	Offset(dx, dy int) Region
	MoveTo(x, y int) Region
	SetSize(w, h int) Region
	Contains(p Point) bool
	ContainsRegion(other Region) bool
	Union(other Region) Region
	Intersection(other Region) Region
	Find(source *Image, pattern *Pattern) (Match, error)
	Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)
	WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
	FindAll(source *Image, pattern *Pattern) ([]Match, error)
	FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)
	FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)
	FindAnyList(source *Image, patterns []*Pattern) ([]Match, error)
	FindBestList(source *Image, patterns []*Pattern) (Match, error)
	WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error)
	ReadText(source *Image, params OCRParams) (string, error)
	FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)
	CollectWords(source *Image, params OCRParams) ([]OCRWord, error)
	CollectLines(source *Image, params OCRParams) ([]OCRLine, error)
}
    RegionAPI defines region geometry and region-scoped automation operations.
    It maps to familiar SikuliX Region methods while keeping the same explicit
    Go miss/timeout contract used by FinderAPI.

type Runtime struct {
	// Has unexported fields.
}

func NewRuntime(address string, opts ...RuntimeOption) (*Runtime, error)
    NewRuntime connects to a running sikuli-go API runtime and exposes live
    screen operations.

func (r *Runtime) Address() string

func (r *Runtime) Capture() (*Image, error)

func (r *Runtime) CaptureRegion(region Region) (*Image, error)

func (r *Runtime) Close() error

func (r *Runtime) PrimaryScreen() (Screen, error)

func (r *Runtime) Region(region Region) LiveRegion

func (r *Runtime) Screen(id int) (Screen, error)

func (r *Runtime) Screens() ([]Screen, error)

type RuntimeAPI interface {
	Address() string
	Close() error
	Screens() ([]Screen, error)
	PrimaryScreen() (Screen, error)
	Screen(id int) (Screen, error)
	Capture() (*Image, error)
	CaptureRegion(region Region) (*Image, error)
	Region(region Region) LiveRegion
}
    RuntimeAPI defines the public live runtime client used to discover screens
    and capture/search them.

type RuntimeOption func(*runtimeConfig)

func WithRuntimeAuthToken(token string) RuntimeOption

func WithRuntimeConn(conn *grpc.ClientConn) RuntimeOption

func WithRuntimeContextDialer(dialer func(context.Context, string) (net.Conn, error)) RuntimeOption

func WithRuntimeDialTimeout(timeout time.Duration) RuntimeOption

func WithRuntimeMatcherEngine(engine MatcherEngine) RuntimeOption

func WithRuntimeRPCTimeout(timeout time.Duration) RuntimeOption

type RuntimeSettings struct {
	ImageCache      int
	ShowActions     bool
	WaitScanRate    float64
	ObserveScanRate float64
	AutoWaitTimeout float64
	MinSimilarity   float64
	// FindFailedThrows is retained as parity metadata for SikuliX-style ports.
	// The Go API uses explicit return values for misses and timeouts regardless of this flag.
	FindFailedThrows bool
}

func GetSettings() RuntimeSettings

func ResetSettings() RuntimeSettings

func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings

type Screen struct {
	ID      int
	Name    string
	Bounds  Rect
	Primary bool
	// Has unexported fields.
}

func NewScreen(id int, bounds Rect) Screen
    NewScreen constructs a logical screen descriptor.

func (s Screen) Capture() (*Image, error)

func (s Screen) Click(opts InputOptions) error

func (s Screen) CollectLines(params OCRParams) ([]OCRLine, error)

func (s Screen) CollectWords(params OCRParams) ([]OCRWord, error)

func (s Screen) DoubleClick(opts InputOptions) error

func (s Screen) DragDrop(target TargetPointProvider, opts InputOptions) error

func (s Screen) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)

func (s Screen) Find(pattern *Pattern) (Match, error)

func (s Screen) FindAnyList(patterns []*Pattern) ([]Match, error)

func (s Screen) FindBestList(patterns []*Pattern) (Match, error)

func (s Screen) FindText(query string, params OCRParams) ([]TextMatch, error)

func (s Screen) FullRegion() LiveRegion

func (s Screen) Has(pattern *Pattern, timeout time.Duration) (bool, error)

func (s Screen) Hover(opts InputOptions) error

func (s Screen) KeyDown(keys ...string) error

func (s Screen) KeyUp(keys ...string) error

func (s Screen) Live() bool

func (s Screen) MouseDown(opts InputOptions) error

func (s Screen) MouseUp(opts InputOptions) error

func (s Screen) Paste(text string, opts InputOptions) error

func (s Screen) ReadText(params OCRParams) (string, error)

func (s Screen) Region(x, y, w, h int) LiveRegion

func (s Screen) RegionRect(rect Rect) LiveRegion

func (s Screen) RightClick(opts InputOptions) error

func (s Screen) TargetPoint() Point

func (s Screen) TypeText(text string, opts InputOptions) error

func (s Screen) Wait(pattern *Pattern, timeout time.Duration) (Match, error)

func (s Screen) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)

func (s Screen) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)

func (s Screen) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)

func (s Screen) Wheel(direction WheelDirection, steps int, opts InputOptions) error

type ScreenAPI interface {
	Live() bool
	TargetPoint() Point
	FullRegion() LiveRegion
	Region(x, y, w, h int) LiveRegion
	RegionRect(rect Rect) LiveRegion
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}
    ScreenAPI defines live screen selection, capture, and search operations.

type SearchProbe func() (Match, error)
    SearchProbe returns the next match candidate for a parity search operation.
    Returning ErrFindFailed indicates a miss; any other error aborts the search.

type TargetPointProvider interface {
	TargetPoint() Point
}
    TargetPointProvider describes values that can resolve to a concrete action
    target.

type TextMatch struct {
	Rect
	Text       string
	Confidence float64
	Index      int
}

type WheelDirection string

const (
	WheelDirectionUp    WheelDirection = "up"
	WheelDirectionDown  WheelDirection = "down"
	WheelDirectionLeft  WheelDirection = "left"
	WheelDirectionRight WheelDirection = "right"
)
type Window struct {
	ID      string
	App     string
	PID     int
	Title   string
	Bounds  Rect
	Focused bool
}

type WindowQuery struct {
	ID            string
	TitleExact    string
	TitleContains string
	FocusedOnly   bool
	Index         int
}

```
