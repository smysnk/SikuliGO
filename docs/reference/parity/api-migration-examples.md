# API Migration Examples

Focused migration examples for the parity areas closed in API phases 1-8.

## Search Semantics and Exception/Null Behavior

### Image-Backed Flow

```go
source := preparedImage
pattern := preparedPattern
region := sikuli.NewRegion(0, 0, source.Width(), source.Height())

match, ok, err := region.Exists(source, pattern, 500*time.Millisecond)
if err != nil {
    return err
}
if !ok {
    return nil
}

_, err = region.Wait(source, pattern, 2*time.Second)
if errors.Is(err, sikuli.ErrTimeout) {
    // explicit timeout handling
}
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
pattern := preparedPattern

match, ok, err := screen.Exists(pattern, 500*time.Millisecond)
if err != nil {
    return err
}
if !ok {
    return nil
}

_, err = screen.Wait(pattern, 2*time.Second)
if errors.Is(err, sikuli.ErrTimeout) {
    // explicit timeout handling
}
```

## Live Screen and Region Surface

### Image-Backed Flow

```go
source := preparedImage
region := sikuli.NewRegion(100, 80, 400, 300)
snippet, _ := source.Crop(sikuli.NewRect(region.X, region.Y, region.W, region.H))
_ = snippet
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screens, _ := runtime.Screens()
primary, _ := runtime.PrimaryScreen()

capture, _ := primary.Capture()
regionCapture, _ := primary.Region(100, 80, 400, 300).Capture()

_ = screens
_ = capture
_ = regionCapture
```

## Match as a First-Class Action Target

### Image-Backed Flow

```go
source := preparedImage
pattern := preparedPattern
finder, _ := sikuli.NewFinder(source)

match, _ := finder.Find(pattern)
target := match.TargetPoint()
bounds := match.Bounds()

_ = target
_ = bounds
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
pattern := preparedPattern

match, _ := screen.Find(pattern)
matchImage, _ := match.Capture()
_ = match.Click(sikuli.InputOptions{})

_ = matchImage
```

## Direct Action Surface

### Image-Backed Flow

```go
source := preparedImage
pattern := preparedPattern
finder, _ := sikuli.NewFinder(source)
match, _ := finder.Find(pattern)

input := sikuli.NewInputController()
target := match.TargetPoint()
_ = input.Click(target.X, target.Y, sikuli.InputOptions{})
_ = input.Paste("hello", sikuli.InputOptions{})
_ = input.KeyDown("cmd", "shift")
_ = input.KeyUp("cmd", "shift")
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
region := screen.Region(100, 80, 400, 300)
pattern, _ := sikuli.NewPatternFromPath("send-button.png")

match, _ := region.Find(pattern)
_ = match.DoubleClick(sikuli.InputOptions{})
_ = match.Paste("hello", sikuli.InputOptions{})
_ = region.Wheel(sikuli.WheelDirectionDown, 3, sikuli.InputOptions{})
_ = region.DragDrop(match, sikuli.InputOptions{})
```

## Finder Traversal and Lifecycle

### Image-Backed Flow

```go
source := preparedImage
pattern := preparedPattern
finder, _ := sikuli.NewFinder(source)

_ = finder.IterateAll(pattern)
for finder.HasNext() {
    match, _ := finder.Next()
    _ = match
}
finder.Destroy()
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
pattern := preparedPattern

capture, _ := screen.Region(0, 0, 800, 200).Capture()
finder, _ := sikuli.NewFinder(capture)
_ = finder.IterateAll(pattern)
for finder.HasNext() {
    match, _ := finder.Next()
    _ = match
}
```

## Multi-Target Search Helpers

### Image-Backed Flow

```go
source := preparedImage
region := sikuli.NewRegion(0, 0, source.Width(), source.Height())
patterns := preparedPatterns

matches, _ := region.FindAnyList(source, patterns)
best, _ := region.FindBestList(source, patterns)

_ = matches
_ = best
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
patterns := preparedPatterns

matches, _ := screen.FindAnyList(patterns)
best, _ := screen.WaitBestList(patterns, 2*time.Second)

_ = matches
_ = best
```

## OCR Collection Surface

### Image-Backed Flow

```go
source := preparedImage
finder, _ := sikuli.NewFinder(source)

words, _ := finder.CollectWords(sikuli.OCRParams{Language: "eng"})
lines, _ := finder.CollectLines(sikuli.OCRParams{Language: "eng"})

_ = words
_ = lines
```

### Live-Screen Flow

```go
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
screen, _ := runtime.PrimaryScreen()
region := screen.Region(100, 120, 600, 300)

words, _ := region.CollectWords(sikuli.OCRParams{Language: "eng"})
lines, _ := region.CollectLines(sikuli.OCRParams{Language: "eng"})

_ = words
_ = lines
```

## App and Window Surface

### Image-Backed Flow

Window discovery has no pure image-backed analogue. The migration shape is: resolve the runtime window first, capture its bounds once, then run the same image-backed search/OCR helpers against that snapshot.

```go
apps := sikuli.NewAppController()
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
window, ok, _ := apps.FocusedWindow("Mail", sikuli.AppOptions{})
if !ok {
    return nil
}

snapshot, _ := runtime.CaptureRegion(window.Bounds)
finder, _ := sikuli.NewFinder(snapshot)
match, _ := finder.Find(sendPattern)

_ = match
```

### Live-Screen Flow

```go
apps := sikuli.NewAppController()
runtime, _ := sikuli.NewRuntime("127.0.0.1:50051")
window, ok, _ := apps.GetWindow("Mail", sikuli.WindowQuery{TitleContains: "Compose"}, sikuli.AppOptions{})
if !ok {
    return nil
}

screen, _ := runtime.PrimaryScreen()
composeRegion := screen.RegionRect(window.Bounds)
match, _ := composeRegion.Find(sendPattern)

_ = match.Click(sikuli.InputOptions{})
```
