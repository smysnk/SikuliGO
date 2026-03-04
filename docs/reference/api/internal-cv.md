# API: `internal/cv`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package cv // import "github.com/smysnk/sikuligo/internal/cv"`

## Symbol Index

### Types

- <span class="api-type">[`AKAZEMatcher`](#type-akazematcher)</span>
- <span class="api-type">[`BRISKMatcher`](#type-briskmatcher)</span>
- <span class="api-type">[`HybridMatcher`](#type-hybridmatcher)</span>
- <span class="api-type">[`KAZEMatcher`](#type-kazematcher)</span>
- <span class="api-type">[`MatcherEngine`](#type-matcherengine)</span>
- <span class="api-type">[`NCCMatcher`](#type-nccmatcher)</span>
- <span class="api-type">[`ORBMatcher`](#type-orbmatcher)</span>
- <span class="api-type">[`OpenCVMatcher`](#type-opencvmatcher)</span>
- <span class="api-type">[`SADMatcher`](#type-sadmatcher)</span>
- <span class="api-type">[`SIFTMatcher`](#type-siftmatcher)</span>

### Functions

- <span class="api-func">[`NewDefaultMatcher`](#func-newdefaultmatcher)</span>
- <span class="api-func">[`NewMatcherForEngine`](#func-newmatcherforengine)</span>
- <span class="api-func">[`OpenCVEnabled`](#func-opencvenabled)</span>
- <span class="api-func">[`NewAKAZEMatcher`](#func-newakazematcher)</span>
- <span class="api-func">[`NewBRISKMatcher`](#func-newbriskmatcher)</span>
- <span class="api-func">[`NewHybridMatcher`](#func-newhybridmatcher)</span>
- <span class="api-func">[`NewKAZEMatcher`](#func-newkazematcher)</span>
- <span class="api-func">[`ParseMatcherEngine`](#func-parsematcherengine)</span>
- <span class="api-func">[`NewNCCMatcher`](#func-newnccmatcher)</span>
- <span class="api-func">[`NewORBMatcher`](#func-neworbmatcher)</span>
- <span class="api-func">[`NewOpenCVMatcher`](#func-newopencvmatcher)</span>
- <span class="api-func">[`NewSADMatcher`](#func-newsadmatcher)</span>
- <span class="api-func">[`NewSIFTMatcher`](#func-newsiftmatcher)</span>

### Methods

- <span class="api-method">[`AKAZEMatcher.Find`](#method-akazematcher-find)</span>
- <span class="api-method">[`BRISKMatcher.Find`](#method-briskmatcher-find)</span>
- <span class="api-method">[`HybridMatcher.Find`](#method-hybridmatcher-find)</span>
- <span class="api-method">[`KAZEMatcher.Find`](#method-kazematcher-find)</span>
- <span class="api-method">[`NCCMatcher.Find`](#method-nccmatcher-find)</span>
- <span class="api-method">[`ORBMatcher.Find`](#method-orbmatcher-find)</span>
- <span class="api-method">[`OpenCVMatcher.Find`](#method-opencvmatcher-find)</span>
- <span class="api-method">[`SADMatcher.Find`](#method-sadmatcher-find)</span>
- <span class="api-method">[`SIFTMatcher.Find`](#method-siftmatcher-find)</span>

## Declarations

### Types

#### <a id="type-akazematcher"></a><span class="api-type">Type</span> `AKAZEMatcher`

- Signature: <span class="api-signature">`type AKAZEMatcher struct{}`</span>

#### <a id="type-briskmatcher"></a><span class="api-type">Type</span> `BRISKMatcher`

- Signature: <span class="api-signature">`type BRISKMatcher struct{}`</span>

#### <a id="type-hybridmatcher"></a><span class="api-type">Type</span> `HybridMatcher`

- Signature: <span class="api-signature">`type HybridMatcher struct {`</span>

#### <a id="type-kazematcher"></a><span class="api-type">Type</span> `KAZEMatcher`

- Signature: <span class="api-signature">`type KAZEMatcher struct{}`</span>

#### <a id="type-matcherengine"></a><span class="api-type">Type</span> `MatcherEngine`

- Signature: <span class="api-signature">`type MatcherEngine string`</span>

#### <a id="type-nccmatcher"></a><span class="api-type">Type</span> `NCCMatcher`

- Signature: <span class="api-signature">`type NCCMatcher struct{}`</span>

#### <a id="type-orbmatcher"></a><span class="api-type">Type</span> `ORBMatcher`

- Signature: <span class="api-signature">`type ORBMatcher struct{}`</span>

#### <a id="type-opencvmatcher"></a><span class="api-type">Type</span> `OpenCVMatcher`

- Signature: <span class="api-signature">`type OpenCVMatcher struct{}`</span>

#### <a id="type-sadmatcher"></a><span class="api-type">Type</span> `SADMatcher`

- Signature: <span class="api-signature">`type SADMatcher struct{}`</span>

#### <a id="type-siftmatcher"></a><span class="api-type">Type</span> `SIFTMatcher`

- Signature: <span class="api-signature">`type SIFTMatcher struct{}`</span>

### Functions

#### <a id="func-newdefaultmatcher"></a><span class="api-func">Function</span> `NewDefaultMatcher`

- Signature: <span class="api-signature">`func NewDefaultMatcher() core.Matcher`</span>
- Notes: NewDefaultMatcher returns the matcher backend used by default in Sikuli flows.

#### <a id="func-newmatcherforengine"></a><span class="api-func">Function</span> `NewMatcherForEngine`

- Signature: <span class="api-signature">`func NewMatcherForEngine(engine MatcherEngine) (core.Matcher, error)`</span>
- Uses: [`MatcherEngine`](#type-matcherengine)

#### <a id="func-opencvenabled"></a><span class="api-func">Function</span> `OpenCVEnabled`

- Signature: <span class="api-signature">`func OpenCVEnabled() bool`</span>
- Notes: OpenCVEnabled reports whether this binary was built with the "opencv" build tag.

#### <a id="func-newakazematcher"></a><span class="api-func">Function</span> `NewAKAZEMatcher`

- Signature: <span class="api-signature">`func NewAKAZEMatcher() *AKAZEMatcher`</span>
- Uses: [`AKAZEMatcher`](#type-akazematcher)

#### <a id="func-newbriskmatcher"></a><span class="api-func">Function</span> `NewBRISKMatcher`

- Signature: <span class="api-signature">`func NewBRISKMatcher() *BRISKMatcher`</span>
- Uses: [`BRISKMatcher`](#type-briskmatcher)

#### <a id="func-newhybridmatcher"></a><span class="api-func">Function</span> `NewHybridMatcher`

- Signature: <span class="api-signature">`func NewHybridMatcher(primary, fallback core.Matcher) *HybridMatcher`</span>
- Uses: [`HybridMatcher`](#type-hybridmatcher)

#### <a id="func-newkazematcher"></a><span class="api-func">Function</span> `NewKAZEMatcher`

- Signature: <span class="api-signature">`func NewKAZEMatcher() *KAZEMatcher`</span>
- Uses: [`KAZEMatcher`](#type-kazematcher)

#### <a id="func-parsematcherengine"></a><span class="api-func">Function</span> `ParseMatcherEngine`

- Signature: <span class="api-signature">`func ParseMatcherEngine(raw string) (MatcherEngine, error)`</span>
- Uses: [`MatcherEngine`](#type-matcherengine)

#### <a id="func-newnccmatcher"></a><span class="api-func">Function</span> `NewNCCMatcher`

- Signature: <span class="api-signature">`func NewNCCMatcher() *NCCMatcher`</span>
- Uses: [`NCCMatcher`](#type-nccmatcher)

#### <a id="func-neworbmatcher"></a><span class="api-func">Function</span> `NewORBMatcher`

- Signature: <span class="api-signature">`func NewORBMatcher() *ORBMatcher`</span>
- Uses: [`ORBMatcher`](#type-orbmatcher)

#### <a id="func-newopencvmatcher"></a><span class="api-func">Function</span> `NewOpenCVMatcher`

- Signature: <span class="api-signature">`func NewOpenCVMatcher() *OpenCVMatcher`</span>
- Uses: [`OpenCVMatcher`](#type-opencvmatcher)

#### <a id="func-newsadmatcher"></a><span class="api-func">Function</span> `NewSADMatcher`

- Signature: <span class="api-signature">`func NewSADMatcher() *SADMatcher`</span>
- Uses: [`SADMatcher`](#type-sadmatcher)

#### <a id="func-newsiftmatcher"></a><span class="api-func">Function</span> `NewSIFTMatcher`

- Signature: <span class="api-signature">`func NewSIFTMatcher() *SIFTMatcher`</span>
- Uses: [`SIFTMatcher`](#type-siftmatcher)

### Methods

#### <a id="method-akazematcher-find"></a><span class="api-method">Method</span> `AKAZEMatcher.Find`

- Signature: <span class="api-signature">`func (m *AKAZEMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-briskmatcher-find"></a><span class="api-method">Method</span> `BRISKMatcher.Find`

- Signature: <span class="api-signature">`func (m *BRISKMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-hybridmatcher-find"></a><span class="api-method">Method</span> `HybridMatcher.Find`

- Signature: <span class="api-signature">`func (m *HybridMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-kazematcher-find"></a><span class="api-method">Method</span> `KAZEMatcher.Find`

- Signature: <span class="api-signature">`func (m *KAZEMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-nccmatcher-find"></a><span class="api-method">Method</span> `NCCMatcher.Find`

- Signature: <span class="api-signature">`func (m *NCCMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-orbmatcher-find"></a><span class="api-method">Method</span> `ORBMatcher.Find`

- Signature: <span class="api-signature">`func (m *ORBMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-opencvmatcher-find"></a><span class="api-method">Method</span> `OpenCVMatcher.Find`

- Signature: <span class="api-signature">`func (m *OpenCVMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-sadmatcher-find"></a><span class="api-method">Method</span> `SADMatcher.Find`

- Signature: <span class="api-signature">`func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-siftmatcher-find"></a><span class="api-method">Method</span> `SIFTMatcher.Find`

- Signature: <span class="api-signature">`func (m *SIFTMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

## Raw Package Doc

```text
package cv // import "github.com/smysnk/sikuligo/internal/cv"


FUNCTIONS

func NewDefaultMatcher() core.Matcher
    NewDefaultMatcher returns the matcher backend used by default in Sikuli
    flows.

func NewMatcherForEngine(engine MatcherEngine) (core.Matcher, error)
func OpenCVEnabled() bool
    OpenCVEnabled reports whether this binary was built with the "opencv" build
    tag.


TYPES

type AKAZEMatcher struct{}

func NewAKAZEMatcher() *AKAZEMatcher

func (m *AKAZEMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type BRISKMatcher struct{}

func NewBRISKMatcher() *BRISKMatcher

func (m *BRISKMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type HybridMatcher struct {
	// Has unexported fields.
}

func NewHybridMatcher(primary, fallback core.Matcher) *HybridMatcher

func (m *HybridMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type KAZEMatcher struct{}

func NewKAZEMatcher() *KAZEMatcher

func (m *KAZEMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type MatcherEngine string

const (
	MatcherEngineTemplate MatcherEngine = "template"
	MatcherEngineORB      MatcherEngine = "orb"
	MatcherEngineAKAZE    MatcherEngine = "akaze"
	MatcherEngineBRISK    MatcherEngine = "brisk"
	MatcherEngineKAZE     MatcherEngine = "kaze"
	MatcherEngineSIFT     MatcherEngine = "sift"
	MatcherEngineHybrid   MatcherEngine = "hybrid"
)
func ParseMatcherEngine(raw string) (MatcherEngine, error)

type NCCMatcher struct{}

func NewNCCMatcher() *NCCMatcher

func (m *NCCMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type ORBMatcher struct{}

func NewORBMatcher() *ORBMatcher

func (m *ORBMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type OpenCVMatcher struct{}

func NewOpenCVMatcher() *OpenCVMatcher

func (m *OpenCVMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type SADMatcher struct{}

func NewSADMatcher() *SADMatcher

func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type SIFTMatcher struct{}

func NewSIFTMatcher() *SIFTMatcher

func (m *SIFTMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

```
