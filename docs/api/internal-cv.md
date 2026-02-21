# API: `internal/cv`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package cv // import "github.com/sikulix/portgo/internal/cv"`

## Symbol Index

### Types

- <span class="api-type">[`NCCMatcher`](#type-nccmatcher)</span>
- <span class="api-type">[`SADMatcher`](#type-sadmatcher)</span>

### Functions

- <span class="api-func">[`NewNCCMatcher`](#func-newnccmatcher)</span>
- <span class="api-func">[`NewSADMatcher`](#func-newsadmatcher)</span>

### Methods

- <span class="api-method">[`NCCMatcher.Find`](#method-nccmatcher-find)</span>
- <span class="api-method">[`SADMatcher.Find`](#method-sadmatcher-find)</span>

## Declarations

### Types

#### <a id="type-nccmatcher"></a><span class="api-type">Type</span> `NCCMatcher`

- Signature: <span class="api-signature">`type NCCMatcher struct{}`</span>

#### <a id="type-sadmatcher"></a><span class="api-type">Type</span> `SADMatcher`

- Signature: <span class="api-signature">`type SADMatcher struct{}`</span>

### Functions

#### <a id="func-newnccmatcher"></a><span class="api-func">Function</span> `NewNCCMatcher`

- Signature: <span class="api-signature">`func NewNCCMatcher() *NCCMatcher`</span>
- Uses: [`NCCMatcher`](#type-nccmatcher)

#### <a id="func-newsadmatcher"></a><span class="api-func">Function</span> `NewSADMatcher`

- Signature: <span class="api-signature">`func NewSADMatcher() *SADMatcher`</span>
- Uses: [`SADMatcher`](#type-sadmatcher)

### Methods

#### <a id="method-nccmatcher-find"></a><span class="api-method">Method</span> `NCCMatcher.Find`

- Signature: <span class="api-signature">`func (m *NCCMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

#### <a id="method-sadmatcher-find"></a><span class="api-method">Method</span> `SADMatcher.Find`

- Signature: <span class="api-signature">`func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)`</span>

## Raw Package Doc

```text
package cv // import "github.com/sikulix/portgo/internal/cv"


TYPES

type NCCMatcher struct{}

func NewNCCMatcher() *NCCMatcher

func (m *NCCMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

type SADMatcher struct{}

func NewSADMatcher() *SADMatcher

func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error)

```
